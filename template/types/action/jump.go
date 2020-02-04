package action

import (
	"github.com/GoAdminGroup/go-admin/context"
	"html/template"
)

type JumpAction struct {
	BtnId       string
	Url         string
	Ext         template.HTML
	JS          template.JS
	NewTabTitle string
}

func Jump(url string, ext ...template.HTML) *JumpAction {
	if len(ext) > 0 {
		return &JumpAction{Url: url, Ext: ext[0]}
	}
	return &JumpAction{Url: url, NewTabTitle: ""}
}

func JumpInNewTab(url, title string, ext ...template.HTML) *JumpAction {
	if len(ext) > 0 {
		return &JumpAction{Url: url, NewTabTitle: title, Ext: ext[0]}
	}
	return &JumpAction{Url: url, NewTabTitle: title}
}

func (jump *JumpAction) GetCallbacks() context.Node {
	return context.Node{}
}

func (jump *JumpAction) SetBtnId(btnId string) {
	jump.BtnId = btnId
}

func (jump *JumpAction) Js() template.JS {
	return jump.JS
}

func (jump *JumpAction) BtnAttribute() template.HTML {
	if jump.NewTabTitle != "" {
		return template.HTML(`href="` + jump.Url + `" class="new-tab-link" data-title="` + jump.NewTabTitle + `"`)
	}
	return template.HTML(`href="` + jump.Url + `"`)
}

func (jump *JumpAction) ExtContent() template.HTML {
	return jump.Ext
}
