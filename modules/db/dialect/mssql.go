// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package dialect

import "fmt"

type mssql struct {
	commonDialect
}

func (mssql) GetName() string {
	return "mssql"
}

func (mssql) ShowColumns(table string) string {
	return fmt.Sprintf("select column_name, data_type from information_schema.columns where table_name = '%s'", table)
}

func (mssql) ShowTables() string {
	return "select * from information_schema.TABLES"
}
