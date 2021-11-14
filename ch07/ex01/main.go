package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (count int, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (count int, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
