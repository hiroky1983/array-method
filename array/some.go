package array

func Some[T any](slice []T, condition func(T) bool) bool {
	for _, v := range slice {
		if condition(v) {
			return true
		}
	}
	return false
}
