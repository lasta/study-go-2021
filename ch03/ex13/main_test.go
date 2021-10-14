package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ListIEC600272(t *testing.T) {
	assert.Equal(t, KiB, 1024)
	assert.Equal(t, MiB, 1048576)
	assert.Equal(t, GiB, 1073741824)
	assert.Equal(t, TiB, 1099511627776)
	assert.Equal(t, PiB, 1125899906842624)
	assert.Equal(t, EiB, 1152921504606846976)
	// cannot test with `assert.Equal` because of overflowing (int64)
	assert.True(t, ZiB == 1180591620717411303424)
	assert.True(t, YiB == 1208925819614629174706176)
}

func Test_ListSIPrefix(t *testing.T) {
	assert.Equal(t, KB, 1e3)
	assert.Equal(t, MB, 1e6)
	assert.Equal(t, GB, 1e9)
	assert.Equal(t, TB, 1e12)
	assert.Equal(t, PB, 1e15)
	assert.Equal(t, EB, 1e18)
	assert.Equal(t, ZB , 1e21)
	assert.Equal(t, YB , 1e24)
}
