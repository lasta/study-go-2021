package ex10

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	parts := strings.Split(s, ".")
	integerPart := parts[0]

	hasSign := strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+")
	absIntegerPart := integerPart
	if hasSign {
		absIntegerPart = integerPart[1:]
	}

	if len(absIntegerPart) <= 3 {
		return s
	}

	var buf bytes.Buffer
	n := len(absIntegerPart)
	shifter := n % 3

	// sign
	if hasSign {
		buf.WriteString(string(s[0]))
	}

	// integer part
	for i, c := range absIntegerPart {
		if i != 0 && i % 3 == shifter {
			buf.WriteString(",")
		}
		buf.WriteString(string(c))
	}

	// decimal part
	if len(parts) == 2 {
		buf.WriteString(".")
		buf.WriteString(parts[1])
	}

	return buf.String()
}
