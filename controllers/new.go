package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/menu"
	"goAdmin/template"
	"goAdmin/models"
)

// 显示新建表单
func ShowNewForm(ctx *fasthttp.RequestCtx) {
	defer GlobalDeferHandler(ctx)

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

	url := "/new/" + prefix + "?id=" + id
	previous := "/info/" + prefix + "?page=" + page + "&pageSize=" + pageSize

	if string(ctx.Request.Header.Peek("X-PJAX")[:]) == "true" {
		template.NewPanelPjax(models.GlobalTableList[prefix].Form.FormList, url, previous, id, models.GlobalTableList[prefix].Form.Title, models.GlobalTableList[prefix].Form.Description, buffer)
	} else {
		template.NewPanel(models.GlobalTableList[prefix].Form.FormList, url, previous, id, (*menu.GlobalMenu).GlobalMenuList, models.GlobalTableList[prefix].Form.Title, models.GlobalTableList[prefix].Form.Description, buffer)
	}

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 新建数据
func NewForm(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	prefix := ctx.UserValue("prefix").(string)

	previous := string(ctx.FormValue("_previous_")[:])

	form, _ := ctx.MultipartForm()

	models.GlobalTableList[prefix].InsertDataFromDatabase(prefix, (*form).Value)

	// TODO: 增加反馈

	ctx.Response.SetStatusCode(302)
	ctx.Response.Header.Add("Location", previous)
}
