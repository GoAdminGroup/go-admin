package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type LinkAttribute struct {
	Name    string
	URL     string
	Title   template.HTML
	NewTab  bool
	Content template.HTML
	types.Attribute
}

func (compo *LinkAttribute) OpenInNewTab() types.LinkAttribute {
	compo.NewTab = true
	return compo
}

func (compo *LinkAttribute) SetURL(value string) types.LinkAttribute {
	compo.URL = value
	return compo
}

func (compo *LinkAttribute) SetTabTitle(value template.HTML) types.LinkAttribute {
	compo.Title = value
	return compo
}

func (compo *LinkAttribute) SetContent(value template.HTML) types.LinkAttribute {
	compo.Content = value
	return compo
}

func (compo *LinkAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "link")
}
