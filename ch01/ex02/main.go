package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	const delimiter = ": "
	for index, arg := range os.Args {
		s += sep + strconv.Itoa(index) + delimiter + arg
		sep = "\n"
	}
	fmt.Println(s)
}
