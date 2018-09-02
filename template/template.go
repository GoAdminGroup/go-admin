package template

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/adminlte"
	"html/template"
)

type Template interface {
	Form() types.FormAttribute
	Box() types.BoxAttribute
	Col() types.ColAttribute
	Image() types.ImgAttribute
	InfoBox() types.InfoBoxAttribute
	Label() types.LabelAttribute
	Row() types.RowAttribute
	Table() types.TableAttribute
	DataTable() types.DataTableAttribute
	Tree() types.TreeAttribute
	Paninator() types.PaninatorAttribute
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAsset() ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
}

func Get(theme string) Template {
	switch theme {
	case "adminlte":
		return adminlte.GetAdminlte()
	default:
		panic("wrong theme name!")
	}
}
