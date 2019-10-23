// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
	"strings"
)

// Attribute is the component interface of template. Every component of
// template should implement it.
type Attribute struct {
	TemplateList map[string]string
}

// Page used in the template as a top variable.
type Page struct {
	// User is the login user.
	User models.UserModel

	// Menu is the left side menu of the template.
	Menu menu.Menu

	// Panel is the main content of template.
	Panel Panel

	// System contains some system info.
	System SystemInfo

	// UrlPrefix is the prefix of url.
	UrlPrefix string

	// Title is the title of the web page.
	Title string

	// Logo is the logo of the template.
	Logo template.HTML

	// MiniLogo is the downsizing logo of the template.
	MiniLogo template.HTML

	// ColorScheme is the color scheme of the template.
	ColorScheme string

	// IndexUrl is the home page url of the site.
	IndexUrl string

	// AssetUrl is the cdn link of assets
	CdnUrl string

	// Custom html in the tag head.
	CustomHeadHtml template.HTML

	// Custom html after body.
	CustomFootHtml template.HTML
}

func NewPage(user models.UserModel, menu menu.Menu, panel Panel, cfg config.Config) Page {
	return Page{
		User:  user,
		Menu:  menu,
		Panel: panel,
		System: SystemInfo{
			Version: system.Version,
		},
		UrlPrefix:      cfg.Prefix(),
		Title:          cfg.Title,
		Logo:           cfg.Logo,
		MiniLogo:       cfg.MiniLogo,
		ColorScheme:    cfg.ColorScheme,
		IndexUrl:       cfg.GetIndexUrl(),
		CdnUrl:         cfg.AssetUrl,
		CustomHeadHtml: cfg.CustomHeadHtml,
		CustomFootHtml: cfg.CustomFootHtml,
	}
}

// SystemInfo contains basic info of system.
type SystemInfo struct {
	Version string
}

// Panel contains the main content of the template which used as pjax.
type Panel struct {
	Content     template.HTML
	Title       string
	Description string
	Url         string
}

type GetPanel func(ctx interface{}) (Panel, error)

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

type FieldDisplay struct {
	Display              FieldFilterFn
	DisplayProcessChains DisplayProcessFnChains
}

