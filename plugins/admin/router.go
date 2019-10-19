package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/template"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	route := app.Group(prefix, globalErrorHandler)

	// auth
	route.GET("/login", controller.ShowLogin)
	route.POST("/signin", controller.Auth)

	// auto install
	route.GET("/install", controller.ShowInstall)
	route.POST("/install/database/check", controller.CheckDatabase)

	for _, path := range template.Get(config.Get().Theme).GetAssetList() {
		route.GET("/assets"+path, controller.Assets)
	}

	for _, path := range template.GetAssetLists() {
		route.GET("/assets"+path, controller.Assets)
	}

	authRoute := route.Group("/", auth.Middleware)

	// auth
	authRoute.GET("/logout", controller.Logout)

	// menus
	authRoute.POST("/menu/delete", guard.MenuDelete, controller.DeleteMenu)
	authRoute.POST("/menu/new", guard.MenuNew, controller.NewMenu)
	authRoute.POST("/menu/edit", guard.MenuEdit, controller.EditMenu)
	authRoute.POST("/menu/order", controller.MenuOrder)
	authRoute.GET("/menu", controller.ShowMenu)
	authRoute.GET("/menu/edit/show", controller.ShowEditMenu)

	//authRoute.GET("/menu/new", controller.ShowMenu) // TODO: this will cause a bug of the tire

	// add delete modify query
	authRoute.GET("/info/:__prefix/edit", guard.ShowForm, controller.ShowForm)
	authRoute.GET("/info/:__prefix/new", guard.ShowNewForm, controller.ShowNewForm)
	authRoute.POST("/edit/:__prefix", guard.EditForm, controller.EditForm)
	authRoute.POST("/new/:__prefix", guard.NewForm, controller.NewForm)
	authRoute.POST("/delete/:__prefix", guard.Delete, controller.Delete)
	authRoute.POST("/export/:__prefix", guard.Export, controller.Export)
	authRoute.GET("/info/:__prefix", controller.ShowInfo)

	return app
}

func globalErrorHandler(ctx *context.Context) {
	defer controller.GlobalDeferHandler(ctx)
	ctx.Next()
}
