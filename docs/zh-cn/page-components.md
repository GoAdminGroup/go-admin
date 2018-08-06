## 列

```go
col1 := components.Col.GetContent("")
```

## 行

```go
col1 := components.Col.GetContent("")
col2 := components.Col.GetContent("")
row := components.Row.GetContent(col1 + col2)
```

## 盒子

```go
box := components.GetBox().SetUrl("/").SetTitle("用户总数").SetValue("1000").GetContent()
```