func (f FieldDisplay) ToDisplay(value FieldModel) interface{} {
	val := f.Display(value)

	if valStr, ok := val.(string); ok {
		for _, process := range f.DisplayProcessChains {
			valStr = process(valStr)
		}
		return valStr
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
	Fixed      bool
	Filterable bool
	Hide       bool

	FieldDisplay
}

type Join struct {
	Table     string
	Field     string
	JoinField string
}

func (j Join) Valid() bool {
	return j.Table != "" && j.Field != "" && j.JoinField != ""
}

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

// InfoPanel
type InfoPanel struct {
	FieldList         []Field
	curFieldListIndex int

	Table       string
	Title       string
	Description string

	// Warn: may be deprecated future.
	TabGroups  TabGroups
	TabHeaders TabHeaders

	Sort Sort

	IsHideNewButton    bool
	IsHideExportButton bool
	IsHideEditButton   bool
	IsHideDeleteButton bool
	IsHideFilterButton bool
	IsHideRowSelector  bool
	IsHidePagination   bool

	Action     template.HTML
	HeaderHtml template.HTML
	FooterHtml template.HTML
}

func NewInfoPanel() *InfoPanel {
	return &InfoPanel{curFieldListIndex: -1}
}

func (i *InfoPanel) AddField(head, field string, typeName db.DatabaseType) *InfoPanel {
	i.FieldList = append(i.FieldList, Field{
		Head:     head,
		Field:    field,
		TypeName: typeName,
		Sortable: false,
		FieldDisplay: FieldDisplay{
			Display: func(value FieldModel) interface{} {
				return value.Value
			},
			DisplayProcessChains: make(DisplayProcessFnChains, 0),
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

func (i *InfoPanel) FieldFixed() *InfoPanel {
	i.FieldList[i.curFieldListIndex].Fixed = true
	return i
}

func (i *InfoPanel) FieldFilterable() *InfoPanel {
	i.FieldList[i.curFieldListIndex].Filterable = true
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

// InfoPanel attribute setting functions
// ====================================================

func (i *InfoPanel) SetTable(table string) *InfoPanel {
	i.Table = table
	return i
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

func (i *InfoPanel) SetSort(sort Sort) *InfoPanel {
	i.Sort = sort
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

func (i *InfoPanel) HideEditButton() *InfoPanel {
	i.IsHideEditButton = true
	return i
}

func (i *InfoPanel) HideDeleteButton() *InfoPanel {
	i.IsHideDeleteButton = true
	return i
}

type Sort uint8

const (
	SortDesc Sort = iota
	SortAsc
)

// FormField is the form field with different options.
type FormField struct {
	Field    string
	TypeName db.DatabaseType
	Head     string
	FormType form.Type

	Default                string
	Value                  string
	Options                []map[string]string
	DefaultOptionDelimiter string

	Editable    bool
	NotAllowAdd bool
	Must        bool

	FieldDisplay
	PostFilterFn PostFieldFilterFn
}

// FormPanel
type FormPanel struct {
	FieldList         FormFields
	curFieldListIndex int

	// Warn: may be deprecated future.
	TabGroups  TabGroups
	TabHeaders TabHeaders

	Table       string
	Title       string
	Description string

	Validator FormValidator
	PostHook  FormPostHookFn

	HeaderHtml template.HTML
	FooterHtml template.HTML
}

func NewFormPanel() *FormPanel {
	return &FormPanel{curFieldListIndex: -1}
}

func (f *FormPanel) AddField(head, field string, filedType db.DatabaseType, formType form.Type) *FormPanel {
	f.FieldList = append(f.FieldList, FormField{
		Head:     head,
		Field:    field,
		TypeName: filedType,
		Editable: true,
		FormType: formType,
		FieldDisplay: FieldDisplay{
			Display: func(value FieldModel) interface{} {
				return value.Value
			},
			DisplayProcessChains: make(DisplayProcessFnChains, 0),
		},
	})
	f.curFieldListIndex++
	return f
}

// Field attribute setting functions
// ====================================================

func (f *FormPanel) FieldDisplay(filter FieldFilterFn) *FormPanel {
	f.FieldList[f.curFieldListIndex].Display = filter
	return f
}

func (f *FormPanel) SetTable(table string) *FormPanel {
	f.Table = table
	return f
}

func (f *FormPanel) FieldMust() *FormPanel {
	f.FieldList[f.curFieldListIndex].Must = true
	return f
}

func (f *FormPanel) FieldDefault(def string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Default = def
	return f
}

func (f *FormPanel) FieldNotAllowEdit() *FormPanel {
	f.FieldList[f.curFieldListIndex].Editable = false
	return f
}

func (f *FormPanel) FieldNotAllowAdd() *FormPanel {
	f.FieldList[f.curFieldListIndex].NotAllowAdd = true
	return f
}

func (f *FormPanel) FieldFormType(formType form.Type) *FormPanel {
	f.FieldList[f.curFieldListIndex].FormType = formType
	return f
}

func (f *FormPanel) FieldValue(value string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Value = value
	return f
}

func (f *FormPanel) FieldOptions(options []map[string]string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Options = options
	return f
}

func (f *FormPanel) FieldDefaultOptionDelimiter(delimiter string) *FormPanel {
	f.FieldList[f.curFieldListIndex].DefaultOptionDelimiter = delimiter
	return f
}

func (f *FormPanel) FieldPostFilterFn(post PostFieldFilterFn) *FormPanel {
	f.FieldList[f.curFieldListIndex].PostFilterFn = post
	return f
}

func (f *FormPanel) FieldLimit(limit int) *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddLimit(limit)
	return f
}

func (f *FormPanel) FieldTrimSpace() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddTrimSpace()
	return f
}

func (f *FormPanel) FieldSubstr(start int, end int) *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddSubstr(start, end)
	return f
}

func (f *FormPanel) FieldToTitle() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddToTitle()
	return f
}

func (f *FormPanel) FieldToUpper() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddToUpper()
	return f
}

func (f *FormPanel) FieldToLower() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddToLower()
	return f
}

// FormPanel attribute setting functions
// ====================================================

func (f *FormPanel) SetTitle(title string) *FormPanel {
	f.Title = title
	return f
}

func (f *FormPanel) SetTabGroups(groups TabGroups) *FormPanel {
	f.TabGroups = groups
	return f
}

func (f *FormPanel) SetTabHeaders(headers ...string) *FormPanel {
	f.TabHeaders = headers
	return f
}

func (f *FormPanel) SetDescription(desc string) *FormPanel {
	f.Description = desc
	return f
}

func (f *FormPanel) SetHeaderHtml(header template.HTML) *FormPanel {
	f.HeaderHtml = header
	return f
}

func (f *FormPanel) SetFooterHtml(footer template.HTML) *FormPanel {
	f.FooterHtml = footer
	return f
}

func (f *FormPanel) SetPostValidator(va FormValidator) *FormPanel {
	f.Validator = va
	return f
}

func (f *FormPanel) SetPostHook(po FormPostHookFn) *FormPanel {
	f.PostHook = po
	return f
}

type FormValidator func(values form2.Values) error

type FormPostHookFn func(values form2.Values)

type FormFields []FormField

func (f FormFields) Copy() FormFields {
	formList := make(FormFields, len(f))
	copy(formList, f)
	for i := 0; i < len(formList); i++ {
		formList[i].Options = make([]map[string]string, len(f[i].Options))
		for j := 0; j < len(f[i].Options); j++ {
			formList[i].Options[j] = modules.CopyMap(f[i].Options[j])
		}
	}
	return formList
}

func (f FormFields) FindByFieldName(field string) FormField {
	for i := 0; i < len(f); i++ {
		if f[i].Field == field {
			return f[i]
		}
	}
	return FormField{}
}
