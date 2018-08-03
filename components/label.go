package components

type LabelAttribute struct {
	Name  string
	Color string
}

var Label = &LabelAttribute{
	"label",
	"success",
}

func (compo *LabelAttribute) GetContent(value interface{}) string {
	if valueStr, ok := value.(string); ok {
		return `<span class="label label-` + (*compo).Color + `">` + valueStr + `</span>`
	}
	return ""
}
