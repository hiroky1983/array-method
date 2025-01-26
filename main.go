package main

import (
	"github.com/davecgh/go-spew/spew"
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
	// reduce
	personWithIsAdult := array.Reduce(pc, func(prev []NewPerson, current Person) []NewPerson {
		return append(prev, NewPerson{Person: Person{Name: current.Name, Age: current.Age}, IsAdult: current.Age >= 20})
	}, []NewPerson{})

	spew.Dump(personWithIsAdult)

	nums := []int{1, 2, 3, 4, 5}

	incrimentNumber := array.Reduce(nums, func(prev int, current int) int {
		return prev + current
	}, 0)
	spew.Dump(incrimentNumber)
}
