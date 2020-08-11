package chartjs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	template2 "github.com/GoAdminGroup/go-admin/template"
)

type LineChart struct {
	*Chart

	JsContent LineJsContent
}

type LineJsContent struct {
	JsContent

	Data LineAttributes `json:"data"`
}

type LineAttributes struct {
	Attributes

	DataSets LineDataSets `json:"datasets"`
}

type LineDataSets []*LineDataSet

func (l LineDataSets) Add(ds *LineDataSet) LineDataSets {
	return append(l, ds)
}

type LineDataSet struct {
	Label                     string    `json:"label"`
	Data                      []float64 `json:"data"`
	Type                      string    `json:"type,omitempty"`
	BackgroundColor           Color     `json:"backgroundColor,omitempty"`
	BorderCapStyle            string    `json:"borderCapStyle,omitempty"`
	BorderColor               Color     `json:"borderColor,omitempty"`
	BorderDash                []int     `json:"borderDash,omitempty"`
	BorderDashOffset          float64   `json:"borderDashOffset,omitempty"`
	BorderJoinStyle           string    `json:"borderJoinStyle,omitempty"`
	BorderWidth               float64   `json:"borderWidth,omitempty"`
	CubicInterpolationMode    string    `json:"cubicInterpolationMode,omitempty"`
	Fill                      bool      `json:"fill"`
	HoverBackgroundColor      Color     `json:"hoverBackgroundColor,omitempty"`
	HoverBorderCapStyle       string    `json:"hoverBorderCapStyle,omitempty"`
	HoverBorderColor          Color     `json:"hoverBorderColor,omitempty"`
	HoverBorderDash           float64   `json:"hoverBorderDash,omitempty"`
	HoverBorderDashOffset     float64   `json:"hoverBorderDashOffset,omitempty"`
	HoverBorderJoinStyle      string    `json:"hoverBorderJoinStyle,omitempty"`
	HoverBorderWidth          float64   `json:"hoverBorderWidth,omitempty"`
	LineTension               float64   `json:"lineTension,omitempty"`
	Order                     float64   `json:"order,omitempty"`
	PointBackgroundColor      Color     `json:"pointBackgroundColor,omitempty"`
	PointBorderColor          Color     `json:"pointBorderColor,omitempty"`
	PointBorderWidth          float64   `json:"pointBorderWidth,omitempty"`
	PointHitRadius            float64   `json:"pointHitRadius,omitempty"`
	PointHoverBackgroundColor Color     `json:"pointHoverBackgroundColor,omitempty"`
	PointHoverBorderColor     Color     `json:"pointHoverBorderColor,omitempty"`
	PointHoverBorderWidth     float64   `json:"pointHoverBorderWidth,omitempty"`
	PointHoverRadius          float64   `json:"pointHoverRadius,omitempty"`
	PointRadius               float64   `json:"pointRadius,omitempty"`
	PointRotation             float64   `json:"pointRotation,omitempty"`
	PointStyle                string    `json:"pointStyle,omitempty"`
	ShowLine                  bool      `json:"showLine,omitempty"`
	SpanGaps                  bool      `json:"spanGaps,omitempty"`
	SteppedLine               bool      `json:"steppedLine,omitempty"`
	XAxisID                   string    `json:"xAxisID,omitempty"`
	YAxisID                   string    `json:"yAxisID,omitempty"`
}

func (l *LineDataSet) SetLabel(label string) *LineDataSet {
	l.Label = label
	return l
}

func (l *LineDataSet) SetData(data []float64) *LineDataSet {
	l.Data = data
	return l
}

func (l *LineDataSet) SetType(t string) *LineDataSet {
	l.Type = t
	return l
}

func (l *LineDataSet) SetBackgroundColor(backgroundColor Color) *LineDataSet {
	l.BackgroundColor = backgroundColor
	return l
}

func (l *LineDataSet) SetBorderCapStyle(borderCapStyle string) *LineDataSet {
	l.BorderCapStyle = borderCapStyle
	return l
}

func (l *LineDataSet) SetBorderColor(borderColor Color) *LineDataSet {
	l.BorderColor = borderColor
	return l
}

func (l *LineDataSet) SetBorderDash(borderDash []int) *LineDataSet {
	l.BorderDash = borderDash
	return l
}

func (l *LineDataSet) SetBorderDashOffset(borderDashOffset float64) *LineDataSet {
	l.BorderDashOffset = borderDashOffset
	return l
}

func (l *LineDataSet) SetBorderJoinStyle(borderJoinStyle string) *LineDataSet {
	l.BorderJoinStyle = borderJoinStyle
	return l
}

func (l *LineDataSet) SetBorderWidth(borderWidth float64) *LineDataSet {
	l.BorderWidth = borderWidth
	return l
}

func (l *LineDataSet) SetCubicInterpolationMode(cubicInterpolationMode string) *LineDataSet {
	l.CubicInterpolationMode = cubicInterpolationMode
	return l
}

func (l *LineDataSet) SetFill(fill bool) *LineDataSet {
	l.Fill = fill
	return l
}

