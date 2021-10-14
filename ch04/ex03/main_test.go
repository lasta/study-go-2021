package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	want := [10]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	actual := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&actual)
	assert.Equal(t, want, actual)
}
