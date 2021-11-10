package main

import (
	"fmt"
	"log"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	//"linear algebra": {"calculus"}, // circular reference

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func findIndex(l []string, s string) int {
	for i, elem := range l {
		if elem == s {
			return i
		}
	}

	return -1
}

func topoSort(m map[string][]string) (order []string, err error) {
	seen := make(map[string]bool)
	var visitAll func(items []string, parents []string)

	visitAll = func(items []string, parents []string) {
		for _, item := range items {
			wasSolved, found := seen[item]

			if found && !wasSolved {
				begin := findIndex(parents, item)
				err = fmt.Errorf(
					"circular reference was found: %s",
					strings.Join(append(parents[begin:], item), " -> "),
				)
				return
			}

			if !found {
				seen[item] = false
				visitAll(m[item], append(parents, item))
				seen[item] = true
				order = append(order, item)
			}
		}
	}

	for key := range m {
		visitAll([]string{key}, nil)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}
