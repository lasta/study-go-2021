package main

import "math"

func max0(values ...int) (maxNum int) {
	if len(values) == 0 {
		return math.MinInt
	}
	maxNum = values[0]
	if len(values) == 1 {
		return
	}
	for _, value := range values[1:] {
		if maxNum < value {
			maxNum = value
		}
	}
	return
}

func max1(value0 int, values ...int) (maxNum int) {
	if len(values) == 0 {
		return value0
	}

	return max0(append(values, value0)...)
}

func min0(values ...int) (minNum int) {
	if len(values) == 0 {
		return math.MaxInt
	}
	minNum = values[0]
	if len(values) == 1 {
		return
	}
	for _, value := range values[1:] {
		if minNum > value {
			minNum = value
		}
	}
	return
}

func min1(value0 int, values ...int) (minNum int) {
	if len(values) == 0 {
		return value0
	}
	return min0(append(values, value0)...)
}
