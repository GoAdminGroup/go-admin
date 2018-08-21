package admin

import (
	"github.com/buaazp/fasthttprouter"
	"goAdmin/plugins/admin/controllers"
	"goAdmin/modules/auth"
)

func RegisterAdmin(router *fasthttprouter.Router) *fasthttprouter.Router {

	// 仪表盘
	router.GET("/", auth.AuthMiddleware(controller.ShowDashboard))

	// 授权认证
	router.GET("/login", controller.ShowLogin)
	router.POST("/signin", controller.Auth)
	router.GET("/logout",  auth.AuthMiddleware(controller.Logout))

	// 菜单管理
	router.GET("/menu",  auth.AuthMiddleware(controller.ShowMenu))
	router.POST("/menu/delete",  auth.AuthMiddleware(controller.DeleteMenu))
	router.POST("/menu/new",  auth.AuthMiddleware(controller.NewMenu))
	router.GET("/menu/new",  auth.AuthMiddleware(controller.ShowMenu))
	router.POST("/menu/edit",  auth.AuthMiddleware(controller.EditMenu))
	router.GET("/menu/edit/show",  auth.AuthMiddleware(controller.ShowEditMenu))
	router.POST("/menu/order",  auth.AuthMiddleware(controller.MenuOrder))

	// 增删改查管理
	router.GET("/info/:prefix",  auth.AuthMiddleware(controller.ShowInfo))
	router.GET("/info/:prefix/edit",  auth.AuthMiddleware(controller.ShowForm))
	router.GET("/info/:prefix/new",  auth.AuthMiddleware(controller.ShowNewForm))
	router.POST("/edit/:prefix",  auth.AuthMiddleware(controller.EditForm))
	router.POST("/delete/:prefix",  auth.AuthMiddleware(controller.DeleteData))
	router.POST("/new/:prefix",  auth.AuthMiddleware(controller.NewForm))

	// 自动化安装
	router.GET("/install", controller.ShowInstall)
	router.POST("/install/database/check", controller.CheckDatabase)

	return router
}