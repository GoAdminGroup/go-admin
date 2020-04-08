// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package dialect

import (
	"fmt"
	"strings"
)

type postgresql struct {
	commonDialect
}

func (postgresql) GetName() string {
	return "postgresql"
}

func (postgresql) ShowTables() string {
	return "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"
}

func (postgresql) ShowColumns(table string) string {
	tableArr := strings.Split(table, ".")
	if len(tableArr) > 1 {
		return fmt.Sprintf("select * from information_schema.columns where table_name = '%s' and table_schema = '%s'", tableArr[1], tableArr[0])
	} else {
		return fmt.Sprintf("select * from information_schema.columns where table_name = '%s'", table)
	}
}
