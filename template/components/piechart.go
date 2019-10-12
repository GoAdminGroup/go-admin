package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type PieChartAttribute struct {
	Name   string
	ID     string
	Height int
	Data   string
	Title  template.HTML
	types.Attribute
}

func (compo *PieChartAttribute) SetID(value string) types.PieChartAttribute {
	compo.ID = value
	return compo
}

func (compo *PieChartAttribute) SetTitle(value template.HTML) types.PieChartAttribute {
	compo.Title = value
	return compo
}

func (compo *PieChartAttribute) SetData(value string) types.PieChartAttribute {
	compo.Data = value
	return compo
}

func (compo *PieChartAttribute) SetHeight(value int) types.PieChartAttribute {
	compo.Height = value
	return compo
}

func (compo *PieChartAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "pie-chart")
}
