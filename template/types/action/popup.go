package action

import (
	"github.com/GoAdminGroup/go-admin/modules/utils"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"html/template"
)

type PopUpAction struct {
	BtnId string
	Url   string
	Id    string
	Title string
}

func PopUp(url, title string) *PopUpAction {
	return &PopUpAction{Url: url, Title: title, Id: "info-popup-model-" + utils.Uuid(10)}
}

func (pop *PopUpAction) SetBtnId(btnId string) {
	pop.BtnId = btnId
}

func (pop *PopUpAction) Js() template.JS {
	return template.JS(`$('#` + pop.BtnId + `').on('click', function (event) {
						$.ajax({
                            method: 'post',
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
