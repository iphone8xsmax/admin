package routers

import (
	"github.com/gin-gonic/gin"
	"gowith/config"
	"gowith/handler/menu"
	"gowith/handler/permission"
	"gowith/handler/role"
	"gowith/handler/user"
	"gowith/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	//使用自定义日志中间件
	r.Use(gin.Logger())
	//使用错误处理中间件
	r.Use(gin.Recovery())
	//http跨域中间件
	r.Use(middleware.Cors())

	//设置运行模式
	gin.SetMode(config.RunMode)

	//路由分组
	v1 := r.Group("/v1")

	//使用token鉴权中间件
//	caracara.Use(jwt.JWT())
	{
		//创建用户
		v1.POST("/user/create", user.CreateUser)
		//登录
		v1.POST("/user/login", user.Login)


		//用户模块
		//1.更新用户
		v1.POST("/user/update", user.UpdateUser)
		//2.更新用户部分字段
		v1.POST("/user/update_field", user.UpdateFiledOfUser)
		//3.用户详情
		v1.POST("/user/find", user.FindUser)
		//4.用户列表
		v1.GET("/user/search", user.SearchUser)
		//5.用户拥有的菜单
		v1.GET("/user/menu", user.MenuofUser)
		//6.退出登录
		v1.POST("/user/logout", user.Logout)
		//7.检查用户权限
		v1.GET("/user/check_permission", user.CheckPermission)

		//菜单模块
		//1.创建菜单
		v1.POST("/menu/create", menu.CreateMenu)
		//2.更新菜单
		v1.POST("/menu/update", menu.UpdateMenu)
		//3.菜单详情
		v1.GET("/menu/find", menu.FindMenu)
		//4.更新菜单部分字段
		v1.POST("/menu/update_field", menu.UpdateFiledOfMenu)
		//5.菜单列表
		v1.GET("/menu/search", menu.SearchMenu)

		//权限模块
		//1.创建权限
		v1.POST("/permission/create", permission.CreatePermission)
		//2.更新权限
		v1.POST("/permission/update", permission.UpdatePermission)
		//3.更新权限部分字段
		v1.POST("/permission/update_field", permission.UpdateFieldOfPermission)
		//4.权限详情
		v1.GET("/permission/find", permission.FindPermission)
		//5.权限列表
		v1.GET("/permission/search", permission.SearchPermission)

		//角色模块
		//1.创建角色
		v1.POST("/role/create", role.CreateRole)
		//2.更新角色
		v1.POST("/role/update", role.UpdateRole)
		//3.更新角色部分字段
		v1.POST("/role/update_field", role.UpdateFieldOfRole)
		//4.角色详情
		v1.GET("/role/find", role.FindRole)
		//5.角色列表
		v1.GET("/role/search", role.SearchRole)
	}
	return r
}