## 表单类型说明

配置编辑和新建表单，需要了解```FormStruct```结构

```go
type FormStruct struct {
	Field    string
	TypeName string
	Head     string
	Default  string
	Editable bool
	FormType string
	Value    string
	Options  []map[string]string
	ExcuFun  FieldValueFun
}
```

其中```Field```代表字段名，```TypeName```代表字段类型名，```Head```代表表头名，
```Value```代表编辑的值，```Options```代表类型为选择时的选项，```FormType```代表表单类型
```Default```代表编辑时是否可编辑。
```ExcuFun```代表字段值过滤函数，定义如下：

```go
type FieldValueFun func(value RowModel) interface{}

type RowModel struct {
	ID    int64
	Value string
}
```

传入的```RowModel```有两个属性，一个是ID，一个是Value。ID是该行数据记录的主键ID，Value代表该字段的值。

## 简单例子

```go
// 表单显示配置
userTable.Form.FormList = []FormStruct{
    {
        Head:     "姓名",
        Field:    "name",
        TypeName: "varchar",
        Default:  "",
        Editable: true,
        FormType: "default",
        ExcuFun: func(model RowModel) string {
            return model.Value
        },
    }, {
        Head:     "性别",
        Field:    "sex",
        TypeName: "tinyint",
        Default:  "",
        Editable: true,
        FormType: "text",
        ExcuFun: func(model RowModel) string {
            return model.Value
        },
    },
}
```