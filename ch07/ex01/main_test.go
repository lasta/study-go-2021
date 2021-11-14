package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteCounter_Write(t *testing.T) {
	testCase := struct {
		input         string
		expectedCount int
		expectedError error
	}{
		input:         "hello",
		expectedCount: 5,
		expectedError: nil,
	}

	var c ByteCounter
	writtenCount, err := c.Write([]byte(testCase.input))
	assert.Equal(t, testCase.expectedCount, writtenCount)
	assert.Nil(t, err)

	actualCount := int(c)
	assert.Equal(t, testCase.expectedCount, actualCount)
}

func TestWordCounter_Write(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expected int
	}{
		{
			name: "空文字の場合, 0",
			input: "",
			expected: 0,
		},
		{
			name: "1文字 ` ` の場合, 0",
			input: " ",
			expected: 0,
		},
		{
			name: "1文字 `a` の場合, 1",
			input: "a",
			expected: 1,
		},
		{
			name: "2文字 `  ` の場合, 0",
			input: "  ",
			expected: 0,
		},
		{
			name: "2文字 `a ` の場合, 1",
			input: "a ",
			expected: 1,
		},
		{
			name: "2文字 ` b` の場合, 1",
			input: " b",
			expected: 1,
		},
		{
			name: "2文字 `ab` の場合, 1",
			input: "ab",
			expected: 1,
		},
		{
			name: "3文字 `   ` の場合, 0",
			input: "   ",
			expected: 0,
		},
		{
			name: "3文字 `a  ` の場合, 1",
			input: "a  ",
			expected: 1,
		},
		{
			name: "3文字 ` b ` の場合, 1",
			input: " b ",
			expected: 1,
		},
		{
			name: "3文字 `  c` の場合, 1",
			input: "  c",
			expected: 1,
		},
		{
			name: "3文字 `ab ` の場合, 1",
			input: "ab ",
			expected: 1,
		},
		{
			name: "3文字 `a c` の場合, 2",
			input: "a c",
			expected: 2,
		},
		{
			name: "3文字 ` bc` の場合, 1",
			input: " bc",
			expected: 1,
		},
		{
			name: "3文字 `abc` の場合, 1",
			input: "abc",
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var c WordCounter
			actualWrittenBytes, actualError := c.Write([]byte(testCase.input))
			assert.Equal(t, testCase.expected, actualWrittenBytes)
			assert.Nil(t, actualError)

			assert.Equal(t, testCase.expected, int(c))
		})
	}
}

func TestLineCounter_Write(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expected int
	}{
		{
			name: "空文字の場合, 0",
			input: "",
			expected: 0,
		},
		{
			name: "1文字 `\n` の場合, 0",
			input: "\n",
			expected: 1,
		},
		{
			name: "1文字 `a` の場合, 1",
			input: "a",
			expected: 1,
		},
		{
			name: "2文字 `\n\n` の場合, 2",
			input: "\n\n",
			expected: 2,
		},
		{
			name: "2文字 `a\n` の場合, 1",
			input: "a\n",
			expected: 1,
		},
		{
			name: "2文字 `\nb` の場合, 2",
			input: "\nb",
			expected: 2,
		},
		{
			name: "2文字 `ab` の場合, 1",
			input: "ab",
			expected: 1,
		},
		{
			name: "3文字 `\n\n\n` の場合, 3",
			input: "\n\n\n",
			expected: 3,
		},
		{
			name: "3文字 `a\n\n` の場合, 2",
			input: "a\n\n",
			expected: 2,
		},
		{
			name: "3文字 `\nb\n` の場合, 2",
			input: "\nb\n",
			expected: 2,
		},
		{
			name: "3文字 `\n\nc` の場合, 3",
			input: "\n\nc",
			expected: 3,
		},
		{
			name: "3文字 `ab\n` の場合, 1",
			input: "ab\n",
			expected: 1,
		},
		{
			name: "3文字 `a\nc` の場合, 2",
			input: "a\nc",
			expected: 2,
		},
		{
			name: "3文字 `\nbc` の場合, 2",
			input: "\nbc",
			expected: 2,
		},
		{
			name: "3文字 `abc` の場合, 1",
			input: "abc",
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var c LineCounter
			actualWrittenBytes, actualError := c.Write([]byte(testCase.input))
			assert.Equal(t, testCase.expected, actualWrittenBytes)
			assert.Nil(t, actualError)

			assert.Equal(t, testCase.expected, int(c))
		})
	}
}
