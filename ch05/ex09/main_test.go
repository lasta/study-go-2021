package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_expand(t *testing.T) {

	testCases := []struct {
		name     string
		text     string
		f        func(string) string
		expected string
	}{
		{
			name: "空文字の場合, 空文字が返却される",
			text: "",
			f: func(s string) string {
				return fmt.Sprintf("f(%s)", s)
			},
			expected: "",
		},
		{
			name: "文字列が $ のみ場合, 変換されない",
			text: "$",
			f: func(s string) string {
				return fmt.Sprintf("f(%s)", s)
			},
			expected: "$",
		},
		{
			name: "$ 単体の場合, 変換されない",
			text: "$ foo bar foo $ buz",
			f: func(s string) string {
				return fmt.Sprintf("f(%s)", s)
			},
			expected: "$ foo bar foo $ buz",
		},
		{
			name:     "変換関数が nil の場合, 指定した文字列がそのまま返却される",
			text:     "$foo",
			f:        nil,
			expected: "$foo",
		},
		{
			name: "文字列が $foo を含む場合, 変換関数で指定された形式に置換される",
			text: "$foo",
			f: func(s string) string {
				return fmt.Sprintf("f(%s)", s)
			},
			expected: "f(foo)",
		},
		{
			name: "文字列が $foo を複数含む場合, 変換関数で指定された形式にそれぞれ置換される",
			text: "$foo bar foo $buz",
			f: func(s string) string {
				return fmt.Sprintf("f(%s)", s)
			},
			expected: "f(foo) bar foo f(buz)",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := expand(testCase.text, testCase.f)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
