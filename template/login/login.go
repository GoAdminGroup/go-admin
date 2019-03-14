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

func (*Login) GetAsset(string) ([]byte, error) {
	panic("implement me")
}
