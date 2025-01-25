# Array Method Package

[![Go Reference](https://pkg.go.dev/badge/github.com/hiroky1983/array-method.svg)](https://pkg.go.dev/github.com/hiroky1983/array-method)

[Go Doc](https://pkg.go.dev/github.com/hiroky1983/array-method)

This package provides convenient methods for arrays and slices in Go, inspired by JavaScript's Array methods. It allows for easy operations such as filtering, mapping, searching, and more.

## Installation

To use this package, install it with the following command:

```bash
go get github.com/hiroky1983/array-method/array
```

## Usage

Below are examples of how to use the main features of this package.

### Filter

#### Function Call

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

#### Method Chaining

Use `NewArray` to filter with method chaining.

```go
nums := []int{1, 2, 3, 4, 5}
evens := array.NewArray(nums).Filter(func(n int) bool {
    return n%2 == 0
})
// evens will be [2, 4]
```

### Map

#### Function Call

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

#### Method Chaining

```go
people := []Person{
    {"Alice", 25},
    {"Bob", 30},
}
names := array.NewArray(people).Map(func(p Person) string {
    return p.Name
})
// names will be ["Alice", "Bob"]
```

### Find

#### Function Call

Find the first element in a slice that meets a condition.

```go
person := array.Find(people, func(p Person) bool {
    return p.Name == "Alice"
})
// person will be &{Alice 25}
```

#### Method Chaining

```go
person := array.NewArray(people).Find(func(p Person) bool {
    return p.Name == "Alice"
})
// person will be &{Alice 25}
```

### FindIndex

#### Function Call

Find the index of the first element in a slice that meets a condition.

```go
index := array.FindIndex(people, func(p Person) bool {
    return p.Name == "Bob"
})
// index will be 1
```

#### Method Chaining

```go
index := array.NewArray(people).FindIndex(func(p Person) bool {
    return p.Name == "Bob"
})
// index will be 1
```

### ForEach

#### Function Call

Perform an action for each element in a slice.

```go
array.ForEach(people, func(p Person) {
    fmt.Println(p.Name)
})
// Output: Alice Bob
```

#### Method Chaining

```go
array.NewArray(people).ForEach(func(p Person) {
    fmt.Println(p.Name)
})
// Output: Alice Bob
```
