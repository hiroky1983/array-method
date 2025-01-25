package array

import "github.com/davecgh/go-spew/spew"

type ArrayType[T any] []T

func NewArray[T any](slice []T) ArrayType[T] {
	return slice
}

func (a ArrayType[T]) Filter(condition func(T) bool) ArrayType[T] {
	return Filter(a, condition)
}

func (a ArrayType[T]) Find(condition func(T) bool) *T {
	return Find(a, condition)
}

func (a ArrayType[T]) FindIndex(condition func(T) bool) int {
	return FindIndex(a, condition)
}

func (a ArrayType[T]) Map(condition func(T) T) ArrayType[T] {
	return Map(a, condition)
}

func (a ArrayType[T]) ForEach(condition func(T)) {
	ForEach(a, condition)
}

func (a ArrayType[T]) Some(condition func(T) bool) bool {
	return Some(a, condition)
}

// copyValue は、ポインタ型の要素を新しいメモリにコピーします
func copyValue[T any](v T) T {
	// ポインタ型の場合、新しいメモリにコピー
	newV := new(T)
	*newV = v
	spew.Dump(newV)
	spew.Dump(v)
	if ptr, ok := any(v).(*T); ok {
		newPtr := new(T)
		*newPtr = *ptr
		return *newPtr
	}
	return *newV
}
