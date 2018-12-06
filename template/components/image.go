package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
	Src    string
	types.Attribute
}

func (compo *ImgAttribute) SetWidth(value string) types.ImgAttribute {
	compo.Witdh = value
	return compo
}

func (compo *ImgAttribute) SetHeight(value string) types.ImgAttribute {
	compo.Height = value
	return compo
}

func (compo *ImgAttribute) SetSrc(value string) types.ImgAttribute {
	compo.Src = value
	return compo
}

func (compo *ImgAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "image")
}
