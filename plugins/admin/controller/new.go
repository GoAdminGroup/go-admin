package controller

import (
	"fmt"
	template2 "html/template"
	"net/http"

	"github.com/GoAdminGroup/go-admin/template"

	"github.com/GoAdminGroup/go-admin/modules/logger"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// ShowNewForm show a new form page.
func (h *Handler) ShowNewForm(ctx *context.Context) {
	param := guard.GetShowNewFormParam(ctx)
	h.showNewForm(ctx, "", param.Prefix, param.Param.GetRouteParamStr(), false)
}

func (h *Handler) showNewForm(ctx *context.Context, alert template2.HTML, prefix, paramStr string, isNew bool) {

	var (
		user        = auth.Auth(ctx)
		panel       = h.table(prefix, ctx)
		formInfo    = panel.GetNewFormInfo()
		infoUrl     = h.routePathWithPrefix("info", prefix) + paramStr
		newUrl      = h.routePathWithPrefix("new", prefix)
		showNewUrl  = h.routePathWithPrefix("show_new", prefix) + paramStr
		referer     = ctx.Referer()
		f           = panel.GetActualNewForm()
		isNotIframe = ctx.Query(constant.IframeKey) != "true"
	)

	if referer != "" && !isInfoUrl(referer) && !isNewUrl(referer, ctx.Query(constant.PrefixKey)) {
		infoUrl = referer
	}

	hiddenFields := map[string]string{
		form2.TokenKey:    h.authSrv().AddToken(),
		form2.PreviousKey: infoUrl,
	}

	if ctx.Query(constant.IframeKey) != "" {
		hiddenFields[constant.IframeKey] = ctx.Query(constant.IframeKey)
	}

	if ctx.Query(constant.IframeIDKey) != "" {
		hiddenFields[constant.IframeIDKey] = ctx.Query(constant.IframeIDKey)
	}

	content := formContent(aForm().
		SetPrefix(h.config.PrefixFixSlash()).
		SetFieldsHTML(f.HTMLContent).
		SetContent(formInfo.FieldList).
		SetTabContents(formInfo.GroupFieldList).
		SetTabHeaders(formInfo.GroupFieldHeaders).
		SetUrl(newUrl).
		SetAjax(f.AjaxSuccessJS, f.AjaxErrorJS).
		SetInputWidth(f.InputWidth).
		SetHeadWidth(f.HeadWidth).
		SetLayout(f.Layout).
		SetPrimaryKey(panel.GetPrimaryKey().Name).
		SetHiddenFields(hiddenFields).
		SetTitle(f.FormNewTitle).
		SetOperationFooter(formFooter("new", f.IsHideContinueEditCheckBox, f.IsHideContinueNewCheckBox,
			f.IsHideResetButton, f.FormNewBtnWord)).
		SetHeader(f.HeaderHtml).
		SetFooter(f.FooterHtml), len(formInfo.GroupFieldHeaders) > 0, !isNotIframe, f.IsHideBackButton, f.Header)

	if f.Wrapper != nil {
		content = f.Wrapper(content)
	}

	h.HTML(ctx, user, types.Panel{
		Content:     alert + content,
		Description: template2.HTML(f.Description),
		Title:       modules.AorBHTML(isNotIframe, template2.HTML(f.Title), ""),
		MiniSidebar: f.HideSideBar,
	}, template.ExecuteOptions{Animation: alert == ""})

	if isNew {
		ctx.AddHeader(constant.PjaxUrlHeader, showNewUrl)
	}
}

// NewForm insert a table row into database.
func (h *Handler) NewForm(ctx *context.Context) {

	param := guard.GetNewFormParam(ctx)

	// process uploading files, only support local storage
	if len(param.MultiForm.File) > 0 {
		err := file.GetFileEngine(h.config.FileUploadEngine.Name).Upload(param.MultiForm)
		if err != nil {
			logger.Error("get file engine error: ", err)
			if ctx.WantJSON() {
				response.Error(ctx, err.Error())
			} else {
				h.showNewForm(ctx, aAlert().Warning(err.Error()), param.Prefix, param.Param.GetRouteParamStr(), true)
			}
			return
		}
	}

	err := param.Panel.InsertData(param.Value())
	if err != nil {
		logger.Error("insert data error: ", err)
		if ctx.WantJSON() {
			response.Error(ctx, err.Error(), map[string]interface{}{
				"token": h.authSrv().AddToken(),
			})
		} else {
			h.showNewForm(ctx, aAlert().Warning(err.Error()), param.Prefix, param.Param.GetRouteParamStr(), true)
		}
		return
	}

	f := param.Panel.GetActualNewForm()

	if f.Responder != nil {
		f.Responder(ctx)
		return
	}

	if ctx.WantJSON() && !param.IsIframe {
		response.OkWithData(ctx, map[string]interface{}{
			"url":   param.PreviousPath,
			"token": h.authSrv().AddToken(),
		})
		return
	}

	if !param.FromList {

		if isNewUrl(param.PreviousPath, param.Prefix) {
			h.showNewForm(ctx, param.Alert, param.Prefix, param.Param.GetRouteParamStr(), true)
			return
		}

		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>location.href="%s"</script>`, param.PreviousPath))
		ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
		return
	}

	if param.IsIframe {
		ctx.HTML(http.StatusOK, fmt.Sprintf(`<script>
		swal('%s', '', 'success');
		setTimeout(function(){
			$("#%s", window.parent.document).hide();
			$('.modal-backdrop.fade.in', window.parent.document).hide();
		}, 1000)
</script>`, language.Get("success"), param.IframeID))
		return
	}

	buf := h.showTable(ctx, param.Prefix, param.Param, nil)

	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, h.routePathWithPrefix("info", param.Prefix)+param.Param.GetRouteParamStr())
}
