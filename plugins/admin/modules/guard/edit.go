package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	template2 "html/template"
	"mime/multipart"
	"regexp"
	"strings"
)

type ShowFormParam struct {
	Panel  table.Table
	Id     string
	Prefix string
	Param  parameter.Parameters
}

func (g *Guard) ShowForm(ctx *context.Context) {

	panel, prefix := g.table(ctx)

	if !panel.GetEditable() {
		alert(ctx, panel, "operation not allow", g.conn)
		ctx.Abort()
		return
	}

	id := ctx.Query(constant.EditPKKey)
	if id == "" {
		alert(ctx, panel, "wrong "+panel.GetPrimaryKey().Name, g.conn)
		ctx.Abort()
		return
	}

	ctx.SetUserValue("show_form_param", &ShowFormParam{
		Panel:  panel,
		Id:     id,
		Prefix: prefix,
		Param: parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
			panel.GetInfo().GetSort()).WithPKs(id),
	})
	ctx.Next()
}

func GetShowFormParam(ctx *context.Context) *ShowFormParam {
	return ctx.UserValue["show_form_param"].(*ShowFormParam)
}

type EditFormParam struct {
	Panel        table.Table
	Id           string
	Prefix       string
	Param        parameter.Parameters
	Path         string
	MultiForm    *multipart.Form
	PreviousPath string
	Alert        template2.HTML
	FromList     bool
}

func (e EditFormParam) Value() form.Values {
	return e.MultiForm.Value
}

func (e EditFormParam) HasAlert() bool {
	return e.Alert != template2.HTML("")
}

func (e EditFormParam) IsManage() bool {
	return e.Prefix == "manager"
}

func (e EditFormParam) IsRole() bool {
	return e.Prefix == "roles"
}

func (g *Guard) EditForm(ctx *context.Context) {
	previous := ctx.FormValue(form.PreviousKey)
	panel, prefix := g.table(ctx)
	multiForm := ctx.Request.MultipartForm

	if !panel.GetEditable() {
		alert(ctx, panel, "operation not allow", g.conn)
		ctx.Abort()
		return
	}
	token := ctx.FormValue(form.TokenKey)

	if !auth.GetTokenService(g.services.Get(auth.TokenServiceKey)).CheckToken(token) {
		alert(ctx, panel, "edit fail, wrong token", g.conn)
		ctx.Abort()
		return
	}

	fromList := isInfoUrl(previous)

	param := parameter.GetParamFromURL(previous, panel.GetInfo().DefaultPageSize,
		panel.GetInfo().GetSort(), panel.GetPrimaryKey().Name)

	if fromList {
		previous = config.Get().Url("/info/" + prefix + param.GetRouteParamStr())
	}

	id := multiForm.Value[panel.GetPrimaryKey().Name][0]

	ctx.SetUserValue("edit_form_param", &EditFormParam{
		Panel:        panel,
		Id:           id,
		Prefix:       prefix,
		Param:        param.WithPKs(id),
		Path:         strings.Split(previous, "?")[0],
		MultiForm:    multiForm,
		PreviousPath: previous,
		FromList:     fromList,
	})
	ctx.Next()
}

func isInfoUrl(s string) bool {
	reg, _ := regexp.Compile("(.*?)info/(.*?)$")
	sub := reg.FindStringSubmatch(s)
	return len(sub) > 2 && !strings.Contains(sub[2], "/")
}

func GetEditFormParam(ctx *context.Context) *EditFormParam {
	return ctx.UserValue["edit_form_param"].(*EditFormParam)
}

func alert(ctx *context.Context, panel table.Table, msg string, conn db.Connection) {
	response.Alert(ctx, config.Get(), panel.GetInfo().Description, panel.GetInfo().Title, msg, conn)
}

func alertWithTitleAndDesc(ctx *context.Context, title, desc, msg string, conn db.Connection) {
	response.Alert(ctx, config.Get(), desc, title, msg, conn)
}

func getAlert(msg string) template2.HTML {
	return template.Get(config.Get().Theme).Alert().
		SetTitle(constant.DefaultErrorMsg).
		SetTheme("warning").
		SetContent(template2.HTML(msg)).
		GetContent()
}
