package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/template"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	// 仪表盘
	if prefix != "" {
		app.GET(prefix, auth.AuthMiddleware(controller.ShowDashboard, prefix))
	} else {
		app.GET("/", auth.AuthMiddleware(controller.ShowDashboard, prefix))
	}

	// 授权认证
	app.GET(prefix + "/login", controller.ShowLogin)
	app.POST(prefix + "/signin", controller.Auth)
	app.GET(prefix + "/logout",  auth.AuthMiddleware(controller.Logout))

	// 菜单管理
	app.GET(prefix + "/menu",  auth.AuthMiddleware(controller.ShowMenu, prefix))
	app.POST(prefix + "/menu/delete",  auth.AuthMiddleware(controller.DeleteMenu, prefix))
	app.POST(prefix + "/menu/new",  auth.AuthMiddleware(controller.NewMenu, prefix))
	app.GET(prefix + "/menu/new",  auth.AuthMiddleware(controller.ShowMenu, prefix))
	app.POST(prefix + "/menu/edit",  auth.AuthMiddleware(controller.EditMenu, prefix))
	app.GET(prefix + "/menu/edit/show",  auth.AuthMiddleware(controller.ShowEditMenu, prefix))
	app.POST(prefix + "/menu/order",  auth.AuthMiddleware(controller.MenuOrder, prefix))

	// 增删改查管理
	app.GET(prefix + "/info/:prefix",  auth.AuthMiddleware(controller.ShowInfo, prefix))
	app.GET(prefix + "/info/:prefix/edit",  auth.AuthMiddleware(controller.ShowForm, prefix))
	app.GET(prefix + "/info/:prefix/new",  auth.AuthMiddleware(controller.ShowNewForm, prefix))
	app.POST(prefix + "/edit/:prefix",  auth.AuthMiddleware(controller.EditForm, prefix))
	app.POST(prefix + "/delete/:prefix",  auth.AuthMiddleware(controller.DeleteData, prefix))
	app.POST(prefix + "/new/:prefix",  auth.AuthMiddleware(controller.NewForm, prefix))

	// 自动化安装
	app.GET(prefix + "/install", controller.ShowInstall)
	app.POST(prefix + "/install/database/check", controller.CheckDatabase)

	for _, path := range template.Get("adminlte").GetAssetList() {
		app.GET(prefix + "/assets" + path, controller.Assert)
	}

	for _, path := range template.GetComp("login").GetAssetList() {
		app.GET(prefix + "/assets" + path, controller.Assert)
	}

	return app
}