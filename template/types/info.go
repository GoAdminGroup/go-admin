package types

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/GoAdminGroup/go-admin/template/types/table"
	"html"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// FieldModel is the single query result.
type FieldModel struct {
	// The primaryKey of the table.
	ID string

	// The value of the single query result.
	Value string

	// The current row data.
	Row map[string]interface{}
}

// PostFieldModel contains ID and value of the single query result and the current row data.
type PostFieldModel struct {
	ID    string
	Value FieldModelValue
	Row   map[string]interface{}
}

type Callbacks []context.Node

func (c Callbacks) AddCallback(node context.Node) Callbacks {
	if node.Path != "" && node.Method != "" && len(node.Handlers) > 0 {
		return append(c, node)
	}
	return c
}

type FieldModelValue []string

func (r FieldModelValue) Value() string {
	return r.First()
}

func (r FieldModelValue) First() string {
	return r[0]
}

// FieldDisplay is filter function of data.
type FieldFilterFn func(value FieldModel) interface{}

// PostFieldFilterFn is filter function of data.
type PostFieldFilterFn func(value PostFieldModel) interface{}

// Field is the table field.
type Field struct {
	Head     string
	Field    string
	TypeName db.DatabaseType

	Join Join

	Width      int
	Sortable   bool
	EditAble   bool
	Fixed      bool
	Filterable bool
	Hide       bool

	EditType    table.Type
	EditOptions []map[string]string

	FilterType      form.Type
	FilterOptions   FieldOptions
	FilterOperator  FilterOperator
	FilterOptionExt template.JS
	FilterHead      string
	FilterHelpMsg   template.HTML
	FilterProcess   func(string) string

	FieldDisplay
}

func (f Field) FilterFormField(params parameter.Parameters, headField string) []FormField {
	var value, value2 string

	if f.FilterType.IsRange() {
		value = params.GetFilterFieldValueStart(headField)
		value2 = params.GetFilterFieldValueEnd(headField)
	} else {
		if f.FilterOperator == FilterOperatorFree {
			value2 = GetOperatorFromValue(params.GetFieldOperator(headField)).String()
		}
		value = params.GetFieldValue(headField)
	}

	var filterForm = make([]FormField, 0)

	filterForm = append(filterForm, FormField{
		Field:     headField,
		Head:      modules.AorB(f.FilterHead == "", f.Head, f.FilterHead),
		TypeName:  f.TypeName,
		HelpMsg:   f.FilterHelpMsg,
		FormType:  f.FilterType,
		Editable:  true,
		Value:     template.HTML(value),
		Value2:    value2,
		Options:   f.FilterOptions.SetSelected(params.GetFieldValue(f.Field), f.FilterType.SelectedLabel()),
		OptionExt: f.FilterOptionExt,
		Label:     f.FilterOperator.Label(),
	})

	if f.FilterOperator.AddOrNot() {
		filterForm = append(filterForm, FormField{
			Field:    headField + parameter.FilterParamOperatorSuffix,
			Head:     f.Head,
			TypeName: f.TypeName,
			Value:    template.HTML(f.FilterOperator.Value()),
			FormType: f.FilterType,
			Hide:     true,
		})
	}

	return filterForm
}

func (f Field) GetEditOptions() string {
	if len(f.EditOptions) == 0 {
		return ""
	}
	eo, err := json.Marshal(f.EditOptions)

	if err != nil {
		return ""
	}

	return string(eo)
}

func (f Field) Exist() bool {
	return f.Field != ""
}

type FieldList []Field

type TableInfo struct {
	Table      string
	PrimaryKey string
	Delimiter  string
	Driver     string
}

