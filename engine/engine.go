// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	template2 "html/template"
	"net/http"
	"strings"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
)

// Engine is the core component of goAdmin. It has two attributes.
// PluginList is an array of plugin. Adapter is the adapter of
// web framework context and goAdmin context. The relationship of adapter and
// plugin is that the adapter use the plugin which contains routers and
// controller methods to inject into the framework entity and make it work.
type Engine struct {
	PluginList []plugins.Plugin
	Adapter    adapter.WebFrameWork
	Services   service.List
	NavButtons types.Buttons

	config *config.Config
}

// Default return the default engine instance.
func Default() *Engine {
	return &Engine{
		Adapter:  defaultAdapter,
		Services: service.GetServices(),
	}
}

// Use enable the adapter.
func (eng *Engine) Use(router interface{}) error {
	if eng.Adapter == nil {
		panic("adapter is nil, import the default adapter or use AddAdapter method add the adapter")
	}

	if len(eng.PluginList) == 0 {
		eng.PluginList = append(eng.PluginList, admin.NewAdmin())
	}

	eng.Services.Add("config", config.SrvWithConfig(eng.config))
	errors.Init()

	if !eng.config.HideConfigCenterEntrance {
		btn := types.GetNavButton("", icon.Gear, action.Jump(eng.config.Url("/info/site/edit")))
		eng.AdminPlugin().AddNavButton(btn)
		eng.NavButtons = append(eng.NavButtons, btn)
		navButtons = append(navButtons, btn)
	}

	// Initialize plugins
	for i := range eng.PluginList {
		eng.PluginList[i].InitPlugin(eng.Services)
	}

	return eng.Adapter.Use(router, eng.PluginList)
}

// AddPlugins add the plugins and initialize them.
func (eng *Engine) AddPlugins(plugs ...plugins.Plugin) *Engine {

	if len(plugs) == 0 {
		panic("wrong plugins")
	}

	eng.PluginList = append(eng.PluginList, plugs...)

	return eng
}

func (eng *Engine) FindPluginByName(name string) (plugins.Plugin, bool) {
	for _, plug := range eng.PluginList {
		if plug.Name() == name {
			return plug, true
		}
	}
	return nil, false
}

func (eng *Engine) AddAuthService(processor auth.Processor) *Engine {
	eng.Services.Add("auth", auth.NewService(processor))
	return eng
}

// ============================
// Config APIs
// ============================

// AddConfig set the global config.
func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	return eng.setConfig(cfg).InitDatabase()
}

// setConfig set the config of engine.
func (eng *Engine) setConfig(cfg config.Config) *Engine {
	eng.config = config.Set(cfg)
	if !template.CheckRequirements() {
		panic(fmt.Sprintf("Wrong GoAdmin version, theme %s required GoAdmin version are %s",
			cfg.Theme, strings.Join(template.Default().GetRequirements(), ",")))
	}
	return eng
}

// AddConfigFromJSON set the global config from json file.
func (eng *Engine) AddConfigFromJSON(path string) *Engine {
	return eng.setConfig(config.ReadFromJson(path)).InitDatabase()
}

// AddConfigFromYAML set the global config from yaml file.
func (eng *Engine) AddConfigFromYAML(path string) *Engine {
	return eng.setConfig(config.ReadFromYaml(path)).InitDatabase()
}

// AddConfigFromINI set the global config from ini file.
func (eng *Engine) AddConfigFromINI(path string) *Engine {
	return eng.setConfig(config.ReadFromINI(path)).InitDatabase()
}

// InitDatabase initialize all database connection.
func (eng *Engine) InitDatabase() *Engine {
	for driver, databaseCfg := range eng.config.Databases.GroupByDriver() {
		eng.Services.Add(driver, db.GetConnectionByDriver(driver).InitDB(databaseCfg))
	}
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
	defaultConnection := db.GetConnection(eng.Services)
	defaultAdapter.SetConnection(defaultConnection)
	eng.Adapter.SetConnection(defaultConnection)
	return eng
}

// AddAdapter add the adapter of engine.
func (eng *Engine) AddAdapter(ada adapter.WebFrameWork) *Engine {
	eng.Adapter = ada
	defaultAdapter = ada
	return eng
}

// defaultAdapter is the default adapter of engine.
var defaultAdapter adapter.WebFrameWork

var navButtons = make([]types.Button, 0)

// Register set default adapter of engine.
func Register(ada adapter.WebFrameWork) {
	if ada == nil {
		panic("adapter is nil")
	}
	defaultAdapter = ada
}

// User call the User method of defaultAdapter.
func User(ci interface{}) (models.UserModel, bool) {
	return defaultAdapter.User(ci)
}

