package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义用户权限类
type RolePermission struct {
	BaseModel

	RoleID  		int 	`gorm:"size:10" json:"role_id"` 		//角色ID
	PermissionID    int 	`gorm:"size:10" json:"permission_id"`	//权限ID
}

//匹配表名
func (RolePermission) TableName() string {
	return "t_passport_role_permission"
}


func (rp RolePermission)Create() error {
	err := mysql.DB.Create(&rp).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//根据RoleID获取PermissionID
func (rp RolePermission)GetPermissionIDByRoleID() []int {
	var rolePermissions []RolePermission
	var permissionIDs []int
	err := mysql.DB.Where("role_id = ?", rp.RoleID).Find(&rolePermissions).Error
	if err != nil{
		logging.Info(err)
		return []int{}
	}

	for _, value := range rolePermissions{
		permissionIDs = append(permissionIDs, value.PermissionID)
	}
	return permissionIDs
}

