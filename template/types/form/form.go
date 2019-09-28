package form

import (
	"github.com/chenhg5/go-admin/modules/db"
)

const (
	Default      = "default"
	Text         = "text"
	SelectSingle = "select_single"
	Select       = "select"
	IconPicker   = "iconpicker"
	SelectBox    = "selectbox"
	File         = "file"
	Password     = "password"
	RichText     = "richtext"
	Datetime     = "datetime"
	Radio        = "radio"
	Email        = "email"
	Url          = "url"
	Ip           = "ip"
	Color        = "color"
	Currency     = "currency"
	Number       = "number"
)

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
