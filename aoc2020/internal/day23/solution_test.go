package day23

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleStartOrder = []int{
	3, 8, 9, 1, 2, 5, 4, 6, 7,
}

func TestCupRingMove(t *testing.T) {
	testCases := []struct {
		nMoves         int
		expectedOutput string
	}{
		{
			nMoves:         0,
			expectedOutput: "25467389",
		},
		{
			nMoves:         1,
			expectedOutput: "54673289",
		},
		{
			nMoves:         2,
			expectedOutput: "32546789",
		},
		{
			nMoves:         3,
			expectedOutput: "34672589",
		},
		{
			nMoves:         4,
			expectedOutput: "32584679",
		},
		{
			nMoves:         5,
			expectedOutput: "36792584",
		},
		{
			nMoves:         6,
			expectedOutput: "93672584",
		},
		{
			nMoves:         7,
			expectedOutput: "92583674",
		},
		{
			nMoves:         8,
			expectedOutput: "58392674",
		},
		{
			nMoves:         9,
			expectedOutput: "83926574",
		},
		{
			nMoves:         10,
			expectedOutput: "92658374",
		},
		{
			nMoves:         100,
			expectedOutput: "67384529",
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d moves results in output %s", tc.nMoves, tc.expectedOutput)
		t.Run(name, func(t *testing.T){
			// Arrange.
			ring := newCupRing2(sampleStartOrder, len(sampleStartOrder))

			// Act.
			for i:=0;i<tc.nMoves;i++ {
				ring.MakeMove()
			}

			// Assert.
			assert.Equal(t, tc.expectedOutput, ring.ReadCupsStartingAfter(1))
		})
	}

}
