package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name         string
		Input        interface{}
		ExpectedCall []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "Tokyo"},
			[]string{"Chris", "Tokyo"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields v.1",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"nested fields v.2",
			Person{
				"Chris",
				Profile{33, "Tokyo"},
			},
			[]string{"Chris", "Tokyo"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "Tokyo"},
			},
			[]string{"Chris", "Tokyo"},
		},
		{
			"slices",
			[]Profile{
				{33, "Ural"},
				{34, "Reykjavík"},
			},
			[]string{"Ural", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Ural"},
				{34, "Reykjavík"},
			},
			[]string{"Ural", "Reykjavík"},
		},
		{
			"map",
			map[string]string{
				"Bar": "Boo",
				"Foo": "Faz",
			},
			[]string{"Boo", "Faz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCall) {
				t.Errorf("got %v but want %v", got, test.ExpectedCall)
			}
		})
	}
}