func (f FieldList) GetTheadAndFilterForm(info TableInfo, params parameter.Parameters, columns []string) ([]map[string]string,
	string, string, []string, []FormField) {
	var (
		thead      = make([]map[string]string, 0)
		fields     = ""
		joins      = ""
		joinTables = make([]string, 0)
		filterForm = make([]FormField, 0)
	)
	for _, field := range f {
		if field.Field != info.PrimaryKey && modules.InArray(columns, field.Field) &&
			!field.Join.Valid() {
			fields += info.Table + "." + modules.FilterField(field.Field, info.Delimiter) + ","
		}

		headField := field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + "_goadmin_join_" + field.Field
			fields += db.GetAggregationExpression(info.Driver, field.Join.Table+"."+
				modules.FilterField(field.Field, info.Delimiter), headField, JoinFieldValueDelimiter) + ","
			if !modules.InArray(joinTables, field.Join.Table) {
				joinTables = append(joinTables, field.Join.Table)
				joins += " left join " + modules.FilterField(field.Join.Table, info.Delimiter) + " on " +
					field.Join.Table + "." + modules.FilterField(field.Join.JoinField, info.Delimiter) + " = " +
					info.Table + "." + modules.FilterField(field.Join.Field, info.Delimiter)
			}
		}

		if field.Filterable {
			filterForm = append(filterForm, field.FilterFormField(params, headField)...)
		}

		if field.Hide {
			continue
		}
		thead = append(thead, map[string]string{
			"head":       field.Head,
			"sortable":   modules.AorB(field.Sortable, "1", "0"),
			"field":      headField,
			"hide":       modules.AorB(modules.InArrayWithoutEmpty(params.Columns, headField), "0", "1"),
			"editable":   modules.AorB(field.EditAble, "true", "false"),
			"edittype":   field.EditType.String(),
			"editoption": field.GetEditOptions(),
			"width":      strconv.Itoa(field.Width),
		})
	}

	return thead, fields, joins, joinTables, filterForm
}

func (f FieldList) GetThead(info TableInfo, params parameter.Parameters, columns []string) ([]map[string]string, string, string) {
	var (
		thead      = make([]map[string]string, 0)
		fields     = ""
		joins      = ""
		joinTables = make([]string, 0)
	)
	for _, field := range f {
		if field.Field != info.PrimaryKey && modules.InArray(columns, field.Field) &&
			!field.Join.Valid() {
			fields += info.Table + "." + modules.FilterField(field.Field, info.Delimiter) + ","
		}

		headField := field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + "_goadmin_join_" + field.Field
			fields += db.GetAggregationExpression(info.Driver, field.Join.Table+"."+
				modules.FilterField(field.Field, info.Delimiter), headField, JoinFieldValueDelimiter) + ","
			if !modules.InArray(joinTables, field.Join.Table) {
				joinTables = append(joinTables, field.Join.Table)
				joins += " left join " + modules.FilterField(field.Join.Table, info.Delimiter) + " on " +
					field.Join.Table + "." + modules.FilterField(field.Join.JoinField, info.Delimiter) + " = " +
					info.Table + "." + modules.FilterField(field.Join.Field, info.Delimiter)
			}
		}

		if field.Hide {
			continue
		}
		thead = append(thead, map[string]string{
			"head":       field.Head,
			"sortable":   modules.AorB(field.Sortable, "1", "0"),
			"field":      headField,
			"hide":       modules.AorB(modules.InArrayWithoutEmpty(params.Columns, headField), "0", "1"),
			"editable":   modules.AorB(field.EditAble, "true", "false"),
			"edittype":   field.EditType.String(),
			"editoption": field.GetEditOptions(),
			"width":      strconv.Itoa(field.Width),
		})
	}

	return thead, fields, joins
}

func (f FieldList) GetFieldFilterProcessValue(key string, value string) string {
	field := f.GetFieldByFieldName(key)
	if field.FilterProcess != nil {
		value = field.FilterProcess(value)
	}
	return value
}

func (f FieldList) GetFieldJoinTable(key string) string {
	field := f.GetFieldByFieldName(key)
	if field.Exist() {
		return field.Join.Table
	}
	return ""
}

func (f FieldList) GetFieldByFieldName(name string) Field {
	for _, field := range f {
		if field.Field == name {
			return field
		}
	}
	return Field{}
}

type Join struct {
	Table     string
	Field     string
	JoinField string
}

func (j Join) Valid() bool {
	return j.Table != "" && j.Field != "" && j.JoinField != ""
}

var JoinFieldValueDelimiter = utils.Uuid(8)

type TabGroups [][]string

func (t TabGroups) Valid() bool {
	return len(t) > 0
}

func NewTabGroups(items ...string) TabGroups {
	var t = make(TabGroups, 0)
	return append(t, items)
}

func (t TabGroups) AddGroup(items ...string) TabGroups {
	return append(t, items)
}

type TabHeaders []string

func (t TabHeaders) Add(header string) TabHeaders {
	return append(t, header)
}

type GetDataFn func(param parameter.Parameters) ([]map[string]interface{}, int)

type DeleteFn func(ids []string) error

type Sort uint8

const (
	SortDesc Sort = iota
	SortAsc
)

