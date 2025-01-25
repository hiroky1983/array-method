package array

// Find searches for the first element in the provided slice that satisfies the given condition.
// If the input slice is nil, it returns nil.
func Find[T any](slice []T, condition func(T) bool) *T {
	if slice == nil {
		return nil
	}
	for _, v := range slice {
		if condition(v) {
			return &v
		}
	}
	return nil
}
