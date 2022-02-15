package models

import (
	"database/sql"

	"github.com/GoAdminGroup/go-admin/modules/db"
)

type Context struct {
	Conn db.Connection
	Tx   *sql.Tx
}

func (ctx *Context) WithConn(conn db.Connection) {
	ctx.Conn = conn
}

func (ctx *Context) WithTx(tx *sql.Tx) {
	ctx.Tx = tx
}

// Base is base model structure.
type Base struct {
	TableName string

	Conn db.Connection
	Tx   *sql.Tx
}

func (b Base) SetConn(con db.Connection) Base {
	b.Conn = con
	return b
}

func (b Base) Table(table string) *db.SQL {
	return db.Table(table).WithDriver(b.Conn)
}

type TableNamer interface {
	TableName() string
}
