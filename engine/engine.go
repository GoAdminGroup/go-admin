// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

import (
	"bytes"
	"encoding/json"
	errors2 "errors"
	"fmt"
	template2 "html/template"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"

	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/modules/ui"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Engine is the core component of goAdmin. It has two attributes.
// PluginList is an array of plugin. Adapter is the adapter of
// web framework context and goAdmin context. The relationship of adapter and
// plugin is that the adapter use the plugin which contains routers and
// controller methods to inject into the framework entity and make it work.
type Engine struct {
	PluginList plugins.Plugins
	Adapter    adapter.WebFrameWork
	Services   service.List
	NavButtons *types.Buttons
	config     *config.Config
}

// Default return the default engine instance.
func Default() *Engine {
	engine = &Engine{
		Adapter:    defaultAdapter,
		Services:   service.GetServices(),
		NavButtons: new(types.Buttons),
	}
	return engine
}

// Use enable the adapter.
func (eng *Engine) Use(router interface{}) error {
	if eng.Adapter == nil {
		panic("adapter is nil, import the default adapter or use AddAdapter method add the adapter")
	}

	eng.PluginList = eng.PluginList.Add(admin.NewAdmin())

	eng.AddPluginList(plugins.Get())

	logger.Info("=====> " + language.Get("initialize configuration"))

	// init site setting
	site := models.Site().SetConn(eng.DefaultConnection())
	site.Init(eng.config.ToMap())
	_ = eng.config.Update(site.AllToMap())
	eng.Services.Add("config", config.SrvWithConfig(eng.config))

	logger.Info("=====> " + language.Get("initialize error"))

	errors.Init()

	logger.Info("=====> " + language.Get("initialize navigation buttons"))

	if !eng.config.HideConfigCenterEntrance {
		*eng.NavButtons = (*eng.NavButtons).AddNavButton(icon.Gear, types.NavBtnSiteName,
			action.JumpInNewTab(config.Url("/info/site/edit"),
				language.GetWithScope("site setting", "config")))
	}

	if !eng.config.HideToolEntrance {
		*eng.NavButtons = (*eng.NavButtons).AddNavButton(icon.Wrench, types.NavBtnToolName,
			action.JumpInNewTab(config.Url("/info/generate/new"),
				language.GetWithScope("tool", "tool")))
	}

	if !eng.config.HideAppInfoEntrance {
		*eng.NavButtons = (*eng.NavButtons).AddNavButton(icon.Info, types.NavBtnInfoName,
			action.JumpInNewTab(config.Url("/application/info"),
				language.GetWithScope("system info", "system")))
	}

	if !eng.config.HidePluginEntrance {
		*eng.NavButtons = (*eng.NavButtons).AddNavButton(icon.Th, types.NavBtnPlugName,
			action.JumpInNewTab(config.Url("/plugins"),
				language.GetWithScope("plugin", "plugin")))
	}

	navButtons = eng.NavButtons

	eng.Services.Add(ui.ServiceKey, ui.NewService(eng.NavButtons))

	defaultConnection := db.GetConnection(eng.Services)
	defaultAdapter.SetConnection(defaultConnection)
	eng.Adapter.SetConnection(defaultConnection)

	logger.Info("=====> " + language.Get("initialize plugins"))

	var plugGenerators = make(table.GeneratorList)

	// Initialize plugins
	for i := range eng.PluginList {
		if eng.PluginList[i].Name() != "admin" {
			logger.Info("=====> " + eng.PluginList[i].Name())
			eng.PluginList[i].InitPlugin(eng.Services)
			skip, gen := eng.PluginList[i].GetInstallationPage()
			if !skip && gen != nil {
				eng.AddGenerator("plugin_"+eng.PluginList[i].Name(), gen)
			}
			plugGenerators = plugGenerators.Combine(eng.PluginList[i].GetGenerators())
		}
	}
	adm := eng.AdminPlugin().AddGenerators(plugGenerators)
	adm.InitPlugin(eng.Services)
	plugins.Add(adm)

	return eng.Adapter.Use(router, eng.PluginList)
}

// AddPlugins add the plugins
func (eng *Engine) AddPlugins(plugs ...plugins.Plugin) *Engine {

	if len(plugs) == 0 {
		return eng
	}

	for _, plug := range plugs {
		eng.PluginList = eng.PluginList.Add(plug)
	}

	return eng
}

// AddPluginList add the plugins
func (eng *Engine) AddPluginList(plugs plugins.Plugins) *Engine {

	if len(plugs) == 0 {
		return eng
	}

	for _, plug := range plugs {
		eng.PluginList = eng.PluginList.Add(plug)
	}

	return eng
}

// FindPluginByName find the register plugin by given name.
func (eng *Engine) FindPluginByName(name string) (plugins.Plugin, bool) {
	for _, plug := range eng.PluginList {
		if plug.Name() == name {
			return plug, true
		}
	}
	return nil, false
}

// AddAuthService customize the auth logic with given callback function.
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
	sysCheck, themeCheck := template.CheckRequirements()
	if !sysCheck {
		panic(fmt.Sprintf("wrong GoAdmin version, theme %s required GoAdmin version are %s",
			eng.config.Theme, strings.Join(template.Default().GetRequirements(), ",")))
	}
	if !themeCheck {
		panic(fmt.Sprintf("wrong Theme version, GoAdmin %s required Theme version are %s",
			system.Version(), strings.Join(system.RequireThemeVersion()[eng.config.Theme], ",")))
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
	logger.Info("=====> " + language.Get("initialize database connections"))
	for driver, databaseCfg := range eng.config.Databases.GroupByDriver() {
		eng.Services.Add(driver, db.GetConnectionByDriver(driver).InitDB(databaseCfg))
	}
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
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

var engine *Engine

// navButtons is the default buttons in the navigation bar.
var navButtons = new(types.Buttons)

// Register set default adapter of engine.
func Register(ada adapter.WebFrameWork) {
	if ada == nil {
		panic("adapter is nil")
	}
	defaultAdapter = ada
}

// User call the User method of defaultAdapter.
func User(ctx interface{}) (models.UserModel, bool) {
	return defaultAdapter.User(ctx)
}

// User call the User method of engine adapter.
func (eng *Engine) User(ctx interface{}) (models.UserModel, bool) {
	return eng.Adapter.User(ctx)
}

// ============================
// DB Connection APIs
// ============================

// DB return the db connection of given driver.
func (eng *Engine) DB(driver string) db.Connection {
	return db.GetConnectionFromService(eng.Services.Get(driver))
}

// DefaultConnection return the default db connection.
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

// ResolveConnection resolve the specified driver connection.
func (eng *Engine) ResolveConnection(setter ConnectionSetter, driver string) *Engine {
	setter(eng.DB(driver))
	return eng
}

// ResolveMysqlConnection resolve the mysql connection.
func (eng *Engine) ResolveMysqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverMysql)
	return eng
}

