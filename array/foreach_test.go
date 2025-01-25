package array

import (
	"reflect"
	"testing"
)

func TestForEach(t *testing.T) {
	name1 := "Alice"
	age1 := 25
	name2 := "Bob"
	age2 := 30
	personsPointer := []TestStructPointer{
		{Name: &name1, Age: &age1},
		{Name: &name2, Age: &age2},
	}

	names := []string{}
	ForEach(personsPointer, func(v TestStructPointer) {
		names = append(names, *v.Name)
	})

	if !reflect.DeepEqual(names, []string{"Alice", "Bob"}) {
		t.Errorf("ForEach() = %v, want %v", names, []string{"Alice", "Bob"})
	}
}
