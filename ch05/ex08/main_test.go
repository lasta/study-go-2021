package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_forEachNode(t *testing.T) {
	type Expected struct{
		found bool
		isNil bool
		tag string // FIXME
	}
	testCases := []struct {
		name     string
		HTML     string
		id string
		expected Expected
	}{
		{
			name: "指定した ID がない場合, false",
			HTML: `
<html>
<head><title>title</title></head>
<body><h1 id="Header 1">Header 1</h1><p>Paragraph</p></body>
</html>
`,
			id: "foo" ,
			expected: Expected{
				found: false,
				isNil: true,
			},
		},
		{
			name: "指定した ID がある場合, true",
			HTML: `
<html>
<head><title>title</title></head>
<body><h1 id="bar">Header 1</h1><p>Paragraph</p></body>
</html>
`,
			id: "bar" ,
			expected: Expected{
				found: true,
				isNil: false,
				tag: "h1",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(testCase.HTML))
			if err != nil {
				t.Error(err)
				t.Fail()
			}

			node := ElementByID(doc, testCase.id)

			if !testCase.expected.found {
				assert.Equal(t, testCase.expected.isNil, node == nil)
				return
			}

			assert.Equal(t, testCase.expected.isNil, node == nil)
			assert.Equal(t, testCase.expected.tag, node.Data)
		})
	}
}
