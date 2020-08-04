package main

var pluginTemplate = map[string]string{
	"main": `package {{.PluginName}}

import (
	"{{.ModulePath}}/controller"
	"{{.ModulePath}}/guard"
	language2 "{{.ModulePath}}/modules/language"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type {{.PluginTitle}} struct {
	*plugins.Base

	handler *controller.Handler
	guard   *guard.Guardian

	// ...
}

func init() {
	plugins.Add(&{{.PluginTitle}}{
		Base: &plugins.Base{PlugName: "{{.PluginName}}", URLPrefix: "{{.PluginName}}"},
		// ....
	})
}

func New{{.PluginTitle}}(/*...*/) *{{.PluginTitle}} {
	return &{{.PluginTitle}}{
		Base: &plugins.Base{PlugName: "{{.PluginName}}", URLPrefix: "{{.PluginName}}"},
		// ...
	}
}

func (plug *{{.PluginTitle}}) IsInstalled() bool {
	// ... implement it
	return true
}

func (plug *{{.PluginTitle}}) GetIndexURL() string {
	return config.Url("/{{.PluginName}}/example?param=helloworld")
}

func (plug *{{.PluginTitle}}) InitPlugin(srv service.List) {

	// DO NOT DELETE
	plug.InitBase(srv, "{{.PluginName}}")

	plug.handler = controller.NewHandler(/*...*/)
	plug.guard = guard.New(/*...*/)
	plug.App = plug.initRouter(srv)
	plug.handler.HTML = plug.HTML
	plug.handler.HTMLMenu = plug.HTMLMenu

	language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)

	plug.SetInfo(info)
}

var info = plugins.Info{
	Website:     "",
	Title:       "{{.PluginTitle}}",
	Description: "",
	Version:     "",
	Author:      "",
	Url:         "",
	Cover:       "",
	Agreement:   "",
	Uuid:        "",
	Name:        "",
	ModulePath:  "",
	CreateDate:  utils.ParseTime("2000-01-11"),
	UpdateDate:  utils.ParseTime("2000-01-11"),
}

func (plug *{{.PluginTitle}}) GetSettingPage() table.Generator {
	return func(ctx *context.Context) ({{.PluginName}}Configuration table.Table) {

		cfg := table.DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver)

		if !plug.IsInstalled() {
			cfg = cfg.SetOnlyNewForm()
		} else {
			cfg = cfg.SetOnlyUpdateForm()
		}

		{{.PluginName}}Configuration = table.NewDefaultTable(cfg)

		formList := {{.PluginName}}Configuration.GetForm().
			AddXssJsFilter().
			HideBackButton().
			HideContinueNewCheckBox().
			HideResetButton()

		// formList.AddField(...)

		formList.SetInsertFn(func(values form2.Values) error {
			// TODO: finish the installation
			return nil
		})

		formList.EnableAjaxData(types.AjaxData{
			SuccessTitle:   language2.Get("install success"),
			ErrorTitle:     language2.Get("install fail"),
			SuccessJumpURL: "...",
		}).SetFormNewTitle(language2.GetHTML("{{.PluginName}} installation")).
			SetTitle(language2.Get("{{.PluginName}} installation")).
			SetFormNewBtnWord(language2.GetHTML("install"))

		return
	}
}
`,
	"controller": `package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Handler struct {
	HTML     func(ctx *context.Context, panel types.Panel, ops ...template.ExecuteOptions)
	HTMLMenu func(ctx *context.Context, panel types.Panel, ops ...template.ExecuteOptions)
}

func NewHandler(/*...*/) *Handler {
	return &Handler{
		// ...
	}
}

func (h *Handler) Update(/*...*/) {
	// ...
}`,
	"controller_example": `package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"{{.ModulePath}}/guard"
)

func (h *Handler) Example(ctx *context.Context) {
	var param = guard.GetExampleParam(ctx)
	h.HTML(ctx, types.Panel{
		Title:       "Example",
		Description: "example",
		Content:     template.HTML(param.Param),
	})
}
`,
	"guard": `package guard

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
)

type Guardian struct {
	conn db.Connection
}

func New(/*...*/) *Guardian {
	return &Guardian{
		// ...
	}
}

func (g *Guardian) Update(/*...*/) {
	// ...
}`,
	"guard_example": `package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
)

type ExampleParam struct {
	Param string
}

func (g *Guardian) Example(ctx *context.Context) {

	param := ctx.Query("param")

	ctx.SetUserValue("example", &ExampleParam{
		Param: param,
	})
	ctx.Next()
}

func GetExampleParam(ctx *context.Context) *ExampleParam {
	return ctx.UserValue["example"].(*ExampleParam)
}`,
	"makefile": `GOCMD = go

all: fmt

fmt:
	GO111MODULE=off $(GOCMD) fmt ./...

test:
	gotest -v ./tests`,
	"router": `package {{.PluginName}}

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/service"
)

func (plug *{{.PluginTitle}}) initRouter(srv service.List) *context.App {

	app := context.NewApp()
	authRoute := app.Group("/", auth.Middleware(plug.Conn))
	
	authRoute.GET("/example", plug.guard.Example, plug.handler.Example)

	return app
}`,
	"language": `package language

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"html/template"
)

func Get(key string) string {
	return language.GetWithScope(key, "{{.PluginName}}")
}

func GetHTML(key string) template.HTML {
	return template.HTML(language.GetWithScope(key, "{{.PluginName}}"))
}`,
	"language_cn": `package language

import "github.com/GoAdminGroup/go-admin/modules/language"

var CN = language.LangSet{
	"{{.PluginName}}.aaa": "aaa",
}`,
	"language_en": `package language

import "github.com/GoAdminGroup/go-admin/modules/language"

var EN = language.LangSet{
	"{{.PluginName}}.aaa": "aaa",
}`,
}
