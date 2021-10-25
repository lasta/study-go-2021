package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

var tagAttrMap = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
	"link":   "href", // rel = "stylesheet"
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	extractUrl(os.Stdout, doc)
}

func extractUrl(writer io.Writer, node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		tag := node.Data
		if attributeName, url, extracted := extractValue(node, tag); extracted {
			io.WriteString(writer, fmt.Sprintf("%-8s%-8s%s\n", tag, attributeName, url))
		}
	}
	extractUrl(writer, node.FirstChild)
	extractUrl(writer, node.NextSibling)
}

func extractValue(node *html.Node, tag string) (attributeName, value string, extracted bool) {
	attributeName = tagAttrMap[tag]
	if attributeName == "" {
		return
	}
	for _, attribute := range node.Attr {
		if attribute.Key != attributeName {
			continue
		}
		value = attribute.Val
		if value == "" {
			return
		}
		return attributeName, value, true
	}
	return
}
