package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template/adminlte/components"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"net/http"
	"github.com/chenhg5/go-admin/modules/menu"
	"strings"
	"path"
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
	editUrl = AssertRootUrl + "/info/" + prefix + "/edit?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)
	newUrl := AssertRootUrl + "/info/" + prefix + "/new?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)

	tmpl := components.GetTemplate(string(ctx.Request.Header.Get("X-PJAX")) == "true")

	menu.GlobalMenu.SetActiveClass(ctx.Path())

	dataTable := components.DataTable().SetInfoList(infoList).SetThead(thead).SetEditUrl(editUrl).SetNewUrl(newUrl)
	table := dataTable.GetContent()

	box := components.Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(paninator.GetContent()).
		GetContent()

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, "layout", components.Page{
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
		AssertRootUrl: AssertRootUrl,
	})
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())
}


func Assert(ctx *context.Context) {
	filepath := "../../../template/adminlte/resource" + strings.Replace(ctx.Request.URL.Path, AssertRootUrl, "", -1)
	data, err := Asset(filepath)
	fileSuffix := path.Ext(filepath)
	fileSuffix = strings.Replace(fileSuffix, ".", "", -1)

	var contentType = ""
	if fileSuffix == "css" || fileSuffix == "js" {
		contentType = "text/" + fileSuffix + "; charset=utf-8"
	} else {
		contentType = "image/" + fileSuffix
	}

	if err != nil {
		ctx.Write(http.StatusNotFound, map[string]string{}, "")
	} else {
		ctx.Write(http.StatusOK, map[string]string{
			"content-type": contentType,
		}, string(data))
	}
}