// InfoPanel
type InfoPanel struct {
	FieldList         FieldList
	curFieldListIndex int

	Table       string
	Title       string
	Description string

	// Warn: may be deprecated future.
	TabGroups  TabGroups
	TabHeaders TabHeaders

	Sort      Sort
	SortField string

	PageSizeList    []int
	DefaultPageSize int

	IsHideNewButton    bool
	IsHideExportButton bool
	IsHideEditButton   bool
	IsHideDeleteButton bool
	IsHideDetailButton bool
	IsHideFilterButton bool
	IsHideRowSelector  bool
	IsHidePagination   bool
	IsHideFilterArea   bool
	FilterFormLayout   form.Layout

	Wheres    Wheres
	WhereRaws WhereRaw

	Callbacks Callbacks

	Buttons Buttons

	TableLayout string

	DeleteHook  DeleteFn
	PreDeleteFn DeleteFn
	DeleteFn    DeleteFn

	GetDataFn GetDataFn

	processChains DisplayProcessFnChains

	ActionButtons Buttons
	Action        template.HTML
	HeaderHtml    template.HTML
	FooterHtml    template.HTML
}

type Where struct {
	Join     string
	Field    string
	Operator string
	Arg      interface{}
}

type Wheres []Where

func (whs Wheres) Statement(wheres, delimiter string, whereArgs []interface{}, existKeys, columns []string) (string, []interface{}) {
	for k, wh := range whs {

		whFieldArr := strings.Split(wh.Field, ".")
		whField := ""
		whTable := ""
		if len(whFieldArr) > 1 {
			whField = whFieldArr[1]
			whTable = whFieldArr[0]
		} else {
			whField = whFieldArr[0]
		}

		if modules.InArray(existKeys, whField) {
			continue
		}

		// TODO: support like operation and join table
		if modules.InArray(columns, whField) {

			joinMark := ""
			if k != len(whs)-1 {
				joinMark = whs[k+1].Join
			}

			if whTable != "" {
				wheres += whTable + "." + modules.FilterField(whField, delimiter) + " " + wh.Operator + " ? " + joinMark + " "
			} else {
				wheres += modules.FilterField(whField, delimiter) + " " + wh.Operator + " ? " + joinMark + " "
			}
			whereArgs = append(whereArgs, wh.Arg)
		}
	}
	return wheres, whereArgs
}

type WhereRaw struct {
	Raw  string
	Args []interface{}
}

func (wh WhereRaw) check() int {
	index := 0
	for i := 0; i < len(wh.Raw); i++ {
		if wh.Raw[i] == ' ' {
			continue
		} else {
			if wh.Raw[i] == 'a' {
				if len(wh.Raw) < i+3 {
					break
				} else {
					if wh.Raw[i+1] == 'n' && wh.Raw[i+2] == 'd' {
						index = i + 3
					}
				}
			} else if wh.Raw[i] == 'o' {
				if len(wh.Raw) < i+2 {
					break
				} else {
					if wh.Raw[i+1] == 'r' {
						index = i + 2
					}
				}
			} else {
				break
			}
		}
	}
	return index
}

func (wh WhereRaw) Statement(wheres string, whereArgs []interface{}) (string, []interface{}) {

	if wh.Raw == "" {
		return wheres, whereArgs
	}

	if wheres != "" {
		if wh.check() != 0 {
			wheres += wh.Raw + " "
		} else {
			wheres += " and " + wh.Raw + " "
		}

		whereArgs = append(whereArgs, wh.Args...)
	} else {
		wheres += wh.Raw[wh.check():] + " "
		whereArgs = append(whereArgs, wh.Args...)
	}

	return wheres, whereArgs
}

type Handler func(ctx *context.Context) (success bool, msg string, data interface{})

func (h Handler) Wrap() context.Handler {
	return func(ctx *context.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"code": 500,
					"data": "",
					"msg":  "error",
				})
			}
		}()

		code := 0
		s, m, d := h(ctx)

		if !s {
			code = 500
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": code,
			"data": d,
			"msg":  m,
		})
	}
}

type Action interface {
	Js() template.JS
	BtnAttribute() template.HTML
	BtnClass() template.HTML
	ExtContent() template.HTML
	SetBtnId(btnId string)
	GetCallbacks() context.Node
}

type DefaultAction struct {
	Attr template.HTML
	JS   template.JS
	Ext  template.HTML
}

func NewDefaultAction(attr, ext template.HTML, js template.JS) *DefaultAction {
	return &DefaultAction{Attr: attr, Ext: ext, JS: js}
}

