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
	if len(os.Args) != 3 {
		log.Fatal("[Usage]: go run main.go <URL> <ID value>")
	}

	url := os.Args[1]
	id := os.Args[2]
	doc, err := FetchDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	foundNode := ElementByID(doc, id)

	if foundNode == nil {
		fmt.Printf("id=%s is not found", id)
		return
	}
	fmt.Printf("id=%s is found.\n", id)
	printNode(foundNode, os.Stdout)
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

func ElementByID(doc *html.Node, id string) *html.Node {
	var foundNode *html.Node
	forEachNode(doc, func(node *html.Node) (found bool) {
		if node.Type == html.ElementNode {
			for _, attribute := range node.Attr {
				if attribute.Key != "id" {
					continue
				}
				if attribute.Val != id {
					return false
				}
				foundNode = node
				return false
			}
		}
		return true
	}, nil)
	return foundNode
}

func forEachNode(
	n *html.Node,
	pre func(n *html.Node) (found bool),
	post func(n *html.Node) (found bool),
) (found bool) {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if !forEachNode(child, pre, post) {
			return false
		}
	}

	if post != nil {
		if post != nil {
			if !post(n) {
				return false
			}
		}
	}

	return true
}
