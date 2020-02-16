package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
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

func ShowForm(conn db.Connection) context.Handler {
	return func(ctx *context.Context) {

		prefix := ctx.Query(constant.PrefixKey)
		panel := table.Get(prefix, ctx)

		if !panel.GetEditable() {
			alert(ctx, panel, "operation not allow", conn)
			ctx.Abort()
			return
		}

		id := ctx.Query(constant.EditPKKey)
		if id == "" {
			alert(ctx, panel, "wrong "+panel.GetPrimaryKey().Name, conn)
			ctx.Abort()
			return
		}

		ctx.SetUserValue("show_form_param", &ShowFormParam{
			Panel:  panel,
			Id:     id,
			Prefix: prefix,
			Param: parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
				panel.GetInfo().GetSort()),
		})
		ctx.Next()
	}
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

func EditForm(srv service.List) context.Handler {
	return func(ctx *context.Context) {
		prefix := ctx.Query(constant.PrefixKey)
		previous := ctx.FormValue("_previous_")
		panel := table.Get(prefix, ctx)
		multiForm := ctx.Request.MultipartForm

		conn := db.GetConnection(srv)

		if !panel.GetEditable() {
			alert(ctx, panel, "operation not allow", conn)
			ctx.Abort()
			return
		}
		token := ctx.FormValue("_t")

		if !auth.GetTokenService(srv.Get(auth.TokenServiceKey)).CheckToken(token) {
			alert(ctx, panel, "edit fail, wrong token", conn)
			ctx.Abort()
			return
		}

		fromList := isInfoUrl(previous)

		param := parameter.GetParamFromUrl(previous, fromList, panel.GetInfo().DefaultPageSize,
			panel.GetPrimaryKey().Name, panel.GetInfo().GetSort())

		if fromList {
			previous = config.Get().Url("/info/" + prefix + param.GetRouteParamStr())
		}

		ctx.SetUserValue("edit_form_param", &EditFormParam{
			Panel:        panel,
			Id:           multiForm.Value[panel.GetPrimaryKey().Name][0],
			Prefix:       prefix,
			Param:        param,
			Path:         strings.Split(previous, "?")[0],
			MultiForm:    multiForm,
			PreviousPath: previous,
			FromList:     fromList,
		})
		ctx.Next()
	}
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
