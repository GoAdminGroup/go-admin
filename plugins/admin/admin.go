package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Admin struct {
	app      *context.App
	tableCfg table.GeneratorList
}

func (admin *Admin) InitPlugin() {

	cfg := config.Get()

	// Init database
	for driver, databaseCfg := range cfg.Databases.GroupByDriver() {
		db.GetConnectionByDriver(driver).InitDB(databaseCfg)
	}

	// Init router
	App.app = InitRouter(cfg.Prefix())

	table.SetGenerators(table.GeneratorList{
		"manager":    table.GetManagerTable,
		"permission": table.GetPermissionTable,
		"roles":      table.GetRolesTable,
		"op":         table.GetOpTable,
		"menu":       table.GetMenuTable,
	})
	table.SetGenerators(admin.tableCfg)
	table.InitTableList()

	controller.SetConfig(cfg)
}

var App = &Admin{
	tableCfg: make(table.GeneratorList),
}

func NewAdmin(tableCfg table.GeneratorList) *Admin {
	App.tableCfg = tableCfg
	return App
}

func (admin *Admin) AddGenerator(key string, g table.Generator) *Admin {
	admin.tableCfg.Add(key, g)
	return admin
}

func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

func (admin *Admin) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, admin.app)
}

func (admin *Admin) AddGlobalDisplayProcessFn(f types.DisplayProcessFn) *Admin {
	types.AddGlobalDisplayProcessFn(f)
	return admin
}

func (admin *Admin) AddDisplayFilterLimit(limit int) *Admin {
	types.AddLimit(limit)
	return admin
}

func (admin *Admin) AddDisplayFilterTrimSpace() *Admin {
	types.AddTrimSpace()
	return admin
}

func (admin *Admin) AddDisplayFilterSubstr(start int, end int) *Admin {
	types.AddSubstr(start, end)
	return admin
}

func (admin *Admin) AddDisplayFilterToTitle() *Admin {
	types.AddToTitle()
	return admin
}

func (admin *Admin) AddDisplayFilterToUpper() *Admin {
	types.AddToUpper()
	return admin
}

func (admin *Admin) AddDisplayFilterToLower() *Admin {
	types.AddToUpper()
	return admin
}

func (admin *Admin) AddDisplayFilterXssFilter() *Admin {
	types.AddXssFilter()
	return admin
}

func (admin *Admin) AddDisplayFilterXssJsFilter() *Admin {
	types.AddXssJsFilter()
	return admin
}
