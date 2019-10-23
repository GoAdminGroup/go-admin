package controller

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

func ShowInfo(ctx *context.Context) {

	prefix := ctx.Query("__prefix")
	panel := table.List[prefix]

	params := parameter.GetParam(ctx.Request.URL.Query())

	editUrl := modules.AorB(panel.GetEditable() && !panel.GetInfo().IsHideEditButton,
		config.Url("/info/"+prefix+"/edit"+params.GetRouteParamStr()), "")
	deleteUrl := modules.AorB(panel.GetDeletable() && !panel.GetInfo().IsHideDeleteButton,
		config.Url("/delete/"+prefix), "")
	exportUrl := modules.AorB(panel.GetExportable() && !panel.GetInfo().IsHideExportButton,
		config.Url("/export/"+prefix+params.GetRouteParamStr()), "")
	newUrl := modules.AorB(panel.GetCanAdd() && !panel.GetInfo().IsHideNewButton,
		config.Url("/info/"+prefix+"/new"+params.GetRouteParamStr()), "")
	infoUrl := config.Url("/info/" + prefix)

	buf := showTable(ctx, panel, ctx.Path(), params, exportUrl, newUrl, deleteUrl, infoUrl, editUrl)
	ctx.Html(http.StatusOK, buf.String())
}

func showTable(ctx *context.Context, panel table.Table, path string, params parameter.Parameters,
	exportUrl, newUrl, deleteUrl, infoUrl, editUrl string) *bytes.Buffer {

	panelInfo, err := panel.GetDataFromDatabase(path, params)

	if err != nil {
		tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
		user := auth.Auth(ctx)
		alert := aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		return template.Execute(tmpl, tmplName, user, types.Panel{
			Content:     alert,
			Description: language.Get("error"),
			Title:       language.Get("error"),
		}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))
	}

	var (
		body      template2.HTML
		dataTable types.DataTableAttribute
	)

	if panel.GetInfo().TabGroups.Valid() {

		dataTable = aDataTable().
			SetThead(panelInfo.Thead).
			SetDeleteUrl(deleteUrl).
			SetNewUrl(newUrl).
			SetExportUrl(exportUrl)

		var (
			tabsHtml    = make([]map[string]template2.HTML, len(panel.GetInfo().TabHeaders))
			infoListArr = panelInfo.InfoList.GroupBy(panel.GetInfo().TabGroups)
			theadArr    = panelInfo.Thead.GroupBy(panel.GetInfo().TabGroups)
		)
		for key, header := range panel.GetInfo().TabHeaders {
			tabsHtml[key] = map[string]template2.HTML{
				"title": template2.HTML(header),
				"content": aDataTable().
					SetInfoList(infoListArr[key]).
					SetFilters(panel.GetFiltersMap()).
					SetInfoUrl(infoUrl).
					SetAction(panel.GetInfo().Action).
					SetIsTab(key != 0).
					SetPrimaryKey(panel.GetPrimaryKey().Name).
					SetThead(theadArr[key]).
					SetExportUrl(exportUrl).
					SetNewUrl(newUrl).
					SetEditUrl(editUrl).
					SetDeleteUrl(deleteUrl).GetContent(),
			}
		}
		body = aTab().SetData(tabsHtml).GetContent()
	} else {
		dataTable = aDataTable().
			SetInfoList(panelInfo.InfoList).
			SetFilters(panel.GetFiltersMap()).
			SetInfoUrl(infoUrl).
			SetAction(panel.GetInfo().Action).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetThead(panelInfo.Thead).
			SetExportUrl(exportUrl).
			SetNewUrl(newUrl).
			SetEditUrl(editUrl).
			SetDeleteUrl(deleteUrl)
		body = dataTable.GetContent()
	}

	box := aBox().
		SetBody(body).
		SetHeader(dataTable.GetDataTableHeader() + panel.GetInfo().HeaderHtml).
		WithHeadBorder(false).
		SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))

	return template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))
}

func Assets(ctx *context.Context) {
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

	var contentType string
	if fileSuffix == "css" || fileSuffix == "js" {
		contentType = "text/" + fileSuffix + "; charset=utf-8"
	} else {
		contentType = "image/" + fileSuffix
	}

	ctx.Write(http.StatusOK, map[string]string{
		"content-type": contentType,
	}, string(data))
}

func Export(ctx *context.Context) {
	param := guard.GetExportParam(ctx)

	tableName := "Sheet1"
	prefix := ctx.Query("__prefix")
	panel := table.List[prefix]

	f := excelize.NewFile()
	index := f.NewSheet(tableName)
	f.SetActiveSheet(index)

	// TODO: support any numbers of fields.
	orders := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var (
		panelInfo table.PanelInfo
		fileName  string
		err       error
	)

	if len(param.Id) == 1 {
		params := parameter.GetParam(ctx.Request.URL.Query())
		panelInfo, err = panel.GetDataFromDatabase(ctx.Path(), params)
		fileName = fmt.Sprintf("%s-%d-page-%s-pageSize-%s.xlsx", panel.GetInfo().Title, time.Now().Unix(), params.Page, params.PageSize)
	} else {
		panelInfo, err = panel.GetDataFromDatabaseWithIds(ctx.Path(), parameter.GetParam(ctx.Request.URL.Query()), param.Id)
		fileName = fmt.Sprintf("%s-%d-id-%s.xlsx", panel.GetInfo().Title, time.Now().Unix(), strings.Join(param.Id, "_"))
	}

	if err != nil {
		response.Error(ctx, "export error")
		return
	}

	for key, head := range panelInfo.Thead {
		f.SetCellValue(tableName, orders[key]+"1", head["head"])
	}

	count := 2
	for _, info := range panelInfo.InfoList {
		for key, head := range panelInfo.Thead {
			f.SetCellValue(tableName, orders[key]+strconv.Itoa(count), info[head["field"]])
		}
		count++
	}

	buf, err := f.WriteToBuffer()

	if err != nil || buf == nil {
		response.Error(ctx, "export error")
		return
	}

	ctx.AddHeader("content-disposition", `attachment; filename=`+fileName)
	ctx.Data(200, "application/vnd.ms-excel", buf.Bytes())
}
