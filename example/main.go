package main

import "github.com/hiroky1983/array-method/array"

func ExampleFilter() {
	array.Filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	// Output: [2 4]
}
