package array

func FindIndex[T any](slice []T, condition func(T) bool) int {
	for i, v := range slice {
		if condition(v) {
			return i
		}
	}
	return -1
}
