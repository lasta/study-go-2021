package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name string
		l    []int
		n    int
		want []int
	}{
		{
			name: "When given empty slice then returns as origin",
			l:    []int{},
			n:    0,
			want: []int{},
		},
		{
			name: "When given empty slice and rotate once then returns as origin",
			l:    []int{},
			n:    1,
			want: []int{},
		},
		{
			name: "When given 1 element slice and never rotate then returns as origin",
			l:    []int{1},
			n:    0,
			want: []int{1},
		},
		{
			name: "When given 1 element slice and rotate once then returns as origin",
			l:    []int{1},
			n:    1,
			want: []int{1},
		},
		{
			name: "When given 1 element slice and rotate twice then returns as origin",
			l:    []int{1},
			n:    2,
			want: []int{1},
		},
		{
			name: "When given 2 element slice and never rotate then returns as origin",
			l:    []int{1, 2},
			n:    0,
			want: []int{1, 2},
		},
		{
			name: "When given 2 element slice and rotate once then returns as reversed",
			l:    []int{1, 2},
			n:    1,
			want: []int{2, 1},
		},
		{
			name: "When given 2 element slice and rotate twice then returns as origin",
			l:    []int{1, 2},
			n:    2,
			want: []int{1, 2},
		},
		{
			name: "When given 2 element slice and rotate 3 times then returns as reversed",
			l:    []int{1, 2},
			n:    3,
			want: []int{2, 1},
		},
		{
			name: "When given 3 element slice and once then returns [2, 3, 1]",
			l:    []int{1, 2, 3},
			n:    1,
			want: []int{2, 3, 1},
		},
		{
			name: "When given 3 element slice and twice then returns [3, 1, 2]",
			l:    []int{1, 2, 3},
			n:    2,
			want: []int{3, 1, 2},
		},
		{
			name: "When given 3 element slice and 3 times then returns as origin",
			l:    []int{1, 2, 3},
			n:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "When given 3 element slice and 4 times then returns [2, 3, 1]",
			l:    []int{1, 2, 3},
			n:    4,
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.l, tt.n)
			assert.Equal(t, tt.want, tt.l)
		})
	}
}
