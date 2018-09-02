package template

import (
	"github.com/chenhg5/go-admin/template/adminlte/components"
	"github.com/chenhg5/go-admin/template/types"
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
}

func Get(theme string) Template {
	switch theme {
	case "adminlte":
		return components.GetAdminlte()
	default:
		panic("wrong theme name!")
	}
}