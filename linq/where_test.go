package linq_test

import (
	"testing"

	"github.com/MasonEllis/go-linq/linq"
	"github.com/stretchr/testify/assert"
)

func Test_Where_Int(t *testing.T) {
	testCases := []struct {
		name     string
		s        []int
		f        func(n int) bool
		expected []int
	}{
		{
			name: "Int, Empty Slice, Returns Empty Slice",
			s:    make([]int, 0),
			f: func(n int) bool {
				return false
			},
			expected: make([]int, 0),
		},
		{
			name: "Int, No Matching Items, Returns Empty Slice",
			s:    []int{1, 2, 3},
			f: func(n int) bool {
				return false
			},
			expected: make([]int, 0),
		},
		{
			name: "Int, Many Matching Items, Returns Matching Items",
			s:    []int{1, 2, 3, 4, 5},
			f: func(n int) bool {
				return n == 2 || n == 4
			},
			expected: []int{2, 4},
		},
		{
			name: "Int, All Matching Items, Returns All Items",
			s:    []int{1, 2, 3},
			f: func(n int) bool {
				return true
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			subset := linq.Where(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, subset)
		})
	}
}

func Test_Where_Float(t *testing.T) {
	testCases := []struct {
		name     string
		s        []float64
		f        func(x float64) bool
		expected []float64
	}{
		{
			name: "Float, Empty Slice, Returns Empty Slice",
			s:    make([]float64, 0),
			f: func(x float64) bool {
				return false
			},
			expected: make([]float64, 0),
		},
		{
			name: "Float, No Matching Items, Returns Empty Slice",
			s:    []float64{1.0, 2.0, 3.0},
			f: func(x float64) bool {
				return false
			},
			expected: make([]float64, 0),
		},
		{
			name: "Float, Many Matching Items, Returns Matching Items",
			s:    []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			f: func(x float64) bool {
				return x == 2.0 || x == 4.0
			},
			expected: []float64{2.0, 4.0},
		},
		{
			name: "Float, All Matching Items, Returns All Items",
			s:    []float64{1.0, 2.0, 3.0},
			f: func(x float64) bool {
				return true
			},
			expected: []float64{1.0, 2.0, 3.0},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			subset := linq.Where(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, subset)
		})
	}
}

func Test_Where_String(t *testing.T) {
	testCases := []struct {
		name     string
		s        []string
		f        func(s string) bool
		expected []string
	}{
		{
			name: "String, Empty Slice, Returns Empty Slice",
			s:    make([]string, 0),
			f: func(s string) bool {
				return false
			},
			expected: make([]string, 0),
		},
		{
			name: "String, No Matching Items, Returns Empty Slice",
			s:    []string{"a", "b", "c"},
			f: func(s string) bool {
				return false
			},
			expected: make([]string, 0),
		},
		{
			name: "String, Many Matching Items, Returns Matching Items",
			s:    []string{"a", "b", "c", "d", "e"},
			f: func(s string) bool {
				return s == "b" || s == "d"
			},
			expected: []string{"b", "d"},
		},
		{
			name: "String, All Matching Items, Returns All Items",
			s:    []string{"a", "b", "c"},
			f: func(s string) bool {
				return true
			},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			subset := linq.Where(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, subset)
		})
	}
}

func Test_Where_Struct(t *testing.T) {
	type someObject struct {
		id int
	}

	testCases := []struct {
		name     string
		s        []someObject
		f        func(s someObject) bool
		expected []someObject
	}{
		{
			name: "Struct, Empty Slice, Returns Empty Slice",
			s:    make([]someObject, 0),
			f: func(s someObject) bool {
				return false
			},
			expected: make([]someObject, 0),
		},
		{
			name: "Struct, No Matching Items, Returns Empty Slice",
			s: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			f: func(s someObject) bool {
				return false
			},
			expected: make([]someObject, 0),
		},
		{
			name: "Struct, Many Matching Items, Returns Matching Items",
			s: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
				{id: 4},
				{id: 5},
			},
			f: func(s someObject) bool {
				return s.id == 2 || s.id == 4
			},
			expected: []someObject{
				{id: 2},
				{id: 4},
			},
		},
		{
			name: "Struct, All Matching Items, Returns All Items",
			s: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			f: func(s someObject) bool {
				return true
			},
			expected: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			subset := linq.Where(testCase.s, testCase.f)

			assert.Equal(t, testCase.expected, subset)
		})
	}
}
