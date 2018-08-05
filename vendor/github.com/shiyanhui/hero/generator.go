package hero

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

const (
	TypeBytesBuffer = "bytes.Buffer"
	TypeIOWriter    = "io.Writer"
)

var errExpectParam = errors.New(
	"The last parameter should be *bytes.Buffer or io.Writer type",
)

func formatMap(t, bufName string) string {
	switch t {
	case String:
		return "%s"
	case Interface:
		return "fmt.Sprintf(\"%%v\", %s)"
	}

	m := map[string]string{
		Int:   "hero.FormatInt(int64(%%s), %s)",
		Uint:  "hero.FormatUint(uint64(%%s), %s)",
		Float: "hero.FormatFloat(float64(%%s), %s)",
		Bool:  "hero.FormatBool(%%s, %s)",
	}

	format, ok := m[t]
	if !ok {
		log.Fatal("Unknown type ", t)
	}
	return fmt.Sprintf(format, bufName)
}

func writeToFile(path string, buffer *bytes.Buffer) {
	err := ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func genAbsPath(path string) string {
	if !filepath.IsAbs(path) {
		var err error
		if path, err = filepath.Abs(path); err != nil {
			log.Fatal(err)
		}
	}
	return path
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// parseDefinition parses the function definition.
func parseDefinition(definition string) (*ast.FuncDecl, error) {
	src := fmt.Sprintf("package hero\n%s", definition)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	funcDecl, ok := file.Decls[0].(*ast.FuncDecl)
	if !ok {
		return nil, errors.New("Definition is not function type")
	}
	return funcDecl, nil
}

// parseParams parses parameters in the function definition.
func parseParams(funcDecl *ast.FuncDecl) (name, t string, err error) {
	params := funcDecl.Type.Params.List
	if len(params) == 0 {
		err = errors.New(
			"Definition parameters should not be empty",
		)
		return
	}

	lastParam := params[len(params)-1]

	expr := lastParam.Type
	if starExpr, ok := expr.(*ast.StarExpr); ok {
		expr = starExpr.X
	}

	selectorExpr, ok := expr.(*ast.SelectorExpr)
	if !ok {
		err = errExpectParam
		return
	}

	t = fmt.Sprintf("%s.%s", selectorExpr.X, selectorExpr.Sel)
	if t != TypeBytesBuffer && t != TypeIOWriter {
		err = fmt.Errorf(
			"'%s' expected to be '%s' or '%s'",
			t, TypeBytesBuffer, TypeIOWriter,
		)
		return
	}

	if n := len(lastParam.Names); n > 0 {
		name = lastParam.Names[n-1].Name
	}

	return
}

// parseResults parses the returned results in the function definition.
func parseResults(funcDecl *ast.FuncDecl) (types []string) {
	if results := funcDecl.Type.Results; results != nil {
		for _, field := range results.List {
			types = append(types, fmt.Sprintf("%s", field.Type))
		}
	}
	return
}

func escapeBackQuote(str string) string {
	return strings.Replace(str, "`", "`+\"`\"+`", -1)
}

// gen generates code to buffer.
func gen(n *node, buffer *bytes.Buffer, bufName string) {
	for _, child := range n.children {
		switch child.t {
		case TypeCode:
			buffer.Write(child.chunk.Bytes())
		case TypeHTML:
			buffer.WriteString(fmt.Sprintf(
				"%s.WriteString(`%s`)",
				bufName,
				escapeBackQuote(child.chunk.String()),
			))
		case TypeRawValue, TypeEscapedValue:
			var format string

			switch child.subtype {
			case Int, Uint, Float, Bool, String, Interface:
				format = formatMap(child.subtype, bufName)
				if child.subtype != String &&
					child.subtype != Interface {
					goto WriteFormat
				}
			case Bytes:
				if child.t == TypeRawValue {
					format = fmt.Sprintf("%s.Write(%%s)", bufName)
					goto WriteFormat
				}
				format = "*(*string)(unsafe.Pointer(&(%s)))"
			default:
				log.Fatal("unknown value type: " + child.subtype)
			}

			if child.t == TypeEscapedValue {
				format = fmt.Sprintf(
					"hero.EscapeHTML(%s, %s)", format, bufName,
				)
			} else {
				format = fmt.Sprintf(
					"%s.WriteString(%s)", bufName, format,
				)
			}

		WriteFormat:
			buffer.WriteString(fmt.Sprintf(format, child.chunk.String()))
		case TypeBlock, TypeInclude:
			gen(child, buffer, bufName)
		default:
			continue
		}

		buffer.WriteByte(BreakLine)
	}
}

// Generate generates Go code from source to test. pkgName represents the
// package name of the generated code.
func Generate(source, dest, pkgName string, extensions []string) {
	defer cleanGlobal()

	source, dest = genAbsPath(source), genAbsPath(dest)
	sourceDir := source

	srcStat, err := os.Stat(source)
	checkError(err)

	fmt.Println("Parsing...")
	if srcStat.IsDir() {
		parseDir(source, extensions)
	} else {
		sourceDir = filepath.Dir(source)
		source, file := filepath.Split(source)
		parseFile(source, file)
	}
	rebuild()

	destStat, err := os.Stat(dest)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dest, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	} else if !destStat.IsDir() {
		if srcStat.IsDir() {
			log.Fatal(dest + " is not a directory")
		}
		dest = filepath.Dir(dest)
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generating...")

	var wg sync.WaitGroup
	for path, n := range parsedNodes {
		wg.Add(1)

		fileName := filepath.Join(dest, fmt.Sprintf(
			"%s.go",
			strings.Join(strings.Split(
				path[len(sourceDir)+1:],
				string(filepath.Separator),
			), "_"),
		))

		go func(n *node, source, fileName string) {
			defer wg.Done()

			buffer := bytes.NewBufferString(`
				// Code generated by hero.
			`)
			buffer.WriteString(fmt.Sprintf("// source: %s", source))
			buffer.WriteString(`
				// DO NOT EDIT!
			`)
			buffer.WriteString(fmt.Sprintf("package %s\n", pkgName))
			buffer.WriteString(`
				import "html"
				import "unsafe"

				import "github.com/shiyanhui/hero"
			`)

			imports := n.childrenByType(TypeImport)
			for _, item := range imports {
				buffer.Write(item.chunk.Bytes())
			}

			definitions := n.childrenByType(TypeDefinition)
			if len(definitions) == 0 {
				writeToFile(fileName, buffer)
				return
			}

			definition := definitions[0].chunk.String()

			funcDecl, err := parseDefinition(definition)
			checkError(err)

			buffer.WriteString(definition)
			buffer.WriteString(`{
			`)

			paramName, paramType, err := parseParams(funcDecl)
			checkError(err)

			if paramType == TypeIOWriter {
				bufName := "_buffer"

				buffer.WriteString(
					fmt.Sprintf(
						"%s := hero.GetBuffer()\ndefer hero.PutBuffer(%s)\n",
						bufName, bufName,
					),
				)
				gen(n, buffer, bufName)

				results, ret := parseResults(funcDecl), ""
				if reflect.DeepEqual(results, []string{"int", "error"}) {
					ret = "return"
				}

				buffer.WriteString(
					fmt.Sprintf(
						"%s %s.Write(%s.Bytes())\n",
						ret, paramName, bufName,
					),
				)
			} else {
				gen(n, buffer, paramName)
			}

			buffer.WriteString(`
			}`)

			writeToFile(fileName, buffer)
		}(n, path, fileName)
	}
	wg.Wait()

	fmt.Println("Executing goimports...")
	execCommand("goimports -w " + dest)

	fmt.Println("Executing go vet...")
	execCommand("go tool vet -v " + dest)
}
