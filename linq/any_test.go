package linq_test

import (
	"go-linq/linq"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Any_Int_EmptySlice_ReturnsFalse(t *testing.T) {
	s := make([]int, 0)
	f := func(n int) bool {
		return n == 0
	}

	a := linq.Any(s, f)

	assert.False(t, a)
}

func Test_Any_Int_NoValidItems_ReturnsFalse(t *testing.T) {
	s := []int {
		1, 2, 3,
	}
	f := func(n int) bool {
		return n == 0
	}

	a := linq.Any(s, f)

	assert.False(t, a)
}

func Test_Any_Int_ValidItemAtStart_ReturnsTrue(t *testing.T) {
	s := []int {
		0, 1, 2, 3,
	}
	f := func(n int) bool {
		return n == 0
	}

	a := linq.Any(s, f)

	assert.True(t, a)
}

func Test_Any_Int_ValidItemAtEnd_ReturnsTrue(t *testing.T) {
	s := []int {
		1, 2, 3, 0,
	}
	f := func(n int) bool {
		return n == 0
	}

	a := linq.Any(s, f)

	assert.True(t, a)
}

func Test_Any_Int_ValidItemInMiddle_ReturnsTrue(t *testing.T) {
	s := []int {
		1, 2, 0, 3, 4,
	}
	f := func(n int) bool {
		return n == 0
	}

	a := linq.Any(s, f)

	assert.True(t, a)
}

func Test_Any_Int_MultipleValidItems_ReturnsTrue(t *testing.T) {
	s := []int {
		1, 0, 2, 0, 3,
	}
	f := func(n int) bool {
		return n == 0
	}

	a := linq.Any(s, f)

	assert.True(t, a)
}

