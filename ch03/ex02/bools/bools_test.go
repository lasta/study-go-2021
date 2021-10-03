package bools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	type args struct {
		inputs []bool
		output bool
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "When given input is empty then returns false",
			args: args{
				inputs: []bool{},
				output: false,
			},
		},
		{
			name: "When given array is {false} then returns false",
			args: args{
				inputs: []bool{false},
				output: false,
			},
		},
		{
			name: "When given array is {true} then returns true",
			args: args{
				inputs: []bool{true},
				output: true,
			},
		},
		{
			name: "When given array has only false then returns false",
			args: args{
				inputs: []bool{false, false},
				output: false,
			},
		},
		{
			name: "When given array is {true, false} then returns true",
			args: args{
				inputs: []bool{true, false},
				output: true,
			},
		},
		{
			name: "When given array is {false, true} then returns true",
			args: args{
				inputs: []bool{false, true},
				output: true,
			},
		},
		{
			name: "When given array has only true then returns true",
			args: args{
				inputs: []bool{true, true},
				output: true,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			inputs := testCase.args.inputs
			actual := Any(inputs)
			expected := testCase.args.output
			assert.Equal(t, expected, actual)
		})
	}
}
