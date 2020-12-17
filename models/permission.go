package models

import (
	"gowith/logging"
	"gowith/models/mysql"
	"strings"
)

//定义权限类
type Permission struct {
	BaseModel

	Name  		string 	`gorm:"size:50"   json:"name"` 		//权限名称
	Url     	string 	`gorm:"size:2000" json:"url"`   		//路由（多个之间用英文分号隔开）
	Sort		int		`gorm:"size:10"   json:"sort"`		//排序（正序）
}

//匹配表名
func (Permission) TableName() string {
	return "t_passport_permission"
}

type PerList struct {
	ID int `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	MTime string `json:"mtime"`
	CTime string `json:"ctime"`
	UrlList []string `json:"url_list"`
}

//权限列表类
type PermissionList struct {
	P	int `json:"p"`
	Size int `json:"size"`
	Total 	int `json:"total"`
	Next 	int `json:"next"`
	List	[]PerList
}

//创建权限
func (p Permission) Create() error {
	err := mysql.DB.Create(&p).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新权限
func (p Permission) Update() error{
	err := mysql.DB.Updates(&p).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新用户部分字段
func (p Permission) UpdateField() error{
	err := mysql.DB.Model(&p).Where("id=?", p.ID).UpdateColumn("status", p.Status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}
//获取权限详情
func (p Permission) Find() Permission{
	err := mysql.DB.Where("id=?", p.ID).Find(&p).Error
	if err != nil{
		logging.Info(err)
		return p
	}
	return p
}

//获取权限列表
func (p Permission) Search() PermissionList{
	var permissions []Permission
	var permissionList PermissionList
	var perList PerList
	var perLists []PerList
	if p.ID != -1{
		err := mysql.DB.Where("id=?", p.ID).Find(&permissions).Error
		if err != nil{
			logging.Info(err)
			return permissionList
		}
	}else{
		err := mysql.DB.Find(&permissions).Error
		if err != nil{
			logging.Info(err)
			return permissionList
		}
	}

	for _, value := range permissions{
		perList.ID = value.ID
		perList.Status = value.Status
		perList.Name = value.Name
		perList.Sort = value.Sort
		perList.URL = value.Url
		perList.UrlList = strings.Split(perList.URL, ";")

		perList.CTime = value.CTime.Format("2006-01-02 15:04:05")
		perList.MTime = value.MTime.Format("2006-01-02 15:04:05")

		perLists = append(perLists, perList)
	}

	permissionList.P = 1
	permissionList.Size = 20
	permissionList.Total = len(perLists)
	if permissionList.Total - (permissionList.P * permissionList.Size)  > 0{
		permissionList.Next = 1
	}else{
		permissionList.Next = 0
	}
	permissionList.List = perLists

	return permissionList
}

//根据ID验证记录是否存在
func (p Permission) CheckIsExitByID() bool{
	var count int
	mysql.DB.Model(&p).Where("id = ? AND status = ?", p.ID, STATUS_NORMAL).Count(&count)
	return count > 0
}