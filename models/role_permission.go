package models

//定义用户权限类
type RolePermission struct {
	BaseModel

	RoleID  		int 	`gorm:"size:10" json:"role_id"` 		//角色ID
	PermissionID    int 	`gorm:"size:10" json:"permission_id"`	//权限ID
	Status    		int 	`gorm:"size:4"  json:"status"`   		//状态{-1：删除；1：正常；}
}
