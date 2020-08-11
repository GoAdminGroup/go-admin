package chartjs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	template2 "github.com/GoAdminGroup/go-admin/template"
)

type PieChart struct {
	*Chart

	JsContent PieJsContent
}

type PieJsContent struct {
	JsContent

	Data PieAttributes `json:"data"`
}

type PieAttributes struct {
	Attributes

	DataSets PieDataSets `json:"datasets"`
}

type PieDataSets []*PieDataSet

func (l PieDataSets) Add(ds *PieDataSet) PieDataSets {
	return append(l, ds)
}

type PieDataSet struct {
	Label           string    `json:"label"`
	Data            []float64 `json:"data"`
	Type            string    `json:"type,omitempty"`
	BackgroundColor []Color   `json:"backgroundColor,omitempty"`
	BorderColor     Color     `json:"borderColor,omitempty"`

	BorderWidth float64 `json:"borderWidth,omitempty"`
	BorderAlign string  `json:"borderAlign,omitempty"`

	HoverBackgroundColor Color   `json:"hoverBackgroundColor,omitempty"`
	HoverBorderColor     Color   `json:"hoverBorderColor,omitempty"`
	HoverBorderWidth     float64 `json:"hoverBorderWidth,omitempty"`

	Weight int `json:"weight,omitempty"`
}

func (l *PieDataSet) SetLabel(label string) *PieDataSet {
	l.Label = label
	return l
}

func (l *PieDataSet) SetData(data []float64) *PieDataSet {
	l.Data = data
	return l
}

func (l *PieDataSet) SetType(t string) *PieDataSet {
	l.Type = t
	return l
}

func (l *PieDataSet) SetBackgroundColor(backgroundColor []Color) *PieDataSet {
	l.BackgroundColor = backgroundColor
	return l
}

func (l *PieDataSet) SetBorderAlign(align string) *PieDataSet {
	l.BorderAlign = align
	return l
}

func (l *PieDataSet) SetBorderColor(borderColor Color) *PieDataSet {
	l.BorderColor = borderColor
	return l
}

func (l *PieDataSet) SetBorderWidth(borderWidth float64) *PieDataSet {
	l.BorderWidth = borderWidth
	return l
}

func (l *PieDataSet) SetWeight(weight int) *PieDataSet {
	l.Weight = weight
	return l
}

func (l *PieDataSet) SetHoverBackgroundColor(hoverBackgroundColor Color) *PieDataSet {
	l.HoverBackgroundColor = hoverBackgroundColor
	return l
}

func (l *PieDataSet) SetHoverBorderColor(hoverBorderColor Color) *PieDataSet {
	l.HoverBorderColor = hoverBorderColor
	return l
}

func (l *PieDataSet) SetHoverBorderWidth(hoverBorderWidth float64) *PieDataSet {
	l.HoverBorderWidth = hoverBorderWidth
	return l
}

func Pie() *PieChart {
	return &PieChart{
		Chart: &Chart{
			BaseComponent: &template2.BaseComponent{
				Name:     "chartjs",
				HTMLData: List["chartjs"],
			},
			dataSetIndex: -1,
		},
		JsContent: PieJsContent{
			JsContent: JsContent{
				Type: "pie",
			},
			Data: PieAttributes{
				Attributes: Attributes{
					Labels: make([]string, 0),
				},
				DataSets: make(PieDataSets, 0),
			},
		},
	}
}

func (l *PieChart) SetID(s string) *PieChart {
	l.ID = s
	return l
}

func (l *PieChart) SetTitle(s template.HTML) *PieChart {
	l.Title = s
	return l
}

func (l *PieChart) SetHeight(s int) *PieChart {
	l.Height = s
	return l
}

func (l *PieChart) SetLabels(s []string) *PieChart {
	l.JsContent.Data.Labels = s
	return l
}

func (l *PieChart) AddDataSet(s string) *PieChart {
	l.dataSetIndex++
	l.JsContent.Data.DataSets = l.JsContent.Data.DataSets.Add(&PieDataSet{
		Type:  "pie",
		Label: s,
	})
	return l
}

func (l *PieChart) DSLabel(s string) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetLabel(s)
	return l
}

func (l *PieChart) DSData(data []float64) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetData(data)
	return l
}

func (l *PieChart) DSType(t string) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetType(t)
	return l
}

func (l *PieChart) DSBackgroundColor(backgroundColor []Color) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBackgroundColor(backgroundColor)
	return l
}

func (l *PieChart) DSBorderColor(borderColor Color) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderColor(borderColor)
	return l
}

func (l *PieChart) DSBorderWidth(borderWidth float64) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderWidth(borderWidth)
	return l
}

func (l *PieChart) DSWeight(weight int) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetWeight(weight)
	return l
}

func (l *PieChart) DSHoverBackgroundColor(hoverBackgroundColor Color) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBackgroundColor(hoverBackgroundColor)
	return l
}

func (l *PieChart) DSHoverBorderColor(hoverBorderColor Color) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderColor(hoverBorderColor)
	return l
}

func (l *PieChart) DSHoverBorderWidth(hoverBorderWidth float64) *PieChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderWidth(hoverBorderWidth)
	return l
}

func (l *PieChart) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := l.GetTemplate()

	if l.JsContentOptions != nil {
		l.JsContent.Options = l.JsContentOptions
	}

	jsonByte, _ := json.Marshal(l.JsContent)
	l.Js = template.JS(string(jsonByte))

	err := tmpl.ExecuteTemplate(buffer, defineName, l)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
