package main

import (
	"fmt"
	"strings"
	"time"
)

// 実行時間の計測結果等は README.md に記載しました
func main() {
	joinTo := make([]string, 10, 1000)
	for i := 0; i < 1000; i++ {
		joinTo = append(joinTo, "a")
	}

	for n := 0; n < 10; n++ {
		fmt.Println("[BEGIN]: inefficient joiner")
		start := time.Now()
		for i := 0; i < 10000; i++ {
			joinInefficiently(joinTo)
		}
		elapsedTime1 := time.Since(start)
		fmt.Printf("[END]: inefficient joiner. elapsed : %.2f [s]\n", elapsedTime1.Seconds())

		fmt.Println("[BEGIN]: efficient joiner")
		start = time.Now()
		for i := 0; i < 10000; i++ {
			joinEfficiently(joinTo)
		}
		elapsedTime2 := time.Since(start)
		fmt.Printf("[END]: efficient joiner. elapsed : %.2f [s]\n", elapsedTime2.Seconds())

		fmt.Printf("Elapsed time 1 - Elapsed time 2 = %.2f [s]\n", (elapsedTime1 - elapsedTime2).Seconds())
	}
}

func joinInefficiently(values []string) string {
	var joined string
	var separator string
	for i := 0; i < len(values); i++ {
		joined += separator + values[i]
		separator = " "
	}

	return joined
}

func joinEfficiently(values []string) string {
	return strings.Join(values, " ")
}
