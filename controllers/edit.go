package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/menu"
	"goAdmin/models"
	"goAdmin/template"
)

// 显示表单
func ShowForm(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue("cur_user").(auth.User)
	prefix := ctx.UserValue("prefix").(string)

	buffer := new(bytes.Buffer)

	id := string(ctx.QueryArgs().Peek("id")[:])

	page := string(ctx.QueryArgs().Peek("page")[:])
	if page == "" {
		page = "1"
	}
	pageSize := string(ctx.QueryArgs().Peek("pageSize")[:])
	if pageSize == "" {
		pageSize = "10"
	}

	formData, title, description := models.GlobalTableList[prefix].GetDataFromDatabaseWithId(prefix, id)

	url := "/edit/" + prefix + "?id=" + id
	previous := "/info/" + prefix + "?page=" + page + "&pageSize=" + pageSize

	if string(ctx.Request.Header.Peek("X-PJAX")[:]) == "true" {
		template.EditPanelPjax(formData, url, previous, id, title, description, buffer)
	} else {
		template.EditPanel(formData, url, previous, id, (*menu.GlobalMenu).GlobalMenuList, title, description, user, buffer)
	}

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 编辑数据
func EditForm(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	prefix := ctx.UserValue("prefix").(string)

	previous := string(ctx.FormValue("_previous_")[:])

	form, _ := ctx.MultipartForm()

	models.GlobalTableList[prefix].UpdateDataFromDatabase(prefix, (*form).Value)

	// TODO: 增加反馈

	ctx.Response.SetStatusCode(302)
	ctx.Response.Header.Add("Location", previous)
}
