package array

// Map applies a given transformation function to each element of the provided slice
// and returns a new slice containing the results.
// If the input slice is nil, it returns nil.
func Map[Input any, Output any](slice []Input, transform func(Input) Output) []Output {
	if slice == nil {
		return nil
	}
	result := make([]Output, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}
