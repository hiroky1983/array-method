package array

// FindIndex searches for the index of the first element in the provided slice that satisfies the given condition.
// If the input slice is nil, it returns -1.
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
