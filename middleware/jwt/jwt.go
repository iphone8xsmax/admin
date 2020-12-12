package jwt

import (
	"github.com/gin-gonic/gin"
	"gowith/util"
	"time"
)
import "gowith/e"

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		token := c.Request.Header.Get("token")

		code = e.SUCCESS
		if token == ""{
			code = e.INVALID_PARAMS
		}else {
			claims, err := util.ParseToken(token)
			if err != nil{
				code = e.ERROR_USER_CHECK_TOKEN_FAIL
			}else if time.Now().Unix() > claims.ExpiresAt{
				code = e.ERROR_USER_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS{
			c.JSON(500, gin.H{
				"code": code,
				"msg": "账号密码不正确",
			})
			c.Abort()
			return
		}
		//c.JSON(200, gin.H{
		//	"code": code,
		//	"token": token,
		//})

		c.Next()
	}
}
