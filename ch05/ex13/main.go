package main

import (
	"fmt"
	"io"
	"log"
	"me.lasta/study-go-2021/ch05/ex13/links"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("[Usage]: go run main.go <URL>")
	}
	breadthFirst(crawl, os.Args[1:])
}

func ExtractHost(rawUrl string) (domain string, err error) {
	var parsedURL *url.URL
	parsedURL, err = url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}

func FilterByHost(list []string, host string) (filteredURL []string) {
	for _, elem := range list {
		elemHost, _ := ExtractHost(elem)
		if elemHost == host {
			filteredURL = append(filteredURL, elem)
		}
	}
	return
}

func crawl(url string) []string {
	host, err := ExtractHost(url)
	if err != nil {
		log.Fatalf("failed to resolve host: %v", err)
	}

	fmt.Println(url)
	download(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return FilterByHost(list, host)
}

func download(rawURL string) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to resolve domain: %s\n", rawURL)
		return
	}

	response, err := http.Get(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to fetch content: %v\n", err)
		return
	}
	defer response.Body.Close()

	currentDir := "./"  // filepath.Dir(execFilePath)
	dirPath := filepath.Join(currentDir, parsedURL.Hostname(), filepath.Dir(parsedURL.Path))
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create directory: %v\n", err)
		return
	}
	filePath := filepath.Join(currentDir, parsedURL.Hostname(), parsedURL.Path)
	fileInfo, err := os.Stat(filePath)
	if err == nil && fileInfo.IsDir() {
		filePath = filepath.Join(filePath, "index.html")
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create file: %v\n", err)
		return
	}

	io.Copy(file, response.Body)
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
