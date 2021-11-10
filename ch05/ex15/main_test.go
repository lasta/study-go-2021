package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_max0(t *testing.T) {
	testCases := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "引数が0個の場合, int の最小値を返却",
			values:   []int{},
			expected: math.MinInt,
		},
		{
			name:     "引数が1個の場合, 渡した値を返却",
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "引数が2個で1個目の値が大きい場合, 1個目の値を返却",
			values:   []int{1, -1},
			expected: 1,
		},
		{
			name:     "引数が2個で2個目の値が大きい場合, 2個目の値を返却",
			values:   []int{-1, 1},
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := max0(testCase.values...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func Test_max1(t *testing.T) {
	testCases := []struct {
		name     string
		value0   int
		values   []int
		expected int
	}{
		{
			name:     "可変長引数が0個の場合, value0 を返却",
			value0:   1,
			values:   []int{},
			expected: 1,
		},
		{
			name:     "可変長引数が1個渡され, 可変長引数のほうが大きい場合, 可変長引数の値を返却",
			value0:   -1,
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "可変長引数が1個渡され, value0 のほうが大きい場合, value0 を返却",
			value0:   1,
			values:   []int{-1},
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := max1(testCase.value0, testCase.values...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func Test_min0(t *testing.T) {
	testCases := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "引数が0個の場合, int の最大値を返却",
			values:   []int{},
			expected: math.MaxInt,
		},
		{
			name:     "引数が1個の場合, 渡した値を返却",
			values:   []int{1},
			expected: 1,
		},
		{
			name:     "引数が2個で1個目の値が小さい場合, 1個目の値を返却",
			values:   []int{-1, 1},
			expected: -1,
		},
		{
			name:     "引数が2個で2個目の値が小さい場合, 2個目の値を返却",
			values:   []int{1, -1},
			expected: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := min0(testCase.values...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func Test_min1(t *testing.T) {
	testCases := []struct {
		name     string
		value0   int
		values   []int
		expected int
	}{
		{
			name:     "可変長引数が0個の場合, value0 を返却",
			value0:   1,
			values:   []int{},
			expected: 1,
		},
		{
			name:     "可変長引数が1個渡され, 可変長引数のほうが小さい場合, 可変長引数の値を返却",
			value0:   1,
			values:   []int{-1},
			expected: -1,
		},
		{
			name:     "可変長引数が1個渡され, value0 のほうが小さい場合, value0 を返却",
			value0:   -1,
			values:   []int{1},
			expected: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := min1(testCase.value0, testCase.values...)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
