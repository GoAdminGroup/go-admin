package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/template"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	route := app.Group(prefix, GlobalErrorHandler)

	// auth
	route.GET("/login", controller.ShowLogin)
	route.POST("/signin", controller.Auth)

	// auto install
	route.GET("/install", controller.ShowInstall)
	route.POST("/install/database/check", controller.CheckDatabase)

	for _, path := range template.Get("adminlte").GetAssetList() {
		route.GET("/assets"+path, controller.Assert)
	}

	for _, path := range template.GetComp("login").GetAssetList() {
		route.GET("/assets"+path, controller.Assert)
	}

	authRoute := route.Group("/", auth.Middleware)

	// auth
	authRoute.GET("/logout", controller.Logout)

	// menus
	authRoute.GET("/menu", controller.ShowMenu)
	authRoute.POST("/menu/delete", controller.DeleteMenu)
	authRoute.POST("/menu/new", controller.NewMenu)
	//authRoute.GET("/menu/new", controller.ShowMenu) // TODO: this will cause a bug of the tire
	authRoute.POST("/menu/edit", controller.EditMenu)
	authRoute.GET("/menu/edit/show", controller.ShowEditMenu)
	authRoute.POST("/menu/order", controller.MenuOrder)

	// add delete modify query
	authRoute.GET("/info/:prefix", controller.ShowInfo)
	authRoute.GET("/info/:prefix/edit", controller.ShowForm)
	authRoute.GET("/info/:prefix/new", controller.ShowNewForm)
	authRoute.POST("/edit/:prefix", controller.EditForm)
	authRoute.POST("/delete/:prefix", controller.Delete)
	authRoute.POST("/new/:prefix", controller.NewForm)

	return app
}

func GlobalErrorHandler(ctx *context.Context) {
	defer controller.GlobalDeferHandler(ctx)
	ctx.Next()
	return
}
