package models

//定义用户角色类
type UserRole struct {
	BaseModel

	UserID  		int 	`gorm:"size:10" json:"role_id"` 		//用户ID
	RoleID  		int 	`gorm:"size:10" json:"role_id"` 		//角色ID
	Status    		int 	`gorm:"size:4"  json:"status"`   		//状态{-1：删除；1：正常；}
}

//检查用户权限
func (u UserRole) ChenkPermission(id int) bool {

	return true
}


//验证role_id是否有效
func (ur *UserRole) CheckRoleID(roleID string) bool {
	return true
}