package day24

import (
	"fmt"
	"testing"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/stretchr/testify/assert"
)

const sampleInput = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestTranslateDirectionSequenceToLoc(t *testing.T) {
	testCases := []struct {
		directionSequence []direction
		expectedLocation  grid.Point
	}{
		{
			directionSequence: []direction{directionEast, directionSouthEast, directionWest},
			expectedLocation:  grid.Point{I: 1, J: -1},
		},
		{
			directionSequence: []direction{directionNorthWest, directionWest, directionSouthWest, directionEast, directionEast},
			expectedLocation:  grid.Point{I: 0, J: 0},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%v leads to %v", tc.directionSequence, tc.expectedLocation)
		t.Run(name, func(t *testing.T) {
			loc := translateDirectionSequenceToLoc(tc.directionSequence)

			assert.Equal(t, tc.expectedLocation, loc)
		})
	}
}

func TestLivingTileFloor(t *testing.T) {
	testCases := []struct {
		nDays               int
		expectedNBlackTiles int
	}{
		{
			nDays:               0,
			expectedNBlackTiles: 10,
		},
		{
			nDays:               1,
			expectedNBlackTiles: 15,
		},
		{
			nDays:               10,
			expectedNBlackTiles: 37,
		},
		{
			nDays:               100,
			expectedNBlackTiles: 2208,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("after %d days, there should be %d black tiles", tc.nDays, tc.expectedNBlackTiles)
		t.Run(name, func(t *testing.T) {
			directionSequences, err := parseInputLines(sampleInput)
			assert.NoError(t, err)

			tiling := newHexagonTiling()

			for _, sequence := range directionSequences {
				tiling.FlipTile(sequence)
			}

			for day := 0; day < tc.nDays; day++ {
				tiling.ChangeTilesForNextDay()
			}

			assert.Equal(t, tc.expectedNBlackTiles, tiling.BlackTiles())
		})
	}
}
