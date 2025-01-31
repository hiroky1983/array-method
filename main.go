package main

import (
	"fmt"

	"github.com/hiroky1983/array-method/array"
)

type Person struct {
	Name string
	Age  int
}

type PersonPointer struct {
	Name *string
	Age  *int
}

func main() {
	type NewPerson struct {
		Person Person
		IsAdult bool
	}

	pc := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 15},
		{Name: "David", Age: 18},
		{Name: "Eve", Age: 21},
	}
	// Map
	personMap := array.Map(pc, func(person Person) NewPerson {
		return NewPerson{Person: person, IsAdult: person.Age >= 20}
	})

	fmt.Println(personMap)

	// reduce
	personWithIsAdult := array.Reduce(pc, func(prev []NewPerson, current Person) []NewPerson {
		return append(prev, NewPerson{Person: Person{Name: current.Name, Age: current.Age}, IsAdult: current.Age >= 20})
	}, []NewPerson{})

	fmt.Println(personWithIsAdult)

	nums := []int{1, 2, 3, 4, 5}

	incrimentNumber := array.Reduce(nums, func(prev int, current int) int {
		return prev + current
	}, 0)
	fmt.Println(incrimentNumber)

	// Filter
	evenNumbers := array.Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evenNumbers)

	// Find
	person := array.Find(pc, func(p Person) bool {
		return p.Name == "Alice"
	})
	fmt.Println(person)

	// FindIndex
	index := array.FindIndex(pc, func(p Person) bool {
		return p.Name == "Alice"
	})
	fmt.Println(index)

	// Some
	some := array.Some(pc, func(p Person) bool {
		return p.Name == "Alice"
	})
	fmt.Println(some)


	arr := []int{1, 2, 3, 4, 5}
	mapped := array.Map(arr, func(v int) int {
		return v * 3
	})
	fmt.Println("Mapped:", mapped) // Output: [3 6 9 12 15]
}
