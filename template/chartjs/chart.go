package chartjs

import (
	"html/template"

	template2 "github.com/GoAdminGroup/go-admin/template"
)

type Chart struct {
	*template2.BaseComponent

	ID     string
	Title  template.HTML
	Js     template.JS
	Height int

	JsContentOptions *Options

	dataSetIndex int
}

func (c *Chart) SetID(id string) *Chart {
	c.ID = id
	return c
}

func (c *Chart) SetTitle(title template.HTML) *Chart {
	c.Title = title
	return c
}

func (c *Chart) SetHeight(height int) *Chart {
	c.Height = height
	return c
}

func (c *Chart) SetOptionAnimationDuration(duration int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Animation == nil {
		c.JsContentOptions.Animation = new(OptionAnimation)
	}
	c.JsContentOptions.Animation.Duration = duration
}

func (c *Chart) SetOptionAnimationEasing(easing string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Animation == nil {
		c.JsContentOptions.Animation = new(OptionAnimation)
	}
	c.JsContentOptions.Animation.Easing = easing
}

func (c *Chart) SetOptionLayoutPaddingLeft(left int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Layout == nil {
		c.JsContentOptions.Layout = new(OptionLayout)
	}
	c.JsContentOptions.Layout.Padding.Left = left
}

func (c *Chart) SetOptionLayoutPaddingRight(right int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Layout == nil {
		c.JsContentOptions.Layout = new(OptionLayout)
	}
	c.JsContentOptions.Layout.Padding.Right = right
}

func (c *Chart) SetOptionLayoutPaddingTop(top int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Layout == nil {
		c.JsContentOptions.Layout = new(OptionLayout)
	}
	c.JsContentOptions.Layout.Padding.Top = top
}

func (c *Chart) SetOptionLayoutPaddingBottom(bottom int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Layout == nil {
		c.JsContentOptions.Layout = new(OptionLayout)
	}
	c.JsContentOptions.Layout.Padding.Bottom = bottom
}

func (c *Chart) SetOptionLegendDisplay(display bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.Display = display
}

func (c *Chart) SetOptionLegendPosition(position string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.Position = position
}

func (c *Chart) SetOptionLegendAlign(align string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.Align = align
}

func (c *Chart) SetOptionLegendFullWidt(fullWidth bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.FullWidth = fullWidth
}

func (c *Chart) SetOptionLegendRevers(reverse bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.Reverse = reverse
}

func (c *Chart) SetOptionLegendRt(rtl bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.Rtl = rtl
}

func (c *Chart) SetOptionLegendTextDirection(textDirection string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.TextDirection = textDirection
}

func (c *Chart) SetOptionLegendLabels(labels *OptionLegendLabel) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Legend == nil {
		c.JsContentOptions.Legend = new(OptionLegend)
	}
	c.JsContentOptions.Legend.Labels = labels
}

func (c *Chart) SetOptionTitleDisplay(display bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.Display = display
}

func (c *Chart) SetOptionTitleFontSize(fontSize int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.FontSize = fontSize
}

func (c *Chart) SetOptionTitlePosition(position string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.Position = position
}

func (c *Chart) SetOptionTitleFontFamily(fontFamily string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.FontFamily = fontFamily
}

func (c *Chart) SetOptionTitleFontColor(fontColor Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.FontColor = fontColor
}

func (c *Chart) SetOptionTitleFontStyle(fontStyle string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.FontStyle = fontStyle
}

func (c *Chart) SetOptionTitlePadding(padding int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.Padding = padding
}

func (c *Chart) SetOptionTitleLineHeight(lineHeight int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.LineHeight = lineHeight
}

func (c *Chart) SetOptionTitleText(text string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Title == nil {
		c.JsContentOptions.Title = new(OptionTitle)
	}
	c.JsContentOptions.Title.Text = text
}

func (c *Chart) SetOptionTooltipsEnabled(enabled bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.Enabled = enabled
}

func (c *Chart) SetOptionTooltipsMode(mode string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.Mode = mode
}

func (c *Chart) SetOptionTooltipsIntersect(intersect bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.Intersect = intersect
}

func (c *Chart) SetOptionTooltipsPosition(position string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.Position = position
}

func (c *Chart) SetOptionTooltipsBackgroundColor(backgroundColor Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BackgroundColor = backgroundColor
}

func (c *Chart) SetOptionTooltipsTitleFontFamily(titleFontFamily string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleFontFamily = titleFontFamily
}

