package admin

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin/modules/language"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/modules/connections"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
)

type Admin struct {
	app      *context.App
	config   config.Config
	tableCfg map[string]models.GetTableDataFunc
}

func (admin *Admin) InitPlugin(cfg config.Config) {

	connections.GetConnection().InitDB(map[string]config.Database{
		"default": cfg.DATABASE,
	})

	AdminApp.config = cfg
	AdminApp.app = InitRouter("/" + cfg.ADMIN_PREFIX)

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

	cfg.ADMIN_PREFIX = "/" + cfg.ADMIN_PREFIX
	if cfg.THEME == "" {
		cfg.THEME = "adminlte"
	}
	controller.SetConfig(cfg)

	menu.InitMenu()
}

var AdminApp = new(Admin)

func NewAdmin(tableCfg map[string]models.GetTableDataFunc) *Admin {
	AdminApp.tableCfg = tableCfg
	return AdminApp
}

func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

func (admin *Admin) GetHandler(url, method string) context.Handler {
	return plugins.GetHandler(url, method, &admin.app.HandlerList)
}

func (admin *Admin) GetLocales() map[string]string {
	return language.Locales[admin.config.LANGUAGE]
}
