package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type BarChartAttribute struct {
	Name  string
	Title string
	Data  string
	ID    string
	Width int
	types.Attribute
}

func (compo *BarChartAttribute) SetID(value string) types.BarChartAttribute {
	compo.ID = value
	return compo
}

func (compo *BarChartAttribute) SetTitle(value string) types.BarChartAttribute {
	compo.Title = value
	return compo
}

func (compo *BarChartAttribute) SetWidth(value int) types.BarChartAttribute {
	compo.Width = value
	return compo
}

func (compo *BarChartAttribute) SetData(value string) types.BarChartAttribute {
	compo.Data = value
	return compo
}

func (compo *BarChartAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "bar-chart")
}
