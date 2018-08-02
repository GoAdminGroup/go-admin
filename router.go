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

	return router
}

func AuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		var (
			authOk   bool
			filterOk bool
			user     auth.User
		)

		if user, authOk, filterOk = auth.Filter(ctx); authOk && filterOk {
			ctx.SetUserValue("cur_user", user)
			h(ctx)
			return
		}

		if !authOk {
			ctx.Response.Header.Add("Location", "/login")
			ctx.Response.SetStatusCode(302)
			return
		}

		if !filterOk {
			ctx.Response.SetStatusCode(403)
			ctx.WriteString(`{"code":403, "msg":"权限不够"}`)
			return
		}
	})
}
