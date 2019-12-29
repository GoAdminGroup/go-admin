package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	config2 "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	template2 "html/template"
	"net/http"
)

func ShowDetail(ctx *context.Context) {
	prefix := ctx.Query("__prefix")
	id := ctx.Query("id")
	panel := table.Get(prefix)
	user := auth.Auth(ctx)

	newPanel := panel.Copy()

	formModel := newPanel.GetForm()

	formModel.FieldList = make([]types.FormField, len(panel.GetInfo().FieldList))

	for i, field := range panel.GetInfo().FieldList {
		formModel.FieldList[i] = types.FormField{
			Field:        field.Field,
			TypeName:     field.TypeName,
			Head:         field.Head,
			FormType:     form.Default,
			FieldDisplay: field.FieldDisplay,
		}
	}

	formData, _, _, _, _, err := newPanel.GetDataFromDatabaseWithId(id)

	var alert template2.HTML

	if err != nil && alert == "" {
		alert = aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
	}

	param := parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize, panel.GetPrimaryKey().Name,
		panel.GetInfo().GetSort())

	title := language.Get("Detail")

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetTitle(template.HTML(title)).
			SetContent(formData).
			SetInfoUrl(config2.Get().Url("/info/"+prefix+param.GetRouteParamStrWithoutId())).
			SetPrefix(config.PrefixFixSlash())),
		Description: title,
		Title:       title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}
