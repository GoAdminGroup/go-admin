package table

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"sync"
	"sync/atomic"
)

type Generator func(ctx *context.Context) Table

type GeneratorList map[string]Generator

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
	GetDetailFromInfo() *types.InfoPanel
	GetForm() *types.FormPanel

	GetCanAdd() bool
	GetEditable() bool
	GetDeletable() bool
	GetExportable() bool

	GetPrimaryKey() PrimaryKey

	GetData(params parameter.Parameters) (PanelInfo, error)
	GetDataWithIds(params parameter.Parameters) (PanelInfo, error)
	GetDataWithId(params parameter.Parameters) (FormInfo, error)
	UpdateData(dataList form.Values) error
	InsertData(dataList form.Values) error
	DeleteData(pk string) error

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
	return base.Info.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetDetail() *types.InfoPanel {
	return base.Detail.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetDetailFromInfo() *types.InfoPanel {
	base.Detail.FieldList = base.GetInfo().FieldList
	return base.GetDetail()
}

func (base *BaseTable) GetForm() *types.FormPanel {
	return base.Form.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetCanAdd() bool {
	return base.CanAdd
}

func (base *BaseTable) GetPrimaryKey() PrimaryKey {
	return base.PrimaryKey
}

func (base *BaseTable) GetEditable() bool {
	return base.Editable
}

func (base *BaseTable) GetDeletable() bool {
	return base.Deletable
}

func (base *BaseTable) GetExportable() bool {
	return base.Exportable
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
