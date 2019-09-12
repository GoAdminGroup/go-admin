package pages

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetForm1Content() types.Panel {

	components := template2.Get(config.Get().THEME)

	aform := components.Form().
		SetContent([]types.Form{
			{
				Field:    "name",
				TypeName: "varchar",
				Head:     "Name",
				Default:  "jane",
				Editable: true,
				FormType: form.Text,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "age",
				TypeName: "int",
				Head:     "Age",
				Default:  "11",
				Editable: true,
				FormType: form.Number,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "homepage",
				TypeName: "varchar",
				Head:     "HomePage",
				Default:  "http://google.com",
				Editable: true,
				FormType: form.Url,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "email",
				TypeName: "varchar",
				Head:     "Email",
				Default:  "xxxx@xxx.com",
				Editable: true,
				FormType: form.Email,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "birthday",
				TypeName: "varchar",
				Head:     "Birthday",
				Default:  "2010-09-05",
				Editable: true,
				FormType: form.Datetime,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "gender",
				TypeName: "tinyint",
				Head:     "Gender",
				Default:  "boy",
				Editable: true,
				FormType: form.Radio,
				Value:    "",
				Options: []map[string]string{
					{
						"field":    "gender",
						"label":    "male",
						"value":    "0",
						"selected": "true",
					}, {
						"field":    "gender",
						"label":    "female",
						"value":    "1",
						"selected": "false",
					},
				},
			}, {
				Field:    "password",
				TypeName: "varchar",
				Head:     "Password",
				Default:  "",
				Editable: true,
				FormType: form.Password,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "ip",
				TypeName: "varchar",
				Head:     "Ip",
				Default:  "",
				Editable: true,
				FormType: form.Ip,
				Value:    "",
				Options:  []map[string]string{},
			}, {
				Field:    "currency",
				TypeName: "int",
				Head:     "Currency",
				Default:  "",
				Editable: true,
				FormType: form.Currency,
				Value:    "",
				Options:  []map[string]string{},
			},
		}).
		SetPrefix(config.Get().PrefixFixSlash()).
		SetUrl("/").
		SetTitle("Form1").
		SetToken(auth.TokenHelper.AddToken()).
		SetInfoUrl("/admin").
		GetContent()

	return types.Panel{
		Content:     aform,
		Title:       "Form1",
		Description: "this is a form example",
	}
}
