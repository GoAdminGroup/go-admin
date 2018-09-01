package components

import (
	"html/template"
)

type LabelAttribute struct {
	Name    string
	Color   string
	Content string
}

func Label() *LabelAttribute {
	return &LabelAttribute{
		"label",
		"success",
		"",
	}
}

func (compo *LabelAttribute) SetContent(value string) *LabelAttribute {
	(*compo).Content = value
	return compo
}

func (compo *LabelAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "label")
}
