package tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/GoAdminGroup/go-admin/tests/frameworks/fasthttp"
	"github.com/gavv/httpexpect"
	fasthttp2 "github.com/valyala/fasthttp"
)

func Cleaner(config config.DatabaseList) {

	checkStatement := ""

	if config.GetDefault().Driver != "sqlite" {
		if config.GetDefault().Dsn == "" {
			checkStatement = config.GetDefault().Name
		} else {
			checkStatement = config.GetDefault().Dsn
		}
	} else {
		if config.GetDefault().Dsn == "" {
			checkStatement = config.GetDefault().File
		} else {
			checkStatement = config.GetDefault().Dsn
		}
	}

	if !strings.Contains(checkStatement, "test") {
		panic("wrong database")
	}

	var allTables = [...]string{
		"goadmin_users",
		"goadmin_user_permissions",
		"goadmin_session",
		"goadmin_roles",
		"goadmin_role_users",
		"goadmin_role_permissions",
		"goadmin_role_menu",
		"goadmin_permissions",
		"goadmin_operation_log",
		"goadmin_menu",
	}
	var autoIncrementTable = [...]string{
		"goadmin_menu",
		"goadmin_permissions",
		"goadmin_roles",
		"goadmin_users",
	}
	var insertData = map[string][]dialect.H{
		"goadmin_users": {
			{"username": "admin", "name": "admin", "password": "$2a$10$TEDU/aUxLkr2wCxGxI62/.yOtzrzfv426DLLdyha9H2GpWRggB0di", "remember_token": "tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh"},
			{"username": "operator", "name": "operator", "password": "$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.", "remember_token": "tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh"},
		},
		"goadmin_roles": {
			{"name": "Administrator", "slug": "administrator"},
			{"name": "Operator", "slug": "operator"},
		},
		"goadmin_permissions": {
			{"name": "All permission", "slug": "*", "http_method": "", "http_path": "*"},
			{"name": "Dashboard", "slug": "dashboard", "http_method": "GET,PUT,POST,DELETE", "http_path": "/"},
		},
		"goadmin_menu": {
			{"parent_id": 0, "type": 1, "order": 2, "title": "Admin", "icon": "fa-tasks", "uri": ""},
			{"parent_id": 1, "type": 1, "order": 2, "title": "Users", "icon": "fa-users", "uri": "/info/manager"},
			{"parent_id": 0, "type": 1, "order": 3, "title": "test2 menu", "icon": "fa-angellist", "uri": "/example/test"},
			{"parent_id": 1, "type": 1, "order": 4, "title": "Permission", "icon": "fa-ban", "uri": "/info/permission"},
			{"parent_id": 1, "type": 1, "order": 5, "title": "Menu", "icon": "fa-bars", "uri": "/menu"},
			{"parent_id": 1, "type": 1, "order": 6, "title": "Operation log", "icon": "fa-history", "uri": "/info/op"},
			{"parent_id": 0, "type": 1, "order": 1, "title": "Dashboard", "icon": "fa-bar-chart", "uri": "/"},
			{"parent_id": 0, "type": 1, "order": 7, "title": "User", "icon": "fa-users", "uri": "/info/user"},
		},
		"goadmin_role_users": {
			{"user_id": 1, "role_id": 1},
			{"user_id": 2, "role_id": 2},
		},
		"goadmin_user_permissions": {
			{"user_id": 1, "permission_id": 1},
			{"user_id": 2, "permission_id": 2},
		},
		"goadmin_role_permissions": {
			{"role_id": 1, "permission_id": 1},
			{"role_id": 1, "permission_id": 2},
			{"role_id": 2, "permission_id": 2},
		},
		"goadmin_role_menu": {
			{"role_id": 1, "menu_id": 1},
			{"role_id": 1, "menu_id": 7},
			{"role_id": 2, "menu_id": 7},
			{"role_id": 1, "menu_id": 8},
			{"role_id": 2, "menu_id": 8},
			{"role_id": 1, "menu_id": 3},
		},
	}
	conn := db.GetConnectionByDriver(config.GetDefault().Driver).InitDB(config)
	// clean data
	for _, t := range allTables {
		_ = db.WithDriver(conn).Table(t).Delete()
	}
	// reset auto increment
	switch config.GetDefault().Driver {
	case db.DriverMysql:
		for _, t := range autoIncrementTable {
			checkErr(conn.Exec(`ALTER TABLE ` + t + ` AUTO_INCREMENT = 1`))
		}
	case db.DriverMssql:
		for _, t := range autoIncrementTable {
			checkErr(conn.Exec(`DBCC CHECKIDENT (` + t + `, RESEED, 0)`))
		}
	case db.DriverPostgresql:
		for _, t := range autoIncrementTable {
			checkErr(conn.Exec(`ALTER SEQUENCE ` + t + `_myid_seq RESTART WITH  1`))
		}
	case db.DriverSqlite:
		for _, t := range autoIncrementTable {
			checkErr(conn.Exec(`update sqlite_sequence set seq = 0 where name = '` + t + `'`))
		}
	}
	// insert data
	for t, data := range insertData {
		for _, d := range data {
			checkErr(db.WithDriver(conn).Table(t).Insert(d))
		}
	}
}

func BlackBoxTestSuitOfBuiltInTables(t *testing.T, fn HandlerGenFn, config config.DatabaseList, isFasthttp ...bool) {
	BlackBoxTestSuit(t, fn, config, nil, Cleaner, common.Test, isFasthttp...)
}

func checkErr(_ interface{}, err error) {
	if err != nil {
		panic(err)
	}
}

func BlackBoxTestSuit(t *testing.T, fn HandlerGenFn,
	config config.DatabaseList,
	gens table.GeneratorList,
	cleaner DataCleaner,
	tester Tester, isFasthttp ...bool) {
	// Clean Data
	cleaner(config)
	// Test
	if len(isFasthttp) > 0 && isFasthttp[0] {
		tester(httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewFastBinder(fasthttp.NewHandler(config, gens)),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		}))
	} else {
		tester(httpexpect.WithConfig(httpexpect.Config{
			Client: &http.Client{
				Transport: httpexpect.NewBinder(fn(config, gens)),
				Jar:       httpexpect.NewJar(),
			},
			Reporter: httpexpect.NewAssertReporter(t),
		}))
	}
}

type Tester func(e *httpexpect.Expect)
type DataCleaner func(config config.DatabaseList)
type HandlerGenFn func(config config.DatabaseList, gens table.GeneratorList) http.Handler
type FasthttpHandlerGenFn func(config config.DatabaseList, gens table.GeneratorList) fasthttp2.RequestHandler
