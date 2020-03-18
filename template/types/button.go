package types

import (
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"html/template"
)

type Button interface {
	Content() (template.HTML, template.JS)
	GetAction() Action
	ID() string
}

type BaseButton struct {
	Id     string
	Title  template.HTML
	Action Action
}

func (b *BaseButton) Content() (template.HTML, template.JS) {
	return "", ""
}

func (b *BaseButton) GetAction() Action {
	return b.Action
}

func (b *BaseButton) ID() string {
	return b.Id
}

type DefaultButton struct {
	*BaseButton
	Color     template.HTML
	TextColor template.HTML
	Icon      string
}

func GetDefaultButton(title template.HTML, icon string, action Action, colors ...template.HTML) *DefaultButton {

	id := btnUUID()
	action.SetBtnId(id)

	var color, textColor template.HTML
	if len(colors) > 0 {
		color = colors[0]
	}
	if len(colors) > 1 {
		textColor = colors[1]
	}
	return &DefaultButton{
		BaseButton: &BaseButton{
			Id:     id,
			Title:  title,
			Action: action,
		},
		Color:     color,
		TextColor: textColor,
		Icon:      icon,
	}
}

func (b *DefaultButton) Content() (template.HTML, template.JS) {

	color := template.HTML("")
	if b.Color != template.HTML("") {
		color = template.HTML(`background-color:`) + b.Color + template.HTML(`;`)
	}
	textColor := template.HTML("")
	if b.TextColor != template.HTML("") {
		textColor = template.HTML(`color:`) + b.TextColor + template.HTML(`;`)
	}

	style := template.HTML("")
	addColor := color + textColor

	if addColor != template.HTML("") {
		style = template.HTML(`style="`) + addColor + template.HTML(`"`)
	}

	h := `<div class="btn-group pull-right" style="margin-right: 10px">
                <a ` + style + ` class="` + template.HTML(b.Id) + ` btn btn-sm btn-default ` + b.Action.BtnClass() + `" ` + b.Action.BtnAttribute() + `>
                    <i class="fa ` + template.HTML(b.Icon) + `"></i>&nbsp;&nbsp;` + b.Title + `
                </a>
        </div>` + b.Action.ExtContent()
	return h, b.Action.Js()
}

type ActionButton struct {
	*BaseButton
}

func GetActionButton(title template.HTML, action Action, ids ...string) *ActionButton {

	id := ""
	if len(ids) > 0 {
		id = ids[0]
	} else {
		id = "action-info-btn-" + utils.Uuid(10)
	}

	action.SetBtnId(id)

	return &ActionButton{
		BaseButton: &BaseButton{
			Id:     id,
			Title:  title,
			Action: action,
		},
	}
}

func (b *ActionButton) Content() (template.HTML, template.JS) {
	h := template.HTML(`<li style="cursor: pointer;"><a data-id="{{.Id}}" class="`+template.HTML(b.Id)+`" `+b.Action.BtnAttribute()+`>`+b.Title+`</a></li>`) + b.Action.ExtContent()
	return h, b.Action.Js()
}

type Buttons []Button

func (b Buttons) Content() (template.HTML, template.JS) {
	h := template.HTML("")
	j := template.JS("")

	for _, btn := range b {
		hh, jj := btn.Content()
		h += hh
		j += jj
	}
	return h, j
}

type NavButton struct {
	*BaseButton
	Icon string
}

func GetNavButton(title template.HTML, icon string, action Action) *NavButton {

	id := btnUUID()
	action.SetBtnId(id)

	return &NavButton{
		BaseButton: &BaseButton{
			Id:     id,
			Title:  title,
			Action: action,
		},
		Icon: icon,
	}
}

func (n *NavButton) Content() (template.HTML, template.JS) {

	icon := template.HTML("")
	title := template.HTML("")

	if n.Icon != "" {
		icon = template.HTML(`<i class="fa ` + n.Icon + `"></i>`)
	}

	if n.Title != "" {
		title = `<span>` + n.Title + `</span>`
	}

	h := template.HTML(`<li>
    <a class="`+template.HTML(n.Id)+`" `+n.Action.BtnAttribute()+`>
      `+icon+`
      `+title+`
    </a>
</li>`) + n.Action.ExtContent()
	return h, n.Action.Js()
}
