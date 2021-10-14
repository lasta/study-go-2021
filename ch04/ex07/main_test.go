package ex07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	tests := []struct {
		input string
		want string
	}{
		{
			input: "",
			want: "",
		},
		{
			input: "a",
			want: "a",
		},
		{
			input: "ab",
			want: "ba",
		},
		{
			input: "abc",
			want: "cba",
		},
		{
			input: "あ",
			want: "あ",
		},
		{
			input: "あい",
			want: "いあ",
		},
		{
			input: "あいう",
			want: "ういあ",
		},
		{
			input: "あいu",
			want: "uいあ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual := []byte(tt.input)
			reverseStr(actual)
			assert.Equal(t, tt.want, string(actual))
		})
	}
}
