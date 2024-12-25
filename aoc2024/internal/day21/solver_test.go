package day21

import (
	"testing"
)

const exampleCodes = `029A
980A
179A
456A
379A`

const code197A = "197A"

func TestPartOne(t *testing.T) {
	testCases := []struct{
		name     string
		input    string
		expected string
	}{
		{
			name: "PartOne finds the sum of complexities for entering the codes",
			input: exampleCodes,
			expected: "126384",
		},
		{
			name: "PartOne finds the sum of complexities for entering the code '197A'",
			input: code197A,
			expected: "17730",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			solver := Solver{}

			result, err := solver.SolvePartOne(tc.input)
			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %s, wanted %s", result, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct{
		name     string
		input    string
		expected string
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			solver := Solver{}

			result, err := solver.SolvePartTwo(tc.input)
			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %s, wanted %s", result, tc.expected)
			}
		})
	}
}
