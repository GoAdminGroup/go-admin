package datamodel

import "html/template"

// custom your own login page

type LoginPage struct {
}

func (*LoginPage) GetTemplate() (*template.Template, string) {
	panic("implement me")
}

func (*LoginPage) GetAssetList() []string {
	panic("implement me")
}

func (*LoginPage) GetAsset(string) ([]byte, error) {
	panic("implement me")
}
