# Array Method Package

[![Go Reference](https://pkg.go.dev/badge/github.com/hiroky1983/array-method.svg)](https://pkg.go.dev/github.com/hiroky1983/array-method)

This package provides convenient methods for arrays and slices in Go, inspired by JavaScript's Array methods. It allows for easy operations such as filtering, mapping, searching, and more.

## Installation

To use this package, install it with the following command:

```bash
go get github.com/hiroky1983/array-method/array
```

## Usage

Below are examples of how to use the main features of this package.

### Filter

Extract elements from a slice that meet a condition.

```go
import (
    "github.com/hiroky1983/array-method/array"
)

nums := []int{1, 2, 3, 4, 5}
evens := array.Filter(nums, func(n int) bool {
    return n%2 == 0
})
// evens will be [2, 4]
```

### Map

Apply a function to each element of a slice to create a new slice.

```go
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
}
names := array.Map(people, func(p Person) string {
    return p.Name
})
// names will be ["Alice", "Bob"]
```

### Find

Find the first element in a slice that meets a condition.

```go
person := array.Find(people, func(p Person) bool {
    return p.Name == "Alice"
})
// person will be &{Alice 25}
```

### FindIndex

Find the index of the first element in a slice that meets a condition.

```go
index := array.FindIndex(people, func(p Person) bool {
    return p.Name == "Bob"
})
// index will be 1
```

### ForEach

Perform an action for each element in a slice.

```go
array.ForEach(people, func(p Person) {
    fmt.Println(p.Name)
})
// Output: Alice Bob
```

### Reduce

Reduce a slice to a single value.

```go
result := array.Reduce(nums, func(prev int, current int) int {
    return prev + current
}, 0)
// result will be 15
```

### Some

Check if at least one element in a slice meets a condition.

```go
result := array.Some(nums, func(n int) bool {
    return n%2 == 0
})
// result will be true
```
