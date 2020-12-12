package models

//定义角色菜单类
type RoleMenu struct {
	BaseModel

	RoleID  		int 	`gorm:"size:10" json:"role_id"` 		//角色ID
	MenuID    		int 	`gorm:"size:10" json:"menu_id"`			//菜单ID
	Status    		int 	`gorm:"size:4"  json:"status"`   		//状态{-1：删除；1：正常；}
}

