package day12

import (
	"testing"
)

const smallExampleGarden = `AAAA
BBCD
BBCC
EEEC`

const exampleGarden = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const exampleEGarden = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`

const exampleCheckeredGarden = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "PartOne computes the total cost for fencing for a small garden",
			input:    smallExampleGarden,
			expected: "140",
		},
		{
			name:     "PartOne computes the total cost for fencing for a garden",
			input:    exampleGarden,
			expected: "1930",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

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
			name:     "PartTwo computes the total cost with bulk discount applied for fencing for a small garden",
			input:    smallExampleGarden,
			expected: "80",
		},
		{
			name:     "PartTwo computes the total cost with bulk discount applied for fencing for a garden",
			input:    exampleGarden,
			expected: "1206",
		},
		{
			name:     "PartTwo computes the total cost with bulk discount applied for fencing for an E garden",
			input:    exampleEGarden,
			expected: "236",
		},
		{
			name:     "PartTwo computes the total cost with bulk discount applied for fencing for a checkered garden",
			input:    exampleCheckeredGarden,
			expected: "368",
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
