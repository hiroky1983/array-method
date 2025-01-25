package array

import "reflect"

type ArrayType[T any] []T

func NewArray[T any](slice []T) ArrayType[T] {
	result := make([]T, len(slice))
	copy(result, slice)
	return result
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
