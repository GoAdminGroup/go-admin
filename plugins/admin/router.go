package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/template"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	app.Group(prefix, GlobalErrorHandler)
	{
		// auth
		app.GET("/login", controller.ShowLogin)
		app.POST("/signin", controller.Auth)

		// auto install
		app.GET("/install", controller.ShowInstall)
		app.POST("/install/database/check", controller.CheckDatabase)

		for _, path := range template.Get("adminlte").GetAssetList() {
			app.GET("/assets"+path, controller.Assert)
		}

		for _, path := range template.GetComp("login").GetAssetList() {
			app.GET("/assets"+path, controller.Assert)
		}

		authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
			ctx.Write(302, map[string]string{
				"Location": prefix + "/login",
			}, ``)
		}).SetPermissionDenyCallback(func(ctx *context.Context) {
			controller.ShowErrorPage(ctx, "permission denied")
		})

		app.Group("", authenticator.Middleware)
		{
			// auth
			app.GET("/logout", controller.Logout)

			// menus
			app.GET("/menu", controller.ShowMenu)
			app.POST("/menu/delete", controller.DeleteMenu)
			app.POST("/menu/new", controller.NewMenu)
			//app.GET("/menu/new", controller.ShowMenu) // TODO: this will cause a bug of the tire
			app.POST("/menu/edit", controller.EditMenu)
			app.GET("/menu/edit/show", controller.ShowEditMenu)
			app.POST("/menu/order", controller.MenuOrder)

			// add delete modify query
			app.GET("/info/:prefix", controller.ShowInfo)
			app.GET("/info/:prefix/edit", controller.ShowForm)
			app.GET("/info/:prefix/new", controller.ShowNewForm)
			app.POST("/edit/:prefix", controller.EditForm)
			app.POST("/delete/:prefix", controller.DeleteData)
			app.POST("/new/:prefix", controller.NewForm)
		}
	}

	return app
}

func GlobalErrorHandler(h context.Handler) context.Handler {
	return func(ctx *context.Context) {
		defer controller.GlobalDeferHandler(ctx)
		h(ctx)
		return
	}
}
