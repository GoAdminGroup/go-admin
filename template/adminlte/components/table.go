package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type TableAttribute struct {
	Name      string
	Thead     []map[string]string
	InfoList  []map[string]template.HTML
	Type      string
	EditUrl   string
	DeleteUrl string
}

func (compo *TableAttribute) SetThead(value []map[string]string) types.TableAttribute {
	compo.Thead = value
	return compo
}

func (compo *TableAttribute) SetInfoList(value []map[string]template.HTML) types.TableAttribute {
	compo.InfoList = value
	return compo
}

func (compo *TableAttribute) SetType(value string) types.TableAttribute {
	compo.Type = value
	return compo
}

func (compo *TableAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "table")
}

type DataTableAttribute struct {
	TableAttribute
	EditUrl   string
	NewUrl    string
	DeleteUrl string
}

func (compo *DataTableAttribute) GetDataTableHeader() template.HTML {
	return ComposeHtml(*compo, "table/box-header")
}

func (compo *DataTableAttribute) SetThead(value []map[string]string) types.DataTableAttribute {
	compo.Thead = value
	return compo
}

func (compo *DataTableAttribute) SetInfoList(value []map[string]template.HTML) types.DataTableAttribute {
	compo.InfoList = value
	return compo
}

func (compo *DataTableAttribute) SetEditUrl(value string) types.DataTableAttribute {
	compo.EditUrl = value
	return compo
}

func (compo *DataTableAttribute) SetDeleteUrl(value string) types.DataTableAttribute {
	compo.DeleteUrl = value
	return compo
}

func (compo *DataTableAttribute) SetNewUrl(value string) types.DataTableAttribute {
	compo.NewUrl = value
	return compo
}

func (compo *DataTableAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "table")
}
