package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
)

type Admin struct {
	app      *context.App
	tableCfg map[string]table.Generator
}

func (admin *Admin) InitPlugin() {

	cfg := config.Get()

	// Init database
	for driver, databaseCfg := range cfg.DATABASE.GroupByDriver() {
		db.GetConnectionByDriver(driver).InitDB(databaseCfg)
	}

	// Init router
	App.app = InitRouter(cfg.Prefix())

	table.SetGenerators(map[string]table.Generator{
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

var App = new(Admin)

func NewAdmin(tableCfg map[string]table.Generator) *Admin {
	App.tableCfg = tableCfg
	return App
}

func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

func (admin *Admin) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, admin.app)
}
