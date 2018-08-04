package components

type ColAttribute struct {
	Name  string
	Width string
}

var Col = &ColAttribute{
	"col",
	"2",
}

func (compo *ColAttribute) SetWidth(value string) *ColAttribute {
	(*compo).Width = value
	return compo
}

func (compo *ColAttribute) GetContent(value interface{}) string {
	if valueStr, ok := value.(string); ok {
		return `<div class="col-md-` + (*compo).Width + `">` + valueStr + `</div>`
	}
	return ""
}
