package array

func Filter[T any](slice []T, condition func(T) bool) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, 0, len(slice))

	for _, v := range slice {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}
