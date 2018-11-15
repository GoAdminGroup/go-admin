package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type PopupAttribute struct {
	Name   string
	ID     string
	Height int
	Data   string
	Title  string
}

func (compo *PopupAttribute) SetID(value string) types.PopupAttribute {
	compo.ID = value
	return compo
}

func (compo *PopupAttribute) SetTitle(value string) types.PopupAttribute {
	compo.Title = value
	return compo
}

func (compo *PopupAttribute) SetData(value string) types.PopupAttribute {
	compo.Data = value
	return compo
}

func (compo *PopupAttribute) SetHeight(value int) types.PopupAttribute {
	compo.Height = value
	return compo
}

func (compo *PopupAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "popup")
}
