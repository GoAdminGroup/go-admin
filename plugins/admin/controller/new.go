package controller

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/menu"
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
	showNewForm(ctx, "", param.Prefix, param.Param.GetRouteParamStr(), false)
}

func showNewForm(ctx *context.Context, alert template2.HTML, prefix string, paramStr string, isNew bool) {

	user := auth.Auth(ctx)

	panel := table.Get(prefix, ctx)

	formList, groupFormList, groupHeaders := table.GetNewFormList(panel.GetForm().TabHeaders, panel.GetForm().TabGroups,
		panel.GetForm().FieldList)

	infoUrl := routePathWithPrefix("info", prefix) + paramStr
	newUrl := routePathWithPrefix("new", prefix)
	showNewUrl := routePathWithPrefix("show_new", prefix) + paramStr

	referer := ctx.Headers("Referer")

	if referer != "" && !isInfoUrl(referer) && !isNewUrl(referer, ctx.Query(constant.PrefixKey)) {
		infoUrl = referer
	}

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetPrefix(config.PrefixFixSlash()).
			SetContent(formList).
			SetTabContents(groupFormList).
			SetTabHeaders(groupHeaders).
			SetUrl(newUrl).
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

	if isNew {
		ctx.AddHeader(constant.PjaxUrlHeader, showNewUrl)
	}
}

// NewForm insert a table row into database.
func NewForm(ctx *context.Context) {

	param := guard.GetNewFormParam(ctx)

	paramStr := param.Param.GetRouteParamStr()

	if param.HasAlert() {
		showNewForm(ctx, param.Alert, param.Prefix, paramStr, true)
		return
	}

	// process uploading files, only support local storage
	if len(param.MultiForm.File) > 0 {
		err := file.GetFileEngine(config.FileUploadEngine.Name).Upload(param.MultiForm)
		if err != nil {
			alert := aAlert().SetTitle(constant.DefaultErrorMsg).
				SetTheme("warning").
				SetContent(template2.HTML(err.Error())).
				GetContent()
			showNewForm(ctx, alert, param.Prefix, paramStr, true)
			return
		}
	}

	err := param.Panel.InsertDataFromDatabase(param.Value())
	if err != nil {
		alert := aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		showNewForm(ctx, alert, param.Prefix, paramStr, true)
		return
	}

	if !param.FromList {
		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>location.href="%s"</script>`, param.PreviousPath))
		ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
		return
	}

	buf := showTable(ctx, param.Prefix, param.Path, param.Param)

	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, routePathWithPrefix("info", param.Prefix)+paramStr)
}
