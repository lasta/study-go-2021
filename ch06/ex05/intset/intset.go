package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const uintSize = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintSize, uint(x%uintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintSize, uint(x%uintSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns how many elements in s
func (s *IntSet) Len() (sum int) {
	for _, word := range s.words {
		sum += bits.OnesCount(word)
	}
	return
}

// Remove removes x in s
func (s *IntSet) Remove(x int) {
	word, bit := x/uintSize, uint(x%uintSize)

	// x is not in s
	if word >= len(s.words) {
		return
	}

	// &^ : bit clear
	s.words[word] &^= 1 << bit
}

// Clear clears all elements from s
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy returns copy of s
func (s *IntSet) Copy() *IntSet {
	srcWords := s.words
	dstWords := make([]uint, len(srcWords))
	copy(dstWords, srcWords)

	return &IntSet{dstWords}
}

// AddAll adds all given integers
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// IntersectWith sets s to the intersection of s and other
func (s *IntSet) IntersectWith(other *IntSet) {
	for i := range s.words {
		if i < len(other.words) {
			s.words[i] &= other.words[i]
			continue
		}
		s.words[i] = 0
	}
}

// DifferenceWith sets s to the difference of s to other
func (s *IntSet) DifferenceWith(other *IntSet) {
	for i := range s.words {
		if i < len(other.words) {
			s.words[i] &^= other.words[i]
		}
	}
}

// SymmetricDifference sets s to the symmetric difference of s and other
func (s *IntSet) SymmetricDifference(other *IntSet) {
	for i := range s.words {
		if i < len(other.words) {
			s.words[i] ^= other.words[i]
		}
	}
	if len(s.words) < len(other.words) {
		s.words = append(s.words, other.words[len(s.words):]...)
	}
}

// Elems returns slice contains all values
func (s *IntSet) Elems() []uint {
	l := make([]uint, 0, len(s.words))

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				l = append(l, uint(uintSize*i+j))
			}
		}
	}
	return l
}