// User call the User method of engine adapter.
func (eng *Engine) User(ci interface{}) (models.UserModel, bool) {
	return eng.Adapter.User(ci)
}

// ============================
// DB Connection APIs
// ============================

// db return the db connection of given driver.
func (eng *Engine) DB(driver string) db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(driver))
}

func (eng *Engine) DefaultConnection() db.Connection {
	return eng.DB(eng.config.Databases.GetDefault().Driver)
}

// MysqlConnection return the mysql db connection of given driver.
func (eng *Engine) MysqlConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverMysql))
}

// MssqlConnection return the mssql db connection of given driver.
func (eng *Engine) MssqlConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverMssql))
}

// PostgresqlConnection return the postgresql db connection of given driver.
func (eng *Engine) PostgresqlConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverPostgresql))
}

// SqliteConnection return the sqlite db connection of given driver.
func (eng *Engine) SqliteConnection() db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(db.DriverSqlite))
}

type ConnectionSetter func(db.Connection)

func (eng *Engine) ResolveConnection(setter ConnectionSetter, driver string) *Engine {
	setter(eng.DB(driver))
	return eng
}

func (eng *Engine) ResolveMysqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverMysql)
	return eng
}

func (eng *Engine) ResolveMssqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverMssql)
	return eng
}

func (eng *Engine) ResolveSqliteConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverSqlite)
	return eng
}

func (eng *Engine) ResolvePostgresqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverPostgresql)
	return eng
}

type Setter func(*Engine)

func (eng *Engine) Clone(e *Engine) *Engine {
	e = eng
	return eng
}

func (eng *Engine) ClonedBySetter(setter Setter) *Engine {
	setter(eng)
	return eng
}

func (eng *Engine) wrapWithAuthMiddleware(handler context.Handler) context.Handlers {
	return []context.Handler{response.OffLineHandler, auth.Middleware(db.GetConnection(eng.Services)), handler}
}

// ============================
// HTML Content Render APIs
// ============================

func (eng *Engine) AddNavButtons(title template2.HTML, icon string, action types.Action) *Engine {
	btn := types.GetNavButton(title, icon, action)
	eng.AdminPlugin().AddNavButton(btn)
	eng.NavButtons = append(eng.NavButtons, btn)
	navButtons = append(navButtons, btn)
	return eng
}

// Content call the Content method of engine adapter.
// If adapter is nil, it will panic.
func (eng *Engine) Content(ctx interface{}, panel types.GetPanelFn) {
	if eng.Adapter == nil {
		panic("adapter is nil")
	}
	eng.Adapter.Content(ctx, panel, eng.NavButtons...)
}

// Content call the Content method of defaultAdapter.
// If defaultAdapter is nil, it will panic.
func Content(ctx interface{}, panel types.GetPanelFn) {
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
	defaultAdapter.Content(ctx, panel, navButtons...)
}

func (eng *Engine) Data(method, url string, handler context.Handler) {
	eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(handler))
}

func (eng *Engine) HTML(method, url string, fn types.GetPanelInfoFn) {

	eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(func(ctx *context.Context) {
		panel, err := fn(ctx)
		if err != nil {
			panel = template.WarningPanel(err.Error())
		}

		tmpl, tmplName := template.Default().GetTemplate(ctx.IsPjax())

		user := auth.Auth(ctx)

		buf := new(bytes.Buffer)
		hasError := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
			User:    user,
			Menu:    menu.GetGlobalMenu(user, eng.Adapter.GetConnection()).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
			Panel:   panel.GetContent(eng.config.IsProductionEnvironment()),
			Assets:  template.GetComponentAssetListsHTML(),
			Buttons: eng.NavButtons.CheckPermission(user),
		}))

		if hasError != nil {
			logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
		}

		ctx.HTMLByte(http.StatusOK, buf.Bytes())
	}))
}

func (eng *Engine) HTMLFile(method, url, path string, data map[string]interface{}) {
	eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(func(ctx *context.Context) {

		cbuf := new(bytes.Buffer)

		t, err := template2.ParseFiles(path)
		if err != nil {
			eng.errorPanelHTML(ctx, cbuf, err)
		} else {
			if err := t.Execute(cbuf, data); err != nil {
				eng.errorPanelHTML(ctx, cbuf, err)
			}
		}

		tmpl, tmplName := template.Default().GetTemplate(ctx.IsPjax())

		user := auth.Auth(ctx)

		buf := new(bytes.Buffer)
		hasError := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
			User: user,
			Menu: menu.GetGlobalMenu(user, eng.Adapter.GetConnection()).SetActiveClass(eng.config.URLRemovePrefix(ctx.Path())),
			Panel: types.Panel{
				Content: template.HTML(cbuf.String()),
			},
			Assets:  template.GetComponentAssetListsHTML(),
			Buttons: eng.NavButtons.CheckPermission(user),
		}))

		if hasError != nil {
			logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
		}

		ctx.HTMLByte(http.StatusOK, buf.Bytes())
	}))
}

