package types

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	form2 "github.com/GoAdminGroup/go-admin/template/types/form"
	"html"
	"html/template"
	"strconv"
	"strings"
)

type FieldOptions []map[string]string

func (fo FieldOptions) SetSelected(val interface{}, labels []string) FieldOptions {

	if valArr, ok := val.([]string); ok {
		for _, v := range fo {
			if utils.InArray(valArr, v["value"]) || utils.InArray(valArr, v["field"]) {
				v["selected"] = labels[0]
			} else {
				v["selected"] = labels[1]
			}
		}
	} else {
		for _, v := range fo {
			if v["value"] == val || v["field"] == val {
				v["selected"] = labels[0]
			} else {
				v["selected"] = labels[1]
			}
		}
	}

	return fo
}

// FormField is the form field with different options.
type FormField struct {
	Field    string
	TypeName db.DatabaseType
	Head     string
	FormType form2.Type

	Default                template.HTML
	Value                  template.HTML
	Value2                 string
	Options                FieldOptions
	DefaultOptionDelimiter string
	Label                  template.HTML

	CustomContent template.HTML
	CustomJs      template.JS
	CustomCss     template.CSS

	Editable    bool
	NotAllowAdd bool
	Must        bool
	Hide        bool

	HelpMsg   template.HTML
	OptionExt template.JS

	FieldDisplay
	PostFilterFn PostFieldFilterFn
}

func (f FormField) UpdateValue(id, val string, res map[string]interface{}) FormField {
	if f.FormType.IsSelect() {
		f.Options.SetSelected(f.ToDisplay(FieldModel{
			ID:    id,
			Value: val,
			Row:   res,
		}), f.FormType.SelectedLabel())
	} else {
		value := f.ToDisplay(FieldModel{
			ID:    id,
			Value: val,
			Row:   res,
		})
		if v, ok := value.(template.HTML); ok {
			f.Value = v
		} else {
			f.Value = template.HTML(value.(string))
		}
	}
	return f
}

func (f FormField) UpdateDefaultValue() FormField {
	f.Value = f.Default
	if f.FormType.IsSelect() {
		f.Options.SetSelected(string(f.Value), f.FormType.SelectedLabel())
	}
	return f
}

// FormPanel
type FormPanel struct {
	FieldList         FormFields
	curFieldListIndex int

	// Warn: may be deprecated in the future.
	TabGroups  TabGroups
	TabHeaders TabHeaders

	Table       string
	Title       string
	Description string

	Validator    FormPostFn
	PostHook     FormPostFn
	PreProcessFn FormPreProcessFn

	Callbacks Callbacks

	UpdateFn FormPostFn
	InsertFn FormPostFn

	processChains DisplayProcessFnChains

	HeaderHtml template.HTML
	FooterHtml template.HTML
}

func NewFormPanel() *FormPanel {
	return &FormPanel{
		curFieldListIndex: -1,
		Callbacks:         make(Callbacks, 0),
	}
}

func (f *FormPanel) AddLimitFilter(limit int) *FormPanel {
	f.processChains = addLimit(limit, f.processChains)
	return f
}

func (f *FormPanel) AddTrimSpaceFilter() *FormPanel {
	f.processChains = addTrimSpace(f.processChains)
	return f
}

func (f *FormPanel) AddSubstrFilter(start int, end int) *FormPanel {
	f.processChains = addSubstr(start, end, f.processChains)
	return f
}

func (f *FormPanel) AddToTitleFilter() *FormPanel {
	f.processChains = addToTitle(f.processChains)
	return f
}

func (f *FormPanel) AddToUpperFilter() *FormPanel {
	f.processChains = addToUpper(f.processChains)
	return f
}

func (f *FormPanel) AddToLowerFilter() *FormPanel {
	f.processChains = addToLower(f.processChains)
	return f
}

func (f *FormPanel) AddXssFilter() *FormPanel {
	f.processChains = addXssFilter(f.processChains)
	return f
}

func (f *FormPanel) AddXssJsFilter() *FormPanel {
	f.processChains = addXssJsFilter(f.processChains)
	return f
}

func (f *FormPanel) AddField(head, field string, filedType db.DatabaseType, formType form2.Type) *FormPanel {
	f.FieldList = append(f.FieldList, FormField{
		Head:     head,
		Field:    field,
		TypeName: filedType,
		Editable: true,
		Hide:     false,
		FormType: formType,
		FieldDisplay: FieldDisplay{
			Display: func(value FieldModel) interface{} {
				return value.Value
			},
			DisplayProcessChains: chooseDisplayProcessChains(f.processChains),
		},
	})
	f.curFieldListIndex++
	return f
}

