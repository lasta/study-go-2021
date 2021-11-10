package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("[Usage]: go run main.go <URL>")
	}

	filename, n, err := fetch(os.Args[1])
	if err != nil {
		log.Fatalf("failed to fetch: %v", err)
	}

	fmt.Printf("filename: %s\n", filename)
	fmt.Printf("size: %d\n", n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}
