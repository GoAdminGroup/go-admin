package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/models"
	"goAdmin/auth"
	"goAdmin/menu"
	"bytes"
	"goAdmin/template"
)

// 设置页面内容
func SetPageContent(ctx *fasthttp.RequestCtx, c func() models.Page) {
	user := ctx.UserValue("cur_user").(auth.User)
	path := string(ctx.Path())

	page := c()

	menu.GlobalMenu.SetActiveClass(path)

	buffer := new(bytes.Buffer)

	template.BaseContent(page.Content, (*menu.GlobalMenu).GlobalMenuList, page.Title, page.Description, user, buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}