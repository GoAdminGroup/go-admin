package chartjs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	template2 "github.com/GoAdminGroup/go-admin/template"
)

type BarChart struct {
	*Chart

	JsContent BarJsContent
}

type BarJsContent struct {
	JsContent

	Data BarAttributes `json:"data"`
}

type BarAttributes struct {
	Attributes

	DataSets BarDataSets `json:"datasets"`
}

type BarDataSets []*BarDataSet

func (l BarDataSets) Add(ds *BarDataSet) BarDataSets {
	return append(l, ds)
}

type BarDataSet struct {
	Label           string    `json:"label"`
	Data            []float64 `json:"data"`
	Type            string    `json:"type,omitempty"`
	BackgroundColor Color     `json:"backgroundColor,omitempty"`
	BorderCapStyle  string    `json:"borderCapStyle,omitempty"`
	BorderColor     Color     `json:"borderColor,omitempty"`

	BorderSkipped string  `json:"borderSkipped,omitempty"`
	BorderWidth   float64 `json:"borderWidth,omitempty"`

	HoverBackgroundColor Color   `json:"hoverBackgroundColor,omitempty"`
	HoverBorderColor     Color   `json:"hoverBorderColor,omitempty"`
	HoverBorderWidth     float64 `json:"hoverBorderWidth,omitempty"`

	Order   float64 `json:"order,omitempty"`
	XAxisID string  `json:"xAxisID,omitempty"`
	YAxisID string  `json:"yAxisID,omitempty"`
}

func (l *BarDataSet) SetLabel(label string) *BarDataSet {
	l.Label = label
	return l
}

func (l *BarDataSet) SetData(data []float64) *BarDataSet {
	l.Data = data
	return l
}

func (l *BarDataSet) SetType(t string) *BarDataSet {
	l.Type = t
	return l
}

func (l *BarDataSet) SetBackgroundColor(backgroundColor Color) *BarDataSet {
	l.BackgroundColor = backgroundColor
	return l
}

func (l *BarDataSet) SetBorderCapStyle(borderCapStyle string) *BarDataSet {
	l.BorderCapStyle = borderCapStyle
	return l
}

func (l *BarDataSet) SetBorderColor(borderColor Color) *BarDataSet {
	l.BorderColor = borderColor
	return l
}

func (l *BarDataSet) SetBorderWidth(borderWidth float64) *BarDataSet {
	l.BorderWidth = borderWidth
	return l
}

func (l *BarDataSet) SetBorderSkipped(skip string) *BarDataSet {
	l.BorderSkipped = skip
	return l
}

func (l *BarDataSet) SetHoverBackgroundColor(hoverBackgroundColor Color) *BarDataSet {
	l.HoverBackgroundColor = hoverBackgroundColor
	return l
}

func (l *BarDataSet) SetHoverBorderColor(hoverBorderColor Color) *BarDataSet {
	l.HoverBorderColor = hoverBorderColor
	return l
}

func (l *BarDataSet) SetHoverBorderWidth(hoverBorderWidth float64) *BarDataSet {
	l.HoverBorderWidth = hoverBorderWidth
	return l
}

func (l *BarDataSet) SetOrder(order float64) *BarDataSet {
	l.Order = order
	return l
}

func (l *BarDataSet) SetXAxisID(xAxisID string) *BarDataSet {
	l.XAxisID = xAxisID
	return l
}

func (l *BarDataSet) SetYAxisID(yAxisID string) *BarDataSet {
	l.YAxisID = yAxisID
	return l
}

func Bar() *BarChart {
	return &BarChart{
		Chart: &Chart{
			BaseComponent: &template2.BaseComponent{
				Name:     "chartjs",
				HTMLData: List["chartjs"],
			},
			dataSetIndex: -1,
		},
		JsContent: BarJsContent{
			JsContent: JsContent{
				Type: "bar",
			},
			Data: BarAttributes{
				Attributes: Attributes{
					Labels: make([]string, 0),
				},
				DataSets: make(BarDataSets, 0),
			},
		},
	}
}

func (l *BarChart) SetID(s string) *BarChart {
	l.ID = s
	return l
}

func (l *BarChart) SetTitle(s template.HTML) *BarChart {
	l.Title = s
	return l
}

func (l *BarChart) SetHeight(s int) *BarChart {
	l.Height = s
	return l
}

func (l *BarChart) SetLabels(s []string) *BarChart {
	l.JsContent.Data.Labels = s
	return l
}

func (l *BarChart) AddDataSet(s string) *BarChart {
	l.dataSetIndex++
	l.JsContent.Data.DataSets = l.JsContent.Data.DataSets.Add(&BarDataSet{
		Type:  "bar",
		Label: s,
	})
	return l
}

func (l *BarChart) DSLabel(s string) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetLabel(s)
	return l
}

func (l *BarChart) DSData(data []float64) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetData(data)
	return l
}

func (l *BarChart) DSType(t string) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetType(t)
	return l
}

func (l *BarChart) DSBackgroundColor(backgroundColor Color) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBackgroundColor(backgroundColor)
	return l
}

func (l *BarChart) DSBorderCapStyle(borderCapStyle string) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderCapStyle(borderCapStyle)
	return l
}

func (l *BarChart) DSBorderSkipped(skip string) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderSkipped(skip)
	return l
}

func (l *BarChart) DSBorderColor(borderColor Color) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderColor(borderColor)
	return l
}

func (l *BarChart) DSBorderWidth(borderWidth float64) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderWidth(borderWidth)
	return l
}

func (l *BarChart) DSHoverBackgroundColor(hoverBackgroundColor Color) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBackgroundColor(hoverBackgroundColor)
	return l
}

func (l *BarChart) DSHoverBorderColor(hoverBorderColor Color) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderColor(hoverBorderColor)
	return l
}

func (l *BarChart) DSHoverBorderWidth(hoverBorderWidth float64) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderWidth(hoverBorderWidth)
	return l
}

func (l *BarChart) DSOrder(order float64) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetOrder(order)
	return l
}

func (l *BarChart) DSXAxisID(xAxisID string) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetXAxisID(xAxisID)
	return l
}

func (l *BarChart) DSYAxisID(yAxisID string) *BarChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetYAxisID(yAxisID)
	return l
}

func (l *BarChart) GetContent() template.HTML {
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