func (l *LineDataSet) SetHoverBackgroundColor(hoverBackgroundColor Color) *LineDataSet {
	l.HoverBackgroundColor = hoverBackgroundColor
	return l
}

func (l *LineDataSet) SetHoverBorderCapStyle(hoverBorderCapStyle string) *LineDataSet {
	l.HoverBorderCapStyle = hoverBorderCapStyle
	return l
}

func (l *LineDataSet) SetHoverBorderColor(hoverBorderColor Color) *LineDataSet {
	l.HoverBorderColor = hoverBorderColor
	return l
}

func (l *LineDataSet) SetHoverBorderDash(hoverBorderDash float64) *LineDataSet {
	l.HoverBorderDash = hoverBorderDash
	return l
}

func (l *LineDataSet) SetHoverBorderDashOffset(hoverBorderDashOffset float64) *LineDataSet {
	l.HoverBorderDashOffset = hoverBorderDashOffset
	return l
}

func (l *LineDataSet) SetHoverBorderJoinStyle(hoverBorderJoinStyle string) *LineDataSet {
	l.HoverBorderJoinStyle = hoverBorderJoinStyle
	return l
}

func (l *LineDataSet) SetHoverBorderWidth(hoverBorderWidth float64) *LineDataSet {
	l.HoverBorderWidth = hoverBorderWidth
	return l
}

func (l *LineDataSet) SetLineTension(lineTension float64) *LineDataSet {
	l.LineTension = lineTension
	return l
}

func (l *LineDataSet) SetOrder(order float64) *LineDataSet {
	l.Order = order
	return l
}

func (l *LineDataSet) SetPointBackgroundColor(pointBackgroundColor Color) *LineDataSet {
	l.PointBackgroundColor = pointBackgroundColor
	return l
}

func (l *LineDataSet) SetPointBorderColor(pointBorderColor Color) *LineDataSet {
	l.PointBorderColor = pointBorderColor
	return l
}

func (l *LineDataSet) SetPointBorderWidth(pointBorderWidth float64) *LineDataSet {
	l.PointBorderWidth = pointBorderWidth
	return l
}

func (l *LineDataSet) SetPointHitRadius(pointHitRadius float64) *LineDataSet {
	l.PointHitRadius = pointHitRadius
	return l
}

func (l *LineDataSet) SetPointHoverBackgroundColor(pointHoverBackgroundColor Color) *LineDataSet {
	l.PointHoverBackgroundColor = pointHoverBackgroundColor
	return l
}

func (l *LineDataSet) SetPointHoverBorderColor(pointHoverBorderColor Color) *LineDataSet {
	l.PointHoverBorderColor = pointHoverBorderColor
	return l
}

func (l *LineDataSet) SetPointHoverBorderWidth(pointHoverBorderWidth float64) *LineDataSet {
	l.PointHoverBorderWidth = pointHoverBorderWidth
	return l
}

func (l *LineDataSet) SetPointHoverRadius(pointHoverRadius float64) *LineDataSet {
	l.PointHoverRadius = pointHoverRadius
	return l
}

func (l *LineDataSet) SetPointRadius(pointRadius float64) *LineDataSet {
	l.PointRadius = pointRadius
	return l
}

func (l *LineDataSet) SetPointRotation(pointRotation float64) *LineDataSet {
	l.PointRotation = pointRotation
	return l
}

func (l *LineDataSet) SetPointStyle(pointStyle string) *LineDataSet {
	l.PointStyle = pointStyle
	return l
}

func (l *LineDataSet) SetShowLine(showLine bool) *LineDataSet {
	l.ShowLine = showLine
	return l
}

func (l *LineDataSet) SetSpanGaps(spanGaps bool) *LineDataSet {
	l.SpanGaps = spanGaps
	return l
}

func (l *LineDataSet) SetSteppedLine(steppedLine bool) *LineDataSet {
	l.SteppedLine = steppedLine
	return l
}

func (l *LineDataSet) SetXAxisID(xAxisID string) *LineDataSet {
	l.XAxisID = xAxisID
	return l
}

func (l *LineDataSet) SetYAxisID(yAxisID string) *LineDataSet {
	l.YAxisID = yAxisID
	return l
}

func Line() *LineChart {
	return &LineChart{
		Chart: &Chart{
			BaseComponent: &template2.BaseComponent{
				Name:     "chartjs",
				HTMLData: List["chartjs"],
			},
			dataSetIndex: -1,
		},
		JsContent: LineJsContent{
			JsContent: JsContent{
				Type: "line",
			},
			Data: LineAttributes{
				Attributes: Attributes{
					Labels: make([]string, 0),
				},
				DataSets: make(LineDataSets, 0),
			},
		},
	}
}

func (l *LineChart) SetID(s string) *LineChart {
	l.ID = s
	return l
}

func (l *LineChart) SetTitle(s template.HTML) *LineChart {
	l.Title = s
	return l
}

func (l *LineChart) SetHeight(s int) *LineChart {
	l.Height = s
	return l
}

