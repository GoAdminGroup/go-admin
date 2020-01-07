package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
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

func (e *ShowNewFormParam) GetUrl() string {
	return config.Get().Url("/new/" + e.Prefix)
}

func (e *ShowNewFormParam) GetInfoUrl() string {
	return config.Get().Url("/info/" + e.Prefix + e.Param.GetRouteParamStr())
}

func ShowNewForm(conn db.Connection) context.Handler {
	return func(ctx *context.Context) {

		prefix := ctx.Query("__prefix")
		panel := table.Get(prefix)

		if !panel.GetCanAdd() {
			alert(ctx, panel, "operation not allow", conn)
			ctx.Abort()
			return
		}

		ctx.SetUserValue("show_new_form_param", &ShowNewFormParam{
			Panel:  panel,
			Prefix: prefix,
			Param: parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize, panel.GetPrimaryKey().Name,
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

func (e NewFormParam) GetEditUrl() string {
	return e.getUrl("edit")
}

func (e NewFormParam) GetNewUrl() string {
	return e.getUrl("new")
}

func (e NewFormParam) GetUpdateUrl() string {
	return config.Get().Url("/update/" + e.Prefix)
}

func (e NewFormParam) GetDeleteUrl() string {
	return config.Get().Url("/delete/" + e.Prefix)
}

func (e NewFormParam) GetExportUrl() string {
	return config.Get().Url("/export/" + e.Prefix + e.Param.GetRouteParamStr())
}

func (e NewFormParam) getUrl(kind string) string {
	return config.Get().Url("/info/" + e.Prefix + "/" + kind + e.Param.GetRouteParamStr())
}

func (e NewFormParam) IsManage() bool {
	return e.Prefix == "manager"
}

func (e *NewFormParam) GetUrl() string {
	return config.Get().Url("/new/" + e.Prefix)
}

func (e *NewFormParam) GetInfoUrl() string {
	return config.Get().Url("/info/" + e.Prefix + e.Param.GetRouteParamStr())
}

func (e *NewFormParam) GetDetailUrl() string {
	return config.Get().Url("/info/" + e.Prefix + "/detail" + e.Param.GetRouteParamStr())
}

func (e NewFormParam) HasAlert() bool {
	return e.Alert != template.HTML("")
}

func (e NewFormParam) IsRole() bool {
	return e.Prefix == "roles"
}

func NewForm(srv service.List) context.Handler {
	return func(ctx *context.Context) {
		prefix := ctx.Query("__prefix")
		previous := ctx.FormValue("_previous_")
		panel := table.Get(prefix)

		conn := db.GetConnection(srv)

		if !panel.GetCanAdd() {
			alert(ctx, panel, "operation not allow", conn)
			ctx.Abort()
			return
		}
		token := ctx.FormValue("_t")

		if !auth.GetService(srv.Get("auth")).CheckToken(token) {
			alert(ctx, panel, "edit fail, wrong token", conn)
			ctx.Abort()
			return
		}

		fromList := modules.IsInfoUrl(previous)

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
