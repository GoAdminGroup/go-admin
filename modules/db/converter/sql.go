// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package converter

import (
	"database/sql"
)

func SetColVarType(colVar *[]interface{}, i int, typeName string) {
	switch typeName {
	case "INT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "TINYINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "MEDIUMINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "SMALLINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "BIGINT":
		var s sql.NullInt64
		(*colVar)[i] = &s
	case "FLOAT":
		var s sql.NullFloat64
		(*colVar)[i] = &s
	case "DOUBLE":
		var s sql.NullFloat64
		(*colVar)[i] = &s
	case "DECIMAL":
		var s []uint8
		(*colVar)[i] = &s
	case "DATE":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TIME":
		var s sql.NullString
		(*colVar)[i] = &s
	case "YEAR":
		var s sql.NullString
		(*colVar)[i] = &s
	case "DATETIME":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TIMESTAMP":
		var s sql.NullString
		(*colVar)[i] = &s
	case "VARCHAR":
		var s sql.NullString
		(*colVar)[i] = &s
	case "CHAR":
		var s sql.NullString
		(*colVar)[i] = &s
	case "MEDIUMTEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	case "LONGTEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TINYTEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	case "TEXT":
		var s sql.NullString
		(*colVar)[i] = &s
	default:
		var s interface{}
		(*colVar)[i] = &s
	}
}

func SetResultValue(result *map[string]interface{}, index string, colVar interface{}, typeName string) {
	switch typeName {
	case "INT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "TINYINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "MEDIUMINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "SMALLINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "BIGINT":
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case "FLOAT":
		temp := *(colVar.(*sql.NullFloat64))
		if temp.Valid {
			(*result)[index] = temp.Float64
		} else {
			(*result)[index] = nil
		}
	case "DOUBLE":
		temp := *(colVar.(*sql.NullFloat64))
		if temp.Valid {
			(*result)[index] = temp.Float64
		} else {
			(*result)[index] = nil
		}
	case "DECIMAL":
		//temp := *(colVar.(*sql.NullInt64))
		//if temp.Valid {
		//	(*result)[index] = temp.Int64
		//} else {
		//	(*result)[index] = nil
		//}
		(*result)[index] = *(colVar.(*[]uint8))
	case "DATE":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TIME":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "YEAR":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "DATETIME":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TIMESTAMP":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "VARCHAR":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "CHAR":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "MEDIUMTEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "LONGTEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TINYTEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	case "TEXT":
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	default:
		var ok bool

		if colVar, ok = (*(colVar.(*interface{}))).(int64); ok {
			(*result)[index] = colVar
		} else if colVar, ok = (*(colVar.(*interface{}))).(string); ok {
			(*result)[index] = colVar
		} else if colVar, ok = (*(colVar.(*interface{}))).(float64); ok {
			(*result)[index] = colVar
		} else if colVar, ok = (*(colVar.(*interface{}))).([]uint8); ok {
			(*result)[index] = colVar
		} else {
			(*result)[index] = colVar
		}
	}
}
