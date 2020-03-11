package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

type Guard struct {
	services  service.List
	conn      db.Connection
	tableList table.GeneratorList
}

func New(s service.List, c db.Connection, t table.GeneratorList) *Guard {
	return &Guard{
		services:  s,
		conn:      c,
		tableList: t,
	}
}

func (g *Guard) table(ctx *context.Context) (table.Table, string) {
	prefix := ctx.Query(constant.PrefixKey)
	return g.tableList[prefix](ctx), prefix
}
