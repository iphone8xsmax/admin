package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/logging"
	"gowith/models/cache"
)

func Logout(c *gin.Context) {
	var data = make(map[string]string)
	//解析参数
	param, _ := c.GetRawData()
	_ = json.Unmarshal(param, &data)

	accessToken := data["access_token"]

	err := cache.RDS.Delete(accessToken)
	if err != nil{
		logging.Info(err.Error())
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg": "退出登录成功",
		"data": 0,
	})
}