package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func Test_countTags(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		delimiter string
		expected  string
	}{
		{
			name:      "空HTMLの場合, 空文字を生成",
			input:     "",
			delimiter: "\n",
			expected:  "",
		},
		{
			name:      "テキストが空の場合, 空文字を生成 (<html> のみ)",
			input:     "<html></html>",
			delimiter: "\n",
			expected:  "",
		},
		{
			name:      "非テキストタグ群のみの場合, 空文字を生成 (<html>, <head>, <body>)",
			input:     "<html><head></head><body></body></html>",
			delimiter: "\n",
			expected:  "",
		},
		{
			name:      "Blacklist にあるタグのみの場合, 空文字を生成 (<script>)",
			input:     `<html><head></head><body><script>document.write("script")</script></body></html>`,
			delimiter: "\n",
			expected:  "",
		},
		{
			name:      "Blacklist にあるタグのみの場合, 空文字を生成 (<style>)",
			input:     `<html><head></head><body><style>H1 { color:black }</style></body></html>`,
			delimiter: "\n",
			expected:  "",
		},
		{
			name: "Blacklist にあるタグと Blacklist にないテキストタグが指定された場合, Blacklist に指定されていないタグのデータのみで生成",
			input: "<html>" +
				"<head></head>" +
				"<body>" +
				"<style>H1 { color:black }</style>" +
				"<h1>Header 1</h1>" +
				"<script>document.write(\"script\")></script>" +
				"</body>" +
				"</html>",
			delimiter: "\n",
			expected:  "Header 1\n",
		},
		{
			name: "テキストノードが複数ある場合, デリミタで結合して生成",
			input: "<html>" +
				"<header>" +
				"<title>Title</title>" +
				"</header>" +
				"<body>" +
				"<p>" +
				"<h1>Header 1</h1>" +
				"Paragraph1" +
				"</p>" +
				"</body>" +
				"</html>",
			delimiter: "\n",
			expected: "Title\n" +
				"Header 1\n" +
				"Paragraph1\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(testCase.input))
			if err != nil {
				t.Fatal(err)
			}
			var writer bytes.Buffer
			extractTextNodes(&writer, doc, testCase.delimiter)
			actual := writer.String()
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func Test_contains(t *testing.T) {
	testCases := []struct {
		name      string
		blacklist TagsBlackList
		tag       HtmlTag
		expected  bool
	}{
		{
			name:      "Blacklist が nil かつ調査対象が空文字の場合, false",
			blacklist: nil,
			tag:       "",
			expected:  false,
		},
		{
			name:      "Blacklist が nil かつ調査対象が空文字ではない場合, false",
			blacklist: nil,
			tag:       "a",
			expected:  false,
		},
		{
			name:      "Blacklist が空かつ調査対象が空文字の場合, false",
			blacklist: TagsBlackList{},
			tag:       "",
			expected:  false,
		},
		{
			name:      "Blacklist が空かつ調査対象が空文字ではない場合, false",
			blacklist: TagsBlackList{},
			tag:       "a",
			expected:  false,
		},
		{
			name:      "Blacklist が空ではないが調査対象のものが登録されていない場合, false (長さ1)",
			blacklist: TagsBlackList{"a"},
			tag:       "b",
			expected:  false,
		},
		{
			name:      "Blacklist が空ではないが調査対象のものが登録されていない場合, false (長さ2)",
			blacklist: TagsBlackList{"a", "h1"},
			tag:       "b",
			expected:  false,
		},
		{
			name:      "Blacklist に登録されているタグが指定された場合, true",
			blacklist: TagsBlackList{"a"},
			tag:       "a",
			expected:  true,
		},
		{
			name:      "Blacklist が指定されたタグを含んでいる場合, true",
			blacklist: TagsBlackList{"a", "b"},
			tag:       "b",
			expected:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.blacklist.contains(testCase.tag)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
