package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type LinkAttribute struct {
	Name       string
	URL        string
	Class      template.HTML
	Title      template.HTML
	Attributes template.HTMLAttr
	Content    template.HTML
	types.Attribute
}

func (compo *LinkAttribute) OpenInNewTab() types.LinkAttribute {
	compo.Class += " new-tab-link"
	return compo
}

func (compo *LinkAttribute) SetURL(value string) types.LinkAttribute {
	compo.URL = value
	return compo
}

func (compo *LinkAttribute) SetClass(class template.HTML) types.LinkAttribute {
	compo.Class = class
	return compo
}

func (compo *LinkAttribute) SetAttributes(attr template.HTMLAttr) types.LinkAttribute {
	compo.Attributes = attr
	return compo
}

func (compo *LinkAttribute) NoPjax() types.LinkAttribute {
	compo.Class += " no-pjax"
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
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "link")
}
