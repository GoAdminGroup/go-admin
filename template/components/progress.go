package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type ProgressGroupAttribute struct {
	Name        string
	Title       string
	Molecular   int
	Denominator int
	Color       string
	Percent     int
	types.Attribute
}

func (compo *ProgressGroupAttribute) SetTitle(value string) types.ProgressGroupAttribute {
	compo.Title = value
	return compo
}

func (compo *ProgressGroupAttribute) SetColor(value string) types.ProgressGroupAttribute {
	compo.Color = value
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