// ResolveMssqlConnection resolve the mssql connection.
func (eng *Engine) ResolveMssqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverMssql)
	return eng
}

// ResolveSqliteConnection resolve the sqlite connection.
func (eng *Engine) ResolveSqliteConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverSqlite)
	return eng
}

// ResolvePostgresqlConnection resolve the postgres connection.
func (eng *Engine) ResolvePostgresqlConnection(setter ConnectionSetter) *Engine {
	eng.ResolveConnection(setter, db.DriverPostgresql)
	return eng
}

type Setter func(*Engine)

// Clone copy a new Engine.
func (eng *Engine) Clone(e *Engine) *Engine {
	e = eng
	return eng
}

// ClonedBySetter copy a new Engine by a setter callback function.
func (eng *Engine) ClonedBySetter(setter Setter) *Engine {
	setter(eng)
	return eng
}

func (eng *Engine) deferHandler(conn db.Connection) context.Handler {
	return func(ctx *context.Context) {
		defer func(ctx *context.Context) {
			if user, ok := ctx.UserValue["user"].(models.UserModel); ok {
				var input []byte
				form := ctx.Request.MultipartForm
				if form != nil {
					input, _ = json.Marshal((*form).Value)
				}

				models.OperationLog().SetConn(conn).New(user.Id, ctx.Path(), ctx.Method(), ctx.LocalIP(), string(input))
			}

			if err := recover(); err != nil {
				logger.Error(err)
				logger.Error(string(debug.Stack()[:]))

				var (
					errMsg string
					ok     bool
					e      error
				)

				if errMsg, ok = err.(string); !ok {
					if e, ok = err.(error); ok {
						errMsg = e.Error()
					}
				}

				if errMsg == "" {
					errMsg = "system error"
				}

				if ctx.WantJSON() {
					response.Error(ctx, errMsg)
					return
				}

				eng.errorPanelHTML(ctx, new(bytes.Buffer), errors2.New(errMsg))
			}
		}(ctx)
		ctx.Next()
	}
}

