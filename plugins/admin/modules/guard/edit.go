package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	template2 "html/template"
	"mime/multipart"
	"strings"
)

type ShowFormParam struct {
	Panel  table.Table
	Id     string
	Prefix string
	Param  parameter.Parameters
}

func (e *ShowFormParam) GetUrl() string {
	return config.Get().Url("/edit/" + e.Prefix)
}

func (e *ShowFormParam) GetInfoUrl() string {
	return config.Get().Url("/info/" + e.Prefix + e.Param.GetRouteParamStrWithoutId())
}

func ShowForm(ctx *context.Context) {

	prefix := ctx.Query("__prefix")
	panel := table.List[prefix]
	if !panel.GetEditable() {
		alert(ctx, panel, "operation not allow")
		ctx.Abort()
		return
	}

	id := ctx.Query("id")
	if id == "" {
		alert(ctx, panel, "wrong id")
		ctx.Abort()
		return
	}

	ctx.SetUserValue("show_form_param", &ShowFormParam{
		Panel:  panel,
		Id:     id,
		Prefix: prefix,
		Param:  parameter.GetParam(ctx.Request.URL.Query()),
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
	Previous     string
	Path         string
	MultiForm    *multipart.Form
	PreviousPath string
	Alert        template2.HTML
}

func (e EditFormParam) Value() form.Values {
	return e.MultiForm.Value
}

func (e EditFormParam) GetEditUrl() string {
	return e.getUrl("edit")
}

func (e EditFormParam) HasAlert() bool {
	return e.Alert != template2.HTML("")
}

func (e EditFormParam) GetNewUrl() string {
	return e.getUrl("new")
}

func (e EditFormParam) GetExportUrl() string {
	return config.Get().Url("/export/" + e.Prefix + e.Param.GetRouteParamStr())
}

func (e EditFormParam) GetDeleteUrl() string {
	return config.Get().Url("/delete/" + e.Prefix)
}

func (e *EditFormParam) GetUrl() string {
	return config.Get().Url("/edit/" + e.Prefix)
}

func (e *EditFormParam) GetInfoUrl() string {
	return config.Get().Url("/info/" + e.Prefix + e.Param.GetRouteParamStrWithoutId())
}

func (e EditFormParam) getUrl(kind string) string {
	return config.Get().Url("/info/" + e.Prefix + "/" + kind + e.Param.GetRouteParamStr())
}

func (e EditFormParam) IsManage() bool {
	return e.Prefix == "manager"
}

func (e EditFormParam) IsRole() bool {
	return e.Prefix == "roles"
}

func EditForm(ctx *context.Context) {
	prefix := ctx.Query("__prefix")
	previous := ctx.FormValue("_previous_")
	panel := table.List[prefix]
	multiForm := ctx.Request.MultipartForm

	if !panel.GetEditable() {
		alert(ctx, panel, "operation not allow")
		ctx.Abort()
		return
	}
	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		alert(ctx, panel, "edit fail, wrong token")
		ctx.Abort()
		return
	}

	if prefix != "manager" && prefix != "roles" {
		for _, f := range panel.GetForm().FieldList {
			if f.Editable {
				continue
			}
			if len(multiForm.Value[f.Field]) > 0 && f.Field != "id" {
				alert(ctx, panel, "field["+f.Field+"]is not editable")
				ctx.Abort()
				return
			}
		}
	}

	param := parameter.GetParamFromUrl(previous)

	ctx.SetUserValue("edit_form_param", &EditFormParam{
		Panel:        panel,
		Id:           "",
		Prefix:       prefix,
		Param:        param,
		Path:         strings.Split(previous, "?")[0],
		MultiForm:    multiForm,
		PreviousPath: config.Get().Url("/info/" + prefix + param.GetRouteParamStrWithoutId()),
	})
	ctx.Next()
}

func GetEditFormParam(ctx *context.Context) *EditFormParam {
	return ctx.UserValue["edit_form_param"].(*EditFormParam)
}

func alert(ctx *context.Context, panel table.Table, msg string) {
	response.Alert(ctx, config.Get(), panel.GetInfo().Description, panel.GetInfo().Title, msg)
}

func alertWithTitleAndDesc(ctx *context.Context, title, desc, msg string) {
	response.Alert(ctx, config.Get(), desc, title, msg)
}

func getAlert(msg string) template2.HTML {
	return template.Get(config.Get().Theme).Alert().
		SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
		SetTheme("warning").
		SetContent(template2.HTML(msg)).
		GetContent()
}
