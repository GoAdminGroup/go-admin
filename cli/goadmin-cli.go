package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func main()  {
	content := `package template

var Adminlte = map[string]string{`
	files, _ := ioutil.ReadDir("./resources/adminlte/pages/")
	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./resources/adminlte/pages/" + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)

		suffix := path.Ext(f.Name())
		fmt.Println("f.Name()", f.Name(),"suffix",suffix)
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			content += `"` + onlyName + `":` + "`" + str + "`,"
		}
	}

	files, _ = ioutil.ReadDir("./resources/adminlte/pages/components/")

	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./resources/adminlte/pages/components/" + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)

		suffix := path.Ext(f.Name())
		fmt.Println("f.Name()", f.Name(),"suffix",suffix)
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			content += `"components/` + onlyName + `":` + "`" + str + "`,"
		}
	}

	files, _ = ioutil.ReadDir("./resources/adminlte/pages/components/form/")

	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./resources/adminlte/pages/components/form/" + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)

		suffix := path.Ext(f.Name())
		fmt.Println("f.Name()", f.Name(),"suffix",suffix)
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			content += `"components/form/` + onlyName + `":` + "`" + str + "`,"
		}
	}

	files, _ = ioutil.ReadDir("./resources/login/")

	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./resources/login/" + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)

		suffix := path.Ext(f.Name())
		fmt.Println("f.Name()", f.Name(),"suffix",suffix)
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			content += `"login/` + onlyName + `":` + "`" + str + "`,"
		}
	}

	content += `}`

	ioutil.WriteFile("./template/adminlte.go", []byte(content), 0644)
}