// Field attribute setting functions
// ====================================================

func (f *FormPanel) FieldDisplay(filter FieldFilterFn) *FormPanel {
	f.FieldList[f.curFieldListIndex].Display = filter
	return f
}

func (f *FormPanel) SetTable(table string) *FormPanel {
	f.Table = table
	return f
}

func (f *FormPanel) FieldMust() *FormPanel {
	f.FieldList[f.curFieldListIndex].Must = true
	return f
}

func (f *FormPanel) FieldHide() *FormPanel {
	f.FieldList[f.curFieldListIndex].Hide = true
	return f
}

func (f *FormPanel) FieldHelpMsg(s template.HTML) *FormPanel {
	f.FieldList[f.curFieldListIndex].HelpMsg = s
	return f
}

func (f *FormPanel) FieldOptionExt(m map[string]interface{}) *FormPanel {
	s, _ := json.Marshal(m)

	if f.FieldList[f.curFieldListIndex].OptionExt != template.JS("") {
		ss := string(f.FieldList[f.curFieldListIndex].OptionExt)
		ss = strings.Replace(ss, "}", ",", strings.Count(ss, "}"))
		ss = strings.TrimRight(ss, " ")
		ss += ","
		f.FieldList[f.curFieldListIndex].OptionExt = template.JS(ss) + template.JS(strings.Replace(string(s), "{", "", 1))
	} else {
		f.FieldList[f.curFieldListIndex].OptionExt = template.JS(string(s))
	}

	return f
}

func (f *FormPanel) FieldOptionExtJS(js template.JS) *FormPanel {
	f.FieldList[f.curFieldListIndex].OptionExt = js
	return f
}

func (f *FormPanel) FieldDefault(def string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Default = template.HTML(def)
	return f
}

func (f *FormPanel) FieldNotAllowEdit() *FormPanel {
	f.FieldList[f.curFieldListIndex].Editable = false
	return f
}

func (f *FormPanel) FieldNotAllowAdd() *FormPanel {
	f.FieldList[f.curFieldListIndex].NotAllowAdd = true
	return f
}

func (f *FormPanel) FieldFormType(formType form2.Type) *FormPanel {
	f.FieldList[f.curFieldListIndex].FormType = formType
	return f
}

func (f *FormPanel) FieldValue(value string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Value = template.HTML(value)
	return f
}

func (f *FormPanel) FieldOptions(options []map[string]string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Options = options
	return f
}

func (f *FormPanel) FieldDefaultOptionDelimiter(delimiter string) *FormPanel {
	f.FieldList[f.curFieldListIndex].DefaultOptionDelimiter = delimiter
	return f
}

func (f *FormPanel) FieldPostFilterFn(post PostFieldFilterFn) *FormPanel {
	f.FieldList[f.curFieldListIndex].PostFilterFn = post
	return f
}

func (f *FormPanel) FieldLimit(limit int) *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddLimit(limit)
	return f
}

func (f *FormPanel) FieldTrimSpace() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddTrimSpace()
	return f
}

func (f *FormPanel) FieldSubstr(start int, end int) *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddSubstr(start, end)
	return f
}

func (f *FormPanel) FieldToTitle() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddToTitle()
	return f
}

func (f *FormPanel) FieldToUpper() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddToUpper()
	return f
}

func (f *FormPanel) FieldToLower() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].AddToLower()
	return f
}

func (f *FormPanel) FieldXssFilter() *FormPanel {
	f.FieldList[f.curFieldListIndex].DisplayProcessChains = f.FieldList[f.curFieldListIndex].DisplayProcessChains.
		Add(func(s string) string {
			return html.EscapeString(s)
		})
	return f
}

func (f *FormPanel) FieldCustomContent(content template.HTML) *FormPanel {
	f.FieldList[f.curFieldListIndex].CustomContent = content
	return f
}

func (f *FormPanel) FieldCustomJs(js template.JS) *FormPanel {
	f.FieldList[f.curFieldListIndex].CustomJs = js
	return f
}

func (f *FormPanel) FieldCustomCss(css template.CSS) *FormPanel {
	f.FieldList[f.curFieldListIndex].CustomCss = css
	return f
}

