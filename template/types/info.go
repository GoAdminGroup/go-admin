package types

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/GoAdminGroup/go-admin/template/types/table"
	"html"
	"html/template"
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
type PostFieldFilterFn func(value PostFieldModel) string

type DisplayProcessFn func(string) string

type DisplayProcessFnChains []DisplayProcessFn

func (d DisplayProcessFnChains) Valid() bool {
	return len(d) > 0
}

func (d DisplayProcessFnChains) Add(f DisplayProcessFn) DisplayProcessFnChains {
	return append(d, f)
}

func (d DisplayProcessFnChains) Append(f DisplayProcessFnChains) DisplayProcessFnChains {
	return append(d, f...)
}

func (d DisplayProcessFnChains) Copy() DisplayProcessFnChains {
	if len(d) == 0 {
		return make(DisplayProcessFnChains, 0)
	} else {
		var newDisplayProcessFnChains = make(DisplayProcessFnChains, len(d))
		copy(newDisplayProcessFnChains, d)
		return newDisplayProcessFnChains
	}
}

func chooseDisplayProcessChains(internal DisplayProcessFnChains) DisplayProcessFnChains {
	if len(internal) > 0 {
		return internal
	}
	return globalDisplayProcessChains.Copy()
}

var globalDisplayProcessChains = make(DisplayProcessFnChains, 0)

func AddGlobalDisplayProcessFn(f DisplayProcessFn) {
	globalDisplayProcessChains = globalDisplayProcessChains.Add(f)
}

func AddLimit(limit int) DisplayProcessFnChains {
	return addLimit(limit, globalDisplayProcessChains)
}

func AddTrimSpace() DisplayProcessFnChains {
	return addTrimSpace(globalDisplayProcessChains)
}

func AddSubstr(start int, end int) DisplayProcessFnChains {
	return addSubstr(start, end, globalDisplayProcessChains)
}

func AddToTitle() DisplayProcessFnChains {
	return addToTitle(globalDisplayProcessChains)
}

func AddToUpper() DisplayProcessFnChains {
	return addToUpper(globalDisplayProcessChains)
}

func AddToLower() DisplayProcessFnChains {
	return addToLower(globalDisplayProcessChains)
}

func AddXssFilter() DisplayProcessFnChains {
	return addXssFilter(globalDisplayProcessChains)
}

func AddXssJsFilter() DisplayProcessFnChains {
	return addXssJsFilter(globalDisplayProcessChains)
}

func addLimit(limit int, chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		if limit > len(value) {
			return value
		} else if limit < 0 {
			return ""
		} else {
			return value[:limit]
		}
	})
	return chains
}

func addTrimSpace(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.TrimSpace(value)
	})
	return chains
}

func addSubstr(start int, end int, chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		if start > end || start > len(value) || end < 0 {
			return ""
		}
		if start < 0 {
			start = 0
		}
		if end > len(value) {
			end = len(value)
		}
		return value[start:end]
	})
	return chains
}

func addToTitle(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.Title(value)
	})
	return chains
}

func addToUpper(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.ToUpper(value)
	})
	return chains
}

func addToLower(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return strings.ToLower(value)
	})
	return chains
}

func addXssFilter(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		return html.EscapeString(value)
	})
	return chains
}

func addXssJsFilter(chains DisplayProcessFnChains) DisplayProcessFnChains {
	chains = chains.Add(func(value string) string {
		s := strings.Replace(value, "<script>", "&lt;script&gt;", -1)
		return strings.Replace(s, "</script>", "&lt;/script&gt;", -1)
	})
	return chains
}

type FieldDisplay struct {
	Display              FieldFilterFn
	DisplayProcessChains DisplayProcessFnChains
}

func (f FieldDisplay) ToDisplay(value FieldModel) interface{} {
	val := f.Display(value)

	if _, ok := val.(template.HTML); !ok {
		if _, ok2 := val.([]string); !ok2 {
			valStr := fmt.Sprintf("%v", val)
			for _, process := range f.DisplayProcessChains {
				valStr = process(valStr)
			}
			return valStr
		}
	}

	return val
}

func (f FieldDisplay) AddLimit(limit int) DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		if limit > len(value) {
			return value
		} else if limit < 0 {
			return ""
		} else {
			return value[:limit]
		}
	})
}

func (f FieldDisplay) AddTrimSpace() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.TrimSpace(value)
	})
}

func (f FieldDisplay) AddSubstr(start int, end int) DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		if start > end || start > len(value) || end < 0 {
			return ""
		}
		if start < 0 {
			start = 0
		}
		if end > len(value) {
			end = len(value)
		}
		return value[start:end]
	})
}

func (f FieldDisplay) AddToTitle() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.Title(value)
	})
}

func (f FieldDisplay) AddToUpper() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.ToUpper(value)
	})
}

func (f FieldDisplay) AddToLower() DisplayProcessFnChains {
	return f.DisplayProcessChains.Add(func(value string) string {
		return strings.ToLower(value)
	})
}

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

	FieldDisplay
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

	Sort Sort

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

	Wheres []Where

	Buttons Buttons

	DeleteHook  DeleteFn
	PreDeleteFn DeleteFn
	DeleteFn    DeleteFn

	processChains DisplayProcessFnChains

	ActionButtons Buttons
	Action        template.HTML
	HeaderHtml    template.HTML
	FooterHtml    template.HTML
}

