package login

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"html/template"
)

type Login struct {
}

func GetLoginComponent() *Login {
	return new(Login)
}

func (*Login) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("login_theme1").
		Funcs(template.FuncMap{
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
		}).
		Parse(List["login/theme1"])

	if err != nil {
		logger.Error("Login GetTemplate Error: ", err)
	}

	return tmpl, "login_theme1"
}

func (*Login) GetAssetList() []string {
	return AssetsList
}

func (*Login) GetAsset(name string) ([]byte, error) {
	return Asset(name[1:])
}
