package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type ColAttribute struct {
	Name    string
	Width   string
	Content template.HTML
	Type    string
}

func (*AdminlteStruct) Col() types.ColAttribute {
	return &ColAttribute{
		"col",
		"2",
		"",
		"md",
	}
}

func (compo *ColAttribute) SetWidth(value string) types.ColAttribute {
	(*compo).Width = value
	return compo
}

func (compo *ColAttribute) SetContent(value template.HTML) types.ColAttribute {
	(*compo).Content = value
	return compo
}

func (compo *ColAttribute) SetType(value string) types.ColAttribute {
	(*compo).Type = value
	return compo
}

func (compo *ColAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "col")
}