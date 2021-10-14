package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deduplicateSpaces(t *testing.T) {
	s := []byte("\t\n\v\f\r あ\t\n\v\f\r い \t\n\v\f\r ")
	assert.Equal(t, " あ い ", string(deduplicateSpaces(s)))
}
