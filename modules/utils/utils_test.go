package utils

import (
	"html/template"
	"testing"
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
