// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
)

// SetColVarType set the column type.
func SetColVarType(colVar *[]interface{}, i int, typeName string) {
	dt := DT(typeName)
	switch {
	case Contains(dt, BoolTypeList):
		var s sql.NullBool
		(*colVar)[i] = &s
	case Contains(dt, IntTypeList):
		var s sql.NullInt64
		(*colVar)[i] = &s
	case Contains(dt, FloatTypeList):
		var s sql.NullFloat64
		(*colVar)[i] = &s
	case Contains(dt, UintTypeList):
		var s []uint8
		(*colVar)[i] = &s
	case Contains(dt, StringTypeList):
		var s sql.NullString
		(*colVar)[i] = &s
	default:
		var s interface{}
		(*colVar)[i] = &s
	}
}

// SetResultValue set the result value.
func SetResultValue(result *map[string]interface{}, index string, colVar interface{}, typeName string) {
	dt := DT(typeName)
	switch {
	case Contains(dt, BoolTypeList):
		temp := *(colVar.(*sql.NullBool))
		if temp.Valid {
			(*result)[index] = temp.Bool
		} else {
			(*result)[index] = nil
		}
	case Contains(dt, IntTypeList):
		temp := *(colVar.(*sql.NullInt64))
		if temp.Valid {
			(*result)[index] = temp.Int64
		} else {
			(*result)[index] = nil
		}
	case Contains(dt, FloatTypeList):
		temp := *(colVar.(*sql.NullFloat64))
		if temp.Valid {
			(*result)[index] = temp.Float64
		} else {
			(*result)[index] = nil
		}
	case Contains(dt, UintTypeList):
		(*result)[index] = *(colVar.(*[]uint8))
	case Contains(dt, StringTypeList):
		temp := *(colVar.(*sql.NullString))
		if temp.Valid {
			(*result)[index] = temp.String
		} else {
			(*result)[index] = nil
		}
	default:
		if colVar2, ok := colVar.(*interface{}); ok {
			if colVar, ok = (*colVar2).(int64); ok {
				(*result)[index] = colVar
			} else if colVar, ok = (*colVar2).(string); ok {
				(*result)[index] = colVar
			} else if colVar, ok = (*colVar2).(float64); ok {
				(*result)[index] = colVar
			} else if colVar, ok = (*colVar2).([]uint8); ok {
				(*result)[index] = colVar
			} else {
				(*result)[index] = colVar
			}
		}
	}
}
