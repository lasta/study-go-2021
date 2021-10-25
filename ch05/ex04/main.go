package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var tagAttrMap = map[string]string{
	"a": "href",
	"img": "src",
	"script": "src",
	"link": "href", // rel = "stylesheet"
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// TODO implement
func visit(links []string, node *html.Node) []string {
	if node == nil {
		return links
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, node.FirstChild)
	return visit(links, node.NextSibling)
}
