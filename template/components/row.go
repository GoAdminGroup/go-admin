package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type RowAttribute struct {
	Name    string
	Content template.HTML
	types.Attribute
}

func (compo *RowAttribute) SetContent(value template.HTML) types.RowAttribute {
	compo.Content = value
	return compo
}

func (compo *RowAttribute) AddContent(value template.HTML) types.RowAttribute {
	compo.Content += value
	return compo
}

func (compo *RowAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "row")
}
