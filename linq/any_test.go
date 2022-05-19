package linq_test

import (
	"go-linq/linq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Any_Int(t *testing.T) {
	testCases := []struct {
		name string
		s    []int
		want bool
	}{
		{
			name: "Int, Empty Slice, Returns False",
			s:    make([]int, 0),
			want: false,
		},
		{
			name: "Int, No Valid Items, Returns False",
			s:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "Int, Valid Item at Start, Returns True",
			s:    []int{0, 1, 2, 3},
			want: true,
		},
		{
			name: "Int, Valid Item at End, Returns True",
			s:    []int{1, 2, 3, 0},
			want: true,
		},
		{
			name: "Int, Valid Item in Middle, Returns True",
			s:    []int{1, 2, 0, 3, 4},
			want: true,
		},
		{
			name: "Int, Multiple Valid Items, Returns True",
			s:    []int{1, 0, 2, 0, 3},
			want: true,
		},
	}

	f := func(n int) bool {
		return n == 0
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Any(testCase.s, f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_Any_Float64(t *testing.T) {
	testCases := []struct {
		name string
		s    []float64
		want bool
	}{
		{
			name: "Int, Empty Slice, Returns False",
			s:    make([]float64, 0),
			want: false,
		},
		{
			name: "Int, No Valid Items, Returns False",
			s:    []float64{1.0, 2.0, 3.0},
			want: false,
		},
		{
			name: "Int, Valid Item at Start, Returns True",
			s:    []float64{0.0, 1.0, 2.0, 3.0},
			want: true,
		},
		{
			name: "Int, Valid Item at End, Returns True",
			s:    []float64{1.0, 2.0, 3.0, 0.0},
			want: true,
		},
		{
			name: "Int, Valid Item in Middle, Returns True",
			s:    []float64{1.0, 2.0, 0.0, 3.0, 4.0},
			want: true,
		},
		{
			name: "Int, Multiple Valid Items, Returns True",
			s:    []float64{1.0, 0.0, 2.0, 0.0, 3.0},
			want: true,
		},
	}

	f := func(n float64) bool {
		return n == 0.0
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Any(testCase.s, f)

			assert.Equal(t, testCase.want, a)
		})
	}
}

func Test_Any_String(t *testing.T) {
	testCases := []struct {
		name string
		s    []string
		want bool
	}{
		{
			name: "Int, Empty Slice, Returns False",
			s:    make([]string, 0),
			want: false,
		},
		{
			name: "Int, No Valid Items, Returns False",
			s:    []string{"a", "b", "c"},
			want: false,
		},
		{
			name: "Int, Valid Item at Start, Returns True",
			s:    []string{"", "a", "b", "c"},
			want: true,
		},
		{
			name: "Int, Valid Item at End, Returns True",
			s:    []string{"a", "b", "c", ""},
			want: true,
		},
		{
			name: "Int, Valid Item in Middle, Returns True",
			s:    []string{"a", "b", "", "c", "d"},
			want: true,
		},
		{
			name: "Int, Multiple Valid Items, Returns True",
			s:    []string{"a", "", "b", "", "c"},
			want: true,
		},
	}

	f := func(n string) bool {
		return n == ""
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := linq.Any(testCase.s, f)

			assert.Equal(t, testCase.want, a)
		})
	}
}
