package types

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
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

type InfoList []map[string]InfoItem

type InfoItem struct {
	Content template.HTML `json:"content"`
	Value   string        `json:"value"`
}

func (i InfoList) GroupBy(groups TabGroups) []InfoList {

	var res = make([]InfoList, len(groups))

	for key, value := range groups {
		var newInfoList = make(InfoList, len(i))

		for index, info := range i {
			var newRow = make(map[string]InfoItem)
			for mk, m := range info {
				if modules.InArray(value, mk) {
					newRow[mk] = m
				}
			}
			newInfoList[index] = newRow
		}

		res[key] = newInfoList
	}

	return res
}

type Callbacks []context.Node

func (c Callbacks) AddCallback(node context.Node) Callbacks {
	if node.Path != "" && node.Method != "" && len(node.Handlers) > 0 {
		for _, item := range c {
			if strings.ToUpper(item.Path) == strings.ToUpper(node.Path) &&
				strings.ToUpper(item.Method) == strings.ToUpper(node.Method) {
				return c
			}
		}
		parr := strings.Split(node.Path, "?")
		if len(parr) > 1 {
			node.Path = parr[0]
			return append(c, node)
		}
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
	EditOptions FieldOptions

	FilterFormFields []FilterFormField

	FieldDisplay
}

type FilterFormField struct {
	Type        form.Type
	Options     FieldOptions
	OptionTable OptionTable
	Width       int
	Operator    FilterOperator
	OptionExt   template.JS
	Head        string
	Placeholder string
	HelpMsg     template.HTML
	ProcessFn   func(string) string
}

func (f Field) GetFilterFormFields(params parameter.Parameters, headField string, sqls ...*db.SQL) []FormField {

	var (
		filterForm = make([]FormField, 0)
		options    = make(FieldOptions, 0)

		value, value2, keySuffix string
	)

	for index, filter := range f.FilterFormFields {

		if index > 0 {
			keySuffix = parameter.FilterParamCountInfix + strconv.Itoa(index)
		}

		if filter.Type.IsRange() {
			value = params.GetFilterFieldValueStart(headField)
			value2 = params.GetFilterFieldValueEnd(headField)
		} else if filter.Type.IsMultiSelect() {
			value = params.GetFieldValuesStr(headField)
		} else {
			if filter.Operator == FilterOperatorFree {
				value2 = GetOperatorFromValue(params.GetFieldOperator(headField, keySuffix)).String()
			}
			value = params.GetFieldValue(headField + keySuffix)
		}

		options = make(FieldOptions, 0)

		if len(filter.Options) == 0 && filter.OptionTable.Table != "" && len(sqls) > 0 && sqls[0] != nil {
			sqls[0].Table(filter.OptionTable.Table).Select(filter.OptionTable.ValueField, filter.OptionTable.TextField)

			if filter.OptionTable.QueryProcessFn != nil {
				filter.OptionTable.QueryProcessFn(sqls[0])
			}

			queryRes, err := sqls[0].All()
			if err == nil {
				for _, item := range queryRes {
					filter.Options = append(filter.Options, FieldOption{
						Value: fmt.Sprintf("%v", item[filter.OptionTable.ValueField]),
						Text:  fmt.Sprintf("%v", item[filter.OptionTable.TextField]),
					})
				}
			}

			if filter.OptionTable.ProcessFn != nil {
				filter.Options = filter.OptionTable.ProcessFn(filter.Options)
			}
		}

		if filter.Type.IsSingleSelect() {
			options = filter.Options.SetSelected(params.GetFieldValue(f.Field), filter.Type.SelectedLabel())
		}

		if filter.Type.IsMultiSelect() {
			options = filter.Options.SetSelected(params.GetFieldValues(f.Field), filter.Type.SelectedLabel())
		}

		filterForm = append(filterForm, FormField{
			Field:       headField + keySuffix,
			Head:        filter.Head,
			TypeName:    f.TypeName,
			HelpMsg:     filter.HelpMsg,
			FormType:    filter.Type,
			Editable:    true,
			Width:       filter.Width,
			Placeholder: filter.Placeholder,
			Value:       template.HTML(value),
			Value2:      value2,
			Options:     options,
			OptionExt:   filter.OptionExt,
			OptionTable: filter.OptionTable,
			Label:       filter.Operator.Label(),
		})

		if filter.Operator.AddOrNot() {
			filterForm = append(filterForm, FormField{
				Field:    headField + parameter.FilterParamOperatorSuffix + keySuffix,
				Head:     f.Head,
				TypeName: f.TypeName,
				Value:    template.HTML(filter.Operator.Value()),
				FormType: filter.Type,
				Hide:     true,
			})
		}
	}

	return filterForm
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

func (f FieldList) GetTheadAndFilterForm(info TableInfo, params parameter.Parameters, columns []string, sql ...func() *db.SQL) (Thead,
	string, string, string, []string, []FormField) {
	var (
		thead      = make(Thead, 0)
		fields     = ""
		joinFields = ""
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
			headField = field.Join.Table + parameter.FilterParamJoinInfix + field.Field
			joinFields += db.GetAggregationExpression(info.Driver, field.Join.Table+"."+
				modules.FilterField(field.Field, info.Delimiter), headField, JoinFieldValueDelimiter) + ","
			if !modules.InArray(joinTables, field.Join.Table) {
				joinTables = append(joinTables, field.Join.Table)
				joins += " left join " + modules.FilterField(field.Join.Table, info.Delimiter) + " on " +
					field.Join.Table + "." + modules.FilterField(field.Join.JoinField, info.Delimiter) + " = " +
					info.Table + "." + modules.FilterField(field.Join.Field, info.Delimiter)
			}
		}

		if field.Filterable {
			if len(sql) > 0 {
				filterForm = append(filterForm, field.GetFilterFormFields(params, headField, sql[0]())...)
			} else {
				filterForm = append(filterForm, field.GetFilterFormFields(params, headField)...)
			}
		}

		if field.Hide {
			continue
		}
		thead = append(thead, TheadItem{
			Head:       field.Head,
			Sortable:   field.Sortable,
			Field:      headField,
			Hide:       !modules.InArrayWithoutEmpty(params.Columns, headField),
			Editable:   field.EditAble,
			EditType:   field.EditType.String(),
			EditOption: field.EditOptions,
			Width:      field.Width,
		})
	}

	return thead, fields, joinFields, joins, joinTables, filterForm
}

func (f FieldList) GetThead(info TableInfo, params parameter.Parameters, columns []string) (Thead, string, string) {
	var (
		thead      = make(Thead, 0)
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
			headField = field.Join.Table + parameter.FilterParamJoinInfix + field.Field
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
		thead = append(thead, TheadItem{
			Head:       field.Head,
			Sortable:   field.Sortable,
			Field:      headField,
			Hide:       !modules.InArrayWithoutEmpty(params.Columns, headField),
			Editable:   field.EditAble,
			EditType:   field.EditType.String(),
			EditOption: field.EditOptions,
			Width:      field.Width,
		})
	}

	return thead, fields, joins
}

func (f FieldList) GetFieldFilterProcessValue(key, value, keyIndex string) string {
	field := f.GetFieldByFieldName(key)
	index := 0
	if keyIndex != "" {
		index, _ = strconv.Atoi(keyIndex)
	}
	if field.FilterFormFields[index].ProcessFn != nil {
		value = field.FilterFormFields[index].ProcessFn(value)
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
		if JoinField(field.Join.Table, field.Field) == name {
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

func JoinField(table, field string) string {
	return table + parameter.FilterParamJoinInfix + field
}

func GetJoinField(field string) string {
	return strings.Split(field, parameter.FilterParamJoinInfix)[1]
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
type DeleteFnWithRes func(ids []string, res error) error

type Sort uint8

const (
	SortDesc Sort = iota
	SortAsc
)

type primaryKey struct {
	Type db.DatabaseType
	Name string
}

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

	ExportType int

	primaryKey primaryKey

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

	DeleteHookWithRes DeleteFnWithRes

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
	pwheres := ""
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
				pwheres += whTable + "." + modules.FilterField(whField, delimiter) + " " + wh.Operator + " ? " + joinMark + " "
			} else {
				pwheres += modules.FilterField(whField, delimiter) + " " + wh.Operator + " ? " + joinMark + " "
			}
			whereArgs = append(whereArgs, wh.Arg)
		}
	}
	if wheres != "" && pwheres != "" {
		wheres += " and "
	}
	return wheres + pwheres, whereArgs
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
	FooterContent() template.HTML
	SetBtnId(btnId string)
	SetBtnData(data interface{})
	GetCallbacks() context.Node
}

type DefaultAction struct {
	Attr   template.HTML
	JS     template.JS
	Ext    template.HTML
	Footer template.HTML
}

func NewDefaultAction(attr, ext, footer template.HTML, js template.JS) *DefaultAction {
	return &DefaultAction{Attr: attr, Ext: ext, Footer: footer, JS: js}
}

func (def *DefaultAction) SetBtnId(btnId string)        {}
func (def *DefaultAction) SetBtnData(data interface{})  {}
func (def *DefaultAction) Js() template.JS              { return def.JS }
func (def *DefaultAction) BtnAttribute() template.HTML  { return def.Attr }
func (def *DefaultAction) BtnClass() template.HTML      { return "" }
func (def *DefaultAction) ExtContent() template.HTML    { return def.Ext }
func (def *DefaultAction) FooterContent() template.HTML { return def.Footer }
func (def *DefaultAction) GetCallbacks() context.Node   { return context.Node{} }

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

func (i *InfoPanel) AddSelectBox(placeholder string, options FieldOptions, action Action, width ...int) *InfoPanel {
	options = append(FieldOptions{{Value: "", Text: language.Get("All")}}, options...)
	action.SetBtnData(options)
	i.addButton(GetDefaultSelection(placeholder, options, action, width...)).
		addFooterHTML(action.FooterContent()).
		addCallback(action.GetCallbacks())

	return i
}

func (i *InfoPanel) ExportValue() *InfoPanel {
	i.ExportType = 1
	return i
}

func (i *InfoPanel) IsExportValue() bool {
	return i.ExportType == 1
}

func (i *InfoPanel) AddButtonRaw(btn Button, action Action) *InfoPanel {
	i.Buttons = append(i.Buttons, btn)
	i.addFooterHTML(action.FooterContent()).addCallback(action.GetCallbacks())
	return i
}

func (i *InfoPanel) AddButton(title template.HTML, icon string, action Action, color ...template.HTML) *InfoPanel {
	i.addButton(GetDefaultButton(title, icon, action, color...)).
		addFooterHTML(action.FooterContent()).
		addCallback(action.GetCallbacks())
	return i
}

func (i *InfoPanel) AddActionButton(title template.HTML, action Action, ids ...string) *InfoPanel {
	i.addActionButton(GetActionButton(title, action, ids...)).
		addFooterHTML(action.FooterContent()).
		addCallback(action.GetCallbacks())

	return i
}

func (i *InfoPanel) AddActionButtonFront(title template.HTML, action Action, ids ...string) *InfoPanel {
	i.ActionButtons = append([]Button{GetActionButton(title, action, ids...)}, i.ActionButtons...)
	i.addFooterHTML(action.FooterContent()).
		addCallback(action.GetCallbacks())
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

func (i *InfoPanel) SetDeleteHookWithRes(fn DeleteFnWithRes) *InfoPanel {
	i.DeleteHookWithRes = fn
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

func (i *InfoPanel) SetPrimaryKey(name string, typ db.DatabaseType) *InfoPanel {
	i.primaryKey = primaryKey{Name: name, Type: typ}
	return i
}

func (i *InfoPanel) SetTableFixed() *InfoPanel {
	i.TableLayout = "fixed"
	return i
}

func (i *InfoPanel) AddColumn(head string, fun FieldFilterFn) *InfoPanel {
	i.FieldList = append(i.FieldList, Field{
		Head:     head,
		Field:    head,
		TypeName: db.Varchar,
		Sortable: false,
		EditAble: false,
		EditType: table.Text,
		FieldDisplay: FieldDisplay{
			Display:              fun,
			DisplayProcessChains: chooseDisplayProcessChains(i.processChains),
		},
	})
	i.curFieldListIndex++
	return i
}

func (i *InfoPanel) AddColumnButtons(head string, buttons ...Button) *InfoPanel {
	var content, js template.HTML
	for _, btn := range buttons {
		btn.GetAction().SetBtnId(btn.ID())
		btnContent, btnJs := btn.Content()
		content += btnContent
		js += template.HTML(btnJs)
		i.FooterHtml += template.HTML(ParseTableDataTmpl(btn.GetAction().FooterContent()))
		i.Callbacks = i.Callbacks.AddCallback(btn.GetAction().GetCallbacks())
	}
	i.FooterHtml += template.HTML("<script>") + template.HTML(ParseTableDataTmpl(js)) + template.HTML("</script>")
	i.FieldList = append(i.FieldList, Field{
		Head:     head,
		Field:    head,
		TypeName: db.Varchar,
		Sortable: false,
		EditAble: false,
		EditType: table.Text,
		FieldDisplay: FieldDisplay{
			Display: func(value FieldModel) interface{} {
				pk := db.GetValueFromDatabaseType(i.primaryKey.Type, value.Row[i.primaryKey.Name], i.isFromJSON())
				return template.HTML(ParseTableDataTmplWithID(pk.HTML(), string(content)))
			},
			DisplayProcessChains: chooseDisplayProcessChains(i.processChains),
		},
	})
	i.curFieldListIndex++
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

func (i *InfoPanel) FieldEditOptions(options FieldOptions, extra ...map[string]string) *InfoPanel {
	if i.FieldList[i.curFieldListIndex].EditType.IsSwitch() {
		if len(extra) == 0 {
			options[0].Extra = map[string]string{
				"size":     "small",
				"onColor":  "primary",
				"offColor": "default",
			}
		} else {
			if extra[0]["size"] == "" {
				extra[0]["size"] = "small"
			}
			if extra[0]["onColor"] == "" {
				extra[0]["onColor"] = "primary"
			}
			if extra[0]["offColor"] == "" {
				extra[0]["offColor"] = "default"
			}
			options[0].Extra = extra[0]
		}
	}
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
	FormType    form.Type
	Operator    FilterOperator
	Head        string
	Placeholder string
	NoHead      bool
	Width       int
	HelpMsg     template.HTML
	Options     FieldOptions
	Process     func(string) string
	OptionExt   map[string]interface{}
}

func (i *InfoPanel) FieldFilterable(filterType ...FilterType) *InfoPanel {
	i.FieldList[i.curFieldListIndex].Filterable = true

	if len(filterType) == 0 {
		i.FieldList[i.curFieldListIndex].FilterFormFields = append(i.FieldList[i.curFieldListIndex].FilterFormFields,
			FilterFormField{
				Type:        form.Text,
				Head:        i.FieldList[i.curFieldListIndex].Head,
				Placeholder: language.Get("input") + " " + i.FieldList[i.curFieldListIndex].Head,
			})
	}

	for _, filter := range filterType {
		var ff FilterFormField
		ff.Operator = filter.Operator
		if filter.FormType == form.Default {
			ff.Type = form.Text
		} else {
			ff.Type = filter.FormType
		}
		ff.Head = modules.AorB(!filter.NoHead && filter.Head == "",
			i.FieldList[i.curFieldListIndex].Head, filter.Head)
		ff.Width = filter.Width
		ff.HelpMsg = filter.HelpMsg
		ff.ProcessFn = filter.Process
		ff.Placeholder = modules.AorB(filter.Placeholder == "", language.Get("input")+" "+ff.Head, filter.Placeholder)
		ff.Options = filter.Options
		if len(filter.OptionExt) > 0 {
			s, _ := json.Marshal(filter.OptionExt)
			ff.OptionExt = template.JS(s)
		}
		i.FieldList[i.curFieldListIndex].FilterFormFields = append(i.FieldList[i.curFieldListIndex].FilterFormFields, ff)
	}

	return i
}

func (i *InfoPanel) FieldFilterOptions(options FieldOptions) *InfoPanel {
	i.FieldList[i.curFieldListIndex].FilterFormFields[0].Options = options
	i.FieldList[i.curFieldListIndex].FilterFormFields[0].OptionExt = `{"allowClear": "true"}`
	return i
}

func (i *InfoPanel) FieldFilterOptionsFromTable(table, textFieldName, valueFieldName string, process ...OptionTableQueryProcessFn) *InfoPanel {
	var fn OptionTableQueryProcessFn
	if len(process) > 0 {
		fn = process[0]
	}
	i.FieldList[i.curFieldListIndex].FilterFormFields[0].OptionTable = OptionTable{
		Table:          table,
		TextField:      textFieldName,
		ValueField:     valueFieldName,
		QueryProcessFn: fn,
	}
	return i
}

func (i *InfoPanel) FieldFilterProcess(process func(string) string) *InfoPanel {
	i.FieldList[i.curFieldListIndex].FilterFormFields[0].ProcessFn = process
	return i
}

func (i *InfoPanel) FieldFilterOptionExt(m map[string]interface{}) *InfoPanel {
	s, _ := json.Marshal(m)
	i.FieldList[i.curFieldListIndex].FilterFormFields[0].OptionExt = template.JS(s)
	return i
}

func (i *InfoPanel) FieldFilterOnSearch(url string, handler Handler, delay ...int) *InfoPanel {
	ext, callback := searchJS(i.FieldList[i.curFieldListIndex].FilterFormFields[0].OptionExt, url, handler, delay...)
	i.FieldList[i.curFieldListIndex].FilterFormFields[0].OptionExt = ext
	i.Callbacks = append(i.Callbacks, callback)
	return i
}

func (i *InfoPanel) FieldFilterOnChooseCustom(js template.HTML) *InfoPanel {
	i.FooterHtml += chooseCustomJS(i.FieldList[i.curFieldListIndex].Field, js)
	return i
}

func (i *InfoPanel) FieldFilterOnChooseMap(m map[string]LinkField) *InfoPanel {
	i.FooterHtml += chooseMapJS(i.FieldList[i.curFieldListIndex].Field, m)
	return i
}

func (i *InfoPanel) FieldFilterOnChoose(val, field string, value template.HTML) *InfoPanel {
	i.FooterHtml += chooseJS(i.FieldList[i.curFieldListIndex].Field, field, val, value)
	return i
}

func (i *InfoPanel) FieldFilterOnChooseAjax(field, url string, handler Handler) *InfoPanel {
	js, callback := chooseAjax(i.FieldList[i.curFieldListIndex].Field, field, url, handler)
	i.FooterHtml += js
	i.Callbacks = append(i.Callbacks, callback)
	return i
}

func (i *InfoPanel) FieldFilterOnChooseHide(value string, field ...string) *InfoPanel {
	i.FooterHtml += chooseHideJS(i.FieldList[i.curFieldListIndex].Field, value, field...)
	return i
}

func (i *InfoPanel) FieldFilterOnChooseDisable(value string, field ...string) *InfoPanel {
	i.FooterHtml += chooseDisableJS(i.FieldList[i.curFieldListIndex].Field, value, field...)
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

func (i *InfoPanel) addFooterHTML(footer template.HTML) *InfoPanel {
	i.FooterHtml += template.HTML(ParseTableDataTmpl(footer))
	return i
}

func (i *InfoPanel) addCallback(node context.Node) *InfoPanel {
	i.Callbacks = i.Callbacks.AddCallback(node)
	return i
}

func (i *InfoPanel) addButton(btn Button) *InfoPanel {
	i.Buttons = append(i.Buttons, btn)
	return i
}

func (i *InfoPanel) addActionButton(btn Button) *InfoPanel {
	i.ActionButtons = append(i.ActionButtons, btn)
	return i
}

func (i *InfoPanel) isFromJSON() bool {
	return i.GetDataFn != nil
}
