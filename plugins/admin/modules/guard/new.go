package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"html/template"
	"mime/multipart"
	"strings"
)

type ShowNewFormParam struct {
	Panel  table.Table
	Prefix string
	Param  parameter.Parameters
}

func (g *Guard) ShowNewForm(ctx *context.Context) {

	panel, prefix := g.table(ctx)

	if !panel.GetCanAdd() {
		alert(ctx, panel, "operation not allow", g.conn)
		ctx.Abort()
		return
	}

	ctx.SetUserValue("show_new_form_param", &ShowNewFormParam{
		Panel:  panel,
		Prefix: prefix,
		Param: parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
			panel.GetInfo().GetSort()),
	})
	ctx.Next()
}

func GetShowNewFormParam(ctx *context.Context) *ShowNewFormParam {
	return ctx.UserValue["show_new_form_param"].(*ShowNewFormParam)
}

type NewFormParam struct {
	Panel        table.Table
	Id           string
	Prefix       string
	Param        parameter.Parameters
	Path         string
	MultiForm    *multipart.Form
	PreviousPath string
	FromList     bool
	Alert        template.HTML
}

func (e NewFormParam) Value() form.Values {
	return e.MultiForm.Value
}

func (e NewFormParam) IsManage() bool {
	return e.Prefix == "manager"
}

func (e NewFormParam) HasAlert() bool {
	return e.Alert != template.HTML("")
}

func (e NewFormParam) IsRole() bool {
	return e.Prefix == "roles"
}

func (g *Guard) NewForm(ctx *context.Context) {
	previous := ctx.FormValue(form.PreviousKey)
	panel, prefix := g.table(ctx)

	conn := db.GetConnection(g.services)

	if !panel.GetCanAdd() {
		alert(ctx, panel, "operation not allow", conn)
		ctx.Abort()
		return
	}
	token := ctx.FormValue(form.TokenKey)

	if !auth.GetTokenService(g.services.Get(auth.TokenServiceKey)).CheckToken(token) {
		alert(ctx, panel, "create fail, wrong token", conn)
		ctx.Abort()
		return
	}

	fromList := isInfoUrl(previous)

	param := parameter.GetParamFromURL(previous, panel.GetInfo().DefaultPageSize,
		panel.GetInfo().GetSort(), panel.GetPrimaryKey().Name)

	if fromList {
		previous = config.Get().Url("/info/" + prefix + param.GetRouteParamStr())
	}

	ctx.SetUserValue("new_form_param", &NewFormParam{
		Panel:        panel,
		Id:           "",
		Prefix:       prefix,
		Param:        param,
		Path:         strings.Split(previous, "?")[0],
		MultiForm:    ctx.Request.MultipartForm,
		PreviousPath: previous,
		FromList:     fromList,
	})
	ctx.Next()
}

func GetNewFormParam(ctx *context.Context) *NewFormParam {
	return ctx.UserValue["new_form_param"].(*NewFormParam)
}
