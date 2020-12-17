package permission

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/models"
	"strconv"
)

//创建权限
func CreatePermission(c *gin.Context)  {
	var permission models.Permission
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &permission)

	err := permission.Create()
	if err != nil{
		c.JSON(200, gin.H{
			"code": 8888,
			"msg": "创建失败",
			"date": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": 0,
	})
}

//更新权限
func UpdatePermission(c *gin.Context)  {
	var permission models.Permission
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &permission)

	err := permission.Update()
	if err != nil{
		c.JSON(200, gin.H{
			"code": 8889,
			"msg": "更新失败",
			"date": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": 1,
	})
}

//更新权限部分字段
func UpdateFieldOfPermission(c *gin.Context)  {
	var permission models.Permission
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &permission)

	isValid := permission.IsValidStatus()
	if !isValid{
		c.JSON(200, gin.H{
			"code": 6666,
			"msg": "status无效",
			"date": 0,
		})
	}

	err := permission.UpdateField()
	if err != nil{
		c.JSON(200, gin.H{
			"code": 8889,
			"msg": "更新失败",
			"date": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": 1,
	})
}

//权限详情
func FindPermission(c *gin.Context)  {
	id := c.Query("id")
	var permission models.Permission
	var result = make(map[string]interface{})
	permission.ID, _ = strconv.Atoi(id)

	data := permission.Find()
	byteData, _ := json.Marshal(data)
	_ = json.Unmarshal(byteData, &result)

	result["ctime"] = data.CTime.Format("2006-01-02 15:04:05")
	result["mtime"] = data.MTime.Format("2006-01-02 15:04:05")

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": result,
	})
}

//权限列表
func SearchPermission(c *gin.Context)  {
	param := c.Query("id")
	var permission models.Permission
	if len(param) != 0{
		permission.ID, _ = strconv.Atoi(param)
	}else{
		permission.ID = -1
	}
	data := permission.Search()
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": data,
	})
}
