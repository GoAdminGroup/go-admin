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
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/html"
	template2 "html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

// ShowInfo show info page.
func ShowInfo(ctx *context.Context) {

	prefix := ctx.Query(constant.PrefixKey)
	panel := table.Get(prefix, ctx)

	params := parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
		panel.GetInfo().GetSort())

	buf := showTable(ctx, prefix, ctx.Path(), params)
	ctx.HTML(http.StatusOK, buf.String())
}

func showTable(ctx *context.Context, prefix, path string, params parameter.Parameters) *bytes.Buffer {

	panel := table.Get(prefix, ctx)

	panelInfo, err := panel.GetData(path, params, false)

	if err != nil {
		tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
		user := auth.Auth(ctx)
		alert := aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		errMsg := language.Get("error")
		return template.Execute(tmpl, tmplName, user, types.Panel{
			Content:     alert,
			Description: errMsg,
			Title:       errMsg,
		}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))
	}

	paramStr := params.GetRouteParamStr()

	editUrl := modules.AorEmpty(panel.GetEditable(), routePathWithPrefix("show_edit", prefix)+paramStr)
	newUrl := modules.AorEmpty(panel.GetCanAdd(), routePathWithPrefix("show_new", prefix)+paramStr)
	deleteUrl := modules.AorEmpty(panel.GetDeletable(), routePathWithPrefix("delete", prefix))
	exportUrl := modules.AorEmpty(panel.GetExportable(), routePathWithPrefix("export", prefix)+paramStr)
	detailUrl := modules.AorEmpty(panel.IsShowDetail(), routePathWithPrefix("detail", prefix)+paramStr)

	infoUrl := routePathWithPrefix("info", prefix)
	updateUrl := routePathWithPrefix("update", prefix)

	user := auth.Auth(ctx)

	editUrl = user.GetCheckPermissionByUrlMethod(editUrl, route("show_edit").Method())
	newUrl = user.GetCheckPermissionByUrlMethod(newUrl, route("show_new").Method())
	deleteUrl = user.GetCheckPermissionByUrlMethod(deleteUrl, route("delete").Method())
	exportUrl = user.GetCheckPermissionByUrlMethod(exportUrl, route("export").Method())
	detailUrl = user.GetCheckPermissionByUrlMethod(detailUrl, route("detail").Method())

	var (
		body       template2.HTML
		dataTable  types.DataTableAttribute
		info       = panel.GetInfo()
		actionBtns = info.Action
		actionJs   template2.JS
	)

	btns, btnsJs := info.Buttons.Content()

	if actionBtns == template.HTML("") && len(info.ActionButtons) > 0 {
		ext := template.HTML("")
		if deleteUrl != "" {
			ext = html.LiEl().SetClass("divider").Get()
			info.AddActionButtonFront(language.GetFromHtml("delete"), types.NewDefaultAction(`data-id='{%id}' style="cursor: pointer;"`,
				ext, ""), "grid-row-delete")
		}
		ext = template.HTML("")
		if detailUrl != "" {
			if editUrl == "" && deleteUrl == "" {
				ext = html.LiEl().SetClass("divider").Get()
			}
			info.AddActionButtonFront(language.GetFromHtml("detail"), action.Jump(detailUrl+"&"+constant.DetailPKKey+"={%id}", ext))
		}
		if editUrl != "" {
			if detailUrl == "" && deleteUrl == "" {
				ext = html.LiEl().SetClass("divider").Get()
			}
			info.AddActionButtonFront(language.GetFromHtml("edit"), action.Jump(editUrl+"&"+constant.EditPKKey+"={%id}", ext))
		}

		var content template2.HTML
		content, actionJs = info.ActionButtons.Content()

		actionBtns = html.Div(
			html.A(icon.Icon(icon.EllipsisV),
				html.M{"color": "#676565"},
				html.M{"class": "dropdown-toggle", "href": "#", "data-toggle": "dropdown"},
			)+html.Ul(content,
				html.M{"min-width": "20px !important", "left": "-32px", "overflow": "hidden"},
				html.M{"class": "dropdown-menu", "role": "menu", "aria-labelledby": "dLabel"}),
			html.M{"text-align": "center"}, html.M{"class": "dropdown"})
	}

	if info.TabGroups.Valid() {

		dataTable = aDataTable().
			SetThead(panelInfo.Thead).
			SetDeleteUrl(deleteUrl).
			SetNewUrl(newUrl).
			SetExportUrl(exportUrl)

		var (
			tabsHtml    = make([]map[string]template2.HTML, len(info.TabHeaders))
			infoListArr = panelInfo.InfoList.GroupBy(info.TabGroups)
			theadArr    = panelInfo.Thead.GroupBy(info.TabGroups)
		)
		for key, header := range info.TabHeaders {
			tabsHtml[key] = map[string]template2.HTML{
				"title": template2.HTML(header),
				"content": aDataTable().
					SetInfoList(infoListArr[key]).
					SetInfoUrl(infoUrl).
					SetButtons(btns).
					SetActionJs(btnsJs + actionJs).
					SetHasFilter(len(panelInfo.FormData) > 0).
					SetAction(actionBtns).
					SetIsTab(key != 0).
					SetPrimaryKey(panel.GetPrimaryKey().Name).
					SetThead(theadArr[key]).
					SetHideRowSelector(info.IsHideRowSelector).
					SetLayout(info.TableLayout).
					SetExportUrl(exportUrl).
					SetNewUrl(newUrl).
					SetEditUrl(editUrl).
					SetUpdateUrl(updateUrl).
					SetDetailUrl(detailUrl).
					SetDeleteUrl(deleteUrl).
					GetContent(),
			}
		}
		body = aTab().SetData(tabsHtml).GetContent()
	} else {
		dataTable = aDataTable().
			SetInfoList(panelInfo.InfoList).
			SetInfoUrl(infoUrl).
			SetButtons(btns).
			SetLayout(info.TableLayout).
			SetActionJs(btnsJs + actionJs).
			SetAction(actionBtns).
			SetHasFilter(len(panelInfo.FormData) > 0).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetThead(panelInfo.Thead).
			SetExportUrl(exportUrl).
			SetHideRowSelector(info.IsHideRowSelector).
			SetHideFilterArea(info.IsHideFilterArea).
			SetNewUrl(newUrl).
			SetEditUrl(editUrl).
			SetUpdateUrl(updateUrl).
			SetDetailUrl(detailUrl).
			SetDeleteUrl(deleteUrl)
		body = dataTable.GetContent()
	}

	boxModel := aBox().
		SetBody(body).
		SetNoPadding().
		SetHeader(dataTable.GetDataTableHeader() + info.HeaderHtml).
		WithHeadBorder().
		SetFooter(panelInfo.Paginator.GetContent())

	if len(panelInfo.FormData) > 0 {
		boxModel = boxModel.SetSecondHeaderClass("filter-area").
			SetSecondHeader(aForm().
				SetContent(panelInfo.FormData).
				SetPrefix(config.PrefixFixSlash()).
				SetMethod("get").
				SetLayout(info.FilterFormLayout).
				SetUrl(infoUrl).
				SetOperationFooter(filterFormFooter(infoUrl)).
				GetContent())
	}

	box := boxModel.GetContent()

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))

	return template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))
}

