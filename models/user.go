package models

import (
	"gowith/logging"
	"gowith/models/mysql"
)

//定义用户类
type User struct {
	BaseModel

	Name  		string 	`gorm:"size:50" json:"name"` 		//姓名
	Email     	string 	`gorm:"size:50" json:"email"`   	//邮箱
	Mobile     	string 	`gorm:"size:11" json:"mobile"`  	//手机号
	Position    string  `gorm:"size:50" json:"position"`    //职位
	Password    string 	`gorm:"size:32" json:"password"`    //密码
	Status    	int 	`gorm:"size:4"  json:"status"`   	//状态{-1：删除；1：正常；}
	Root       	int 	`gorm:"size:4"  json:"root"`      	//ROOT 用户 {0：否；1：是；}
	Sort		int		`gorm:"size:10" json:"sort"`		//排序（正序）
}

//匹配表名
func (User) TableName() string {
	return "t_passport_user"
}

//创建用户
func (u *User) Create(user User) error{
	err := mysql.DB.Model(&u).Create(&user).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新用户
func (u *User) Update(user User) error{
	err := mysql.DB.Model(&u).Updates(user).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新用户部分字段
func (u *User) UpdateField(id int, status int) error{
	err := mysql.DB.Model(&u).Where("id=?", id).UpdateColumn("status", status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//获取用户详情
func (u User) Find(id int) (User, error){
	err := mysql.DB.Where("id=?", id).Find(&u).Error
	if err != nil{
		logging.Info(err)
		return u, err
	}
	return u, nil
}

//获取用户列表
func (u *User) Search() ([]User, error){
	var users []User
	err := mysql.DB.Find(&users).Error
	if err != nil{
		logging.Info(err)
		return users, err
	}
	return users, nil
}

//获取用户拥有的菜单
func (u *User) Menu(id int) ([]Menu, error){
	var menus []Menu
	return menus, nil
}

//登录
func (u User) Login(email, password string) bool {
	err := mysql.DB.Where("email=? AND password=?", email, password).Find(&u).Error
	if err == nil{
		return true
	}
	return false
}

//登出
func (u User) Logout(email, password string) bool {
	err := mysql.DB.Where("email=? AND password=?", email, password).Find(&u).Error
	if err == nil{
		return true
	}
	return false
}

