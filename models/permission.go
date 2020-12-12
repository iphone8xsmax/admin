package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义权限类
type Permission struct {
	BaseModel

	Name  		string 	`gorm:"size:50"   json:"name"` 		//权限名称
	Url     	string 	`gorm:"size:2000" json:"url"`   		//路由（多个之间用英文分号隔开）
	Status    	int 	`gorm:"size:4"    json:"status"`   	//状态{-1：删除；1：正常；}
	Sort		int		`gorm:"size:10"   json:"sort"`		//排序（正序）
}

//匹配表名
func (Permission) TableName() string {
	return "t_passport_permission"
}

//创建权限
func (p *Permission) Create(permission Permission) error {
	err := mysql.DB.Model(&p).Create(&permission).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新权限
func (p *Permission) Update(permission Permission) error{
	err := mysql.DB.Model(&p).Updates(permission).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新用户部分字段
func (p *Permission) UpdateField(id int, status int) error{
	err := mysql.DB.Model(&p).Where("id=?", id).UpdateColumn("status", status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}
//获取权限详情
func (p Permission) Find(id int) (Permission, error){
	err := mysql.DB.Where("id=?", id).Find(&p).Error
	if err != nil{
		logging.Info(err)
		return p, err
	}
	return p, nil
}

//获取权限列表
func (p Permission) Search() ([]Permission, error){
	var pers []Permission
	err := mysql.DB.Find(&pers).Error
	if err != nil{
		logging.Info(err)
		return pers, err
	}
	return pers, nil
}
