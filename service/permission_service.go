package service

import (
	"github.com/druidcaesa/gotool"
	"monkey-admin/dao"
	"monkey-admin/models"
	"monkey-admin/models/response"
)

type PermissionService struct {
	roleDao dao.RoleDao
	menuDao dao.MenuDao
}

// GetRolePermissionByUserId 查询用户角色集合
func (s PermissionService) GetRolePermissionByUserId(user *response.UserResponse) *[]string {
	admin := models.SysUser{}.IsAdmin(user.UserId)
	roleKeys := s.roleDao.GetRolePermissionByUserId(user.UserId)
	if admin && roleKeys != nil {
		*roleKeys = append(*roleKeys, "admin")
	}
	duplication := gotool.StrArrayUtils.ArrayDuplication(*roleKeys)
	return &duplication
}

// GetMenuPermission 获取菜单数据权限
func (s PermissionService) GetMenuPermission(user *response.UserResponse) *[]string {
	flag := models.SysUser{}.IsAdmin(user.UserId)
	//查询菜单数据权限
	permission := s.menuDao.GetMenuPermission(user.UserId)
	if flag && permission != nil {
		*permission = append(*permission, "*:*:*")
	}
	var ret []string
	duplication := gotool.StrArrayUtils.ArrayDuplication(*permission)
	for i := 0; i < len(duplication); i++ {
		if (i > 0 && duplication[i-1] == duplication[i]) || len(duplication[i]) == 0 {
			continue
		}
		ret = append(ret, duplication[i])
	}
	return &ret
}
