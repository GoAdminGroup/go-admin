package main

import (
	"github.com/buaazp/fasthttprouter"
	"goAdmin/controllers"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
)

func InitRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()
	router.GET("/login", controller.ShowLogin)
	router.POST("/signin", controller.Auth)
	router.GET("/install", controller.ShowInstall)

	router.POST("/logout", AuthMiddleware(controller.Logout))
	router.GET("/menu", AuthMiddleware(controller.ShowMenu))
	router.POST("/menu/delete", AuthMiddleware(controller.DeleteMenu))
	router.POST("/menu/new", AuthMiddleware(controller.NewMenu))
	router.GET("/menu/new", AuthMiddleware(controller.ShowMenu))
	router.POST("/menu/edit", AuthMiddleware(controller.EditMenu))
	router.GET("/menu/edit/show", AuthMiddleware(controller.ShowEditMenu))

	router.GET("/info/:prefix", AuthMiddleware(controller.ShowInfo))
	router.GET("/info/:prefix/edit", AuthMiddleware(controller.ShowForm))
	router.GET("/info/:prefix/new", AuthMiddleware(controller.ShowNewForm))
	router.POST("/edit/:prefix", AuthMiddleware(controller.EditForm))
	router.POST("/delete/:prefix", AuthMiddleware(controller.DeleteData))
	router.POST("/new/:prefix", AuthMiddleware(controller.NewForm))

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