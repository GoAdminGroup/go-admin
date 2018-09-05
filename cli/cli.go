package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func main()  {
	content := `package tmpl

var List = map[string]string{`
	files, _ := ioutil.ReadDir("./template/adminlte/resource/pages/")
	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./template/adminlte/resource/pages/" + f.Name())
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

	files, _ = ioutil.ReadDir("./template/adminlte/resource/pages/components/")

	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./template/adminlte/resource/pages/components/" + f.Name())
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

	files, _ = ioutil.ReadDir("./template/adminlte/resource/pages/components/form/")

	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./template/adminlte/resource/pages/components/form/" + f.Name())
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

	files, _ = ioutil.ReadDir("./template/adminlte/resource/pages/components/table/")

	for _, f := range files {
		fmt.Println(f.Name())
		b, err := ioutil.ReadFile("./template/adminlte/resource/pages/components/table/" + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)

		suffix := path.Ext(f.Name())
		fmt.Println("f.Name()", f.Name(),"suffix",suffix)
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			content += `"components/table/` + onlyName + `":` + "`" + str + "`,"
		}
	}
	//
	//files, _ = ioutil.ReadDir("./template/login/")
	//
	//for _, f := range files {
	//	fmt.Println(f.Name())
	//	b, err := ioutil.ReadFile("./template/login/" + f.Name())
	//	if err != nil {
	//		fmt.Print(err)
	//	}
	//	str := string(b)
	//	fmt.Println(str)
	//
	//	suffix := path.Ext(f.Name())
	//	fmt.Println("f.Name()", f.Name(),"suffix",suffix)
	//	onlyName := strings.TrimSuffix(f.Name(), suffix)
	//
	//	if suffix == ".tmpl" {
	//		content += `"login/` + onlyName + `":` + "`" + str + "`,"
	//	}
	//}

	content += `}`

	ioutil.WriteFile("./template/adminlte/template/template.go", []byte(content), 0644)
}
