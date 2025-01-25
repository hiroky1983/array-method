package array

func ForEach[T any](slice []T, callback func(T)) {
	if slice == nil {
		return
	}
	for _, v := range slice {
		callback(v)
	}
}
