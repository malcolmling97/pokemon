package main

import "testing"

func TestCleanInput(t *testing.T) {

	// create some test cases
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  bye  world  ",
			expected: []string{"bye", "world"},
		},
		{
			input:    "  i am not from here  ",
			expected: []string{"i", "am", "not", "from", "here"},
		},
	}

	// loop thorugh and evaluate
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		if len(c.expected) != len(actual) {
			t.Errorf("Lengths don't match: '%v' vs '%v' ", actual, c.expected)
		}
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
