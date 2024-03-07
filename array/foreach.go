package array

func ForEach[T any](slice []T, callback func(T)) {
	for _, v := range slice {
		callback(v)
	}
}
