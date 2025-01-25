package array

func FindIndex[T any](slice []T, condition func(T) bool) int {
	if slice == nil {
		return -1
	}
	for i, v := range slice {
		if condition(v) {
			return i
		}
	}
	return -1
}
