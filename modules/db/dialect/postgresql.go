// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package dialect

type postgresql struct {
	commonDialect
}

func (postgresql) GetName() string {
	return "postgresql"
}

func (postgresql) ShowTables() string {
	return "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"
}
