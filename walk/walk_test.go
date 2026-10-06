package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Description   string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 21},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{21, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{
					21,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{21, "London"},
				{22, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
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
		t.Run(test.Description, func(t *testing.T) {
			var got []string

			ourFunc := func(input string) {
				got = append(got, input)
			}

			walk(test.Input, ourFunc)

			// if len(got) != 1 {
			// 	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
			// }

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string

		ourFunc := func(input string) {
			got = append(got, input)
		}

		walk(aMap, ourFunc)

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}
		var got []string

		ourFunc := func(input string) {
			got = append(got, input)
		}

		walk(aMap, ourFunc)

		assertLength(t, got, len(aMap))
		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		want := []string{"Berlin", "Katowice"}
		var got []string

		ourFunc := func(input string) {
			got = append(got, input)
		}

		walk(aChannel, ourFunc)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		want := []string{"Berlin", "Katowice"}
		var got []string

		ourFunc := func(input string) {
			got = append(got, input)
		}

		walk(aFunction, ourFunc)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("circular references", func(t *testing.T) {
		type Person struct {
			Name   string
			Friend *Person
		}
		p := Person{Name: "Alice"}
		p.Friend = &p // cycle!

		var got []string

		ourFunc := func(input string) {
			got = append(got, input)
		}

		walk(p, ourFunc)
	})
}

func assertLength(t testing.TB, got []string, want int) {
	t.Helper()
	if len(got) != want {
		t.Errorf("got %d values but expected %d", len(got), want)
	}
}

func assertContains(t testing.TB, got []string, expected string) {
	t.Helper()

	contains := false
	for _, value := range got {
		if value == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", got, expected)
	}
}
