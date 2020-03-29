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
	prefix := ctx.Query(constant.PrefixKey)
	id := ctx.Query(constant.DetailPKKey)
	panel := h.table(prefix, ctx)
	user := auth.Auth(ctx)

	newPanel := panel.Copy()

	formModel := newPanel.GetForm()

	var fieldList types.FieldList

	if len(panel.GetDetail().FieldList) == 0 {
		fieldList = panel.GetInfo().FieldList
	} else {
		fieldList = panel.GetDetail().FieldList
	}

	formModel.FieldList = make([]types.FormField, len(fieldList))

	for i, field := range fieldList {
		formModel.FieldList[i] = types.FormField{
			Field:        field.Field,
			TypeName:     field.TypeName,
			Head:         field.Head,
			Hide:         field.Hide,
			Join:         field.Join,
			FormType:     form.Default,
			FieldDisplay: field.FieldDisplay,
		}
	}

	param := parameter.GetParam(ctx.Request.URL,
		panel.GetInfo().DefaultPageSize,
		panel.GetInfo().SortField,
		panel.GetInfo().GetSort())

	paramStr := param.DeleteDetailPk().GetRouteParamStr()

	editUrl := modules.AorEmpty(panel.GetEditable(), h.routePathWithPrefix("show_edit", prefix)+paramStr+
		"&"+constant.EditPKKey+"="+ctx.Query(constant.DetailPKKey))
	deleteUrl := modules.AorEmpty(panel.GetDeletable(), h.routePathWithPrefix("delete", prefix)+paramStr)
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

</script>`, language.Get("are you sure to delete"), language.Get("yes"), language.Get("cancel"), deleteUrl, infoUrl, id)
	}

	title := panel.GetDetail().Title

	if title == "" {
		title = panel.GetInfo().Title + language.Get("Detail")
	}

	desc := panel.GetDetail().Description

	if desc == "" {
		desc = panel.GetInfo().Description + language.Get("Detail")
	}

	formInfo, err := newPanel.GetDataWithId(param.WithPKs(id))

	if err != nil {
		h.HTML(ctx, user, types.Panel{
			Content:     aAlert().Warning(err.Error()),
			Description: desc,
			Title:       title,
		}, param.Animation)
		return
	}

	h.HTML(ctx, user, types.Panel{
		Content: detailContent(aForm().
			SetTitle(template.HTML(title)).
			SetContent(formInfo.FieldList).
			SetFooter(template.HTML(deleteJs)).
			SetHiddenFields(map[string]string{
				form2.PreviousKey: infoUrl,
			}).
			SetPrefix(h.config.PrefixFixSlash()), editUrl, deleteUrl),
		Description: desc,
		Title:       title,
	}, param.Animation)
}
