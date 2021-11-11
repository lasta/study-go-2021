package intset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewIntSet(elements ...int) *IntSet {
	obj := &IntSet{}
	for _, element := range elements {
		obj.Add(element)
	}
	return obj
}

func (s *IntSet) Equals(other *IntSet) bool {
	if s == nil && other == nil {
		return true
	}
	if s == nil {
		return false
	}
	if other == nil {
		return false
	}

	var shorterLength int
	if len(s.words) < len(other.words) {
		shorterLength = len(s.words)
	} else {
		shorterLength = len(other.words)
	}

	for i := 0; i < shorterLength; i++ {
		if s.words[i] != other.words[i] {
			return false
		}
	}

	if len(s.words) > shorterLength {
		for i := shorterLength; i < len(s.words); i++ {
			if s.words[i] != 0 {
				return false
			}
		}
	}

	if len(other.words) > shorterLength {
		for i := shorterLength; i < len(other.words); i++ {
			if other.words[i] != 0 {
				return false
			}
		}
	}

	return true
}

func TestIntSet_Len(t *testing.T) {
	testCases := []struct {
		name     string
		s        *IntSet
		expected int
	}{
		{
			name:     "空の場合, 長さ0",
			s:        NewIntSet(),
			expected: 0,
		},
		{
			name:     "1 のみ渡された場合, 長さ1",
			s:        NewIntSet(1),
			expected: 1,
		},
		{
			name:     "2 のみ渡された場合, 長さ1",
			s:        NewIntSet(2),
			expected: 1,
		},
		{
			name:     "1, 2 が渡された場合, 長さ2",
			s:        NewIntSet(1, 2),
			expected: 2,
		},
		{
			name:     "1, 64 が渡された場合, 長さ2",
			s:        NewIntSet(1, 64),
			expected: 2,
		},
		{
			name:     "1, 63, 64 が渡された場合, 長さ3",
			s:        NewIntSet(1, 63, 64),
			expected: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.s.Len()
			assert.Equal(t, testCase.expected, actual)
		})
	}
}

func TestIntSet_Remove(t *testing.T) {
	testCases := []struct {
		name     string
		s        *IntSet
		x        int
		expected *IntSet
	}{
		{
			name:     "空の IntSet から要素を削除した場合, 空のまま変化しない",
			s:        NewIntSet(),
			x:        1,
			expected: NewIntSet(),
		},
		{
			name:     "要素が 1 である IntSet から 1 を削除した場合, 空の IntSet になる",
			s:        NewIntSet(1),
			x:        1,
			expected: NewIntSet(),
		},
		{
			name:     "要素が 1 である IntSet から 2 を削除した場合, 削除されない",
			s:        NewIntSet(1),
			x:        2,
			expected: NewIntSet(1),
		},
		{
			name:     "要素が 1, 64 である IntSet から 128 を削除した場合, 削除されない",
			s:        NewIntSet(1, 64),
			x:        128,
			expected: NewIntSet(1, 64),
		},
		{
			name:     "要素が 1, 64 である IntSet から 64 を削除した場合, 64 が削除される",
			s:        NewIntSet(1, 64),
			x:        64,
			expected: NewIntSet(1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.s.Remove(testCase.x)
			assert.True(
				t,
				testCase.s.Equals(testCase.expected),
				"expected %v, but got %v",
				testCase.expected,
				testCase.s,
			)
		})
	}
}

func TestIntSet_Clear(t *testing.T) {
	testCases := []struct {
		name string
		s    *IntSet
	}{
		{
			name: "空の IntSet を clear できる",
			s:    NewIntSet(),
		},
		{
			name: "要素が1 (words の長さ1) の IntSet を clear できる",
			s:    NewIntSet(1),
		},
		{
			name: "要素が63 (words の長さ1) の IntSet を clear できる",
			s:    NewIntSet(63),
		},
		{
			name: "要素が64 (words の長さ2) の IntSet を clear できる",
			s:    NewIntSet(64),
		},
		{
			name: "要素が1, 64 (words の長さ2) の IntSet を clear できる",
			s:    NewIntSet(1, 64),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.s.Clear()
			assert.True(t,
				testCase.s.Equals(NewIntSet()),
				"empty IntSet is expected, but got %v",
				testCase.s,
			)
		})
	}
}

func TestIntSet_Copy(t *testing.T) {
	srcIntSet := NewIntSet(1, 64, 128)
	destIntSet := srcIntSet.Copy()

	// asserts same elements
	assert.True(
		t,
		srcIntSet.Equals(destIntSet),
		"srcIntSet is %v, but destIntSet is %v",
		srcIntSet,
		destIntSet,
	)

	// asserts not same instances
	assert.False(t,
		srcIntSet == destIntSet,
		"srcIntSet address is %p, but destIntSet address is %p",
		srcIntSet,
		destIntSet,
	)
}

func TestIntSet_AddAll(t *testing.T) {
	testCases := []struct {
		name     string
		s        *IntSet
		xs       []int
		expected *IntSet
	}{
		{
			name:     "空の IntSet に 1 つ要素を追加できる",
			s:        NewIntSet(),
			xs:       []int{1},
			expected: NewIntSet(1),
		},
		{
			name:     "空の IntSet に 2 つ要素を追加できる",
			s:        NewIntSet(),
			xs:       []int{1, 64},
			expected: NewIntSet(1, 64),
		},
		{
			name:     "要素を 1 つ持つの IntSet に 1 つ要素を追加できる",
			s:        NewIntSet(1),
			xs:       []int{64},
			expected: NewIntSet(1, 64),
		},
		{
			name:     "要素を 2 つ持つの IntSet に 2 つ要素を追加できる",
			s:        NewIntSet(1, 63),
			xs:       []int{64, 127},
			expected: NewIntSet(1, 63, 64, 127),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.s.AddAll(testCase.xs...)
			assert.True(t,
				testCase.s.Equals(testCase.expected),
				"expected: %v, but got %v",
				testCase.expected,
				testCase.s,
			)
		})
	}
}
