package role

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gowith/logging"
	"gowith/models"
	"strconv"
	"strings"
	"gowith/models/mysql"
)

//创建角色
func CreateRole(c *gin.Context)  {
	var role models.Role
	var menu models.Menu
	var permission models.Permission
	var roleMenu models.RoleMenu
	var rolePermission models.RolePermission
	//保存权限 ID 集合,菜单 ID 集合
	var extractData = make(map[string]string)
	var menuIDs []string
	var permissionIDs []string
	//解析数据
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &role)
	_ = json.Unmarshal(data, extractData)

	menuIDs = strings.Split(extractData["menu_id"], ",")
	permissionIDs = strings.Split(extractData["permission_id"], ",")

	//检查角色名是否重复
	repeatedRole := role.IsRepeatedRole()
	if repeatedRole{
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "角色名重复",
			"date": 0,
		})
		return
	}

	//检查菜单是否已删除
	for _, menuID := range menuIDs{
		menu.ID, _ = strconv.Atoi(menuID)
		isExistMenu := menu.CheckIsExitByID()
		if !isExistMenu{
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "菜单不存在" + menuID,
				"date": 0,
			})
			return
		}
	}
	//检查权限是否已删除
	for _, permissionID:= range permissionIDs{
		permission.ID, _ = strconv.Atoi(permissionID)
		isExistPermission := permission.CheckIsExitByID()
		if !isExistPermission{
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "权限不存在" + permissionID,
				"date": 0,
			})
			return
		}
	}

	//在事务中执行创建角色、创建角色权限、创建角色菜单
	tx := mysql.DB.Begin()
	err := tx.Create(&role).Error
	if err != nil{
		tx.Rollback()
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 3001,
			"msg": "创建角色失败",
			"date": 0,
		})
	}

	//插入用户菜单表
	for _, menuID := range menuIDs{
		roleMenu.RoleID = role.ID
		roleMenu.MenuID, _ = strconv.Atoi(menuID)
		err = tx.Create(&roleMenu).Error
		if err != nil{
			tx.Rollback()
			logging.Fatal(err)
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "创建用户菜单失败" + menuID,
				"date": 0,
			})
			return
		}
	}

	//插入用户权限表
	for _, permissionID := range permissionIDs{
		rolePermission.RoleID = role.ID
		rolePermission.PermissionID, _ = strconv.Atoi(permissionID)
		err = tx.Create(&rolePermission).Error
		if err != nil{
			//回滚
			tx.Rollback()
			logging.Fatal(err)
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "创建用户菜单失败" + permissionID,
				"date": 0,
			})
			return
		}
	}
	//提交事务
	tx.Commit()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "创建角色成功",
		"date": 1,
	})
}

//更新角色
func UpdateRole(c *gin.Context)  {
	var role models.Role
	var roleMenu models.RoleMenu
	var rolePermission models.RolePermission

	//保存权限 ID 集合,菜单 ID 集合
	var extractData = make(map[string]string)
	var menuIDs []string
	var permissionIDs []string

	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &role)
	_ = json.Unmarshal(data, &extractData)

	isExistRole := role.CheckIsExitByID()
	if !isExistRole{
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "角色不存在",
			"date": 0,
		})
		return
	}

	//根据ID获取角色信息
	roleInfo, _ := role.Find()
	if roleInfo.IsAdmin(){
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "超级管理员角色不允许更新",
			"date": 0,
		})
		return
	}

	isRepeatedName := role.IsRepeatedRole()
	if isRepeatedName{
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "角色名称重复",
			"date": 0,
		})
		return
	}

	//在事务中执行更新角色、更新角色权限、更新角色菜单
	tx := mysql.DB.Begin()
	err := tx.Updates(&role).Error
	if err != nil{
		tx.Rollback()
		logging.Fatal(err)
		c.JSON(200, gin.H{
			"code": 3001,
			"msg": "更新角色失败",
			"date": 0,
		})
	}

	//插入用户菜单表
	for _, menuID := range menuIDs{
		roleMenu.RoleID = role.ID
		roleMenu.MenuID, _ = strconv.Atoi(menuID)
		err = tx.Updates(&roleMenu).Error
		if err != nil{
			tx.Rollback()
			logging.Fatal(err)
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "更新用户菜单失败" + menuID,
				"date": 0,
			})
			return
		}
	}

	//插入用户权限表
	for _, permissionID := range permissionIDs{
		rolePermission.RoleID = role.ID
		rolePermission.PermissionID, _ = strconv.Atoi(permissionID)
		err = tx.Updates(&rolePermission).Error
		if err != nil{
			//回滚
			tx.Rollback()
			logging.Fatal(err)
			c.JSON(200, gin.H{
				"code": 0,
				"msg": "更新用户菜单失败" + permissionID,
				"date": 0,
			})
			return
		}
	}
	//提交事务
	tx.Commit()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "更新角色成功",
		"date": 1,
	})
}

//更新角色部分字段
func UpdateFieldOfRole(c *gin.Context)  {
	var role models.Role
	data, _ := c.GetRawData()
	_ = json.Unmarshal(data, &role)

	isExist := role.CheckIsExitByID()
	if !isExist{
		c.JSON(200, gin.H{
			"code": 400,
			"msg": "角色不存在",
			"date": 0,
		})
	}
	//根据ID获取角色信息
	roleInfo, _ := role.Find()
	if roleInfo.IsAdmin(){
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "超级管理员角色不允许更新",
			"date": 0,
		})
		return
	}

	isRepeatedName := role.IsRepeatedRole()
	if isRepeatedName{
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "角色名称重复",
			"date": 0,
		})
		return
	}

	isValidStatus := role.IsValidStatus()
	if !isValidStatus{
		c.JSON(200, gin.H{
			"code": 200,
			"msg": "status无效",
			"date": 0,
		})
		return
	}

	err := role.UpdateField()
	if err != nil{
		c.JSON(200, gin.H{
			"code": 0,
			"msg": "更新失败",
			"date": 0,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg": "更新成功",
		"date": 0,
	})
}

//角色详情
func FindRole(c *gin.Context)  {
	//获取参数
	data := c.Query("id")
	id, _ := strconv.Atoi(data)

	var role models.Role

	role.ID = id
	result := role.FindRole()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": result,
	})
}

//角色列表
func SearchRole(c *gin.Context)  {
	var role models.Role
	result := role.SearchRole()

	c.JSON(200, gin.H{
		"code": 0,
		"msg": "",
		"date": result,
	})
}
