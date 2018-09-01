package components

import (
	"html/template"
)

type FormAttribute struct {
	Name      string
	Content   []FormStruct
	Url       string
	Method    string
	InfoUrl   string
	CSRFToken string
	Title     string
	Prefix    string
}

func Form() *FormAttribute {
	return &FormAttribute{
		Name:    "form",
		Content: []FormStruct{},
		Url:     "/",
		Method:  "post",
		InfoUrl: "",
		Title:   "edit",
	}
}

func (compo *FormAttribute) SetContent(value []FormStruct) *FormAttribute {
	(*compo).Content = value
	return compo
}

func (compo *FormAttribute) SetPrefix(value string) *FormAttribute {
	(*compo).Prefix = value
	return compo
}

func (compo *FormAttribute) SetUrl(value string) *FormAttribute {
	(*compo).Url = value
	return compo
}

func (compo *FormAttribute) SetInfoUrl(value string) *FormAttribute {
	(*compo).InfoUrl = value
	return compo
}

func (compo *FormAttribute) SetMethod(value string) *FormAttribute {
	(*compo).Method = value
	return compo
}

func (compo *FormAttribute) SetTitle(value string) *FormAttribute {
	(*compo).Title = value
	return compo
}

func (compo *FormAttribute) SetToken(value string) *FormAttribute {
	(*compo).CSRFToken = value
	return compo
}

func (compo *FormAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "form",
		"form/default", "form/file", "form/textarea",
		"form/selectbox", "form/text",
		"form/password", "form/select", "form/iconpicker")
}
