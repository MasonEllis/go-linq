package linq

func FindFirst[T interface{}](s []T, f func(a T) bool) *T {
	for _, a := range s {
		if f(a) {
			return &a
		}
	}

	return nil
}