package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type LabelAttribute struct {
	Name    string
	Color   string
	Content string
}

func (compo *LabelAttribute) SetContent(value string) types.LabelAttribute {
	compo.Content = value
	return compo
}

func (compo *LabelAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "label")
}
