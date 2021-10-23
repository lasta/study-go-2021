package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	tagsCount := map[string]int{}
	countTags(tagsCount, doc)

	fmt.Printf("%-12s%s\n", "tag", "count")
	for tag, count := range tagsCount {
		fmt.Printf("%-12s%d\n", tag, count)
	}
}

func countTags(tagsCount map[string]int, node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode {
		tagsCount[node.Data]++
	}
	countTags(tagsCount, node.FirstChild)
	countTags(tagsCount, node.NextSibling)
}
