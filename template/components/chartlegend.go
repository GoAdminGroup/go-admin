package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type ChartLegendAttribute struct {
	Name string
	Data []map[string]string
	types.Attribute
}

func (compo *ChartLegendAttribute) SetData(value []map[string]string) types.ChartLegendAttribute {
	compo.Data = value
	return compo
}

func (compo *ChartLegendAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "chart-legend")
}
