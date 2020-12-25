package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleResponses = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestParseResponses(t *testing.T) {
	// Arrange.
	expectedOutput := [][]string{
		{"abc"},
		{"a", "b", "c",},
		{"ab", "ac",},
		{"a", "a", "a", "a"},
		{"b"},
	}

	// Act.
	responses := parseResponses(sampleResponses)

	// Assert.
	assert.Equal(t, expectedOutput, responses)
}	
