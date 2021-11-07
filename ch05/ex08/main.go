package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	var id = "id"
	for _, url := range os.Args[1:] {
		doc, err := FetchDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		foundNode := ElementByID(doc, id)

		if foundNode == nil {
			fmt.Printf("id=%s is not found", id)
			continue
		}
		fmt.Printf("id=%s is found. node: %v", id, foundNode)
	}
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
	node, ok := forEachNode(doc, pre, nil, id)
	if ok {
		return node
	}
	return nil
}

func pre(n *html.Node, id string) (foundNode *html.Node, found bool) {
	if n.Type != html.ElementNode {
		return n, false
	}
	for _, attribute := range n.Attr {
		if attribute.Key != "id" {
			continue
		}
		if attribute.Val != id {
			return n, false
		}
		return n, true
	}
	return n, false
}

func forEachNode(
	n *html.Node,
	pre func(n *html.Node, id string) (foundNode *html.Node, found bool),
	post func(n *html.Node, id string) (foundNode *html.Node, found bool),
	id string,
) (node *html.Node, found bool) {
	if pre != nil {
		node, found = pre(n, id)
		if !found {
			return node, found
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node, found = forEachNode(c, pre, post, id)
		if !found {
			return node, found
		}
	}

	if post != nil {
		node, found = post(n, id)
		if !found {
			return node, found
		}
	}

	return node, found
}


