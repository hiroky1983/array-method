package array

import (
	"testing"
)

func TestFind(t *testing.T) {
	name1 := "Alice"
	age1 := 25
	name2 := "Bob"
	age2 := 30
	personsPointer := []TestStructPointer{
		{Name: &name1, Age: &age1},
		{Name: &name2, Age: &age2},
	}

	found := Find(personsPointer, func(v TestStructPointer) bool {
		return *v.Name == "Alice"
	})

	if found == nil || *found.Name != "Alice" {
		t.Errorf("Find() = %v, want %v", found, "Alice")
	}
}

func TestFind_Primitive(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	found := Find(nums, func(v int) bool {
		return v == 3
	})

	if found == nil || *found != 3 {
		t.Errorf("Find() = %v, want %v", found, 3)
	}
}

func TestFind_Nil(t *testing.T) {
	var nums []int

	found := Find(nums, func(v int) bool {
		return v == 3
	})

	if found != nil {
		t.Errorf("Find() = %v, want %v", found, nil)
	}
}
