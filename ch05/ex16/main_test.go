package ex16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_join(t *testing.T) {
	testCases := []struct {
		name string
		sep string
		values []string
		expected string
	}{
		{
			name: "可変長引数が空の場合, 空文字を返却",
			sep: ", ",
			values: []string{},
			expected: "",
		},
		{
			name: "可変長引数が1つの場合, 渡された文字列をそのまま返却",
			sep: ", ",
			values: []string{"foo"},
			expected: "foo",
		},
		{
			name: "可変長引数が2つの場合, sep で結合して返却",
			sep: ", ",
			values: []string{"foo", "bar"},
			expected: "foo, bar",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := join(testCase.sep, testCase.values...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
