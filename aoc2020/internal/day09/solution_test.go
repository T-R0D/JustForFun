package day09

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleXMASData = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

var sampleData = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}

func TestParseXMASData(t *testing.T) {
	data, err := parseXMASData(sampleXMASData)

	assert.NoError(t, err)
	assert.Equal(t, sampleData, data)
}

func TestXMASProcessor(t *testing.T) {
	processor := newXMASProcessor(sampleData, 5)
	currentData := processor.CurrentData()

	iterations := 0
	for currentData != 127 && iterations < 9 {
		currentValueValid, err := processor.CurrentDataIsValidXMAS()
		assert.NoError(t, err)
		assert.True(t, currentValueValid, "data is not valid: %d", currentData)

		advanced, err := processor.Advance()
		assert.NoError(t, err)
		assert.True(t, advanced)

		currentData = processor.CurrentData()
		iterations++
	}

	assert.Equal(t, 127, processor.CurrentData())
}

func TestFindMaxMinInSumRange(t *testing.T) {
	testCases := []struct{
		target int
		expectedMax int
		expectedMin int
	}{
		{
			target: 127,
			expectedMax: 47,
			expectedMin: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d results in max of %d and min of %d", tc.target, tc.expectedMax, tc.expectedMin), func(t *testing.T) {
			result, err := findMaxMinInSumRange(sampleData, tc.target)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedMax, result.max)
			assert.Equal(t, tc.expectedMin, result.min)
		})
	}
}