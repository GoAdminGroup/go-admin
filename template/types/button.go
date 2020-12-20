package types

import (
	"html/template"
	"net/url"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
)

type Button interface {
	Content() (template.HTML, template.JS)
	GetAction() Action
	URL() string
	METHOD() string
	ID() string
	Type() string
	GetName() string
	SetName(name string)
	IsType(t string) bool
}

type BaseButton struct {
	Id, Url, Method, Name, TypeName string
	Title                           template.HTML
	Action                          Action
}

func (b *BaseButton) Content() (template.HTML, template.JS) { return "", "" }
func (b *BaseButton) GetAction() Action                     { return b.Action }
func (b *BaseButton) ID() string                            { return b.Id }
func (b *BaseButton) URL() string                           { return b.Url }
func (b *BaseButton) Type() string                          { return b.TypeName }
func (b *BaseButton) IsType(t string) bool                  { return b.TypeName == t }
func (b *BaseButton) METHOD() string                        { return b.Method }
func (b *BaseButton) GetName() string                       { return b.Name }
func (b *BaseButton) SetName(name string)                   { b.Name = name }

type DefaultButton struct {
	*BaseButton
	Color     template.HTML
	TextColor template.HTML
	Icon      string
	Direction template.HTML
	Group     bool
}

func GetDefaultButton(title template.HTML, icon string, action Action, colors ...template.HTML) *DefaultButton {
	return defaultButton(title, "right", icon, action, false, colors...)
}

func GetDefaultButtonGroup(title template.HTML, icon string, action Action, colors ...template.HTML) *DefaultButton {
	return defaultButton(title, "right", icon, action, true, colors...)
}

func defaultButton(title, direction template.HTML, icon string, action Action, group bool, colors ...template.HTML) *DefaultButton {
	id := btnUUID()
	action.SetBtnId("." + id)

	var color, textColor template.HTML
	if len(colors) > 0 {
		color = colors[0]
	}
	if len(colors) > 1 {
		textColor = colors[1]
	}
	node := action.GetCallbacks()
	return &DefaultButton{
		BaseButton: &BaseButton{
			Id:     id,
			Title:  title,
			Action: action,
			Url:    node.Path,
			Method: node.Method,
		},
		Group:     group,
		Color:     color,
		TextColor: textColor,
		Icon:      icon,
		Direction: direction,
	}
}

func GetColumnButton(title template.HTML, icon string, action Action, colors ...template.HTML) *DefaultButton {
	return defaultButton(title, "", icon, action, true, colors...)
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

	h := template.HTML("")
	if b.Group {
		h += `<div class="btn-group pull-` + b.Direction + `" style="margin-right: 10px">`
	}

	h += `<a ` + style + ` class="` + template.HTML(b.Id) + ` btn btn-sm btn-default ` + b.Action.BtnClass() + `" ` + b.Action.BtnAttribute() + `>
                    <i class="fa ` + template.HTML(b.Icon) + `"></i>&nbsp;&nbsp;` + b.Title + `
                </a>`
	if b.Group {
		h += `</div>`
	}
	return h + b.Action.ExtContent(), b.Action.Js()
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

	action.SetBtnId("." + id)
	node := action.GetCallbacks()

	return &ActionButton{
		BaseButton: &BaseButton{
			Id:     id,
			Title:  title,
			Action: action,
			Url:    node.Path,
			Method: node.Method,
		},
	}
}

func (b *ActionButton) Content() (template.HTML, template.JS) {
	h := template.HTML(`<li style="cursor: pointer;"><a data-id="{{.Id}}" class="`+template.HTML(b.Id)+` `+
		b.Action.BtnClass()+`" `+b.Action.BtnAttribute()+`>`+b.Title+`</a></li>`) + b.Action.ExtContent()
	return h, b.Action.Js()
}

type ActionIconButton struct {
	Icon template.HTML
	*BaseButton
}

