// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package main

import (
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"

	"fmt"
	cli "github.com/jawher/mow.cli"
	"github.com/mgutz/ansi"
	"os"
	"runtime"
	"runtime/debug"
)

func main() {

	var verbose *bool

	defer func() {
		if err := recover(); err != nil {
			if errs, ok := err.(error); ok {
				fmt.Println()
				if runtime.GOOS == "windows" && errs.Error() == "Incorrect function." {
					fmt.Println(ansi.Color("GoAdmin CLI error: CLI has not supported MINGW64 for now, "+
						"please use cmd terminal instead.", "red"))
					fmt.Println("know more here: http://forum.go-admin.cn/threads/2")
				} else {
					fmt.Println(ansi.Color("GoAdmin CLI error: "+errs.Error(), "red"))

					if *verbose {
						fmt.Println(string(debug.Stack()))
					}
				}
				fmt.Println()
			}
		}
	}()

	app := cli.App("adm", "GoAdmin CLI tool for developing and generating")

	app.Spec = "[-v]"

	verbose = app.BoolOpt("v verbose", false, "debug info output")
	// quiet

	app.Command("-V version", "display this application version", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			cliInfo()
		}
	})

	app.Command("combine", "combine assets", func(cmd *cli.Cmd) {
		cmd.Command("css", "combine css assets", func(cmd *cli.Cmd) {
			var (
				rootPath   = cmd.StringOpt("s src", "./resource/assets/src/css/combine/", "css src path")
				outputPath = cmd.StringOpt("d dist", "./resource/assets/dist/css/all.min.css", "css output path")
			)

			cmd.Action = func() {
				cssMinifier(*rootPath, *outputPath)
			}
		})

		cmd.Command("js", "combine js assets", func(cmd *cli.Cmd) {
			var (
				rootPath   = cmd.StringOpt("s src", "./resource/assets/src/js/combine/", "js src path")
				outputPath = cmd.StringOpt("d dist", "./resource/assets/dist/js/all.min.js", "js output path")
			)

			cmd.Action = func() {
				jsMinifier(*rootPath, *outputPath)
			}
		})
	})

	app.Command("compile", "compile template files or assets to one go file", func(cmd *cli.Cmd) {
		cmd.Command("tpl", "compile template files", func(cmd *cli.Cmd) {
			var (
				rootPath    = cmd.StringOpt("s src", "./resource/pages/", "template files src path")
				outputPath  = cmd.StringOpt("d dist", "./template.go", "compile file output path")
				packageName = cmd.StringOpt("p package", "newTmplTheme", "the package name")
			)

			cmd.Action = func() {
				compileTmpl(*rootPath, *outputPath, *packageName)
			}
		})

		cmd.Command("asset", "compile assets", func(cmd *cli.Cmd) {
			var (
				rootPath    = cmd.StringOpt("s src", "./resource/assets/dist/", "assets root path")
				outputPath  = cmd.StringOpt("d dist", "./resource/", "compile file output path")
				packageName = cmd.StringOpt("p package", "resource", "package name of the output golang file")
			)

			cmd.Action = func() {
				compileAsset(*rootPath, *outputPath, *packageName)
			}
		})
	})

	app.Command("develop", "commands for developing", func(cmd *cli.Cmd) {
		cmd.Command("tpl", "generate a theme project from a remote template", func(cmd *cli.Cmd) {
			var (
				moduleName = cmd.StringOpt("m module", "github.com/GoAdminGroup/themes/newTmpl", "the module name of your theme")
				themeName  = cmd.StringOpt("n name", "newTmplTheme", "the name of your theme")
			)

			cmd.Action = func() {
				getThemeTemplate(*moduleName, *themeName)
			}
		})
	})

	app.Command("generate", "generate table model files", func(cmd *cli.Cmd) {

		var (
			config = cmd.StringOpt("c config", "", "config ini path")
			lang   = cmd.StringOpt("l language", "en", "language")
		)

		cmd.Action = func() {
			setDefaultLangSet(*lang)
			generating(*config)
		}
	})

	app.Command("add", "generate user/permission/roles", func(cmd *cli.Cmd) {

		cmd.Command("user", "generate users", func(cmd *cli.Cmd) {
			var (
				config = cmd.StringOpt("c config", "", "config ini path")
				lang   = cmd.StringOpt("l language", "en", "language")
			)

			cmd.Action = func() {
				setDefaultLangSet(*lang)
				addUser(*config)
			}
		})

		cmd.Command("permission", "generate permissions of table", func(cmd *cli.Cmd) {
			var (
				config = cmd.StringOpt("c config", "", "config ini path")
				lang   = cmd.StringOpt("l language", "en", "language")
			)

			cmd.Action = func() {
				setDefaultLangSet(*lang)
				addPermission(*config)
			}
		})
	})

	_ = app.Run(os.Args)
}
