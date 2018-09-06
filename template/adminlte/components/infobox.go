package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type SmallBoxAttribute struct {
	Name  string
	Title string
	Value string
	Url   string
	Color string
}

func (compo *SmallBoxAttribute) SetTitle(value string) types.SmallBoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *SmallBoxAttribute) SetValue(value string) types.SmallBoxAttribute {
	(*compo).Value = value
	return compo
}

func (compo *SmallBoxAttribute) SetUrl(value string) types.SmallBoxAttribute {
	(*compo).Url = value
	return compo
}

func (compo *SmallBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "infobox")
}
