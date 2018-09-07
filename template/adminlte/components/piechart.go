package components

import (
	"goAdmin/template/types"
	"html/template"
)

type PieChartAttribute struct {
	Name   string
	ID     string
	Height int
	Data   string
	Prefix string
}

func (compo *PieChartAttribute) SetID(value string) types.PieChartAttribute {
	(*compo).ID = value
	return compo
}

func (compo *PieChartAttribute) SetData(value string) types.PieChartAttribute {
	(*compo).Data = value
	return compo
}

func (compo *PieChartAttribute) SetPrefix(value string) types.PieChartAttribute {
	(*compo).Prefix = value
	return compo
}

func (compo *PieChartAttribute) SetHeight(value int) types.PieChartAttribute {
	(*compo).Height = value
	return compo
}

func (compo *PieChartAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "pie-chart")
}