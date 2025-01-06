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
	}
	// add more cases here

	for _, c := range cases {
		actual := cleanInput(c.input)

		length := len(actual)
		expectedLength := len(c.expected)
		if length != expectedLength {
			t.Errorf("got %d slices but expected %d", length, expectedLength)
		}

		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("got word '%s' but expected '%s'", word, expectedWord)
			}
		}
	}

}