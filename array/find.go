package array

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
