package login

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
)

type Installation struct {
	Name string
}

func Get() *Installation {
	return &Installation{
		Name: "installation",
	}
}

var DefaultFuncMap = template.FuncMap{
	"lang":     language.Get,
	"langHtml": language.GetFromHtml,
	"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
		if cdnUrl == "" {
			return prefixUrl + assetsUrl
		}
		return cdnUrl + assetsUrl
	},
	"isLinkUrl": func(s string) bool {
		return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
	},
	"render": func(s, old, repl template.HTML) template.HTML {
		return template.HTML(strings.ReplaceAll(string(s), string(old), string(repl)))
	},
	"renderJS": func(s template.JS, old, repl template.HTML) template.JS {
		return template.JS(strings.ReplaceAll(string(s), string(old), string(repl)))
	},
	"divide": func(a, b int) int {
		return a / b
	},
}

func (i *Installation) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("installation").
		Funcs(DefaultFuncMap).
		Parse(List["installation"])

	if err != nil {
		logger.Error("Installation GetTemplate Error: ", err)
	}

	return tmpl, "installation"
}

func (i *Installation) GetAssetList() []string               { return AssetsList }
func (i *Installation) GetAsset(name string) ([]byte, error) { return Asset(name[1:]) }
func (i *Installation) IsAPage() bool                        { return true }
func (i *Installation) GetName() string                      { return "login" }

func (i *Installation) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := i.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, i)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
