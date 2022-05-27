package linq_test

import (
	"go-linq/linq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Count_Int(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		f    func(n int) bool
		want int
	}{
		{
			name: "Int, Empty Slice, Returns 0",
			s: make([]int, 0),
			f : func(n int) bool {
				return n == 0
			},
			want: 0,
		},
		{
			name: "Int, 1 Matching Item, Returns 1",
			s: []int {1, 2, 3},
			f : func(n int) bool {
				return n == 2
			},
			want: 1,
		},
		{
			name: "Int, Multiple Matching Item, Returns Correct Count",
			s: []int {1, 2, 3, 4, 5, 6},
			f : func(n int) bool {
				return n % 2 == 0
			},
			want: 3,
		},
		{
			name: "Int, Every Item Matches, Returns Slice Length",
			s: []int {1, 2, 3, 4, 5, 6},
			f : func(n int) bool {
				return true
			},
			want: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Count(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_Count_Float(t *testing.T) {
	testCases := []struct {
		name string
		s    []float64
		f    func(n float64) bool
		want int
	}{
		{
			name: "Float, Empty Slice, Returns 0",
			s: make([]float64, 0),
			f : func(n float64) bool {
				return n == 0.0
			},
			want: 0,
		},
		{
			name: "Float, 1 Matching Item, Returns 1",
			s: []float64 {1.0, 2.0, 3.0},
			f : func(n float64) bool {
				return n == 2.0
			},
			want: 1,
		},
		{
			name: "Float, Multiple Matching Item, Returns Correct Count",
			s: []float64 {1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			f : func(n float64) bool {
				return n == 2.0 || n == 4.0 || n == 6.0
			},
			want: 3,
		},
		{
			name: "Float, Every Item Matches, Returns Slice Length",
			s: []float64 {1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			f : func(n float64) bool {
				return true
			},
			want: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Count(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_Count_String(t *testing.T) {
	testCases := []struct {
		name string
		s    []string
		f    func(n string) bool
		want int
	}{
		{
			name: "String, Empty Slice, Returns 0",
			s: make([]string, 0),
			f : func(s string) bool {
				return s == "a"
			},
			want: 0,
		},
		{
			name: "String, 1 Matching Item, Returns 1",
			s: []string {"a", ""},
			f : func(s string) bool {
				return s == "a"
			},
			want: 1,
		},
		{
			name: "String, Multiple Matching Item, Returns Correct Count",
			s: []string {"a", "b", "c", "d", "e", "f"},
			f : func(s string) bool {
				return s == "a" || s == "c"|| s == "e"
			},
			want: 3,
		},
		{
			name: "String, Every Item Matches, Returns Slice Length",
			s: []string {"a", "b", "c", "d", "e", "f"},
			f : func(n string) bool {
				return true
			},
			want: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Count(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_Count_Struct(t *testing.T) {
	type someObject struct {
		id int
	}

	testCases := []struct {
		name string
		s    []someObject
		f    func(s someObject) bool
		want int
	}{
		{
			name: "Struct, Empty Slice, Returns 0",
			s: make([]someObject, 0),
			f : func(s someObject) bool {
				return s.id == 0
			},
			want: 0,
		},
		{
			name: "Struct, 1 Matching Item, Returns 1",
			s: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
			},
			f : func(s someObject) bool {
				return s.id == 2
			},
			want: 1,
		},
		{
			name: "Struct, Multiple Matching Item, Returns Correct Count",
			s: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
				{id: 4},
				{id: 5},
				{id: 6},
			},
			f : func(s someObject) bool {
				return s.id == 2 || s.id == 4 || s.id == 6
			},
			want: 3,
		},
		{
			name: "Struct, Every Item Matches, Returns Slice Length",
			s: []someObject{
				{id: 1},
				{id: 2},
				{id: 3},
				{id: 4},
				{id: 5},
				{id: 6},
			},
			f : func(s someObject) bool {
				return true
			},
			want: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Count(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}