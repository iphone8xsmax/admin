package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义角色菜单类
type RoleMenu struct {
	BaseModel

	RoleID  		int 	`gorm:"size:10" json:"role_id"` 		//角色ID
	MenuID    		int 	`gorm:"size:10" json:"menu_id"`			//菜单ID
}

//匹配表名
func (RoleMenu) TableName() string {
	return "t_passport_role_menu"
}

//根据RoleID获取MenuID
func (rm RoleMenu)GetMenuIDByRoleID() []int {
	var roleMenus []RoleMenu
	var menuIDs []int

	err := mysql.DB.Where("role_id = ?", rm.RoleID).Find(&roleMenus).Error
	if err != nil{
		logging.Info(err)
		return []int{}
	}
	for _, value := range roleMenus{
		menuIDs = append(menuIDs, value.MenuID)
	}
	return menuIDs
}

