## 表格类型说明

配置展示表格，需要先知道展示表格的数据类型```FieldStruct```

```go
type FieldStruct struct {
	ExcuFun  FieldValueFun
	Field    string
	TypeName string
	Head     string
}
```

其中```Field```代表字段名，```TypeName```代表字段类型名，```Head```代表表头名。
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
// 列显示配置
userTable.Info.FieldList = []FieldStruct{
    {
        Head:     "姓名",
        Field:    "name",
        TypeName: "varchar",
        ExcuFun: func(model RowModel) string {
            return model.Value
        },
    },
    {
        Head:     "性别",
        Field:    "sex",
        TypeName: "tinyint",
        ExcuFun: func(model RowModel) string {
            if model.Value == "1" {
                return "男"
            }
            if model.Value == "2" {
                return "女"
            }
            return "未知"
        },
    },
}
```