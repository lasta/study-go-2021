package main

func deduplicate(l []string) (deduplicated []string) {
	if len(l) <= 1 {
		return l
	}

	previous := l[0]
	deduplicated = append(deduplicated, previous)

	for _, str := range l {
		if str != previous {
			previous = str
			deduplicated = append(deduplicated, str)
		}
	}
	return
}
