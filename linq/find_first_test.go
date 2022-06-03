package linq_test

import (
	"testing"

	"github.com/MasonEllis/go-linq/linq"
	"github.com/stretchr/testify/assert"
)

func GetPointer[T interface{}](a T) *T {
	return &a
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
				return n%2 == 0
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

func Test_FindFirst_Float(t *testing.T) {
	testCases := []struct {
		name string
		s    []float64
		f    func(n float64) bool
		want *float64
	}{
		{
			name: "Float, No Matching Item, Returns Nil",
			s: []float64{
				1.0, 2.0, 3.0,
			},
			f: func(x float64) bool {
				return x == 0.0
			},
			want: nil,
		},
		{
			name: "Float, Starts With Matching Item, Returns First Item",
			s: []float64{
				1.0, 2.0, 3.0,
			},
			f: func(x float64) bool {
				return x == 1.0
			},
			want: GetPointer(1.0),
		},
		{
			name: "Float, First Matching Item is Last, Returns Last Item",
			s: []float64{
				1.0, 2.0, 3.0,
			},
			f: func(x float64) bool {
				return x == 3.0
			},
			want: GetPointer(3.0),
		},
		{
			name: "Float, Multiple Matching Items, Returns First",
			s: []float64{
				1.0, 2.0, 3.0, 4.0, 5.0, 6.0,
			},
			f: func(x float64) bool {
				return x == 2.0 || x == 4.0 || x == 6.0
			},
			want: GetPointer(2.0),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.FindFirst(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_FindFirst_String(t *testing.T) {
	testCases := []struct {
		name string
		s    []string
		f    func(n string) bool
		want *string
	}{
		{
			name: "String, No Matching Item, Returns Nil",
			s: []string{
				"a", "b", "c",
			},
			f: func(s string) bool {
				return s == ""
			},
			want: nil,
		},
		{
			name: "String, Starts With Matching Item, Returns First Item",
			s: []string{
				"a", "b", "c",
			},
			f: func(s string) bool {
				return s == "a"
			},
			want: GetPointer("a"),
		},
		{
			name: "String, First Matching Item is Last, Returns Last Item",
			s: []string{
				"a", "b", "c",
			},
			f: func(s string) bool {
				return s == "c"
			},
			want: GetPointer("c"),
		},
		{
			name: "String, Multiple Matching Items, Returns First",
			s: []string{
				"a", "b", "c", "d", "e", "f",
			},
			f: func(s string) bool {
				return s == "b" || s == "d" || s == "f"
			},
			want: GetPointer("b"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.FindFirst(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_FindFirst_Struct(t *testing.T) {
	type someObject struct {
		id int
	}

	testCases := []struct {
		name string
		s    []someObject
		f    func(n someObject) bool
		want *someObject
	}{
		{
			name: "Struct, No Matching Item, Returns Nil",
			s: []someObject{
				{1},
				{2},
				{3},
			},
			f: func(s someObject) bool {
				return s.id == 0
			},
			want: nil,
		},
		{
			name: "Struct, Starts With Matching Item, Returns First Item",
			s: []someObject{
				{1},
				{2},
				{3},
			},
			f: func(s someObject) bool {
				return s.id == 1
			},
			want: GetPointer(someObject{1}),
		},
		{
			name: "Struct, First Matching Item is Last, Returns Last Item",
			s: []someObject{
				{1},
				{2},
				{3},
			},
			f: func(s someObject) bool {
				return s.id == 3
			},
			want: GetPointer(someObject{3}),
		},
		{
			name: "Struct, Multiple Matching Items, Returns First",
			s: []someObject{
				{1},
				{2},
				{3},
				{4},
				{5},
				{6},
			},
			f: func(s someObject) bool {
				return s.id == 2 || s.id == 4 || s.id == 6
			},
			want: GetPointer(someObject{2}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.FindFirst(testCase.s, testCase.f)

			assert.Equal(t, testCase.want, a)
		})
	}
}
