package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	form2 "github.com/GoAdminGroup/go-admin/template/types/form"
)

type FieldOption struct {
	Text          string            `json:"text"`
	Value         string            `json:"value"`
	TextHTML      template.HTML     `json:"-"`
	Selected      bool              `json:"-"`
	SelectedLabel template.HTML     `json:"-"`
	Extra         map[string]string `json:"-"`
}

type FieldOptions []FieldOption

func (fo FieldOptions) Copy() FieldOptions {
	newOptions := make(FieldOptions, len(fo))
	copy(newOptions, fo)
	return newOptions
}

func (fo FieldOptions) SetSelected(val interface{}, labels []template.HTML) FieldOptions {

	if valArr, ok := val.([]string); ok {
		for k := range fo {
			text := fo[k].Text
			if text == "" {
				text = string(fo[k].TextHTML)
			}
			fo[k].Selected = utils.InArray(valArr, fo[k].Value) || utils.InArray(valArr, text)
			if fo[k].Selected {
				fo[k].SelectedLabel = labels[0]
			} else {
				fo[k].SelectedLabel = labels[1]
			}
		}
	} else {
		for k := range fo {
			text := fo[k].Text
			if text == "" {
				text = string(fo[k].TextHTML)
			}
			fo[k].Selected = fo[k].Value == val || text == val
			if fo[k].Selected {
				fo[k].SelectedLabel = labels[0]
			} else {
				fo[k].SelectedLabel = labels[1]
			}
		}
	}

	return fo
}

func (fo FieldOptions) SetSelectedLabel(labels []template.HTML) FieldOptions {
	for k := range fo {
		if fo[k].Selected {
			fo[k].SelectedLabel = labels[0]
		} else {
			fo[k].SelectedLabel = labels[1]
		}
	}
	return fo
}

func (fo FieldOptions) Marshal() string {
	if len(fo) == 0 {
		return ""
	}
	eo, err := json.Marshal(fo)

	if err != nil {
		return ""
	}

	return string(eo)
}

type (
	OptionInitFn              func(val FieldModel) FieldOptions
	OptionArrInitFn           func(val FieldModel) []FieldOptions
	OptionTableQueryProcessFn func(sql *db.SQL) *db.SQL
	OptionProcessFn           func(options FieldOptions) FieldOptions

	OptionTable struct {
		Table          string
		TextField      string
		ValueField     string
		QueryProcessFn OptionTableQueryProcessFn
		ProcessFn      OptionProcessFn
	}
)

// FormField is the form field with different options.
type FormField struct {
	Field          string          `json:"field"`
	FieldClass     string          `json:"field_class"`
	TypeName       db.DatabaseType `json:"type_name"`
	Head           string          `json:"head"`
	Foot           template.HTML   `json:"foot"`
	FormType       form2.Type      `json:"form_type"`
	FatherFormType form2.Type      `json:"father_form_type"`
	FatherField    string          `json:"father_field"`

	RowWidth int
	RowFlag  uint8

	Default                template.HTML  `json:"default"`
	DefaultArr             interface{}    `json:"default_arr"`
	Value                  template.HTML  `json:"value"`
	Value2                 string         `json:"value_2"`
	ValueArr               []string       `json:"value_arr"`
	Value2Arr              []string       `json:"value_2_arr"`
	Options                FieldOptions   `json:"options"`
	OptionsArr             []FieldOptions `json:"options_arr"`
	DefaultOptionDelimiter string         `json:"default_option_delimiter"`
	Label                  template.HTML  `json:"label"`
	HideLabel              bool           `json:"hide_label"`

	Placeholder string `json:"placeholder"`

	CustomContent template.HTML `json:"custom_content"`
	CustomJs      template.JS   `json:"custom_js"`
	CustomCss     template.CSS  `json:"custom_css"`

	Editable    bool `json:"editable"`
	NotAllowAdd bool `json:"not_allow_add"`
	Must        bool `json:"must"`
	Hide        bool `json:"hide"`

	Width int `json:"width"`

	InputWidth int `json:"input_width"`
	HeadWidth  int `json:"head_width"`

	Joins Joins `json:"-"`

	Divider      bool   `json:"divider"`
	DividerTitle string `json:"divider_title"`

	HelpMsg template.HTML `json:"help_msg"`

	TableFields FormFields

	OptionExt       template.JS     `json:"option_ext"`
	OptionExt2      template.JS     `json:"option_ext_2"`
	OptionInitFn    OptionInitFn    `json:"-"`
	OptionArrInitFn OptionArrInitFn `json:"-"`
	OptionTable     OptionTable     `json:"-"`

	FieldDisplay `json:"-"`
	PostFilterFn PostFieldFilterFn `json:"-"`
}

func (f *FormField) GetRawValue(columns []string, v interface{}) string {
	isJSON := len(columns) == 0
	return modules.AorB(isJSON || modules.InArray(columns, f.Field),
		db.GetValueFromDatabaseType(f.TypeName, v, isJSON).String(), "")
}

