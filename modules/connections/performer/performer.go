package performer

import (
	"database/sql"
	"github.com/chenhg5/go-admin/modules/connections/converter"
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

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			converter.SetColVarType(&colVar, i, typeVal[i].DatabaseTypeName())
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			rs.Close()
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			converter.SetResultValue(&result, col[j], colVar[j], typeVal[j].DatabaseTypeName())
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
