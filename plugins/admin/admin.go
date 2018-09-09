package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/modules/connections"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
)

type Admin struct {
	app      *context.App
	tableCfg map[string]models.GetTableDataFunc
}

func (admin *Admin) InitPlugin() {

	cfg := config.Get()

	for _, databaseCfg := range cfg.DATABASE {
		connections.GetConnectionByDriver(databaseCfg.DRIVER).InitDB(map[string]config.Database{
			"default": databaseCfg,
		})
	}

	App.app = InitRouter("/" + cfg.PREFIX)

	models.SetTableFuncConfig(map[string]models.GetTableDataFunc{
		// 管理员管理部分
		"manager":    models.GetManagerTable,    // 管理员管理
		"permission": models.GetPermissionTable, // 权限管理
		"roles":      models.GetRolesTable,      // 角色管理
		"op":         models.GetOpTable,         // 操作日志管理
		"menu":       models.GetMenuTable,       // 菜单管理
	})
	models.SetTableFuncConfig(admin.tableCfg)
	models.InitGlobalTableList()

	cfg.PREFIX = "/" + cfg.PREFIX
	if cfg.THEME == "" {
		cfg.THEME = "adminlte"
	}
	controller.SetConfig(cfg)

	menu.InitMenu()
}

var App = new(Admin)

func NewAdmin(tableCfg map[string]models.GetTableDataFunc) *Admin {
	App.tableCfg = tableCfg
	return App
}

func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

func (admin *Admin) GetHandler(url, method string) context.Handler {
	return plugins.GetHandler(url, method, &admin.app.HandlerList)
}