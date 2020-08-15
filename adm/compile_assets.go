package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jteeuwen/go-bindata"
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
		pathsContent := `package ` + packageName + `

var AssetPaths = map[string]string{
`

		fileNames, err := getAllFiles(rootPath)

		if err != nil {
			return
		}

		for _, name := range fileNames {
			listContent += `	"` + rootPathArr[1] + strings.ReplaceAll(name, rootPath, "")[1:] + `",
`
			ext := filepath.Ext(name)
			if ext == ".css" || ext == ".js" {
				fileName := filepath.Base(name)
				reg, _ := regexp.Compile(".min.(.*?)" + ext)
				pathsContent += `	"` + reg.ReplaceAllString(fileName, ".min"+ext) + `":"` +
					rootPathArr[1] + strings.ReplaceAll(name, rootPath, "")[1:] + `",
`
			}
		}

		pathsContent += `
}`

		listContent += `
}`

		err = ioutil.WriteFile(outputPath+"/assets_list.go", []byte(listContent), 0644)
		if err != nil {
			return
		}
		err = ioutil.WriteFile(outputPath+"/assets_path.go", []byte(pathsContent), 0644)
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
		if fi.IsDir() {
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			_, _ = getAllFiles(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

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
	}
	return bindata.InputConfig{
		Path:      filepath.Clean(path),
		Recursive: false,
	}
}
