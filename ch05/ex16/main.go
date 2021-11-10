package ex16

import "strings"

func join(sep string, values ...string) string {
	if len(values) == 0 {
		return ""
	}
	return strings.Join(values, sep)
}
