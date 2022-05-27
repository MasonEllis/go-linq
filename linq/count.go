package linq 

func Count[T interface{}](s []T, f func(a T) bool) int {
	count := 0

	for _, a := range s {
		if f(a) {
			count++
		}
	}

	return count
}