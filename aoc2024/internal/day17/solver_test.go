package day17

import (
	"fmt"
	"strconv"
	"testing"
)

const exampleInitializationAndProgram = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

const a2024Program = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

const a10Program = `Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`

const a2024v2Program = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "PartOne computes the output of the given program",
			input:    exampleInitializationAndProgram,
			expected: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			name:     "PartOne computes the output of the A=2024 program",
			input:    a2024Program,
			expected: "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			name:     "PartOne computes the output of the A=10 program",
			input:    a10Program,
			expected: "0,1,2",
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
		name                       string
		input                      string
		extractedArbitaryOperation func(int) int
		expected                   string
	}{
		{
			name:  "PartTwo finds the initial register A value to make the program output a copy of itself",
			input: a2024v2Program,
			extractedArbitaryOperation: func(dividend int) int {
				return (dividend / 8) % 8
			},
			expected: "117440",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Copy pasted `PartTwo` to allow for changing some parameters.
			result, err := func() (string, error) {
				initialRegistersAndProgram, err := parseInitialRegistersAndProgram(tc.input)
				if err != nil {
					return "", err
				}

				quineInitValue, ok := findQuineInitValue(
					initialRegistersAndProgram.Program, tc.extractedArbitaryOperation)
				if !ok {
					return "", fmt.Errorf("a quine-producing initial value was not found")
				}

				return strconv.Itoa(quineInitValue), nil
			}()

			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %s, wanted %s", result, tc.expected)
			}
		})
	}
}
