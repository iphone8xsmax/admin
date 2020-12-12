package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gowith/models"
	"gowith/util"
)

//创建用户
func CreateUser(c *gin.Context) {
	var user *models.User
	var userRole *models.UserRole
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &userRole)
	_ = json.Unmarshal(body, &user)

	fmt.Println(userRole)

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
	//isValid = userRole.CheckRoleID(userRole.RoleID)


	c.JSON(200, gin.H{
		"code": 200,
		"msg": "",
		"date": 2,
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