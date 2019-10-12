package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"strings"
)

type ProgressGroupAttribute struct {
	Name        string
	Title       template.HTML
	Molecular   int
	Denominator int
	Color       template.HTML
	IsHexColor  bool
	Percent     int
	types.Attribute
}

func (compo *ProgressGroupAttribute) SetTitle(value template.HTML) types.ProgressGroupAttribute {
	compo.Title = value
	return compo
}

func (compo *ProgressGroupAttribute) SetColor(value template.HTML) types.ProgressGroupAttribute {
	compo.Color = value
	if strings.Contains(string(value), "#") {
		compo.IsHexColor = true
	}
	return compo
}

func (compo *ProgressGroupAttribute) SetPercent(value int) types.ProgressGroupAttribute {
	compo.Percent = value
	return compo
}

func (compo *ProgressGroupAttribute) SetDenominator(value int) types.ProgressGroupAttribute {
	compo.Denominator = value
	return compo
}

func (compo *ProgressGroupAttribute) SetMolecular(value int) types.ProgressGroupAttribute {
	compo.Molecular = value
	return compo
}

func (compo *ProgressGroupAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "progress-group")
}
