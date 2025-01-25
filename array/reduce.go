package array

func Reduce[T any, U any](slice []T, initial U, reducer func(U, T) U) U {
	if slice == nil {
		return initial
	}
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}
