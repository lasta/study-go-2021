package ex10

import "bytes"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	shifter := n % 3

	for i, c := range s {
		if i != 0 && i % 3 == shifter {
			buf.WriteString(",")
		}
		buf.WriteString(string(c))
	}
	return buf.String()
}
