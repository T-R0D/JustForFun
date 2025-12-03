package day03

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
			name: "Part one sums max possible joltages (from 2 batteries) properly",
			input: testBatteries,
			expected: "357",
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
			name: "Part two sums max possible joltages (from 12 batteries) properly",
			input: testBatteries,
			expected: "3121910778619",
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

const testBatteries = `987654321111111
811111111111119
234234234234278
818181911112111`
