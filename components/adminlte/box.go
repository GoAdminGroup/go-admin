package adminlte

import (
	"html/template"
)

type BoxAttribute struct {
	Name  string
	Title string
	Value string
	Url   string
	Color string
}

func (AdminlteComponents) Box() *BoxAttribute {
	return &BoxAttribute{
		"box",
		"标题",
		"值",
		"/",
		"aqua",
	}
}

func (compo *BoxAttribute) SetTitle(value string) *BoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *BoxAttribute) SetValue(value string) *BoxAttribute {
	(*compo).Value = value
	return compo
}

func (compo *BoxAttribute) SetUrl(value string) *BoxAttribute {
	(*compo).Url = value
	return compo
}

func (compo *BoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "box")
}
