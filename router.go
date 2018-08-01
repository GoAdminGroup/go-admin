package main

import (
	"github.com/buaazp/fasthttprouter"
	"goAdmin/controllers"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
)

func InitRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()

	// 授权认证
	router.GET("/login", controller.ShowLogin)
	router.POST("/signin", controller.Auth)
	router.POST("/logout", AuthMiddleware(controller.Logout))

	// 菜单管理
	router.GET("/menu", AuthMiddleware(controller.ShowMenu))
	router.POST("/menu/delete", AuthMiddleware(controller.DeleteMenu))
	router.POST("/menu/new", AuthMiddleware(controller.NewMenu))
	router.GET("/menu/new", AuthMiddleware(controller.ShowMenu))
	router.POST("/menu/edit", AuthMiddleware(controller.EditMenu))
	router.GET("/menu/edit/show", AuthMiddleware(controller.ShowEditMenu))

	// 增删改查管理
	router.GET("/info/:prefix", AuthMiddleware(controller.ShowInfo))
	router.GET("/info/:prefix/edit", AuthMiddleware(controller.ShowForm))
	router.GET("/info/:prefix/new", AuthMiddleware(controller.ShowNewForm))
	router.POST("/edit/:prefix", AuthMiddleware(controller.EditForm))
	router.POST("/delete/:prefix", AuthMiddleware(controller.DeleteData))
	router.POST("/new/:prefix", AuthMiddleware(controller.NewForm))

	// 自动化安装
	router.GET("/install", controller.ShowInstall)
	router.POST("/install/database/check", controller.CheckDatabase)

	// 管理员管理
	router.GET("/manager/list", controller.ShowManager)
	router.POST("/manager/new", controller.NewManager)
	router.POST("/manager/delete", controller.DeleteManager)
	router.POST("/manager/edit", controller.EditManager)

	// 管理权限管理
	router.GET("/manager/rules/list", controller.ShowManagerRules)
	router.POST("/manager/rules/new", controller.NewManagerRules)
	router.POST("/manager/rules/delete", controller.DeleteManagerRules)
	router.POST("/manager/rules/edit", controller.EditManagerRules)

	// 管理角色管理
	router.GET("/manager/roles/list", controller.ShowManagerRoles)
	router.POST("/manager/roles/new", controller.NewManagerRoles)
	router.POST("/manager/roles/delete", controller.DeleteManagerRoles)
	router.POST("/manager/roles/edit", controller.EditManagerRoles)

	// 操作日志管理
	router.GET("/operation/log/list", controller.ShowOperationLog)
	router.POST("/operation/log/new", controller.NewOperationLog)
	router.POST("/operation/log/delete", controller.DeleteOperationLog)
	router.POST("/operation/log/edit", controller.EditOperationLog)

	return router
}


func AuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		if user, ok := auth.Filter(ctx); ok {
			ctx.SetUserValue("cur_user", user)
			h(ctx)
			return
		}

		ctx.Response.Header.Add("Location", "/login")
		ctx.Response.SetStatusCode(302)
	})
}