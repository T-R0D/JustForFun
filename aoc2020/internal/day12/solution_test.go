package day12

import (
	"fmt"
	"testing"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/stretchr/testify/assert"
)

func TestNewHeading(t *testing.T) {
	testCases := []struct {
		startingHeading int
		direction       rune
		magnitude       int
		expectedHeading int
	}{
		{
			startingHeading: 0,
			direction:       left,
			magnitude:       90,
			expectedHeading: 90,
		},
		{
			startingHeading: 0,
			direction:       left,
			magnitude:       180,
			expectedHeading: 180,
		},
		{
			startingHeading: 90,
			direction:       left,
			magnitude:       180,
			expectedHeading: 270,
		},
		{
			startingHeading: 180,
			direction:       left,
			magnitude:       180,
			expectedHeading: 0,
		},
		{
			startingHeading: 0,
			direction:       right,
			magnitude:       90,
			expectedHeading: 270,
		},
		{
			startingHeading: 0,
			direction:       right,
			magnitude:       180,
			expectedHeading: 180,
		},
		{
			startingHeading: 270,
			direction:       right,
			magnitude:       180,
			expectedHeading: 90,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("starting at %d, turning %c %d degrees, ends at %d", tc.startingHeading, tc.direction, tc.magnitude, tc.expectedHeading)
		t.Run(name, func(t *testing.T) {
			actualHeading, err := findNewHeading(tc.startingHeading, tc.direction, tc.magnitude)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedHeading, actualHeading)
		})
	}
}

func TestShipTakeInstruction(t *testing.T) {
	testCases := []struct {
		givenInstruction    instruction
		expectedI           int
		expectedJ           int
		expectedOrientation int
	}{
		{
			givenInstruction:    instruction{key: 'F', value: 10},
			expectedI:           0,
			expectedJ:           10,
			expectedOrientation: 0,
		},
		{
			givenInstruction:    instruction{key: 'N', value: 10},
			expectedI:           -10,
			expectedJ:           0,
			expectedOrientation: 0,
		},
		{
			givenInstruction:    instruction{key: 'S', value: 10},
			expectedI:           10,
			expectedJ:           0,
			expectedOrientation: 0,
		},
		{
			givenInstruction:    instruction{key: 'E', value: 10},
			expectedI:           0,
			expectedJ:           10,
			expectedOrientation: 0,
		},
		{
			givenInstruction:    instruction{key: 'W', value: 10},
			expectedI:           0,
			expectedJ:           -10,
			expectedOrientation: 0,
		},
		{
			givenInstruction:    instruction{key: 'R', value: 90},
			expectedI:           0,
			expectedJ:           0,
			expectedOrientation: 270,
		},
		{
			givenInstruction:    instruction{key: 'L', value: 90},
			expectedI:           0,
			expectedJ:           0,
			expectedOrientation: 90,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf(
			"instruction %v results in ship at %d, %d and orientation %c",
			tc.givenInstruction, tc.expectedI, tc.expectedJ, tc.expectedOrientation)

		t.Run(name, func(t *testing.T) {
			// Arrange.
			s := newShip()

			// Act.
			err := s.TakeInstruction(tc.givenInstruction.key, tc.givenInstruction.value)
			position := s.Position()

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedI, position.I)
			assert.Equal(t, tc.expectedJ, position.J)
			assert.Equal(t, tc.expectedOrientation, s.orientation)
		})
	}
}

func TestGuidedShipTakeInstruction(t *testing.T) {
	testCases := []struct {
		startingPosition grid.Point
		startingWaypoint grid.Point
		givenInstruction instruction
		expectedPosition grid.Point
		expectedWaypoint grid.Point
	}{
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'F', value: 10},
			expectedPosition: grid.Point{I: -10, J: 100},
			expectedWaypoint: grid.Point{I: -1, J: 10},
		},
		{
			startingPosition: grid.Point{I: -10, J: 100},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'N', value: 3},
			expectedPosition: grid.Point{I: -10, J: 100},
			expectedWaypoint: grid.Point{I: -4, J: 10},
		},
		{
			startingPosition: grid.Point{I: -10, J: 100},
			startingWaypoint: grid.Point{I: -4, J: 10},
			givenInstruction: instruction{key: 'F', value: 7},
			expectedPosition: grid.Point{I: -38, J: 170},
			expectedWaypoint: grid.Point{I: -4, J: 10},
		},
		{
			startingPosition: grid.Point{I: -38, J: 170},
			startingWaypoint: grid.Point{I: -4, J: 10},
			givenInstruction: instruction{key: 'R', value: 90},
			expectedPosition: grid.Point{I: -38, J: 170},
			expectedWaypoint: grid.Point{I: 10, J: 4},
		},
		{
			startingPosition: grid.Point{I: -38, J: 170},
			startingWaypoint:grid.Point{I: 10, J: 4},
			givenInstruction: instruction{key: 'F', value: 11},
			expectedPosition: grid.Point{I: 72, J: 214},
			expectedWaypoint: grid.Point{I: 10, J: 4},
		},
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'L', value: 90},
			expectedPosition: grid.Point{I: 0, J: 0},
			expectedWaypoint: grid.Point{I: -10, J: -1},
		},
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'L', value: 180},
			expectedPosition: grid.Point{I: 0, J: 0},
			expectedWaypoint: grid.Point{I: 1, J: -10},
		},
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'L', value: 270},
			expectedPosition: grid.Point{I: 0, J: 0},
			expectedWaypoint: grid.Point{I: 10, J: 1},
		},
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'R', value: 90},
			expectedPosition: grid.Point{I: 0, J: 0},
			expectedWaypoint: grid.Point{I: 10, J: 1},
		},
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'R', value: 180},
			expectedPosition: grid.Point{I: 0, J: 0},
			expectedWaypoint: grid.Point{I: 1, J: -10},
		},
		{
			startingPosition: grid.Point{I: 0, J: 0},
			startingWaypoint: grid.Point{I: -1, J: 10},
			givenInstruction: instruction{key: 'R', value: 270},
			expectedPosition: grid.Point{I: 0, J: 0},
			expectedWaypoint: grid.Point{I: -10, J: -1},
		},
		{
			startingPosition: grid.Point{I: 199, J: 806},
			startingWaypoint: grid.Point{I: -6, J: -4},
			givenInstruction: instruction{key: 'R', value: 90},
			expectedPosition: grid.Point{I: 199, J: 806},
			expectedWaypoint: grid.Point{I: -4, J: 6},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf(
			"ship starting at position %v and waypoint %v, taking instruction %v, moves to position %v and waypoint %v",
		tc.startingPosition, tc.startingWaypoint, tc.givenInstruction, tc.expectedPosition, tc.expectedWaypoint)
		t.Run(name, func(t *testing.T){
			// Arrange.
			s := newGuidedShip(tc.startingPosition, tc.startingWaypoint)

			// Act.
			err := s.TakeInstruction(tc.givenInstruction.key, tc.givenInstruction.value)

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedPosition, s.Position())
			assert.Equal(t, tc.expectedWaypoint, s.Waypoint())	
		})
	}
}
