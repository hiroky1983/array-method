package array

import (
	"testing"
)

func TestSome(t *testing.T) {
	name1 := "Alice"
	age1 := 25
	name2 := "Bob"
	age2 := 30
	personsPointer := []TestStructPointer{
		{Name: &name1, Age: &age1},
		{Name: &name2, Age: &age2},
	}

	hasAlice := Some(personsPointer, func(v TestStructPointer) bool {
		return *v.Name == "Alice"
	})

	if !hasAlice {
		t.Errorf("Some() = %v, want %v", hasAlice, true)
	}
}
