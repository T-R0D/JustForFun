package day13

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	// Arrange.
	input := "939\n7,13,x,x,59,x,31,19"

	// Act.
	inputPair, err := parseInput(input)

	// Assert.
	assert.NoError(t, err)
	assert.Equal(t, 939, inputPair.EarliestDepartureTime)
	assert.Equal(t, []int{7, 13, -1, -1, 59, -1, 31, 19}, inputPair.BusIDs)
}

func TestFindTimeToWaitForBus(t *testing.T) {
	testCases := []struct {
		earliestDeparture int
		busID             int
		expectedWaitTime  int
	}{
		{
			earliestDeparture: 939,
			busID:             7,
			expectedWaitTime:  6,
		},
		{
			earliestDeparture: 939,
			busID:             13,
			expectedWaitTime:  10,
		},
		{
			earliestDeparture: 939,
			busID:             59,
			expectedWaitTime:  5,
		},
		{
			earliestDeparture: 939,
			busID:             31,
			expectedWaitTime:  22,
		},
		{
			earliestDeparture: 939,
			busID:             19,
			expectedWaitTime:  11,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf(
			"leaving as early as %d, bus %d has to be waited %d for",
			tc.earliestDeparture, tc.busID, tc.expectedWaitTime)
		t.Run(name, func(t *testing.T) {
			timeToWait := findTimeToWaitForBus(tc.earliestDeparture, tc.busID)

			assert.Equal(t, tc.expectedWaitTime, timeToWait)
		})
	}
}

func TestFindMagicTimeT(t *testing.T) {
	testCases := []struct {
		busSchedule        []int
		expectedMagicTimeT uint64
	}{
		{
			busSchedule:        []int{7, 13, brokenDownBusID, brokenDownBusID, 59, brokenDownBusID, 31, 19},
			expectedMagicTimeT: 1068781,
		},
		{
			busSchedule:        []int{17, brokenDownBusID, 13, 19},
			expectedMagicTimeT: 3417,
		},
		{
			busSchedule:        []int{67, 7, 59, 61},
			expectedMagicTimeT: 754018,
		},
		{
			busSchedule:        []int{67, brokenDownBusID, 7, 59, 61},
			expectedMagicTimeT: 779210,
		},
		{
			busSchedule:        []int{67, 7, brokenDownBusID, 59, 61},
			expectedMagicTimeT: 1261476,
		},
		{
			busSchedule:        []int{1789, 37, 47, 1889},
			expectedMagicTimeT: 1202161486,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf(
			"bus shedule %v has magic time t %d",
			tc.busSchedule, tc.expectedMagicTimeT)
		t.Run(name, func(t *testing.T) {
			magicTimeT := findMagicTimeT(tc.busSchedule)

			assert.Equal(t, tc.expectedMagicTimeT, magicTimeT)
		})
	}
}
