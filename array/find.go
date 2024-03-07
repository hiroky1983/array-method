package array

func Find[T any](slice []T, condition func(T) bool) *T {
	for _, v := range slice {
		if condition(v) {
			return &v
		}
	}
	return nil
}