func GetActionIconButton(icon string, action Action, ids ...string) *ActionIconButton {

	id := ""
	if len(ids) > 0 {
		id = ids[0]
	} else {
		id = "action-info-btn-" + utils.Uuid(10)
	}

	action.SetBtnId("." + id)
	node := action.GetCallbacks()

	return &ActionIconButton{
		Icon: template.HTML(icon),
		BaseButton: &BaseButton{
			Id:     id,
			Action: action,
			Url:    node.Path,
			Method: node.Method,
		},
	}
}

func (b *ActionIconButton) Content() (template.HTML, template.JS) {
	h := template.HTML(`<a data-id="{{.Id}}" class="`+template.HTML(b.Id)+` `+
		b.Action.BtnClass()+`" `+b.Action.BtnAttribute()+`><i class="fa `+b.Icon+`" style="font-size: 16px;"></i></a>`) + b.Action.ExtContent()
	return h, b.Action.Js()
}

type Buttons []Button

func (b Buttons) Add(btn Button) Buttons {
	return append(b, btn)
}

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

func (b Buttons) Copy() Buttons {
	var c = make(Buttons, len(b))
	copy(c, b)
	return c
}

func (b Buttons) FooterContent() template.HTML {
	footer := template.HTML("")

	for _, btn := range b {
		footer += btn.GetAction().FooterContent()
	}
	return footer
}

func (b Buttons) CheckPermission(user models.UserModel) Buttons {
	btns := make(Buttons, 0)
	for _, btn := range b {
		if btn.IsType(ButtonTypeNavDropDown) {
			items := make([]Button, 0)
			for _, navItem := range btn.(*NavDropDownButton).Items {
				if user.CheckPermissionByUrlMethod(btn.URL(), btn.METHOD(), url.Values{}) {
					items = append(items, navItem)
				}
			}
			if len(items) > 0 {
				btns = append(btns, btn)
			}
		} else if user.CheckPermissionByUrlMethod(btn.URL(), btn.METHOD(), url.Values{}) {
			btns = append(btns, btn)
		}
	}
	return btns
}

func (b Buttons) CheckPermissionWhenURLAndMethodNotEmpty(user models.UserModel) Buttons {
	btns := make(Buttons, 0)
	for _, b := range b {
		if b.URL() == "" || b.METHOD() == "" || user.CheckPermissionByUrlMethod(b.URL(), b.METHOD(), url.Values{}) {
			btns = append(btns, b)
		}
	}
	return btns
}

func (b Buttons) AddNavButton(ico, name string, action Action) Buttons {
	if !b.CheckExist(name) {
		return append(b, GetNavButton("", ico, action, name))
	}
	return b
}

func (b Buttons) RemoveButtonByName(name string) Buttons {
	if name == "" {
		return b
	}

	for i := 0; i < len(b); i++ {
		if b[i].GetName() == name {
			b = append(b[:i], b[i+1:]...)
		}
	}
	return b
}

func (b Buttons) CheckExist(name string) bool {
	if name == "" {
		return false
	}
	for i := 0; i < len(b); i++ {
		if b[i].GetName() == name {
			return true
		}
	}
	return false
}

func (b Buttons) Callbacks() []context.Node {
	cbs := make([]context.Node, 0)
	for _, btn := range b {
		cbs = append(cbs, btn.GetAction().GetCallbacks())
	}
	return cbs
}

const (
	NavBtnSiteName = "go_admin_site_navbtn"
	NavBtnInfoName = "go_admin_info_navbtn"
	NavBtnToolName = "go_admin_tool_navbtn"
	NavBtnPlugName = "go_admin_plug_navbtn"
)

func (b Buttons) RemoveSiteNavButton() Buttons {
	return b.RemoveButtonByName(NavBtnSiteName)
}

func (b Buttons) RemoveInfoNavButton() Buttons {
	return b.RemoveButtonByName(NavBtnInfoName)
}

func (b Buttons) RemoveToolNavButton() Buttons {
	return b.RemoveButtonByName(NavBtnToolName)
}

func (b Buttons) RemovePlugNavButton() Buttons {
	return b.RemoveButtonByName(NavBtnPlugName)
}

type NavButton struct {
	*BaseButton
	Icon string
}

