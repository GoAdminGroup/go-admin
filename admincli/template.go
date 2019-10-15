package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func getThemeTemplate(moduleName, themeName string) {

	defer func() {
		_ = os.Remove("tmp.zip")
	}()

	url := "http://file.go-admin.cn/go_admin/template/template.zip"

	req, err := http.NewRequest("GET", url, nil)

	checkError(err)

	res, err := http.DefaultClient.Do(req)

	checkError(err)

	defer func() {
		_ = res.Body.Close()
	}()

	file, err := os.Create("tmp.zip")

	checkError(err)

	_, err = io.Copy(file, res.Body)

	checkError(err)

	unzipDir("tmp.zip", ".")

	checkError(os.Rename("./template", "./"+themeName))

	replaceContents("./"+themeName, moduleName, themeName)

	if runtime.GOOS == "darwin" {
		checkError(os.Remove("__MACOSX"))
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("generate Template Theme success!")
	fmt.Println()
}

func unzipDir(zipFile, dir string) {

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		checkError(err)
	}
	defer func() {
		_ = r.Close()
	}()

	for _, f := range r.File {
		func() {
			path := dir + string(filepath.Separator) + f.Name
			checkError(os.MkdirAll(filepath.Dir(path), 0755))
			fDest, err := os.Create(path)
			checkError(err)
			defer func() {
				_ = fDest.Close()
			}()

			fSrc, err := f.Open()
			checkError(err)
			defer func() {
				_ = fSrc.Close()
			}()

			_, err = io.Copy(fDest, fSrc)
			checkError(err)
		}()
	}
}

func replaceContents(fileDir, moduleName, themeName string) {
	files, err := ioutil.ReadDir(fileDir)
	checkError(err)
	for _, file := range files {
		path := fileDir + "/" + file.Name()
		buf, err := ioutil.ReadFile(path)
		checkError(err)
		content := string(buf)

		newContent := strings.Replace(content, "github.com/GoAdminGroup/themes/adminlte", moduleName, -1)
		newContent = strings.Replace(newContent, "adminlte", themeName, -1)
		newContent = strings.Replace(newContent, "Adminlte", strings.Title(themeName), -1)

		checkError(ioutil.WriteFile(path, []byte(newContent), 0))
	}
}
