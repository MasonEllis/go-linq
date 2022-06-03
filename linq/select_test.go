package linq_test

import (
	"strconv"
	"testing"

	"github.com/MasonEllis/go-linq/linq"
	"github.com/stretchr/testify/assert"
)

func Test_Where_IntToString(t *testing.T) {
	testCases := []struct{
		name string
		s []int
		f func(n int) string
		expected []string
	}{
		{
			name: "Int to String, Empty Slice, Returns Empty Slice",
			s: make([]int, 0),
			f: func(n int) string {
				return strconv.Itoa(n)
			},
			expected: make([]string, 0),
		},
		{
			name: "Int to String, Non Empty Slice",
			s: []int {1, 2, 3},
			f: func(n int) string {
				return strconv.Itoa(n)
			},
			expected: []string {"1", "2", "3"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := linq.Select(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, res)
		})
	}
}

func Test_Where_IntToStruct(t *testing.T) {
	type someType struct {
		id int
	}
	testCases := []struct{
		name string
		s []int
		f func(n int) someType
		expected []someType
	}{
		{
			name: "Int to Struct, Empty Slice, Returns Empty Slice",
			s: make([]int, 0),
			f: func(n int) someType {
				return someType{
					id: n,
				}
			},
			expected: make([]someType, 0),
		},
		{
			name: "Int to Struct, Non Empty Slice",
			s: []int {1, 2, 3},
			f: func(n int) someType {
				return someType{
					id: n,
				}
			},
			expected: []someType {
				{ id: 1},
				{ id: 2},
				{ id: 3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := linq.Select(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, res)
		})
	}
}

func Test_Where_StructToInt(t *testing.T) {
	type someType struct {
		id int
	}
	testCases := []struct{
		name string
		s []someType
		f func(n someType) int
		expected []int
	}{
		{
			name: "Int to Struct, Empty Slice, Returns Empty Slice",
			s: make([]someType, 0),
			f: func(s someType) int {
				return s.id
			},
			expected: make([]int, 0),
		},
		{
			name: "Int to Struct, Non Empty Slice",
			s: []someType {
				{ id: 1},
				{ id: 2},
				{ id: 3},
			},
			f: func(s someType) int {
				return s.id
			},
			expected: []int {1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := linq.Select(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, res)
		})
	}
}