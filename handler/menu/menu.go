package menu

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/models"
	"strconv"
)

//创建菜单
func CreateMenu(c *gin.Context)  {
	var menu models.Menu
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &menu)

	msg := menu.Create()

	c.JSON(200, gin.H{
		"code": 3005,
		"msg": msg,
		"date": 0,
	})
	return
}

//更新菜单
func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &menu)

	msg := menu.Update()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": msg,
		"date": 0,
	})
}

//更新菜单部分字段
func UpdateFiledOfMenu(c *gin.Context) {
	var menu models.Menu
	body, _ := c.GetRawData()
	_ = json.Unmarshal(body, &menu)

	isValid := menu.IsValidStatus()
	if !isValid{
		c.JSON(200, gin.H{
			"code": 6666,
			"msg": "status无效",
			"date": 0,
		})
	}

	err := menu.UpdateField()
	if err != nil{
		c.JSON(200, gin.H{
			"code": 6668,
			"msg": "更新失败",
			"date": 1,
		})
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": 1,
	})
}

//菜单详情
func FindMenu(c *gin.Context)  {
	id := c.Query("id")
	var menu models.Menu
	var result = make(map[string]interface{})
	menu.ID, _ = strconv.Atoi(id)

	data := menu.Find()
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

//菜单列表
func SearchMenu(c *gin.Context)  {
	param := c.Query("pid")
	var menu models.Menu
	if len(param) != 0{
		 menu.PID, _ = strconv.Atoi(param)
	}else{
		menu.PID = -1
	}
	data := menu.Search()
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": data,
	})
}
