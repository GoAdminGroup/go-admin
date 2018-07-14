package components

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
}

var Image = &ImgAttribute{
	"image",
	"50",
	"50",
}

func (compo *ImgAttribute) GetContent(value interface{}) string {
	if valueStr, ok := value.(string); ok {
		return `<img src="` + valueStr + `" width="50" height="50">`
	}
	if valueMap, ok := value.(map[string]string); ok {
		return `<img src="` + valueMap["url"] + `" width="` + valueMap["width"] + `" height="` + valueMap["height"] + `">`
	}
	return ""
}
