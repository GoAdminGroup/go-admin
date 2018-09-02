package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type BoxAttribute struct {
	Name       string
	Header     template.HTML
	Body       template.HTML
	Footer     template.HTML
	Title      template.HTML
	HeadBorder string
}

func (*AdminlteStruct) Box() types.BoxAttribute {
	return &BoxAttribute{
		"box",
		template.HTML(""),
		template.HTML(""),
		template.HTML(""),
		"",
		"",
	}
}

func (compo *BoxAttribute) SetHeader(value template.HTML) types.BoxAttribute {
	(*compo).Header = value
	return compo
}

func (compo *BoxAttribute) SetBody(value template.HTML) types.BoxAttribute {
	(*compo).Body = value
	return compo
}

func (compo *BoxAttribute) SetFooter(value template.HTML) types.BoxAttribute {
	(*compo).Footer = value
	return compo
}

func (compo *BoxAttribute) SetTitle(value template.HTML) types.BoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *BoxAttribute) WithHeadBorder(has bool) types.BoxAttribute {
	if has {
		(*compo).HeadBorder = "with-border"
	} else {
		(*compo).HeadBorder = ""
	}
	return compo
}

func (compo *BoxAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "box")
}
