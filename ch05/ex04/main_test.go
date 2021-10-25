package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_visit(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "空のHTMLの場合, 空文字",
			input:    "",
			expected: "",
		},
		{
			name: "抽出対象の要素を含まない場合, 空文字",
			input: `
<html>
<head><title>Title</title></head>
<body>
    <h1>Header 1</h1>
    <p>Paragraph 1</p>
</body>
</html>
`,
			expected: "",
		},
		{
			name: "抽出対象のタグを含むが URL 要素がない場合, 空文字",
			input: `

<html>
<head><title>Title</title></head>
<body><a>empty anchor</a></body>
</html>
`,
			expected: "",
		},
		{
			name: "抽出対象のタグを含み, URL 要素がある場合, それを含めて出力",
			input: `
<html>
<head>
   <link rel="prev" href="prev.html" />
   <script src="script1.js"></script>
   <link rel="next" href="next.html" />
   <script src="script2.js"></script>
</head>
<body>
   <img src="image1.png" alt="image1" />
   <a href="./a.html">a.html</a>
   <img src="image2.png" alt="image2" />
   <a href="./b.html">b.html</a>
</body>
</html>
`,
			expected: "link    href    prev.html\n" +
				"script  src     script1.js\n" +
				"link    href    next.html\n" +
				"script  src     script2.js\n" +
				"img     src     image1.png\n" +
				"a       href    ./a.html\n" +
				"img     src     image2.png\n" +
				"a       href    ./b.html\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}
			var buf bytes.Buffer
			extractUrl(&buf, doc)
			actual := buf.String()
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
