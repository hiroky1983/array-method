package array

func Map[Input any, Output any](slice []Input, transform func(Input) Output) []Output {
	var result []Output
	for _, v := range slice {
		result = append(result, transform(v))
	}
	return result
}
