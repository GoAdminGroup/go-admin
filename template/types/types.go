// Copyright 2019 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/modules/system"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	form2 "github.com/chenhg5/go-admin/plugins/admin/modules/form"
	"github.com/chenhg5/go-admin/template/types/form"
	"html/template"
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

	// CdnUrl is the cdn link of assets
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
		CdnUrl:         cfg.CdnUrl,
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

type GetPanel func() Panel

// RowModel contains ID and value of the single query result.
type RowModel struct {
	ID    string
	Value string
	Row   map[string]interface{}
}

// PostRowModel contains ID and value of the single query result.
type PostRowModel struct {
	ID    string
	Value RowModelValue
	Row   map[string]interface{}
}

type RowModelValue []string

func (r RowModelValue) Value() string {
	return r.First()
}

func (r RowModelValue) First() string {
	return r[0]
}

// FieldFilterFn is filter function of data.
type FieldFilterFn func(value RowModel) interface{}

// ProcessFn process the data and store into the database.
type ProcessFn func(value PostRowModel)

// PostFieldFilterFn is filter function of data.
type PostFieldFilterFn func(value PostRowModel) string

// Field is the table field.
type Field struct {
	FilterFn   FieldFilterFn
	Field      string
	TypeName   db.DatabaseType
	Head       string
	Width      int
	Join       Join
	Sortable   bool
	Fixed      bool
	Filterable bool
	Hide       bool
}

type Join struct {
	Table     string
	Field     string
	JoinField string
	HasChild  bool
	JoinTable *Join
}

func (j Join) Valid() bool {
	return j.Table != "" && j.Field != "" && j.JoinField != ""
}

// InfoPanel
type InfoPanel struct {
	FieldList     []Field
	curFieldIndex int
	Table         string
	Title         string
	Sort          Sort
	Group         [][]string
	GroupHeaders  []string
	Description   string
	Action        template.HTML
	HeaderHtml    template.HTML
	FooterHtml    template.HTML
}

func NewInfoPanel() *InfoPanel {
	return &InfoPanel{curFieldIndex: -1}
}

func (i *InfoPanel) AddField(head, field string, typeName db.DatabaseType) *InfoPanel {
	i.FieldList = append(i.FieldList, Field{
		Head:     head,
		Field:    field,
		TypeName: typeName,
		Sortable: false,
		FilterFn: func(model RowModel) interface{} {
			return model.Value
		},
	})
	i.curFieldIndex++
	return i
}

func (i *InfoPanel) FieldFilterFn(filter FieldFilterFn) *InfoPanel {
	i.FieldList[i.curFieldIndex].FilterFn = filter
	return i
}

func (i *InfoPanel) FieldWidth(width int) *InfoPanel {
	i.FieldList[i.curFieldIndex].Width = width
	return i
}

func (i *InfoPanel) FieldSortable(sort bool) *InfoPanel {
	i.FieldList[i.curFieldIndex].Sortable = sort
	return i
}

func (i *InfoPanel) FieldFixed(fixed bool) *InfoPanel {
	i.FieldList[i.curFieldIndex].Fixed = fixed
	return i
}

func (i *InfoPanel) FieldFilterable(filter bool) *InfoPanel {
	i.FieldList[i.curFieldIndex].Filterable = filter
	return i
}

func (i *InfoPanel) FieldHide(hide bool) *InfoPanel {
	i.FieldList[i.curFieldIndex].Hide = hide
	return i
}

func (i *InfoPanel) FieldJoin(join Join) *InfoPanel {
	i.FieldList[i.curFieldIndex].Join = join
	return i
}

func (i *InfoPanel) SetTable(table string) *InfoPanel {
	i.Table = table
	return i
}

func (i *InfoPanel) SetTitle(title string) *InfoPanel {
	i.Title = title
	return i
}

func (i *InfoPanel) SetGroup(group [][]string) *InfoPanel {
	i.Group = group
	return i
}

