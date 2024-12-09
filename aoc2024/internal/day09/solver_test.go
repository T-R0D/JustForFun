package day09

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
			name:     "PartOne computes the checksum of the compacted disk",
			input:    "2333133121414131402",
			expected: "1928",
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
			name:     "PartTwo computes the checksum of the compacted disk (without fragmenting files)",
			input:    "2333133121414131402",
			expected: "2858",
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
