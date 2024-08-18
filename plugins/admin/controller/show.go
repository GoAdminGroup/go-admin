package controller

import (
	"bytes"
	"crypto/md5"
	"fmt"
	template2 "html/template"
	"mime"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
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
	form2 "github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/GoAdminGroup/html"
)

// ShowInfo show info page.
func (h *Handler) ShowInfo(ctx *context.Context) {

	prefix := ctx.Query(constant.PrefixKey)

	panel := h.table(prefix, ctx)

	if panel.GetOnlyUpdateForm() {
		ctx.Redirect(h.routePathWithPrefix("show_edit", prefix))
		return
	}

	if panel.GetOnlyNewForm() {
		ctx.Redirect(h.routePathWithPrefix("show_new", prefix))
		return
	}

	if panel.GetOnlyDetail() {
		ctx.Redirect(h.routePathWithPrefix("detail", prefix))
		return
	}

	params := parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
		panel.GetInfo().GetSort())

	buf := h.showTable(ctx, prefix, params, panel)
	ctx.HTML(http.StatusOK, buf.String())
}

func (h *Handler) showTableData(ctx *context.Context, prefix string, params parameter.Parameters,
	panel table.Table, urlNamePrefix string) (table.Table, table.PanelInfo, []string, error) {
	if panel == nil {
		panel = h.table(prefix, ctx)
	}

	panelInfo, err := panel.GetData(ctx, params.WithIsAll(false))

	if err != nil {
		return panel, panelInfo, nil, err
	}

	var (
		paramStr = params.DeleteIsAll().GetRouteParamStr()

		editUrl   = modules.AorEmpty(!panel.GetInfo().IsHideEditButton, h.routePathWithPrefix(urlNamePrefix+"show_edit", prefix)+paramStr)
		newUrl    = modules.AorEmpty(!panel.GetInfo().IsHideNewButton, h.routePathWithPrefix(urlNamePrefix+"show_new", prefix)+paramStr)
		deleteUrl = modules.AorEmpty(!panel.GetInfo().IsHideDeleteButton, h.routePathWithPrefix(urlNamePrefix+"delete", prefix)+paramStr)
		exportUrl = modules.AorEmpty(!panel.GetInfo().IsHideExportButton, h.routePathWithPrefix(urlNamePrefix+"export", prefix)+paramStr)
		detailUrl = modules.AorEmpty(!panel.GetInfo().IsHideDetailButton, h.routePathWithPrefix(urlNamePrefix+"detail", prefix)+paramStr)

		infoUrl   = h.routePathWithPrefix(urlNamePrefix+"info", prefix)
		updateUrl = h.routePathWithPrefix(urlNamePrefix+"update", prefix) + paramStr

		user = auth.Auth(ctx)
	)

	editUrl = user.GetCheckPermissionByUrlMethod(editUrl, h.route(urlNamePrefix+"show_edit").Method())
	newUrl = user.GetCheckPermissionByUrlMethod(newUrl, h.route(urlNamePrefix+"show_new").Method())
	deleteUrl = user.GetCheckPermissionByUrlMethod(deleteUrl, h.route(urlNamePrefix+"delete").Method())
	exportUrl = user.GetCheckPermissionByUrlMethod(exportUrl, h.route(urlNamePrefix+"export").Method())
	detailUrl = user.GetCheckPermissionByUrlMethod(detailUrl, h.route(urlNamePrefix+"detail").Method())

	return panel, panelInfo, []string{editUrl, newUrl, deleteUrl, exportUrl, detailUrl, infoUrl, updateUrl}, nil
}

