package controller

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

// ShowNewForm show a new form page.
func ShowNewForm(ctx *context.Context) {
	param := guard.GetShowNewFormParam(ctx)
	showNewForm(ctx, "", param.Prefix, param.GetUrl(), param.GetInfoUrl(), "")
}

func showNewForm(ctx *context.Context, alert template2.HTML, prefix string, url, infoUrl, newUrl string) {

	user := auth.Auth(ctx)

	table.RefreshTableList()
	panel := table.Get(prefix)

	formList, groupFormList, groupHeaders := table.GetNewFormList(panel.GetForm().TabHeaders, panel.GetForm().TabGroups,
		panel.GetForm().FieldList)

	referer := ctx.Headers("Referer")

	if referer != "" && !modules.IsInfoUrl(referer) && !modules.IsNewUrl(referer, ctx.Query("__prefix")) {
		infoUrl = referer
	}

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetPrefix(config.PrefixFixSlash()).
			SetContent(formList).
			SetTabContents(groupFormList).
			SetTabHeaders(groupHeaders).
			SetUrl(url).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetToken(authSrv().AddToken()).
			SetOperationFooter(formFooter()).
			SetTitle("New").
			SetInfoUrl(infoUrl).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml)),
		Description: panel.GetForm().Description,
		Title:       panel.GetForm().Title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))
	ctx.HTML(http.StatusOK, buf.String())

	if newUrl != "" {
		ctx.AddHeader(constant.PjaxUrlHeader, newUrl)
	}
}

// NewForm insert a table row into database.
func NewForm(ctx *context.Context) {

	param := guard.GetNewFormParam(ctx)

	if param.HasAlert() {
		showNewForm(ctx, param.Alert, param.Prefix, param.GetUrl(), param.GetInfoUrl(), param.GetNewUrl())
		return
	}

	// process uploading files, only support local storage
	if len(param.MultiForm.File) > 0 {
		err := file.GetFileEngine(config.FileUploadEngine.Name).Upload(param.MultiForm)
		if err != nil {
			alert := aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
				SetTheme("warning").
				SetContent(template2.HTML(err.Error())).
				GetContent()
			showNewForm(ctx, alert, param.Prefix, param.GetUrl(), param.GetInfoUrl(), param.GetNewUrl())
			return
		}
	}

	err := param.Panel.InsertDataFromDatabase(param.Value())
	if err != nil {
		alert := aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		showNewForm(ctx, alert, param.Prefix, param.GetUrl(), param.GetInfoUrl(), param.GetNewUrl())
		return
	}

	if !param.FromList {
		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>location.href="%s"</script>`, param.PreviousPath))
		ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
		return
	}

	editUrl := modules.AorB(param.Panel.GetEditable(), param.GetEditUrl(), "")
	deleteUrl := modules.AorB(param.Panel.GetDeletable(), param.GetDeleteUrl(), "")
	exportUrl := modules.AorB(param.Panel.GetExportable(), param.GetExportUrl(), "")
	newUrl := modules.AorB(param.Panel.GetCanAdd(), param.GetNewUrl(), "")
	infoUrl := param.GetInfoUrl()
	updateUrl := modules.AorB(param.Panel.GetEditable(), param.GetUpdateUrl(), "")
	detailUrl := param.GetDetailUrl()

	buf := showTable(ctx, param.Panel, param.Path, param.Param, exportUrl, newUrl, deleteUrl,
		infoUrl, editUrl, updateUrl, detailUrl)

	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.GetInfoUrl())
}
