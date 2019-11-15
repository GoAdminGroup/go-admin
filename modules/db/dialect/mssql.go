// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package dialect

type mssql struct {
	commonDialect
}

func (mssql) GetName() string {
	return "mssql"
}
