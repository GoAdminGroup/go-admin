package action

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"html/template"
	"net/http"
	"strings"
)

type AjaxAction struct {
	BtnId    string
	Url      string
	Method   string
	Title    string
	Data     AjaxData
	Handlers []context.Handler
}

type AjaxData map[string]interface{}

func NewAjaxData() AjaxData {
	return AjaxData{"ids": "{%ids}"}
}

func (a AjaxData) Add(m map[string]interface{}) AjaxData {
	for k, v := range m {
		a[k] = v
	}
	return a
}

func (a AjaxData) JSON() string {
	b, _ := json.Marshal(a)
	return strings.Replace(string(b), `"{%ids}"`, "{%ids}", -1)
}

type Handler func(ctx *context.Context) (success bool, data, msg string)

func (h Handler) Wrap() context.Handler {
	return func(ctx *context.Context) {
		s, d, m := h(ctx)
		code := 0
		if !s {
			code = 500
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": code,
			"data": d,
			"msg":  m,
		})
	}
}

func Ajax(url, title string, handler Handler) *AjaxAction {
	return &AjaxAction{
		Url:      url,
		Title:    title,
		Method:   "post",
		Data:     NewAjaxData(),
		Handlers: context.Handlers{handler.Wrap()},
	}
}

func (ajax *AjaxAction) SetData(data map[string]interface{}) *AjaxAction {
	ajax.Data = ajax.Data.Add(data)
	return ajax
}

func (ajax *AjaxAction) SetUrl(url string) *AjaxAction {
	ajax.Url = url
	return ajax
}

func (ajax *AjaxAction) SetMethod(method string) *AjaxAction {
	ajax.Method = method
	return ajax
}

func (ajax *AjaxAction) GetCallbacks() context.Node {
	return context.Node{
		Path:     ajax.Url,
		Method:   ajax.Method,
		Handlers: ajax.Handlers,
		Value:    map[string]interface{}{auth.ContextNodeNeedAuth: 1},
	}
}

func (ajax *AjaxAction) SetBtnId(btnId string) {
	ajax.BtnId = btnId
}

func (ajax *AjaxAction) Js() template.JS {
	return template.JS(`$('#` + ajax.BtnId + `').on('click', function (event) {
						$.ajax({
                            method: '` + ajax.Method + `',
                            url: "` + ajax.Url + `",
                            data: ` + ajax.Data.JSON() + `,
                            success: function (data) { 
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code === 0) {
                                    swal(data.msg, '', 'success');
                                } else {
                                    swal(data.msg, '', 'error');
                                }
                            }
                        });
            		});`)
}

func (ajax *AjaxAction) BtnAttribute() template.HTML { return template.HTML(``) }
func (ajax *AjaxAction) ExtContent() template.HTML   { return template.HTML(``) }
