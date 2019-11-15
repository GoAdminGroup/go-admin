// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package dialect

type sqlite struct {
	commonDialect
}

func (sqlite) GetName() string {
	return "sqlite"
}

func (sqlite) ShowColumns(table string) string {
	return "PRAGMA table_info(" + table + ");"
}

func (sqlite) ShowTables() string {
	return "SELECT name as tablename FROM sqlite_master WHERE type ='table'"
}
