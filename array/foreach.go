package array

// ForEach executes a given callback function for each element in the provided slice.
// If the input slice is nil, the function does nothing.
func ForEach[T any](slice []T, callback func(T)) {
	if slice == nil {
		return
	}
	for _, v := range slice {
		callback(v)
	}
}
