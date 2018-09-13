package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/modules/auth"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	app.Group(prefix)
	{
		// 授权认证
		app.GET("/login", controller.ShowLogin)
		app.POST("/signin", controller.Auth)

		// 自动化安装
		app.GET("/install", controller.ShowInstall)
		app.POST("/install/database/check", controller.CheckDatabase)

		for _, path := range template.Get("adminlte").GetAssetList() {
			app.GET("/assets" + path, controller.Assert)
		}

		for _, path := range template.GetComp("login").GetAssetList() {
			app.GET("/assets" + path, controller.Assert)
		}

		authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func (ctx *context.Context)  {
			ctx.Write(302, map[string]string{
				"Location": prefix + "/login",
			}, ``)
		}).SetPermissionDenyCallback(func (ctx *context.Context)  {
			controller.ShowErrorPage(ctx, "permission denied")
		})

		app.Group("", authenticator.Middleware)
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
	}

	return app
}