package form

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
)

type Type uint8

const (
	Default Type = iota
	Text
	SelectSingle
	Select
	IconPicker
	SelectBox
	File
	Multifile
	Password
	RichText
	Datetime
	DatetimeRange
	Radio
	Checkbox
	CheckboxStacked
	CheckboxSingle
	Email
	Date
	DateRange
	Url
	Ip
	Color
	Array
	Currency
	Rate
	Number
	Table
	NumberRange
	TextArea
	Custom
	Switch
	Code
	Slider
)

var AllType = []Type{Default, Text, Array, SelectSingle, Select, IconPicker, SelectBox, File, Multifile, Password,
	RichText, Datetime, DatetimeRange, Checkbox, CheckboxStacked, Radio, Table, Email, Url, Ip, Color, Currency, Number, NumberRange,
	TextArea, Custom, Switch, Code, Rate, Slider, Date, DateRange, CheckboxSingle}

func CheckType(t, def Type) Type {
	for _, item := range AllType {
		if t == item {
			return t
		}
	}
	return def
}

type Layout uint8

const (
	LayoutDefault Layout = iota
	LayoutTwoCol
	LayoutThreeCol
	LayoutFourCol
	LayoutFiveCol
	LayoutSixCol
	LayoutFlow
	LayoutTab
)

func (l Layout) Col() int {
	if l == LayoutTwoCol {
		return 2
	}
	if l == LayoutThreeCol {
		return 3
	}
	if l == LayoutFourCol {
		return 4
	}
	if l == LayoutFiveCol {
		return 5
	}
	if l == LayoutSixCol {
		return 6
	}
	return 0
}

func (l Layout) Flow() bool {
	return l == LayoutFlow
}

func (l Layout) Default() bool {
	return l == LayoutDefault
}

func (l Layout) String() string {
	switch l {
	case LayoutDefault:
		return "LayoutDefault"
	case LayoutTwoCol:
		return "LayoutTwoCol"
	case LayoutThreeCol:
		return "LayoutThreeCol"
	case LayoutFourCol:
		return "LayoutFourCol"
	case LayoutFiveCol:
		return "LayoutFiveCol"
	case LayoutSixCol:
		return "LayoutSixCol"
	case LayoutFlow:
		return "LayoutFlow"
	case LayoutTab:
		return "LayoutTab"
	default:
		return "LayoutDefault"
	}
}

func GetLayoutFromString(s string) Layout {
	switch s {
	case "LayoutDefault":
		return LayoutDefault
	case "LayoutTwoCol":
		return LayoutTwoCol
	case "LayoutThreeCol":
		return LayoutThreeCol
	case "LayoutFourCol":
		return LayoutFourCol
	case "LayoutFiveCol":
		return LayoutFiveCol
	case "LayoutSixCol":
		return LayoutSixCol
	case "LayoutFlow":
		return LayoutFlow
	case "LayoutTab":
		return LayoutTab
	default:
		return LayoutDefault
	}
}

func (t Type) Name() string {
	switch t {
	case Default:
		return "Default"
	case Text:
		return "Text"
	case SelectSingle:
		return "SelectSingle"
	case Select:
		return "Select"
	case IconPicker:
		return "IconPicker"
	case SelectBox:
		return "SelectBox"
	case File:
		return "File"
	case Table:
		return "Table"
	case Multifile:
		return "Multifile"
	case Password:
		return "Password"
	case RichText:
		return "RichText"
	case Rate:
		return "Rate"
	case Checkbox:
		return "Checkbox"
	case CheckboxStacked:
		return "CheckboxStacked"
	case CheckboxSingle:
		return "CheckboxSingle"
	case Date:
		return "Date"
	case DateRange:
		return "DateRange"
	case Datetime:
		return "Datetime"
	case DatetimeRange:
		return "DatetimeRange"
	case Radio:
		return "Radio"
	case Slider:
		return "Slider"
	case Array:
		return "Array"
	case Email:
		return "Email"
	case Url:
		return "Url"
	case Ip:
		return "Ip"
	case Color:
		return "Color"
	case Currency:
		return "Currency"
	case Number:
		return "Number"
	case NumberRange:
		return "NumberRange"
	case TextArea:
		return "TextArea"
	case Custom:
		return "Custom"
	case Switch:
		return "Switch"
	case Code:
		return "Code"
	default:
		panic("wrong form type")
	}
}

