package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/template"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	// 授权认证
	app.GET(prefix + "/login", controller.ShowLogin)
	app.POST(prefix + "/signin", controller.Auth)

	// 自动化安装
	app.GET(prefix + "/install", controller.ShowInstall)
	app.POST(prefix + "/install/database/check", controller.CheckDatabase)

	for _, path := range template.Get("adminlte").GetAssetList() {
		app.GET(prefix + "/assets" + path, controller.Assert)
	}

	for _, path := range template.GetComp("login").GetAssetList() {
		app.GET(prefix + "/assets" + path, controller.Assert)
	}

	// 仪表盘
	if prefix != "" {
		app.GET(prefix, auth.SetPrefix(prefix).Middleware(controller.ShowDashboard))
	} else {
		app.GET("/", auth.SetPrefix(prefix).Middleware(controller.ShowDashboard))
	}

	app.Group(prefix, auth.SetPrefix(prefix).Middleware)
	{
		// 授权认证
		app.GET("/logout",  controller.Logout)

		// 菜单管理
		app.GET("/menu",  controller.ShowMenu)
		app.POST("/menu/delete",  controller.DeleteMenu)
		app.POST("/menu/new",  controller.NewMenu)
		app.GET("/menu/new",  controller.ShowMenu)
		app.POST("/menu/edit",  controller.EditMenu)
		app.GET("/menu/edit/show",  controller.ShowEditMenu)
		app.POST("/menu/order",  controller.MenuOrder)

		// 增删改查管理
		app.GET("/info/:prefix",  controller.ShowInfo)
		app.GET("/info/:prefix/edit",  controller.ShowForm)
		app.GET("/info/:prefix/new",  controller.ShowNewForm)
		app.POST("/edit/:prefix",  controller.EditForm)
		app.POST("/delete/:prefix",  controller.DeleteData)
		app.POST("/new/:prefix",  controller.NewForm)
	}

	return app
}