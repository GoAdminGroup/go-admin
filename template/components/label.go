package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type LabelAttribute struct {
	Name    string
	Color   string
	Content template.HTML
	types.Attribute
}

func (compo *LabelAttribute) SetContent(value template.HTML) types.LabelAttribute {
	compo.Content = value
	return compo
}

func (compo *LabelAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "label")
}
