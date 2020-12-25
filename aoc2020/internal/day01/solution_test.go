package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Sample input produces 514579",
			input: "1721\n979\n366\n299\n675\n1456",
			expectedOutput: "514579",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange.
			s := Solver{}

			// Act.
			output, err := s.Part1(tc.input)

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name string
		input string
		expectedOutput string
	}{
		{
			name: "Sample input produces 241861950",
			input: "1721\n979\n366\n299\n675\n1456",
			expectedOutput: "241861950",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange.
			s := Solver{}

			// Act.
			output, err := s.Part2(tc.input)

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
