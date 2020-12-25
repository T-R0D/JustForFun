package day05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardingIDFromString(t *testing.T) {
	testCases := []struct {
		inputID        string
		expectedRow    int
		expectedColumn int
		expectedSeatID int
	}{
		{
			inputID:        "BFFFBBFRRR",
			expectedRow:    70,
			expectedColumn: 7,
			expectedSeatID: 567,
		},
		{
			inputID:        "FFFBBBFRRR",
			expectedRow:    14,
			expectedColumn: 7,
			expectedSeatID: 119,
		},
		{
			inputID:        "BBFFBBFRLL",
			expectedRow:    102,
			expectedColumn: 4,
			expectedSeatID: 820,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s has seat ID %d", tc.inputID, tc.expectedSeatID), func(t *testing.T) {
			// Arrange.

			// Act.
			bID, err := boardingIDFromString(tc.inputID)
			seatID := bID.SeatID()

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedRow, bID.row)
			assert.Equal(t, tc.expectedColumn, bID.column)
			assert.Equal(t, tc.expectedSeatID, seatID)
		})
	}
}
