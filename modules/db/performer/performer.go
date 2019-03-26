// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package performer

import (
	"database/sql"
	"fmt"
	"github.com/chenhg5/go-admin/modules/db/converter"
	"regexp"
	"strings"
)

func Query(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {

	rs, err := db.Query(query, args...)

	if err != nil {
		if rs != nil {
			rs.Close()
		}
		panic(err)
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		if rs != nil {
			rs.Close()
		}
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		if rs != nil {
			rs.Close()
		}
		panic(err)
	}

	// TODO: regular expressions for sqlite, use the dialect module
	// tell the drive to reduce the performance loss
	results := make([]map[string]interface{}, 0)

	r, _ := regexp.Compile("\\((.*)\\)")
	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			typeName := strings.ToUpper(r.ReplaceAllString(typeVal[i].DatabaseTypeName(), ""))
			converter.SetColVarType(&colVar, i, typeName)
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			rs.Close()
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			typeName := strings.ToUpper(r.ReplaceAllString(typeVal[j].DatabaseTypeName(), ""))
			converter.SetResultValue(&result, col[j], colVar[j], typeName)
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		if rs != nil {
			rs.Close()
		}
		panic(err)
	}
	rs.Close()
	return results, rs
}

func Exec(db *sql.DB, query string, args ...interface{}) sql.Result {

	rs, err := db.Exec(query, args...)
	if err != nil {
		panic(err.Error())
	}
	return rs
}
