package adminlte

import (
	"html/template"
	"goAdmin/components"
)

type FormAttribute struct {
	Name      string
	Content   []components.FormStruct
	Url       string
	Method    string
	InfoUrl   string
	CSRFToken string
}

func (AdminlteComponents) Form() *FormAttribute {
	return &FormAttribute{
		Name:    "form",
		Content: []components.FormStruct{},
		Url:     "/",
		Method:  "post",
		InfoUrl: "",
	}
}

func (compo *FormAttribute) SetContent(value []components.FormStruct) *FormAttribute {
	(*compo).Content = value
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

func (compo *FormAttribute) SetToken(value string) *FormAttribute {
	(*compo).CSRFToken = value
	return compo
}

func (compo *FormAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "form",
		"form/default", "form/file", "form/textarea",
		"form/selectbox", "form/text",
		"form/password", "form/select")
}
