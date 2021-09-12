package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppendProtocol(t *testing.T) {
	const (
		urlStartsWithHttp  = "http://example.com"
		urlStartsWithHttps = "https://example.com"
		urlStartsWithFtp   = "ftp://example.com"
		urlHasNoProtocol   = "example.com"
	)

	type args struct {
		inputs  []string
		outputs []string
	}
	testCases := []struct {
		name string
		args args
	}{
		// len(input) == 0
		{
			name: "When given input is empty, then returns empty slice.",
			args: args{
				inputs:  []string{},
				outputs: []string{},
			},
		},

		// len(input) == 1
		{
			name: "When given input starts with 'http://' then returns itself",
			args: args{
				inputs:  []string{urlStartsWithHttp},
				outputs: []string{urlStartsWithHttp},
			},
		},
		{
			name: "When given input starts with 'https://' then returns itself",
			args: args{
				inputs:  []string{urlStartsWithHttps},
				outputs: []string{urlStartsWithHttps},
			},
		},
		{
			name: "When given input has no protocol then append 'http://'",
			args: args{
				inputs:  []string{urlHasNoProtocol},
				outputs: []string{urlStartsWithHttp},
			},
		},
		{
			name: "When given input doesn't starts with 'http://' nor 'https://' then append 'http://'",
			args: args{
				inputs:  []string{urlStartsWithFtp},
				outputs: []string{"http://" + urlStartsWithFtp},
			},
		},

		// len(input) > 1
		{
			name: "When given input has plural urls, then append protocol if required",
			args: args{
				inputs: []string{urlHasNoProtocol, urlStartsWithHttp, urlStartsWithHttps, urlStartsWithFtp},
				outputs: []string{
					"http://" + urlHasNoProtocol,
					urlStartsWithHttp,
					urlStartsWithHttps,
					"http://" + urlStartsWithFtp,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			inputs := testCase.args.inputs
			actual := AppendProtocol(inputs)
			expected := testCase.args.outputs
			assert.Equal(t, expected, actual)
		})
	}
}
