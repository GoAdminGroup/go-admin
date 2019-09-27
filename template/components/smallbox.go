package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type SmallBoxAttribute struct {
	Name  string
	Title template.HTML
	Value template.HTML
	Url   string
	Color string
	types.Attribute
}

func (compo *SmallBoxAttribute) SetTitle(value template.HTML) types.SmallBoxAttribute {
	compo.Title = value
	return compo
}

func (compo *SmallBoxAttribute) SetValue(value template.HTML) types.SmallBoxAttribute {
	compo.Value = value
	return compo
}

func (compo *SmallBoxAttribute) SetUrl(value string) types.SmallBoxAttribute {
	compo.Url = value
	return compo
}

func (compo *SmallBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "smallbox")
}
