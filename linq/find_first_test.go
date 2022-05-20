package linq_test

import (
	"go-linq/linq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetPointer(n int) *int {
	return &n
}

func Test_FindFirst_Int(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		f    func(n int) bool
		want *int
	}{
		{
			name: "Int, No Matching Item, Returns Nil",
			s: []int{
				1, 2, 3,
			},
			f: func(n int) bool {
				return n == 0
			},
			want: nil,
		},
		{
			name: "Int, Starts With Matching Item, Returns First Item",
			s: []int{
				1, 2, 3,
			},
			f: func(n int) bool {
				return n == 1
			},
			want: GetPointer(1),
		},
		{
			name: "Int, First Matching Item is Last, Returns Last Item",
			s: []int{
				1, 2, 3,
			},
			f: func(n int) bool {
				return n == 3
			},
			want: GetPointer(3),
		},
		{
			name: "Int, Multiple Matching Items, Returns First",
			s: []int{
				1, 2, 3, 4, 5, 6,
			},
			f: func(n int) bool {
				return n % 2 == 0
			},
			want: GetPointer(2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.FindFirst(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}
