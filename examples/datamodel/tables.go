package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "counties" => http://localhost:9033/admin/info/counties
//
var Generators = map[string]table.Generator{
	"counties": GetCountiesTable,
}
