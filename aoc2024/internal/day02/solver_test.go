package day02

import (
	"testing"
)

const exampleReports = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "PartOne counts safe reports",
			input:    exampleReports,
			expected: "2",
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
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "PartTwo counts safe (with dampening) reports",
			input:    exampleReports,
			expected: "4",
		},
		{
			name:     "Oversize start is 'removed'",
			input:    "30 1 2 3 4 5",
			expected: "1",
		},
		{
			name:     "Oversize end is 'removed'",
			input:    "1 2 3 4 30",
			expected: "1",
		},
		{
			name:     "Two problematic levels results in rejection",
			input:    "30 1 2 3 4 30",
			expected: "0",
		},
		{
			name:     "Two consecutive problematic levels results in rejection",
			input:    "1 2 30 40 3 4",
			expected: "0",
		},
		{
			name:     "A level breaking monotonicity is dampened",
			input:    "3 5 2 1",
			expected: "1",
		},
		{
			name:     "A starting level breaking monotonicity is dampened",
			input:    "5 4 6 7 8",
			expected: "1",
		},
		{
			name:     "Too many duplicates and a break in monotonicity is rejected",
			input:    "1 3 4 2 4 6 7 7",
			expected: "0",
		},
		{
			name:     "A single too-high value is dampened",
			input:    "39 40 42 44 45 51 48",
			expected: "1",
		},
		{
			name:     "A single break in monotonicity is dampened",
			input:    "66 64 62 63 59",
			expected: "1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

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