func (c *Chart) SetOptionTooltipsTitleFontSize(titleFontSize int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleFontSize = titleFontSize
}

func (c *Chart) SetOptionTooltipsTitleFontStyle(titleFontStyle string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleFontStyle = titleFontStyle
}

func (c *Chart) SetOptionTooltipsTitleFontColor(titleFontColor Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleFontColor = titleFontColor
}

func (c *Chart) SetOptionTooltipsTitleAlign(titleAlign string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleAlign = titleAlign
}

func (c *Chart) SetOptionTooltipsTitleSpacing(titleSpacing int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleSpacing = titleSpacing
}

func (c *Chart) SetOptionTooltipsTitleMarginBottom(titleMarginBottom int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TitleMarginBottom = titleMarginBottom
}

func (c *Chart) SetOptionTooltipsBodyFontFamily(bodyFontFamily string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BodyFontFamily = bodyFontFamily
}

func (c *Chart) SetOptionTooltipsBodyFontSize(bodyFontSize int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BodyFontSize = bodyFontSize
}

func (c *Chart) SetOptionTooltipsBodyFontStyle(bodyFontStyle string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BodyFontStyle = bodyFontStyle
}

func (c *Chart) SetOptionTooltipsBodyFontColor(bodyFontColor Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BodyFontColor = bodyFontColor
}

func (c *Chart) SetOptionTooltipsBodyAlign(bodyAlign string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BodyAlign = bodyAlign
}

func (c *Chart) SetOptionTooltipsBodySpacing(bodySpacing int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BodySpacing = bodySpacing
}

func (c *Chart) SetOptionTooltipsFooterFontFamily(footerFontFamily string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterFontFamily = footerFontFamily
}

func (c *Chart) SetOptionTooltipsFooterFontSize(footerFontSize int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterFontSize = footerFontSize
}

func (c *Chart) SetOptionTooltipsFooterFontStyle(footerFontStyle string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterFontStyle = footerFontStyle
}

func (c *Chart) SetOptionTooltipsFooterFontColor(footerFontColor Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterFontColor = footerFontColor
}

func (c *Chart) SetOptionTooltipsFooterAlign(footerAlign string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterAlign = footerAlign
}

func (c *Chart) SetOptionTooltipsFooterSpacing(footerSpacing int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterSpacing = footerSpacing
}

func (c *Chart) SetOptionTooltipsFooterMarginTop(footerMarginTop int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.FooterMarginTop = footerMarginTop
}

func (c *Chart) SetOptionTooltipsXPadding(xPadding int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.XPadding = xPadding
}

func (c *Chart) SetOptionTooltipsYPadding(yPadding int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.YPadding = yPadding
}

func (c *Chart) SetOptionTooltipsCaretPadding(caretPadding int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.CaretPadding = caretPadding
}

func (c *Chart) SetOptionTooltipsCaretSize(caretSize int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.CaretSize = caretSize
}

func (c *Chart) SetOptionTooltipsCornerRadius(cornerRadius int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.CornerRadius = cornerRadius
}

func (c *Chart) SetOptionTooltipsMultiKeyBackground(multiKeyBackground Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.MultiKeyBackground = multiKeyBackground
}

func (c *Chart) SetOptionTooltipsDisplayColors(displayColors bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.DisplayColors = displayColors
}

func (c *Chart) SetOptionTooltipsBorderColor(borderColor Color) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BorderColor = borderColor
}

func (c *Chart) SetOptionTooltipsBorderWidth(borderWidth int) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.BorderWidth = borderWidth
}

func (c *Chart) SetOptionTooltipsRtl(rtl bool) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.Rtl = rtl
}

func (c *Chart) SetOptionTooltipsTextDirection(textDirection string) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Tooltips == nil {
		c.JsContentOptions.Tooltips = new(OptionTooltips)
	}
	c.JsContentOptions.Tooltips.TextDirection = textDirection
}

func (c *Chart) SetOptionElementPoint(point *OptionElementPoint) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Elements == nil {
		c.JsContentOptions.Elements = new(OptionElement)
	}
	c.JsContentOptions.Elements.Point = point
}

func (c *Chart) SetOptionElementLine(line *OptionElementLine) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Elements == nil {
		c.JsContentOptions.Elements = new(OptionElement)
	}
	c.JsContentOptions.Elements.Line = line
}