func (h *Handler) showTable(ctx *context.Context, prefix string, params parameter.Parameters, panel table.Table) *bytes.Buffer {

	panel, panelInfo, urls, err := h.showTableData(ctx, prefix, params, panel, "")
	if err != nil {
		return h.Execute(ctx, auth.Auth(ctx),
			template.WarningPanelWithDescAndTitle(ctx, err.Error(), errors.Msg, errors.Msg), "",
			template.ExecuteOptions{Animation: params.Animation})
	}

	if panel.GetInfo().HasError() {
		if panel.GetInfo().PageErrorHTML != template2.HTML("") {
			return h.Execute(ctx, auth.Auth(ctx),
				types.Panel{Content: panel.GetInfo().PageErrorHTML}, "",
				template.ExecuteOptions{Animation: params.Animation})
		}
		return h.Execute(ctx, auth.Auth(ctx),
			template.WarningPanel(ctx, panel.GetInfo().PageError.Error(),
				template.GetPageTypeFromPageError(panel.GetInfo().PageError)), "",
			template.ExecuteOptions{Animation: params.Animation})
	}

	editUrl, newUrl, deleteUrl, exportUrl, detailUrl, infoUrl,
		updateUrl := urls[0], urls[1], urls[2], urls[3], urls[4], urls[5], urls[6]

	var (
		actionJs  template2.JS
		body      template2.HTML
		dataTable types.DataTableAttribute

		user          = auth.Auth(ctx)
		info          = panel.GetInfo()
		actionBtns    = info.Action
		allActionBtns = info.ActionButtons.CheckPermissionWhenURLAndMethodNotEmpty(user)
	)

	hasExtAction := actionBtns == template.HTML("") && len(allActionBtns) > 0
	hasAction := hasExtAction || (editUrl != "" || newUrl != "" || deleteUrl != "")

	if hasExtAction {
		if info.ActionButtonFold {
			ext := template2.HTML("")
			if deleteUrl != "" {
				ext = html.LiEl().SetClass("divider").Get()
				allActionBtns = append([]types.Button{types.GetActionButton(language.GetFromHtml("delete"),
					types.NewDefaultAction(`data-id='{{.Id}}' data-param='{{(index .Value "__goadmin_delete_params").Content}}' style="cursor: pointer;"`,
						ext, "", ""), "grid-row-delete")}, allActionBtns...)
			}
			ext = template2.HTML("")
			if detailUrl != "" {
				if editUrl == "" && deleteUrl == "" {
					ext = html.LiEl().SetClass("divider").Get()
				}
				allActionBtns = append([]types.Button{types.GetActionButton(language.GetFromHtml("detail"),
					action.Jump(detailUrl+"&"+constant.DetailPKKey+`={{.Id}}{{(index .Value "__goadmin_detail_params").Content}}`, ext))}, allActionBtns...)
			}
			if editUrl != "" {
				if detailUrl == "" && deleteUrl == "" {
					ext = html.LiEl().SetClass("divider").Get()
				}
				allActionBtns = append([]types.Button{types.GetActionButton(language.GetFromHtml("edit"),
					action.Jump(editUrl+"&"+constant.EditPKKey+`={{.Id}}{{(index .Value "__goadmin_edit_params").Content}}`, ext))}, allActionBtns...)
			}

			var content template2.HTML
			content, actionJs = allActionBtns.Content(ctx)

			actionBtns = html.Div(html.Div(
				html.A(icon.Icon(icon.EllipsisV),
					html.M{"color": "#676565"},
					html.M{"href": "#"},
				), html.M{"cursor": "pointer", "width": "100%"}, html.M{"class": "dropdown-toggle", "data-toggle": "dropdown"})+
				html.Ul(content,
					html.M{"min-width": "20px !important", "left": "-112px", "overflow": "hidden"},
					html.M{"class": "dropdown-menu", "role": "menu", "aria-labelledby": "dLabel"}),
				html.M{"text-align": "center"}, html.M{"class": "dropdown"})
		} else {
			actionBtns, actionJs = allActionBtns.Content(ctx)
		}
	} else {
		info.ActionButtonFold = false
	}

	btns, btnsJs := info.Buttons.CheckPermissionWhenURLAndMethodNotEmpty(user).Content(ctx)

	if info.TabGroups.Valid() {

		dataTable = aDataTable(ctx).
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
				"content": aDataTable(ctx).
					SetInfoList(infoListArr[key]).
					SetInfoUrl(infoUrl).
					SetButtons(btns).
					SetSticky(hasAction).
					SetActionJs(btnsJs + actionJs).
					SetHasFilter(len(panelInfo.FilterFormData) > 0).
					SetAction(actionBtns).
					SetActionFold(info.ActionButtonFold).
					SetIsTab(key != 0).
					SetPrimaryKey(panel.GetPrimaryKey().Name).
					SetThead(theadArr[key]).
					SetHideRowSelector(info.IsHideRowSelector).
					SetLayout(info.TableLayout).
					SetExportUrl(exportUrl).
					SetNewUrl(newUrl).
					SetSortUrl(params.GetFixedParamStrWithoutSort()).
					SetEditUrl(editUrl).
					SetUpdateUrl(updateUrl).
					SetDetailUrl(detailUrl).
					SetDeleteUrl(deleteUrl).
					GetContent(),
			}
		}
		body = aTab(ctx).SetData(tabsHtml).GetContent()
	} else {
		dataTable = aDataTable(ctx).
			SetInfoList(panelInfo.InfoList).
			SetInfoUrl(infoUrl).
			SetButtons(btns).
			SetSticky(hasAction).
			SetLayout(info.TableLayout).
			SetActionJs(btnsJs + actionJs).
			SetAction(actionBtns).
			SetHasFilter(len(panelInfo.FilterFormData) > 0).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetThead(panelInfo.Thead).
			SetExportUrl(exportUrl).
			SetActionFold(info.ActionButtonFold).
			SetHideRowSelector(info.IsHideRowSelector).
			SetHideFilterArea(info.IsHideFilterArea).
			SetNewUrl(newUrl).
			SetEditUrl(editUrl).
			SetSortUrl(params.GetFixedParamStrWithoutSort()).
			SetUpdateUrl(updateUrl).
			SetDetailUrl(detailUrl).
			SetDeleteUrl(deleteUrl)
		body = dataTable.GetContent()
	}

	isNotIframe := ctx.Query(constant.IframeKey) != "true"
	paginator := panelInfo.Paginator

	if !isNotIframe {
		paginator = paginator.SetEntriesInfo("")
	}

	boxModel := aBox(ctx).
		SetBody(body).
		SetStyle(template2.HTMLAttr(`overflow-x: auto;overflow-y: hidden;`)).
		SetNoPadding().
		SetHeader(dataTable.GetDataTableHeader() + info.HeaderHtml).
		WithHeadBorder().
		SetIframeStyle(!isNotIframe).
		SetFooter(paginator.GetContent() + info.FooterHtml + `
		<script>
		$(document).ready(function() {
			var tableWrapper = $(".table");
			var lastTh = tableWrapper.find('tbody th:last-child');
			var lastTd = tableWrapper.find('tbody td:last-child');

			var minWidth = parseInt(tableWrapper.css('min-width'));

			var resize = function() {
				if (tableWrapper.width() <= minWidth) {
					lastTh.addClass('last_th_td_ele');
					lastTd.addClass('last_th_td_ele');
				} else {
					lastTh.removeClass('last_th_td_ele');
					lastTd.removeClass('last_th_td_ele');
				}
			}

			resize();

			$(window).resize(function() {
				resize();
			});
		});
</script>
		`)

	content := template2.HTML("")

	if len(panelInfo.FilterFormData) > 0 && info.FilterFormLayout == form2.LayoutFilter {
		filterBoxModel := aBox(ctx).SetClass("filter-area").
			SetAttr(`style="padding-top: 20px;margin-top: -10px;margin-bottom: 12px;padding-left: 20px;display: block;padding-bottom: 5px;"`).
			SetStyle(`padding: 0px;`).
			SetBody(aForm(ctx).
				SetContent(panelInfo.FilterFormData).
				SetHorizontal(true).
				SetPrefix(h.config.PrefixFixSlash()).
				SetInputWidth(info.FilterFormInputWidth).
				SetHeadWidth(info.FilterFormHeadWidth).
				SetMethod("get").
				SetLayout(info.FilterFormLayout).
				SetUrl(infoUrl). //  + params.GetFixedParamStrWithoutColumnsAndPage()
				SetHiddenFields(map[string]string{
					form.NoAnimationKey: "true",
				}).
				GetContent())
		content = filterBoxModel.GetContent() + boxModel.GetContent()
	} else {
		if len(panelInfo.FilterFormData) > 0 {
			boxModel = boxModel.SetSecondHeaderClass("filter-area").
				SetSecondHeader(aForm(ctx).
					SetContent(panelInfo.FilterFormData).
					SetPrefix(h.config.PrefixFixSlash()).
					SetInputWidth(info.FilterFormInputWidth).
					SetHeadWidth(info.FilterFormHeadWidth).
					SetMethod("get").
					SetLayout(info.FilterFormLayout).
					SetUrl(infoUrl). //  + params.GetFixedParamStrWithoutColumnsAndPage()
					SetHiddenFields(map[string]string{
						form.NoAnimationKey: "true",
					}).
					SetOperationFooter(filterFormFooter(ctx, infoUrl)).
					GetContent())
		}
		content = boxModel.GetContent()
	}

	if info.Wrapper != nil {
		content = info.Wrapper(content)
	}

	interval := make([]int, 0)
	autoRefresh := info.AutoRefresh != uint(0)
	if autoRefresh {
		interval = append(interval, int(info.AutoRefresh))
	}

	return h.Execute(ctx, user, types.Panel{
		Content:         content,
		Description:     template2.HTML(panelInfo.Description),
		Title:           modules.AorBHTML(isNotIframe, template2.HTML(panelInfo.Title), ""),
		MiniSidebar:     info.HideSideBar,
		AutoRefresh:     autoRefresh,
		RefreshInterval: interval,
	}, "", template.ExecuteOptions{Animation: params.Animation, NoCompress: info.NoCompress})
}

