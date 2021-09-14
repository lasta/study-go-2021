package main

import (
	"fmt"
	"me.lasta/study-go-2021/ch02/ex03/popcount"
	"time"
)

func main() {
	const (
		beginNum = 0
		endNum   = 1_000_000_000
	)
	startOrigin := time.Now()
	test(beginNum, endNum, popcount.PopCountOrigin)
	exceededByOrigin := time.Since(startOrigin).Seconds()

	startIterating := time.Now()
	test(beginNum, endNum, popcount.PopCount)
	exceededByIterating := time.Since(startIterating).Seconds()

	fmt.Printf(
		"Origin: %g[s], Iterating: %g[s], difference: %g[s]\n",
		exceededByOrigin,
		exceededByIterating,
		exceededByIterating-exceededByOrigin,
	)
	// => Origin: 0.372712629[s], Iterating: 4.568632618[s], difference: 4.195919989[s]
}

func test(beginIncludedNum int, endExcludedNum int, block func(num uint64) int) {
	for i := beginIncludedNum; i < endExcludedNum; i++ {
		_ = block(uint64(i))
	}
}
