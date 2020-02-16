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
	"github.com/GoAdminGroup/go-admin/template/types/form"
	template2 "html/template"
	"net/http"
)

// ShowForm show form page.
func ShowForm(ctx *context.Context) {
	param := guard.GetShowFormParam(ctx)
	showForm(ctx, "", param.Prefix, param.Id, param.Param.GetRouteParamStr(), false)
}

func showForm(ctx *context.Context, alert template2.HTML, prefix string, id string, paramStr string, isEdit bool) {

	panel := table.Get(prefix, ctx)

	formData, groupFormData, groupHeaders, title, description, err := panel.GetDataWithId(id)

	if err != nil && alert == "" {
		alert = aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
	}

	user := auth.Auth(ctx)

	infoUrl := routePathWithPrefix("info", prefix) + paramStr
	editUrl := routePathWithPrefix("edit", prefix)
	showEditUrl := routePathWithPrefix("show_edit", prefix) + paramStr

	referer := ctx.Headers("Referer")

	if referer != "" && !isInfoUrl(referer) && !isEditUrl(referer, ctx.Query(constant.PrefixKey)) {
		infoUrl = referer
	}

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetContent(formData).
			SetTabContents(groupFormData).
			SetTabHeaders(groupHeaders).
			SetPrefix(config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(editUrl).
			SetToken(authSrv().AddToken()).
			SetInfoUrl(infoUrl).
			SetOperationFooter(formFooter()).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml)),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())

	if isEdit {
		ctx.AddHeader(constant.PjaxUrlHeader, showEditUrl)
	}
}

func EditForm(ctx *context.Context) {

	param := guard.GetEditFormParam(ctx)

	if param.HasAlert() {
		showForm(ctx, param.Alert, param.Prefix, param.Id, param.Param.GetRouteParamStr(), true)
		return
	}

	// process uploading files, only support local storage for now.
	if len(param.MultiForm.File) > 0 {
		err := file.GetFileEngine(config.FileUploadEngine.Name).Upload(param.MultiForm)
		if err != nil {
			alert := aAlert().SetTitle(constant.DefaultErrorMsg).
				SetTheme("warning").
				SetContent(template2.HTML(err.Error())).
				GetContent()
			showForm(ctx, alert, param.Prefix, param.Id, param.Param.GetRouteParamStr(), true)
			return
		}
	}

	for _, field := range param.Panel.GetForm().FieldList {
		if field.FormType == form.File &&
			len(param.MultiForm.File[field.Field]) == 0 &&
			param.MultiForm.Value[field.Field+"__delete_flag"][0] != "1" {
			delete(param.MultiForm.Value, field.Field)
		}
	}

	err := param.Panel.UpdateDataFromDatabase(param.Value())
	if err != nil {
		alert := aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		showForm(ctx, alert, param.Prefix, param.Id, param.Param.GetRouteParamStr(), true)
		return
	}

	if !param.FromList {
		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>location.href="%s"</script>`, param.PreviousPath))
		ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
		return
	}

	buf := showTable(ctx, param.Prefix, param.Path, param.Param)

	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
}
