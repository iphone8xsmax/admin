package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/logging"
	"gowith/models"
	"gowith/models/mysql"
	"gowith/util"
	"strconv"
)

//创建用户
func CreateUser(c *gin.Context) {
	var user models.User
	var userRole models.UserRole
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
	roleID, _ := strconv.Atoi(data["role_id"])

	//验证role_id是否有效
	isValid = userRole.CheckRoleID(roleID)
	if !isValid{
		c.JSON(200, gin.H{
			"code": 3005,
			"msg": "role_id无效",
			"date": 0,
		})
		return
	}

	//验证邮箱是否重复
	isExist := user.IsExistEmail(user.Email)
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
		"msg": "",
		"date": 1,
	})
}

//更新用户
func UpdateUser(c *gin.Context) {

}

//更新用户部分字段
func UpdateFiledOfUser(c *gin.Context) {

}

//用户详情
func FindUser(c *gin.Context) {

}

//用户列表
func SearchUser(c *gin.Context) {

}

//用户拥有的菜单
func MenuofUser(c *gin.Context) {

}

//登录
func Login(c *gin.Context) {

}

//退出登录
func Logout(c *gin.Context) {

}

//检查用户权限
func CheckPermission(c *gin.Context) {

}