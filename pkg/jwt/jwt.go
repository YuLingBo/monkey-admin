package jwt

import (
	"errors"
	"monkey-admin/config"
	"monkey-admin/dao"
	"monkey-admin/models/response"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//调用过滤去将放行的请求先放行
		if doSquare(c) {
			return
		}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			gotool.Logs.InfoLog().Printf("请求未携带token，无权限访问 111")
			c.JSON(http.StatusOK, gin.H{
				"status": 401,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		s := strings.Split(token, " ")
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(s[1])
		if err != nil {
			if err == TokenExpired {
				// token过期，尝试刷新
				newToken, refreshErr := j.RefreshToken(s[1])
				if refreshErr == nil {
					// 刷新成功，设置新的token
					c.Header("New-Token", newToken)
					claims, _ = j.ParseToken(newToken)
				} else {
					gotool.Logs.InfoLog().Printf("token解析失败 222")
					c.JSON(http.StatusOK, gin.H{
						"status": 401,
						"msg":    err.Error(),
					})
					c.Abort()
					return
				}
			} else {
				gotool.Logs.InfoLog().Printf("token解析失败 222")
				c.JSON(http.StatusOK, gin.H{
					"status": 401,
					"msg":    err.Error(),
				})
				c.Abort()
				return
			}
		}
		appServer := config.GetServerCfg()
		lock := appServer.Lock
		if lock == "0" {
			get, err := dao.RedisDB.GET(claims.UserInfo.UserName)
			if err == nil {
				if !(get == s[1]) {
					gotool.Logs.InfoLog().Printf("您的账号已在其他终端登录，请重新登录 333")
					c.JSON(http.StatusOK, gin.H{
						"status": 401,
						"msg":    "您的账号已在其他终端登录，请重新登录",
					})
					c.Abort()
					return
				}
			}
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("授权已过期")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("令牌非法")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "0df9b8db-6f7c-d713-eeab-ecb317696042"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserInfo response.UserResponse `json:"userInfo"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateUserToken 生成含有用户信息的token
func (j *JWT) CreateUserToken(u *response.UserResponse) (string, error) {
	jwtConfig := config.GetJwtConfig()
	now := time.Now()
	// 增加 5 分钟的缓冲时间，避免边界情况
	expiresAt := now.Add(jwtConfig.TimeOut * time.Hour).Add(5 * time.Minute)
	gotool.Logs.InfoLog().Printf("生成token - 当前时间: %v, 过期时间: %v", now, expiresAt)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserInfo: *u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    jwtConfig.Issuer,
		},
	})
	token, err := claims.SignedString(j.SigningKey)
	if err != nil {
		gotool.Logs.InfoLog().Printf("生成token失败: %v", err)
		return "", err
	}
	gotool.Logs.InfoLog().Printf("生成token成功: %s", token)
	return token, nil
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	gotool.Logs.InfoLog().Printf("开始解析token: %s", tokenString)
	gotool.Logs.InfoLog().Printf("当前时间: %v", time.Now())
	
	// 设置验证选项
	parser := jwt.Parser{
		ValidMethods: []string{jwt.SigningMethodHS256.Name},
		SkipClaimsValidation: true, // 跳过 claims 验证
	}
	
	token, err := parser.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		gotool.Logs.InfoLog().Printf("验证token签名方法: %v", token.Method)
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
		gotool.Logs.InfoLog().Printf("token解析失败，错误类型: %T", err)
		return nil, TokenInvalid
	}
	
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 手动验证过期时间
		now := time.Now().Unix()
		gotool.Logs.InfoLog().Printf("token过期时间: %v", time.Unix(claims.ExpiresAt, 0))
		if claims.ExpiresAt < now {
			gotool.Logs.InfoLog().Printf("token已过期，过期时间: %v, 当前时间: %v", 
				time.Unix(claims.ExpiresAt, 0), time.Unix(now, 0))
			return nil, TokenExpired
		}
		gotool.Logs.InfoLog().Printf("token解析成功，过期时间: %v", time.Unix(claims.ExpiresAt, 0))
		return claims, nil
	}
	gotool.Logs.InfoLog().Printf("token无效或claims类型错误")
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	now := time.Now()
	jwt.TimeFunc = func() time.Time {
		return now
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		// 刷新时也增加 5 分钟的缓冲时间
		claims.StandardClaims.ExpiresAt = now.Add(1 * time.Hour).Add(5 * time.Minute).Unix()
		claims.StandardClaims.IssuedAt = now.Unix()
		claims.StandardClaims.NotBefore = now.Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