// wrapWithAuthMiddleware wrap a auth middleware to the given handler.
func (eng *Engine) wrapWithAuthMiddleware(handler context.Handler) context.Handlers {
	conn := db.GetConnection(eng.Services)
	return []context.Handler{eng.deferHandler(conn), response.OffLineHandler, auth.Middleware(conn), handler}
}

// wrapWithAuthMiddleware wrap a auth middleware to the given handler.
func (eng *Engine) wrap(handler context.Handler) context.Handlers {
	conn := db.GetConnection(eng.Services)
	return []context.Handler{eng.deferHandler(conn), response.OffLineHandler, handler}
}

// ============================
// HTML Content Render APIs
// ============================

// AddNavButtons add the nav buttons.
func (eng *Engine) AddNavButtons(title template2.HTML, icon string, action types.Action) *Engine {
	btn := types.GetNavButton(title, icon, action)
	*eng.NavButtons = append(*eng.NavButtons, btn)
	return eng
}

// Content call the Content method of engine adapter.
// If adapter is nil, it will panic.
func (eng *Engine) Content(ctx interface{}, panel types.GetPanelFn) {
	if eng.Adapter == nil {
		panic("adapter is nil")
	}
	eng.Adapter.Content(ctx, panel, eng.AdminPlugin().GetAddOperationFn(), *eng.NavButtons...)
}

// Content call the Content method of defaultAdapter.
// If defaultAdapter is nil, it will panic.
func Content(ctx interface{}, panel types.GetPanelFn) {
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
	defaultAdapter.Content(ctx, panel, engine.AdminPlugin().GetAddOperationFn(), *navButtons...)
}

// Data inject the route and corresponding handler to the web framework.
func (eng *Engine) Data(method, url string, handler context.Handler, noAuth ...bool) {
	if len(noAuth) > 0 && noAuth[0] {
		eng.Adapter.AddHandler(method, url, eng.wrap(handler))
	} else {
		eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(handler))
	}
}

// HTML inject the route and corresponding handler wrapped by the given function to the web framework.
func (eng *Engine) HTML(method, url string, fn types.GetPanelInfoFn, noAuth ...bool) {

	var handler = func(ctx *context.Context) {
		panel, err := fn(ctx)
		if err != nil {
			panel = template.WarningPanel(err.Error())
		}

		eng.AdminPlugin().GetAddOperationFn()(panel.Callbacks...)

		tmpl, tmplName := template.Default().GetTemplate(ctx.IsPjax())

		user := auth.Auth(ctx)

		buf := new(bytes.Buffer)
		hasError := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
			User:         user,
			Menu:         menu.GetGlobalMenu(user, eng.Adapter.GetConnection()).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
			Panel:        panel.GetContent(eng.config.IsProductionEnvironment()),
			Assets:       template.GetComponentAssetImportHTML(),
			Buttons:      eng.NavButtons.CheckPermission(user),
			TmplHeadHTML: template.Default().GetHeadHTML(),
			TmplFootJS:   template.Default().GetFootJS(),
		}))

		if hasError != nil {
			logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
		}

		ctx.HTMLByte(http.StatusOK, buf.Bytes())
	}

	if len(noAuth) > 0 && noAuth[0] {
		eng.Adapter.AddHandler(method, url, eng.wrap(handler))
	} else {
		eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(handler))
	}
}