// Assets return front-end assets according the request path.
func (h *Handler) Assets(ctx *context.Context) {
	filepath := h.config.URLRemovePrefix(ctx.Path())

	theme := h.assetsTheme[filepath]
	if theme == "" {
		theme = h.config.Theme
	}

	data, err := aTemplateByTheme(ctx, theme).GetAsset(filepath)

	if err != nil {
		data, err = template.GetAsset(filepath)
		if err != nil {
			logger.ErrorCtx(ctx, "asset err %+v", err)
			ctx.Write(http.StatusNotFound, map[string]string{}, "")
			return
		}
	}

	var contentType = mime.TypeByExtension(path.Ext(filepath))

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	etag := fmt.Sprintf("%x", md5.Sum(data))

	if match := ctx.Headers("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			ctx.SetStatusCode(http.StatusNotModified)
			return
		}
	}

	ctx.DataWithHeaders(http.StatusOK, map[string]string{
		"Content-Type":   contentType,
		"Cache-Control":  "max-age=2592000",
		"Content-Length": strconv.Itoa(len(data)),
		"ETag":           etag,
	}, data)
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

	var (
		infoData  table.PanelInfo
		fileName  string
		err       error
		tableInfo = panel.GetInfo()
		params    parameter.Parameters
	)

	if fn := panel.GetInfo().ExportProcessFn; fn != nil {
		params = parameter.GetParam(ctx.Request.URL, tableInfo.DefaultPageSize, tableInfo.SortField,
			tableInfo.GetSort())
		p, err := fn(params.WithIsAll(param.IsAll))
		if err != nil {
			response.Error(ctx, "export error")
			return
		}
		infoData.Thead = p.Thead
		infoData.InfoList = p.InfoList
	} else {
		if len(param.Id) == 0 {
			params = parameter.GetParam(ctx.Request.URL, tableInfo.DefaultPageSize, tableInfo.SortField,
				tableInfo.GetSort())
			infoData, err = panel.GetData(ctx, params.WithIsAll(param.IsAll))
			fileName = fmt.Sprintf("%s-%d-page-%s-pageSize-%s.xlsx", tableInfo.Title, time.Now().Unix(),
				params.Page, params.PageSize)
		} else {
			infoData, err = panel.GetDataWithIds(ctx, parameter.GetParam(ctx.Request.URL,
				tableInfo.DefaultPageSize, tableInfo.SortField, tableInfo.GetSort()).WithPKs(param.Id...))
			fileName = fmt.Sprintf("%s-%d-id-%s.xlsx", tableInfo.Title, time.Now().Unix(), strings.Join(param.Id, "_"))
		}
		if err != nil {
			response.Error(ctx, "export error")
			return
		}
	}

	// TODO: support any numbers of fields.
	orders := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if len(infoData.Thead) > 26 {
		j := -1
		for i := 0; i < len(infoData.Thead)-26; i++ {
			if i%26 == 0 {
				j++
			}
			letter := orders[j] + orders[i%26]
			orders = append(orders, letter)
		}
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
