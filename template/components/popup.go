package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type PopupAttribute struct {
	Name   string
	ID     string
	Body   template.HTML
	Footer string
	Title  string
	types.Attribute
}

func (compo *PopupAttribute) SetID(value string) types.PopupAttribute {
	compo.ID = value
	return compo
}

func (compo *PopupAttribute) SetTitle(value string) types.PopupAttribute {
	compo.Title = value
	return compo
}

func (compo *PopupAttribute) SetFooter(value string) types.PopupAttribute {
	compo.Footer = value
	return compo
}

func (compo *PopupAttribute) SetBody(value template.HTML) types.PopupAttribute {
	compo.Body = value
	return compo
}

func (compo *PopupAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "popup")
}
