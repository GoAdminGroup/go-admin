package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"path"
	"strings"
)

// 显示列表
func ShowInfo(ctx *context.Context) {

	prefix := ctx.Query("prefix")
	panel := models.TableList[prefix]

	params := models.GetParam(ctx.Request.URL.Query())

	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	panelInfo := panel.GetDataFromDatabase(ctx.Path(), params)

	var box template2.HTML
	if prefix != "op" {
		dataTable := template.Get(Config.THEME).
			DataTable().
			SetInfoList(panelInfo.InfoList).
			SetFilters(panel.GetFiltersMap()).
			SetInfoUrl(Config.PREFIX + "/info/" + prefix).
			SetThead(panelInfo.Thead)

		if panelInfo.CanAdd {
			dataTable.SetNewUrl(newUrl)
		}
		if panelInfo.Editable {
			dataTable.SetEditUrl(editUrl)
		}
		if panelInfo.Deletable {
			dataTable.SetDeleteUrl(deleteUrl)
		}

		table := dataTable.GetContent()

		box = template.Get(Config.THEME).Box().
			SetBody(table).
			SetHeader(dataTable.GetDataTableHeader() + panel.GetInfo().HeaderHtml).
			WithHeadBorder(false).
			SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
			GetContent()
	} else {
		dataTable := template.Get(Config.THEME).
			Table().
			SetType("table").
			SetThead(panelInfo.Thead).
			SetInfoList(panelInfo.InfoList)

		table := dataTable.GetContent()

		box = template.Get(Config.THEME).Box().
			SetBody(table).
			WithHeadBorder(false).
			SetHeader(panel.GetInfo().HeaderHtml).
			SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
			GetContent()
	}

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())
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
		logger.Error("asset err", err)
		ctx.Write(http.StatusNotFound, map[string]string{}, "")
	} else {
		ctx.Write(http.StatusOK, map[string]string{
			"content-type": contentType,
		}, string(data))
	}
}
