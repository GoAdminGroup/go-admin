package components

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
	Src    string
}

func GetImage() *ImgAttribute {
	return &ImgAttribute{
		"image",
		"50",
		"50",
		"",
	}
}

func (compo *ImgAttribute) SetWidth(value string) *ImgAttribute {
	(*compo).Witdh = value
	return compo
}

func (compo *ImgAttribute) SetHeight(value string) *ImgAttribute {
	(*compo).Height = value
	return compo
}

func (compo *ImgAttribute) SetSrc(value string) *ImgAttribute {
	(*compo).Src = value
	return compo
}

func (compo *ImgAttribute) GetContent() string {
	return `<img src="` + (*compo).Src + `" width="` + (*compo).Witdh + `" height="` + (*compo).Height + `">`
}
