package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniq(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "When given empty slice then do nothing.",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "When given 1 element slice then do nothing.",
			input: []string{"a"},
			want:  []string{"a"},
		},
		{
			name:  "When given same 2 elements then uniq.",
			input: []string{"a", "a"},
			want:  []string{"a"},
		},
		{
			name:  "When given different 2 elements then not uniq.",
			input: []string{"a", "b"},
			want:  []string{"a", "b"},
		},
		{
			name:  "When given [a, b, a] then not uniq.",
			input: []string{"a", "b", "a"},
			want:  []string{"a", "b", "a"},
		},
		{
			name:  "When given [a, b, c] then not uniq.",
			input: []string{"a", "b", "c"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "When given [a, a, b] then uniq.",
			input: []string{"a", "a", "b"},
			want:  []string{"a", "b"},
		},
		{
			name:  "When given [a, b, b] then uniq.",
			input: []string{"a", "b", "b"},
			want:  []string{"a", "b"},
		},
		{
			name:  "When given [a, a, a] then uniq.",
			input: []string{"a", "a", "a"},
			want:  []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := deduplicate(tt.input)
			assert.Equal(t, tt.want, actual)
		})
	}
}
