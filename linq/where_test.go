package linq_test

import (
	"go-linq/linq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Where_Int(t *testing.T){
	testCases := []struct {
		name string
		s []int
		f func(n int) bool
		expected []int
	}{
		{
			name: "Int, Empty Slice, Returns Empty Slice",
			s: make([]int, 0),
			f: func (n int) bool{
				return false
			},
			expected: make([]int, 0),
		},
		{
			name: "Int, No Matching Items, Returns Empty Slice",
			s: []int {1, 2, 3},
			f: func (n int) bool{
				return false
			},
			expected: make([]int, 0),
		},
		{
			name: "Int, Many Matching Items, Returns Matching Items",
			s: []int {1, 2, 3, 4, 5},
			f: func (n int) bool{
				return n == 2 || n == 4
			},
			expected: []int {2, 4},
		},
		{
			name: "Int, All Matching Items, Returns All Items",
			s: []int {1, 2, 3},
			f: func (n int) bool{
				return true
			},
			expected: []int {1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			subset := linq.Where(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, subset)
		})
	}
}