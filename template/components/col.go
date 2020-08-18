package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type ColAttribute struct {
	Name    string
	Content template.HTML
	Size    string
	types.Attribute
}

func (compo *ColAttribute) SetContent(value template.HTML) types.ColAttribute {
	compo.Content = value
	return compo
}

func (compo *ColAttribute) AddContent(value template.HTML) types.ColAttribute {
	compo.Content += value
	return compo
}

func (compo *ColAttribute) SetSize(value types.S) types.ColAttribute {
	compo.Size = ""
	for key, size := range value {
		compo.Size += "col-" + key + "-" + size + " "
	}
	return compo
}

func (compo *ColAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "col")
}
