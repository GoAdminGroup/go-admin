package components

import (
	"goAdmin/template/types"
	"html/template"
)

type ChartLegendAttribute struct {
	Name   string
	Data   []map[string]string
}

func (compo *ChartLegendAttribute) SetData(value []map[string]string) types.ChartLegendAttribute {
	(*compo).Data = value
	return compo
}

func (compo *ChartLegendAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "chart-legend")
}