type Where struct {
	Field    string
	Operator string
	Arg      interface{}
}

type Action interface {
	Js() template.JS
	BtnAttribute() template.HTML
	ExtContent() template.HTML
	SetBtnId(btnId string)
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
func (def *DefaultAction) ExtContent() template.HTML   { return def.Ext }

type Button interface {
	Content() (template.HTML, template.JS)
}

type DefaultButton struct {
	Id        string
	Title     template.HTML
	Color     template.HTML
	TextColor template.HTML
	Action    Action
	Icon      string
}

func (b DefaultButton) Content() (template.HTML, template.JS) {

	color := template.HTML("")
	if b.Color != template.HTML("") {
		color = template.HTML(`background-color:`) + b.Color + template.HTML(`;`)
	}
	textColor := template.HTML("")
	if b.TextColor != template.HTML("") {
		textColor = template.HTML(`color:`) + b.TextColor + template.HTML(`;`)
	}

	style := template.HTML("")
	addColor := color + textColor

	if addColor != template.HTML("") {
		style = template.HTML(`style="`) + addColor + template.HTML(`"`)
	}

	h := `<div class="btn-group pull-right" style="margin-right: 10px">
                <a id="` + template.HTML(b.Id) + `" ` + style + `class="btn btn-sm btn-default" ` + b.Action.BtnAttribute() + `>
                    <i class="fa ` + template.HTML(b.Icon) + `"></i>&nbsp;&nbsp;` + b.Title + `
                </a>
        </div>` + b.Action.ExtContent()
	return h, b.Action.Js()
}

type ActionButton struct {
	Id     string
	Title  template.HTML
	Action Action
}

func (b ActionButton) Content() (template.HTML, template.JS) {
	h := template.HTML(`<li><a class="`+template.HTML(b.Id)+`" `+b.Action.BtnAttribute()+`>`+b.Title+`</a></li>`) + b.Action.ExtContent()
	return h, b.Action.Js()
}

type Buttons []Button

func (b Buttons) Content() (template.HTML, template.JS) {
	h := template.HTML("")
	j := template.JS("")

	for _, btn := range b {
		hh, jj := btn.Content()
		h += hh
		j += jj
	}
	return h, j
}

var DefaultPageSizeList = []int{10, 20, 30, 50, 100}

const DefaultPageSize = 10

func NewInfoPanel() *InfoPanel {
	return &InfoPanel{
		curFieldListIndex: -1,
		PageSizeList:      DefaultPageSizeList,
		DefaultPageSize:   DefaultPageSize,
		processChains:     make(DisplayProcessFnChains, 0),
		Buttons:           make(Buttons, 0),
		Wheres:            make([]Where, 0),
	}
}

func (i *InfoPanel) Where(field string, operator string, arg interface{}) *InfoPanel {
	i.Wheres = append(i.Wheres, Where{Field: field, Operator: operator, Arg: arg})
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

type FilterOperator string

const (
	FilterOperatorLike           FilterOperator = "like"
	FilterOperatorGreater        FilterOperator = ">"
	FilterOperatorGreaterOrEqual FilterOperator = ">="
	FilterOperatorEqual          FilterOperator = "="
	FilterOperatorNotEqual       FilterOperator = "!="
	FilterOperatorLess           FilterOperator = "<"
	FilterOperatorLessOrEqual    FilterOperator = "<="
	FilterOperatorFree           FilterOperator = "free"
)

func GetOperatorFromValue(value string) FilterOperator {
	switch value {
	case "like":
		return FilterOperatorLike
	case "gr":
		return FilterOperatorGreater
	case "gq":
		return FilterOperatorGreaterOrEqual
	case "eq":
		return FilterOperatorEqual
	case "ne":
		return FilterOperatorNotEqual
	case "le":
		return FilterOperatorLess
	case "lq":
		return FilterOperatorLessOrEqual
	case "free":
		return FilterOperatorFree
	default:
		return FilterOperatorEqual
	}
}

func (o FilterOperator) Value() string {
	switch o {
	case FilterOperatorLike:
		return "like"
	case FilterOperatorGreater:
		return "gr"
	case FilterOperatorGreaterOrEqual:
		return "gq"
	case FilterOperatorEqual:
		return "eq"
	case FilterOperatorNotEqual:
		return "ne"
	case FilterOperatorLess:
		return "le"
	case FilterOperatorLessOrEqual:
		return "lq"
	case FilterOperatorFree:
		return "free"
	default:
		return "eq"
	}
}

func (o FilterOperator) String() string {
	return string(o)
}

func (o FilterOperator) Label() template.HTML {
	if o == FilterOperatorLike {
		return ""
	}
	return template.HTML(o)
}

func (o FilterOperator) AddOrNot() bool {
	return string(o) != "" && o != FilterOperatorFree
}

func (o FilterOperator) Valid() bool {
	switch o {
	case FilterOperatorLike, FilterOperatorGreater, FilterOperatorGreaterOrEqual,
		FilterOperatorLess, FilterOperatorLessOrEqual, FilterOperatorFree:
		return true
	default:
		return false
	}
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
