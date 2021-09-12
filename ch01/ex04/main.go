package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineHasFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesWithFileName(f, counts, lineHasFile)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			if lineHasFile == nil {
				fmt.Printf("%d\t%s\n", n, line)
			} else {
				fileName := strings.Join(lineHasFile[line], ", ")
				fmt.Printf("%d\t%s\t%s\n", n, line, fileName)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	// FIXME: input.Err() からのエラーの可能性を無視している
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countLinesWithFileName(f *os.File, counts map[string]int, textHasFile map[string][]string) {
	input := bufio.NewScanner(f)
	// FIXME: input.Err() からのエラーの可能性を無視している
	for input.Scan() {
		text := input.Text()
		counts[text]++

		if len(textHasFile[text]) == 0 {
			textHasFile[text] = []string{f.Name()}
			continue
		}
		if Contains(textHasFile[text], f.Name()) {
			continue
		}
		textHasFile[text] = append(textHasFile[text], f.Name())
	}
}

func Contains(slice []string, value string) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}
