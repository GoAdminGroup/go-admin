package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/controller"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Admin is a GoAdmin plugin.
type Admin struct {
	app      *context.App
	tableCfg table.GeneratorList
}

// InitPlugin implements Plugin.InitPlugin.
func (admin *Admin) InitPlugin(services service.List) {

	cfg := config.Get()

	// Init router
	App.app = InitRouter(cfg.Prefix(), services)

	table.SetGenerators(table.GeneratorList{
		"manager":        table.GetManagerTable,
		"permission":     table.GetPermissionTable,
		"roles":          table.GetRolesTable,
		"op":             table.GetOpTable,
		"menu":           table.GetMenuTable,
		"normal_manager": table.GetNormalManagerTable,
	})
	table.SetServices(services)
	table.SetGenerators(admin.tableCfg)
	table.InitTableList()

	controller.SetConfig(cfg)
	controller.SetServices(services)
}

// App is the global Admin plugin.
var App = &Admin{
	tableCfg: make(table.GeneratorList),
}

// NewAdmin return the global Admin plugin.
func NewAdmin(tableCfg ...table.GeneratorList) *Admin {
	App.tableCfg.CombineAll(tableCfg)
	return App
}

// SetCaptcha set captcha driver.
func (admin *Admin) SetCaptcha(captcha map[string]string) *Admin {
	controller.SetCaptcha(captcha)
	return admin
}

// AddGenerator add table model generator.
func (admin *Admin) AddGenerator(key string, g table.Generator) *Admin {
	admin.tableCfg.Add(key, g)
	return admin
}

// GetRequest implements Plugin.GetRequest.
func (admin *Admin) GetRequest() []context.Path {
	return admin.app.Requests
}

// GetHandler implements Plugin.GetHandler.
func (admin *Admin) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, admin.app)
}

// AddGlobalDisplayProcessFn call types.AddGlobalDisplayProcessFn
func (admin *Admin) AddGlobalDisplayProcessFn(f types.DisplayProcessFn) *Admin {
	types.AddGlobalDisplayProcessFn(f)
	return admin
}

// AddDisplayFilterLimit call types.AddDisplayFilterLimit
func (admin *Admin) AddDisplayFilterLimit(limit int) *Admin {
	types.AddLimit(limit)
	return admin
}

// AddDisplayFilterTrimSpace call types.AddDisplayFilterTrimSpace
func (admin *Admin) AddDisplayFilterTrimSpace() *Admin {
	types.AddTrimSpace()
	return admin
}

// AddDisplayFilterSubstr call types.AddDisplayFilterSubstr
func (admin *Admin) AddDisplayFilterSubstr(start int, end int) *Admin {
	types.AddSubstr(start, end)
	return admin
}

// AddDisplayFilterToTitle call types.AddDisplayFilterToTitle
func (admin *Admin) AddDisplayFilterToTitle() *Admin {
	types.AddToTitle()
	return admin
}

// AddDisplayFilterToUpper call types.AddDisplayFilterToUpper
func (admin *Admin) AddDisplayFilterToUpper() *Admin {
	types.AddToUpper()
	return admin
}

// AddDisplayFilterToLower call types.AddDisplayFilterToLower
func (admin *Admin) AddDisplayFilterToLower() *Admin {
	types.AddToUpper()
	return admin
}

// AddDisplayFilterXssFilter call types.AddDisplayFilterXssFilter
func (admin *Admin) AddDisplayFilterXssFilter() *Admin {
	types.AddXssFilter()
	return admin
}

// AddDisplayFilterXssJsFilter call types.AddDisplayFilterXssJsFilter
func (admin *Admin) AddDisplayFilterXssJsFilter() *Admin {
	types.AddXssJsFilter()
	return admin
}
