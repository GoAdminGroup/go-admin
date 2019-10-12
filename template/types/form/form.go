package form

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
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
	Password
	RichText
	Datetime
	Radio
	Email
	Url
	Ip
	Color
	Currency
	Number
	TextArea
)

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
	case Password:
		return "password"
	case RichText:
		return "richtext"
	case Datetime:
		return "datetime"
	case Radio:
		return "radio"
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
	case TextArea:
		return "textarea"
	default:
		panic("wrong form type")
	}
}

func (t Type) IsSelect() bool {
	return t == Select || t == SelectSingle || t == SelectBox
}

func (t Type) IsRadio() bool {
	return t == Radio
}

func GetFormTypeFromFieldType(typeName db.DatabaseType, fieldName string) string {

	if fieldName == "password" {
		return "form.Password"
	}

	if fieldName == "id" {
		return "form.Default"
	}

	if fieldName == "ip" {
		return "form.Ip"
	}

	if fieldName == "Url" {
		return "form.Url"
	}

	if fieldName == "email" {
		return "form.Email"
	}

	if fieldName == "color" {
		return "form.Color"
	}

	if fieldName == "money" {
		return "form.Currency"
	}

	switch typeName {
	case db.Int, db.Tinyint, db.Int4, db.Integer, db.Mediumint, db.Smallint,
		db.Numeric, db.Smallserial, db.Serial, db.Bigserial, db.Money, db.Bigint:
		return "form.Number"
	case db.Text, db.Longtext, db.Mediumtext, db.Tinytext:
		return "form.RichText"
	case db.Datetime, db.Date, db.Time, db.Timestamp, db.Timestamptz, db.Year:
		return "form.Datetime"
	}

	return "form.Text"
}