func GetNavButton(title template.HTML, icon string, action Action, names ...string) *NavButton {

	id := btnUUID()
	action.SetBtnId("." + id)
	node := action.GetCallbacks()
	name := ""

	if len(names) > 0 {
		name = names[0]
	}

	return &NavButton{
		BaseButton: &BaseButton{
			Id:     id,
			Title:  title,
			Action: action,
			Url:    node.Path,
			Method: node.Method,
			Name:   name,
		},
		Icon: icon,
	}
}

func (n *NavButton) Content() (template.HTML, template.JS) {

	ico := template.HTML("")
	title := template.HTML("")

	if n.Icon != "" {
		ico = template.HTML(`<i class="fa ` + n.Icon + `"></i>`)
	}

	if n.Title != "" {
		title = `<span>` + n.Title + `</span>`
	}

	h := template.HTML(`<li>
    <a class="`+template.HTML(n.Id)+` `+n.Action.BtnClass()+`" `+n.Action.BtnAttribute()+`>
      `+ico+`
      `+title+`
    </a>
</li>`) + n.Action.ExtContent()
	return h, n.Action.Js()
}

type NavDropDownButton struct {
	*BaseButton
	Icon  string
	Items []*NavDropDownItemButton
}

type NavDropDownItemButton struct {
	*BaseButton
}

func GetDropDownButton(title template.HTML, icon string, items []*NavDropDownItemButton, names ...string) *NavDropDownButton {
	id := btnUUID()
	name := ""

	if len(names) > 0 {
		name = names[0]
	}

	return &NavDropDownButton{
		BaseButton: &BaseButton{
			Id:       id,
			Title:    title,
			Name:     name,
			TypeName: ButtonTypeNavDropDown,
			Action:   new(NilAction),
		},
		Items: items,
		Icon:  icon,
	}
}

func (n *NavDropDownButton) SetItems(items []*NavDropDownItemButton) {
	n.Items = items
}

func (n *NavDropDownButton) AddItem(item *NavDropDownItemButton) {
	n.Items = append(n.Items, item)
}

func (n *NavDropDownButton) Content() (template.HTML, template.JS) {

	ico := template.HTML("")
	title := template.HTML("")

	if n.Icon != "" {
		ico = template.HTML(`<i class="fa ` + n.Icon + `"></i>`)
	}

	if n.Title != "" {
		title = `<span>` + n.Title + `</span>`
	}

	content := template.HTML("")
	js := template.JS("")

	for _, item := range n.Items {
		c, j := item.Content()
		content += c
		js += j
	}

	did := utils.Uuid(10)

	h := template.HTML(`<li class="dropdown" id="` + template.HTML(did) + `">
    <a class="` + template.HTML(n.Id) + ` dropdown-toggle" data-toggle="dropdown" style="cursor:pointer;">
      ` + ico + `
      ` + title + `
    </a>
	<ul class="dropdown-menu"  aria-labelledby="` + template.HTML(did) + `">
    	` + content + `
	</ul>
</li>`)

	return h, js
}

const (
	ButtonTypeNavDropDownItem = "navdropdownitem"
	ButtonTypeNavDropDown     = "navdropdown"
)

func GetDropDownItemButton(title template.HTML, action Action, names ...string) *NavDropDownItemButton {
	id := btnUUID()
	action.SetBtnId("." + id)
	node := action.GetCallbacks()
	name := ""

	if len(names) > 0 {
		name = names[0]
	}

	return &NavDropDownItemButton{
		BaseButton: &BaseButton{
			Id:       id,
			Title:    title,
			Action:   action,
			Url:      node.Path,
			Method:   node.Method,
			Name:     name,
			TypeName: ButtonTypeNavDropDownItem,
		},
	}
}

func (n *NavDropDownItemButton) Content() (template.HTML, template.JS) {

	title := template.HTML("")

	if n.Title != "" {
		title = `<span>` + n.Title + `</span>`
	}

	h := template.HTML(`<li><a class="dropdown-item `+template.HTML(n.Id)+` `+
		n.Action.BtnClass()+`" `+n.Action.BtnAttribute()+`>
      `+title+`
</a></li>`) + n.Action.ExtContent()
	return h, n.Action.Js()
}
