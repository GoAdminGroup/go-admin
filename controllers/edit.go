package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/menu"
	"goAdmin/models"
	"goAdmin/template"
	"goAdmin/modules/file"
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

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // 管理员管理编辑
		EditManager((*form).Value)
	} else if prefix == "roles" { // 管理员角色管理编辑
		EditRole((*form).Value)
	} else {
		models.GlobalTableList[prefix].UpdateDataFromDatabase(prefix, (*form).Value)
	}

	models.RefreshGlobalTableList()

	// TODO: 增加反馈

	previous := string(ctx.FormValue("_previous_")[:])
	ctx.Response.SetStatusCode(302)
	ctx.Response.Header.Add("Location", previous)
}
