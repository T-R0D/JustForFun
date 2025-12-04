package day04

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Part one counts rolls with fewer than 4 neighbors",
			input:    testGrid,
			expected: "13",
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
			name:     "Part two counts rolls that can be removed (cascading)",
			input:    testGrid,
			expected: "43",
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

const testGrid = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
