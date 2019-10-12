package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type PopupAttribute struct {
	Name   string
	ID     string
	Body   template.HTML
	Footer template.HTML
	Title  template.HTML
	Size   string
	types.Attribute
}

func (compo *PopupAttribute) SetID(value string) types.PopupAttribute {
	compo.ID = value
	return compo
}

func (compo *PopupAttribute) SetTitle(value template.HTML) types.PopupAttribute {
	compo.Title = value
	return compo
}

func (compo *PopupAttribute) SetFooter(value template.HTML) types.PopupAttribute {
	compo.Footer = value
	return compo
}

func (compo *PopupAttribute) SetBody(value template.HTML) types.PopupAttribute {
	compo.Body = value
	return compo
}

func (compo *PopupAttribute) SetSize(value string) types.PopupAttribute {
	compo.Size = value
	return compo
}

func (compo *PopupAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "popup")
}
