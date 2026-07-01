package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pashgoipsahgoishgoi",
			expected: []string{"pashgoipsahgoishgoi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Wrong length")
			return
		}
		for i, word := range actual {
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("Incorrect output")
				return
			}
		}
	}

}
