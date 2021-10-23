package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_visit(t *testing.T) {
	testCases := []struct {
		name string
		input    string
		expected []string
	}{
		{
			name: "空の HTML の場合 nil",
			input:    ``,
			expected: []string(nil),
		},
		{
			name: "anchor 要素がない場合 nil",
			input:    `<html><body></body></html>`,
			expected: []string(nil),
		},
		{
			name: "anchor 要素がある場合, URL のリストを生成 (孫なし)",
			input: `
<html>
<body>
    <a href="https://golang.org/">Golang</a>
    <a href="https://github.com/">GitHub</a>
</body>
</html>
`,
			expected: []string{
				"https://golang.org/",
				"https://github.com/",
			},
		},
		{
			name: "anchor 要素がある場合, URL のリストを生成 (孫あり)",
			input: `
<html>
<body>
    <div>
	    <a href="https://golang.org/">Golang</a>
	    <a href="https://github.com/">GitHub</a>
        <div>
            <a href="https://www.youtube.com/">YouTube</a>
        </div>
    </div>
</body>
</html>
`,
			expected: []string{
				"https://golang.org/",
				"https://github.com/",
				"https://www.youtube.com/",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}
			actual := visit(nil, doc)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
