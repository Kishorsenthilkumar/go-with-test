package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile profile
}
type profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"nested fields",
			Person{
				"Chris",
				profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
