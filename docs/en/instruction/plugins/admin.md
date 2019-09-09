# Admin Plugin Usage

The admin plugin can help you to quickly generate a database data table for adding, deleting, and changing database data tables.

## Quick start
- Generate a configuration file corresponding to the data table
- Set access routing
- Initialize and load in the engine
- Set access menu

### Generate the configuration file

Suppose you have a data table users in your database, for example:

```sql
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gender` tinyint(4) DEFAULT NULL,
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3635 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

Use the included command line tools to help you quickly generate configuration files, for example:

- Installation

```bash
go install github.com/chenhg5/go-admin/admincli
```

- Generation

```bash
admincli generate -h=127.0.0.1 -p=3306 -u=root -P=root -pa=main -n=goadmin -o=./

-h  host
-p  port
-u  username
-P  password
-n  database name
-pa package name
-o  output file path

```

The file ```users.go``` will be generated. This is the configuration file corresponding to the data table. Configuration details described later.

### Set access routing

After generating the configuration file, you need to set the route to access the data of this data table, for example:

```go
package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.DOMAIN}}:{{PORT}}/{{config.PREFIX}}/info/{{key}}
//
// example:
//
// "posts"   => http://localhost:9033/admin/info/posts
// "authors" => http://localhost:9033/admin/info/authors
//
var Generators = table.GeneratorList{
	"user":    GetUserTable,
}
```

```"user"``` is the prefix of the table info url ```GetUserTable``` is the method in the configuration file

### Initialize and load the engine

To initialize, you need to call the ```NewAdmin``` method, and then pass the ```Generators``` defined above. Then call the engine's ```AddPlugins``` method to load the engine.

```go
package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chenhg5/go-admin/adapter/gin" // adapter must be imported, if not - you have to implement it yourself
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/examples/datamodel"
)

func main() {
	r := gin.Default()
	eng := engine.Default()
	cfg := config.Config{}

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	_ = eng.AddConfig(cfg).
		AddPlugins(adminPlugin). // Load plugin
		Use(r)

	_ = r.Run(":9033")
}
```

### Set access routing

After running visit the login URL, go to the menu management page. The management menu for setting up the data table can be accessed in the sidebar.

## Profile introduction

Configuration file can looks like:

```go
package main

import (
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
)

func GetUsersTable() (usersTable table.Table) {
	
	usersTable = table.NewDefaultTable(table.DefaultConfig)

	usersTable.GetInfo().FieldList = []types.Field{}

	usersTable.GetInfo().Table = "users"
	usersTable.GetInfo().Title = "Users"
	usersTable.GetInfo().Description = "Users"

	usersTable.GetForm().FormList = []types.Form{}

	usersTable.GetForm().Table = "users"
	usersTable.GetForm().Title = "Users"
	usersTable.GetForm().Description = "Users"

	return
}
```
```GetUsersTable``` returns the object with type ```models.Table```. See the ```models.Table``` definition below:

```go
type Table struct {
	Info             types.InfoPanel
	Form             types.FormPanel
	ConnectionDriver string
}
```

Both ```Info``` and ```Form``` are UI components responsible for displaying table title with description and editing new data respectively

- ```Info```

![](https://ws4.sinaimg.cn/large/006tNbRwly1fxoy26qnc5j31y60u0q91.jpg)

- ```Form```

![](https://ws1.sinaimg.cn/large/006tNbRwly1fxoy2w3cobj318k0ooabv.jpg)

### Info Panel

```go
type InfoPanel struct {
	FieldList   []Field
	Table       string
	Title       string
	Description string
}

type Field struct {
	FilterFn  FieldFilterFn // filter function
	Field    string        // field name
	TypeName string        // field type name
	Head     string        // title
	Sortable bool
	Filter   bool
}
```

### Form Panel

```go
type FormPanel struct {
	FormList   []Form
	Table       string
	Title       string
	Description string
}

type Form struct {
	Field    string  // field name
	TypeName string  // field type name
	Head     string  // title
	Default  string
	Editable bool
	FormType string
	Value    string
	Options  []map[string]string
	FilterFn  FieldFilterFn // filter function
}
```

Currently supported form types are:
- Default
- Text
- SelectSingle
- Select
- IconPicker
- SelectBox
- File
- Password
- RichText
- Datetime
- Radio
- Email
- Url
- Ip
- Color
- Currency
- Number

For example:

```go

import "github.com/chenhg5/go-admin/template/types/form"

...
FormType: form.File,
...

```

For the selection type: `SelectSingle`, `Select`, `SelectBox` you must specify `Options` field：

```go
...
Options: []map[string]string{
	{
        "field": "name", // the name of the field
        "value": "Adam", // value you can select from
    },{
        "field": "name",
        "value": "John",
    },
}
...
```

### Filter function `FilterFn`

```go
// RowModel contains ID and value of the single query result.
type RowModel struct {
	ID    int64
	Value string
}

// FieldFilterFn is filter function of data.
type FieldFilterFn func(value RowModel) interface{}
```

The filter function receives a parameter, RowModel, which represents the current edit target line, contains the id and the displayed value.
The return value of the filter function is the default value displayed by the final form line.

In the table, you can customize the html return.
In the form, for non-selected form types, you **must** return `string`. For single-select, multi-select, etc., select form type, return `[]string`.

[Back to Contents](https://github.com/chenhg5/go-admin/blob/master/docs/en/index.md)<br>
[Previous: Plugins](https://github.com/chenhg5/go-admin/blob/master/docs/en/instruction/plugins/plugins.md)<br>
[Next page: Custom page](https://github.com/chenhg5/go-admin/blob/master/docs/en/instruction/pages/pages.md)
