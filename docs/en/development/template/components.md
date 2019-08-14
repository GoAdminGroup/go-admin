# Component development

New image component development.

## New types and methods of upper interface

- Create a new ```ImgAttribute``` type:

```go
type ImgAttribute interface {
	SetWidth(value string) ImgAttribute
	SetHeight(value string) ImgAttribute
	SetSrc(value string) ImgAttribute
	GetContent() template.HTML
}
```

- Add a method to the ```Template``` interface:

```go
type Template interface {
	...
	Image() types.ImgAttribute
	...
}
```

## Implementation with ```adminlte```

- Implement ```ImgAttribute```

Create a new ```image.go``` file under ```./template/adminlte/components```:

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

- Implement ```Image()``` method

Add following function to ```.template/adminlte/adminlte.go```:

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

- Add a static resource file

Create ```image.tmpl``` file under ```.template/adminlte/resource/pages/components```

- Execute in the root directory:

```shell
admincli compile
```
