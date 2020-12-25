package day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testMap = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestTreeMapCountsTreesOnSlope(t *testing.T) {
	testCases := []struct {
		rise          int
		run           int
		expectedTrees int
	}{
		{
			rise:          1,
			run:           1,
			expectedTrees: 2,
		},
		{
			rise:          1,
			run:           3,
			expectedTrees: 7,
		},
		{
			rise:          1,
			run:           5,
			expectedTrees: 3,
		},
		{
			rise:          1,
			run:           7,
			expectedTrees: 4,
		},
		{
			rise:          2,
			run:           1,
			expectedTrees: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("rise=%d run=%d crosses %d trees", tc.rise, tc.run, tc.expectedTrees), func(t *testing.T) {
			// Arrange.
			treemap, err := newTreeMap(testMap)
			assert.NoError(t, err)

			// Act.
			treesOnSlope := treemap.CountTreesOnSlope(tc.rise, tc.run)

			// Assert.
			assert.Equal(t, tc.expectedTrees, treesOnSlope)
		})
	}
}
