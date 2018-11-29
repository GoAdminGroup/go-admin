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
	Name    string
	Content string
}

var Default = &Attribute{
	"Default",
	"",
}

func (compo *Attribute) GetContent(value interface{}) string {
	return compo.Content + value.(string)
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
type FormStruct struct {
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
type FieldStruct struct {
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
	JoinTable  Join
}

func (field *FieldStruct) SetHead(head string) *FieldStruct {
	field.Head = head
	return field
}

func (field *FieldStruct) SetTypeName(typeName string) *FieldStruct {
	field.TypeName = typeName
	return field
}

func (field *FieldStruct) SetField(fieldName string) *FieldStruct {
	field.Field = fieldName
	return field
}

// 展示面板
type InfoPanel struct {
	FieldList   []FieldStruct
	Table       string
	Title       string
	Description string
}

// 表单面板
type FormPanel struct {
	FormList    []FormStruct
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