func (eng *Engine) HTMLFiles(method, url string, data map[string]interface{}, files ...string) {
	eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(func(ctx *context.Context) {

		cbuf := new(bytes.Buffer)

		t, err := template2.ParseFiles(files...)
		if err != nil {
			eng.errorPanelHTML(ctx, cbuf, err)
		} else {
			if err := t.Execute(cbuf, data); err != nil {
				eng.errorPanelHTML(ctx, cbuf, err)
			}
		}

		tmpl, tmplName := template.Default().GetTemplate(ctx.IsPjax())

		user := auth.Auth(ctx)

		buf := new(bytes.Buffer)
		hasError := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
			User: user,
			Menu: menu.GetGlobalMenu(user, eng.Adapter.GetConnection()).SetActiveClass(eng.config.URLRemovePrefix(ctx.Path())),
			Panel: types.Panel{
				Content: template.HTML(cbuf.String()),
			},
			Assets:  template.GetComponentAssetListsHTML(),
			Buttons: eng.NavButtons.CheckPermission(user),
		}))

		if hasError != nil {
			logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
		}

		ctx.HTMLByte(http.StatusOK, buf.Bytes())
	}))
}

func (eng *Engine) errorPanelHTML(ctx *context.Context, buf *bytes.Buffer, err error) {

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Default().GetTemplate(ctx.IsPjax())

	hasError := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
		User:    user,
		Menu:    menu.GetGlobalMenu(user, eng.Adapter.GetConnection()).SetActiveClass(eng.config.URLRemovePrefix(ctx.Path())),
		Panel:   template.WarningPanel(err.Error()).GetContent(eng.config.IsProductionEnvironment()),
		Assets:  template.GetComponentAssetListsHTML(),
		Buttons: eng.NavButtons.CheckPermission(user),
	}))

	if hasError != nil {
		logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
	}
}

// ============================
// Admin Plugin APIs
// ============================

func (eng *Engine) AddGenerators(list ...table.GeneratorList) *Engine {
	eng.PluginList = append(eng.PluginList, admin.NewAdmin(list...))
	return eng
}

func (eng *Engine) AdminPlugin() *admin.Admin {
	plug, exist := eng.FindPluginByName("admin")
	if exist {
		return plug.(*admin.Admin)
	}
	adm := admin.NewAdmin()
	eng.PluginList = append(eng.PluginList, adm)
	return adm
}

// AddGenerator add table model generator.
func (eng *Engine) AddGenerator(key string, g table.Generator) *Engine {
	eng.AdminPlugin().AddGenerator(key, g)
	return eng
}

// AddGlobalDisplayProcessFn call types.AddGlobalDisplayProcessFn
func (eng *Engine) AddGlobalDisplayProcessFn(f types.DisplayProcessFn) *Engine {
	types.AddGlobalDisplayProcessFn(f)
	return eng
}

// AddDisplayFilterLimit call types.AddDisplayFilterLimit
func (eng *Engine) AddDisplayFilterLimit(limit int) *Engine {
	types.AddLimit(limit)
	return eng
}

// AddDisplayFilterTrimSpace call types.AddDisplayFilterTrimSpace
func (eng *Engine) AddDisplayFilterTrimSpace() *Engine {
	types.AddTrimSpace()
	return eng
}

// AddDisplayFilterSubstr call types.AddDisplayFilterSubstr
func (eng *Engine) AddDisplayFilterSubstr(start int, end int) *Engine {
	types.AddSubstr(start, end)
	return eng
}

// AddDisplayFilterToTitle call types.AddDisplayFilterToTitle
func (eng *Engine) AddDisplayFilterToTitle() *Engine {
	types.AddToTitle()
	return eng
}

// AddDisplayFilterToUpper call types.AddDisplayFilterToUpper
func (eng *Engine) AddDisplayFilterToUpper() *Engine {
	types.AddToUpper()
	return eng
}

// AddDisplayFilterToLower call types.AddDisplayFilterToLower
func (eng *Engine) AddDisplayFilterToLower() *Engine {
	types.AddToUpper()
	return eng
}

// AddDisplayFilterXssFilter call types.AddDisplayFilterXssFilter
func (eng *Engine) AddDisplayFilterXssFilter() *Engine {
	types.AddXssFilter()
	return eng
}

// AddDisplayFilterXssJsFilter call types.AddDisplayFilterXssJsFilter
func (eng *Engine) AddDisplayFilterXssJsFilter() *Engine {
	types.AddXssJsFilter()
	return eng
}
