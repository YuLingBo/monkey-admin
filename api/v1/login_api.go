package v1

import (
	"monkey-admin/models/req"
	"monkey-admin/pkg/cache"
	"monkey-admin/pkg/library/tree/tree_menu"
	"monkey-admin/pkg/library/user_util"
	"monkey-admin/pkg/resp"
	"monkey-admin/service"

	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
	loginService      service.LoginService
	roleService       service.RoleService
	permissionService service.PermissionService
	menuService       service.MenuService
}

// Login 登录
func (a LoginApi) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	gotool.Logs.InfoLog().Println("收到登录请求，开始处理...")
	
	if c.BindJSON(&loginBody) == nil {
		gotool.Logs.InfoLog().Printf("登录请求参数: username=%s", loginBody.UserName)
		
		m := make(map[string]string)
		login, s := a.loginService.Login(loginBody.UserName, loginBody.Password)
		
		if login {
			gotool.Logs.InfoLog().Printf("用户 %s 登录成功", loginBody.UserName)
			//将token存入到redis中
			user_util.SaveRedisToken(loginBody.UserName, s)
			m["token"] = s
			c.JSON(200, resp.Success(m))
		} else {
			gotool.Logs.ErrorLog().Printf("用户 %s 登录失败: %s", loginBody.UserName, s)
			c.JSON(200, resp.ErrorResp(s))
		}
	} else {
		gotool.Logs.ErrorLog().Println("登录请求参数绑定失败")
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}

// GetUserInfo 获取用户信息
func (a LoginApi) GetUserInfo(c *gin.Context) {
	m := make(map[string]interface{})
	user := a.loginService.LoginUser(c)
	//查询用户角色集合
	roleKeys := a.permissionService.GetRolePermissionByUserId(user)
	// 权限集合
	perms := a.permissionService.GetMenuPermission(user)
	m["roles"] = roleKeys
	m["permissions"] = perms
	m["user"] = user
	c.JSON(200, resp.Success(m))
}

// GetRouters 根据用户ID查询菜单
func (a LoginApi) GetRouters(c *gin.Context) {
	//获取等钱登录用户
	user := a.loginService.LoginUser(c)
	menus := a.menuService.GetMenuTreeByUserId(user)
	systemMenus := tree_menu.SystemMenus{}
	systemMenus = *menus
	array := systemMenus.ConvertToINodeArray(menus)
	generateTree := tree_menu.GenerateTree(array, nil)
	c.JSON(200, resp.Success(generateTree))
}

// Logout 退出登录
func (a LoginApi) Logout(c *gin.Context) {
	//删除Redis缓存
	name := user_util.GetUserInfo(c).UserName
	cache.RemoveKey(name)
	resp.OK(c)
}
