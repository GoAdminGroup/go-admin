package adminlte

import "html/template"

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
	Src    string
}

func (AdminlteComponents) Image() *ImgAttribute {
	return &ImgAttribute{
		"image",
		"50",
		"50",
		"",
	}
}

func (compo *ImgAttribute) SetWidth(value string) *ImgAttribute {
	(*compo).Witdh = value
	return compo
}

func (compo *ImgAttribute) SetHeight(value string) *ImgAttribute {
	(*compo).Height = value
	return compo
}

func (compo *ImgAttribute) SetSrc(value string) *ImgAttribute {
	(*compo).Src = value
	return compo
}

func (compo *ImgAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "image")
}
