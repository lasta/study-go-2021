package treesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_String(t *testing.T) {
	testCases := []struct {
		name     string
		tr       *tree
		expected string
	}{
		{
			name:     "tree が nil の場合, 空文字",
			tr:       nil,
			expected: "",
		},
		{
			name:     "tree の要素が1つの場合, その要素を出力",
			tr:       &tree{value: 1},
			expected: "1",
		},
		{
			name: "tree が left を 1 つ持つ場合, left から出力",
			tr: &tree{
				value: 5,
				left:  &tree{value: 4},
			},
			expected: "4 5",
		},
		{
			name: "tree が right を 1 つ持つ場合, 自身から出力",
			tr: &tree{
				value: 5,
				right: &tree{value: 6},
			},
			expected: "5 6",
		},
		{
			name: "tree が left と right を 1 つずつ持つ場合, left, 自身, right の順で出力",
			tr: &tree{
				value: 5,
				left:  &tree{value: 4},
				right: &tree{value: 6},
			},
			expected: "4 5 6",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.tr.String()
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
