package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/utils"
)

func getThemeTemplate(moduleName, themeName string) {

	downloadTo("http://file.go-admin.cn/go_admin/template/template.zip", "tmp.zip")

	checkError(unzipDir("tmp.zip", "."))

	checkError(os.Rename("./QiAtztVk83CwCh", "./"+themeName))

	replaceContents("./"+themeName, moduleName, themeName)

	checkError(os.Rename("./"+themeName+"/template.go", "./"+themeName+"/"+themeName+".go"))

	fmt.Println()
	fmt.Println("generate theme template success!!🍺🍺")
	fmt.Println()
}

func downloadTo(url, output string) {
	defer func() {
		_ = os.Remove(output)
	}()

	req, err := http.NewRequest("GET", url, nil)

	checkError(err)

	res, err := http.DefaultClient.Do(req)

	checkError(err)

	defer func() {
		_ = res.Body.Close()
	}()

	file, err := os.Create(output)

	checkError(err)

	_, err = io.Copy(file, res.Body)

	checkError(err)
}

func unzipDir(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	checkError(os.MkdirAll(dest, 0750))

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			checkError(os.MkdirAll(path, f.Mode()))
		} else {
			checkError(os.MkdirAll(filepath.Dir(path), f.Mode()))
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func replaceContents(fileDir, moduleName, themeName string) {
	files, err := ioutil.ReadDir(fileDir)
	checkError(err)
	for _, file := range files {
		path := fileDir + "/" + file.Name()
		if !file.IsDir() {
			buf, err := ioutil.ReadFile(path)
			checkError(err)
			content := string(buf)

			newContent := utils.ReplaceAll(content, "github.com/GoAdminGroup/themes/adminlte", moduleName,
				"adminlte", themeName, "Adminlte", strings.Title(themeName))

			checkError(ioutil.WriteFile(path, []byte(newContent), 0))
		}
	}
}
