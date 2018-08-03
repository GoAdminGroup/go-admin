package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/menu"
	"bytes"
	"goAdmin/template"
)

// 显示仪表盘
func ShowDashboard(ctx *fasthttp.RequestCtx)  {
	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue("cur_user").(auth.User)
	path := string(ctx.Path())

	title := "仪表盘"
	description := "仪表盘"

	menu.GlobalMenu.SetActiveClass(path)

	buffer := new(bytes.Buffer)

	template.BaseContent("你好", (*menu.GlobalMenu).GlobalMenuList, title, description, user, buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
