package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_CountWordsAndImage(t *testing.T) {
	type expected struct {
		words  int
		images int
		err    error
	}
	type TestCase struct {
		name     string
		url      string
		response string
		expected
	}
	testCases := []TestCase{
		{
			name:     "Get に成功した場合, 適切な値を返却",
			url:      "/",
			response: `<html><head><title>title</title></head><body><p>paragraph 1</p><img src="image.png" /></body></html>`,
			expected: expected{
				words:  3,
				images: 1,
				err:    nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
				io.WriteString(writer, testCase.response)
			})
			server := httptest.NewServer(mux)
			defer server.Close()
			words, images, err := CountWordsAndImages(server.URL + testCase.url)

			assert.Equal(t, testCase.words, words)
			assert.Equal(t, testCase.images, images)
			assert.Equal(t, testCase.err, err)
		})
	}
}

func Test_countWordsAndImage(t *testing.T) {
	type expected struct {
		words  int
		images int
	}
	type TestCase struct {
		name  string
		input string
		expected
	}

	testCases := []TestCase{
		{
			name:  "空のHTMLが渡された場合, 0を返却",
			input: "",
			expected: expected{
				words:  0,
				images: 0,
			},
		},
		{
			name:  "img が含まれる場合, words が0, images が1",
			input: `<html><body><img src="image.png" /></body></html>`,
			expected: expected{
				words:  0,
				images: 1,
			},
		},
		{
			name:  "text node が含まれる場合, words が1, images が0",
			input: `<html><body><p>paragraph</p></body></html>`,
			expected: expected{
				words:  1,
				images: 0,
			},
		},
		{
			name:  "text node (2単語) と image が含まれる場合, words が2, images が1",
			input: `<html><body><p>paragraph 1</p><img src="image.png" /></body></html>`,
			expected: expected{
				words:  2,
				images: 1,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			doc, err := html.Parse(strings.NewReader(c.input))
			if err != nil {
				t.Fatal(err)
			}
			words, images := countWordsAndImages(doc)

			assert.Equal(t, c.words, words)
			assert.Equal(t, c.images, images)
		})
	}
}

func Test_countWords(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{" ", 0},
		{"a", 1},
		{"ab", 1},
		{"a ", 1},
		{"a b", 2},
		{"a\tb", 2},
		{"a\nb", 2},
		{"a\vb", 2},
		{"a\fb", 2},
		{"a\rb", 2},
		{"a\r\nb", 2},
		{"a bc", 2},
		{"a bc\n\n", 2},
		{"a b c", 3},
		{"a b\nc", 3},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("[%d] input: %s, expected: %d", i, testCase.input, testCase.expected), func(t *testing.T) {
			actual := countWords(testCase.input)
			assert.Equal(t, testCase.expected, actual)
		})

	}
}
