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
	Root       	int 	`gorm:"size:4"  json:"root"`      	//ROOT 用户 {0：否；1：是；}
	Sort		int		`gorm:"size:10" json:"sort"`		//排序（正序）
}
//匹配表名
func (User) TableName() string {
	return "t_passport_user"
}

//定义用户拥有的角色列表
type RoleList struct {
	UserID 	int 	`json:"user_id"`
	RoleID 	int 	`json:"role_id"`
	Name	string 	`json:"name"`
	Admin	int 	`json:"admin"`
}

//定义findUser返回值类
type FindUser struct {
	ID	int `json:"id"`
	Name string `json:"name"`
	Email	string `json:"email"`
	Mobile	string `json:"mobile"`
	Position	string `json:"position"`
	Status	int `json:"status"`
	Root	int `json:"root"`
	Sort 	int `json:"sort"`
	Mtime 	string `json:"mtime"`
	Ctime 	string `json:"ctime"`
	RoleList []RoleList `json:"role_list"`
}
//定义searchUser返回类
type SearchUser struct {
	P	int `json:"p"` //页码
	Size 	int `json:"size"` //每页条数
	Total	int `json:"total"` //总条数
	Next	int `json:"next"` // 是否有下一页 {0：没有；1：有；}
	List 	[]FindUser `json:"list"`  // 用户列表
}


//创建用户
func (u User) Create() error{
	err := mysql.DB.Create(&u).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新用户
func (u User) Update() error{
	err := mysql.DB.Updates(&u).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//更新用户部分字段
func (u User) UpdateField() error{
	err := mysql.DB.Model(&u).Where("id=?", u.ID).UpdateColumn("status", u.Status).Error
	if err != nil{
		logging.Info(err)
		return err
	}
	return nil
}

//获取用户详情
func (u User) Find() (User, error){
	err := mysql.DB.Where("id=?", u.ID).Find(&u).Error
	if err != nil{
		logging.Info(err)
		return u, err
	}
	return u, nil
}

//获取用户列表
func (u User) Search() []User{
	var users []User
	err := mysql.DB.Find(&users).Error
	if err != nil{
		logging.Info(err)
		return users
	}
	return users
}

//获取用户拥有的菜单
func (u User) Menu() ([]Menu, error){
	var menus []Menu
	return menus, nil
}

//登出
func (u User) Logout(email, password string) bool {
	err := mysql.DB.Where("email=? AND password=?", email, password).Find(&u).Error
	if err == nil{
		return true
	}
	return false
}


//验证邮箱是否重复
func (u User)IsExistEmail() bool {
	var count int
	tempID := u.ID
	mysql.DB.Model(&u).Where("email=?", u.Email).Count(&count)
	if tempID == u.ID{
		count --
	}
	return count > 0
}

//检查用户是否root用户
func (u User) IsRoot() bool {
	return u.Root == 1
}

//联表查询用户拥有的角色列表
func (u User) GetRoleList() RoleList {
	var userRole UserRole
	var role Role
	var roleList RoleList

	userRole.UserID = u.ID
	role.ID = userRole.GetRoleID()

	userRoleInfo := userRole.GetUserRole()
	roleInfo, _ := role.Find()

	roleList.UserID = userRoleInfo.UserID
	roleList.RoleID = userRoleInfo.RoleID
	roleList.Name = roleInfo.Name
	roleList.Admin =  roleInfo.Admin

	return roleList
}

//获取用户详情
func (u User) UserFind() FindUser{
	var userRole UserRole
	var role Role
	var result FindUser
	userInfo, err := u.Find()
	if err != nil{
		logging.Info("findError" + err.Error())
		return result
	}

	//去用户角色表查看角色ID
	userRole.UserID = u.ID
	role.ID = userRole.GetRoleID()
	logging.Info(role.ID)
	//用户角色表和角色表连表，需要 1、用户角色表记录状态正常；2、角色表记录状态正常
	isExist := role.CheckIsExitByID() && userRole.IsExitUserRoleByUserID()
	if !isExist{
		logging.Info("用户角色不存在" + err.Error())
	}

	//查询出用户所拥有的角色列表
	info := u.GetRoleList()

	result.ID = userInfo.ID
	result.Name = userInfo.Name
	result.Email = userInfo.Email
	result.Mobile = userInfo.Mobile
	result.Position = userInfo.Position
	result.Status = userInfo.Status
	result.Root = userInfo.Root
	result.Sort = userInfo.Sort
	result.Ctime = userInfo.CTime.Format("2006-01-02 15:04:05")
	result.Mtime = userInfo.MTime.Format("2006-01-02 15:04:05")
	result.RoleList = append(result.RoleList, info)

	return result
}

func (u User) UserSearch() SearchUser {
	var res SearchUser
	var userList []FindUser
	//获取到所有用户的信息，并据此查找用户列表
	users := u.Search()
	for _, user := range users{
		u.ID = user.ID
		info := u.UserFind()
		userList = append(userList, info)
	}

	res.P = 1
	res.Size = 20
	res.Total = len(users)
	if res.Total - (res.P * res.Size)  > 0{
		res.Next = 1
	}else{
		res.Next = 0
	}
	res.List = userList
	return res
}

//根据ID验证记录是否存在
func (u User) CheckIsExitByID() bool{
	var count int
	mysql.DB.Model(&u).Where("id = ? AND status = ?", u.ID, STATUS_NORMAL).Count(&count)
	return count > 0
}

//验证登录请求的邮箱和密码
func (u User) Login() bool {
	var count int
	mysql.DB.Model(&u).Where("email = ? AND password = ?", u.Email, u.Password).Count(&count)
	return count > 0
}

//根据邮箱获取用户信息
func (u User) FindUserByEmail() User {
	err := mysql.DB.Where("email = ?", u.Email).Find(&u).Error
	if err != nil{
		logging.Fatal(err)
	}
	return u
}