func (def *DefaultAction) SetBtnId(btnId string)       {}
func (def *DefaultAction) Js() template.JS             { return def.JS }
func (def *DefaultAction) BtnAttribute() template.HTML { return def.Attr }
func (def *DefaultAction) BtnClass() template.HTML     { return "" }
func (def *DefaultAction) ExtContent() template.HTML   { return def.Ext }
func (def *DefaultAction) GetCallbacks() context.Node  { return context.Node{} }

var _ Action = (*DefaultAction)(nil)

var DefaultPageSizeList = []int{10, 20, 30, 50, 100}

const DefaultPageSize = 10

func NewInfoPanel(pk string) *InfoPanel {
	return &InfoPanel{
		curFieldListIndex: -1,
		PageSizeList:      DefaultPageSizeList,
		DefaultPageSize:   DefaultPageSize,
		processChains:     make(DisplayProcessFnChains, 0),
		Buttons:           make(Buttons, 0),
		Callbacks:         make(Callbacks, 0),
		Wheres:            make([]Where, 0),
		WhereRaws:         WhereRaw{},
		SortField:         pk,
		TableLayout:       "auto",
	}
}

func (i *InfoPanel) Where(field string, operator string, arg interface{}) *InfoPanel {
	i.Wheres = append(i.Wheres, Where{Field: field, Operator: operator, Arg: arg, Join: "and"})
	return i
}

func (i *InfoPanel) WhereOr(field string, operator string, arg interface{}) *InfoPanel {
	i.Wheres = append(i.Wheres, Where{Field: field, Operator: operator, Arg: arg, Join: "or"})
	return i
}

func (i *InfoPanel) WhereRaw(raw string, arg ...interface{}) *InfoPanel {
	i.WhereRaws.Raw = raw
	i.WhereRaws.Args = arg
	return i
}

func (i *InfoPanel) AddButton(title template.HTML, icon string, action Action, color ...template.HTML) *InfoPanel {
	id := "info-btn-" + utils.Uuid(10)
	action.SetBtnId(id)
	if len(color) == 0 {
		i.Buttons = append(i.Buttons, DefaultButton{Title: title, Id: id, Action: action, Icon: icon})
	}
	if len(color) == 1 {
		i.Buttons = append(i.Buttons, DefaultButton{Title: title, Color: color[0], Id: id, Action: action, Icon: icon})
	}
	if len(color) >= 2 {
		i.Buttons = append(i.Buttons, DefaultButton{Title: title, Color: color[0], TextColor: color[1], Id: id, Action: action, Icon: icon})
	}
	i.Callbacks = i.Callbacks.AddCallback(action.GetCallbacks())
	return i
}

func (i *InfoPanel) AddActionButton(title template.HTML, action Action, ids ...string) *InfoPanel {
	id := ""
	if len(ids) > 0 {
		id = ids[0]
	} else {
		id = "action-info-btn-" + utils.Uuid(10)
	}
	action.SetBtnId(id)
	i.ActionButtons = append(i.ActionButtons, ActionButton{Title: title, Id: id, Action: action})
	i.Callbacks = i.Callbacks.AddCallback(action.GetCallbacks())
	return i
}

func (i *InfoPanel) AddActionButtonFront(title template.HTML, action Action, ids ...string) *InfoPanel {
	id := ""
	if len(ids) > 0 {
		id = ids[0]
	} else {
		id = "action-info-btn-" + utils.Uuid(10)
	}
	action.SetBtnId(id)
	i.ActionButtons = append([]Button{ActionButton{Title: title, Id: id, Action: action}}, i.ActionButtons...)
	i.Callbacks = i.Callbacks.AddCallback(action.GetCallbacks())
	return i
}

func (i *InfoPanel) AddLimitFilter(limit int) *InfoPanel {
	i.processChains = addLimit(limit, i.processChains)
	return i
}

func (i *InfoPanel) AddTrimSpaceFilter() *InfoPanel {
	i.processChains = addTrimSpace(i.processChains)
	return i
}

func (i *InfoPanel) AddSubstrFilter(start int, end int) *InfoPanel {
	i.processChains = addSubstr(start, end, i.processChains)
	return i
}

func (i *InfoPanel) AddToTitleFilter() *InfoPanel {
	i.processChains = addToTitle(i.processChains)
	return i
}

func (i *InfoPanel) AddToUpperFilter() *InfoPanel {
	i.processChains = addToUpper(i.processChains)
	return i
}

func (i *InfoPanel) AddToLowerFilter() *InfoPanel {
	i.processChains = addToLower(i.processChains)
	return i
}

