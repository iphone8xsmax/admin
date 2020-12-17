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
	Sort		int		`gorm:"size:10" json:"sort"`		//排序（正序）
}
//匹配表名
func (Menu) TableName() string {
	return "t_passport_menu"
}

//定义菜单列表类
type MenuList struct {
	ID	  	int			`json:"id"`		//主键ID
	Status  int 		`json:"status"`   //状态{-1：删除；1：正常；}
	PID			int		`json:"pid"`			//父级ID {0：顶级菜单；}
	Name  		string 	`json:"name"` 		//菜单名称
	Icon		string	`json:"icon"`		//菜单图标
	Url     	string 	`json:"url"`   		//路由
	Sort		int		`json:"sort"`		//排序（正序）

	CreateTime 	string `json:"ctime"`
	UpdateTime 	string `json:"mtime"`
}

//创建菜单
func (m Menu) Create() string{
	var err error
	if m.PID != 0 && len(m.Url) == 0{
		return "二级菜单必须输入路由！"
	}
	err = mysql.DB.Create(&m).Error
	if err != nil{
		logging.Info(err)
		return err.Error()
	}
	return ""
}

//更新菜单
func (m Menu) Update() string{
	var err error
	if m.PID != 0 && len(m.Url) == 0{
		return "二级菜单必须输入路由！"
	}
	err = mysql.DB.Updates(&m).Error
	if err != nil{
		logging.Info(err)
		return err.Error()
	}
	return ""
}

//获取菜单详情
func (m Menu) Find() Menu{
	err := mysql.DB.Where("id=?", m.ID).Find(&m).Error
	if err != nil{
		logging.Info(err)
	}

	return m
}

//更新菜单部分字段
func (m Menu) UpdateField() error{
	//如果删除一级菜单，需要先删除下面的二级菜单
	if m.PID == 0{
		err := mysql.DB.Table("t_passport_menu").Where("pid=?", m.ID).UpdateColumn("status", m.Status).Error
		if err != nil{
			logging.Info(err)
			return err
		}
	}
	err := mysql.DB.Model(&m).Where("id=?", m.ID).UpdateColumn("status", m.Status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//获取菜单列表
func (m Menu) Search() []MenuList{
	var menus []Menu
	var menuList MenuList
	var menuLists []MenuList
	if m.PID != -1{
		err := mysql.DB.Where("pid=?", m.PID).Find(&menus).Error
		if err != nil{
			logging.Info(err)
			return menuLists
		}
	}else{
		err := mysql.DB.Find(&menus).Error
		if err != nil{
			logging.Info(err)
			return menuLists
		}
	}
	for _, value := range menus{
		menuList.ID = value.ID
		menuList.Status = value.Status
		menuList.PID = value.PID
		menuList.Name = value.Name
		menuList.Url = value.Url
		menuList.Icon = value.Icon
		menuList.Sort = value.Sort

		menuList.CreateTime = value.CTime.Format("2006-01-02 15:04:05")
		menuList.UpdateTime = value.MTime.Format("2006-01-02 15:04:05")

		menuLists = append(menuLists, menuList)
	}

	return menuLists
}

//根据ID获取菜单列表
func (m Menu) SearchByID() MenuList{
	var menuList MenuList

	if m.PID != -1{
		err := mysql.DB.Where("id=? AND pid = ?", m.ID, m.PID).Find(&m).Error
		if err != nil{
			logging.Info(err)
			return menuList
		}
	}else{
		err := mysql.DB.Where("id=?", m.ID).Find(&m).Error
		if err != nil{
			logging.Info(err)
			return menuList
		}
	}

	menuList.ID = m.ID
	menuList.Status = m.Status
	menuList.PID = m.PID
	menuList.Name = m.Name
	menuList.Url = m.Url
	menuList.Icon = m.Icon
	menuList.Sort = m.Sort
	menuList.CreateTime = m.CTime.Format("2006-01-02 15:04:05")
	menuList.UpdateTime = m.MTime.Format("2006-01-02 15:04:05")

	return menuList
}
//根据ID验证记录是否存在
func (m Menu) CheckIsExitByID() bool{
	var count int
	mysql.DB.Model(&m).Where("id = ? AND status = ?", m.ID, STATUS_NORMAL).Count(&count)
	return count > 0
}