// HTMLFile inject the route and corresponding handler which returns the panel content of given html file path
// to the web framework.
func (eng *Engine) HTMLFile(method, url, path string, data map[string]interface{}, noAuth ...bool) {

	var handler = func(ctx *context.Context) {

		cbuf := new(bytes.Buffer)

		t, err := template2.ParseFiles(path)
		if err != nil {
			eng.errorPanelHTML(ctx, cbuf, err)
			return
		} else {
			if err := t.Execute(cbuf, data); err != nil {
				eng.errorPanelHTML(ctx, cbuf, err)
				return
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
			Assets:       template.GetComponentAssetImportHTML(),
			Buttons:      eng.NavButtons.CheckPermission(user),
			TmplHeadHTML: template.Default().GetHeadHTML(),
			TmplFootJS:   template.Default().GetFootJS(),
		}))

		if hasError != nil {
			logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
		}

		ctx.HTMLByte(http.StatusOK, buf.Bytes())
	}

	if len(noAuth) > 0 && noAuth[0] {
		eng.Adapter.AddHandler(method, url, eng.wrap(handler))
	} else {
		eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(handler))
	}
}

// HTMLFiles inject the route and corresponding handler which returns the panel content of given html files path
// to the web framework.
func (eng *Engine) HTMLFiles(method, url string, data map[string]interface{}, files ...string) {
	eng.Adapter.AddHandler(method, url, eng.wrapWithAuthMiddleware(eng.htmlFilesHandler(data, files...)))
}

// HTMLFilesNoAuth inject the route and corresponding handler which returns the panel content of given html files path
// to the web framework without auth check.
func (eng *Engine) HTMLFilesNoAuth(method, url string, data map[string]interface{}, files ...string) {
	eng.Adapter.AddHandler(method, url, eng.wrap(eng.htmlFilesHandler(data, files...)))
}

// HTMLFiles inject the route and corresponding handler which returns the panel content of given html files path
// to the web framework.
func (eng *Engine) htmlFilesHandler(data map[string]interface{}, files ...string) context.Handler {
	return func(ctx *context.Context) {

		cbuf := new(bytes.Buffer)

		t, err := template2.ParseFiles(files...)
		if err != nil {
			eng.errorPanelHTML(ctx, cbuf, err)
			return
		} else {
			if err := t.Execute(cbuf, data); err != nil {
				eng.errorPanelHTML(ctx, cbuf, err)
				return
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
			Assets:       template.GetComponentAssetImportHTML(),
			Buttons:      eng.NavButtons.CheckPermission(user),
			TmplHeadHTML: template.Default().GetHeadHTML(),
			TmplFootJS:   template.Default().GetFootJS(),
		}))

		if hasError != nil {
			logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
		}

		ctx.HTMLByte(http.StatusOK, buf.Bytes())
	}
}

