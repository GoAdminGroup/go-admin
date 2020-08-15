package action

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type PopUpAction struct {
	BaseAction
	Url         string
	Method      string
	Id          string
	Title       string
	Draggable   bool
	Width       string
	Height      string
	HasIframe   bool
	HideFooter  bool
	BtnTitle    template.HTML
	ParameterJS template.JS
	Data        AjaxData
	Handlers    []context.Handler
	Event       Event
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
		Event:    EventClick,
	}
}

func (pop *PopUpAction) SetData(data map[string]interface{}) *PopUpAction {
	pop.Data = pop.Data.Add(data)
	return pop
}

func (pop *PopUpAction) SetDraggable() *PopUpAction {
	pop.Draggable = true
	return pop
}

func (pop *PopUpAction) SetWidth(width string) *PopUpAction {
	pop.Width = width
	return pop
}

func (pop *PopUpAction) SetHeight(height string) *PopUpAction {
	pop.Height = height
	return pop
}

func (pop *PopUpAction) SetParameterJS(parameterJS template.JS) *PopUpAction {
	pop.ParameterJS += parameterJS
	return pop
}

func (pop *PopUpAction) SetUrl(url string) *PopUpAction {
	pop.Url = url
	return pop
}

type IframeData struct {
	Width          string
	Height         string
	Src            string
	AddParameterFn func(ctx *context.Context) string
}

func PopUpWithIframe(id, title string, data IframeData, width, height string) *PopUpAction {
	if id == "" {
		panic("wrong popup action parameter, empty id")
	}
	if data.Width == "" {
		data.Width = "100%"
	}
	if data.Height == "" {
		data.Height = "100%"
	}
	if strings.Contains(data.Src, "?") {
		data.Src = data.Src + "&"
	} else {
		data.Src = data.Src + "?"
	}
	modalID := "info-popup-model-" + utils.Uuid(10)
	var handler types.Handler = func(ctx *context.Context) (success bool, msg string, res interface{}) {
		param := ""
		if data.AddParameterFn != nil {
			param = data.AddParameterFn(ctx)
		}
		return true, "ok", fmt.Sprintf(`<iframe style="width:%s;height:%s;" 
			scrolling="auto" 
			allowtransparency="true" 
			frameborder="0"
			src="%s__goadmin_iframe=true&__go_admin_no_animation_=true&__goadmin_iframe_id=%s`+param+`"><iframe>`,
			data.Width, data.Height, data.Src, modalID)
	}
	return &PopUpAction{
		Url:        URL(id),
		Title:      title,
		Method:     "post",
		BtnTitle:   "",
		Height:     height,
		HasIframe:  true,
		HideFooter: isFormURL(data.Src),
		Width:      width,
		Draggable:  true,
		Data:       NewAjaxData(),
		Id:         modalID,
		Handlers:   context.Handlers{handler.Wrap()},
		Event:      EventClick,
	}
}

type PopUpData struct {
	Id     string
	Title  string
	Width  string
	Height string
}

type GetForm func(panel *types.FormPanel) *types.FormPanel

var operationHandlerSetter context.NodeProcessor

func InitOperationHandlerSetter(p context.NodeProcessor) {
	operationHandlerSetter = p
}

func PopUpWithForm(data PopUpData, fn GetForm, url string) *PopUpAction {
	if data.Id == "" {
		panic("wrong popup action parameter, empty id")
	}
	modalID := "info-popup-model-" + utils.Uuid(10)

	var handler types.Handler = func(ctx *context.Context) (success bool, msg string, res interface{}) {
		col1 := template2.Default().Col().GetContent()
		btn1 := template2.Default().Button().SetType("submit").
			SetContent(language.GetFromHtml("Save")).
			SetThemePrimary().
			SetOrientationRight().
			SetLoadingText(icon.Icon("fa-spinner fa-spin", 2) + language.GetFromHtml("Save")).
			GetContent()
		btn2 := template2.Default().Button().SetType("reset").
			SetContent(language.GetFromHtml("Reset")).
			SetThemeWarning().
			SetOrientationLeft().
			GetContent()
		col2 := template2.Default().Col().SetSize(types.SizeMD(8)).
			SetContent(btn1 + btn2).GetContent()
		panel := fn(types.NewFormPanel())

		operationHandlerSetter(panel.Callbacks...)

		fields, tabFields, tabHeaders := panel.GetNewFormFields()

		return true, "ok", template2.Default().Box().
			SetHeader("").
			SetBody(template2.Default().Form().
				SetContent(fields).
				SetTabHeaders(tabHeaders).
				SetTabContents(tabFields).
				SetAjax(panel.AjaxSuccessJS, panel.AjaxErrorJS).
				SetPrefix(config.PrefixFixSlash()).
				SetUrl(url).
				SetOperationFooter(col1 + col2).GetContent()).
			GetContent()
	}
	return &PopUpAction{
		Url:        URL(data.Id),
		Title:      data.Title,
		Method:     "post",
		BtnTitle:   "",
		HideFooter: true,
		Height:     data.Height,
		Width:      data.Width,
		Draggable:  true,
		Data:       NewAjaxData(),
		Id:         modalID,
		Handlers:   context.Handlers{handler.Wrap()},
		Event:      EventClick,
	}
}

func (pop *PopUpAction) SetBtnTitle(title template.HTML) *PopUpAction {
	pop.BtnTitle = title
	return pop
}

func (pop *PopUpAction) SetEvent(event Event) *PopUpAction {
	pop.Event = event
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
	return template.JS(`$('`+pop.BtnId+`').on('`+string(pop.Event)+`', function (event) {
						let data = `+pop.Data.JSON()+`;
						`) + pop.ParameterJS + template.JS(`
						let id = $(this).attr("data-id");
						if (id && id !== "") {
							data["id"] = id;
						}
						data['popup_id'] = "`+pop.Id+`"
						$.ajax({
                            method: '`+pop.Method+`',
                            url: "`+pop.Url+`",
                            data: data,
                            success: function (data) {
                                if (typeof (data) === "string") {
                                    data = JSON.parse(data);
                                }
                                if (data.code === 0) {
                                    $('#`+pop.Id+` .modal-body').html(data.data);
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
									$('#`+pop.Id+`').hide();
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
	up := template2.Default().Popup().SetID(pop.Id).
		SetTitle(template.HTML(pop.Title)).
		SetFooter(pop.BtnTitle).
		SetWidth(pop.Width).
		SetHeight(pop.Height).
		SetBody(template.HTML(``))

	if pop.Draggable {
		if pop.HideFooter {
			up = up.SetHideFooter()
		}
		return up.SetDraggable().GetContent()
	}

	return up.GetContent()
}

func isFormURL(s string) bool {
	reg, _ := regexp.Compile("(.*)info/(.*)/(new|edit)(.*?)")
	return reg.MatchString(s)
}
