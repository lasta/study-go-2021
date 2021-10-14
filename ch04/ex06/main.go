package main

import (
	"unicode"
	"unicode/utf8"
)

func deduplicateSpaces(str []byte) []byte {
	if len(str) < 2 {
		return str
	}

	isLastSpace := false
	readingIndex := 0
	writingIndex := 0

	for readingIndex < len(str) {
		letter, letterLength := utf8.DecodeRune(str[readingIndex:])
		if !unicode.IsSpace(letter) {
			isLastSpace = false
			utf8.EncodeRune(str[writingIndex:], letter)
			writingIndex += letterLength
			readingIndex += letterLength
			continue
		}

		if !isLastSpace {
			letterLength = utf8.EncodeRune(str[writingIndex:], ' ')
			writingIndex += letterLength
		}
		isLastSpace = true
		readingIndex += letterLength
	}
	return str[:writingIndex]
}
