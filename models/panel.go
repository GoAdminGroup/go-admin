package models

// 表单列
type FormStruct struct {
	Field    string
	TypeName string
	Head     string
	Default  string
	Editable bool
	FormType string
	Value    string
}

type FieldValueFun func(value string) string

// 展示列
type FieldStruct struct {
	ExcuFun  FieldValueFun
	Field    string
	TypeName string
	Head     string
}

func (field *FieldStruct) SetHead(head string) *FieldStruct {
	(*field).Head = head
	return field
}

func (field *FieldStruct) SetTypeName(typeName string) *FieldStruct {
	(*field).TypeName = typeName
	return field
}

func (field *FieldStruct) SetField(fieldName string) *FieldStruct {
	(*field).Field = fieldName
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

type GlobalTable struct {
	Info InfoPanel
	Form FormPanel
}

type FunCaller struct{}
