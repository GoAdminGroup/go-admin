package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/modules/parameter"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"path"
	"strings"
)

func ShowInfo(ctx *context.Context) {

	prefix := ctx.Query("prefix")
	panel := table.List[prefix]

	params := parameter.GetParam(ctx.Request.URL.Query())

	editUrl := config.Url("/info/" + prefix + "/edit" + params.GetRouteParamStr())
	newUrl := config.Url("/info/" + prefix + "/new" + params.GetRouteParamStr())
	deleteUrl := config.Url("/delete/" + prefix)

	panelInfo := panel.GetDataFromDatabase(ctx.Path(), params)

	var box template2.HTML
	if prefix != "op" {
		dataTable := aDataTable().
			SetInfoList(panelInfo.InfoList).
			SetFilters(panel.GetFiltersMap()).
			SetInfoUrl(config.Url("/info/" + prefix)).
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

		box = aBox().
			SetBody(dataTable.GetContent()).
			SetHeader(dataTable.GetDataTableHeader() + panel.GetInfo().HeaderHtml).
			WithHeadBorder(false).
			SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
			GetContent()
	} else {
		dataTable := aTemplate().
			Table().
			SetType("table").
			SetThead(panelInfo.Thead).
			SetInfoList(panelInfo.InfoList)

		box = aBox().
			SetBody(dataTable.GetContent()).
			WithHeadBorder(false).
			SetHeader(panel.GetInfo().HeaderHtml).
			SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
			GetContent()
	}

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))
	ctx.Html(http.StatusOK, buf.String())
}

func Assert(ctx *context.Context) {
	filepath := config.UrlRemovePrefix(ctx.Path())
	data, err := aTemplate().GetAsset(filepath)

	if err != nil {
		data, err = loginComponent().GetAsset(filepath)
		if err != nil {
			logger.Error("asset err", err)
			ctx.Write(http.StatusNotFound, map[string]string{}, "")
			return
		}
	}

	fileSuffix := path.Ext(filepath)
	fileSuffix = strings.Replace(fileSuffix, ".", "", -1)

	var contentType = ""
	if fileSuffix == "css" || fileSuffix == "js" {
		contentType = "text/" + fileSuffix + "; charset=utf-8"
	} else {
		contentType = "image/" + fileSuffix
	}

	ctx.Write(http.StatusOK, map[string]string{
		"content-type": contentType,
	}, string(data))
}