func (t Type) String() string {
	switch t {
	case Default:
		return "default"
	case Text:
		return "text"
	case SelectSingle:
		return "select_single"
	case Select:
		return "select"
	case IconPicker:
		return "iconpicker"
	case SelectBox:
		return "selectbox"
	case File:
		return "file"
	case Table:
		return "table"
	case Multifile:
		return "multi_file"
	case Password:
		return "password"
	case RichText:
		return "richtext"
	case Rate:
		return "rate"
	case Checkbox:
		return "checkbox"
	case CheckboxStacked:
		return "checkbox_stacked"
	case CheckboxSingle:
		return "checkbox_single"
	case Date:
		return "datetime"
	case DateRange:
		return "datetime_range"
	case Datetime:
		return "datetime"
	case DatetimeRange:
		return "datetime_range"
	case Radio:
		return "radio"
	case Slider:
		return "slider"
	case Array:
		return "array"
	case Email:
		return "email"
	case Url:
		return "url"
	case Ip:
		return "ip"
	case Color:
		return "color"
	case Currency:
		return "currency"
	case Number:
		return "number"
	case NumberRange:
		return "number_range"
	case TextArea:
		return "textarea"
	case Custom:
		return "custom"
	case Switch:
		return "switch"
	case Code:
		return "code"
	default:
		panic("wrong form type")
	}
}

func (t Type) IsSelect() bool {
	return t == Select || t == SelectSingle || t == SelectBox || t == Radio || t == Switch ||
		t == Checkbox || t == CheckboxStacked || t == CheckboxSingle
}

func (t Type) IsArray() bool {
	return t == Array
}

func (t Type) IsTable() bool {
	return t == Table
}

func (t Type) IsSingleSelect() bool {
	return t == SelectSingle || t == Radio || t == Switch || t == CheckboxSingle
}

func (t Type) IsMultiSelect() bool {
	return t == Select || t == SelectBox || t == Checkbox || t == CheckboxStacked
}

func (t Type) IsMultiFile() bool {
	return t == Multifile
}

func (t Type) IsRange() bool {
	return t == DatetimeRange || t == NumberRange
}

func (t Type) IsFile() bool {
	return t == File || t == Multifile
}

func (t Type) IsSlider() bool {
	return t == Slider
}

func (t Type) IsDateTime() bool {
	return t == Datetime
}

func (t Type) IsDateTimeRange() bool {
	return t == DatetimeRange
}

func (t Type) IsDate() bool {
	return t == Date
}

func (t Type) IsDateRange() bool {
	return t == DateRange
}

func (t Type) IsCode() bool {
	return t == Code
}

func (t Type) IsRichText() bool {
	return t == RichText
}

func (t Type) IsTextarea() bool {
	return t == TextArea
}

func (t Type) IsEditor() bool {
	return t == TextArea || t == Code || t == RichText
}

func (t Type) IsCustom() bool {
	return t == Custom
}

func (t Type) FixOptions(m map[string]interface{}) map[string]interface{} {
	if t == Slider {
		if _, ok := m["type"]; !ok {
			m["type"] = "single"
		}
		if _, ok := m["prettify"]; !ok {
			m["prettify"] = false
		}
		if _, ok := m["hasGrid"]; !ok {
			m["hasGrid"] = true
		}
		return m
	}
	return m
}

