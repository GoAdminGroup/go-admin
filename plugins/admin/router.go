package admin

import (
	"goAdmin/context"
	"goAdmin/modules/auth"
	"goAdmin/plugins/admin/controller"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	controller.AssertRootUrl = prefix

	// 仪表盘
	app.GET(prefix + "/", auth.AuthMiddleware(controller.ShowDashboard))

	// 授权认证
	app.GET(prefix + "/login", controller.ShowLogin)
	app.POST(prefix + "/signin", controller.Auth)
	app.GET(prefix + "/logout",  auth.AuthMiddleware(controller.Logout))

	// 菜单管理
	app.GET(prefix + "/menu",  auth.AuthMiddleware(controller.ShowMenu))
	app.POST(prefix + "/menu/delete",  auth.AuthMiddleware(controller.DeleteMenu))
	app.POST(prefix + "/menu/new",  auth.AuthMiddleware(controller.NewMenu))
	app.GET(prefix + "/menu/new",  auth.AuthMiddleware(controller.ShowMenu))
	app.POST(prefix + "/menu/edit",  auth.AuthMiddleware(controller.EditMenu))
	app.GET(prefix + "/menu/edit/show",  auth.AuthMiddleware(controller.ShowEditMenu))
	app.POST(prefix + "/menu/order",  auth.AuthMiddleware(controller.MenuOrder))

	// 增删改查管理
	app.GET(prefix + "/info/:prefix",  auth.AuthMiddleware(controller.ShowInfo))
	app.GET(prefix + "/info/:prefix/edit",  auth.AuthMiddleware(controller.ShowForm))
	app.GET(prefix + "/info/:prefix/new",  auth.AuthMiddleware(controller.ShowNewForm))
	app.POST(prefix + "/edit/:prefix",  auth.AuthMiddleware(controller.EditForm))
	app.POST(prefix + "/delete/:prefix",  auth.AuthMiddleware(controller.DeleteData))
	app.POST(prefix + "/new/:prefix",  auth.AuthMiddleware(controller.NewForm))

	// 自动化安装
	app.GET(prefix + "/install", controller.ShowInstall)
	app.POST(prefix + "/install/database/check", controller.CheckDatabase)

	for _, path := range asserts {
		app.GET(prefix + "/assets" + path, controller.Assert)
	}

	return app
}