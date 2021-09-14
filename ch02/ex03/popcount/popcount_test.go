package popcount

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func generateNumberPops1Bit() []uint64 {
	var numbers []uint64
	for i := 0; i < 64; i++ {
		numbers = append(numbers, uint64(math.Pow(2, float64(i))))
	}
	return numbers
}

func TestPopCountOrigin(t *testing.T) {
	for _, n := range generateNumberPops1Bit() {
		actual := PopCountOrigin(n)
		t.Run(
			fmt.Sprintf("Test: %d", n),
			func(t *testing.T) {
				assert.Equal(t, 1, actual)
			},
		)
	}
}

func TestPopCount(t *testing.T) {
	for _, n := range generateNumberPops1Bit() {
		actual := PopCount(n)
		t.Run(
			fmt.Sprintf("Test: %d", n),
			func(t *testing.T) {
				assert.Equal(t, 1, actual)
			},
		)
	}
}

