package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/logging"
	"gowith/models"
	"gowith/models/cache"
	"gowith/models/mysql"
	"gowith/util"
	"strconv"
	"strings"
)

//创建用户
func CreateUser(c *gin.Context) {
	//定义实体类
	var user models.User
	var userRole models.UserRole
	var role models.Role

	//参数解析
	var data = make(map[string]string)
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &data)
	_ = json.Unmarshal(body, &user)

	//验证邮箱是否合法
	isValid := util.CheckEamil(user.Email)
	if !isValid{
		c.JSON(200, gin.H{
			"code": 2001,
			"msg": "邮箱不合法",
			"date": 0,
		})
		return
	}

	//验证role_id是否有效
	role.ID, _ = strconv.Atoi(data["role_id"])
	isExist := role.CheckIsExitByID()
	if !isValid{
		c.JSON(200, gin.H{
			"code": 3005,
			"msg": "role_id无效",
			"date": 0,
		})
		return
	}

	//验证邮箱是否重复
	isExist = user.IsExistEmail()
	if isExist{
		c.JSON(200, gin.H{
			"code": 2002,
			"msg": "邮箱重复使用",
			"date": 0,
		})
		return
	}

	//密码需要加密
	user.Password = util.MD5(user.Password)

	//创建用户和用户角色，事务中执行
	tx := mysql.DB.Begin()
	err := tx.Create(&user).Error
	if err != nil{
		tx.Rollback()
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 3001,
			"msg": "创建用户失败",
			"date": 0,
		})
	}

	userRole.UserID = user.ID
	userRole.RoleID = role.ID
	err = tx.Create(&userRole).Error
	if err != nil{
		tx.Rollback()
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 3002,
			"msg": "创建用户角色失败",
			"date": 0,
		})
	}
	//提交事务
	tx.Commit()

	c.JSON(200, gin.H{
		"code": 200,
		"msg": "创建用户成功",
		"date": 1,
	})
}

//更新用户
func UpdateUser(c *gin.Context) {
	var user models.User
	var userRole models.UserRole
	var role models.Role
	var data = make(map[string]string)
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &data)
	_ = json.Unmarshal(body, &user)

	//验证邮箱是否合法
	isValid := util.CheckEamil(user.Email)
	if !isValid{
		c.JSON(200, gin.H{
			"code": 2001,
			"msg": "邮箱不合法",
			"date": false,
		})
		return
	}

	//验证用户是否存在
	isExist := user.CheckIsExitByID()
	if !isExist{
		c.JSON(200, gin.H{
			"code": 2005,
			"msg": "用户不存在",
			"date": false,
		})
		return
	}

	//检查是否root用户
	root := user.IsRoot()
	if root{
		c.JSON(200, gin.H{
			"code": 3008,
			"msg": "root用户不允许更新",
			"date": false,
		})
		return
	}

	//验证role_id是否有效
	role.ID, _ = strconv.Atoi(data["role_id"])
	isExist = role.CheckIsExitByID()
	if !isValid{
		c.JSON(200, gin.H{
			"code": 3005,
			"msg": "角色不存在",
			"date": 0,
		})
		return
	}

	//验证邮箱是否重复
	isExist = user.IsExistEmail()
	if isExist{
		c.JSON(200, gin.H{
			"code": 2002,
			"msg": "邮箱重复使用",
			"date": false,
		})
		return
	}

	//密码需要加密
	user.Password = util.MD5(user.Password)


	//创建用户和用户角色，事务中执行
	tx := mysql.DB.Begin()
	err := tx.Updates(&user).Error
	if err != nil{
		tx.Rollback()
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 3001,
			"msg": "创建用户失败",
			"date": false,
		})
	}

	userRole.UserID = user.ID
	userRole.RoleID = role.ID
	err = tx.Updates(&userRole).Error
	if err != nil{
		tx.Rollback()
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 3002,
			"msg": "创建用户角色失败",
			"date": false,
		})
	}
	//提交事务
	tx.Commit()

	c.JSON(200, gin.H{
		"code": 200,
		"msg": "更新成功",
		"date": true,
	})

}

