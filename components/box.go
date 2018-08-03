package components

type BoxAttribute struct {
	Name  string
	Title string
	Value string
	Url   string
}

func GetBox() *BoxAttribute {
	return &BoxAttribute{
		"box",
		"2",
		"2",
		"2",
	}
}

func (compo *BoxAttribute) SetTitle(value string) *BoxAttribute {
	(*compo).Title = value
	return compo
}

func (compo *BoxAttribute) SetValue(value string) *BoxAttribute {
	(*compo).Value = value
	return compo
}

func (compo *BoxAttribute) SetUrl(value string) *BoxAttribute {
	(*compo).Url = value
	return compo
}

func (compo *BoxAttribute) GetContent() string {
		return `<div class="small-box bg-aqua">
<div class="inner">
<h3>` + (*compo).Value + `</h3>

<p>` + (*compo).Title + `</p>
</div>
<div class="icon">
<i class="fa fa-users"></i>
</div>
<a href="` + (*compo).Url + `" class="small-box-footer">
More&nbsp;
<i class="fa fa-arrow-circle-right"></i>
</a>
</div>`
}
