package array

func Some[T any](slice []T, condition func(T) bool) bool {
	if slice == nil {
		return false
	}
	for _, v := range slice {
		if condition(v) {
			return true
		}
	}
	return false
}
