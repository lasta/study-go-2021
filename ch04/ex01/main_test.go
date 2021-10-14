package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func generate(t *testing.T, numOfBits int) (bytes [32]byte) {
	if numOfBits < 0 {
		require.Fail(t, "numOfBits must be between 0 to 256")
	}
	if numOfBits > 256 {
		require.Fail(t, "numOfBits must be between 0 to 256")
	}

	for i := 0; i < numOfBits/8; i++ {
		if numOfBits/((i+1)*8) > 0 {
			bytes[i] = byte(math.Pow(2, 8) - 1)
			continue
		}
	}
	for i := 0; i < numOfBits%8; i++ {
		bytes[numOfBits%8] = bytes[numOfBits%8] | 2<<i
	}
	return
}

func TestCountDifferentBytes(t *testing.T) {
	type args struct {
		a [32]byte
		b [32]byte
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name:      "When given bytes are same then returns 256",
			args:      args{a: [32]byte{}, b: [32]byte{}},
			wantCount: 0,
		},
		{
			name:      "When given bytes are bit all digit then returns 256",
			args:      args{a: generate(t, 256), b: generate(t, 256)},
			wantCount: 0,
		},
		{
			name:      "When given 0 and 1 then returns 255",
			args:      args{a: generate(t, 0), b: generate(t, 1)},
			wantCount: 1,
		},
		{
			name:      "When given all black and all white then returns 256",
			args:      args{a: generate(t, 0), b: generate(t, 256)},
			wantCount: 256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("a: %x\nb: %x\n", tt.args.a, tt.args.b)
			if gotCount := CountDifferentBytes(tt.args.a, tt.args.b); gotCount != tt.wantCount {
				t.Errorf("CountDifferentBytes() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
