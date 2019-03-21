// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	_ "github.com/chenhg5/go-admin/modules/db/mysql"
	_ "github.com/chenhg5/go-admin/modules/db/postgresql"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/config"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	rootPath    string
	outputPath  string
	driver      string
	port        string
	user        string
	password    string
	host        string
	name        string
	packageName string
)

func main() {

	if len(os.Args) > 1 {
        switch os.Args[1] {
        case "compile":

            if len(os.Args) > 2 {
                compileFlag := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
                compileFlag.StringVar(&rootPath, "path", "./template/adminlte/resource/pages/", "compile root path")
                compileFlag.StringVar(&outputPath, "path", "./template/adminlte/tmpl/template.go", "compile output path")
                compileFlag.Parse(os.Args[2:])
            } else {
                rootPath = "./template/adminlte/resource/pages/"
                outputPath = "./template/adminlte/tmpl/template.go"
            }

            CompileTmpl()

        case "generate":
            if len(os.Args) < 2 {
                panic("need options")
            }

            generateFlag := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
            generateFlag.StringVar(&driver, "d", "mysql", "database driver")
            generateFlag.StringVar(&port, "p", "3306", "database port")
            generateFlag.StringVar(&user, "u", "root", "database user")
            generateFlag.StringVar(&password, "P", "root", "database password")
            generateFlag.StringVar(&host, "h", "127.0.0.1", "database host")
            generateFlag.StringVar(&outputPath, "o", "template/", "database output path")
            generateFlag.StringVar(&name, "n", "", "database name")
            generateFlag.StringVar(&packageName, "pa", "main", "package name") //把原Pa改成pa，使与说明文档一致
            generateFlag.Parse(os.Args[2:])

            conn := db.GetConnectionByDriver(driver)
            if conn == nil {
                panic("Erorr db init")
            }
            cfg := map[string]config.Database{
                "default": config.Database{
                    HOST:         host,
                    PORT:         port,
                    USER:         user,
                    PWD:          password,
                    NAME:         name,
                    MAX_IDLE_CON: 50,
                    MAX_OPEN_CON: 150,
                    DRIVER:       driver,
                },
            }
            // step 1. test connection
            conn.InitDB(cfg)

            // step 2. show tables
            tables, _ := conn.Query("show tables")

            key := "Tables_in_" + name
            fieldField := "Field"
            typeField := "Type"
            if driver == "postgresql" {
                key = "tablename"
                fieldField = "column_name"
                typeField = "udt_name"
            }

            // step 3. show columns
            // step 4. generate file
            for i := 0; i < len(tables); i++ {
                if tables[i][key].(string) != "goadmin_menu" && tables[i][key].(string) != "goadmin_operation_log" &&
                tables[i][key].(string) != "goadmin_permissions" && tables[i][key].(string) != "goadmin_role_menu" &&
                tables[i][key].(string) != "goadmin_roles" && tables[i][key].(string) != "goadmin_session" &&
                tables[i][key].(string) != "goadmin_users" && tables[i][key].(string) != "goadmin_role_permissions" &&
                tables[i][key].(string) != "goadmin_role_users" && tables[i][key].(string) != "goadmin_user_permissions" {
                    GenerateFile(tables[i][key].(string), conn, fieldField, typeField)
                }
            }
        default:
            panic("wrong option")
        }
	}
}

func CompileTmpl() {
	content := `package tmpl

var List = map[string]string{`

	content = GetContentFromDir(content, rootPath)

	content += `}`

	ioutil.WriteFile(outputPath, []byte(content), 0644)
}

func GetContentFromDir(content string, dirPath string) string {
	files, _ := ioutil.ReadDir(dirPath)

	for _, f := range files {

		if f.IsDir() {
			content = GetContentFromDir(content, dirPath+f.Name()+"/")
			continue
		}

		b, err := ioutil.ReadFile(dirPath + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)

		suffix := path.Ext(f.Name())
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			fmt.Println(dirPath + f.Name())
			content += `"` + strings.Replace(dirPath, rootPath, "", -1) + onlyName + `":` + "`" + str + "`,"
		}
	}

	return content
}

func GenerateFile(table string, conn db.Connection, fieldField, typeField string) {
	columnsModel, _ := conn.ShowColumns(table)

	content := `package ` + packageName + `

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func Get` + strings.Title(table) + `Table() (` + table + `Table models.Table) {

	` + table + `Table.Info.FieldList = []types.Field{`

	for _, model := range columnsModel {
		content += `{
			Head:     "` + strings.Title(model[fieldField].(string)) + `",
			Field:    "` + model[fieldField].(string) + `",
			TypeName: "` + GetType(model[typeField].(string)) + `",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},`
	}

	content += `}

	` + table + `Table.Info.Table = "` + table + `"
	` + table + `Table.Info.Title = "` + strings.Title(table) + `"
	` + table + `Table.Info.Description = "` + strings.Title(table) + `"

	` + table + `Table.Form.FormList = []types.Form{`

	for _, model := range columnsModel {

		formType := "text"
		if model[fieldField].(string) == "id" {
			formType = "default"
		}

		content += `{
			Head:     "` + strings.Title(model[fieldField].(string)) + `",
			Field:    "` + model[fieldField].(string) + `",
			TypeName: "` + GetType(model[typeField].(string)) + `",
			Default:  "",
			Editable: false,
			FormType: "` + formType + `",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},`
	}

	content += `	}

	` + table + `Table.Form.Table = "` + table + `"
	` + table + `Table.Form.Title = "` + strings.Title(table) + `"
	` + table + `Table.Form.Description = "` + strings.Title(table) + `"

	` + table + `Table.ConnectionDriver = "` + driver + `"

	return
}`

	fmt.Println(outputPath + "/" + table + ".go")
	ioutil.WriteFile(outputPath+table+".go", []byte(content), 0644)
}

func GetType(typeName string) string {
	r, _ := regexp.Compile("\\(.*\\)")
	typeName = r.ReplaceAllString(typeName, "")
	return strings.ToLower(strings.Replace(typeName, " unsigned", "", -1))
}
