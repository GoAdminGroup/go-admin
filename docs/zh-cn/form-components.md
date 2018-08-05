## 多选

```go
permissionsModel, _ := mysql.Query("select `id`, `slug` from goadmin_permissions where id > ?", 0)
for _, v := range permissionsModel {
    permissions = append(permissions, map[string]string{
        "field": v["slug"].(string),
        "value": strconv.FormatInt(v["id"].(int64), 10),
    })
}

userTable.Form.FormList = []FormStruct{
    {
        Head:     "权限",
        Field:    "permission_id",
        TypeName: "varchar",
        Default:  "",
        Editable: true,
        FormType: "select",
        Options:  permissions,
        ExcuFun: func(model RowModel) interface{} {
            permissionModel, _ := mysql.Query("select permission_id from goadmin_user_permissions where user_id = ?", model.ID)
            var permissions []string
            for _, v := range permissionModel {
                permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
            }
            return permissions
        },
    }
}
```

## 多选盒子

```go
var permissions []map[string]string
permissionsModel, _ := mysql.Query("select `id`, `slug` from goadmin_permissions where id > ?", 0)
for _, v := range permissionsModel {
    permissions = append(permissions, map[string]string{
        "field": v["slug"].(string),
        "value": strconv.FormatInt(v["id"].(int64), 10),
    })
}

userTable.Form.FormList = []FormStruct{
    {
        Head:     "权限",
        Field:    "permission_id",
        TypeName: "varchar",
        Default:  "",
        Editable: true,
        FormType: "selectbox",
        Options:  permissions,
        ExcuFun: func(model RowModel) interface{} {
            perModel, _ := mysql.Query("select permission_id from goadmin_role_permissions where role_id = ?", model.ID)
            var permissions []string
            for _, v := range perModel {
                permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
            }
            return permissions
        },
    }
}
```

## 密码

```go
userTable.Form.FormList = []FormStruct{
    {
        Head:     "密码",
        Field:    "password",
        TypeName: "varchar",
        Default:  "",
        Editable: true,
        FormType: "password",
        ExcuFun: func(model RowModel) interface{} {
            return model.Value
        },
    }
}
```

## 文本框

```go
userTable.Form.FormList = []FormStruct{
    {
        Head:     "简介",
        Field:    "description",
        TypeName: "varchar",
        Default:  "",
        Editable: true,
        FormType: "textarea",
        ExcuFun: func(model RowModel) interface{} {
            return model.Value
        },
    }
}
```