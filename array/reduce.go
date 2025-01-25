package array

// Reduce applies a reducer function to each element of the slice and returns the accumulated result.
// If the input slice is nil, it returns the initial value.
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