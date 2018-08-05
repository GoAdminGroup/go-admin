package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/menu"
	"goAdmin/models"
	"goAdmin/modules/file"
	"goAdmin/template"
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

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // 管理员管理新建
		NewManager((*form).Value)
	} else if prefix == "roles" { // 管理员角色管理新建
		NewRole((*form).Value)
	} else {
		models.GlobalTableList[prefix].InsertDataFromDatabase(prefix, (*form).Value)
	}

	models.RefreshGlobalTableList()

	// TODO: 增加反馈

	previous := string(ctx.FormValue("_previous_")[:])
	ctx.Response.SetStatusCode(302)
	ctx.Response.Header.Add("Location", previous)
}
