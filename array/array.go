package array

func Filter[T any](slice []T, condition func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[Input any, Output any](slice []Input, transform func(Input) Output) []Output {
	var result []Output
	for _, v := range slice {
		result = append(result, transform(v))
	}
	return result
}

func Reduce[T any, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

func Find[T any](slice []T, condition func(T) bool) *T {
	for _, v := range slice {
		if condition(v) {
			return &v
		}
	}
	return nil
}

func FindIndex[T any](slice []T, condition func(T) bool) int {
	for i, v := range slice {
		if condition(v) {
			return i
		}
	}
	return -1
}
