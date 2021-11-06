package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_forEachNode(t *testing.T) {
	testCases := []struct {
		name     string
		HTML     string
		expected string
	}{
		{
			name: "各要素がある HTML の場合, 値も含めて出力",
			HTML: `
<html>

<head><title> Title </title></head>

<body>
  <h1>Header 1</h1>

  <p> Paragraph 1
  <img src="image.png" alt="this is PNG image" />
  </p>
</body>
</html>
`,
			expected: `<html>
  <head>
    <title>
      Title
    </title>
  </head>
  <body>
    <h1>
      Header 1
    </h1>
    <p>
      Paragraph 1
      <img src="image.png" alt="this is PNG image" />
    </p>
  </body>
</html>
`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(testCase.HTML))
			if err != nil {
				t.Error(err)
				t.Fail()
			}

			writer := bytes.NewBuffer(nil)
			forEachNode(doc, startElement, endElement, writer)
			actual := writer.String()
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
