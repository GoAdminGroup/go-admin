# Hero

Hero是一个高性能、强大并且易用的go模板引擎，工作原理是把模板预编译为go代码。Hero目前已经在[bthub.io](http://bthub.io)的线上环境上使用。

[![GoDoc](https://godoc.org/github.com/shiyanhui/hero?status.svg)](https://godoc.org/github.com/shiyanhui/hero)
[![Go Report Card](https://goreportcard.com/badge/github.com/shiyanhui/hero)](https://goreportcard.com/report/github.com/shiyanhui/hero)

- [Features](#features)
- [Install](#install)
- [Usage](#usage)
- [Quick Start](#quick-start)
- [Template Syntax](#template-syntax)
- [License](#license)

## Features

- 高性能.
- 非常易用.
- 功能强大，支持模板继承和模板include.
- 自动编译.

## Performance

Hero在目前已知的模板引擎中是速度是最快的，并且内存使用是最少的。下面Benchmark图表的数据
来源于[https://github.com/SlinSo/goTemplateBenchmark](https://github.com/SlinSo/goTemplateBenchmark#full-featured-template-engines-2)，
关于更多的细节和Benchmarks请到上述项目中查看。

<img src='http://i.imgur.com/93D7T5C.png' width="600">
<img src='http://i.imgur.com/EIGtYyF.png' width="600">

## Install

    go get github.com/shiyanhui/hero
    go get github.com/shiyanhui/hero/hero

    // Hero需要goimports处理生成的go代码，所以需要安装goimports.
    go get golang.org/x/tools/cmd/goimports

## Usage

```shell
hero [options]

options:
	- source:     模板目录，默认为当前目录
	- dest:       生成的go代码的目录，如果没有设置的话，和source一样
	- pkgname:    生成的go代码包的名称，默认为template
  - extensions: source文件的后缀, 如果有多个则用英文逗号隔开, 默认为.html
	- watch:      是否监控模板文件改动并自动编译

example:
	hero -source="./"
	hero -source="$GOPATH/src/app/template" -watch
```

## Quick Start

假设我们现在要渲染一个用户列表模板`userlist.html`, 它继承自`index.html`, 并且一个用户的模板是`user.html`. 我们还假设所有的模板都在`$GOPATH/src/app/template`目录下。

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

然后我们编译这些模板:

```shell
hero -source="$GOPATH/src/app/template"
```

编译后，我们将在同一个目录下得到三个go文件，分别是`index.html.go`, `user.html.go`
和 `userlist.html.go`, 然后我们在http server里边去调用模板：

### main.go

```go
package main

import (
	"bytes"
	"net/http"

	"github.com/shiyanhui/hero/examples/app/template"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		var userList = []string{
			"Alice",
			"Bob",
			"Tom",
		}

		buffer := new(bytes.Buffer)
		template.UserList(userList, buffer)

		w.Write(buffer.Bytes())
	})

	http.ListenAndServe(":8080", nil)
}
```

最后，运行这个http server，访问`http://localhost:8080/users`，我们就能得到我们期待的结果了！

## Template syntax

Hero总共有九种语句，他们分别是：

- 函数定义语句 `<%: func define %>`
  - 该语句定义了该模板所对应的函数，如果一个模板中没有函数定义语句，那么最终结果不会生成对应的函数。
  - 该函数最后一个参数必须为`*bytes.Buffer`或者`io.Writer`, hero会自动识别该参数的名字，并把把结果写到该参数里。
  - 例:
    - `<%: func UserList(userList []string, buffer *bytes.Buffer) %>`
    - `<%: func UserList(userList []string, w io.Writer) %>`
    - `<%: func UserList(userList []string, w io.Writer) (int, error) %>`

- 模板继承语句 `<%~ "parent template" %>`
  - 该语句声明要继承的模板。
  - 例: `<%~ "index.html" >`

- 模板include语句 `<%+ "sub template" %>`
  - 该语句把要include的模板加载进该模板，工作原理和`C++`中的`#include`有点类似。
  - 例: `<%+ "user.html" >`

- 包导入语句 `<%! go code %>`
  - 该语句用来声明所有在函数外的代码，包括依赖包导入、全局变量、const等。

  - 该语句不会被子模板所继承

  - 例:

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

- 块语句 `<%@ blockName { %> <% } %>`

  - 块语句是用来在子模板中重写父模中的同名块，进而实现模板的继承。

  - 例:

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

- Go代码语句 `<% go code %>`

  - 该语句定义了函数内部的代码部分。

  - 例:

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

- 原生值语句 `<%==[t] variable %>`

  - 该语句把变量转换为string。

  - `t`是变量的类型，hero会自动根据`t`来选择转换函数。`t`的待选值有:
    - `b`: bool
    - `i`: int, int8, int16, int32, int64
    - `u`: byte, uint, uint8, uint16, uint32, uint64
    - `f`: float32, float64
    - `s`: string
    - `bs`: []byte
    - `v`: interface

    注意：
    - 如果`t`没有设置，那么`t`默认为`s`.
    - 最好不要使用`v`，因为其对应的转换函数为`fmt.Sprintf("%v", variable)`，该函数很慢。

  - 例:

    ```go
    <%== "hello" %>
    <%==i 34  %>
    <%==u Add(a, b) %>
    <%==s user.Name %>
    ```

- 转义值语句 `<%= statement %>`

  - 该语句把变量转换为string后，又通过`html.EscapesString`记性转义。
  - `t`跟上面原生值语句中的`t`一样。
  - 例:

    ```go
    <%= a %>
    <%= a + b %>
    <%= Add(a, b) %>
    <%= user.Name %>
    ```

- 注释语句 `<%# note %>`

  - 该语句注释相关模板，注释不会被生成到go代码里边去。
  - 例: `<# 这是一个注释 >`.

## License

Hero is licensed under the Apache License.