func (f *FormField) UpdateValue(id, val string, res map[string]interface{}, sql *db.SQL) *FormField {
	return f.updateValue(id, val, res, PostTypeUpdate, sql)
}

func (f *FormField) UpdateDefaultValue(sql *db.SQL) *FormField {
	f.Value = f.Default
	return f.updateValue("", string(f.Value), make(map[string]interface{}), PostTypeCreate, sql)
}

func (f *FormField) setOptionsFromSQL(sql *db.SQL) {
	if sql != nil && f.OptionTable.Table != "" && len(f.Options) == 0 {

		sql.Table(f.OptionTable.Table).Select(f.OptionTable.ValueField, f.OptionTable.TextField)

		if f.OptionTable.QueryProcessFn != nil {
			f.OptionTable.QueryProcessFn(sql)
		}

		queryRes, err := sql.All()
		if err == nil {
			for _, item := range queryRes {
				f.Options = append(f.Options, FieldOption{
					Value: fmt.Sprintf("%v", item[f.OptionTable.ValueField]),
					Text:  fmt.Sprintf("%v", item[f.OptionTable.TextField]),
				})
			}
		}

		if f.OptionTable.ProcessFn != nil {
			f.Options = f.OptionTable.ProcessFn(f.Options)
		}
	}
}

func (f *FormField) isBelongToATable() bool {
	return f.FatherField != "" && f.FatherFormType.IsTable()
}

func (f *FormField) isNotBelongToATable() bool {
	return f.FatherField == "" && !f.FatherFormType.IsTable()
}

func (f *FormField) allowAdd() bool {
	return !f.NotAllowAdd
}

func (f *FormField) updateValue(id, val string, res map[string]interface{}, typ PostType, sql *db.SQL) *FormField {

	m := FieldModel{
		ID:       id,
		Value:    val,
		Row:      res,
		PostType: typ,
	}

	if f.isBelongToATable() {
		if f.FormType.IsSelect() {
			if len(f.OptionsArr) == 0 && f.OptionArrInitFn != nil {
				f.OptionsArr = f.OptionArrInitFn(m)
				for i := 0; i < len(f.OptionsArr); i++ {
					f.OptionsArr[i] = f.OptionsArr[i].SetSelectedLabel(f.FormType.SelectedLabel())
				}
			} else {

				f.setOptionsFromSQL(sql)

				if f.FormType.IsSingleSelect() {
					values := f.ToDisplayStringArray(m)
					f.OptionsArr = make([]FieldOptions, len(values))
					for k, value := range values {
						f.OptionsArr[k] = f.Options.Copy().SetSelected(value, f.FormType.SelectedLabel())
					}
				} else {
					values := f.ToDisplayStringArrayArray(m)
					f.OptionsArr = make([]FieldOptions, len(values))
					for k, value := range values {
						f.OptionsArr[k] = f.Options.Copy().SetSelected(value, f.FormType.SelectedLabel())
					}
				}
			}
		} else {
			f.ValueArr = f.ToDisplayStringArray(m)
		}
	} else {
		if f.FormType.IsSelect() {
			if len(f.Options) == 0 && f.OptionInitFn != nil {
				f.Options = f.OptionInitFn(m).SetSelectedLabel(f.FormType.SelectedLabel())
			} else {
				f.setOptionsFromSQL(sql)
				f.Options.SetSelected(f.ToDisplay(m), f.FormType.SelectedLabel())
			}
		} else if f.FormType.IsArray() {
			f.ValueArr = f.ToDisplayStringArray(m)
		} else {
			f.Value = f.ToDisplayHTML(m)
			if f.FormType.IsFile() {
				if f.Value != template.HTML("") {
					f.Value2 = config.GetStore().URL(string(f.Value))
				}
			}
		}
	}

	return f
}

func (f *FormField) FillCustomContent() *FormField {
	// TODO: optimize
	if f.CustomContent != "" {
		f.CustomContent = template.HTML(f.fillCustom(string(f.CustomContent)))
	}
	if f.CustomJs != "" {
		f.CustomJs = template.JS(f.fillCustom(string(f.CustomJs)))
	}
	if f.CustomCss != "" {
		f.CustomCss = template.CSS(f.fillCustom(string(f.CustomCss)))
	}
	return f
}

