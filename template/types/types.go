// Copyright 2019 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
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

// Form is the form field with different options.
type Form struct {
	Field                  string
	TypeName               db.DatabaseType
	Head                   string
	Default                string
	Editable               bool
	FormType               form.Type
	Value                  string
	Options                []map[string]string
	DefaultOptionDelimiter string
	FilterFn               FieldFilterFn
	PostFilterFn           PostFieldFilterFn
	ProcessFn              ProcessFn
}

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
	FieldList    []Field
	Table        string
	Title        string
	Sort         Sort
	Group        [][]string
	GroupHeaders []string
	Description  string
	HeaderHtml   template.HTML
	FooterHtml   template.HTML
}

type Sort uint8

const (
	SortDesc Sort = iota
	SortAsc
)

// FormPanel
type FormPanel struct {
	FormList     FormList
	Group        [][]string
	GroupHeaders []string
	Table        string
	Title        string
	Description  string
	HeaderHtml   template.HTML
	FooterHtml   template.HTML
}

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
