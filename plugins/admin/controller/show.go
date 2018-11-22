package controller

import (
	"fmt"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"net/http"
	"path"
	"strings"
)

// 显示列表
func ShowInfo(ctx *context.Context) {

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Query("prefix")

	params := models.GetParam(ctx.Request.URL.Query())

	panelInfo := models.TableList[prefix].GetDataFromDatabase(ctx.Path(), params)

	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	menu.GlobalMenu.SetActiveClass(ctx.Path())

	dataTable := template.Get(Config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
		SetFilters(models.TableList[prefix].GetFiltersMap()).
		SetInfoUrl(Config.PREFIX + "/info/" + prefix).
		SetThead(panelInfo.Thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(panelInfo.Paginator.GetContent()).
		GetContent()

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, Config)

	ctx.WriteString(buf.String())
}

func Assert(ctx *context.Context) {
	filepath := "template/adminlte/resource" + strings.Replace(ctx.Path(), Config.PREFIX, "", 1)
	data, err := template.Get(Config.THEME).GetAsset(filepath)
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