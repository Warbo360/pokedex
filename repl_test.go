package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Mario Luigi Toad",
			expected: []string{"mario", "luigi", "toad"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
        	t.Errorf("expected: %v, got: %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
            expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
				t.FailNow()
			}
		}
	}
}
