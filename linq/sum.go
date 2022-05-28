package linq

import (
	"golang.org/x/exp/constraints"
)

type summable interface {
	constraints.Integer |
	constraints.Float | 
	constraints.Complex
}

func Sum[T summable](s []T) T {
	sum := *new(T)

	for _, n := range s {
		sum += n
	}

	return sum
}
