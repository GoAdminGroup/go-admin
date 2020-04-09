package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/template"
)

// initRouter initialize the router and return the context.
func (admin *Admin) initRouter() *Admin {
	app := context.NewApp()

	route := app.Group(config.Prefix(), admin.globalErrorHandler)

	// auth
	route.GET(config.GetLoginUrl(), admin.handler.ShowLogin)
	route.POST("/signin", admin.handler.Auth)

	// auto install
	route.GET("/install", admin.handler.ShowInstall)
	route.POST("/install/database/check", admin.handler.CheckDatabase)

	for _, path := range template.Get(config.GetTheme()).GetAssetList() {
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
	authRoute.POST("/menu/delete", admin.guardian.MenuDelete, admin.handler.DeleteMenu).Name("menu_delete")
	authRoute.POST("/menu/new", admin.guardian.MenuNew, admin.handler.NewMenu).Name("menu_new")
	authRoute.POST("/menu/edit", admin.guardian.MenuEdit, admin.handler.EditMenu).Name("menu_edit")
	authRoute.POST("/menu/order", admin.handler.MenuOrder).Name("menu_order")
	authRoute.GET("/menu", admin.handler.ShowMenu).Name("menu")
	authRoute.GET("/menu/edit/show", admin.handler.ShowEditMenu).Name("menu_edit_show")
	authRoute.GET("/menu/new", admin.handler.ShowNewMenu).Name("menu_new_show")

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

	route.ANY("/operation/:__goadmin_op_id", auth.Middleware(admin.conn), admin.handler.Operation)

	if config.GetOpenAdminApi() {

		// crud json apis
		apiRoute := route.Group("/api", auth.Middleware(admin.conn), admin.guardian.CheckPrefix)
		apiRoute.GET("/list/:__prefix", admin.handler.ApiList).Name("api_info")
		apiRoute.GET("/detail/:__prefix", admin.handler.ApiDetail).Name("api_detail")
		apiRoute.POST("/delete/:__prefix", admin.guardian.Delete, admin.handler.Delete).Name("api_delete")
		apiRoute.POST("/update/:__prefix", admin.guardian.EditForm, admin.handler.ApiUpdate).Name("api_edit")
		apiRoute.GET("/update/form/:__prefix", admin.guardian.ShowForm, admin.handler.ApiUpdateForm).Name("api_show_edit")
		apiRoute.POST("/create/:__prefix", admin.guardian.NewForm, admin.handler.ApiCreate).Name("api_new")
		apiRoute.GET("/create/form/:__prefix", admin.guardian.ShowNewForm, admin.handler.ApiCreateForm).Name("api_show_new")
		apiRoute.POST("/export/:__prefix", admin.guardian.Export, admin.handler.Export).Name("api_export")

		apiRoute.POST("/update/:__prefix", admin.guardian.Update, admin.handler.Update).Name("api_update")
	}

	admin.app = app
	return admin
}

func (admin *Admin) globalErrorHandler(ctx *context.Context) {
	defer admin.handler.GlobalDeferHandler(ctx)
	response.OffLineHandler(ctx)
	ctx.Next()
}
