package adminlte

import (
	"html/template"
)

type InfoBoxAttribute struct {
	Name  string
	Title string
	Value string
	Url   string
	Color string
}

func (AdminlteComponents) InfoBox() *InfoBoxAttribute {
	return &InfoBoxAttribute{
		"infobox",
		"标题",
		"值",
		"/",
		"aqua",
	}
}

func (compo *InfoBoxAttribute) SetTitle(value string) *InfoBoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *InfoBoxAttribute) SetValue(value string) *InfoBoxAttribute {
	(*compo).Value = value
	return compo
}

func (compo *InfoBoxAttribute) SetUrl(value string) *InfoBoxAttribute {
	(*compo).Url = value
	return compo
}

func (compo *InfoBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "infobox")
}
