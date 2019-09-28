package db

import (
	"fmt"
	"strconv"
)

type DatabaseType string

const (
	// =================================
	// integer
	// =================================

	Int       DatabaseType = "INT"
	TinyInt   DatabaseType = "TINYINT"
	MediumInt DatabaseType = "MEDIUMINT"
	SmallInt  DatabaseType = "SMALLINT"
	Bigint    DatabaseType = "BIGINT"
	Bit       DatabaseType = "BIT"
	Int4      DatabaseType = "INT4"

	Integer     DatabaseType = "INTEGER"
	Numeric     DatabaseType = "NUMERIC"
	SmallSerial DatabaseType = "SMALLSERIAL"
	Serial      DatabaseType = "SERIAL"
	BigSerial   DatabaseType = "BIGSERIAL"
	Money       DatabaseType = "MONEY"

	// =================================
	// float
	// =================================

	Real    DatabaseType = "REAL"
	Float   DatabaseType = "FLOAT"
	Double  DatabaseType = "DOUBLE"
	Decimal DatabaseType = "DECIMAL"

	DoublePrecision DatabaseType = "DOUBLEPRECISION"

	// =================================
	// string
	// =================================

	Date      DatabaseType = "DATE"
	Time      DatabaseType = "TIME"
	Year      DatabaseType = "YEAR"
	Datetime  DatabaseType = "DATETIME"
	Timestamp DatabaseType = "TIMESTAMP"

	Text       DatabaseType = "TEXT"
	LongText   DatabaseType = "LONGTEXT"
	MediumText DatabaseType = "MEDIUMTEXT"
	TinyText   DatabaseType = "TINYTEXT"

	Varchar DatabaseType = "VARCHAR"
	Char    DatabaseType = "CHAR"
	Json    DatabaseType = "JSON"

	Blob       DatabaseType = "BLOB"
	TinyBlob   DatabaseType = "TINYBLOB"
	MediumBlob DatabaseType = "MEDIUMBLOB"
	LongBlob   DatabaseType = "LONGBLOB"

	Interval DatabaseType = "INTERVAL"
	Boolean  DatabaseType = "BOOLEAN"
	Bool     DatabaseType = "Bool"

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
	VaryingCharacter DatabaseType = "VARYINGCHARACTER"
	Nchar            DatabaseType = "NCHAR"
	NativeCharacter  DatabaseType = "NATIVECHARACTER"
	Nvarchar         DatabaseType = "NVARCHAR"
	Clob             DatabaseType = "CLOB"

	Binary    DatabaseType = "BINARY"
	VarBinary DatabaseType = "VARBINARY"
	Enum      DatabaseType = "ENUM"
	Set       DatabaseType = "SET"

	Geometry DatabaseType = "GEOMETRY"

	Multilinestring    DatabaseType = "MULTILINESTRING"
	Multipolygon       DatabaseType = "MULTIPOLYGON"
	Linestring         DatabaseType = "LINESTRING"
	Multipoint         DatabaseType = "MULTIPOINT"
	Geometrycollection DatabaseType = "GEOMETRYCOLLECTION"

	Name DatabaseType = "NAME"
	Uuid DatabaseType = "UUID"

	Timestamptz DatabaseType = "TIMESTAMPTZ"
)

func DT(s string) DatabaseType {
	return DatabaseType(s)
}

func GetDTAndCheck(s string) DatabaseType {
	ss := DatabaseType(s)
	if !Contains(ss, BoolTypeList) &&
		!Contains(ss, IntTypeList) &&
		!Contains(ss, FloatTypeList) &&
		!Contains(ss, UintTypeList) &&
		!Contains(ss, StringTypeList) {
		panic("wrong type")
	}
	return ss
}

var (
	StringTypeList = []DatabaseType{Date, Time, Year, Datetime, Timestamptz, Timestamp,
		Varchar, Char, MediumText, LongText, TinyText,
		Text, Json, Blob, TinyBlob, MediumBlob, LongBlob,
		Interval, Point,
		Line, Lseg, Box, Path, Polygon, Circle, Cidr, Inet, Macaddr, Character, VaryingCharacter,
		Nchar, NativeCharacter, Nvarchar, Clob, Binary, VarBinary, Enum, Set, Geometry, Multilinestring,
		Multipolygon, Linestring, Multipoint, Geometrycollection, Name, Uuid, Timestamptz,
		Name, Uuid, Inet}
	BoolTypeList = []DatabaseType{Bool, Boolean}
	IntTypeList  = []DatabaseType{Int4,
		Int,
		TinyInt,
		MediumInt,
		SmallInt,
		Numeric, SmallSerial, Serial, BigSerial, Money,
		Integer,
		Bigint}
	FloatTypeList = []DatabaseType{Float, Double, Real, DoublePrecision}
	UintTypeList  = []DatabaseType{Decimal, Bit}
)

func Contains(v DatabaseType, a []DatabaseType) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

type Value string

func (v Value) ToInt64() int64 {
	value, err := strconv.ParseInt(string(v), 10, 64)
	if err != nil {
		panic("wrong value")
	}
	return value
}

func (v Value) String() string {
	return string(v)
}

func GetValueFromDatabaseType(typ DatabaseType, value interface{}) Value {
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
			} else {
				return "false"
			}
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
	}
	panic("wrong type")
}
