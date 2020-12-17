package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义用户角色类
type UserRole struct {
	BaseModel

	UserID  		int 	`gorm:"size:10" json:"role_id"` 		//用户ID
	RoleID  		int 	`gorm:"size:10" json:"role_id"` 		//角色ID
}

//匹配表名
func (UserRole) TableName() string {
	return "t_passport_user_role"
}

//创建用户角色
func (ur UserRole) Create() error{
	err := mysql.DB.Create(&ur).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//检查用户权限
func (ur UserRole) ChenkPermission() bool {

	return true
}

//根据userRole.ID验证用户角色是否存在
func (ur UserRole) IsExitUserRole() bool {
	var count int
	mysql.DB.Model(&ur).Where("id = ? AND status = ?", ur.ID, STATUS_NORMAL).Count(&count)
	return count > 0
}

//根据userRole.UserID验证用户角色是否存在
func (ur UserRole) IsExitUserRoleByUserID() bool {
	var count int
	mysql.DB.Model(&ur).Where("user_id = ? AND status = ?", ur.UserID, STATUS_NORMAL).Count(&count)
	return count > 0
}

//根据userRole.UserID查询角色ID
func (ur UserRole) GetRoleID() int {
	err := mysql.DB.Where("user_id = ? AND status = ?", ur.UserID, STATUS_NORMAL).Find(&ur).Error
	if err != nil{
		logging.Error(err)
		return ur.RoleID
	}
	return ur.RoleID
}

//查询用户权限信息
func (ur UserRole) GetUserRole() UserRole {
	err := mysql.DB.Where("user_id = ? AND status = ?", ur.UserID, STATUS_NORMAL).Find(&ur).Error
	if err != nil{
		logging.Error(err)
		return ur
	}
	return ur
}



