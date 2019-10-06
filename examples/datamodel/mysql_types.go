package datamodel

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetAll_typesTable() table.Table {

	all_typesTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	all_typesTable.GetInfo().FieldList = []types.Field{{
		Head:     "Id",
		Field:    "id",
		TypeName: db.Int,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_1",
		Field:    "type_1",
		TypeName: db.Tinyint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_2",
		Field:    "type_2",
		TypeName: db.Smallint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_3",
		Field:    "type_3",
		TypeName: db.Mediumint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_4",
		Field:    "type_4",
		TypeName: db.Bigint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_5",
		Field:    "type_5",
		TypeName: db.Float,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_6",
		Field:    "type_6",
		TypeName: db.Double,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_7",
		Field:    "type_7",
		TypeName: db.Double,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_8",
		Field:    "type_8",
		TypeName: db.Double,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_9",
		Field:    "type_9",
		TypeName: db.Decimal,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_10",
		Field:    "type_10",
		TypeName: db.Bit,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_11",
		Field:    "type_11",
		TypeName: db.Tinyint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_12",
		Field:    "type_12",
		TypeName: db.Tinyint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_13",
		Field:    "type_13",
		TypeName: db.Decimal,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_14",
		Field:    "type_14",
		TypeName: db.Decimal,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_15",
		Field:    "type_15",
		TypeName: db.Decimal,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_16",
		Field:    "type_16",
		TypeName: db.Char,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_17",
		Field:    "type_17",
		TypeName: db.Varchar,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_18",
		Field:    "type_18",
		TypeName: db.Tinytext,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_19",
		Field:    "type_19",
		TypeName: db.Text,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_20",
		Field:    "type_20",
		TypeName: db.Mediumtext,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_21",
		Field:    "type_21",
		TypeName: db.Longtext,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_22",
		Field:    "type_22",
		TypeName: db.Tinyblob,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_23",
		Field:    "type_23",
		TypeName: db.Mediumblob,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_24",
		Field:    "type_24",
		TypeName: db.Blob,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_25",
		Field:    "type_25",
		TypeName: db.Longblob,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_26",
		Field:    "type_26",
		TypeName: db.Binary,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_27",
		Field:    "type_27",
		TypeName: db.Varbinary,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_28",
		Field:    "type_28",
		TypeName: db.Enum,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_29",
		Field:    "type_29",
		TypeName: db.Set,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_30",
		Field:    "type_30",
		TypeName: db.Date,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_31",
		Field:    "type_31",
		TypeName: db.Datetime,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_32",
		Field:    "type_32",
		TypeName: db.Timestamp,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_33",
		Field:    "type_33",
		TypeName: db.Time,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_34",
		Field:    "type_34",
		TypeName: db.Year,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_35",
		Field:    "type_35",
		TypeName: db.Geometry,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_36",
		Field:    "type_36",
		TypeName: db.Point,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_39",
		Field:    "type_39",
		TypeName: db.Multilinestring,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_41",
		Field:    "type_41",
		TypeName: db.Multipolygon,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_37",
		Field:    "type_37",
		TypeName: db.Linestring,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_38",
		Field:    "type_38",
		TypeName: db.Polygon,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_40",
		Field:    "type_40",
		TypeName: db.Multipoint,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_42",
		Field:    "type_42",
		TypeName: db.Geometrycollection,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_50",
		Field:    "type_50",
		TypeName: db.Double,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_51",
		Field:    "type_51",
		TypeName: db.Json,
		Sortable: false,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}}

	all_typesTable.GetInfo().Table = "all_types"
	all_typesTable.GetInfo().Title = "All_types"
	all_typesTable.GetInfo().Description = "All_types"

	all_typesTable.GetForm().FormList = []types.Form{{
		Head:     "Id",
		Field:    "id",
		TypeName: db.Int,
		Default:  "",
		Editable: true,
		FormType: form.Default,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_1",
		Field:    "type_1",
		TypeName: db.Tinyint,
		Default:  "",
		Editable: true,
		FormType: form.Number,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_2",
		Field:    "type_2",
		TypeName: db.Smallint,
		Default:  "",
		Editable: true,
		FormType: form.Number,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_3",
		Field:    "type_3",
		TypeName: db.Mediumint,
		Default:  "",
		Editable: true,
		FormType: form.Number,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_4",
		Field:    "type_4",
		TypeName: db.Bigint,
		Default:  "",
		Editable: true,
		FormType: form.Number,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_5",
		Field:    "type_5",
		TypeName: db.Float,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_6",
		Field:    "type_6",
		TypeName: db.Double,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_7",
		Field:    "type_7",
		TypeName: db.Double,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_8",
		Field:    "type_8",
		TypeName: db.Double,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_9",
		Field:    "type_9",
		TypeName: db.Decimal,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_10",
		Field:    "type_10",
		TypeName: db.Bit,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_11",
		Field:    "type_11",
		TypeName: db.Tinyint,
		Default:  "",
		Editable: true,
		FormType: form.Number,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_12",
		Field:    "type_12",
		TypeName: db.Tinyint,
		Default:  "",
		Editable: true,
		FormType: form.Number,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_13",
		Field:    "type_13",
		TypeName: db.Decimal,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_14",
		Field:    "type_14",
		TypeName: db.Decimal,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_15",
		Field:    "type_15",
		TypeName: db.Decimal,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_16",
		Field:    "type_16",
		TypeName: db.Char,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_17",
		Field:    "type_17",
		TypeName: db.Varchar,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_18",
		Field:    "type_18",
		TypeName: db.Tinytext,
		Default:  "",
		Editable: true,
		FormType: form.RichText,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_19",
		Field:    "type_19",
		TypeName: db.Text,
		Default:  "",
		Editable: true,
		FormType: form.RichText,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_20",
		Field:    "type_20",
		TypeName: db.Mediumtext,
		Default:  "",
		Editable: true,
		FormType: form.RichText,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_21",
		Field:    "type_21",
		TypeName: db.Longtext,
		Default:  "",
		Editable: true,
		FormType: form.RichText,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_22",
		Field:    "type_22",
		TypeName: db.Tinyblob,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_23",
		Field:    "type_23",
		TypeName: db.Mediumblob,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_24",
		Field:    "type_24",
		TypeName: db.Blob,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_25",
		Field:    "type_25",
		TypeName: db.Longblob,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_26",
		Field:    "type_26",
		TypeName: db.Binary,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_27",
		Field:    "type_27",
		TypeName: db.Varbinary,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_28",
		Field:    "type_28",
		TypeName: db.Enum,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_29",
		Field:    "type_29",
		TypeName: db.Set,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_30",
		Field:    "type_30",
		TypeName: db.Date,
		Default:  "",
		Editable: true,
		FormType: form.Datetime,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_31",
		Field:    "type_31",
		TypeName: db.Datetime,
		Default:  "",
		Editable: true,
		FormType: form.Datetime,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_32",
		Field:    "type_32",
		TypeName: db.Timestamp,
		Default:  "",
		Editable: true,
		FormType: form.Datetime,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_33",
		Field:    "type_33",
		TypeName: db.Time,
		Default:  "",
		Editable: true,
		FormType: form.Datetime,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_34",
		Field:    "type_34",
		TypeName: db.Year,
		Default:  "",
		Editable: true,
		FormType: form.Datetime,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_35",
		Field:    "type_35",
		TypeName: db.Geometry,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_36",
		Field:    "type_36",
		TypeName: db.Point,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_39",
		Field:    "type_39",
		TypeName: db.Multilinestring,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_41",
		Field:    "type_41",
		TypeName: db.Multipolygon,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_37",
		Field:    "type_37",
		TypeName: db.Linestring,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_38",
		Field:    "type_38",
		TypeName: db.Polygon,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_40",
		Field:    "type_40",
		TypeName: db.Multipoint,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_42",
		Field:    "type_42",
		TypeName: db.Geometrycollection,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_50",
		Field:    "type_50",
		TypeName: db.Double,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}, {
		Head:     "Type_51",
		Field:    "type_51",
		TypeName: db.Json,
		Default:  "",
		Editable: true,
		FormType: form.Text,
		FilterFn: func(model types.RowModel) interface{} {
			return model.Value
		},
	}}

	all_typesTable.GetForm().Table = "all_types"
	all_typesTable.GetForm().Title = "All Mysql Types"
	all_typesTable.GetForm().Description = "all mysql types"

	return all_typesTable
}