func (t Type) SelectedLabel() []template.HTML {
	if t == Select || t == SelectSingle || t == SelectBox {
		return []template.HTML{"selected", ""}
	}
	if t == Radio || t == Switch || t == Checkbox || t == CheckboxStacked || t == CheckboxSingle {
		return []template.HTML{"checked", ""}
	}
	return []template.HTML{"", ""}
}

func (t Type) GetDefaultOptions(field string) (map[string]interface{}, map[string]interface{}, template.JS) {
	switch t {
	case File, Multifile:
		return map[string]interface{}{
			"overwriteInitial":     true,
			"initialPreviewAsData": true,
			"browseLabel":          language.Get("Browse"),
			"showRemove":           false,
			"previewClass":         "preview-" + field,
			"showUpload":           false,
			"allowedFileTypes":     []string{"image"},
		}, nil, ""
	case Slider:
		return map[string]interface{}{
			"type":     "single",
			"prettify": false,
			"hasGrid":  true,
			"max":      100,
			"min":      1,
			"step":     1,
			"postfix":  "",
		}, nil, ""
	case DatetimeRange:
		op1, op2 := getDateTimeRangeOptions(DatetimeRange)
		return op1, op2, ""
	case Datetime:
		return getDateTimeOptions(Datetime), nil, ""
	case Date:
		return getDateTimeOptions(Date), nil, ""
	case DateRange:
		op1, op2 := getDateTimeRangeOptions(DateRange)
		return op1, op2, ""
	case Code:
		return nil, nil, `
	theme = "monokai";
	font_size = 14;
	language = "html";
	options = {useWorker: false};
`
	}

	return nil, nil, ""
}

func getDateTimeOptions(f Type) map[string]interface{} {
	format := "YYYY-MM-DD HH:mm:ss"
	if f == Date {
		format = "YYYY-MM-DD"
	}
	m := map[string]interface{}{
		"format":           format,
		"locale":           "en",
		"allowInputToggle": true,
	}
	if config.GetLanguage() == language.CN || config.GetLanguage() == "cn" {
		m["locale"] = "zh-CN"
	}
	return m
}

func getDateTimeRangeOptions(f Type) (map[string]interface{}, map[string]interface{}) {
	format := "YYYY-MM-DD HH:mm:ss"
	if f == DateRange {
		format = "YYYY-MM-DD"
	}
	m := map[string]interface{}{
		"format": format,
		"locale": "en",
	}
	m1 := map[string]interface{}{
		"format":     format,
		"locale":     "en",
		"useCurrent": false,
	}
	if config.GetLanguage() == language.CN || config.GetLanguage() == "cn" {
		m["locale"] = "zh-CN"
		m1["locale"] = "zh-CN"
	}
	return m, m1
}

func GetFormTypeFromFieldType(typeName db.DatabaseType, fieldName string) string {

	if fieldName == "password" {
		return "Password"
	}

	if fieldName == "id" {
		return "Default"
	}

	if fieldName == "ip" {
		return "Ip"
	}

	if fieldName == "Url" {
		return "Url"
	}

	if fieldName == "email" {
		return "Email"
	}

	if fieldName == "color" {
		return "Color"
	}

	if fieldName == "money" {
		return "Currency"
	}

	switch typeName {
	case db.Int, db.Tinyint, db.Int4, db.Integer, db.Mediumint, db.Smallint,
		db.Numeric, db.Smallserial, db.Serial, db.Bigserial, db.Money, db.Bigint:
		return "Number"
	case db.Text, db.Longtext, db.Mediumtext, db.Tinytext:
		return "RichText"
	case db.Datetime, db.Date, db.Time, db.Timestamp, db.Timestamptz, db.Year:
		return "Datetime"
	}

	return "Text"
}

func DefaultHTML(value string) template.HTML {
	return template.HTML(`<div class="box box-solid box-default no-margin"><div class="box-body" style="min-height: 40px;">` +
		value + `</div></div>`)
}

func HiddenInputHTML(field, value string) template.HTML {
	return template.HTML(`<input type="hidden" name="` + field + `" value="` + value + `" class="form-control">`)
}
