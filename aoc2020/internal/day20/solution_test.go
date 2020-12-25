package day20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullImageConstruction(t *testing.T) {
	tiles, err := readTiles(smallExample)
	assert.NoError(t, err)

	fullImage, err := constructFullImage(tiles)
	assert.NoError(t, err)

	assert.Equal(t, uint64(1951), fullImage.TopLeftTileID(), "top left")
	assert.Equal(t, uint64(3079), fullImage.TopRightTileID(), "top right")
	assert.Equal(t, uint64(2971), fullImage.BottomLeftTileID(), "bottom left")
	assert.Equal(t, uint64(1171), fullImage.BottomRightTileID(), "bottom right")

	checkSum := fullImage.TopLeftTileID() * fullImage.TopRightTileID() *
		fullImage.BottomLeftTileID() * fullImage.BottomRightTileID()
	assert.Equal(t, uint64(20899048083289), checkSum)
}

func TestMergeTiles(t *testing.T) {
	tiles, err := readTiles(smallExample)
	assert.NoError(t, err)

	fullImage, err := constructFullImage(tiles)
	assert.NoError(t, err)

	mergedImage := fullImage.MergeTiles()

	assert.Equal(t, mergedSmallExample, mergedImage.String())
}

const smallExample = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

const mergedSmallExample = `.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###
`

func TestImageTileRotate90(t *testing.T) {
	testCases := []struct {
		startPixels    [tileDim][tileDim]rune
		expectedPixels [tileDim][tileDim]rune
	}{
		{
			startPixels: [tileDim][tileDim]rune{
				{'#', '.', '#', '#', '.', '.', '.', '#', '#', '.'},
				{'#', '.', '#', '#', '#', '#', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '#', '.', '.', '#', '#'},
				{'#', '.', '.', '.', '#', '#', '#', '#', '#', '#'},
				{'.', '#', '#', '.', '#', '.', '.', '.', '.', '#'},
				{'.', '#', '#', '#', '.', '#', '#', '#', '#', '#'},
				{'#', '#', '#', '.', '#', '#', '.', '#', '#', '.'},
				{'.', '#', '#', '#', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '#', '.', '.', '#', '.', '#'},
				{'#', '.', '.', '.', '#', '#', '.', '#', '.', '.'},
			},
			expectedPixels: [tileDim][tileDim]rune{
				{'#', '.', '.', '#', '.', '.', '#', '.', '#', '#'},
				{'.', '.', '#', '#', '#', '#', '.', '.', '.', '.'},
				{'.', '#', '#', '#', '#', '#', '.', '.', '#', '#'},
				{'.', '.', '#', '.', '#', '.', '.', '.', '#', '#'},
				{'#', '#', '.', '#', '.', '#', '#', '.', '#', '.'},
				{'#', '.', '.', '#', '#', '.', '#', '#', '#', '.'},
				{'.', '.', '.', '.', '#', '.', '#', '.', '.', '.'},
				{'#', '#', '.', '#', '#', '.', '#', '.', '.', '#'},
				{'.', '.', '#', '#', '#', '.', '#', '#', '.', '#'},
				{'.', '#', '.', '.', '#', '#', '#', '#', '#', '.'},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tile := &imageTile{
				pixels: tc.startPixels,
			}

			newTile := tile.Rotate90()

			assert.Equal(t, tc.expectedPixels, newTile.pixels)
		})
	}
}

func TestImageTileFlipAcrossX(t *testing.T) {
	testCases := []struct {
		startPixels    [tileDim][tileDim]rune
		expectedPixels [tileDim][tileDim]rune
	}{
		{
			startPixels: [tileDim][tileDim]rune{
				{'#', '.', '#', '#', '.', '.', '.', '#', '#', '.'},
				{'#', '.', '#', '#', '#', '#', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '#', '.', '.', '#', '#'},
				{'#', '.', '.', '.', '#', '#', '#', '#', '#', '#'},
				{'.', '#', '#', '.', '#', '.', '.', '.', '.', '#'},
				{'.', '#', '#', '#', '.', '#', '#', '#', '#', '#'},
				{'#', '#', '#', '.', '#', '#', '.', '#', '#', '.'},
				{'.', '#', '#', '#', '.', '.', '.', '.', '#', '.'},
				{'.', '.', '#', '.', '#', '.', '.', '#', '.', '#'},
				{'#', '.', '.', '.', '#', '#', '.', '#', '.', '.'},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tile := &imageTile{
				pixels: tc.startPixels,
			}

			newTile := tile.FlipAcrossXAxis()

			assert.Equal(t, tc.expectedPixels, newTile.pixels)
		})
	}
}
