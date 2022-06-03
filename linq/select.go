package linq 

func Select[T, U any](s []T, f func(a T) U) []U {
	res := make([]U, len(s))

	for i, a := range s {
		b := f(a)
		res[i] = b
	}
	
	return res
}