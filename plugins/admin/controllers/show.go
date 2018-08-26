package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/modules/auth"
	"goAdmin/components/menu"
	"goAdmin/plugins/admin/models"
	"goAdmin/components"
	"goAdmin/app"
	"goAdmin/components/adminlte"
)

// 显示列表
func ShowInfo(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue("user").(auth.User)

	path := string(ctx.Path())
	prefix := ctx.UserValue("prefix").(string)

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

	thead, infoList, _, title, description := models.GlobalTableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      string(page),
		"path":      string(path),
		"sortField": string(sortField),
		"sortType":  string(sortType),
		"prefix":    prefix,
		"pageSize":  string(pageSize),
	})

	var (
		editUrl string
		//newUrl  string
	)
	editUrl = "/info/" + prefix + "/edit?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)
	newUrl := "/info/" + prefix + "/new?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)

	tmpl := adminlte.GetTemplate(string(ctx.Request.Header.Peek("X-PJAX")) == "true")

	menu.GlobalMenu.SetActiveClass(path)

	dataTable := app.GetComponents().DataTable().SetInfoList(infoList).SetThead(thead).SetEditUrl(editUrl).SetNewUrl(newUrl)
	table := dataTable.GetContent()
	paninator := app.GetComponents().Paninator().GetContent()

	box := app.GetComponents().Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(paninator).
		GetContent()

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(ctx.Response.BodyWriter(), "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel: components.Panel{
			Content:     box,
			Description: description,
			Title:       title,
		},
	})
}
