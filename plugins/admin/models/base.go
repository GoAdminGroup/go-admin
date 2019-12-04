package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
)

// Base is base model structure.
type Base struct {
	TableName string

	Conn db.Connection
}

func (b Base) SetConn(con db.Connection) Base {
	b.Conn = con
	return b
}

func (b Base) Table(table string) *db.SQL {
	return db.Table(table).WithDriver(b.Conn)
}
