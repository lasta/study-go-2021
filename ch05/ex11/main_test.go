package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topoSort(t *testing.T) {
	const (
		intro                = "intro to programming"
		discreteMath         = "discrete math"
		dataStructure        = "data structures"
		algorithms           = "algorithms"
		formalLanguages      = "formal languages"
		computerOrganization = "computer organization"
		os                   = "operating systems"
		networks             = "networks"
		linearAlgebra        = "linear algebra"
		calculus             = "calculus"
		compilers            = "compilers"
		db                   = "databases"
		programmingLanguages = "programming languages"
	)

	sorted, _ := topoSort(prereqs)

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
func Test_topoSortCircularRef(t *testing.T) {
	var prereqsHasCircularRef = map[string][]string{
		"calculus":       {"linear algebra"},
		"linear algebra": {"calculus"},
	}
	_, err := topoSort(prereqsHasCircularRef)

	assert.NotNil(t, err)

	// 順序不定 のため assert を用いることができない
	// assert.EqualError(t, err, "circular reference was found: calculus -> linear algebra -> calculus")
	actualMessage := err.Error()
	if actualMessage != "circular reference was found: calculus -> linear algebra -> calculus" &&
		actualMessage != "circular reference was found: linear algebra -> calculus -> linear algebra" {
		t.Errorf("Error message not equal. %s nor %s",
			"circular reference was found: calculus -> linear algebra -> calculus",
			"circular reference was found: linear algebra -> calculus -> linear algebra",
		)
	}

}
