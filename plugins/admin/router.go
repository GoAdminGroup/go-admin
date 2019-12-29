package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/template"
)

// InitRouter initialize the router and return the context.
func InitRouter(prefix string, srv service.List) *context.App {
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

	for _, path := range template.GetComponentAssetLists() {
		route.GET("/assets"+path, controller.Assets)
	}

	authRoute := route.Group("/", auth.Middleware(db.GetConnection(srv)))

	// auth
	authRoute.GET("/logout", controller.Logout)

	conn := db.GetConnection(srv)

	// menus
	authRoute.POST("/menu/delete", guard.MenuDelete(conn), controller.DeleteMenu)
	authRoute.POST("/menu/new", guard.MenuNew(srv), controller.NewMenu)
	authRoute.POST("/menu/edit", guard.MenuEdit(srv), controller.EditMenu)
	authRoute.POST("/menu/order", controller.MenuOrder)
	authRoute.GET("/menu", controller.ShowMenu)
	authRoute.GET("/menu/edit/show", controller.ShowEditMenu)
	authRoute.GET("/menu/new", controller.ShowNewMenu)

	// add delete modify query
	authRoute.GET("/info/:__prefix/edit", guard.ShowForm(conn), controller.ShowForm)
	authRoute.GET("/info/:__prefix/new", guard.ShowNewForm(conn), controller.ShowNewForm)
	authRoute.POST("/edit/:__prefix", guard.EditForm(srv), controller.EditForm)
	authRoute.POST("/new/:__prefix", guard.NewForm(srv), controller.NewForm)
	authRoute.POST("/delete/:__prefix", guard.Delete(conn), controller.Delete)
	authRoute.POST("/export/:__prefix", guard.Export(conn), controller.Export)
	authRoute.GET("/info/:__prefix", controller.ShowInfo)
	authRoute.GET("/info/:__prefix/detail", controller.ShowDetail)

	authRoute.POST("/update/:__prefix", guard.Update, controller.Update)

	return app
}

func globalErrorHandler(ctx *context.Context) {
	defer controller.GlobalDeferHandler(ctx)
	ctx.Next()
}