func (c *Chart) SetOptionElementArc(arc *OptionElementArc) {
	if c.JsContentOptions == nil {
		c.JsContentOptions = new(Options)
	}
	if c.JsContentOptions.Elements == nil {
		c.JsContentOptions.Elements = new(OptionElement)
	}
	c.JsContentOptions.Elements.Arc = arc
}

func (c *Chart) SetOptionElementRectangle(rectangle *OptionElementRectangle) {
	if c.JsContentOptions.Elements == nil {
		c.JsContentOptions.Elements = new(OptionElement)
	}
	c.JsContentOptions.Elements.Rectangle = rectangle
}

type JsContent struct {
	Type    string   `json:"type,omitempty"`
	Options *Options `json:"options,omitempty"`
}

type OptionAnimation struct {
	Duration int    `json:"duration,omitempty"`
	Easing   string `json:"easing,omitempty"`
}

type OptionLayout struct {
	Padding struct {
		Left   int `json:"left,omitempty"`
		Right  int `json:"right,omitempty"`
		Top    int `json:"top,omitempty"`
		Bottom int `json:"bottom,omitempty"`
	} `json:"padding,omitempty"`
}

type OptionLegend struct {
	Display       bool               `json:"display,omitempty"`
	Position      string             `json:"position,omitempty"`
	Align         string             `json:"align,omitempty"`
	FullWidth     bool               `json:"full_width,omitempty"`
	Reverse       bool               `json:"reverse,omitempty"`
	Rtl           bool               `json:"rtl,omitempty"`
	TextDirection string             `json:"text_direction,omitempty"`
	Labels        *OptionLegendLabel `json:"labels,omitempty"`
}

type OptionLegendLabel struct {
	BoxWidth      int    `json:"box_width,omitempty"`
	FontSize      int    `json:"fontSize,omitempty"`
	FontStyle     string `json:"fontStyle,omitempty"`
	FontColor     Color  `json:"fontColor,omitempty"`
	FontFamily    string `json:"fontFamily,omitempty"`
	Padding       int    `json:"padding,omitempty"`
	UsePointStyle bool   `json:"usePointStyle,omitempty"`
}

type OptionTitle struct {
	Display    bool   `json:"display,omitempty"`
	Position   string `json:"position,omitempty"`
	FontSize   int    `json:"fontSize,omitempty"`
	FontFamily string `json:"fontFamily,omitempty"`
	FontColor  Color  `json:"fontColor,omitempty"`
	FontStyle  string `json:"fontStyle,omitempty"`
	Padding    int    `json:"padding,omitempty"`
	LineHeight int    `json:"lineHeight,omitempty"`
	Text       string `json:"text,omitempty"`
}

type OptionTooltips struct {
	Enabled            bool   `json:"enabled,omitempty"`
	Mode               string `json:"mode,omitempty"`
	Intersect          bool   `json:"intersect,omitempty"`
	Position           string `json:"position,omitempty"`
	BackgroundColor    Color  `json:"backgroundColor,omitempty"`
	TitleFontFamily    string `json:"titleFontFamily,omitempty"`
	TitleFontSize      int    `json:"titleFontSize,omitempty"`
	TitleFontStyle     string `json:"titleFontStyle,omitempty"`
	TitleFontColor     Color  `json:"titleFontColor,omitempty"`
	TitleAlign         string `json:"titleAlign,omitempty"`
	TitleSpacing       int    `json:"titleSpacing,omitempty"`
	TitleMarginBottom  int    `json:"titleMarginBottom,omitempty"`
	BodyFontFamily     string `json:"bodyFontFamily,omitempty"`
	BodyFontSize       int    `json:"bodyFontSize,omitempty"`
	BodyFontStyle      string `json:"bodyFontStyle,omitempty"`
	BodyFontColor      Color  `json:"bodyFontColor,omitempty"`
	BodyAlign          string `json:"bodyAlign,omitempty"`
	BodySpacing        int    `json:"bodySpacing,omitempty"`
	FooterFontFamily   string `json:"footerFontFamily,omitempty"`
	FooterFontSize     int    `json:"footerFontSize,omitempty"`
	FooterFontStyle    string `json:"footerFontStyle,omitempty"`
	FooterFontColor    Color  `json:"footerFontColor,omitempty"`
	FooterAlign        string `json:"footerAlign,omitempty"`
	FooterSpacing      int    `json:"footerSpacing,omitempty"`
	FooterMarginTop    int    `json:"footerMarginTop,omitempty"`
	XPadding           int    `json:"xPadding,omitempty"`
	YPadding           int    `json:"yPadding,omitempty"`
	CaretPadding       int    `json:"caretPadding,omitempty"`
	CaretSize          int    `json:"caretSize,omitempty"`
	CornerRadius       int    `json:"cornerRadius,omitempty"`
	MultiKeyBackground Color  `json:"multiKeyBackground,omitempty"`
	DisplayColors      bool   `json:"displayColors,omitempty"`
	BorderColor        Color  `json:"borderColor,omitempty"`
	BorderWidth        int    `json:"borderWidth,omitempty"`
	Rtl                bool   `json:"rtl,omitempty"`
	TextDirection      string `json:"textDirection,omitempty"`
}

