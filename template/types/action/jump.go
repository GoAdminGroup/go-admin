package action

import (
	"html/template"
)

type JumpAction struct {
	BtnId string
	Url   string
	Ext   template.HTML
	JS    template.JS
}

func Jump(url string, ext ...template.HTML) *JumpAction {
	if len(ext) > 0 {
		return &JumpAction{Url: url, Ext: ext[0]}
	}
	return &JumpAction{Url: url}
}

func (jump *JumpAction) SetBtnId(btnId string) {
	jump.BtnId = btnId
}

func (jump *JumpAction) Js() template.JS {
	return jump.JS
}

func (jump *JumpAction) BtnAttribute() template.HTML {
	return template.HTML(`href="` + jump.Url + `"`)
}

func (jump *JumpAction) ExtContent() template.HTML {
	return jump.Ext
}
