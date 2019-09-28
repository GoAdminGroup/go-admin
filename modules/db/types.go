package db

import (
	"fmt"
	"strconv"
	"strings"
)

type DatabaseType uint8

const (
	Varchar DatabaseType = iota
	Text
	LongText
	Json
	Int
	LongInt
	Float
	Double
	Decimal
	Date
	Time
	Year
	Datetime
	Timestamp
	MediumText
	TinyText
	Tinyint
	Mediumint
	Smallint
	Bigint
)

type Value string

func (v Value) ToInt64() int64 {
	value, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		panic("wrong value")
	}
	return value
}

func GetValueFromDatabaseType(value interface{}, typ DatabaseType) Value {
	switch typ {
	case Varchar, LongText, Json, Text:
		return Value(value.(string))
	case Int, LongInt:
		return Value(fmt.Sprintf("%d", value.(int64)))
	case Float, Double:
		return Value(fmt.Sprintf("%d", value.(float64)))
	}
	panic("wrong type")
}

func GetTypeFromString(typeName string) DatabaseType {
	typeName = strings.ToUpper(typeName)
	switch typeName {
	case "INT":
		return Int
	case "TINYINT":
		return Tinyint
	case "MEDIUMINT":
		return Mediumint
	case "SMALLINT":
		return Smallint
	case "BIGINT":
		return Bigint
	case "FLOAT":
		return Float
	case "DOUBLE":
		return Double
	case "DECIMAL":
		return Decimal
	case "DATE":
		return Date
	case "TIME":
		return Time
	case "YEAR":
		return Year
	case "DATETIME":
		return Datetime
	case "TIMESTAMP":
		return Timestamp
	case "VARCHAR":
		return Varchar
	case "MEDIUMTEXT":
		return MediumText
	case "LONGTEXT":
		return LongText
	case "TINYTEXT":
		return TinyText
	case "TEXT":
		return Text
	default:
		panic("wrong type name")
	}
}
