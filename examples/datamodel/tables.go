package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
var Generators = map[string]table.Generator{
	"user":    GetUserTable,
	"posts":   GetPostsTable,
	"authors": GetAuthorsTable,
}
