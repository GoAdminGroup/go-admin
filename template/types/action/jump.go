package action

import (
	"github.com/GoAdminGroup/go-admin/context"
	"html/template"
	"strings"
)

type JumpAction struct {
	BaseAction
	Url         string
	Ext         template.HTML
	NewTabTitle string
}

func Jump(url string, ext ...template.HTML) *JumpAction {
	url = strings.Replace(url, "{%id}", "{{.Id}}", -1)
	url = strings.Replace(url, "{%ids}", "{{.Ids}}", -1)
	if len(ext) > 0 {
		return &JumpAction{Url: url, Ext: ext[0]}
	}
	return &JumpAction{Url: url, NewTabTitle: ""}
}

func JumpInNewTab(url, title string, ext ...template.HTML) *JumpAction {
	url = strings.Replace(url, "{%id}", "{{.Id}}", -1)
	url = strings.Replace(url, "{%ids}", "{{.Ids}}", -1)
	if len(ext) > 0 {
		return &JumpAction{Url: url, NewTabTitle: title, Ext: ext[0]}
	}
	return &JumpAction{Url: url, NewTabTitle: title}
}

func (jump *JumpAction) GetCallbacks() context.Node {
	return context.Node{}
}

func (jump *JumpAction) BtnAttribute() template.HTML {
	if jump.NewTabTitle != "" {
		return template.HTML(`href="` + jump.Url + `" data-title="` + jump.NewTabTitle + `"`)
	}
	return template.HTML(`href="` + jump.Url + `"`)
}

func (jump *JumpAction) BtnClass() template.HTML {
	if jump.NewTabTitle != "" {
		return "new-tab-link"
	}
	return ""
}

func (jump *JumpAction) ExtContent() template.HTML {
	return jump.Ext
}
