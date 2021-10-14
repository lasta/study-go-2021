package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("different bits: %d\n", CountDifferentBytes(c1, c2))
}

func CountDifferentBytes(a, b [32]byte) (count int) {
	if a == b {
		return 0
	}

	for i := 0; i < 32; i++ {
		count += int(pc[a[i]^b[i]])
	}

	return
}
