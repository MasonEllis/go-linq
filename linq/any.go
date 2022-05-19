package linq

func Any[T interface{}](s []T, f func(a T) bool) bool {
	for _, a := range s {
		if f(a) {
			return true
		}
	}

	return false
}