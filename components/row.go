package components

type RowAttribute struct {
	Name  string
}

var Row = &RowAttribute{
	"row",
}

func (compo *RowAttribute) GetContent(value interface{}) string {
	if valueStr, ok := value.(string); ok {
		return `<div class="row">` + valueStr + `</div>`
	}
	return ""
}