package web

import (
	"fmt"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	"github.com/sclevine/agouti"
	"github.com/stretchr/testify/assert"
)

type Page struct {
	*agouti.Page
	T      *testing.T
	Driver *agouti.WebDriver
	Quit   chan struct{}
}

func (page *Page) Destroy() {
	if r := recover(); r != nil {
		debug.PrintStack()
		fmt.Println("Recovered in f", r)
		_ = page.Page.Destroy()
		_ = page.Driver.Stop()
		page.T.Fail()
		page.Quit <- struct{}{}
	}
}

func (page *Page) Wait(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

func (page *Page) Contain(s string) {
	content, err := page.HTML()
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, strings.Contains(content, s), true)
}

func (page *Page) NoContain(s string) {
	content, err := page.HTML()
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, strings.Contains(content, s), false)
}

func (page *Page) Css(xpath, css, res string) {
	style, err := page.FindByXPath(xpath).CSS(css)
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, style, res)
}

func (page *Page) CssS(s *agouti.Selection, css, res string) {
	style, err := s.CSS(css)
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, style, res)
}

func (page *Page) Text(xpath, text string) {
	mli1, err := page.FindByXPath(xpath).Text()
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, mli1, text)
}

func (page *Page) MoveMouseBy(xOffset, yOffset int) {
	assert.Equal(page.T, page.Page.MoveMouseBy(xOffset, yOffset), nil)
}

func (page *Page) Display(xpath string) {
	page.Css(xpath, "display", "block")
}

func (page *Page) Nondisplay(xpath string) {
	page.Css(xpath, "display", "none")
}

func (page *Page) Value(xpath, value string) {
	val, err := page.FindByXPath(xpath).Attribute("value")
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, val, value)
}

func (page *Page) Click(xpath string, intervals ...int) {
	assert.Equal(page.T, page.FindByXPath(xpath).Click(), nil)
	interval := 1
	if len(intervals) > 0 {
		interval = intervals[0]
	}
	page.Wait(interval)
}

func (page *Page) ClickS(s *agouti.Selection, intervals ...int) {
	assert.Equal(page.T, s.Click(), nil)
	interval := 1
	if len(intervals) > 0 {
		interval = intervals[0]
	}
	page.Wait(interval)
}

func (page *Page) Attr(s *agouti.Selection, attr, res string) {
	style, err := s.Attribute(attr)
	assert.Equal(page.T, err, nil)
	assert.Equal(page.T, style, res)
}

func (page *Page) Fill(xpath, content string) {
	assert.Equal(page.T, page.FindByXPath(xpath).Fill(content), nil)
}

func (page *Page) NavigateTo(path string) {
	assert.Equal(page.T, page.Navigate(path), nil)
	page.Wait(2)
}
