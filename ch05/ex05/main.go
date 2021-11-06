package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("[Usage]: go run main.go URL")
	}
	words, images, err := CountWordsAndImages(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Printf("words: %d, images: %d\n", words, images)
}

// CountWordsAndImages does GET request to url, fetch its content,
// and count words and images.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(node *html.Node) (words, images int) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == "img" {
		images++
	}
	if node.Type == html.TextNode {
		words = countWords(node.Data)
	}

	wordsInFirstChild, imagesInFirstChild := countWordsAndImages(node.FirstChild)
	wordsInSibling, imagesInSibling := countWordsAndImages(node.NextSibling)

	words += wordsInFirstChild + wordsInSibling
	images += imagesInFirstChild + imagesInSibling
	return
}

func countWords(data string) (words int) {
	reader := bufio.NewScanner(strings.NewReader(data))
	reader.Split(bufio.ScanWords)
	for reader.Scan() {
		words++
	}
	return words
}
