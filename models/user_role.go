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
	Status    		int 	`gorm:"size:4"  json:"status"`   		//状态{-1：删除；1：正常；}
}

//匹配表名
func (UserRole) TableName() string {
	return "t_passport_user_role"
}

//创建用户
func (ur UserRole) Create() error{
	err := mysql.DB.Create(&ur).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}


//检查用户权限
func (u UserRole) ChenkPermission(id int) bool {

	return true
}


//验证role_id是否有效
func (ur *UserRole) CheckRoleID(roleID int) bool {
	var count int
	mysql.DB.Model(&ur).Where("role_id=?", roleID).Count(&count)
	if count <= 0{
		if ur.Status == -1 || ur.Status == 1{
			return false
		}
		return true
	}
	return false
}