func (i *InfoPanel) SetGroupHeaders(headers ...string) *InfoPanel {
	i.GroupHeaders = headers
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

type Sort uint8

const (
	SortDesc Sort = iota
	SortAsc
)

// Form is the form field with different options.
type Form struct {
	Field                  string
	TypeName               db.DatabaseType
	Head                   string
	Default                string
	Editable               bool
	NotAllowAdd            bool
	FormType               form.Type
	Value                  string
	Options                []map[string]string
	DefaultOptionDelimiter string
	FilterFn               FieldFilterFn
	PostFilterFn           PostFieldFilterFn
	ProcessFn              ProcessFn
}

// FormPanel
type FormPanel struct {
	FormList      FormList
	curFieldIndex int
	Group         [][]string
	GroupHeaders  []string
	Table         string
	Title         string
	Description   string
	PostValidator PostValidator
	PostHook      PostHookFn
	HeaderHtml    template.HTML
	FooterHtml    template.HTML
}

func NewFormPanel() *FormPanel {
	return &FormPanel{curFieldIndex: -1}
}

func (f *FormPanel) AddField(head, field string, filedType db.DatabaseType, formType form.Type) *FormPanel {
	f.FormList = append(f.FormList, Form{
		Head:     head,
		Field:    field,
		TypeName: filedType,
		Editable: true,
		FormType: formType,
		FilterFn: func(model RowModel) interface{} {
			return model.Value
		},
	})
	f.curFieldIndex++
	return f
}

func (f *FormPanel) FieldFilterFn(filter FieldFilterFn) *FormPanel {
	f.FormList[f.curFieldIndex].FilterFn = filter
	return f
}

func (f *FormPanel) SetTable(table string) *FormPanel {
	f.Table = table
	return f
}

func (f *FormPanel) FieldEditable(edit bool) *FormPanel {
	f.FormList[f.curFieldIndex].Editable = edit
	return f
}

func (f *FormPanel) FieldNotAllowAdd(add bool) *FormPanel {
	f.FormList[f.curFieldIndex].NotAllowAdd = add
	return f
}

func (f *FormPanel) FieldFormType(formType form.Type) *FormPanel {
	f.FormList[f.curFieldIndex].FormType = formType
	return f
}

func (f *FormPanel) FieldValue(value string) *FormPanel {
	f.FormList[f.curFieldIndex].Value = value
	return f
}

func (f *FormPanel) FieldOptions(options []map[string]string) *FormPanel {
	f.FormList[f.curFieldIndex].Options = options
	return f
}

func (f *FormPanel) FieldDefaultOptionDelimiter(delimiter string) *FormPanel {
	f.FormList[f.curFieldIndex].DefaultOptionDelimiter = delimiter
	return f
}

func (f *FormPanel) FieldPostFilterFn(post PostFieldFilterFn) *FormPanel {
	f.FormList[f.curFieldIndex].PostFilterFn = post
	return f
}

func (f *FormPanel) FieldProcessFn(a ProcessFn) *FormPanel {
	f.FormList[f.curFieldIndex].ProcessFn = a
	return f
}

func (f *FormPanel) SetTitle(title string) *FormPanel {
	f.Title = title
	return f
}

func (f *FormPanel) SetGroup(group [][]string) *FormPanel {
	f.Group = group
	return f
}

func (f *FormPanel) SetGroupHeaders(headers ...string) *FormPanel {
	f.GroupHeaders = headers
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

func (f *FormPanel) SetPostValidator(va PostValidator) *FormPanel {
	f.PostValidator = va
	return f
}

func (f *FormPanel) SetPostHook(po PostHookFn) *FormPanel {
	f.PostHook = po
	return f
}

type PostValidator func(values form2.Values) error

type PostHookFn func(values form2.Values)

type FormList []Form

func (f FormList) Copy() FormList {
	formList := make(FormList, len(f))
	copy(formList, f)
	for i := 0; i < len(formList); i++ {
		formList[i].Options = make([]map[string]string, len(f[i].Options))
		for j := 0; j < len(f[i].Options); j++ {
			formList[i].Options[j] = modules.CopyMap(f[i].Options[j])
		}
	}
	return formList
}

func (f FormList) FindByField(field string) Form {
	for i := 0; i < len(f); i++ {
		if f[i].Field == field {
			return f[i]
		}
	}
	return Form{}
}
