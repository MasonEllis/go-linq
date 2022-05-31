package linq 

func Where[T any](s []T, f func(a T) bool) []T {
	res := make([]T, 0)

	for _, a := range s {
		if f(a) {
			res = append(res, a)
		}
	}

	return res
}