package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/menu"
	"goAdmin/models"
	"goAdmin/modules/language"
	"goAdmin/template"
)

// 显示列表
func ShowInfo(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue("cur_user").(auth.User)
	prefix := ctx.UserValue("prefix").(string)
	path := string(ctx.Path())

	page := ctx.QueryArgs().Peek("page")
	if len(page) == 0 {
		page = []byte("1")
	}
	pageSize := ctx.QueryArgs().Peek("pageSize")
	if len(pageSize) == 0 {
		pageSize = []byte("10")
	}

	sortField := ctx.QueryArgs().Peek("sort")
	if len(sortField) == 0 {
		sortField = []byte("id")
	}
	sortType := ctx.QueryArgs().Peek("sort_type")
	if len(sortType) == 0 {
		sortType = []byte("desc")
	}

	thead, infoList, paginator, title, description := models.GlobalTableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      string(page),
		"path":      string(path),
		"sortField": string(sortField),
		"sortType":  string(sortType),
		"prefix":    prefix,
		"pageSize":  string(pageSize),
	})

	menu.GlobalMenu.SetActiveClass(path)

	buffer := new(bytes.Buffer)

	if string(ctx.Request.Header.Peek("X-PJAX")) == "true" {
		template.InfoListPjax(infoList, (*menu.GlobalMenu).GlobalMenuList, thead, paginator, title, description, buffer)
	} else {
		template.InfoList(infoList, (*menu.GlobalMenu).GlobalMenuList, thead, paginator, title, description, user, language.Lang, buffer)
	}

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
