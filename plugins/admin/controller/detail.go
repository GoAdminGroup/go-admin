package controller

import (
	"fmt"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func (h *Handler) ShowDetail(ctx *context.Context) {

	var (
		prefix    = ctx.Query(constant.PrefixKey)
		id        = ctx.Query(constant.DetailPKKey)
		panel     = h.table(prefix, ctx)
		user      = auth.Auth(ctx)
		newPanel  = panel.Copy()
		detail    = panel.GetDetail()
		info      = panel.GetInfo()
		formModel = newPanel.GetForm()
		fieldList = make(types.FieldList, 0)
	)

	if len(detail.FieldList) == 0 {
		fieldList = info.FieldList
	} else {
		fieldList = detail.FieldList
	}

	formModel.FieldList = make([]types.FormField, len(fieldList))

	for i, field := range fieldList {
		formModel.FieldList[i] = types.FormField{
			Field:        field.Field,
			FieldClass:   field.Field,
			TypeName:     field.TypeName,
			Head:         field.Head,
			Hide:         field.Hide,
			Joins:        field.Joins,
			FormType:     form.Default,
			FieldDisplay: field.FieldDisplay,
		}
	}

	if detail.Table != "" {
		formModel.Table = detail.Table
	} else {
		formModel.Table = info.Table
	}

	param := parameter.GetParam(ctx.Request.URL,
		info.DefaultPageSize,
		info.SortField,
		info.GetSort())

	paramStr := param.DeleteDetailPk().GetRouteParamStr()

	editUrl := modules.AorEmpty(!info.IsHideEditButton, h.routePathWithPrefix("show_edit", prefix)+paramStr+
		"&"+constant.EditPKKey+"="+ctx.Query(constant.DetailPKKey))
	deleteUrl := modules.AorEmpty(!info.IsHideDeleteButton, h.routePathWithPrefix("delete", prefix)+paramStr)
	infoUrl := h.routePathWithPrefix("info", prefix) + paramStr

	editUrl = user.GetCheckPermissionByUrlMethod(editUrl, h.route("show_edit").Method())
	deleteUrl = user.GetCheckPermissionByUrlMethod(deleteUrl, h.route("delete").Method())

	deleteJs := ""

	if deleteUrl != "" {
		deleteJs = fmt.Sprintf(`<script>
function DeletePost(id) {
	swal({
			title: '%s',
			type: "warning",
			showCancelButton: true,
			confirmButtonColor: "#DD6B55",
			confirmButtonText: '%s',
			closeOnConfirm: false,
			cancelButtonText: '%s',
		},
		function () {
			$.ajax({
				method: 'post',
				url: '%s',
				data: {
					id: id
				},
				success: function (data) {
					if (typeof (data) === "string") {
						data = JSON.parse(data);
					}
					if (data.code === 200) {
						location.href = '%s'
					} else {
						swal(data.msg, '', 'error');
					}
				}
			});
		});
}

$('.delete-btn').on('click', function (event) {
	DeletePost(%s)
});

</script>`, language.Get("are you sure to delete"), language.Get("yes"),
			language.Get("cancel"), deleteUrl, infoUrl, id)
	}

	title := ""
	desc := ""

	isNotIframe := ctx.Query(constant.IframeKey) != "true"

	if isNotIframe {
		title = detail.Title

		if title == "" {
			title = info.Title + language.Get("Detail")
		}

		desc = detail.Description

		if desc == "" {
			desc = info.Description + language.Get("Detail")
		}
	}

	formInfo, err := newPanel.GetDataWithId(param.WithPKs(id))

	if err != nil {
		h.HTML(ctx, user, template.WarningPanelWithDescAndTitle(err.Error(), desc, title),
			template.ExecuteOptions{Animation: param.Animation})
		return
	}

	h.HTML(ctx, user, types.Panel{
		Content: detailContent(aForm().
			SetTitle(template.HTML(title)).
			SetContent(formInfo.FieldList).
			SetHeader(detail.HeaderHtml).
			SetFooter(template.HTML(deleteJs)+detail.FooterHtml).
			SetHiddenFields(map[string]string{
				form2.PreviousKey: infoUrl,
			}).
			SetPrefix(h.config.PrefixFixSlash()), editUrl, deleteUrl, !isNotIframe),
		Description: template.HTML(desc),
		Title:       template.HTML(title),
	}, template.ExecuteOptions{Animation: param.Animation})
}
