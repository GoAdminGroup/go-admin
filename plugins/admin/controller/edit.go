package controller

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	template2 "html/template"
	"net/http"
	"net/url"
)

// ShowForm show form page.
func ShowForm(ctx *context.Context) {
	param := guard.GetShowFormParam(ctx)
	showForm(ctx, "", param.Prefix, param.Id, param.Param.GetRouteParamStr(), false)
}

func showForm(ctx *context.Context, alert template2.HTML, prefix string, id string, paramStr string, isEdit bool) {

	panel := getTable(prefix, ctx)

	user := auth.Auth(ctx)

	infoUrl := routePathWithPrefix("info", prefix) + paramStr
	editUrl := routePathWithPrefix("edit", prefix)
	showEditUrl := routePathWithPrefix("show_edit", prefix) + paramStr

	referer := ctx.Headers("Referer")

	if referer != "" && !isInfoUrl(referer) && !isEditUrl(referer, ctx.Query(constant.PrefixKey)) {
		infoUrl = referer
	}

	newUrl := modules.AorEmpty(panel.GetCanAdd(), routePathWithPrefix("show_new", prefix)+paramStr)
	footerKind := "edit"
	if newUrl == "" || !user.CheckPermissionByUrlMethod(newUrl, route("show_new").Method(), url.Values{}) {
		footerKind = "edit_only"
	}

	formInfo, err := panel.GetDataWithId(id)

	if err != nil && alert == "" {
		alert = aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
	}

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetContent(formInfo.FieldList).
			SetTabContents(formInfo.GroupFieldList).
			SetTabHeaders(formInfo.GroupFieldHeaders).
			SetPrefix(config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(editUrl).
			SetHiddenFields(map[string]string{
				form2.TokenKey:    authSrv().AddToken(),
				form2.PreviousKey: infoUrl,
			}).
			SetOperationFooter(formFooter(footerKind)).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml)),
		Description: formInfo.Description,
		Title:       formInfo.Title,
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

	err := param.Panel.UpdateData(param.Value())
	if err != nil {
		alert := aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
		showForm(ctx, alert, param.Prefix, param.Id, param.Param.GetRouteParamStr(), true)
		return
	}

	if !param.FromList {

		if isNewUrl(param.PreviousPath, param.Prefix) {
			showNewForm(ctx, param.Alert, param.Prefix, param.Param.GetRouteParamStr(), true)
			return
		}

		if isEditUrl(param.PreviousPath, param.Prefix) {
			showForm(ctx, param.Alert, param.Prefix, param.Id, param.Param.GetRouteParamStr(), true)
			return
		}

		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>location.href="%s"</script>`, param.PreviousPath))
		ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
		return
	}

	buf := showTable(ctx, param.Prefix, param.Param)

	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
}
