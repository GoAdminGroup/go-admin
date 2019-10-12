package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type LineChartAttribute struct {
	Name   string
	Title  template.HTML
	Data   string
	ID     string
	Height int
	types.Attribute
}

func (compo *LineChartAttribute) SetID(value string) types.LineChartAttribute {
	compo.ID = value
	return compo
}

func (compo *LineChartAttribute) SetTitle(value template.HTML) types.LineChartAttribute {
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
	return ComposeHtml(compo.TemplateList, *compo, "line-chart")
}
