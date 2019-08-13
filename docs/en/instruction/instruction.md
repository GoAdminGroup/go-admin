# Go-admin usage

Go-admin makes it easy to use in various web frameworks through various adapters.

## Example

Import ```$GOPATH/github.com/chenhg5/go-admin/examples/datamodel/admin.sql``` into the database.

Gin example:

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

    // initialize the go-admin engine object
	eng := engine.Default()

	// go-admin global configuration
	cfg := config.Config{
		DATABASE: []config.Database{
			{
			HOST:         "127.0.0.1",
			PORT:         "3306",
			USER:         "root",
			PWD:          "root",
			NAME:         "goadmin",
			MAX_IDLE_CON: 50,
			MAX_OPEN_CON: 150,
			DRIVER:       "mysql",
			},
			},
		DOMAIN: "localhost", // cookies will be set for this domain
		PREFIX: "admin",
		// STORE must be set and have write permissions, otherwise new administrator users cannot be added
		STORE: config.Store{
		    PATH:   "./uploads",
		    PREFIX: "uploads",
		},
		LANGUAGE: "en",
	}

	// Generators： see https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// Add configuration and plugins, call the `Use` method to mount to the web framework
	eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	r.Run(":9033")
}
```

The corresponding steps are annotated, the use is very simple, only need to:

- Define an adapter
- Set global configuration items
- Initialize the plugin
- Set up plugins and configurations
- Mount to the web framework

More examples: [https://github.com/chenhg5/go-admin/tree/master/examples](https://github.com/chenhg5/go-admin/tree/master/examples)

## Configuration

```go
package config

import (
	"html/template"
)

type Database struct {
	HOST         string
	PORT         string
	USER         string
	PWD          string
	NAME         string
	MAX_IDLE_CON int
	MAX_OPEN_CON int
	DRIVER       string

	FILE string
}

type Store struct {
	PATH   string
	PREFIX string
}

type Config struct {
	// Database configuration
	DATABASE []Database

	// Login domain name
	DOMAIN string

	// Website language
	LANGUAGE string

	// Global management prefix
	PREFIX string

	// Theme name
	THEME string

	// Uploads storage location
	STORE Store

	// Website title
	TITLE string

	// Sidebar logo
	LOGO template.HTML

	// Logo for collapsed sidebar
	MINILOGO template.HTML

	// Page to redirect to after login
	INDEX string
}

```

[Back to Contents](https://github.com/chenhg5/go-admin/blob/master/docs/en/index.md)<br>
[Next: Plugins](https://github.com/chenhg5/go-admin/blob/master/docs/en/instruction/plugins/plugins.md)
