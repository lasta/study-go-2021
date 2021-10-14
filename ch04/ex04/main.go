package main

func rotate(l []int, n int) {
	if len(l) == 0 {
		return
	}
	n = n % len(l)

	swap := make([]int, n)
	copy(swap, l[:n])
	copy(l, l[n:])
	copy(l[len(l)-n:], swap)
}
