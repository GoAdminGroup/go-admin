# 组件开发

组件开发，以图片组件为例。

## 上层接口新增类型与方法

- 新建一个```ImgAttribute```类型

```go
type ImgAttribute interface {
	SetWidth(value string) ImgAttribute
	SetHeight(value string) ImgAttribute
	SetSrc(value string) ImgAttribute
	GetContent() template.HTML
}
```

- 在```Template```接口中，增加一个方法：

```go
type Template interface {
	...
	Image() types.ImgAttribute
	...
}
```

## 具体实现，以```adminlte```为例子

- 实现```ImgAttribute```

在```./template/adminlte/components```下新建```image.go```文件，内容如下：

```go
package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type ImgAttribute struct {
	Name   string
	Witdh  string
	Height string
	Src    string
}

func (compo *ImgAttribute) SetWidth(value string) types.ImgAttribute {
	compo.Witdh = value
	return compo
}

func (compo *ImgAttribute) SetHeight(value string) types.ImgAttribute {
	compo.Height = value
	return compo
}

func (compo *ImgAttribute) SetSrc(value string) types.ImgAttribute {
	compo.Src = value
	return compo
}

func (compo *ImgAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "image")
}
```

- 实现```Image()```

在```./template/adminlte/adminlte.go```中，增加一个函数：

```go
func (*Theme) Image() types.ImgAttribute {
	return &components.ImgAttribute{
		Name:   "image",
		Witdh:  "50",
		Height: "50",
		Src:    "",
	}
}
```

到这里还是没有完成的，需要增加静态资源文件。

- 增加静态资源文件

在```./template/adminlte/resource/pages/components```增加```image.tmpl```文件

烦人，还有最后一步

- 在根目录下执行：

```shell
admincli compile
```