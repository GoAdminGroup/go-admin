## 图片组件

```go
// 列显示配置
userTable.Info.FieldList = []FieldStruct{
    {
        Head:     "头像",
        Field:    "avatar",
        TypeName: "tinyint",
        ExcuFun: func(model RowModel) string {
            return components.GetImage().SetHeight("50")
                    .SetWidth("50").SetSrc(model.Value).GetContent()
        },
    },
}
```

## 标签组件

```go
// 列显示配置
userTable.Info.FieldList = []FieldStruct{
    {
        Head:     "标签",
        Field:    "label",
        TypeName: "tinyint",
        ExcuFun: func(model RowModel) string {
            return components.Label.GetContent(model.Value)
        },
    },
}
```

