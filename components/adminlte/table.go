package adminlte

import (
	"html/template"
)

type TableAttribute struct {
	Name     string
	Thead    []map[string]string
	InfoList []map[string]template.HTML
	Type     string
}

func (AdminlteComponents) Table() *TableAttribute {
	return &TableAttribute{
		Name:     "table",
		Thead:    []map[string]string{},
		InfoList: []map[string]template.HTML{},
		Type:     "normal",
	}
}

func (compo *TableAttribute) SetThead(value []map[string]string) *TableAttribute {
	(*compo).Thead = value
	return compo
}

func (compo *TableAttribute) SetInfoList(value []map[string]template.HTML) *TableAttribute {
	(*compo).InfoList = value
	return compo
}

func (compo *TableAttribute) SetType(value string) *TableAttribute {
	(*compo).Type = value
	return compo
}

func (compo *TableAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "table")
}

type DataTableAttribute struct {
	TableAttribute
	EditUrl string
	NewUrl  string
}

func (admin AdminlteComponents) DataTable() *DataTableAttribute {
	return &DataTableAttribute{
		TableAttribute: *admin.Table().SetType("data-table"),
		EditUrl:        "",
		NewUrl:         "",
	}
}

func (compo *DataTableAttribute) GetDataTableHeader() template.HTML {
	return ComposeHtml(*compo, "table/box-header")
}

func (compo *DataTableAttribute) SetThead(value []map[string]string) *DataTableAttribute {
	(*compo).Thead = value
	return compo
}

func (compo *DataTableAttribute) SetInfoList(value []map[string]template.HTML) *DataTableAttribute {
	(*compo).InfoList = value
	return compo
}

func (compo *DataTableAttribute) SetEditUrl(value string) *DataTableAttribute {
	(*compo).EditUrl = value
	return compo
}

func (compo *DataTableAttribute) SetNewUrl(value string) *DataTableAttribute {
	(*compo).NewUrl = value
	return compo
}

func (compo *DataTableAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "table")
}
