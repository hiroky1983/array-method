package array

import (
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type NewPerson struct {
	Name string
	Age  int
	Address string
}

type Address struct {
	Address string
}

type Job struct {
	Job string
}

func TestMap(t *testing.T) {
	name1 := "Alice"
	age1 := 25
	name2 := "Bob"
	age2 := 30
	personsPointer := []TestStructPointer{
		{Name: &name1, Age: &age1},
		{Name: &name2, Age: &age2},
	}

	mapped := Map(personsPointer, func(v TestStructPointer) TestStructPointer {
		newName := *v.Name + "!"
		return TestStructPointer{Name: &newName, Age: v.Age}
	})

	if *mapped[0].Name != "Alice!" || *mapped[1].Name != "Bob!" {
		t.Errorf("Map() = %v, want %v", mapped, []string{"Alice!", "Bob!"})
	}
}

func TestMapReturnNewStruct(t *testing.T) {


	persons := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
	}

	mapped := Map(persons, func(p Person) NewPerson {
		return NewPerson{Name: p.Name, Age: p.Age + 1, Address: "Tokyo"}
	})

	if mapped[0].Name != "Alice" || mapped[0].Age != 26 || mapped[0].Address != "Tokyo" || mapped[1].Name != "Bob" || mapped[1].Age != 31 || mapped[1].Address != "Tokyo" {
		t.Errorf("Map() = %v, want %v", mapped, []NewPerson{{Name: "Alice", Age: 26, Address: "Tokyo"}, {Name: "Bob", Age: 31, Address: "Tokyo"}})
	}
}

func TestMapReturnNewStructMerge(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
	}



	type NewPerson struct {
		Name string
		Age  int
		Address
		Job
	}

	mapped := Map(persons, func(p Person) NewPerson {
		return NewPerson{Name: p.Name, Age: p.Age, Address: Address{Address: "Tokyo"}, Job: Job{Job: "Engineer"}}
	})

	if mapped[0].Name != "Alice" || mapped[0].Age != 25 || mapped[0].Address.Address != "Tokyo" || mapped[0].Job.Job != "Engineer" || mapped[1].Name != "Bob" || mapped[1].Age != 30 || mapped[1].Address.Address != "Tokyo" || mapped[1].Job.Job != "Engineer" {
		t.Errorf("Map() = %v, want %v", mapped, []NewPerson{{Name: "Alice", Age: 25, Address: Address{Address: "Tokyo"}, Job: Job{Job: "Engineer"}}, {Name: "Bob", Age: 30, Address: Address{Address: "Tokyo"}, Job: Job{Job: "Engineer"}}})
	}	
}

func TestMapReturnEmpty(t *testing.T) {
	persons := []Person{}

	mapped := Map(persons, func(p Person) Person {
		return Person{Name: p.Name, Age: p.Age}
	})

	if len(mapped) != 0 {
		t.Errorf("Map() = %v, want %v", mapped, []Person{})
	}
}