//更新用户部分字段
func UpdateFiledOfUser(c *gin.Context) {
	var user models.User

	//参数解析
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &user)

	//检查status是否有效
	isValid := user.IsValidStatus()
	if !isValid{
		c.JSON(200, gin.H{
			"code": 20009,
			"msg": "status无效",
			"date": false,
		})
		return
	}

	//root用户不能删除
	isRoot := user.IsRoot()
	if isRoot{
		c.JSON(200, gin.H{
			"code": 3008,
			"msg": "root用户不允许更新",
			"date": false,
		})
		return
	}

	err := user.UpdateField()
	if err != nil{
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 20009,
			"msg": "status无效",
			"date": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "更新成功",
		"date": 1,
	})
}

//用户详情(联表查询还不会用)
func FindUser(c *gin.Context) {
	//获取参数--用户id
	data := c.Query("id")
	id, _ := strconv.Atoi(data)

	var user models.User

	user.ID = id
	result := user.UserFind()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": result,
	})
}

//用户列表
func SearchUser(c *gin.Context) {
	var user models.User
	//data := c.Request.URL.Query()
	result := user.UserSearch()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": result,
	})
}

//用户拥有的菜单
func MenuOfUser(c *gin.Context) {
	var userRole models.UserRole
	var role models.Role
	var user models.User

	data := c.Query("id")
	userRole.UserID, _ = strconv.Atoi(data)
	user.ID = userRole.UserID
	isExist := user.CheckIsExitByID()
	if !isExist{
		c.JSON(200, gin.H{
			"code": 65656,
			"msg": "用户不存在",
			"date": 0,
		})
	}
	role.ID = userRole.GetRoleID()

	result := role.FindRole()
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": result.MenuList,
	})
}

//检查用户权限
func CheckPermission(c *gin.Context) {
	var userRole models.UserRole
	var role models.Role
	var permission models.Permission
	var rolePermission models.RolePermission

	var data = make(map[string]string)
	urls := []string{"/"}
	param, _ := c.GetRawData()
	_ = json.Unmarshal(param, &data)

	url := data["url"]
	accessToken := data["access_token"]

	//根据 access_token 获取用户 ID，如果从缓存中没有获取到用户 ID，抛出异常
	userID, _ := cache.RDS.Get(accessToken)
	if userID == nil{
		c.JSON(200, gin.H{
			"code": 8585,
			"msg": "用户不存在",
			"date": 0,
		})
		return
	}

	//部分路由直接返回不需要校验权限
	for _, v := range urls{
		if v == url{
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "",
				"date": 1,
			})
			return
		}
	}

	//如果用户有超级管理员角色，直接返回
	userRole.UserID, _ = userID.(int) //利用类型断言将interface转为具体类型值
	role.ID = userRole.GetRoleID()
	roleInfo, err := role.Find()
	if err != nil{
		c.JSON(200, gin.H{
			"code": 8959,
			"msg": err,
			"date": 1,
		})
		return
	}
	if roleInfo.Admin ==  1{
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "",
			"date": 1,
		})
		return
	}

	//根据RoleID获取PermissionID
	rolePermission.RoleID = role.ID
	permissionIDs := rolePermission.GetPermissionIDByRoleID()

	//获取对应的权限路由
	for _, id := range permissionIDs{
		permission.ID = id
		//获取路由，并切分
		urlContant := permission.Find().Url
		urlList := strings.Split(urlContant, ";")

		for _, value := range urlList{
			if !util.IsContains(value, urls){ //去重
				urls = append(urls, value)
			}
		}
	}
	//请求的路由在该用户拥有的权限列表中
	for _, value := range urls{
		if value == url{
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "",
				"date": 1,
			})
			return
		}
	}
	//请求的路由不在该用户拥有的权限列表中，则表示该用户无该路由的权限
	c.JSON(200, gin.H{
		"code": 5859,
		"msg": "",
		"date": 0,
	})
	return
}