func (i *InfoPanel) AddXssFilter() *InfoPanel {
	i.processChains = addXssFilter(i.processChains)
	return i
}

func (i *InfoPanel) AddXssJsFilter() *InfoPanel {
	i.processChains = addXssJsFilter(i.processChains)
	return i
}

func (i *InfoPanel) SetDeleteHook(fn DeleteFn) *InfoPanel {
	i.DeleteHook = fn
	return i
}

func (i *InfoPanel) SetPreDeleteFn(fn DeleteFn) *InfoPanel {
	i.PreDeleteFn = fn
	return i
}

func (i *InfoPanel) SetDeleteFn(fn DeleteFn) *InfoPanel {
	i.DeleteFn = fn
	return i
}

func (i *InfoPanel) SetGetDataFn(fn GetDataFn) *InfoPanel {
	i.GetDataFn = fn
	return i
}

func (i *InfoPanel) SetTableFixed() *InfoPanel {
	i.TableLayout = "fixed"
	return i
}

func (i *InfoPanel) AddField(head, field string, typeName db.DatabaseType) *InfoPanel {
	i.FieldList = append(i.FieldList, Field{
		Head:     head,
		Field:    field,
		TypeName: typeName,
		Sortable: false,
		EditAble: false,
		EditType: table.Text,
		FieldDisplay: FieldDisplay{
			Display: func(value FieldModel) interface{} {
				return value.Value
			},
			DisplayProcessChains: chooseDisplayProcessChains(i.processChains),
		},
	})
	i.curFieldListIndex++
	return i
}

// Field attribute setting functions
// ====================================================

func (i *InfoPanel) FieldDisplay(filter FieldFilterFn) *InfoPanel {
	i.FieldList[i.curFieldListIndex].Display = filter
	return i
}

func (i *InfoPanel) FieldWidth(width int) *InfoPanel {
	i.FieldList[i.curFieldListIndex].Width = width
	return i
}

func (i *InfoPanel) FieldSortable() *InfoPanel {
	i.FieldList[i.curFieldListIndex].Sortable = true
	return i
}

func (i *InfoPanel) FieldEditOptions(options []map[string]string) *InfoPanel {
	i.FieldList[i.curFieldListIndex].EditOptions = options
	return i
}

func (i *InfoPanel) FieldEditAble(editType ...table.Type) *InfoPanel {
	i.FieldList[i.curFieldListIndex].EditAble = true
	if len(editType) > 0 {
		i.FieldList[i.curFieldListIndex].EditType = editType[0]
	}
	return i
}

func (i *InfoPanel) FieldFixed() *InfoPanel {
	i.FieldList[i.curFieldListIndex].Fixed = true
	return i
}

type FilterType struct {
	FormType form.Type
	Operator FilterOperator
	Head     string
	HelpMsg  template.HTML
}

func (i *InfoPanel) FieldFilterable(filterType ...FilterType) *InfoPanel {
	i.FieldList[i.curFieldListIndex].Filterable = true
	if len(filterType) > 0 {
		i.FieldList[i.curFieldListIndex].FilterOperator = filterType[0].Operator
		if uint8(filterType[0].FormType) == 0 {
			i.FieldList[i.curFieldListIndex].FilterType = form.Text
		} else {
			i.FieldList[i.curFieldListIndex].FilterType = filterType[0].FormType
		}
		i.FieldList[i.curFieldListIndex].FilterHead = filterType[0].Head
		i.FieldList[i.curFieldListIndex].FilterHelpMsg = filterType[0].HelpMsg
	} else {
		i.FieldList[i.curFieldListIndex].FilterType = form.Text
	}
	return i
}

func (i *InfoPanel) FieldFilterOptions(options []map[string]string) *InfoPanel {
	i.FieldList[i.curFieldListIndex].FilterOptions = options
	i.FieldList[i.curFieldListIndex].FilterOptionExt = `{"allowClear": "true"}`
	return i
}

func (i *InfoPanel) FieldFilterProcess(process func(string) string) *InfoPanel {
	i.FieldList[i.curFieldListIndex].FilterProcess = process
	return i
}

func (i *InfoPanel) FieldFilterOptionExt(m map[string]interface{}) *InfoPanel {
	s, _ := json.Marshal(m)
	i.FieldList[i.curFieldListIndex].FilterOptionExt = template.JS(s)
	return i
}

func (i *InfoPanel) FieldHide() *InfoPanel {
	i.FieldList[i.curFieldListIndex].Hide = true
	return i
}

