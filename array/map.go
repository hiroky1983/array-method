package array


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
