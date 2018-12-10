package types

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"html/template"
)

type Component interface {
	GetContent(interface{}) string
}

type Attribute struct {
	TemplateList map[string]string
}

type Page struct {
	User          auth.User
	Menu          menu.Menu
	Panel         Panel
	System        SystemInfo
	AssertRootUrl string
	Title         string
	Logo          template.HTML
	MiniLogo      template.HTML
}

type SystemInfo struct {
	Version string
}

// 表单列
type Form struct {
	Field    string
	TypeName string
	Head     string
	Default  string
	Editable bool
	FormType string
	Value    string
	Options  []map[string]string
	ExcuFun  FieldValueFun
}

type RowModel struct {
	ID    int64
	Value string
}

// 数据过滤函数
type FieldValueFun func(value RowModel) interface{}

// 展示列
type Field struct {
	ExcuFun   FieldValueFun
	Field     string
	TypeName  string
	Head      string
	JoinTable []Join
	Sortable  bool
	Filter    bool
}

type Join struct {
	Table      string
	Field      string
	TableField string
	HasChild   bool
	JoinTable  *Join
}

func (field *Field) SetHead(head string) *Field {
	field.Head = head
	return field
}

func (field *Field) SetTypeName(typeName string) *Field {
	field.TypeName = typeName
	return field
}

func (field *Field) SetField(fieldName string) *Field {
	field.Field = fieldName
	return field
}

// 展示面板
type InfoPanel struct {
	FieldList   []Field
	Table       string
	Title       string
	Description string
}

// 表单面板
type FormPanel struct {
	FormList    []Form
	Table       string
	Title       string
	Description string
}

type Panel struct {
	Content     template.HTML
	Title       string
	Description string
	Url         string
}

type GetPanel func() Panel
