package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/template"
)

// initRouter initialize the router and return the context.
func (admin *Admin) initRouter(prefix string) *Admin {
	app := context.NewApp()

	route := app.Group(prefix, admin.globalErrorHandler)

	// auth
	route.GET("/login", admin.handler.ShowLogin)
	route.POST("/signin", admin.handler.Auth)

	// auto install
	route.GET("/install", admin.handler.ShowInstall)
	route.POST("/install/database/check", admin.handler.CheckDatabase)

	for _, path := range template.Get(config.Get().Theme).GetAssetList() {
		route.GET("/assets"+path, admin.handler.Assets)
	}

	for _, path := range template.GetComponentAssetLists() {
		route.GET("/assets"+path, admin.handler.Assets)
	}

	authRoute := route.Group("/", auth.Middleware(admin.conn))

	// auth
	authRoute.GET("/logout", admin.handler.Logout)

	authPrefixRoute := route.Group("/", auth.Middleware(admin.conn), admin.guardian.CheckPrefix)

	// menus
	authRoute.POST("/menu/delete", admin.guardian.MenuDelete, admin.handler.DeleteMenu)
	authRoute.POST("/menu/new", admin.guardian.MenuNew, admin.handler.NewMenu)
	authRoute.POST("/menu/edit", admin.guardian.MenuEdit, admin.handler.EditMenu)
	authRoute.POST("/menu/order", admin.handler.MenuOrder)
	authRoute.GET("/menu", admin.handler.ShowMenu)
	authRoute.GET("/menu/edit/show", admin.handler.ShowEditMenu)
	authRoute.GET("/menu/new", admin.handler.ShowNewMenu)

	// add delete modify query
	authPrefixRoute.GET("/info/:__prefix/detail", admin.handler.ShowDetail).Name("detail")
	authPrefixRoute.GET("/info/:__prefix/edit", admin.guardian.ShowForm, admin.handler.ShowForm).Name("show_edit")
	authPrefixRoute.GET("/info/:__prefix/new", admin.guardian.ShowNewForm, admin.handler.ShowNewForm).Name("show_new")
	authPrefixRoute.POST("/edit/:__prefix", admin.guardian.EditForm, admin.handler.EditForm).Name("edit")
	authPrefixRoute.POST("/new/:__prefix", admin.guardian.NewForm, admin.handler.NewForm).Name("new")
	authPrefixRoute.POST("/delete/:__prefix", admin.guardian.Delete, admin.handler.Delete).Name("delete")
	authPrefixRoute.POST("/export/:__prefix", admin.guardian.Export, admin.handler.Export).Name("export")
	authPrefixRoute.GET("/info/:__prefix", admin.handler.ShowInfo).Name("info")

	authPrefixRoute.POST("/update/:__prefix", admin.guardian.Update, admin.handler.Update).Name("update")

	admin.app = app
	return admin
}

func (admin *Admin) globalErrorHandler(ctx *context.Context) {
	defer admin.handler.GlobalDeferHandler(ctx)
	ctx.Next()
}
