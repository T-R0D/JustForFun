package day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryGameSolverRunUntil(t *testing.T) {
	testCases := []struct{
		startingNumbers []int
		until int
		expectedResult int
	}{
		{
			startingNumbers: []int{0,3,6},
			until: 10,
			expectedResult: 0,
		},
		{
			startingNumbers: []int{0,3,6},
			until: 2020,
			expectedResult: 436,
		},
		{
			startingNumbers: []int{1,3,2},
			until: 2020,
			expectedResult: 1,
		},
		{
			startingNumbers: []int{2,1,3},
			until: 2020,
			expectedResult: 10,
		},
		{
			startingNumbers: []int{1,2,3},
			until: 2020,
			expectedResult: 27,
		},
		{
			startingNumbers: []int{2,3,1},
			until: 2020,
			expectedResult: 78,
		},
		{
			startingNumbers: []int{3,2,1},
			until: 2020,
			expectedResult: 438,
		},
		{
			startingNumbers: []int{3,1,2},
			until: 2020,
			expectedResult: 1836,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf(
			"starting with %v, going until %d, results with %d",
			tc.startingNumbers, tc.until, tc.expectedResult)
		t.Run(name, func(t *testing.T){
			// Arrange.
			gameSolver := newMemoryGameSolver()

			// Act.
			gameSolver.Init(tc.startingNumbers)
			result := gameSolver.RunUntil(tc.until)

			// Assert.
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
