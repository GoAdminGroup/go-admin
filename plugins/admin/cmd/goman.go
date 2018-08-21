package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ")
		fmt.Println("gonne [command]")
		fmt.Printf("\n")
		fmt.Println("Available Commands: ")
		fmt.Println("  create      create a config file")
		return
	}

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	createFlag := createCommand.String("table", "", "table config file name")

	switch os.Args[1] {
	case "create":
		createCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	if createCommand.Parsed() {
		if *createFlag == "" {
			fmt.Println("please supply a config file name")
			fmt.Printf("\n")
			fmt.Println("Usage: ")
			fmt.Println("gonne create [flag]")
			fmt.Printf("\n")
			fmt.Println("Available Flags: ")
			fmt.Println("  --table      config file")
			fmt.Printf("\n")
			fmt.Println("Example: ")
			fmt.Println("  goman create --table=user")
			return
		}

		fmt.Printf("your input is %q.\n", *createFlag)

		d1 := []byte(`package models

func Get` + strFirstToUpper(*createFlag) + `Table() (table GlobalTable) {

	table.Info.FieldList = []FieldStruct{
	}

	table.Info.Table = ""
	table.Info.Title = ""
	table.Info.Description = ""

	table.Form.FormList = []FormStruct{
	}

	table.Form.Table = ""
	table.Form.Title = ""
	table.Form.Description = ""

	return
}
`)

		f, _ := os.Create("./models/" + *createFlag + ".go")

		defer f.Close()

		f.Write(d1)
	}
}

func strFirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			vv[i] -= 32
			upperStr += string(vv[i])
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
