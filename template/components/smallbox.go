package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"strings"
)

type SmallBoxAttribute struct {
	Name       string
	Title      template.HTML
	Value      template.HTML
	Url        string
	Color      template.HTML
	IsSvg      bool
	IsHexColor bool
	Icon       template.HTML
	types.Attribute
}

func (compo *SmallBoxAttribute) SetTitle(value template.HTML) types.SmallBoxAttribute {
	compo.Title = value
	return compo
}

func (compo *SmallBoxAttribute) SetValue(value template.HTML) types.SmallBoxAttribute {
	compo.Value = value
	return compo
}

func (compo *SmallBoxAttribute) SetColor(value template.HTML) types.SmallBoxAttribute {
	compo.Color = value
	if strings.Contains(string(value), "#") {
		compo.IsHexColor = true
	}
	return compo
}

func (compo *SmallBoxAttribute) SetIcon(value template.HTML) types.SmallBoxAttribute {
	compo.Icon = value
	if strings.Contains(string(value), "svg") {
		compo.IsSvg = true
	}
	return compo
}

func (compo *SmallBoxAttribute) SetUrl(value string) types.SmallBoxAttribute {
	compo.Url = value
	return compo
}

func (compo *SmallBoxAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "smallbox")
}
