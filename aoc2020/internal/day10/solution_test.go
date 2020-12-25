package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var smallDataSet = adapterSet{
	16: struct{}{},
	10: struct{}{},
	15: struct{}{},
	5:  struct{}{},
	1:  struct{}{},
	11: struct{}{},
	7:  struct{}{},
	19: struct{}{},
	6:  struct{}{},
	12: struct{}{},
	4:  struct{}{},
}

var largerDataSet = adapterSet{
	28: struct{}{},
	33: struct{}{},
	18: struct{}{},
	42: struct{}{},
	31: struct{}{},
	14: struct{}{},
	46: struct{}{},
	20: struct{}{},
	48: struct{}{},
	47: struct{}{},
	24: struct{}{},
	23: struct{}{},
	49: struct{}{},
	45: struct{}{},
	19: struct{}{},
	38: struct{}{},
	39: struct{}{},
	11: struct{}{},
	1:  struct{}{},
	32: struct{}{},
	25: struct{}{},
	35: struct{}{},
	8:  struct{}{},
	17: struct{}{},
	7:  struct{}{},
	9:  struct{}{},
	4:  struct{}{},
	2:  struct{}{},
	34: struct{}{},
	10: struct{}{},
	3:  struct{}{},
}

func TestConnectAllAdapters(t *testing.T) {
	testCases := []struct {
		name                             string
		adapters                         adapterSet
		expectedJoltageDifferenceSummary map[int]int
	}{
		{
			name:     "small adapter set processes successfully",
			adapters: smallDataSet,
			expectedJoltageDifferenceSummary: map[int]int{
				1: 7,
				3: 4,
			},
		},
		{
			name:     "large adapter set processes successfully",
			adapters: largerDataSet,
			expectedJoltageDifferenceSummary: map[int]int{
				1: 22,
				3: 9,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			joltageDifferenceSummary, err := connectAllAdapters(tc.adapters, maxAllowedJoltageDifference)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedJoltageDifferenceSummary, joltageDifferenceSummary)
		})
	}
}

func TestFindNumberOfWaysToLinkAdapters(t *testing.T) {
	testCases := []struct {
		name                               string
		adapters                           adapterSet
		expectedNumberOfWaysToLinkAdapters int64
	}{
		{
			name:                               "ways to link small set is found successfully",
			adapters:                           smallDataSet,
			expectedNumberOfWaysToLinkAdapters: 8,
		},
		{
			name:                               "ways to link larger set is found successfully",
			adapters:                           largerDataSet,
			expectedNumberOfWaysToLinkAdapters: 19208,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			waysToLinkAdapters := findNumberOfWaysToLinkAdapters(tc.adapters, maxAllowedJoltageDifference)

			assert.Equal(t, tc.expectedNumberOfWaysToLinkAdapters, waysToLinkAdapters)
		})
	}
}
