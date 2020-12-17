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
	Sort		int		`gorm:"size:10" json:"sort"`		//排序（正序）
}
//匹配表名
func (Role) TableName() string {
	return "t_passport_role"
}

type FindRole struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Admin int `json:"admin"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	MTime string `json:"mtime"`
	CTime string `json:"ctime"`
	PermissionList []PermissionList `json:"permission_list"`
	MenuList	[]MenuList `json:"menu_list"`
}

type SearchRole struct {
	P	int `json:"p"` //页码
	Size 	int `json:"size"` //每页条数
	Total	int `json:"total"` //总条数
	Next	int `json:"next"` // 是否有下一页 {0：没有；1：有；}
	List 	[]FindRole `json:"list"`  // 用户列表
}

//创建角色
func (r Role) Create() error{
	err := mysql.DB.Create(&r).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新角色
func (r Role) Update() error{
	err := mysql.DB.Model(&r).Updates(r).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新角色部分字段
func (r Role) UpdateField() error{
	err := mysql.DB.Model(&r).Where("id=?", r.ID).UpdateColumn("status", r.Status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//获取角色详情
func (r Role) Find() (Role, error){
	err := mysql.DB.Where("id=?", r.ID).Find(&r).Error
	if err != nil{
		logging.Info(err)
		return r, err
	}
	return r, nil
}

//获取角色列表
func (r Role) Search() []Role{
	var roles []Role
	err := mysql.DB.Find(&roles).Error
	if err != nil{
		logging.Info(err)
		return roles
	}
	return roles
}

//根据角色Name验证角色是否存在
func (r Role) IsRepeatedRole() bool {
	var count int
	mysql.DB.Model(&r).Where("name = ? AND status = ?", r.Name, STATUS_NORMAL).Count(&count)
	return count > 0
}

//检查是否超级管理员
func (r Role) IsAdmin() bool {
	return r.Admin == 1
}

//获取角色详情列表
func (r Role) FindRole() FindRole {
	var findRole FindRole
	var roleMenu RoleMenu
	var rolePermission RolePermission
	var menu Menu
	var permission Permission
	var menuList MenuList
	var permissionList PermissionList
	//var menuLists []MenuList
	//var permissionLists []PermissionList

	info, _ := r.Find()
	findRole.ID = info.ID
	findRole.Name = info.Name
	findRole.Admin = info.Admin
	findRole.Status = info.Status
	findRole.Sort = info.Sort
	findRole.MTime = info.MTime.Format("2006-01-02 15:04:05")
	findRole.CTime = info.CTime.Format("2006-01-02 15:04:05")

	//得到MenuID数组和PermissionID数组
	roleMenu.RoleID = findRole.ID
	menuIDs := roleMenu.GetMenuIDByRoleID()
	rolePermission.RoleID = findRole.ID
	permissionIDs := rolePermission.GetPermissionIDByRoleID()

	for _, menuID := range menuIDs{
		menu.ID = menuID
		menuList = menu.SearchByID()
		findRole.MenuList = append(findRole.MenuList, menuList)
	}

	for _, permissionID := range permissionIDs{
		permission.ID = permissionID
		permissionList = permission.Search()
		findRole.PermissionList = append(findRole.PermissionList, permissionList)
	}

	return findRole
}

//获取角色列表
func (r Role) SearchRole() SearchRole {
	var searchRole SearchRole
	var roleList []FindRole
	//获取到所有用户的信息，并据此查找用户列表
	roles := r.Search()
	for _, role := range roles{
		r.ID = role.ID
		info := r.FindRole()
		roleList = append(roleList, info)
	}

	searchRole.P = 1
	searchRole.Size = 20
	searchRole.Total = len(roles)
	if searchRole.Total - (searchRole.P * searchRole.Size)  > 0{
		searchRole.Next = 1
	}else{
		searchRole.Next = 0
	}
	searchRole.List = roleList

	return searchRole
}

//根据ID验证记录是否存在
func (r Role) CheckIsExitByID() bool{
	var count int
	mysql.DB.Model(&r).Where("id = ? AND status = ?", r.ID, STATUS_NORMAL).Count(&count)
	return count > 0
}
