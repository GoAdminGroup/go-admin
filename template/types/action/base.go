package action

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"strings"
)

type AjaxData map[string]interface{}

func NewAjaxData() AjaxData {
	return AjaxData{"ids": "{{.Ids}}"}
}

func (a AjaxData) Add(m map[string]interface{}) AjaxData {
	for k, v := range m {
		a[k] = v
	}
	return a
}

func (a AjaxData) JSON() string {
	b, _ := json.Marshal(a)
	s := strings.Replace(string(b), `"{%id}"`, "{{.Id}}", -1)
	s = strings.Replace(s, `"{%ids}"`, "{{.Ids}}", -1)
	s = strings.Replace(s, `"{{.Ids}}"`, "{{.Ids}}", -1)
	s = strings.Replace(s, `"{{.Id}}"`, "{{.Id}}", -1)
	return s
}

type BaseAction struct {
	BtnId   string
	BtnData interface{}
	JS      template.JS
}

func (base *BaseAction) SetBtnId(btnId string)        { base.BtnId = btnId }
func (base *BaseAction) Js() template.JS              { return base.JS }
func (base *BaseAction) BtnClass() template.HTML      { return "" }
func (base *BaseAction) BtnAttribute() template.HTML  { return "" }
func (base *BaseAction) GetCallbacks() context.Node   { return context.Node{} }
func (base *BaseAction) ExtContent() template.HTML    { return template.HTML(``) }
func (base *BaseAction) FooterContent() template.HTML { return template.HTML(``) }
func (base *BaseAction) SetBtnData(data interface{})  { base.BtnData = data }

var _ types.Action = (*AjaxAction)(nil)
var _ types.Action = (*PopUpAction)(nil)
var _ types.Action = (*JumpAction)(nil)
var _ types.Action = (*JumpSelectBoxAction)(nil)

func URL(id string) string {
	return config.Url("/operation/" + utils.WrapURL(id))
}
