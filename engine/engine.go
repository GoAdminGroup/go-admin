// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package engine

import (
	"github.com/GoAdminGroup/go-admin/adapter"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/types"
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

	config config.Config
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

	return eng.Adapter.Use(router, eng.PluginList)
}

// AddPlugins add the plugins and initialize them.
func (eng *Engine) AddPlugins(plugs ...plugins.Plugin) *Engine {

	for _, plug := range plugs {
		plug.InitPlugin(eng.Services)
	}

	eng.PluginList = append(eng.PluginList, plugs...)
	return eng
}

// AddConfig set the global config.
func (eng *Engine) AddConfig(cfg config.Config) *Engine {
	return eng.setConfig(cfg).InitDatabase()
}

// setConfig set the config of engine.
func (eng *Engine) setConfig(cfg config.Config) *Engine {
	eng.config = config.Set(cfg)
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

// Register set default adapter of engine.
func Register(ada adapter.WebFrameWork) {
	if ada == nil {
		panic("adapter is nil")
	}
	defaultAdapter = ada
}

// Content call the Content method of engine adapter.
// If adapter is nil, it will panic.
func (eng *Engine) Content(ctx interface{}, panel types.GetPanelFn) {
	if eng.Adapter == nil {
		panic("adapter is nil")
	}
	eng.Adapter.Content(ctx, panel)
}

// Content call the Content method of defaultAdapter.
// If defaultAdapter is nil, it will panic.
func Content(ctx interface{}, panel types.GetPanelFn) {
	if defaultAdapter == nil {
		panic("adapter is nil")
	}
	defaultAdapter.Content(ctx, panel)
}

// User call the User method of defaultAdapter.
func User(ci interface{}) (models.UserModel, bool) {
	return defaultAdapter.User(ci)
}

// User call the User method of engine adapter.
func (eng *Engine) User(ci interface{}) (models.UserModel, bool) {
	return eng.Adapter.User(ci)
}

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