func (l *LineChart) SetLabels(s []string) *LineChart {
	l.JsContent.Data.Labels = s
	return l
}

func (l *LineChart) AddDataSet(s string) *LineChart {
	l.dataSetIndex++
	l.JsContent.Data.DataSets = l.JsContent.Data.DataSets.Add(&LineDataSet{
		Type:  "line",
		Label: s,
	})
	return l
}

func (l *LineChart) DSLabel(s string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetLabel(s)
	return l
}

func (l *LineChart) DSData(data []float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetData(data)
	return l
}

func (l *LineChart) DSType(t string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetType(t)
	return l
}

func (l *LineChart) DSBackgroundColor(backgroundColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBackgroundColor(backgroundColor)
	return l
}

func (l *LineChart) DSBorderCapStyle(borderCapStyle string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderCapStyle(borderCapStyle)
	return l
}

func (l *LineChart) DSBorderColor(borderColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderColor(borderColor)
	return l
}

func (l *LineChart) DSBorderDash(borderDash []int) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderDash(borderDash)
	return l
}

func (l *LineChart) DSBorderDashOffset(borderDashOffset float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderDashOffset(borderDashOffset)
	return l
}

func (l *LineChart) DSBorderJoinStyle(borderJoinStyle string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderJoinStyle(borderJoinStyle)
	return l
}

func (l *LineChart) DSBorderWidth(borderWidth float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetBorderWidth(borderWidth)
	return l
}

func (l *LineChart) DSCubicInterpolationMode(cubicInterpolationMode string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetCubicInterpolationMode(cubicInterpolationMode)
	return l
}

func (l *LineChart) DSFill(fill bool) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetFill(fill)
	return l
}

func (l *LineChart) DSHoverBackgroundColor(hoverBackgroundColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBackgroundColor(hoverBackgroundColor)
	return l
}

func (l *LineChart) DSHoverBorderCapStyle(hoverBorderCapStyle string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderCapStyle(hoverBorderCapStyle)
	return l
}

func (l *LineChart) DSHoverBorderColor(hoverBorderColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderColor(hoverBorderColor)
	return l
}

func (l *LineChart) DSHoverBorderDash(hoverBorderDash float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderDash(hoverBorderDash)
	return l
}

func (l *LineChart) DSHoverBorderDashOffset(hoverBorderDashOffset float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderDashOffset(hoverBorderDashOffset)
	return l
}

func (l *LineChart) DSHoverBorderJoinStyle(hoverBorderJoinStyle string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderJoinStyle(hoverBorderJoinStyle)
	return l
}

func (l *LineChart) DSHoverBorderWidth(hoverBorderWidth float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetHoverBorderWidth(hoverBorderWidth)
	return l
}

func (l *LineChart) DSLineTension(lineTension float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetLineTension(lineTension)
	return l
}

func (l *LineChart) DSOrder(order float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetOrder(order)
	return l
}

func (l *LineChart) DSPointBackgroundColor(pointBackgroundColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointBackgroundColor(pointBackgroundColor)
	return l
}

func (l *LineChart) DSPointBorderColor(pointBorderColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointBorderColor(pointBorderColor)
	return l
}

func (l *LineChart) DSPointBorderWidth(pointBorderWidth float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointBorderWidth(pointBorderWidth)
	return l
}

func (l *LineChart) DSPointHitRadius(pointHitRadius float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointHitRadius(pointHitRadius)
	return l
}

func (l *LineChart) DSPointHoverBackgroundColor(pointHoverBackgroundColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointHoverBackgroundColor(pointHoverBackgroundColor)
	return l
}

func (l *LineChart) DSPointHoverBorderColor(pointHoverBorderColor Color) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointHoverBorderColor(pointHoverBorderColor)
	return l
}

func (l *LineChart) DSPointHoverBorderWidth(pointHoverBorderWidth float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointHoverBorderWidth(pointHoverBorderWidth)
	return l
}

func (l *LineChart) DSPointHoverRadius(pointHoverRadius float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointHoverRadius(pointHoverRadius)
	return l
}

func (l *LineChart) DSPointRadius(pointRadius float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointRadius(pointRadius)
	return l
}

func (l *LineChart) DSPointRotation(pointRotation float64) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointRotation(pointRotation)
	return l
}

func (l *LineChart) DSPointStyle(pointStyle string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetPointStyle(pointStyle)
	return l
}

func (l *LineChart) DSShowLine(showLine bool) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetShowLine(showLine)
	return l
}

func (l *LineChart) DSSpanGaps(spanGaps bool) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetSpanGaps(spanGaps)
	return l
}

func (l *LineChart) DSSteppedLine(steppedLine bool) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetSteppedLine(steppedLine)
	return l
}

func (l *LineChart) DSXAxisID(xAxisID string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetXAxisID(xAxisID)
	return l
}

func (l *LineChart) DSYAxisID(yAxisID string) *LineChart {
	l.JsContent.Data.DataSets[l.dataSetIndex].SetYAxisID(yAxisID)
	return l
}

func (l *LineChart) GetContent() template.HTML {
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
