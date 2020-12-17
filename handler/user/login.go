package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/logging"
	"gowith/models"
	"gowith/models/cache"
	"gowith/util"
	"strconv"
)

//登录
func Login(c *gin.Context) {
	var user models.User
	var data map[string]string
	result := make(map[string]interface{})

	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &data)

	user.Email = data["email"]
	password := data["password"]
	user.Password = util.MD5(password)

	isExist := user.IsExistEmail()
	if !isExist{
		c.JSON(200, gin.H{
			"code": 1111,
			"msg": "用户不存在",
			"date": 0,
		})
		return
	}

	//校验用户名和密码
	isValidUser := user.Login()
	if !isValidUser{
		c.JSON(200, gin.H{
			"code": 1111,
			"msg": "用户名或密码错误",
			"date": 0,
		})
		return
	}

	//根据email获取用户信息
	userInfo := user.FindUserByEmail()
	ID := strconv.Itoa(userInfo.ID)

	//如果存在token，直接返回
	isExistToken, err := cache.RDS.Get("ID")
	if err != nil{
		logging.Info(err)
	}
	if isExistToken != ""{
		result["expire"] = 7200
		result["token"] = isExistToken

		c.JSON(200, gin.H{
			"code": 200,
			"msg": "",
			"data": result,
		})
		return
	}

	//如果不存在，重新生成token
	token, err := util.GenerateToken(user.Email, user.Password)
	if err != nil{
		logging.Fatal("failed to create token!")
		return
	}

	result["expire"] = 7200
	result["token"] = token

	//缓存用户token和管理员信息
	err = cache.RDS.SetEx(ID, token, 7200)
	if err != nil{
		logging.Error(err)
	}
	err = cache.RDS.SetEx(token, userInfo, 7200)
	if err != nil{
		logging.Error(err)
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg": "",
		"data": result,
	})
}
