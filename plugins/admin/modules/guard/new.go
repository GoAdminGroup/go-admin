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

func ShowNewForm(conn db.Connection) context.Handler {
	return func(ctx *context.Context) {

		prefix := ctx.Query(constant.PrefixKey)
		panel := table.Get(prefix, ctx)

		if !panel.GetCanAdd() {
			alert(ctx, panel, "operation not allow", conn)
			ctx.Abort()
			return
		}

		ctx.SetUserValue("show_new_form_param", &ShowNewFormParam{
			Panel:  panel,
			Prefix: prefix,
			Param: parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize, panel.GetInfo().SortField,
				panel.GetInfo().GetSort()),
		})
		ctx.Next()
	}
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

func NewForm(srv service.List) context.Handler {
	return func(ctx *context.Context) {
		prefix := ctx.Query(constant.PrefixKey)
		previous := ctx.FormValue("_previous_")
		panel := table.Get(prefix, ctx)

		conn := db.GetConnection(srv)

		if !panel.GetCanAdd() {
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

		param := parameter.GetParamFromUrl(previous, fromList, panel.GetInfo().DefaultPageSize, panel.GetPrimaryKey().Name, panel.GetInfo().GetSort())

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
}

func GetNewFormParam(ctx *context.Context) *NewFormParam {
	return ctx.UserValue["new_form_param"].(*NewFormParam)
}
