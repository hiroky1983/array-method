package array

// Filter applies a condition function to each element of a slice and returns a new slice containing only the elements that satisfy the condition.
// If the input slice is nil, returns nil.
// If an element contains nil pointer fields, it is skipped.
// The original slice is not modified.
func Filter[T any](slice []T, condition func(T) bool) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if hasNilPointer(v) {
			continue
		}
		if condition(v) {
			copiedValue := copyValue(v)
			result = append(result, copiedValue)
		}
	}
	return result
}
