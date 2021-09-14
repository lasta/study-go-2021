package popcount

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountOrigin(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))],
	)
}

func PopCount(x uint64) int {
	if x == 0 {
		return 0
	}
	for i := 1; i <= 64; i++{
		x = x & (x - 1)
		if x == 0 {
			return i
		}
	}
	return -1 // unreachable
}

