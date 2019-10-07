package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type TableAttribute struct {
	Name       string
	Thead      []map[string]string
	InfoList   []map[string]template.HTML
	Type       string
	PrimaryKey string
	EditUrl    string
	DeleteUrl  string
	IsTab      bool
	ExportUrl  string
	types.Attribute
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
	return ComposeHtml(compo.TemplateList, *compo, "table")
}

type DataTableAttribute struct {
	TableAttribute
	EditUrl    string
	NewUrl     string
	DeleteUrl  string
	PrimaryKey string
	IsTab      bool
	ExportUrl  string
	InfoUrl    string
	FilterUrl  string
	Filters    []map[string]string
	types.Attribute
}

func (compo *DataTableAttribute) GetDataTableHeader() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "table/box-header")
}

func (compo *DataTableAttribute) SetThead(value []map[string]string) types.DataTableAttribute {
	compo.Thead = value
	return compo
}

func (compo *DataTableAttribute) SetIsTab(value bool) types.DataTableAttribute {
	compo.IsTab = value
	return compo
}

func (compo *DataTableAttribute) SetInfoUrl(value string) types.DataTableAttribute {
	compo.InfoUrl = value
	return compo
}

func (compo *DataTableAttribute) SetExportUrl(value string) types.DataTableAttribute {
	compo.ExportUrl = value
	return compo
}

func (compo *DataTableAttribute) SetFilterUrl(value string) types.DataTableAttribute {
	compo.FilterUrl = value
	return compo
}

func (compo *DataTableAttribute) SetPrimaryKey(value string) types.DataTableAttribute {
	compo.PrimaryKey = value
	return compo
}

func (compo *DataTableAttribute) SetFilters(value []map[string]string) types.DataTableAttribute {
	compo.Filters = value
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
	return ComposeHtml(compo.TemplateList, *compo, "table")
}
