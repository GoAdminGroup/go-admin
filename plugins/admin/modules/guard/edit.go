package guard

import (
	tmpl "html/template"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/GoAdminGroup/go-admin/template/types"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
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
		alert(ctx, panel, errors.OperationNotAllow, g.conn, g.navBtns)
		ctx.Abort()
		return
	}

	if panel.GetOnlyInfo() {
		ctx.Redirect(config.Url("/info/" + prefix))
		ctx.Abort()
		return
	}

	if panel.GetOnlyDetail() {
		ctx.Redirect(config.Url("/info/" + prefix + "/detail"))
		ctx.Abort()
		return
	}

	if panel.GetOnlyNewForm() {
		ctx.Redirect(config.Url("/info/" + prefix + "/new"))
		ctx.Abort()
		return
	}

	id := ctx.Query(constant.EditPKKey)

	if id == "" {
		id = "1"
	}

	ctx.SetUserValue(showFormParamKey, &ShowFormParam{
		Panel:  panel,
		Id:     id,
		Prefix: prefix,
		Param: parameter.GetParam(ctx.Request.URL, panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
			panel.GetInfo().GetSort()).WithPKs(id),
	})
	ctx.Next()
}

func GetShowFormParam(ctx *context.Context) *ShowFormParam {
	return ctx.UserValue[showFormParamKey].(*ShowFormParam)
}

type EditFormParam struct {
	Panel        table.Table
	Id           string
	Prefix       string
	Param        parameter.Parameters
	Path         string
	MultiForm    *multipart.Form
	PreviousPath string
	Alert        tmpl.HTML
	FromList     bool
	IsIframe     bool
	IframeID     string
}

func (e EditFormParam) Value() form.Values {
	return e.MultiForm.Value
}

func (g *Guard) EditForm(ctx *context.Context) {

	panel, prefix := g.table(ctx)

	if !panel.GetEditable() {
		alert(ctx, panel, errors.OperationNotAllow, g.conn, g.navBtns)
		ctx.Abort()
		return
	}
	token := ctx.FormValue(form.TokenKey)

	if !auth.GetTokenService(g.services.Get(auth.TokenServiceKey)).CheckToken(token) {
		alert(ctx, panel, errors.EditFailWrongToken, g.conn, g.navBtns)
		ctx.Abort()
		return
	}

	var (
		previous = ctx.FormValue(form.PreviousKey)
		fromList = isInfoUrl(previous)
		param    = parameter.GetParamFromURL(previous, panel.GetInfo().DefaultPageSize,
			panel.GetInfo().GetSort(), panel.GetPrimaryKey().Name)
	)

	if fromList {
		previous = config.Url("/info/" + prefix + param.GetRouteParamStr())
	}

	var (
		multiForm = ctx.Request.MultipartForm
		id        = multiForm.Value[panel.GetPrimaryKey().Name][0]
		values    = ctx.Request.MultipartForm.Value
	)

	ctx.SetUserValue(editFormParamKey, &EditFormParam{
		Panel:        panel,
		Id:           id,
		Prefix:       prefix,
		Param:        param.WithPKs(id),
		Path:         strings.Split(previous, "?")[0],
		MultiForm:    multiForm,
		IsIframe:     form.Values(values).Get(constant.IframeKey) == "true",
		IframeID:     form.Values(values).Get(constant.IframeIDKey),
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
	return ctx.UserValue[editFormParamKey].(*EditFormParam)
}

func alert(ctx *context.Context, panel table.Table, msg string, conn db.Connection, btns *types.Buttons) {
	if ctx.WantJSON() {
		response.BadRequest(ctx, msg)
	} else {
		response.Alert(ctx, panel.GetInfo().Description, panel.GetInfo().Title, msg, conn, btns)
	}
}

func alertWithTitleAndDesc(ctx *context.Context, title, desc, msg string, conn db.Connection, btns *types.Buttons) {
	response.Alert(ctx, desc, title, msg, conn, btns)
}

func getAlert(msg string) tmpl.HTML {
	return template.Get(config.GetTheme()).Alert().Warning(msg)
}
