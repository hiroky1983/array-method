package main

import (
	"fmt"

	"github.com/hiroky1983/array-method/array"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := array.Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evens) // [2 4 6 8 10]

	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Carol", 35},
		{"Dave", 40},
	}
	persons := array.Map(people, func(p Person) Person {
		return p
	})
	fmt.Println(persons) // [Alice Bob Carol Dave]
	names := array.Map(people, func(p Person) string {
		return p.Name
	})
	fmt.Println(names) // [Alice Bob Carol Dave]

	findBob := array.Filter(people, func(p Person) bool {
		return p.Name == "Bob"
	}) 
	fmt.Println(findBob) // [{Bob 30}]
	fmt.Println(people)

	// Reduceを使ってみる
	sum := array.Reduce(nums, 0, func(acc int, n int) int {
		return acc + n
	})
	fmt.Println(sum) // 55

	// Findを使ってみる
	findAlice := array.Find(people, func(p Person) bool {
		return p.Name == "Alice"
	})
	fmt.Println(findAlice) // &{Alice 25}

	notFind := array.Find(people, func(p Person) bool {
		return p.Name == "Oyaji"
	})
	fmt.Println(notFind) // <nil>

	// FindIndexを使ってみる
	findIndexCarol := array.FindIndex(people, func(p Person) bool {
		return p.Name == "Carol"
	})
	fmt.Println(findIndexCarol) // 2

	notFindIndex := array.FindIndex(people, func(p Person) bool {
		return p.Name == "Oyaji"
	})
	fmt.Println(notFindIndex) // -1

	// Someを使ってみる
	someAlice := array.Some(people, func(p Person) bool {
		return p.Name == "Alice"
	})
	fmt.Println(someAlice) // true
}