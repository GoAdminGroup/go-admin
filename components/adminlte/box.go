package adminlte

import (
	"html/template"
)

type BoxAttribute struct {
	Name       string
	Header     template.HTML
	Body       template.HTML
	Footer     template.HTML
	Title      template.HTML
	HeadBorder string
}

func (AdminlteComponents) Box() *BoxAttribute {
	return &BoxAttribute{
		"box",
		template.HTML(""),
		template.HTML(""),
		template.HTML(""),
		"",
		"",
	}
}

func (compo *BoxAttribute) SetHeader(value template.HTML) *BoxAttribute {
	(*compo).Header = value
	return compo
}

func (compo *BoxAttribute) SetBody(value template.HTML) *BoxAttribute {
	(*compo).Body = value
	return compo
}

func (compo *BoxAttribute) SetFooter(value template.HTML) *BoxAttribute {
	(*compo).Footer = value
	return compo
}

func (compo *BoxAttribute) SetTitle(value template.HTML) *BoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *BoxAttribute) WithHeadBorder(has bool) *BoxAttribute {
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
