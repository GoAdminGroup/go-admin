package form

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"html/template"
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

var allType = []Type{Default, Text, Array, SelectSingle, Select, IconPicker, SelectBox, File, Multifile, Password,
	RichText, Datetime, DatetimeRange, Checkbox, CheckboxStacked, Radio, Table, Email, Url, Ip, Color, Currency, Number, NumberRange,
	TextArea, Custom, Switch, Code, Rate, Slider, Date, DateRange}

func CheckType(t, def Type) Type {
	for _, item := range allType {
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
		t == Checkbox || t == CheckboxStacked
}

func (t Type) IsArray() bool {
	return t == Array
}

func (t Type) IsTable() bool {
	return t == Table
}

func (t Type) IsSingleSelect() bool {
	return t == SelectSingle || t == Radio || t == Switch
}

func (t Type) IsMultiSelect() bool {
	return t == Select || t == SelectBox || t == Checkbox || t == CheckboxStacked
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

func (t Type) IsCustom() bool {
	return t == Custom
}

func (t Type) FixOptions(m map[string]interface{}) map[string]interface{} {
	switch t {
	case Slider:
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
	if t == Radio || t == Switch || t == Checkbox || t == CheckboxStacked {
		return []template.HTML{"checked", ""}
	}
	return []template.HTML{"", ""}
}

func (t Type) GetDefaultOptions(field string) (map[string]interface{}, map[string]interface{}) {
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
		}, nil
	case Slider:
		return map[string]interface{}{
			"type":     "single",
			"prettify": false,
			"hasGrid":  true,
			"max":      100,
			"min":      1,
			"step":     1,
			"postfix":  "",
		}, nil
	case DatetimeRange:
		return getDateTimeRangeOptions(DatetimeRange)
	case Datetime:
		return getDateTimeOptions(Datetime), nil
	case Date:
		return getDateTimeOptions(Date), nil
	case DateRange:
		return getDateTimeRangeOptions(DateRange)
	}

	return map[string]interface{}{}, nil
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
