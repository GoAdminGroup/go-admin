package adminlte

import (
	"html/template"
	"goAdmin/components"
)

type FormAttribute struct {
	Name    string
	Content []components.FormStruct
}

func (AdminlteComponents) Form() *FormAttribute {
	return &FormAttribute{
		"form",
		[]components.FormStruct{},
	}
}

func (compo *FormAttribute) SetContent(value []components.FormStruct) *FormAttribute {
	(*compo).Content = value
	return compo
}

func (compo *FormAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "form",
		"form/default", "form/file", "form/textarea",
		"form/selectbox", "form/text",
		"form/password", "form/select")
}
