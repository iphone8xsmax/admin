package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义角色类
type Role struct {
	BaseModel

	Name  		string 	`gorm:"size:50" json:"name"` 		//角色名称
	Admin     	int 	`gorm:"size:4"  json:"admin"`   	//超级管理员 {0：否；1：是；}
	Status    	int 	`gorm:"size:4"  json:"status"`   	//状态{-1：删除；1：正常；}
	Sort		int		`gorm:"size:10" json:"sort"`		//排序（正序）
}


//匹配表名
func (Role) TableName() string {
	return "t_passport_role"
}

//创建角色
func (r *Role) Create(role Role) error{
	err := mysql.DB.Model(&r).Create(&role).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新角色
func (r *Role) Update(role Role) error{
	err := mysql.DB.Model(&r).Updates(role).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新角色部分字段
func (r *Role) UpdateField(id int, status int) error{
	err := mysql.DB.Model(&r).Where("id=?", id).UpdateColumn("status", status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//获取角色详情
func (r Role) Find(id int) (Role, error){
	err := mysql.DB.Where("id=?", id).Find(&r).Error
	if err != nil{
		logging.Info(err)
		return r, err
	}
	return r, nil
}

//获取角色列表
func (r *Role) Search() ([]Role, error){
	var roles []Role
	err := mysql.DB.Find(&roles).Error
	if err != nil{
		logging.Info(err)
		return roles, err
	}
	return roles, nil
}