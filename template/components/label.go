package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type LabelAttribute struct {
	Name    string
	Color   template.HTML
	Type    string
	Content template.HTML
	types.Attribute
}

func (compo *LabelAttribute) SetType(value string) types.LabelAttribute {
	compo.Type = value
	return compo
}

func (compo *LabelAttribute) SetColor(value template.HTML) types.LabelAttribute {
	compo.Color = value
	return compo
}

func (compo *LabelAttribute) SetContent(value template.HTML) types.LabelAttribute {
	compo.Content = value
	return compo
}

func (compo *LabelAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "label")
}
