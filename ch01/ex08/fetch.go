package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stdout, "No arguments passed.")
		return
	}

	urls := AppendProtocol(args)

	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func AppendProtocol(inputs []string) []string {
	outputs := make([]string, 0, len(inputs))
	for _, input := range inputs {
		if strings.HasPrefix(input, "http://") {
			outputs = append(outputs, input)
			continue
		}
		if strings.HasPrefix(input, "https://") {
			outputs = append(outputs, input)
			continue
		}
		outputs = append(outputs, "http://"+input)
	}
	return outputs
}
