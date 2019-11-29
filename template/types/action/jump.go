package action

import (
	"html/template"
)

type JumpAction struct {
	BtnId string
	Url   string
}

func Jump(url string) *JumpAction {
	return &JumpAction{Url: url}
}

func (pop *JumpAction) SetBtnId(btnId string) {
	pop.BtnId = btnId
}

func (pop *JumpAction) Js() template.JS {
	return template.JS(``)
}

func (pop *JumpAction) BtnAttribute() template.HTML {
	return template.HTML(`href="` + pop.Url + `"`)
}

func (pop *JumpAction) ExtContent() template.HTML {
	return ``
}
