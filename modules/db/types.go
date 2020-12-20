// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"html/template"
	"strconv"
)

// DatabaseType is the database field type.
type DatabaseType string

const (
	// =================================
	// integer
	// =================================

	Int       DatabaseType = "INT"
	Tinyint   DatabaseType = "TINYINT"
	Mediumint DatabaseType = "MEDIUMINT"
	Smallint  DatabaseType = "SMALLINT"
	Bigint    DatabaseType = "BIGINT"
	Bit       DatabaseType = "BIT"
	Int8      DatabaseType = "INT8"
	Int4      DatabaseType = "INT4"
	Int2      DatabaseType = "INT2"

	Integer     DatabaseType = "INTEGER"
	Numeric     DatabaseType = "NUMERIC"
	Smallserial DatabaseType = "SMALLSERIAL"
	Serial      DatabaseType = "SERIAL"
	Bigserial   DatabaseType = "BIGSERIAL"
	Money       DatabaseType = "MONEY"

	// =================================
	// float
	// =================================

	Real    DatabaseType = "REAL"
	Float   DatabaseType = "FLOAT"
	Float4  DatabaseType = "FLOAT4"
	Float8  DatabaseType = "FLOAT8"
	Double  DatabaseType = "DOUBLE"
	Decimal DatabaseType = "DECIMAL"

	Doubleprecision DatabaseType = "DOUBLEPRECISION"

	// =================================
	// string
	// =================================

	Date      DatabaseType = "DATE"
	Time      DatabaseType = "TIME"
	Year      DatabaseType = "YEAR"
	Datetime  DatabaseType = "DATETIME"
	Timestamp DatabaseType = "TIMESTAMP"

	Text       DatabaseType = "TEXT"
	Longtext   DatabaseType = "LONGTEXT"
	Mediumtext DatabaseType = "MEDIUMTEXT"
	Tinytext   DatabaseType = "TINYTEXT"

	Varchar DatabaseType = "VARCHAR"
	Char    DatabaseType = "CHAR"
	Bpchar  DatabaseType = "BPCHAR"
	JSON    DatabaseType = "JSON"

	Blob       DatabaseType = "BLOB"
	Tinyblob   DatabaseType = "TINYBLOB"
	Mediumblob DatabaseType = "MEDIUMBLOB"
	Longblob   DatabaseType = "LONGBLOB"

	Interval DatabaseType = "INTERVAL"
	Boolean  DatabaseType = "BOOLEAN"
	Bool     DatabaseType = "BOOL"

	Point   DatabaseType = "POINT"
	Line    DatabaseType = "LINE"
	Lseg    DatabaseType = "LSEG"
	Box     DatabaseType = "BOX"
	Path    DatabaseType = "PATH"
	Polygon DatabaseType = "POLYGON"
	Circle  DatabaseType = "CIRCLE"

	Cidr    DatabaseType = "CIDR"
	Inet    DatabaseType = "INET"
	Macaddr DatabaseType = "MACADDR"

	Character        DatabaseType = "CHARACTER"
	Varyingcharacter DatabaseType = "VARYINGCHARACTER"
	Nchar            DatabaseType = "NCHAR"
	Nativecharacter  DatabaseType = "NATIVECHARACTER"
	Nvarchar         DatabaseType = "NVARCHAR"
	Clob             DatabaseType = "CLOB"

	Binary    DatabaseType = "BINARY"
	Varbinary DatabaseType = "VARBINARY"
	Enum      DatabaseType = "ENUM"
	Set       DatabaseType = "SET"

	Geometry DatabaseType = "GEOMETRY"

	Multilinestring    DatabaseType = "MULTILINESTRING"
	Multipolygon       DatabaseType = "MULTIPOLYGON"
	Linestring         DatabaseType = "LINESTRING"
	Multipoint         DatabaseType = "MULTIPOINT"
	Geometrycollection DatabaseType = "GEOMETRYCOLLECTION"

	Name DatabaseType = "NAME"
	UUID DatabaseType = "UUID"

	Timestamptz DatabaseType = "TIMESTAMPTZ"
	Timetz      DatabaseType = "TIMETZ"
)

// DT turn the string value into DatabaseType.
func DT(s string) DatabaseType {
	return DatabaseType(s)
}

// GetDTAndCheck check the DatabaseType.
func GetDTAndCheck(s string) DatabaseType {
	ss := DatabaseType(s)
	if !Contains(ss, BoolTypeList) &&
		!Contains(ss, IntTypeList) &&
		!Contains(ss, FloatTypeList) &&
		!Contains(ss, UintTypeList) &&
		!Contains(ss, StringTypeList) {
		panic("wrong type: " + s)
	}
	return ss
}

