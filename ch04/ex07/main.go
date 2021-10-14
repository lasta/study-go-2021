package ex07

import "unicode/utf8"

func reverseStr(str []byte) []byte {
	if len(str) == 0 {
		return str
	}

	_, size := utf8.DecodeRune(str)

	reverse(str[size:])
	reverse(str[:size])
	reverse(str)
	return reverseStr(str[:len(str)-size])
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
