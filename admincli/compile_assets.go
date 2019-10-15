package main

import (
	"github.com/jteeuwen/go-bindata"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func compileAsset(rootPath, outputPath, packageName string) {
	cfg := bindata.NewConfig()
	cfg.Package = packageName
	cfg.Output = outputPath + "assets.go"
	cfg.Input = make([]bindata.InputConfig, 0)
	cfg.Input = append(cfg.Input, parseInput(rootPath+"..."))
	checkError(bindata.Translate(cfg))

	rootPathArr := strings.Split(rootPath, "assets")
	if len(rootPathArr) > 0 {
		listContent := `package ` + packageName + `

var AssetsList = []string{
`
		fileNames, err := getAllFiles(rootPath)

		if err != nil {
			return
		}

		for _, name := range fileNames {
			listContent += `	"` + rootPathArr[1] + strings.Replace(name, rootPath, "", -1)[1:] + `",
`
		}

		listContent += `
}`

		err = ioutil.WriteFile(outputPath+"/assets_list.go", []byte(listContent), 0644)
		if err != nil {
			return
		}
	}
}

func getAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			_, _ = getAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := getAllFiles(table)
		files = append(files, temp...)
	}

	return files, nil
}

func parseInput(path string) bindata.InputConfig {
	if strings.HasSuffix(path, "/...") {
		return bindata.InputConfig{
			Path:      filepath.Clean(path[:len(path)-4]),
			Recursive: true,
		}
	} else {
		return bindata.InputConfig{
			Path:      filepath.Clean(path),
			Recursive: false,
		}
	}
}
