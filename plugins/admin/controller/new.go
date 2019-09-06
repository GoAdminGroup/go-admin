package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/plugins/admin/modules/file"
	"github.com/chenhg5/go-admin/plugins/admin/modules/parameter"
	"github.com/chenhg5/go-admin/plugins/admin/modules/response"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"net/http"
	"strings"
)

func ShowNewForm(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	table.RefreshTableList()
	panel := table.List[prefix]
	if !panel.GetCanAdd() {
		response.Alert(ctx, config, panel.GetForm().Description, panel.GetForm().Title, "operation not allow")
		return
	}
	params := parameter.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)

	formList := table.GetNewFormList(panel.GetForm().FormList)
	for i := 0; i < len(formList); i++ {
		formList[i].Editable = true
	}
	tmpl, tmplName := aTemplate().GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: aForm().
			SetPrefix(config.PrefixFixSlash()).
			SetContent(formList).
			SetUrl(config.Url("/new/" + prefix)).
			SetToken(auth.TokenHelper.AddToken()).
			SetTitle("New").
			SetInfoUrl(config.Url("/info/" + prefix + params.GetRouteParamStr())).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml).
			GetContent(),
		Description: panel.GetForm().Description,
		Title:       panel.GetForm().Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.Prefix(), "", 1)))
	ctx.Html(http.StatusOK, buf.String())
}

func NewForm(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := parameter.GetParamFromUrl(previous)
	panel := table.List[prefix]

	if !panel.GetCanAdd() {
		response.Alert(ctx, config, panel.GetInfo().Description, panel.GetInfo().Title, "operation not allow")
		return
	}

	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		response.Alert(ctx, config, panel.GetInfo().Description, panel.GetInfo().Title, "create fail, wrong token")
		return
	}

	form := ctx.Request.MultipartForm

	// process uploading files, only support local storage
	if len(form.File) > 0 {
		_, _ = file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // manager edit
		newManager(form.Value)
	} else if prefix == "roles" { // role edit
		newRole(form.Value)
	} else {
		panel.InsertDataFromDatabase(form.Value)
	}

	table.RefreshTableList()

	editUrl := config.Url("/info/" + prefix + "/edit" + params.GetRouteParamStr())
	newUrl := config.Url("/info/" + prefix + "/new" + params.GetRouteParamStr())
	deleteUrl := config.Url("/delete/" + prefix)

	panelInfo := panel.GetDataFromDatabase(prevUrlArr[0], params)

	dataTable := aDataTable().
		SetInfoList(panelInfo.InfoList).
		SetThead(panelInfo.Thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	box := aBox().
		SetBody(dataTable.GetContent()).
		SetHeader(dataTable.GetDataTableHeader() + panel.GetInfo().HeaderHtml).
		WithHeadBorder(false).
		SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(true)
	buffer := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(previous, config.Prefix(), "", 1)))

	ctx.Html(http.StatusOK, buffer.String())
	ctx.AddHeader(constant.PjaxUrlHeader, previous)
}
