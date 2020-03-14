package table

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
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

func (g GeneratorList) Combine(gg GeneratorList) GeneratorList {
	for key, gen := range gg {
		if _, ok := g[key]; !ok {
			g[key] = gen
		}
	}
	return g
}

func (g GeneratorList) CombineAll(ggg []GeneratorList) GeneratorList {
	for _, gg := range ggg {
		for key, gen := range gg {
			if _, ok := g[key]; !ok {
				g[key] = gen
			}
		}
	}
	return g
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

	GetData(params parameter.Parameters) (PanelInfo, error)
	GetDataWithIds(params parameter.Parameters) (PanelInfo, error)
	GetDataWithId(params parameter.Parameters) (FormInfo, error)
	UpdateData(dataList form.Values) error
	InsertData(dataList form.Values) error
	DeleteData(id string) error

	GetNewForm() FormInfo

	Copy() Table
}

type BaseTable struct {
	Info       *types.InfoPanel
	Form       *types.FormPanel
	Detail     *types.InfoPanel
	CanAdd     bool
	Editable   bool
	Deletable  bool
	Exportable bool
	PrimaryKey PrimaryKey
}

func (base *BaseTable) GetInfo() *types.InfoPanel {
	return base.Info
}

func (base *BaseTable) GetDetail() *types.InfoPanel {
	return base.Detail
}

func (base *BaseTable) GetForm() *types.FormPanel {
	return base.Form
}

func (base *BaseTable) GetCanAdd() bool {
	return base.CanAdd && !base.Info.IsHideNewButton
}

func (base *BaseTable) GetPrimaryKey() PrimaryKey {
	return base.PrimaryKey
}

func (base *BaseTable) GetEditable() bool {
	return base.Editable && !base.Info.IsHideEditButton
}

func (base *BaseTable) GetDeletable() bool {
	return base.Deletable && !base.Info.IsHideDeleteButton
}

func (base *BaseTable) IsShowDetail() bool {
	return !base.Info.IsHideDetailButton
}

func (base *BaseTable) GetExportable() bool {
	return base.Exportable && !base.Info.IsHideExportButton
}

func (base *BaseTable) GetPaginator(size int, params parameter.Parameters, extraHtml ...template.HTML) types.PaginatorAttribute {

	var eh template.HTML

	if len(extraHtml) > 0 {
		eh = extraHtml[0]
	}

	return paginator.Get(paginator.Config{
		Size:         size,
		Param:        params,
		PageSizeList: base.Info.GetPageSizeList(),
	}).SetExtraInfo(eh)
}

type PanelInfo struct {
	Thead          types.Thead
	InfoList       types.InfoList
	FilterFormData types.FormFields
	Paginator      types.PaginatorAttribute
	Title          string
	Description    string
}

type FormInfo struct {
	FieldList         types.FormFields
	GroupFieldList    types.GroupFormFields
	GroupFieldHeaders types.GroupFieldHeaders
	Title             string
	Description       string
}

type PrimaryKey struct {
	Type db.DatabaseType
	Name string
}

const (
	DefaultPrimaryKeyName = "id"
	DefaultConnectionName = "default"
)

var (
	services service.List
	count    uint32
	lock     sync.Mutex
)

func SetServices(srv service.List) {
	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&count) != 0 {
		panic("can not initialize twice")
	}

	services = srv
}
