package utils

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressedContent(t *testing.T) {
	htmlContent1 := template.HTML(`
<html>
	<body>

		<h1>Test</h1>                         



		<p>CompressedContent</p>           

	</body>

</html>
`)
	htmlContent2 := htmlContent1
	CompressedContent(&htmlContent2)
	t.Log(len(htmlContent1) > len(htmlContent2))
}

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, true, CompareVersion("v1.2.4", "v1.2.5"))
	assert.Equal(t, false, CompareVersion("v1.2.4", "v1.2.4"))
	assert.Equal(t, false, CompareVersion("v1.2.4", "v1.2.3"))
	assert.Equal(t, false, CompareVersion("v1.2.4", "v1.1.3"))
	assert.Equal(t, true, CompareVersion("v1.2.4", "v1.3.3"))
	assert.Equal(t, false, CompareVersion("v1.2.4", "v0.3.3"))

	assert.Equal(t, true, CompareVersion("<v1.2.4", "v0.3.3"))
	assert.Equal(t, false, CompareVersion("<v1.2.4", "v1.2.5"))
	assert.Equal(t, true, CompareVersion("<=v1.2.4", "v1.2.4"))
	assert.Equal(t, true, CompareVersion("<=v1.2.4", "v1.2.3"))
	assert.Equal(t, false, CompareVersion("<=v1.2.4", "v1.2.5"))
	assert.Equal(t, true, CompareVersion(">v1.2.4", "v1.2.5"))
	assert.Equal(t, false, CompareVersion(">v1.2.4", "v1.2.4"))
	assert.Equal(t, true, CompareVersion(">=v1.2.4", "v1.2.4"))
	assert.Equal(t, true, CompareVersion(">=v1.2.4", "v1.2.5"))
	assert.Equal(t, false, CompareVersion(">=v1.2.4", "v1.2.3"))
	assert.Equal(t, false, CompareVersion("=v1.2.4", "v1.2.3"))
	assert.Equal(t, true, CompareVersion("=v1.2.4", "v1.2.4"))
	assert.Equal(t, true, CompareVersion("= v1.2.4", "v1.2.4"))
}
