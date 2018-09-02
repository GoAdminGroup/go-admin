package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type InfoBoxAttribute struct {
	Name  string
	Title string
	Value string
	Url   string
	Color string
}

func (*AdminlteStruct) InfoBox() types.InfoBoxAttribute {
	return &InfoBoxAttribute{
		"infobox",
		"标题",
		"值",
		"/",
		"aqua",
	}
}

func (compo *InfoBoxAttribute) SetTitle(value string) types.InfoBoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *InfoBoxAttribute) SetValue(value string) types.InfoBoxAttribute {
	(*compo).Value = value
	return compo
}

func (compo *InfoBoxAttribute) SetUrl(value string) types.InfoBoxAttribute {
	(*compo).Url = value
	return compo
}

func (compo *InfoBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "infobox")
}
