package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type LabelAttribute struct {
	Name    string
	Color   string
	Content string
	types.Attribute
}

func (compo *LabelAttribute) SetContent(value string) types.LabelAttribute {
	compo.Content = value
	return compo
}

func (compo *LabelAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "label")
}
