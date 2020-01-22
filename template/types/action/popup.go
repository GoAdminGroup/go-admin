package action

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"html/template"
)

type PopUpAction struct {
	BtnId    string
	Url      string
	Method   string
	Id       string
	Title    string
	Handlers []context.Handler
}

func PopUp(url, title string, handlers ...context.Handler) *PopUpAction {
	return &PopUpAction{
		Url:      url,
		Title:    title,
		Method:   "post",
		Id:       "info-popup-model-" + utils.Uuid(10),
		Handlers: handlers,
	}
}

func (pop *PopUpAction) SetUrl(url string) {
	pop.Url = url
}

func (pop *PopUpAction) SetMethod(method string) {
	pop.Method = method
}

func (pop *PopUpAction) GetCallbacks() context.Node {
	return context.Node{Path: pop.Url, Method: pop.Method, Handlers: pop.Handlers}
}

func (pop *PopUpAction) SetBtnId(btnId string) {
	pop.BtnId = btnId
}

func (pop *PopUpAction) Js() template.JS {
	return template.JS(`$('#` + pop.BtnId + `').on('click', function (event) {
						$.ajax({
                            method: '` + pop.Method + `',
                            url: "` + pop.Url + `",
                            data: {
								"ids": {%ids}
							},
                            success: function (data) { 
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code === 0) {
                                    $('#` + pop.Id + ` .modal-body').html(data.data);
                                } else {
                                    swal(data.msg, '', 'error');
                                }
                            }
                        });
            		});`)
}

func (pop *PopUpAction) BtnAttribute() template.HTML {
	return template.HTML(`data-toggle="modal" data-target="#` + pop.Id + ` "`)
}

func (pop *PopUpAction) ExtContent() template.HTML {
	return template2.Default().Popup().SetID(pop.Id).
		SetTitle(template.HTML(pop.Title)).
		SetBody(template.HTML(``)).
		GetContent()
}
