// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package dialect

type mysql struct {
	commonDialect
}

func (mysql) GetName() string {
	return "mysql"
}

func (mysql) ShowColumns(table string) string {
	return "show columns in " + table
}

func (mysql) ShowTables() string {
	return "show tables"
}
