package admin

import (
	"goAdmin/context"
	"goAdmin/modules/config"
	"goAdmin/plugins/admin/modules/language"
	"goAdmin/plugins/admin/models"
	"goAdmin/modules/connections/mysql"
	"strconv"
	"goAdmin/modules/menu"
	"goAdmin/plugins"
)

type Admin struct {
	app      *context.App
	config   config.Config
	tableCfg map[string]models.GetTableDataFunc
}

func (admin *Admin) InitPlugin(cfg config.Config) {
	idleCon, _ := strconv.Atoi(cfg.DATABASE_MAX_IDLE_CON)
	openCon, _ := strconv.Atoi(cfg.DATABASE_MAX_OPEN_CON)

	mysql.InitDB(cfg.DATABASE_USER, cfg.DATABASE_PWD, cfg.DATABASE_PORT,
		cfg.DATABASE_IP, cfg.DATABASE_NAME, idleCon, openCon)

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
