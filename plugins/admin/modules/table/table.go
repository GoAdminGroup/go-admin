package table

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
	"net/http"
	"net/url"
)

type Generator func(ctx *context.Context) Table

type GeneratorList map[string]Generator

func (g GeneratorList) InjectRoutes(app *context.App, srv service.List) {
	authHandler := auth.Middleware(db.GetConnection(srv))
	for _, gen := range g {
		table := gen(context.NewContext(&http.Request{
			URL: &url.URL{},
		}))
		for _, cb := range table.GetInfo().Callbacks {
			if cb.Value[constant.ContextNodeNeedAuth] == 1 {
				app.AppendReqAndResp(cb.Path, cb.Method, append([]context.Handler{authHandler}, cb.Handlers...))
			} else {
				app.AppendReqAndResp(cb.Path, cb.Method, cb.Handlers)
			}
		}
		for _, cb := range table.GetForm().Callbacks {
			if cb.Value[constant.ContextNodeNeedAuth] == 1 {
				app.AppendReqAndResp(cb.Path, cb.Method, append([]context.Handler{authHandler}, cb.Handlers...))
			} else {
				app.AppendReqAndResp(cb.Path, cb.Method, cb.Handlers)
			}
		}
	}
}

func (g GeneratorList) Add(key string, gen Generator) {
	g[key] = gen
}

func (g GeneratorList) Combine(gg GeneratorList) {
	for key, gen := range gg {
		if _, ok := g[key]; !ok {
			g[key] = gen
		}
	}
}

func (g GeneratorList) CombineAll(ggg []GeneratorList) {
	for _, gg := range ggg {
		for key, gen := range gg {
			if _, ok := g[key]; !ok {
				g[key] = gen
			}
		}
	}
}

var generators = make(GeneratorList)

func Get(key string, ctx *context.Context) Table {
	return generators[key](ctx)
}

// SetGenerators update generators.
func SetGenerators(gens map[string]Generator) {
	for key, gen := range gens {
		generators[key] = gen
	}
}

type Table interface {
	GetInfo() *types.InfoPanel
	GetDetail() *types.InfoPanel
	GetForm() *types.FormPanel

	GetCanAdd() bool
	GetEditable() bool
	GetDeletable() bool
	GetExportable() bool
	IsShowDetail() bool

	GetPrimaryKey() PrimaryKey

	GetData(path string, params parameter.Parameters, isAll bool) (PanelInfo, error)
	GetDataWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error)
	GetDataWithId(id string) ([]types.FormField, [][]types.FormField, []string, string, string, error)
	UpdateDataFromDatabase(dataList form.Values) error
	InsertDataFromDatabase(dataList form.Values) error
	DeleteDataFromDatabase(id string) error

	Copy() Table
}

type PrimaryKey struct {
	Type db.DatabaseType
	Name string
}

const (
	DefaultPrimaryKeyName = "id"
	DefaultConnectionName = "default"
)

var services service.List

func SetServices(srv service.List) {
	services = srv
}
