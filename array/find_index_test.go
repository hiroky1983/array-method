package array

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	name1 := "Alice"
	age1 := 25
	name2 := "Bob"
	age2 := 30
	personsPointer := []TestStructPointer{
		{Name: &name1, Age: &age1},
		{Name: &name2, Age: &age2},
	}

	index := FindIndex(personsPointer, func(v TestStructPointer) bool {
		return *v.Name == "Bob"
	})

	if index != 1 {
		t.Errorf("FindIndex() = %v, want %v", index, 1)
	}
}

func TestFindIndex_Primitive(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	index := FindIndex(nums, func(v int) bool {
		return v == 3
	})

	if index != 2 {
		t.Errorf("FindIndex() = %v, want %v", index, 2)
	}
}

func TestFindIndex_Nil(t *testing.T) {
	var nums []int

	index := FindIndex(nums, func(v int) bool {
		return v == 3
	})

	if index != -1 {
		t.Errorf("FindIndex() = %v, want %v", index, -1)
	}
}
