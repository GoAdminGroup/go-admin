package components

import (
	"html/template"
	"goAdmin/template/types"
)

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
	Src    string
}

func (compo *ImgAttribute) SetWidth(value string) types.ImgAttribute {
	(*compo).Witdh = value
	return compo
}

func (compo *ImgAttribute) SetHeight(value string) types.ImgAttribute {
	(*compo).Height = value
	return compo
}

func (compo *ImgAttribute) SetSrc(value string) types.ImgAttribute {
	(*compo).Src = value
	return compo
}

func (compo *ImgAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "image")
}
