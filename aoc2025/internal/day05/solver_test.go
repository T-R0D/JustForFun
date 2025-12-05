package day05

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct{
		name     string
		input    string
		expected string
	}{
		{
			name: "Part one counts fresh ingredient IDs",
			input: testDB,
			expected: "3",
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
	}{
		{
			name: "Part two finds count of unique fresh ingredient IDs",
			input: testDB,
			expected: "14",
		},
	}

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

const testDB = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`