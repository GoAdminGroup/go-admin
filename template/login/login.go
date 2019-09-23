package login

import (
	"github.com/chenhg5/go-admin/modules/logger"
	"html/template"
)

type Login struct {
}

func GetLoginComponent() *Login {
	return new(Login)
}

func (*Login) GetTemplate() (*template.Template, string) {
	tmpler, err := template.New("login_theme1").Parse(List["login/theme1"])

	if err != nil {
		logger.Error("Login GetTemplate Error: ", err)
	}

	return tmpler, "login_theme1"
}

func (*Login) GetAssetList() []string {
	return asserts
}

func (*Login) GetAsset(name string) ([]byte, error) {
	name = "template/login" + name
	return Asset(name)
}
