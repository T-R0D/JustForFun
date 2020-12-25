package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasValidPasswordByFrequency(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedResult bool
	}{
		{
			name: "'1-3 a: abcde' has a valid password",
			input: "1-3 a: abcde",
			expectedResult: true,
		},
		{
			name: "'1-3 b: cdefg' has a valid password",
			input: "1-3 b: cdefg",
			expectedResult: false,
		},
		{
			name: "'2-9 c: ccccccccc' has a valid password",
			input: "2-9 c: ccccccccc",
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange.
		
			// Act.
			result, err := hasValidPasswordByFrequency(tc.input)

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestHasValidPasswordByPlacement(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedResult bool
	}{
		{
			name: "'1-3 a: abcde' has a valid password",
			input: "1-3 a: abcde",
			expectedResult: true,
		},
		{
			name: "'1-3 b: cdefg' has a valid password",
			input: "1-3 b: cdefg",
			expectedResult: false,
		},
		{
			name: "'2-9 c: ccccccccc' has a valid password",
			input: "2-9 c: ccccccccc",
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange.
		
			// Act.
			result, err := hasValidPasswordByFrequency(tc.input)

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
