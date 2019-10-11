// Copyright 2019 GoAdmin.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"regexp"
	"strings"
)

func CommonQuery(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {

	rs, err := db.Query(query, args...)

	if err != nil {
		if rs != nil {
			_ = rs.Close()
		}
		panic(err)
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		if rs != nil {
			_ = rs.Close()
		}
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		if rs != nil {
			_ = rs.Close()
		}
		panic(err)
	}

	// TODO: regular expressions for sqlite, use the dialect module
	// tell the drive to reduce the performance loss
	results := make([]map[string]interface{}, 0)

	r, _ := regexp.Compile(`\\((.*)\\)`)
	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			typeName := strings.ToUpper(r.ReplaceAllString(typeVal[i].DatabaseTypeName(), ""))
			SetColVarType(&colVar, i, typeName)
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			_ = rs.Close()
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			typeName := strings.ToUpper(r.ReplaceAllString(typeVal[j].DatabaseTypeName(), ""))
			SetResultValue(&result, col[j], colVar[j], typeName)
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		if rs != nil {
			_ = rs.Close()
		}
		panic(err)
	}
	_ = rs.Close()
	return results, rs
}

func CommonExec(db *sql.DB, query string, args ...interface{}) sql.Result {

	rs, err := db.Exec(query, args...)
	if err != nil {
		panic(err.Error())
	}
	return rs
}
