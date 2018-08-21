package adminlte

import (
	"html/template"
)

type TableAttribute struct {
	Name     string
	Thead    []map[string]string
	InfoList []map[string]string
	EditUrl  string
}

func (AdminlteComponents) Table() *TableAttribute {
	return &TableAttribute{
		"table",
		[]map[string]string{},
		[]map[string]string{},
		"",
	}
}

func (compo *TableAttribute) SetThead(value []map[string]string) *TableAttribute {
	(*compo).Thead = value
	return compo
}

func (compo *TableAttribute) SetInfoList(value []map[string]string) *TableAttribute {
	(*compo).InfoList = value
	return compo
}

func (compo *TableAttribute) SetUrl(value string) *TableAttribute {
	(*compo).EditUrl = value
	return compo
}

func (compo *TableAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "table", "paninator")
}