func (f *FormPanel) FieldOnSearch(url string, handler Handler, delay ...int) *FormPanel {

	delayStr := "500"
	if len(delay) > 0 {
		delayStr = strconv.Itoa(delay[0])
	}

	if f.FieldList[f.curFieldListIndex].OptionExt != template.JS("") {
		s := string(f.FieldList[f.curFieldListIndex].OptionExt)
		s = strings.Replace(s, "{", "", 1)
		s = utils.ReplaceNth(s, "}", "", strings.Count(s, "}"))
		s = strings.TrimRight(s, " ")
		s += ","
		f.FieldList[f.curFieldListIndex].OptionExt = template.JS(s)
	}

	f.FieldList[f.curFieldListIndex].OptionExt = `{
		` + f.FieldList[f.curFieldListIndex].OptionExt + template.JS(`
		ajax: {
		    url: "` + url + `",
		    dataType: 'json',
		    data: function (params) {
			      var query = {
			        	search: params.term,
						page: params.page || 1
			      }
			      return query;
		    },
		    delay: ` + delayStr + `,
		    processResults: function (data, params) {
			      return data.data;
	    	}
	  	}
	}`)

	f.Callbacks = append(f.Callbacks, context.Node{
		Path:     url,
		Method:   "get",
		Handlers: context.Handlers{handler.Wrap()},
		Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
	})

	return f
}

func (f *FormPanel) FieldOnChooseCustom(js template.HTML) *FormPanel {
	f.FooterHtml += `<script>
$(".` + template.HTML(f.FieldList[f.curFieldListIndex].Field) + `").on("select2:select",function(e){
	` + js + `
})
</script>`
	return f
}

type LinkField struct {
	Field   string
	Value   template.HTML
	Hide    bool
	Disable bool
}

func (f *FormPanel) FieldOnChooseMap(m map[string]LinkField) *FormPanel {

	cm := template.HTML("")

	for val, obejct := range m {
		if obejct.Hide {
			cm += `if (e.params.data.text === "` + template.HTML(val) + `") {
		$("label[for='` + template.HTML(obejct.Field) + `']").parent().hide()
	} else {
		$("label[for='` + template.HTML(obejct.Field) + `']").parent().show()
	}`
		} else if obejct.Disable {
			cm += `if (e.params.data.text === "` + template.HTML(val) + `") {
		$("#` + template.HTML(obejct.Field) + `").prop('disabled', true);
	} else {
		$("#` + template.HTML(obejct.Field) + `").prop('disabled', false);
	}`
		} else {
			cm += `if (e.params.data.text === "` + template.HTML(val) + `") {
		if ($(".` + template.HTML(obejct.Field) + `").length > 0) {
			$(".` + template.HTML(obejct.Field) + `").val("` + obejct.Value + `").select2()
		} else {
			$("#` + template.HTML(obejct.Field) + `").val("` + obejct.Value + `")
		}	
	}`
		}
	}

	f.FooterHtml += `<script>
$(".` + template.HTML(f.FieldList[f.curFieldListIndex].Field) + `").on("select2:select",function(e){
	` + cm + `
})
</script>`
	return f
}

func (f *FormPanel) FieldOnChoose(val, field string, value template.HTML) *FormPanel {
	f.FooterHtml += `<script>
$(".` + template.HTML(f.FieldList[f.curFieldListIndex].Field) + `").on("select2:select",function(e){
	if (e.params.data.text === "` + template.HTML(val) + `") {
		if ($(".` + template.HTML(field) + `").length > 0) {
			$(".` + template.HTML(field) + `").val("` + value + `").select2()
		} else {
			$("#` + template.HTML(field) + `").val("` + value + `")
		}	
	}
})
</script>`
	return f
}

