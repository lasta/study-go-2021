package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_countTags(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "空の HTML の場合, 最低限のタグが自動挿入されカウントされる",
			input:    ``,
			expected: map[string]int{"body":1, "head":1, "html":1},
		},
		{
			name:     "最低限の要素を満たしていない場合, 最低限のタグが自動挿入されカウントされる",
			input:    `<html></html>`,
			expected: map[string]int{"body":1, "head":1, "html":1},
		},
		{
			name:     "子要素あり, 孫要素なし",
			input:    `
<html>
<head></head>
<body></body>
</html>
`,
			expected: map[string]int{"html": 1, "head": 1, "body": 1},
		},
		{
			name: "孫要素あり, 重複あり",
			input: `
<html>
<head>
	<title>Title</title>
</head>
<body>
	<h1>Golang</h1>
	<p>
		<a href="http://golang.org/">Golang</a>
	</p>
	<h1>Google</h1>
	<p>
		<a href="http://google.com/">Google</a>
	</p>
	<h1>YouTube</h1>
	<p>
		<a href="http://youtube.com/">YouTube</a>
	</p>
</body>
</html>
`,
			expected: map[string]int{
				"html": 1,
				"head": 1,
				"title": 1,
				"body": 1,
				"h1": 3,
				"p": 3,
				"a": 3,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}
			actual := map[string]int{}
			countTags(actual, doc)

			assert.Equal(t, testCase.expected, actual)
		})
	}
}
