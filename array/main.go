package array

import "reflect"

type ArrayType[T any, U any] []T

func NewArray[T any, U any](slice []T) ArrayType[T, U] {
	return ArrayType[T, U](slice)
}

func (a ArrayType[T, U]) Filter(condition func(T) bool) []T {
	return Filter(a, condition)
}

func (a ArrayType[T, U]) Find(condition func(T) bool) *T {
	return Find(a, condition)
}

func (a ArrayType[T, U]) FindIndex(condition func(T) bool) int {
	return FindIndex(a, condition)
}

func (a ArrayType[T, U]) Map(condition func(T) T) []T {
	return Map(a, condition)
}

func (a ArrayType[T, U]) ForEach(condition func(T)) {
	ForEach(a, condition)
}

func (a ArrayType[T, U]) Some(condition func(T) bool) bool {
	return Some(a, condition)
}

func (a ArrayType[T, U]) Reduce(reducer func(U, T) U, initial U) U {
	return Reduce(a, reducer, initial)
}

// hasNilPointer は、構造体のポインタフィールドがnilの場合にtrueを返します
func hasNilPointer[T any](v T) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			return true
		}
	}
	return false
}

// copyValue は、ポインタ型の要素を新しいメモリにコピーします
func copyValue[T any](v T) T {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	newVal := reflect.New(val.Type()).Elem()
	newVal.Set(val)

	return newVal.Interface().(T)
}
