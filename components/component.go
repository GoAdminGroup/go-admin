package components

type Component interface {
	GetContent(interface{}) string
}

type Attribute struct {
	Name    string
	Content string
}

var Default = &Attribute{
	"Default",
	"",
}

func (compo *Attribute) GetContent(value interface{}) string {
	return (*compo).Content + value.(string)
}
