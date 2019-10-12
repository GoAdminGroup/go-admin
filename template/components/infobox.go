package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"strings"
)

type InfoBoxAttribute struct {
	Name       string
	Icon       template.HTML
	Text       template.HTML
	Number     template.HTML
	Content    template.HTML
	Color      template.HTML
	IsHexColor bool
	IsSvg      bool
	types.Attribute
}

func (compo *InfoBoxAttribute) SetIcon(value template.HTML) types.InfoBoxAttribute {
	compo.Icon = value
	if strings.Contains(string(value), "svg") {
		compo.IsSvg = true
	}
	return compo
}

func (compo *InfoBoxAttribute) SetText(value template.HTML) types.InfoBoxAttribute {
	compo.Text = value
	return compo
}

func (compo *InfoBoxAttribute) SetNumber(value template.HTML) types.InfoBoxAttribute {
	compo.Number = value
	return compo
}

func (compo *InfoBoxAttribute) SetContent(value template.HTML) types.InfoBoxAttribute {
	compo.Content = value
	return compo
}

func (compo *InfoBoxAttribute) SetColor(value template.HTML) types.InfoBoxAttribute {
	compo.Color = value
	if strings.Contains(string(value), "#") {
		compo.IsHexColor = true
	}
	return compo
}

func (compo *InfoBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "infobox")
}
