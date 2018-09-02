package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"net/http"
	"github.com/chenhg5/go-admin/modules/menu"
	"strings"
	"path"
	"fmt"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
)

// 显示列表
func ShowInfo(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Request.URL.Query().Get("prefix")

	page := ctx.Request.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageSize := ctx.Request.URL.Query().Get("pageSize")
	if pageSize == "" {
		pageSize = "10"
	}

	sortField := ctx.Request.URL.Query().Get("sort")
	if sortField == "" {
		sortField = "id"
	}
	sortType := ctx.Request.URL.Query().Get("sort_type")
	if sortType == "" {
		sortType = "desc"
	}

	thead, infoList, paninator, title, description := models.GlobalTableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      page,
		"path":      ctx.Path(),
		"sortField": sortField,
		"sortType":  sortType,
		"prefix":    prefix,
		"pageSize":  pageSize,
	})

	var (
		editUrl string
		//newUrl  string
	)
	editUrl = Config.ADMIN_PREFIX + "/info/" + prefix + "/edit?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)
	newUrl := Config.ADMIN_PREFIX + "/info/" + prefix + "/new?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)

	tmpl, tmplName := template.Get("adminlte").GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

	menu.GlobalMenu.SetActiveClass(ctx.Path())

	dataTable := template.Get(Config.THEME).DataTable().SetInfoList(infoList).SetThead(thead).SetEditUrl(editUrl).SetNewUrl(newUrl)
	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(paninator.GetContent()).
		GetContent()

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel: types.Panel{
			Content:     box,
			Description: description,
			Title:       title,
		},
		AssertRootUrl: Config.ADMIN_PREFIX,
	})
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())
}


func Assert(ctx *context.Context) {
	filepath := "../../../template/adminlte/resource" + strings.Replace(ctx.Request.URL.Path, Config.ADMIN_PREFIX, "", 1)
	data, err := template.Get("adminlte").GetAsset(filepath)
	fileSuffix := path.Ext(filepath)
	fileSuffix = strings.Replace(fileSuffix, ".", "", -1)

	var contentType = ""
	if fileSuffix == "css" || fileSuffix == "js" {
		contentType = "text/" + fileSuffix + "; charset=utf-8"
	} else {
		contentType = "image/" + fileSuffix
	}

	if err != nil {
		fmt.Println("asset err", err)
		ctx.Write(http.StatusNotFound, map[string]string{}, "")
	} else {
		ctx.Write(http.StatusOK, map[string]string{
			"content-type": contentType,
		}, string(data))
	}
}