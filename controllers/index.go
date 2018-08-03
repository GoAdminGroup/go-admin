package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/menu"
	"bytes"
	"goAdmin/template"
	"goAdmin/components"
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

	box := components.GetBox().SetUrl("/").SetTitle("用户总数").SetValue("1000").GetContent()

	col1 := components.Col.GetContent(box)
	col2 := components.Col.GetContent(box)
	col3 := components.Col.GetContent(box)
	col4 := components.Col.GetContent(box)

	row := components.Row.GetContent(col1 + col2 + col3 + col4)

	template.BaseContent(row, (*menu.GlobalMenu).GlobalMenuList, title, description, user, buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
