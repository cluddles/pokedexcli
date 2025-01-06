package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		length := len(actual)
		expectedLength := len(c.expected)
		if length != expectedLength {
			t.Errorf("cleanInput('%s'): got %d slices but expected %d", c.input, length, expectedLength)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput('%s'): got word '%s' but expected '%s'", c.input, word, expectedWord)
			}
		}
	}

}