type OptionElement struct {
	Point     *OptionElementPoint     `json:"point,omitempty"`
	Line      *OptionElementLine      `json:"line,omitempty"`
	Rectangle *OptionElementRectangle `json:"rectangle,omitempty"`
	Arc       *OptionElementArc       `json:"arc,omitempty"`
}

type OptionElementPoint struct {
	Radius           int    `json:"radius,omitempty"`
	PointStyle       string `json:"pointStyle,omitempty"`
	Rotation         int    `json:"rotation,omitempty"`
	BackgroundColor  Color  `json:"backgroundColor,omitempty"`
	BorderWidth      int    `json:"borderWidth,omitempty"`
	BorderColor      Color  `json:"borderColor,omitempty"`
	HitRadius        int    `json:"hitRadius,omitempty"`
	HoverRadius      int    `json:"hoverRadius,omitempty"`
	HoverBorderWidth int    `json:"hoverBorderWidth,omitempty"`
}

type OptionElementLine struct {
	Tension                int    `json:"tension,omitempty"`
	BackgroundColor        Color  `json:"background_color,omitempty"`
	BorderWidth            int    `json:"border_width,omitempty"`
	BorderColor            Color  `json:"border_color,omitempty"`
	BorderCapStyle         string `json:"border_cap_style,omitempty"`
	BorderDash             int    `json:"border_dash,omitempty"`
	BorderDashOffset       int    `json:"border_dash_offset,omitempty"`
	BorderJoinStyle        string `json:"border_join_style,omitempty"`
	CapBezierPoints        bool   `json:"cap_bezier_points,omitempty"`
	CubicInterpolationMode string `json:"cubic_interpolation_mode,omitempty"`
	Fill                   bool   `json:"fill,omitempty"`
	Stepped                bool   `json:"stepped,omitempty"`
}

type OptionElementRectangle struct {
	BackgroundColor Color  `json:"backgroundColor,omitempty"`
	BorderWidth     int    `json:"borderWidth,omitempty"`
	BorderColor     Color  `json:"borderColor,omitempty"`
	BorderSkipped   string `json:"borderSkipped,omitempty"`
}

type OptionElementArc struct {
	Angle           int    `json:"angle,omitempty"`
	BackgroundColor Color  `json:"backgroundColor,omitempty"`
	BorderAlign     string `json:"borderAlign,omitempty"`
	BorderColor     Color  `json:"borderColor,omitempty"`
	BorderWidth     int    `json:"borderWidth,omitempty"`
}

type Options struct {
	Animation *OptionAnimation `json:"animation,omitempty"`
	Layout    *OptionLayout    `json:"layout,omitempty"`
	Legend    *OptionLegend    `json:"legend,omitempty"`
	Title     *OptionTitle     `json:"title,omitempty"`
	Tooltips  *OptionTooltips  `json:"tooltips,omitempty"`
	Elements  *OptionElement   `json:"elements,omitempty"`
}

type Attributes struct {
	Labels []string `json:"labels,omitempty"`
}

type DataSets []DataSet

type DataSet struct {
	Label string    `json:"label,omitempty"`
	Data  []float64 `json:"data,omitempty"`
	Type  string    `json:"type,omitempty"`
}

type Color string

func NewChart() *Chart {
	return &Chart{
		BaseComponent: &template2.BaseComponent{
			Name:     "chartjs",
			HTMLData: List["chartjs"],
		},
	}
}

func (c *Chart) GetAssetList() []string               { return AssetsList }
func (c *Chart) GetAsset(name string) ([]byte, error) { return Asset(name[1:]) }
func (c *Chart) GetContent() template.HTML            { return c.GetContentWithData(c) }
