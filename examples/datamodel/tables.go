package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/models"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
var Generators = map[string]models.TableGenerator{
	"user":    GetUserTable,
	"posts":   GetPostsTable,
	"authors": GetAuthorsTable,
}
