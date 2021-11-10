package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("[Usage]: go run main.go <URL>")
	}

	doc, err := FetchDocument(os.Args[1])

	if err != nil {
		log.Fatalf("failed fetch: %v\n", err)
	}

	images := ElementsByTagName(doc, "img")
	if len(images) == 0 {
		fmt.Println("image nodes not found")
	} else {
		fmt.Println("image nodes")
		for _, image := range images {
			printNode(image, os.Stdout)
		}
	}

	fmt.Println("------------------------------------------------------------")

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4", "h5", "h6")
	if len(headings) == 0 {
		fmt.Println("headings nodes not found")
	} else {
		fmt.Println("heading nodes")
		for _, heading := range headings {
			printNode(heading, os.Stdout)
		}
	}
}

func printNode(node *html.Node, writer io.Writer) {
	if node == nil {
		return
	}

	fmt.Fprintf(writer, "<%s", node.Data)
	for _, attribute := range node.Attr {
		fmt.Fprintf(writer, " %s=%q", attribute.Key, attribute.Val)
	}
	fmt.Fprintln(writer, ">")
}

func FetchDocument(url string) (doc *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err = html.Parse(resp.Body)
	return
}

func ElementsByTagName(doc *html.Node, names ...string) (foundElements []*html.Node) {
	for _, name := range names {
		foundElements = append(foundElements, ElementByTagName(doc, name)...)
	}
	return
}

func ElementByTagName(node *html.Node, name string) (foundNodes []*html.Node)  {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == name {
		foundNodes = append(foundNodes, node)
	}
	foundNodes = append(foundNodes, ElementByTagName(node.FirstChild, name)...)
	foundNodes = append(foundNodes, ElementByTagName(node.NextSibling, name)...)
	return
}
