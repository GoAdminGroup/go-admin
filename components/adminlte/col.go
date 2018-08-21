package adminlte

import (
	"html/template"
)

type ColAttribute struct {
	Name    string
	Width   string
	Content template.HTML
	Type    string
}

func (AdminlteComponents) Col() *ColAttribute {
	return &ColAttribute{
		"col",
		"2",
		"",
		"md",
	}
}

func (compo *ColAttribute) SetWidth(value string) *ColAttribute {
	(*compo).Width = value
	return compo
}

func (compo *ColAttribute) SetContent(value string) *ColAttribute {
	(*compo).Content = template.HTML(value)
	return compo
}

func (compo *ColAttribute) SetType(value string) *ColAttribute {
	(*compo).Type = value
	return compo
}

func (compo *ColAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "col")
}