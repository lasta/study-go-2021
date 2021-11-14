package countwriter

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	testCases := []struct {
		name          string
		inputs        []string
		expectedCount int64
	}{
		{
			name:          "空文字の場合, 0",
			inputs:        []string{},
			expectedCount: 0,
		},
		{
			name:          "1文字 (ASCII) の場合, 1",
			inputs:        []string{"a"},
			expectedCount: 1,
		},
		{
			name:          "1文字 (日本語) の場合, 3",
			inputs:        []string{"あ"},
			expectedCount: 3,
		},
		{
			name:          "2文字 (ASCII) 1単語の場合, 2",
			inputs:        []string{"ab"},
			expectedCount: 2,
		},
		{
			name:          "1文字 (ASCII) 2単語の場合, 2",
			inputs:        []string{"a", "b"},
			expectedCount: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var buf bytes.Buffer
			countingWriter, count := CountingWriter(&buf)

			assert.Equal(t, int64(0), *count)

			for _, input := range testCase.inputs {
				io.WriteString(countingWriter, input)
			}

			assert.Equal(t, testCase.expectedCount, *count)
			assert.Equal(t, strings.Join(testCase.inputs, ""), buf.String())
		})
	}

}
