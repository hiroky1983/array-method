package array

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name string
	Age  int
}

type TestStructPointer struct {
	Name *string
	Age  *int
}

func TestFilterTypeInt(t *testing.T) {
	type args struct {
		slice     []int
		condition func(int) bool
	}
	intTests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case array int primitive1",
			args: args{
				slice:     []int{1, 2, 3, 4, 5},
				condition: func(v int) bool { return v%2 == 0 },
			},
			want: []int{2, 4},
		},
		{
			name: "Test case array int primitive2",
			args: args{
				slice:     []int{1, 2, 3, 4, 5},
				condition: func(v int) bool { return v == 5 },
			},
			want: []int{5},
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTypeString(t *testing.T) {
	type args struct {
		slice     []string
		condition func(string) bool
	}
	stringTests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case array string primitive1",
			args: args{
				slice:     []string{"spray", "elite", "exuberant", "destruction", "present"},
				condition: func(v string) bool { return len(v) > 6 },
			},
			want: []string{"exuberant", "destruction", "present"},
		},
		{
			name: "Test case array string primitive2",
			args: args{
				slice:     []string{"spray", "elite", "exuberant", "destruction", "present"},
				condition: func(v string) bool { return v == "exuberant" || v == "destruction" || v == "present"},
			},
			want: []string{"exuberant", "destruction", "present"},
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTypeStruct(t *testing.T) {
	Persons := []TestStruct{
		{"Bob", 15}, {"Joe", 20}, {"Taro", 3}, {"Minji", 44}, {"Ari", 51},
	}
	type args struct {
		slice     []TestStruct
		condition func(TestStruct) bool
	}
	stringTests := []struct {
		name string
		args args
		want []TestStruct
	}{
		{
			name: "Test case array Struct1",
			args: args{
				slice:     Persons,
				condition: func(v TestStruct) bool { return v.Age > 12},
			},
			want: []TestStruct{{"Bob", 15}, {"Joe", 20}, {"Minji", 44}, {"Ari", 51}},
		},
		{
			name: "Test case array Struct2",
			args: args{
				slice:     Persons,
				condition: func(v TestStruct) bool { return  v.Name == "Bob"},
			},
			want: []TestStruct{{"Bob", 15}},
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTypeStructPointer (t *testing.T) {
	name := "Alice"
	age := 25
	personsPointer := []TestStructPointer{
		{Name: &name, Age: &age},
	}
	type args struct {
		slice     []TestStructPointer
		condition func(TestStructPointer) bool
	}
	stringTests := []struct {
		name string
		args args
		want []TestStructPointer
	}{
		{
			name: "Test case array Struct1",
			args: args{
				slice:     personsPointer,
				condition: func(v TestStructPointer) bool { return *v.Age > 12},
			},
			want: personsPointer,
		},
		{
			name: "Test case array Struct1",
			args: args{
				slice:     personsPointer,
				condition: func(v TestStructPointer) bool { return *v.Age > 12},
			},
			want: personsPointer,
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterTypeStructPointerNil (t *testing.T) {
	personsPointer := []TestStructPointer{
		{Name: nil, Age: nil},
	}

	type args struct {
		slice     []TestStructPointer
		condition func(TestStructPointer) bool
	}
	stringTests := []struct {
		name string
		args args
		want []TestStructPointer
	}{
		{
			name: "Test case array Struct1",
			args: args{
				slice:     personsPointer,
				condition: func(v TestStructPointer) bool { return *v.Age > 12},
			},
			want: personsPointer,
		},
		{
			name: "Test case array Struct1",
			args: args{
				slice:     personsPointer,
				condition: func(v TestStructPointer) bool { return *v.Age > 12},
			},
			want: personsPointer,
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
