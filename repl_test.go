package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		input: "Hello WORLd",
		expected: []string{
			"hello",
			"world",
		},
	},
	for _, cs := range cases {
		actual := cleanInput(cs.input)

		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v", len(actual), len(c.expected))
		}
	}
}
