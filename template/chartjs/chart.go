package chartjs

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"html/template"
)

type Chart struct {
	ID     string
	Title  template.HTML
	Js     template.JS
	Height int

	dataSetIndex int
}

func (c Chart) SetID(id string) Chart {
	c.ID = id
	return c
}

func (c Chart) SetTitle(title template.HTML) Chart {
	c.Title = title
	return c
}

func (c Chart) SetHeight(height int) Chart {
	c.Height = height
	return c
}

type JsContent struct {
	Type    string  `json:"type"`
	Options Options `json:"options"`
}

type Options struct {
	Animation struct {
		Duration int    `json:"duration"`
		Easing   string `json:"easing"`
	} `json:"animation"`
	Layout struct {
		Padding struct {
			Left   int `json:"left"`
			Right  int `json:"right"`
			Top    int `json:"top"`
			Bottom int `json:"bottom"`
		} `json:"padding"`
	} `json:"layout"`
	Legend struct {
		Display   bool   `json:"display"`
		Position  string `json:"position"`
		Align     string `json:"align"`
		FullWidth bool   `json:"full_width"`
		Reverse   bool   `json:"reverse"`
		Labels    struct {
			BoxWidth      int    `json:"box_width"`
			FontSize      int    `json:"fontSize"`
			FontStyle     string `json:"fontStyle"`
			FontColor     Color  `json:"fontColor"`
			FontFamily    string `json:"fontFamily"`
			Padding       int    `json:"padding"`
			UsePointStyle bool   `json:"usePointStyle"`
		} `json:"labels"`
		Rtl           bool   `json:"rtl"`
		TextDirection string `json:"text_direction"`
	} `json:"legend"`
	Title struct {
		Display    bool   `json:"display"`
		Position   string `json:"position"`
		FontSize   int    `json:"fontSize"`
		FontFamily string `json:"fontFamily"`
		FontColor  Color  `json:"fontColor"`
		FontStyle  string `json:"fontStyle"`
		Padding    int    `json:"padding"`
		LineHeight int    `json:"lineHeight"`
		Text       string `json:"text"`
	} `json:"title"`
	Tooltips struct {
		Enabled            bool   `json:"enabled"`
		Mode               string `json:"mode"`
		Intersect          bool   `json:"intersect"`
		Position           string `json:"position"`
		BackgroundColor    Color  `json:"backgroundColor"`
		TitleFontFamily    string `json:"titleFontFamily"`
		TitleFontSize      int    `json:"titleFontSize"`
		TitleFontStyle     string `json:"titleFontStyle"`
		TitleFontColor     Color  `json:"titleFontColor"`
		TitleAlign         string `json:"titleAlign"`
		TitleSpacing       int    `json:"titleSpacing"`
		TitleMarginBottom  int    `json:"titleMarginBottom"`
		BodyFontFamily     string `json:"bodyFontFamily"`
		BodyFontSize       int    `json:"bodyFontSize"`
		BodyFontStyle      string `json:"bodyFontStyle"`
		BodyFontColor      Color  `json:"bodyFontColor"`
		BodyAlign          string `json:"bodyAlign"`
		BodySpacing        int    `json:"bodySpacing"`
		FooterFontFamily   string `json:"footerFontFamily"`
		FooterFontSize     int    `json:"footerFontSize"`
		FooterFontStyle    string `json:"footerFontStyle"`
		FooterFontColor    Color  `json:"footerFontColor"`
		FooterAlign        string `json:"footerAlign"`
		FooterSpacing      int    `json:"footerSpacing"`
		FooterMarginTop    int    `json:"footerMarginTop"`
		XPadding           int    `json:"xPadding"`
		YPadding           int    `json:"yPadding"`
		CaretPadding       int    `json:"caretPadding"`
		CaretSize          int    `json:"caretSize"`
		CornerRadius       int    `json:"cornerRadius"`
		MultiKeyBackground Color  `json:"multiKeyBackground"`
		DisplayColors      bool   `json:"displayColors"`
		BorderColor        Color  `json:"borderColor"`
		BorderWidth        int    `json:"borderWidth"`
		Rtl                bool   `json:"rtl"`
		TextDirection      string `json:"textDirection"`
	} `json:"tooltips"`
	Elements struct {
		Point struct {
			Radius           int    `json:"radius"`
			PointStyle       string `json:"pointStyle"`
			Rotation         int    `json:"rotation"`
			BackgroundColor  Color  `json:"backgroundColor"`
			BorderWidth      int    `json:"borderWidth"`
			BorderColor      Color  `json:"borderColor"`
			HitRadius        int    `json:"hitRadius"`
			HoverRadius      int    `json:"hoverRadius"`
			HoverBorderWidth int    `json:"hoverBorderWidth"`
		} `json:"point"`
		Line struct {
			tension                int
			backgroundColor        Color
			borderWidth            int
			borderColor            Color
			borderCapStyle         string
			borderDash             int
			borderDashOffset       int
			borderJoinStyle        string
			capBezierPoints        bool
			cubicInterpolationMode string
			fill                   bool
			stepped                bool
		} `json:"line"`
		Rectangle struct {
			BackgroundColor Color  `json:"backgroundColor"`
			BorderWidth     int    `json:"borderWidth"`
			BorderColor     Color  `json:"borderColor"`
			BorderSkipped   string `json:"borderSkipped"`
		} `json:"rectangle"`
		Arc struct {
			Angle           int    `json:"angle"`
			BackgroundColor Color  `json:"backgroundColor"`
			BorderAlign     string `json:"borderAlign"`
			BorderColor     Color  `json:"borderColor"`
			BorderWidth     int    `json:"borderWidth"`
		} `json:"arc"`
	} `json:"elements"`
}

type Attributes struct {
	Labels []string `json:"labels"`
}

type DataSets []DataSet

type DataSet struct {
	Label string    `json:"label"`
	Data  []float64 `json:"data"`
	Type  string    `json:"type"`
}

type Color string

func NewChart() *Chart {
	return new(Chart)
}

func (c Chart) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("chartjs").
		Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
			"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
				if cdnUrl == "" {
					return prefixUrl + assetsUrl
				}
				return cdnUrl + assetsUrl
			},
			"isLinkUrl": func(s string) bool {
				return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
			},
		}).
		Parse(List["chartjs"])

	if err != nil {
		logger.Error("Chart GetTemplate Error: ", err)
	}

	return tmpl, "chartjs"
}

func (c Chart) GetAssetList() []string {
	return AssetsList
}

func (c Chart) GetAsset(name string) ([]byte, error) {
	return Asset(name[1:])
}

func (c Chart) IsAPage() bool {
	return false
}

func (c Chart) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := c.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, c)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
