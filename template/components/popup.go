package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type PopupAttribute struct {
	Name       string
	ID         string
	Body       template.HTML
	Footer     template.HTML
	FooterHTML template.HTML
	Title      template.HTML
	Size       string
	HideFooter bool
	Height     string
	Width      string
	Draggable  bool
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

func (compo *PopupAttribute) SetFooterHTML(value template.HTML) types.PopupAttribute {
	compo.FooterHTML = value
	return compo
}

func (compo *PopupAttribute) SetWidth(width string) types.PopupAttribute {
	compo.Width = width
	return compo
}

func (compo *PopupAttribute) SetHeight(height string) types.PopupAttribute {
	compo.Height = height
	return compo
}

func (compo *PopupAttribute) SetDraggable() types.PopupAttribute {
	compo.Draggable = true
	return compo
}

func (compo *PopupAttribute) SetHideFooter() types.PopupAttribute {
	compo.HideFooter = true
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
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "popup")
}
