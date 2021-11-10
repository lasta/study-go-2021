package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_p(t *testing.T) {
	expected := "panic!"
	actual := p()
	assert.Equal(t, expected, actual)
}
