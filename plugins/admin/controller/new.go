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
	panel := table.List[prefix]
	if !panel.GetCanAdd() {
		response.PageNotFound(ctx)
		return
	}
	params := parameter.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)

	formList := table.GetNewFormList(panel.GetForm().FormList)
	for i := 0; i < len(formList); i++ {
		formList[i].Editable = true
	}
	tmpl, tmplName := template.Get(config.THEME).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: template.Get(config.THEME).Form().
			SetPrefix(config.PREFIX).
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
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())
}

func NewForm(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	panel := table.List[prefix]
	if !panel.GetCanAdd() {
		response.PageNotFound(ctx)
		return
	}

	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		response.BadRequest(ctx, "create fail")
		return
	}

	form := ctx.Request.MultipartForm

	// process uploading files, only support local storage
	if len(form.File) > 0 {
		_, _ = file.GetFileEngine("local").Upload(form)
	}

	var err error
	if prefix == "manager" { // manager edit
		if err = newManager(form.Value); err != nil {
			response.Error(ctx, err.Error())
			return
		}
	} else if prefix == "roles" { // role edit
		if err = newRole(form.Value); err != nil {
			response.Error(ctx, err.Error())
			return
		}
	} else {
		panel.InsertDataFromDatabase(form.Value)
	}

	table.RefreshTableList()

	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := parameter.GetParamFromUrl(previous)

	panelInfo := panel.GetDataFromDatabase(prevUrlArr[0], params)

	editUrl := config.Url("/info/" + prefix + "/edit" + params.GetRouteParamStr())
	newUrl := config.Url("/info/" + prefix + "/new" + params.GetRouteParamStr())
	deleteUrl := config.Url("/delete/" + prefix)

	dataTable := template.Get(config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
		SetThead(panelInfo.Thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	box := template.Get(config.THEME).Box().
		SetBody(dataTable.GetContent()).
		SetHeader(dataTable.GetDataTableHeader() + panel.GetInfo().HeaderHtml).
		WithHeadBorder(false).
		SetFooter(panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(config.THEME).GetTemplate(true)
	buffer := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(previous, config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buffer.String())
	ctx.AddHeader(constant.PjaxUrlHeader, previous)
}
