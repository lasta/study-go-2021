package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url, os.Stdout)
	}
}

func outline(url string, writer io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement, writer)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, writer io.Writer), writer io.Writer) {
	if pre != nil {
		pre(n, writer)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, writer)
	}

	if post != nil {
		post(n, writer)
	}
}

var depth int

func startElement(n *html.Node, writer io.Writer) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(writer, "%*s<%s", depth*2, "", n.Data)

		var attributes []string
		for _, attribute := range n.Attr {
			attributes = append(attributes, fmt.Sprintf("%s=%q", attribute.Key, attribute.Val))
		}
		if len(attributes) > 0 {
			fmt.Fprint(writer, " ")
			fmt.Fprintf(writer, strings.Join(attributes, " "))
		}

		if n.FirstChild == nil {
			fmt.Fprintln(writer, " />")
			return
		}
		fmt.Fprintln(writer, ">")
		depth++
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			lines := strings.Split(text, "\n")
			for _, line := range lines {
				fmt.Fprintf(writer, "%*s%s\n", depth*2, "", strings.TrimSpace(line))
			}
		}
		return
	case html.CommentNode:
		fmt.Fprintf(writer, "%*s<!-- %s -->\n", depth*2, "", n.Data)
		return
	}
}

func endElement(n *html.Node, writer io.Writer) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Fprintf(writer, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
