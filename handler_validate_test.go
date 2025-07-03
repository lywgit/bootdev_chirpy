package main

import (
	"testing"
)

// test case by Gemini
func TestReplaceProfane(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "This is a kerfuffle test",
			expected: "This is a **** test",
		},
		{
			input:    "No profane words here",
			expected: "No profane words here",
		},
		{
			input:    "sharbert and Fornax are bad words",
			expected: "**** and **** are bad words",
		},
		{
			input:    "mixedCaseKerfuffle is not a profane word",
			expected: "mixedCaseKerfuffle is not a profane word",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "kerfuffle",
			expected: "****",
		},
		{
			input:    "KERFUFFLE",
			expected: "****",
		},
		{
			input:    "kerfuffle kerfuffle",
			expected: "**** ****",
		},
		{
			input:    "a",
			expected: "a",
		},
		{
			input:    "fOrNaX is a profane word",
			expected: "**** is a profane word",
		},
	}

	for _, test := range tests {
		actual := replaceProfane(test.input)
		if actual != test.expected {
			t.Errorf("Input: \"%s\", Expected: \"%s\", Got: \"%s\"", test.input, test.expected, actual)
		}
	}
}
