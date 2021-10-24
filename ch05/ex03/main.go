package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

type HtmlTag string
type TagsBlackList []HtmlTag

var tagsBlackList = TagsBlackList{"script", "style"}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	extractTextNodes(os.Stdout, doc, "\n")
}

func extractTextNodes(writer io.Writer, node *html.Node, delimiter string) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		if data := node.Data; data != "" {
			io.WriteString(writer, node.Data)
			io.WriteString(writer, delimiter)
		}
	}
	if node.Type != html.ElementNode || !tagsBlackList.contains(HtmlTag(node.Data)) {
		extractTextNodes(writer, node.FirstChild, delimiter)
	}
	extractTextNodes(writer, node.NextSibling, delimiter)
}

func (blacklist TagsBlackList) contains(tag HtmlTag) bool {
	for _, elem := range blacklist {
		if tag == elem {
			return true
		}
	}
	return false
}
