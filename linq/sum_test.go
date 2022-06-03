package linq_test

import (
	"testing"

	"github.com/MasonEllis/go-linq/linq"
	"github.com/stretchr/testify/assert"
)

func Test_Sum_Int(t *testing.T) {
	testCases := []struct {
		name        string
		s           []int
		expectedSum int
	}{
		{
			name:        "Int, Empty Slice, Returns 0",
			s:           make([]int, 0),
			expectedSum: 0,
		},
		{
			name:        "Int, Slice With One Element, Returns Element",
			s:           []int{5},
			expectedSum: 5,
		},
		{
			name:        "Int, Slice With Many Element, Returns Correct Sum",
			s:           []int{1, 2, 3, 4, 5},
			expectedSum: 15,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum := linq.Sum(testCase.s)

			assert.Equal(t, testCase.expectedSum, sum)
		})
	}
}

func Test_Sum_Float(t *testing.T) {
	testCases := []struct {
		name        string
		s           []float64
		expectedSum float64
	}{
		{
			name:        "Float, Empty Slice, Returns 0",
			s:           make([]float64, 0),
			expectedSum: 0.0,
		},
		{
			name:        "Float, Slice With One Element, Returns Element",
			s:           []float64{5.0},
			expectedSum: 5.0,
		},
		{
			name:        "Float, Slice With Many Element, Returns Correct Sum",
			s:           []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			expectedSum: 16.5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum := linq.Sum(testCase.s)

			assert.Equal(t, testCase.expectedSum, sum)
		})
	}
}

func Test_Sum_Complex(t *testing.T) {
	testCases := []struct {
		name        string
		s           []complex128
		expectedSum complex128
	}{
		{
			name:        "Complex, Empty Slice, Returns 0",
			s:           make([]complex128, 0),
			expectedSum: 0.0,
		},
		{
			name:        "Complex, Slice With One Element, Returns Element",
			s:           []complex128{5 + 2i},
			expectedSum: 5 + 2i,
		},
		{
			name:        "Complex, Slice With Many Element, Returns Correct Sum",
			s:           []complex128{1 + 2i, 3 + 4i, 5 + 6i},
			expectedSum: 9 + 12i,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sum := linq.Sum(testCase.s)

			assert.Equal(t, testCase.expectedSum, sum)
		})
	}
}
