package table

import (
	"html/template"
	"sync"
	"sync/atomic"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Generator func(ctx *context.Context) Table

type GeneratorList map[string]Generator

func (g GeneratorList) Add(key string, gen Generator) {
	g[key] = gen
}

func (g GeneratorList) Combine(list GeneratorList) GeneratorList {
	for key, gen := range list {
		if _, ok := g[key]; !ok {
			g[key] = gen
		}
	}
	return g
}

func (g GeneratorList) CombineAll(gens []GeneratorList) GeneratorList {
	for _, list := range gens {
		for key, gen := range list {
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
	GetNewForm() *types.FormPanel
	GetActualNewForm() *types.FormPanel

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

	GetNewFormInfo() FormInfo

	GetOnlyInfo() bool
	GetOnlyDetail() bool
	GetOnlyNewForm() bool
	GetOnlyUpdateForm() bool

	Copy() Table
}

type BaseTable struct {
	Info           *types.InfoPanel
	Form           *types.FormPanel
	NewForm        *types.FormPanel
	Detail         *types.InfoPanel
	CanAdd         bool
	Editable       bool
	Deletable      bool
	Exportable     bool
	OnlyInfo       bool
	OnlyDetail     bool
	OnlyNewForm    bool
	OnlyUpdateForm bool
	PrimaryKey     PrimaryKey
}

func (base *BaseTable) GetInfo() *types.InfoPanel {
	return base.Info.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetDetail() *types.InfoPanel {
	return base.Detail.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetDetailFromInfo() *types.InfoPanel {
	detail := base.GetDetail()
	detail.FieldList = make(types.FieldList, len(base.Info.FieldList))
	copy(detail.FieldList, base.Info.FieldList)
	return detail
}

func (base *BaseTable) GetForm() *types.FormPanel {
	return base.Form.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetNewForm() *types.FormPanel {
	return base.NewForm.SetPrimaryKey(base.PrimaryKey.Name, base.PrimaryKey.Type)
}

func (base *BaseTable) GetActualNewForm() *types.FormPanel {
	if len(base.NewForm.FieldList) == 0 {
		return base.Form
	}
	return base.NewForm
}

func (base *BaseTable) GetCanAdd() bool {
	return base.CanAdd
}

func (base *BaseTable) GetPrimaryKey() PrimaryKey { return base.PrimaryKey }
func (base *BaseTable) GetEditable() bool         { return base.Editable }
func (base *BaseTable) GetDeletable() bool        { return base.Deletable }
func (base *BaseTable) GetExportable() bool       { return base.Exportable }
func (base *BaseTable) GetOnlyInfo() bool         { return base.OnlyInfo }
func (base *BaseTable) GetOnlyDetail() bool       { return base.OnlyDetail }
func (base *BaseTable) GetOnlyNewForm() bool      { return base.OnlyNewForm }
func (base *BaseTable) GetOnlyUpdateForm() bool   { return base.OnlyUpdateForm }

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
	Thead          types.Thead              `json:"thead"`
	InfoList       types.InfoList           `json:"info_list"`
	FilterFormData types.FormFields         `json:"filter_form_data"`
	Paginator      types.PaginatorAttribute `json:"-"`
	Title          string                   `json:"title"`
	Description    string                   `json:"description"`
}

type FormInfo struct {
	FieldList         types.FormFields        `json:"field_list"`
	GroupFieldList    types.GroupFormFields   `json:"group_field_list"`
	GroupFieldHeaders types.GroupFieldHeaders `json:"group_field_headers"`
	Title             string                  `json:"title"`
	Description       string                  `json:"description"`
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