// errorPanelHTML add an error panel html to context response.
func (eng *Engine) errorPanelHTML(ctx *context.Context, buf *bytes.Buffer, err error) {

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Default().GetTemplate(ctx.IsPjax())

	hasError := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
		User:         user,
		Menu:         menu.GetGlobalMenu(user, eng.Adapter.GetConnection()).SetActiveClass(eng.config.URLRemovePrefix(ctx.Path())),
		Panel:        template.WarningPanel(err.Error()).GetContent(eng.config.IsProductionEnvironment()),
		Assets:       template.GetComponentAssetImportHTML(),
		Buttons:      (*eng.NavButtons).CheckPermission(user),
		TmplHeadHTML: template.Default().GetHeadHTML(),
		TmplFootJS:   template.Default().GetFootJS(),
	}))

	if hasError != nil {
		logger.Error(fmt.Sprintf("error: %s adapter content, ", eng.Adapter.Name()), hasError)
	}

	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

// ============================
// Admin Plugin APIs
// ============================

// AddGenerators add the admin generators.
func (eng *Engine) AddGenerators(list ...table.GeneratorList) *Engine {
	plug, exist := eng.FindPluginByName("admin")
	if exist {
		plug.(*admin.Admin).AddGenerators(list...)
		return eng
	}
	eng.PluginList = append(eng.PluginList, admin.NewAdmin(list...))
	return eng
}

// AdminPlugin get the admin plugin. if not exist, create one.
func (eng *Engine) AdminPlugin() *admin.Admin {
	plug, exist := eng.FindPluginByName("admin")
	if exist {
		return plug.(*admin.Admin)
	}
	adm := admin.NewAdmin()
	eng.PluginList = append(eng.PluginList, adm)
	return adm
}

// SetCaptcha set the captcha config.
func (eng *Engine) SetCaptcha(captcha map[string]string) *Engine {
	eng.AdminPlugin().SetCaptcha(captcha)
	return eng
}

// SetCaptchaDriver set the captcha config with driver.
func (eng *Engine) SetCaptchaDriver(driver string) *Engine {
	eng.AdminPlugin().SetCaptcha(map[string]string{"driver": driver})
	return eng
}

// AddGenerator add table model generator.
func (eng *Engine) AddGenerator(key string, g table.Generator) *Engine {
	eng.AdminPlugin().AddGenerator(key, g)
	return eng
}

// AddGlobalDisplayProcessFn call types.AddGlobalDisplayProcessFn.
func (eng *Engine) AddGlobalDisplayProcessFn(f types.FieldFilterFn) *Engine {
	types.AddGlobalDisplayProcessFn(f)
	return eng
}

// AddDisplayFilterLimit call types.AddDisplayFilterLimit.
func (eng *Engine) AddDisplayFilterLimit(limit int) *Engine {
	types.AddLimit(limit)
	return eng
}

// AddDisplayFilterTrimSpace call types.AddDisplayFilterTrimSpace.
func (eng *Engine) AddDisplayFilterTrimSpace() *Engine {
	types.AddTrimSpace()
	return eng
}

// AddDisplayFilterSubstr call types.AddDisplayFilterSubstr.
func (eng *Engine) AddDisplayFilterSubstr(start int, end int) *Engine {
	types.AddSubstr(start, end)
	return eng
}

// AddDisplayFilterToTitle call types.AddDisplayFilterToTitle.
func (eng *Engine) AddDisplayFilterToTitle() *Engine {
	types.AddToTitle()
	return eng
}

// AddDisplayFilterToUpper call types.AddDisplayFilterToUpper.
func (eng *Engine) AddDisplayFilterToUpper() *Engine {
	types.AddToUpper()
	return eng
}

// AddDisplayFilterToLower call types.AddDisplayFilterToLower.
func (eng *Engine) AddDisplayFilterToLower() *Engine {
	types.AddToUpper()
	return eng
}

// AddDisplayFilterXssFilter call types.AddDisplayFilterXssFilter.
func (eng *Engine) AddDisplayFilterXssFilter() *Engine {
	types.AddXssFilter()
	return eng
}

// AddDisplayFilterXssJsFilter call types.AddDisplayFilterXssJsFilter.
func (eng *Engine) AddDisplayFilterXssJsFilter() *Engine {
	types.AddXssJsFilter()
	return eng
}
