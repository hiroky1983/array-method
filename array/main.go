package array

type ArrayType[T any] []T

func NewArray[T any](slice []T) ArrayType[T] {
	return slice
}

func (a ArrayType[T]) Filter(condition func(T) bool) ArrayType[T] {
	return ArrayType[T](Filter(a, condition))
}

func (a ArrayType[T]) Find(condition func(T) bool) *T {
	return Find(a, condition)
}

func (a ArrayType[T]) FindIndex(condition func(T) bool) int {
	return FindIndex(a, condition)
}

func (a ArrayType[T]) Map(condition func(T) T) ArrayType[T] {
	return ArrayType[T](Map(a, condition))
}

func (a ArrayType[T]) ForEach(condition func(T)) {
	ForEach(a, condition)
}

func (a ArrayType[T]) Some(condition func(T) bool) bool {
	return Some(a, condition)
}
