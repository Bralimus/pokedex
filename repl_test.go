package main

import "testing"

func TestCleanInput(t *testing.T) {
	//Defining test cases
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "HOLA AMIGOS",
			expected: []string{"hola", "amigos"},
		}, {
			input:    "   ",
			expected: []string{},
		}, {
			input:    "  PokeMon  <  TemTem ",
			expected: []string{"pokemon", "<", "temtem"},
		},
		// add more cases here
	}

	//Loop and run the tests
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Expected length %d but got %d for input %q", len(c.expected), len(actual), c.input)
			continue //don't return in tests, and already failed here to no need to check the next loop
		}
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %q but got %q for input %q at position %d", expectedWord, word, c.input, i)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
