package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type LineChartAttribute struct {
	Name   string
	Title  string
	Data   string
	ID     string
	Height int
}

func (compo *LineChartAttribute) SetID(value string) types.LineChartAttribute {
	compo.ID = value
	return compo
}

func (compo *LineChartAttribute) SetTitle(value string) types.LineChartAttribute {
	compo.Title = value
	return compo
}

func (compo *LineChartAttribute) SetHeight(value int) types.LineChartAttribute {
	compo.Height = value
	return compo
}

func (compo *LineChartAttribute) SetData(value string) types.LineChartAttribute {
	compo.Data = value
	return compo
}

func (compo *LineChartAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "line-chart")
}
