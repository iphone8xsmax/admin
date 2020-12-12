package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义菜单类
type Menu struct {
	BaseModel

	PID			int		`gorm:"size:10" json:"pid"`			//父级ID {0：顶级菜单；}
	Name  		string 	`gorm:"size:20" json:"name"` 		//菜单名称
	Icon		string	`gorm:"size:50" json:"icon"`		//菜单图标
	Url     	string 	`gorm:"size:50" json:"url"`   		//路由
	Status    	int 	`gorm:"size:4"  json:"status"`   	//状态{-1：删除；1：正常；}
	Sort		int		`gorm:"size:10" json:"sort"`		//排序（正序）
}

//匹配表名
func (Menu) TableName() string {
	return "t_passport_menu"
}

//创建菜单
func (m *Menu) Create(menu Menu) error{
	err := mysql.DB.Model(&m).Create(&menu).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新菜单
func (m *Menu) Update(menu Menu) error{
	err := mysql.DB.Model(&m).Updates(menu).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新菜单部分字段
func (m *Menu) UpdateField(id, sort, status int) error{
	err := mysql.DB.Model(&m).Where("id=?", id).UpdateColumn("status", status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//获取菜单列表
func (m *Menu) Search(pid int) ([]Menu, error){
	var menus []Menu
	if pid != 0{
		err := mysql.DB.Where("pid=?", pid).Find(&menus).Error
		if err != nil{
			logging.Info(err)
			return menus, err
		}
	}else{
		err := mysql.DB.Find(&menus).Error
		if err != nil{
			logging.Info(err)
			return menus, err
		}
	}

	return menus, nil
}