var (
	// StringTypeList is a DatabaseType list of string.
	StringTypeList = []DatabaseType{Date, Time, Year, Datetime, Timestamptz, Timestamp, Timetz,
		Varchar, Char, Mediumtext, Longtext, Tinytext,
		Text, JSON, Blob, Tinyblob, Mediumblob, Longblob,
		Interval, Point, Bpchar,
		Line, Lseg, Box, Path, Polygon, Circle, Cidr, Inet, Macaddr, Character, Varyingcharacter,
		Nchar, Nativecharacter, Nvarchar, Clob, Binary, Varbinary, Enum, Set, Geometry, Multilinestring,
		Multipolygon, Linestring, Multipoint, Geometrycollection, Name, UUID, Timestamptz,
		Name, UUID, Inet}

	// BoolTypeList is a DatabaseType list of bool.
	BoolTypeList = []DatabaseType{Bool, Boolean}

	// IntTypeList is a DatabaseType list of integer.
	IntTypeList = []DatabaseType{Int4, Int2, Int8,
		Int,
		Tinyint,
		Mediumint,
		Smallint,
		Smallserial, Serial, Bigserial,
		Integer,
		Bigint}

	// FloatTypeList is a DatabaseType list of float.
	FloatTypeList = []DatabaseType{Float, Float4, Float8, Double, Real, Doubleprecision}

	// UintTypeList is a DatabaseType list of uint.
	UintTypeList = []DatabaseType{Decimal, Bit, Money, Numeric}
)

// Contains check the given DatabaseType is in the list or not.
func Contains(v DatabaseType, a []DatabaseType) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

// Value is a string.
type Value string

// ToInt64 turn the string to a int64.
func (v Value) ToInt64() int64 {
	value, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		panic("wrong value")
	}
	return value
}

// String return the string value.
func (v Value) String() string {
	return string(v)
}

// HTML return the template.HTML value.
func (v Value) HTML() template.HTML {
	return template.HTML(v)
}

func GetValueFromDatabaseType(typ DatabaseType, value interface{}, json bool) Value {
	if json {
		return GetValueFromJSONOfDatabaseType(typ, value)
	} else {
		return GetValueFromSQLOfDatabaseType(typ, value)
	}
}

// GetValueFromDatabaseType return Value of given DatabaseType and interface.
func GetValueFromSQLOfDatabaseType(typ DatabaseType, value interface{}) Value {
	switch {
	case Contains(typ, StringTypeList):
		if v, ok := value.(string); ok {
			return Value(v)
		}
		return ""
	case Contains(typ, BoolTypeList):
		if v, ok := value.(bool); ok {
			if v {
				return "true"
			}
			return "false"
		}
		if v, ok := value.(int64); ok {
			if v == 0 {
				return "false"
			}
			return "true"
		}
		return "false"
	case Contains(typ, IntTypeList):
		if v, ok := value.(int64); ok {
			return Value(fmt.Sprintf("%d", v))
		}
		return "0"
	case Contains(typ, FloatTypeList):
		if v, ok := value.(float64); ok {
			return Value(fmt.Sprintf("%f", v))
		}
		return "0"
	case Contains(typ, UintTypeList):
		if v, ok := value.([]uint8); ok {
			return Value(string(v))
		}
		return "0"
	}
	panic("wrong type：" + string(typ))
}

// GetValueFromJSONOfDatabaseType return Value of given DatabaseType and interface from JSON string value.
func GetValueFromJSONOfDatabaseType(typ DatabaseType, value interface{}) Value {
	switch {
	case Contains(typ, StringTypeList):
		if v, ok := value.(string); ok {
			return Value(v)
		}
		return ""
	case Contains(typ, BoolTypeList):
		if v, ok := value.(bool); ok {
			if v {
				return "true"
			}
			return "false"
		}
		return "false"
	case Contains(typ, IntTypeList):
		if v, ok := value.(float64); ok {
			return Value(fmt.Sprintf("%d", int64(v)))
		}
		return Value(fmt.Sprintf("%d", value))
	case Contains(typ, FloatTypeList):
		return Value(fmt.Sprintf("%f", value))
	case Contains(typ, UintTypeList):
		if v, ok := value.([]uint8); ok {
			return Value(string(v))
		}
		return "0"
	}
	panic("wrong type：" + string(typ))
}
