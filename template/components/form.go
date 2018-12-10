package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type FormAttribute struct {
	Name      string
	Content   []types.Form
	Url       string
	Method    string
	InfoUrl   string
	CSRFToken string
	Title     string
	Prefix    string
	types.Attribute
}

func (compo *FormAttribute) SetContent(value []types.Form) types.FormAttribute {
	compo.Content = value
	return compo
}

func (compo *FormAttribute) SetPrefix(value string) types.FormAttribute {
	compo.Prefix = value
	return compo
}

func (compo *FormAttribute) SetUrl(value string) types.FormAttribute {
	compo.Url = value
	return compo
}

func (compo *FormAttribute) SetInfoUrl(value string) types.FormAttribute {
	compo.InfoUrl = value
	return compo
}

func (compo *FormAttribute) SetMethod(value string) types.FormAttribute {
	compo.Method = value
	return compo
}

func (compo *FormAttribute) SetTitle(value string) types.FormAttribute {
	compo.Title = value
	return compo
}

func (compo *FormAttribute) SetToken(value string) types.FormAttribute {
	compo.CSRFToken = value
	return compo
}

func (compo *FormAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "form",
		"form/default", "form/file", "form/textarea",
		"form/selectbox", "form/text",
		"form/password", "form/select", "form/singleselect", "form/iconpicker")
}