func (f *FormField) fillCustom(src string) string {
	t := template.New("custom")
	t, err := t.Parse(src)
	if err != nil {
		logger.Error(err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, f)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return buf.String()
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

	primaryKey primaryKey

	UpdateFn FormPostFn
	InsertFn FormPostFn

	IsHideContinueEditCheckBox bool
	IsHideContinueNewCheckBox  bool
	IsHideResetButton          bool
	IsHideBackButton           bool

	Layout form2.Layout

	HTMLContent template.HTML

	Header template.HTML

	InputWidth int
	HeadWidth  int

	FormNewTitle    template.HTML
	FormNewBtnWord  template.HTML
	FormEditTitle   template.HTML
	FormEditBtnWord template.HTML

	Ajax          bool
	AjaxSuccessJS template.JS
	AjaxErrorJS   template.JS

	Responder Responder

	Wrapper ContentWrapper

	processChains DisplayProcessFnChains

	HeaderHtml template.HTML
	FooterHtml template.HTML
}

type Responder func(ctx *context.Context)

func NewFormPanel() *FormPanel {
	return &FormPanel{
		curFieldListIndex: -1,
		Callbacks:         make(Callbacks, 0),
		Layout:            form2.LayoutDefault,
		FormNewTitle:      "New",
		FormEditTitle:     "Edit",
		FormNewBtnWord:    language.GetFromHtml("Save"),
		FormEditBtnWord:   language.GetFromHtml("Save"),
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

func (f *FormPanel) SetPrimaryKey(name string, typ db.DatabaseType) *FormPanel {
	f.primaryKey = primaryKey{Name: name, Type: typ}
	return f
}

func (f *FormPanel) HideContinueEditCheckBox() *FormPanel {
	f.IsHideContinueEditCheckBox = true
	return f
}

func (f *FormPanel) HideContinueNewCheckBox() *FormPanel {
	f.IsHideContinueNewCheckBox = true
	return f
}

func (f *FormPanel) HideResetButton() *FormPanel {
	f.IsHideResetButton = true
	return f
}

func (f *FormPanel) HideBackButton() *FormPanel {
	f.IsHideBackButton = true
	return f
}

func (f *FormPanel) AddField(head, field string, filedType db.DatabaseType, formType form2.Type) *FormPanel {

	f.FieldList = append(f.FieldList, FormField{
		Head:        head,
		Field:       field,
		FieldClass:  field,
		TypeName:    filedType,
		Editable:    true,
		Hide:        false,
		TableFields: make(FormFields, 0),
		Placeholder: language.Get("input") + " " + head,
		FormType:    formType,
		FieldDisplay: FieldDisplay{
			Display: func(value FieldModel) interface{} {
				return value.Value
			},
			DisplayProcessChains: chooseDisplayProcessChains(f.processChains),
		},
	})
	f.curFieldListIndex++

	// Set default options of different form type
	op1, op2, js := formType.GetDefaultOptions(field)
	f.FieldOptionExt(op1)
	f.FieldOptionExt2(op2)
	f.FieldOptionExtJS(js)

	// Set default Display Filter Function of different form type
	setDefaultDisplayFnOfFormType(f, formType)

	return f
}

type AddFormFieldFn func(panel *FormPanel)

func (f *FormPanel) AddTable(head, field string, addFields AddFormFieldFn) *FormPanel {
	index := f.curFieldListIndex
	addFields(f)
	for i := index + 1; i <= f.curFieldListIndex; i++ {
		f.FieldList[i].FatherFormType = form2.Table
		f.FieldList[i].FatherField = field
	}
	fields := make(FormFields, f.curFieldListIndex-index)
	copy(fields, f.FieldList[index+1:f.curFieldListIndex+1])
	f.FieldList = append(f.FieldList, FormField{
		Head:        head,
		Field:       field,
		FieldClass:  field,
		TypeName:    db.Varchar,
		Editable:    true,
		Hide:        false,
		TableFields: fields,
		FormType:    form2.Table,
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

func (f *FormPanel) AddRow(addFields AddFormFieldFn) *FormPanel {
	index := f.curFieldListIndex
	addFields(f)
	if f.curFieldListIndex != index+1 {
		for i := index + 1; i <= f.curFieldListIndex; i++ {
			if i == index+1 {
				f.FieldList[i].RowFlag = 1
			} else if i == f.curFieldListIndex {
				f.FieldList[i].RowFlag = 2
			} else {
				f.FieldList[i].RowFlag = 3
			}
		}
	}
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

func (f *FormPanel) FieldPlaceholder(placeholder string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Placeholder = placeholder
	return f
}

func (f *FormPanel) FieldWidth(width int) *FormPanel {
	f.FieldList[f.curFieldListIndex].Width = width
	return f
}

func (f *FormPanel) FieldInputWidth(width int) *FormPanel {
	f.FieldList[f.curFieldListIndex].InputWidth = width
	return f
}

func (f *FormPanel) FieldHeadWidth(width int) *FormPanel {
	f.FieldList[f.curFieldListIndex].HeadWidth = width
	return f
}

func (f *FormPanel) FieldRowWidth(width int) *FormPanel {
	f.FieldList[f.curFieldListIndex].RowWidth = width
	return f
}

func (f *FormPanel) FieldHideLabel() *FormPanel {
	f.FieldList[f.curFieldListIndex].HideLabel = true
	return f
}

func (f *FormPanel) FieldFoot(foot template.HTML) *FormPanel {
	f.FieldList[f.curFieldListIndex].Foot = foot
	return f
}

func (f *FormPanel) FieldDivider(title ...string) *FormPanel {
	f.FieldList[f.curFieldListIndex].Divider = true
	if len(title) > 0 {
		f.FieldList[f.curFieldListIndex].DividerTitle = title[0]
	}
	return f
}

func (f *FormPanel) FieldHelpMsg(s template.HTML) *FormPanel {
	f.FieldList[f.curFieldListIndex].HelpMsg = s
	return f
}

func (f *FormPanel) FieldOptionInitFn(fn OptionInitFn) *FormPanel {
	f.FieldList[f.curFieldListIndex].OptionInitFn = fn
	return f
}

func (f *FormPanel) FieldOptionExt(m map[string]interface{}) *FormPanel {

	if m == nil {
		return f
	}

	if f.FieldList[f.curFieldListIndex].FormType.IsCode() {
		f.FieldList[f.curFieldListIndex].OptionExt = template.JS(fmt.Sprintf(`
	theme = "%s";
	font_size = %s;
	language = "%s";
	options = %s;
`, m["theme"], m["font_size"], m["language"], m["options"]))
		return f
	}

	m = f.FieldList[f.curFieldListIndex].FormType.FixOptions(m)

	s, _ := json.Marshal(m)

	if f.FieldList[f.curFieldListIndex].OptionExt != template.JS("") {
		ss := string(f.FieldList[f.curFieldListIndex].OptionExt)
		ss = strings.Replace(ss, "}", "", strings.Count(ss, "}"))
		ss = strings.TrimRight(ss, " ")
		ss += ","
		f.FieldList[f.curFieldListIndex].OptionExt = template.JS(ss) + template.JS(strings.Replace(string(s), "{", "", 1))
	} else {
		f.FieldList[f.curFieldListIndex].OptionExt = template.JS(string(s))
	}

	return f
}

func (f *FormPanel) FieldOptionExt2(m map[string]interface{}) *FormPanel {

	if m == nil {
		return f
	}

	m = f.FieldList[f.curFieldListIndex].FormType.FixOptions(m)

	s, _ := json.Marshal(m)

	if f.FieldList[f.curFieldListIndex].OptionExt2 != template.JS("") {
		ss := string(f.FieldList[f.curFieldListIndex].OptionExt2)
		ss = strings.Replace(ss, "}", "", strings.Count(ss, "}"))
		ss = strings.TrimRight(ss, " ")
		ss += ","
		f.FieldList[f.curFieldListIndex].OptionExt2 = template.JS(ss) + template.JS(strings.Replace(string(s), "{", "", 1))
	} else {
		f.FieldList[f.curFieldListIndex].OptionExt2 = template.JS(string(s))
	}

	return f
}

func (f *FormPanel) FieldOptionExtJS(js template.JS) *FormPanel {
	if js != template.JS("") {
		f.FieldList[f.curFieldListIndex].OptionExt = js
	}
	return f
}

func (f *FormPanel) FieldOptionExtJS2(js template.JS) *FormPanel {
	f.FieldList[f.curFieldListIndex].OptionExt2 = js
	return f
}

func (f *FormPanel) FieldEnableFileUpload(data ...interface{}) *FormPanel {

	url := f.OperationURL("/file/upload")

	if len(data) > 0 {
		url = data[0].(string)
	}

	field := f.FieldList[f.curFieldListIndex].Field

	f.FieldList[f.curFieldListIndex].OptionExt = template.JS(fmt.Sprintf(`
	%seditor.customConfig.uploadImgServer = '%s';
	%seditor.customConfig.uploadImgMaxSize = 3 * 1024 * 1024;
	%seditor.customConfig.uploadImgMaxLength = 5;
	%seditor.customConfig.uploadFileName = 'file';
`, field, url, field, field, field))

	var fileUploadHandler context.Handler
	if len(data) > 1 {
		fileUploadHandler = data[1].(context.Handler)
	} else {
		fileUploadHandler = func(ctx *context.Context) {
			if len(ctx.Request.MultipartForm.File) == 0 {
				ctx.JSON(http.StatusOK, map[string]interface{}{
					"errno": 400,
				})
				return
			}

			err := file.GetFileEngine(config.GetFileUploadEngine().Name).Upload(ctx.Request.MultipartForm)
			if err != nil {
				ctx.JSON(http.StatusOK, map[string]interface{}{
					"errno": 500,
				})
				return
			}

			var imgPath = make([]string, len(ctx.Request.MultipartForm.Value["file"]))
			for i, path := range ctx.Request.MultipartForm.Value["file"] {
				imgPath[i] = config.GetStore().URL(path)
			}

			ctx.JSON(http.StatusOK, map[string]interface{}{
				"errno": 0,
				"data":  imgPath,
			})
		}
	}

	f.Callbacks = f.Callbacks.AddCallback(context.Node{
		Path:     url,
		Method:   "post",
		Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
		Handlers: []context.Handler{fileUploadHandler},
	})

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

func (f *FormPanel) FieldOptionsFromTable(table, textFieldName, valueFieldName string, process ...OptionTableQueryProcessFn) *FormPanel {
	var fn OptionTableQueryProcessFn
	if len(process) > 0 {
		fn = process[0]
	}
	f.FieldList[f.curFieldListIndex].OptionTable = OptionTable{
		Table:          table,
		TextField:      textFieldName,
		ValueField:     valueFieldName,
		QueryProcessFn: fn,
	}
	return f
}

func (f *FormPanel) FieldOptionsTableProcessFn(fn OptionProcessFn) *FormPanel {
	f.FieldList[f.curFieldListIndex].OptionTable.ProcessFn = fn
	return f
}

func (f *FormPanel) FieldOptions(options FieldOptions) *FormPanel {
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
		Add(func(value FieldModel) interface{} {
			return html.EscapeString(value.Value)
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
	ext, callback := searchJS(f.FieldList[f.curFieldListIndex].OptionExt, f.OperationURL(url), handler, delay...)
	f.FieldList[f.curFieldListIndex].OptionExt = ext
	f.Callbacks = f.Callbacks.AddCallback(callback)
	return f
}

func (f *FormPanel) FieldOnChooseCustom(js template.HTML) *FormPanel {
	f.FooterHtml += chooseCustomJS(f.FieldList[f.curFieldListIndex].Field, js)
	return f
}

type LinkField struct {
	Field   string
	Value   template.HTML
	Hide    bool
	Disable bool
}

func (f *FormPanel) FieldOnChooseMap(m map[string]LinkField) *FormPanel {
	f.FooterHtml += chooseMapJS(f.FieldList[f.curFieldListIndex].Field, m)
	return f
}

func (f *FormPanel) FieldOnChoose(val, field string, value template.HTML) *FormPanel {
	f.FooterHtml += chooseJS(f.FieldList[f.curFieldListIndex].Field, field, val, value)
	return f
}

func (f *FormPanel) OperationURL(id string) string {
	return config.Url("/operation/" + utils.WrapURL(id))
}

func (f *FormPanel) FieldOnChooseAjax(field, url string, handler Handler, custom ...template.HTML) *FormPanel {
	js, callback := chooseAjax(f.FieldList[f.curFieldListIndex].Field, field, f.OperationURL(url), handler, custom...)
	f.FooterHtml += js
	f.Callbacks = f.Callbacks.AddCallback(callback)
	return f
}

func (f *FormPanel) FieldOnChooseHide(value string, field ...string) *FormPanel {
	f.FooterHtml += chooseHideJS(f.FieldList[f.curFieldListIndex].Field, value, field...)
	return f
}

func (f *FormPanel) FieldOnChooseShow(value string, field ...string) *FormPanel {
	f.FooterHtml += chooseShowJS(f.FieldList[f.curFieldListIndex].Field, value, field...)
	return f
}

func (f *FormPanel) FieldOnChooseDisable(value string, field ...string) *FormPanel {
	f.FooterHtml += chooseDisableJS(f.FieldList[f.curFieldListIndex].Field, value, field...)
	return f
}

func searchJS(ext template.JS, url string, handler Handler, delay ...int) (template.JS, context.Node) {
	delayStr := "500"
	if len(delay) > 0 {
		delayStr = strconv.Itoa(delay[0])
	}

	if ext != template.JS("") {
		s := string(ext)
		s = strings.Replace(s, "{", "", 1)
		s = utils.ReplaceNth(s, "}", "", strings.Count(s, "}"))
		s = strings.TrimRight(s, " ")
		s += ","
		ext = template.JS(s)
	}

	return template.JS(`{
		`) + ext + template.JS(`
		ajax: {
		    url: "`+url+`",
		    dataType: 'json',
		    data: function (params) {
			      var query = {
			        	search: params.term,
						page: params.page || 1
			      }
			      return query;
		    },
		    delay: `+delayStr+`,
		    processResults: function (data, params) {
			      return data.data;
	    	}
	  	}
	}`), context.Node{
			Path:     url,
			Method:   "get",
			Handlers: context.Handlers{handler.Wrap()},
			Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
		}
}

func parseHTML(name string, param interface{}) template.HTML {
	t := template.New(name)
	t, err := t.Parse(tmpls[name])
	if err != nil {
		logger.Error(err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, param)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return template.HTML(buf.String())
}

func chooseCustomJS(field string, js template.HTML) template.HTML {
	return parseHTML("choose_custom", struct {
		Field template.JS
		JS    template.JS
	}{Field: template.JS(field), JS: template.JS(js)})
}

func chooseMapJS(field string, m map[string]LinkField) template.HTML {
	return parseHTML("choose_map", struct {
		Field template.JS
		Data  map[string]LinkField
	}{
		Field: template.JS(field),
		Data:  m,
	})
}

func chooseJS(field, chooseField, val string, value template.HTML) template.HTML {
	return parseHTML("choose", struct {
		Field       template.JS
		ChooseField template.JS
		Val         template.JS
		Value       template.JS
	}{
		Field:       template.JS(field),
		ChooseField: template.JS(chooseField),
		Value:       template.JS(value),
		Val:         template.JS(val),
	})
}

func chooseAjax(field, chooseField, url string, handler Handler, js ...template.HTML) (template.HTML, context.Node) {

	actionJS := template.HTML("")
	passValue := template.JS("")

	if len(js) > 0 {
		actionJS = js[0]
	}

	if len(js) > 1 {
		passValue = template.JS(js[1])
	}

	return parseHTML("choose_ajax", struct {
			Field       template.JS
			ChooseField template.JS
			PassValue   template.JS
			ActionJS    template.JS
			Url         template.JS
		}{
			Url:         template.JS(url),
			Field:       template.JS(field),
			ChooseField: template.JS(chooseField),
			PassValue:   passValue,
			ActionJS:    template.JS(actionJS),
		}), context.Node{
			Path:     url,
			Method:   "post",
			Handlers: context.Handlers{handler.Wrap()},
			Value:    map[string]interface{}{constant.ContextNodeNeedAuth: 1},
		}
}

func chooseHideJS(field, value string, chooseFields ...string) template.HTML {
	if len(chooseFields) == 0 {
		return ""
	}

	return parseHTML("choose_hide", struct {
		Field        template.JS
		Value        template.JS
		ChooseFields []string
	}{
		Field:        template.JS(field),
		Value:        template.JS(value),
		ChooseFields: chooseFields,
	})
}

func chooseShowJS(field, value string, chooseFields ...string) template.HTML {
	if len(chooseFields) == 0 {
		return ""
	}

	return parseHTML("choose_show", struct {
		Field        template.JS
		Value        template.JS
		ChooseFields []string
	}{
		Field:        template.JS(field),
		Value:        template.JS(value),
		ChooseFields: chooseFields,
	})
}

func chooseDisableJS(field, value string, chooseFields ...string) template.HTML {
	if len(chooseFields) == 0 {
		return ""
	}

	return parseHTML("choose_disable", struct {
		Field        template.JS
		Value        template.JS
		ChooseFields []string
	}{
		Field:        template.JS(field),
		Value:        template.JS(value),
		ChooseFields: chooseFields,
	})
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

func (f *FormPanel) SetLayout(layout form2.Layout) *FormPanel {
	f.Layout = layout
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

func (f *FormPanel) SetHTMLContent(content template.HTML) *FormPanel {
	f.HTMLContent = content
	return f
}

func (f *FormPanel) SetHeader(content template.HTML) *FormPanel {
	f.Header = content
	return f
}

func (f *FormPanel) SetInputWidth(width int) *FormPanel {
	f.InputWidth = width
	return f
}

func (f *FormPanel) SetHeadWidth(width int) *FormPanel {
	f.HeadWidth = width
	return f
}

func (f *FormPanel) SetWrapper(wrapper ContentWrapper) *FormPanel {
	f.Wrapper = wrapper
	return f
}

func (f *FormPanel) SetFormNewTitle(title template.HTML) *FormPanel {
	f.FormNewTitle = title
	return f
}

func (f *FormPanel) SetFormNewBtnWord(word template.HTML) *FormPanel {
	f.FormNewBtnWord = word
	return f
}

func (f *FormPanel) SetFormEditTitle(title template.HTML) *FormPanel {
	f.FormEditTitle = title
	return f
}

func (f *FormPanel) SetFormEditBtnWord(word template.HTML) *FormPanel {
	f.FormEditBtnWord = word
	return f
}

func (f *FormPanel) SetResponder(responder Responder) *FormPanel {
	f.Responder = responder
	return f
}

type AjaxData struct {
	SuccessTitle   string
	SuccessText    string
	ErrorTitle     string
	ErrorText      string
	SuccessJumpURL string
}

func (f *FormPanel) EnableAjaxData(data AjaxData) *FormPanel {
	f.Ajax = true
	if f.AjaxSuccessJS == template.JS("") {
		successMsg := modules.AorB(data.SuccessTitle != "", `"`+data.SuccessTitle+`"`, "data.msg")
		errorMsg := modules.AorB(data.ErrorTitle != "", `"`+data.ErrorTitle+`"`, "data.msg")
		jump := modules.AorB(data.SuccessJumpURL != "", `"`+data.SuccessJumpURL+`"`, "data.data.url")
		text := modules.AorB(data.SuccessText != "", `text:"`+data.SuccessText+`",`, "")
		wrongText := modules.AorB(data.ErrorText != "", `text:"`+data.ErrorText+`",`, "")
		f.AjaxSuccessJS = template.JS(`
	if (typeof (data) === "string") {
	    data = JSON.parse(data);
	}
	if (data.code === 200) {
	    swal({
			type: "success",
			title: ` + successMsg + `,
			` + text + `
			showCancelButton: false,
			confirmButtonColor: "#3c8dbc",
			confirmButtonText: '` + language.Get("got it") + `',
        }, function() {
			$(".modal-backdrop.fade.in").remove();
			$.pjax({url: ` + jump + `, container: '#pjax-container'});
        });
	} else {
		swal({
			type: "error",
			title: ` + errorMsg + `,
			` + wrongText + `
			showCancelButton: false,
			confirmButtonColor: "#3c8dbc",
			confirmButtonText: '` + language.Get("got it") + `',
        })
	}
`)
	}
	if f.AjaxErrorJS == template.JS("") {
		errorMsg := modules.AorB(data.ErrorTitle != "", `"`+data.ErrorTitle+`"`, "data.responseJSON.msg")
		error2Msg := modules.AorB(data.ErrorTitle != "", `"`+data.ErrorTitle+`"`, "'"+language.Get("error")+"'")
		wrongText := modules.AorB(data.ErrorText != "", `text:"`+data.ErrorText+`",`, "")
		f.AjaxErrorJS = template.JS(`
	if (data.responseText !== "") {
		swal({
			type: "error",
			title: ` + errorMsg + `,
			` + wrongText + `
			showCancelButton: false,
			confirmButtonColor: "#3c8dbc",
			confirmButtonText: '` + language.Get("got it") + `',
        })								
	} else {
		swal({
			type: "error",
			title: ` + error2Msg + `,
			` + wrongText + `
			showCancelButton: false,
			confirmButtonColor: "#3c8dbc",
			confirmButtonText: '` + language.Get("got it") + `',
        })
	}
`)
	}
	return f
}

func (f *FormPanel) EnableAjax(msgs ...string) *FormPanel {
	var data AjaxData
	if len(msgs) > 0 && msgs[0] != "" {
		data.SuccessTitle = msgs[0]
	}
	if len(msgs) > 1 && msgs[1] != "" {
		data.ErrorTitle = msgs[0]
	}
	if len(msgs) > 2 && msgs[2] != "" {
		data.SuccessJumpURL = msgs[0]
	}
	if len(msgs) > 3 && msgs[3] != "" {
		data.SuccessText = msgs[0]
	}
	if len(msgs) > 4 && msgs[4] != "" {
		data.ErrorText = msgs[0]
	}
	return f.EnableAjaxData(data)
}

func (f *FormPanel) SetAjaxSuccessJS(js template.JS) *FormPanel {
	f.AjaxSuccessJS = js
	return f
}

func (f *FormPanel) SetAjaxErrorJS(js template.JS) *FormPanel {
	f.AjaxErrorJS = js
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

func (f *FormPanel) GroupFieldWithValue(pk, id string, columns []string, res map[string]interface{}, sql func() *db.SQL) ([]FormFields, []string) {
	var (
		groupFormList = make([]FormFields, 0)
		groupHeaders  = make([]string, 0)
		hasPK         = false
	)

	if len(f.TabGroups) > 0 {
		for index, group := range f.TabGroups {
			list := make(FormFields, 0)
			for _, fieldName := range group {
				field := f.FieldList.FindByFieldName(fieldName)
				if field != nil && field.isNotBelongToATable() {
					if field.FormType.IsTable() {
						for z := 0; z < len(field.TableFields); z++ {
							rowValue := field.TableFields[z].GetRawValue(columns, res[field.TableFields[z].Field])
							if field.TableFields[z].Field == pk {
								hasPK = true
							}
							field.TableFields[z] = *(field.TableFields[z].UpdateValue(id, rowValue, res, sql()))
						}
						list = append(list, *field)
					} else {
						if field.Field == pk {
							hasPK = true
						}
						rowValue := field.GetRawValue(columns, res[field.Field])
						list = append(list, *(field.UpdateValue(id, rowValue, res, sql())))
					}
				}
			}

			groupFormList = append(groupFormList, list.FillCustomContent())
			groupHeaders = append(groupHeaders, f.TabHeaders[index])
		}

		if len(groupFormList) > 0 && !hasPK {
			groupFormList[len(groupFormList)-1] = groupFormList[len(groupFormList)-1].Add(FormField{
				Head:       pk,
				FieldClass: pk,
				Field:      pk,
				Value:      template.HTML(id),
				Hide:       true,
			})
		}
	}

	return groupFormList, groupHeaders
}

func (f *FormPanel) GroupField(sql ...func() *db.SQL) ([]FormFields, []string) {
	var (
		groupFormList = make([]FormFields, 0)
		groupHeaders  = make([]string, 0)
	)

	for index, group := range f.TabGroups {
		list := make(FormFields, 0)
		for _, fieldName := range group {
			field := f.FieldList.FindByFieldName(fieldName)
			if field != nil && field.isNotBelongToATable() && field.allowAdd() {
				field.Editable = true
				if field.FormType.IsTable() {
					for z := 0; z < len(field.TableFields); z++ {
						if len(sql) > 0 {
							field.TableFields[z] = *(field.TableFields[z].UpdateDefaultValue(sql[0]()))
						} else {
							field.TableFields[z] = *(field.TableFields[z].UpdateDefaultValue(nil))
						}
					}
					list = append(list, *field)
				} else {
					if len(sql) > 0 {
						list = append(list, *(field.UpdateDefaultValue(sql[0]())))
					} else {
						list = append(list, *(field.UpdateDefaultValue(nil)))
					}
				}
			}
		}
		groupFormList = append(groupFormList, list.FillCustomContent())
		groupHeaders = append(groupHeaders, f.TabHeaders[index])
	}

	return groupFormList, groupHeaders
}

func (f *FormPanel) FieldsWithValue(pk, id string, columns []string, res map[string]interface{}, sql func() *db.SQL) FormFields {
	var (
		list  = make(FormFields, 0)
		hasPK = false
	)
	for _, field := range f.FieldList {
		rowValue := field.GetRawValue(columns, res[field.Field])
		if field.FatherField != "" {
			f.FieldList.FindTableField(field.Field, field.FatherField).UpdateValue(id, rowValue, res, sql())
		} else if field.FormType.IsTable() {
			list = append(list, field)
		} else {
			list = append(list, *(field.UpdateValue(id, rowValue, res, sql())))
		}

		if field.Field == pk {
			hasPK = true
		}
	}
	if !hasPK {
		list = list.Add(FormField{
			Head:       pk,
			FieldClass: pk,
			Field:      pk,
			Value:      template.HTML(id),
			FormType:   form2.Default,
			Hide:       true,
		})
	}
	return list.FillCustomContent()
}

func (f *FormPanel) FieldsWithDefaultValue(sql ...func() *db.SQL) FormFields {
	var list = make(FormFields, 0)
	for _, v := range f.FieldList {
		if v.allowAdd() {
			v.Editable = true
			if v.FatherField != "" {
				if len(sql) > 0 {
					f.FieldList.FindTableField(v.Field, v.FatherField).UpdateDefaultValue(sql[0]())
				} else {
					f.FieldList.FindTableField(v.Field, v.FatherField).UpdateDefaultValue(nil)
				}
			} else if v.FormType.IsTable() {
				list = append(list, v)
			} else {
				if len(sql) > 0 {
					list = append(list, *(v.UpdateDefaultValue(sql[0]())))
				} else {
					list = append(list, *(v.UpdateDefaultValue(nil)))
				}
			}
		}
	}
	return list.FillCustomContent().RemoveNotShow()
}

func (f *FormPanel) GetNewFormFields(sql ...func() *db.SQL) (FormFields, []FormFields, []string) {
	if len(f.TabGroups) > 0 {
		tabFields, tabHeaders := f.GroupField(sql...)
		return make(FormFields, 0), tabFields, tabHeaders
	}
	return f.FieldsWithDefaultValue(sql...), make([]FormFields, 0), make([]string, 0)
}

type (
	FormPreProcessFn  func(values form.Values) form.Values
	FormPostFn        func(values form.Values) error
	FormFields        []FormField
	GroupFormFields   []FormFields
	GroupFieldHeaders []string
)

func (f FormFields) Copy() FormFields {
	formList := make(FormFields, len(f))
	copy(formList, f)
	for i := 0; i < len(formList); i++ {
		formList[i].Options = make(FieldOptions, len(f[i].Options))
		for j := 0; j < len(f[i].Options); j++ {
			formList[i].Options[j] = FieldOption{
				Value:    f[i].Options[j].Value,
				Text:     f[i].Options[j].Text,
				TextHTML: f[i].Options[j].TextHTML,
				Selected: f[i].Options[j].Selected,
			}
		}
	}
	return formList
}

func (f FormFields) FindByFieldName(field string) *FormField {
	for i := 0; i < len(f); i++ {
		if f[i].Field == field {
			return &f[i]
		}
	}
	return nil
}

func (f FormFields) FindIndexByFieldName(field string) int {
	for i := 0; i < len(f); i++ {
		if f[i].Field == field {
			return i
		}
	}
	return -1
}

func (f FormFields) FindTableField(field, father string) *FormField {
	ff := f.FindByFieldName(father)
	return ff.TableFields.FindByFieldName(field)
}

func (f FormFields) FindTableChildren(father string) []*FormField {
	list := make([]*FormField, 0)
	for i := 0; i < len(f); i++ {
		if f[i].FatherField == father {
			list = append(list, &f[i])
		}
	}
	return list
}

func (f FormFields) FillCustomContent() FormFields {
	for i := range f {
		if f[i].FormType.IsCustom() {
			f[i] = *(f[i]).FillCustomContent()
		}
	}
	return f
}

func (f FormFields) Add(field FormField) FormFields {
	return append(f, field)
}

func (f FormFields) RemoveNotShow() FormFields {
	ff := f
	for i := 0; i < len(ff); {
		if ff[i].FatherFormType == form2.Table {
			ff = append(ff[:i], ff[i+1:]...)
		} else {
			i++
		}
	}
	return ff
}
