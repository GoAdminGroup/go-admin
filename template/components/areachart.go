package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type AreaChartAttribute struct {
	Name   string
	Title  string
	Data   string
	ID     string
	Height int
	types.Attribute
}

func (compo *AreaChartAttribute) SetID(value string) types.AreaChartAttribute {
	compo.ID = value
	return compo
}

func (compo *AreaChartAttribute) SetTitle(value string) types.AreaChartAttribute {
	compo.Title = value
	return compo
}

func (compo *AreaChartAttribute) SetHeight(value int) types.AreaChartAttribute {
	compo.Height = value
	return compo
}

func (compo *AreaChartAttribute) SetData(value string) types.AreaChartAttribute {
	compo.Data = value
	return compo
}

func (compo *AreaChartAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "area-chart")
}