func (f *FormPanel) FieldOnChooseAjax(field, url string, handler Handler) *FormPanel {

	f.FooterHtml += `<script>
$(".` + template.HTML(f.FieldList[f.curFieldListIndex].Field) + `").on("select2:select",function(e){
	let id = '` + template.HTML(field) + `'
	let selectObj = $("."+id)
	if (selectObj.length > 0) {
		selectObj.val("").select2()
		selectObj.html('<option value="" selected="selected"></option>')
	}
	$.ajax({
		url:"` + template.HTML(url) + `",
		type: 'post',
		dataType: 'text',
		data: {
			'value':$(this).val()
		},
		success: function (data)  {
			if (typeof (data) === "string") {
				data = JSON.parse(data);
			}
			if (data.code === 0) {
				if (selectObj.length > 0) {
					if (typeof(data.data) === "object") {
						$('.' + id).select2({
							data: data.data
						});
					} else {
						$('.' + id).val(data.data).select2()
					}
				} else {
					$('#` + template.HTML(field) + `').val(data.data);
				}
			} else {
				swal(data.msg, '', 'error');
			}
		},
		error:function(){
			alert('error')
		}
	})
})
</script>`

	f.Callbacks = append(f.Callbacks, context.Node{
		Path:     url,
		Method:   "post",
		Handlers: context.Handlers{handler.Wrap()},
		Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
	})

	return f
}

func (f *FormPanel) FieldOnChooseHide(value string, field ...string) *FormPanel {

	if len(field) == 0 {
		return f
	}

	hideText := template.HTML("")
	showText := template.HTML("")

	for _, f := range field {
		hideText += `$("label[for='` + template.HTML(f) + `']").parent().hide()
`
		showText += `$("label[for='` + template.HTML(f) + `']").parent().show()
`
	}

	f.FooterHtml += `<script>
$(".` + template.HTML(f.FieldList[f.curFieldListIndex].Field) + `").on("select2:select",function(e){
	if (e.params.data.text === "` + template.HTML(value) + `") {
		` + hideText + `
	} else {
		` + showText + `
	}
})
</script>`
	return f
}

func (f *FormPanel) FieldOnChooseDisable(value string, field ...string) *FormPanel {

	if len(field) == 0 {
		return f
	}

	disableText := template.HTML("")
	enableText := template.HTML("")

	for _, f := range field {
		disableText += `$("#` + template.HTML(f) + `").prop('disabled', true);
`
		enableText += `$("#` + template.HTML(f) + `").prop('disabled', false);
`
	}

	f.FooterHtml += `<script>
$(".` + template.HTML(f.FieldList[f.curFieldListIndex].Field) + `").on("select2:select",function(e){
	if (e.params.data.text === "` + template.HTML(value) + `") {
		` + disableText + `
	} else {
		` + enableText + `
	}
})
</script>`
	return f
}

// FormPanel attribute setting functions
// ====================================================

func (f *FormPanel) SetTitle(title string) *FormPanel {
	f.Title = title
	return f
}

func (f *FormPanel) SetTabGroups(groups TabGroups) *FormPanel {
	f.TabGroups = groups
	return f
}

func (f *FormPanel) SetTabHeaders(headers ...string) *FormPanel {
	f.TabHeaders = headers
	return f
}

func (f *FormPanel) SetDescription(desc string) *FormPanel {
	f.Description = desc
	return f
}

func (f *FormPanel) SetHeaderHtml(header template.HTML) *FormPanel {
	f.HeaderHtml += header
	return f
}

func (f *FormPanel) SetFooterHtml(footer template.HTML) *FormPanel {
	f.FooterHtml += footer
	return f
}

func (f *FormPanel) SetPostValidator(va FormPostFn) *FormPanel {
	f.Validator = va
	return f
}

func (f *FormPanel) SetPreProcessFn(fn FormPreProcessFn) *FormPanel {
	f.PreProcessFn = fn
	return f
}

func (f *FormPanel) SetPostHook(fn FormPostFn) *FormPanel {
	f.PostHook = fn
	return f
}

func (f *FormPanel) SetUpdateFn(fn FormPostFn) *FormPanel {
	f.UpdateFn = fn
	return f
}

func (f *FormPanel) SetInsertFn(fn FormPostFn) *FormPanel {
	f.InsertFn = fn
	return f
}

type FormPreProcessFn func(values form.Values) form.Values

type FormPostFn func(values form.Values) error

type FormFields []FormField

func (f FormFields) Copy() FormFields {
	formList := make(FormFields, len(f))
	copy(formList, f)
	for i := 0; i < len(formList); i++ {
		formList[i].Options = make([]map[string]string, len(f[i].Options))
		for j := 0; j < len(f[i].Options); j++ {
			formList[i].Options[j] = utils.CopyMap(f[i].Options[j])
		}
	}
	return formList
}

func (f FormFields) FindByFieldName(field string) FormField {
	for i := 0; i < len(f); i++ {
		if f[i].Field == field {
			return f[i]
		}
	}
	return FormField{}
}
