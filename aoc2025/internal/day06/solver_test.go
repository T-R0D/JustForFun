package day06

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
			name:     "Part one computes the checksum properly for given problems in human arithmetic",
			input:    testWorksheet,
			expected: "4277556",
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
			name:     "Part two computes the checksum properly for given problems in cephalopod arithmetic",
			input:    testWorksheet,
			expected: "3263827",
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

const testWorksheet = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