// Assets return front-end assets according the request path.
func Assets(ctx *context.Context) {
	filepath := config.URLRemovePrefix(ctx.Path())
	data, err := aTemplate().GetAsset(filepath)

	if err != nil {
		data, err = template.GetAsset(filepath)
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

// Export export table rows as excel object.
func Export(ctx *context.Context) {
	param := guard.GetExportParam(ctx)

	tableName := "Sheet1"
	prefix := ctx.Query(constant.PrefixKey)
	panel := table.Get(prefix, ctx)

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
		params := parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
			panel.GetInfo().GetSort())
		panelInfo, err = panel.GetData(ctx.Path(), params, param.IsAll)
		fileName = fmt.Sprintf("%s-%d-page-%s-pageSize-%s.xlsx", panel.GetInfo().Title, time.Now().Unix(),
			params.Page, params.PageSize)
	} else {
		panelInfo, err = panel.GetDataWithIds(ctx.Path(), parameter.GetParam(ctx.Request.URL.Query(),
			panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField, panel.GetInfo().GetSort()), param.Id)
		fileName = fmt.Sprintf("%s-%d-id-%s.xlsx", panel.GetInfo().Title, time.Now().Unix(), strings.Join(param.Id, "_"))
	}

	if err != nil {
		response.Error(ctx, "export error")
		return
	}

	columnIndex := 0
	for _, head := range panelInfo.Thead {
		if head["hide"] != "1" {
			f.SetCellValue(tableName, orders[columnIndex]+"1", head["head"])
			columnIndex++
		}
	}

	count := 2
	for _, info := range panelInfo.InfoList {
		columnIndex = 0
		for _, head := range panelInfo.Thead {
			if head["hide"] != "1" {
				f.SetCellValue(tableName, orders[columnIndex]+strconv.Itoa(count), info[head["field"]])
				columnIndex++
			}
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
