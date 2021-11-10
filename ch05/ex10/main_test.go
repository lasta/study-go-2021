package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findIndex(l []string, s string) int {
	for i, elem := range l {
		if elem == s {
			return i
		}
	}

	return -1
}

func Test_topoSort(t *testing.T) {
	const (
		intro = "intro to programming"
		discreteMath = "discrete math"
		dataStructure = "data structures"
		algorithms = "algorithms"
		formalLanguages = "formal languages"
		computerOrganization = "computer organization"
		os = "operating systems"
		networks = "networks"
		linearAlgebra = "linear algebra"
		calculus = "calculus"
		compilers = "compilers"
		db = "databases"
		programmingLanguages = "programming languages"
	)

	sorted := topoSort(prereqs)

	// see: dependencies graph ./ch05.ex10.png
	assert.Less(t, findIndex(sorted, intro), findIndex(sorted, discreteMath))

	assert.Less(t, findIndex(sorted, discreteMath), findIndex(sorted, dataStructure))
	assert.Less(t, findIndex(sorted, discreteMath), findIndex(sorted, formalLanguages))

	assert.Less(t, findIndex(sorted, dataStructure), findIndex(sorted, db))
	assert.Less(t, findIndex(sorted, dataStructure), findIndex(sorted, algorithms))
	assert.Less(t, findIndex(sorted, dataStructure), findIndex(sorted, compilers))
	assert.Less(t, findIndex(sorted, dataStructure), findIndex(sorted, programmingLanguages))
	assert.Less(t, findIndex(sorted, dataStructure), findIndex(sorted, os))

	assert.Less(t, findIndex(sorted, formalLanguages), findIndex(sorted, compilers))

	assert.Less(t, findIndex(sorted, computerOrganization), findIndex(sorted, compilers))
	assert.Less(t, findIndex(sorted, computerOrganization), findIndex(sorted, programmingLanguages))
	assert.Less(t, findIndex(sorted, computerOrganization), findIndex(sorted, os))

	assert.Less(t, findIndex(sorted, os), findIndex(sorted, networks))

	assert.Less(t, findIndex(sorted, linearAlgebra), findIndex(sorted, calculus))
}
