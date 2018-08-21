package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/components"
	"goAdmin/components/menu"
	"goAdmin/modules/auth"
	"goAdmin/components/adminlte"
)

// 设置页面内容
func SetPageContent(ctx *fasthttp.RequestCtx, c func() components.Panel) {
	user := ctx.UserValue("user").(auth.User)

	panel := c()

	tmpl := adminlte.GetTemplate(string(ctx.Request.Header.Peek("X-PJAX")) == "true")

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	tmpl.ExecuteTemplate(ctx.Response.BodyWriter(), "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel: panel,
	})

}
