package array

import (
	"testing"
)

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