func (i *InfoPanel) FieldJoin(join Join) *InfoPanel {
	i.FieldList[i.curFieldListIndex].Join = join
	return i
}

func (i *InfoPanel) FieldLimit(limit int) *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].AddLimit(limit)
	return i
}

func (i *InfoPanel) FieldTrimSpace() *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].AddTrimSpace()
	return i
}

func (i *InfoPanel) FieldSubstr(start int, end int) *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].AddSubstr(start, end)
	return i
}

func (i *InfoPanel) FieldToTitle() *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].AddToTitle()
	return i
}

func (i *InfoPanel) FieldToUpper() *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].AddToUpper()
	return i
}

func (i *InfoPanel) FieldToLower() *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].AddToLower()
	return i
}

func (i *InfoPanel) FieldXssFilter() *InfoPanel {
	i.FieldList[i.curFieldListIndex].DisplayProcessChains = i.FieldList[i.curFieldListIndex].DisplayProcessChains.
		Add(func(s string) string {
			return html.EscapeString(s)
		})
	return i
}

// InfoPanel attribute setting functions
// ====================================================

func (i *InfoPanel) SetTable(table string) *InfoPanel {
	i.Table = table
	return i
}

func (i *InfoPanel) SetPageSizeList(pageSizeList []int) *InfoPanel {
	i.PageSizeList = pageSizeList
	return i
}

func (i *InfoPanel) SetDefaultPageSize(defaultPageSize int) *InfoPanel {
	i.DefaultPageSize = defaultPageSize
	return i
}

func (i *InfoPanel) GetPageSizeList() []string {
	var pageSizeList = make([]string, len(i.PageSizeList))
	for j := 0; j < len(i.PageSizeList); j++ {
		pageSizeList[j] = strconv.Itoa(i.PageSizeList[j])
	}
	return pageSizeList
}

func (i *InfoPanel) GetSort() string {
	switch i.Sort {
	case SortAsc:
		return "asc"
	default:
		return "desc"
	}
}

func (i *InfoPanel) SetTitle(title string) *InfoPanel {
	i.Title = title
	return i
}

func (i *InfoPanel) SetTabGroups(groups TabGroups) *InfoPanel {
	i.TabGroups = groups
	return i
}

func (i *InfoPanel) SetTabHeaders(headers ...string) *InfoPanel {
	i.TabHeaders = headers
	return i
}

func (i *InfoPanel) SetDescription(desc string) *InfoPanel {
	i.Description = desc
	return i
}

func (i *InfoPanel) SetFilterFormLayout(layout form.Layout) *InfoPanel {
	i.FilterFormLayout = layout
	return i
}

func (i *InfoPanel) SetSortField(field string) *InfoPanel {
	i.SortField = field
	return i
}

func (i *InfoPanel) SetSortAsc() *InfoPanel {
	i.Sort = SortAsc
	return i
}

func (i *InfoPanel) SetSortDesc() *InfoPanel {
	i.Sort = SortDesc
	return i
}

func (i *InfoPanel) SetAction(action template.HTML) *InfoPanel {
	i.Action = action
	return i
}

func (i *InfoPanel) SetHeaderHtml(header template.HTML) *InfoPanel {
	i.HeaderHtml = header
	return i
}

func (i *InfoPanel) SetFooterHtml(footer template.HTML) *InfoPanel {
	i.FooterHtml = footer
	return i
}

func (i *InfoPanel) HideNewButton() *InfoPanel {
	i.IsHideNewButton = true
	return i
}

func (i *InfoPanel) HideExportButton() *InfoPanel {
	i.IsHideExportButton = true
	return i
}

func (i *InfoPanel) HideFilterButton() *InfoPanel {
	i.IsHideFilterButton = true
	return i
}

func (i *InfoPanel) HideRowSelector() *InfoPanel {
	i.IsHideRowSelector = true
	return i
}

func (i *InfoPanel) HidePagination() *InfoPanel {
	i.IsHidePagination = true
	return i
}

func (i *InfoPanel) HideFilterArea() *InfoPanel {
	i.IsHideFilterArea = true
	return i
}

func (i *InfoPanel) HideEditButton() *InfoPanel {
	i.IsHideEditButton = true
	return i
}

func (i *InfoPanel) HideDeleteButton() *InfoPanel {
	i.IsHideDeleteButton = true
	return i
}

func (i *InfoPanel) HideDetailButton() *InfoPanel {
	i.IsHideDetailButton = true
	return i
}
