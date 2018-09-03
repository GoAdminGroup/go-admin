package login

import (
	"html/template"
	"fmt"
)

type Login struct {
}

func GetLoginComponent() *Login {
	return new(Login)
}

func (*Login) GetTemplate() (*template.Template, string) {
	tmpler, err := template.New("content").Parse(List["login/theme1"])

	if err != nil {
		fmt.Println(err)
	}

	return tmpler, "content"
}

func (*Login) GetAssetList() []string {
	return asserts
}

func (*Login) GetAsset(string) ([]byte, error) {
	panic("implement me")
}
