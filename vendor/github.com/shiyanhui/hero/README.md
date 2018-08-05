# Hero

Hero is a handy, fast and powerful go template engine, which pre-compiles the html templates to go code.
It has been used in production environment in [bthub.io](http://bthub.io).

[![GoDoc](https://godoc.org/github.com/shiyanhui/hero?status.svg)](https://godoc.org/github.com/shiyanhui/hero)
[![Go Report Card](https://goreportcard.com/badge/github.com/shiyanhui/hero)](https://goreportcard.com/report/github.com/shiyanhui/hero)
[![Travis CI](https://travis-ci.org/shiyanhui/hero.svg?branch=master)](https://travis-ci.org/shiyanhui/hero.svg?branch=master)

[中文文档](https://github.com/shiyanhui/hero/blob/master/README_CN.md)

- [Features](#features)
- [Install](#install)
- [Usage](#usage)
- [Quick Start](#quick-start)
- [Template Syntax](#template-syntax)
- [License](#license)

## Features

- High performance.
- Easy to use.
- Powerful. template `Extend` and `Include` supported.
- Auto compiling when files change.

## Performance

Hero is the fastest and least-memory used among currently known template engines
in the benchmark. The data of chart comes from [https://github.com/SlinSo/goTemplateBenchmark](https://github.com/SlinSo/goTemplateBenchmark#full-featured-template-engines-2).
You can find more details and benchmarks from that project.

<img src='http://i.imgur.com/93D7T5C.png' width="600">
<img src='http://i.imgur.com/EIGtYyF.png' width="600">

## Install

```shell
go get github.com/shiyanhui/hero/hero

# Hero needs `goimports` to format the generated codes.
go get golang.org/x/tools/cmd/goimports
```

## Usage

```shell
hero [options]

  -source string
        the html template file or dir (default "./")
  -dest string
        generated golang files dir, it will be the same with source if not set
  -extensions string
        source file extensions, comma splitted if many (default ".html")
  -pkgname template
        the generated template package name, default is template (default "template")
  -watch
        whether automatically compile when the source files change

example:
	hero -source="./"
	hero -source="$GOPATH/src/app/template" -dest="./" -extensions=".html,.htm" -pkgname="t" -watch
```

## Quick Start

Assume that we are going to render a user list `userlist.html`. `index.html`
is the layout, and `user.html` is an item in the list.

And assumes that they are all under `$GOPATH/src/app/template`

### index.html

```html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
    </head>

    <body>
        <%@ body { %>
        <% } %>
    </body>
</html>
```

### userlist.html

```html
<%: func UserList(userList []string, buffer *bytes.Buffer) %>

<%~ "index.html" %>

<%@ body { %>
    <% for _, user := range userList { %>
        <ul>
            <%+ "user.html" %>
        </ul>
    <% } %>
<% } %>
```

### user.html

```html
<li>
    <%= user %>
</li>
```

Then we compile the templates to go code.

```shell
hero -source="$GOPATH/src/app/template"
```

We will get three new `.go` files under `$GOPATH/src/app/template`,
i.e. `index.html.go`, `user.html.go` and `userlist.html.go`.

Then we write a http server in `$GOPATH/src/app/main.go`.

### main.go

```go
package main

import (
    "bytes"
    "net/http"

    "app/template"
)

func main() {
    http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
        var userList = []string {
            "Alice",
            "Bob",
            "Tom",
        }

        // Had better use buffer pool. Hero exports `GetBuffer` and `PutBuffer` for this.
        //
        // For convenience, hero also supports `io.Writer`. For example, you can also define
        // the function to `func UserList(userList []string, w io.Writer) (int, error)`,
        // and then:
        //
        //   template.UserList(userList, w)
        //
        buffer := new(bytes.Buffer)
        template.UserList(userList, buffer)
        w.Write(buffer.Bytes())
    })

    http.ListenAndServe(":8080", nil)
}
```

At last, start the server and visit `http://localhost:8080/users` in your browser, we will get what we want!

## Template syntax

There are only nine necessary kinds of statements, which are:

- Function Definition `<%: func define %>`
  - Function definition statement defines the function which represents an html file.
  - The type of the last parameter in the function defined should be `*bytes.Buffer` for manual buffer management or `io.Writer` for automatic buffer management (
    note: if using `io.Writer` you may optionally specify return values `(int, error)` to handle the result of `io.Writer.Write`). Hero will identify the parameter name
    automaticly.
  - Example:
    - `<%: func UserList(userList []string, buffer *bytes.Buffer) %>`
    - `<%: func UserList(userList []string, w io.Writer) %>`
    - `<%: func UserList(userList []string, w io.Writer) (int, error) %>`

- Extend `<%~ "parent template" %>`
  - Extend statement states the parent template the current template extends.
  - The parent template should be quoted with `""`.
  - Example: `<%~ "index.html" >`, which we have mentioned in quick start, too.

- Include `<%+ "sub template" %>`
  - Include statement includes a sub-template to the current template. It works like `#include` in `C++`.
  - The sub-template should be quoted with `""`.
  - Example: `<%+ "user.html" >`, which we also have mentioned in quick start.

- Import `<%! go code %>`
  - Import statement imports the packages used in the defined function, and it also contains everything that is outside of the defined function.
  - Import statement will NOT be inherited by child template.
  - Example:

    ```go
    <%!
      import (
          "fmt"
          "strings"
      )

      var a int

      const b = "hello, world"

      func Add(a, b int) int {
          return a + b
      }

      type S struct {
          Name string
      }

      func (s S) String() string {
          return s.Name
      }
    %>
    ```

- Block `<%@ blockName { %> <% } %>`

  - Block statement represents a block. Child template overwrites blocks to extend parent template.

  - Example:

    ```html
    <!DOCTYPE html>
    <html>
        <head>
            <meta charset="utf-8">
        </head>

        <body>
            <%@ body { %>
            <% } %>
        </body>
    </html>
    ```

- Code `<% go code %>`

  - Code statement states all code inside the defined function. It's just go code.

  - Example:

    ```go
    <% for _, user := userList { %>
        <% if user != "Alice" { %>
            <%= user %>
        <% } %>
    <% } %>

    <%
        a, b := 1, 2
        c := Add(a, b)
    %>
    ```

- Raw Value `<%==[t] variable %>`

  - Raw Value statement will convert the variable to string.
  - `t` is the type of variable, hero will find suitable converting method by `t`. Candidates of `t` are:
    - `b`: bool
    - `i`: int, int8, int16, int32, int64
    - `u`: byte, uint, uint8, uint16, uint32, uint64
    - `f`: float32, float64
    - `s`: string
    - `bs`: []byte
    - `v`: interface

    Note:
    - If `t` is not set, the value of `t` is `s`.
    - Had better not use `v`, cause when `t=v`, the converting method is `fmt.Sprintf("%v", variable)` and it is very slow.
  - Example:

    ```go
    <%== "hello" %>
    <%==i 34  %>
    <%==u Add(a, b) %>
    <%==s user.Name %>
    ```

- Escaped Value `<%=[t] variable %>`

  - Escaped Value statement is similar with Raw Value statement, but after converting, it will be escaped it with `html.EscapesString`.
  - `t` is the same as in `Raw Value Statement`.
  - Example:

    ```go
    <%= a %>
    <%=i a + b %>
    <%=u Add(a, b) %>
    <%=bs []byte{1, 2} %>
    ```

- Note `<%# note %>`

  - Note statement add notes to the template.
  - It will not be added to the generated go source.
  - Example: `<# this is just a note example>`.

## License

Hero is licensed under the Apache License.
