package ex12

import "sort"

func isAnagram(s1, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}

	b1 := []rune(s1)
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
	b2 := []rune(s2)
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	for i, r1 := range b1 {
		if r1 != b2[i] {
			return false
		}
	}
	return true
}
