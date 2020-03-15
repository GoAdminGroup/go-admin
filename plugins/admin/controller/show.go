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
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
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
func (h *Handler) ShowInfo(ctx *context.Context) {

	prefix := ctx.Query(constant.PrefixKey)
	panel := h.table(prefix, ctx)

	params := parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
		panel.GetInfo().GetSort())

	buf := h.showTable(ctx, prefix, params)
	ctx.HTML(http.StatusOK, buf.String())
}

func (h *Handler) showTable(ctx *context.Context, prefix string, params parameter.Parameters) *bytes.Buffer {

	panel := h.table(prefix, ctx)

	panelInfo, err := panel.GetData(params.WithIsAll(false))

	if err != nil {
		tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
		user := auth.Auth(ctx)
		alert := aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		errMsg := language.GetUser("error", user.Id)
		return template.Execute(tmpl, tmplName, user, types.Panel{
			Content:     alert,
			Description: errMsg,
			Title:       errMsg,
		}, h.config, menu.GetGlobalMenu(user, h.conn).SetActiveClass(h.config.URLRemovePrefix(ctx.Path())))
	}

	paramStr := params.GetRouteParamStr()

	editUrl := modules.AorEmpty(panel.GetEditable(), h.routePathWithPrefix("show_edit", prefix)+paramStr)
	newUrl := modules.AorEmpty(panel.GetCanAdd(), h.routePathWithPrefix("show_new", prefix)+paramStr)
	deleteUrl := modules.AorEmpty(panel.GetDeletable(), h.routePathWithPrefix("delete", prefix))
	exportUrl := modules.AorEmpty(panel.GetExportable(), h.routePathWithPrefix("export", prefix)+paramStr)
	detailUrl := modules.AorEmpty(panel.IsShowDetail(), h.routePathWithPrefix("detail", prefix)+paramStr)

	infoUrl := h.routePathWithPrefix("info", prefix)
	updateUrl := h.routePathWithPrefix("update", prefix)

	user := auth.Auth(ctx)

	editUrl = user.GetCheckPermissionByUrlMethod(editUrl, h.route("show_edit").Method())
	newUrl = user.GetCheckPermissionByUrlMethod(newUrl, h.route("show_new").Method())
	deleteUrl = user.GetCheckPermissionByUrlMethod(deleteUrl, h.route("delete").Method())
	exportUrl = user.GetCheckPermissionByUrlMethod(exportUrl, h.route("export").Method())
	detailUrl = user.GetCheckPermissionByUrlMethod(detailUrl, h.route("detail").Method())

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
			info.AddActionButtonFront(language.GetUserFromHtml("delete", user.Id), types.NewDefaultAction(`data-id='{%id}' style="cursor: pointer;"`,
				ext, "", ""), "grid-row-delete")
		}
		ext = template.HTML("")
		if detailUrl != "" {
			if editUrl == "" && deleteUrl == "" {
				ext = html.LiEl().SetClass("divider").Get()
			}
			info.AddActionButtonFront(language.GetUserFromHtml("detail", user.Id), action.Jump(detailUrl+"&"+constant.DetailPKKey+"={%id}", ext))
		}
		if editUrl != "" {
			if detailUrl == "" && deleteUrl == "" {
				ext = html.LiEl().SetClass("divider").Get()
			}
			info.AddActionButtonFront(language.GetUserFromHtml("edit", user.Id), action.Jump(editUrl+"&"+constant.EditPKKey+"={%id}", ext))
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
					SetHasFilter(len(panelInfo.FilterFormData) > 0).
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
			SetHasFilter(len(panelInfo.FilterFormData) > 0).
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
		SetFooter(panelInfo.Paginator.GetContent() + info.FooterHtml)

	if len(panelInfo.FilterFormData) > 0 {
		boxModel = boxModel.SetSecondHeaderClass("filter-area").
			SetSecondHeader(aForm().
				SetContent(panelInfo.FilterFormData).
				SetPrefix(h.config.PrefixFixSlash()).
				SetInputWidth(10).
				SetMethod("get").
				SetLayout(info.FilterFormLayout).
				SetUrl(infoUrl).
				SetHiddenFields(map[string]string{
					form.NoAnimationKey: "true",
				}).
				SetOperationFooter(filterFormFooter(infoUrl)).
				GetContent())
	}

	box := boxModel.GetContent()

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))

	return template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, h.config, menu.GetGlobalMenu(user, h.conn).SetActiveClass(h.config.URLRemovePrefix(ctx.Path())), params.Animation)
}

// Assets return front-end assets according the request path.
func (h *Handler) Assets(ctx *context.Context) {
	filepath := h.config.URLRemovePrefix(ctx.Path())
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
func (h *Handler) Export(ctx *context.Context) {
	param := guard.GetExportParam(ctx)

	tableName := "Sheet1"
	prefix := ctx.Query(constant.PrefixKey)
	panel := h.table(prefix, ctx)

	f := excelize.NewFile()
	index := f.NewSheet(tableName)
	f.SetActiveSheet(index)

	// TODO: support any numbers of fields.
	orders := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var (
		infoData  table.PanelInfo
		fileName  string
		err       error
		tableInfo = panel.GetInfo()
	)

	if len(param.Id) == 0 {
		params := parameter.GetParam(ctx.Request.URL, tableInfo.DefaultPageSize, tableInfo.SortField,
			tableInfo.GetSort())
		infoData, err = panel.GetData(params.WithIsAll(param.IsAll))
		fileName = fmt.Sprintf("%s-%d-page-%s-pageSize-%s.xlsx", tableInfo.Title, time.Now().Unix(),
			params.Page, params.PageSize)
	} else {
		infoData, err = panel.GetDataWithIds(parameter.GetParam(ctx.Request.URL,
			tableInfo.DefaultPageSize, tableInfo.SortField, tableInfo.GetSort()).WithPKs(param.Id...))
		fileName = fmt.Sprintf("%s-%d-id-%s.xlsx", tableInfo.Title, time.Now().Unix(), strings.Join(param.Id, "_"))
	}

	if err != nil {
		response.Error(ctx, "export error")
		return
	}

	columnIndex := 0
	for _, head := range infoData.Thead {
		if !head.Hide {
			f.SetCellValue(tableName, orders[columnIndex]+"1", head.Head)
			columnIndex++
		}
	}

	count := 2
	for _, info := range infoData.InfoList {
		columnIndex = 0
		for _, head := range infoData.Thead {
			if !head.Hide {
				if tableInfo.IsExportValue() {
					f.SetCellValue(tableName, orders[columnIndex]+strconv.Itoa(count), info[head.Field].Value)
				} else {
					f.SetCellValue(tableName, orders[columnIndex]+strconv.Itoa(count), info[head.Field].Content)
				}
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

