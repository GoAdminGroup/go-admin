package action

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type PopUpAction struct {
	BaseAction
	Url      string
	Method   string
	Id       string
	Title    string
	BtnTitle template.HTML
	Data     AjaxData
	Handlers []context.Handler
}

func PopUp(id, title string, handler types.Handler) *PopUpAction {
	if id == "" {
		panic("wrong popup action parameter, empty id")
	}
	return &PopUpAction{
		Url:      URL(id),
		Title:    title,
		Method:   "post",
		BtnTitle: "",
		Data:     NewAjaxData(),
		Id:       "info-popup-model-" + utils.Uuid(10),
		Handlers: context.Handlers{handler.Wrap()},
	}
}

func (pop *PopUpAction) SetData(data map[string]interface{}) *PopUpAction {
	pop.Data = pop.Data.Add(data)
	return pop
}

func (pop *PopUpAction) SetUrl(url string) *PopUpAction {
	pop.Url = url
	return pop
}

func (pop *PopUpAction) SetBtnTitle(title template.HTML) *PopUpAction {
	pop.BtnTitle = title
	return pop
}

func (pop *PopUpAction) SetMethod(method string) *PopUpAction {
	pop.Method = method
	return pop
}

func (pop *PopUpAction) GetCallbacks() context.Node {
	return context.Node{
		Path:     pop.Url,
		Method:   pop.Method,
		Handlers: pop.Handlers,
		Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
	}
}

func (pop *PopUpAction) Js() template.JS {
	return template.JS(`$('.` + pop.BtnId + `').on('click', function (event) {
						let data = ` + pop.Data.JSON() + `;
						let id = $(this).attr("data-id");
						if (id && id !== "") {
							data["id"] = id;
						}
						data['popup_id'] = "` + pop.Id + `"
						$.ajax({
                            method: '` + pop.Method + `',
                            url: "` + pop.Url + `",
                            data: data,
                            success: function (data) {
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code === 0) {
                                    $('#` + pop.Id + ` .modal-body').html(data.data);
                                } else {
                                    swal(data.msg, '', 'error');
                                }
                            },
							error: function (data) {
								if (data.responseText !== "") {
									swal(data.responseJSON.msg, '', 'error');								
								} else {
									swal('error', '', 'error');
								}
								setTimeout(function() {
									$('#` + pop.Id + `').hide();
									$('.modal-backdrop.fade.in').hide();
								}, 500)
							},
                        });
            		});`)
}

func (pop *PopUpAction) BtnAttribute() template.HTML {
	return template.HTML(`data-toggle="modal" data-target="#` + pop.Id + ` " data-id="{{.Id}}" style="cursor: pointer;"`)
}

func (pop *PopUpAction) FooterContent() template.HTML {
	return template2.Default().Popup().SetID(pop.Id).
		SetTitle(template.HTML(pop.Title)).
		SetFooter(pop.BtnTitle).
		SetBody(template.HTML(``)).
		GetContent()
}
