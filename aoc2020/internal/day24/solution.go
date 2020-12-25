package day24

import (
	"fmt"
	"strings"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/pkg/errors"
)

// Solver saves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	directionSequences, err := parseInputLines(input)
	if err != nil {
		return "", err
	}

	tiling := newHexagonTiling()

	for _, directionSequence := range directionSequences {
		tiling.FlipTile(directionSequence)
	}

	return fmt.Sprint(tiling.BlackTiles()), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	directionSequences, err := parseInputLines(input)
	if err != nil {
		return "", err
	}

	tiling := newHexagonTiling()

	for _, directionSequence := range directionSequences {
		tiling.FlipTile(directionSequence)
	}

	for day := 0; day < oneHundredDays; day++ {
		tiling.ChangeTilesForNextDay()
	}

	return fmt.Sprint(tiling.BlackTiles()), nil
}

func parseInputLines(input string) ([][]direction, error) {
	lines := strings.Split(input, "\n")
	directionSequences := make([][]direction, len(lines))
	for i, line := range lines {
		sequence, err := parseLineIntoSequence(line)
		if err != nil {
			return nil, errors.Wrapf(err, "processing line %d", i)
		}
		directionSequences[i] = sequence
	}
	return directionSequences, nil
}

func parseLineIntoSequence(line string) ([]direction, error) {
	sequence := make([]direction, 0, len(line))
	buildingANortherlyDirection := false
	buildingASoutherlyDirection := false
	for i, r := range line {
		switch r {
		case 'e':
			switch {
			case buildingANortherlyDirection:
				sequence = append(sequence, directionNorthEast)
			case buildingASoutherlyDirection:
				sequence = append(sequence, directionSouthEast)
			default:
				sequence = append(sequence, directionEast)
			}
			buildingANortherlyDirection = false
			buildingASoutherlyDirection = false
		case 'n':
			if buildingANortherlyDirection || buildingASoutherlyDirection {
				return nil, errors.Errorf("undexpected 'n' at %d", i)
			}
			buildingANortherlyDirection = true
		case 's':
			if buildingANortherlyDirection || buildingASoutherlyDirection {
				return nil, errors.Errorf("undexpected 's' at %d", i)
			}
			buildingASoutherlyDirection = true
		case 'w':
			switch {
			case buildingANortherlyDirection:
				sequence = append(sequence, directionNorthWest)
			case buildingASoutherlyDirection:
				sequence = append(sequence, directionSouthWest)
			default:
				sequence = append(sequence, directionWest)
			}
			buildingANortherlyDirection = false
			buildingASoutherlyDirection = false
		default:
			return nil, errors.Errorf("unexpected charachter '%c' at position %d", r, i)
		}
	}
	return sequence, nil
}

const (
	oneHundredDays = 100
)

type direction string

const (
	directionEast      direction = "e"
	directionNorthEast direction = "ne"
	directionNorthWest direction = "nw"
	directionSouthEast direction = "se"
	directionSouthWest direction = "sw"
	directionWest      direction = "w"
)

var tileVectors map[direction]grid.Point = map[direction]grid.Point{
	directionEast:      {I: 2, J: 0},
	directionNorthEast: {I: 1, J: 1},
	directionNorthWest: {I: -1, J: 1},
	directionSouthEast: {I: 1, J: -1},
	directionSouthWest: {I: -1, J: -1},
	directionWest:      {I: -2, J: 0},
}

type tileColor string

const (
	tileColorBlack tileColor = "black"
	tileColorWhite tileColor = "white"
)

type hexagonTiling struct {
	tiles map[grid.Point]tileColor
}

func newHexagonTiling() *hexagonTiling {
	return &hexagonTiling{
		tiles: map[grid.Point]tileColor{},
	}
}

func (t *hexagonTiling) GetTileColor(loc grid.Point) tileColor {
	if color, ok := t.tiles[loc]; ok {
		return color
	}
	return tileColorWhite
}

func (t *hexagonTiling) FlipTile(directionSequence []direction) {
	loc := translateDirectionSequenceToLoc(directionSequence)
	if color, ok := t.tiles[loc]; ok {
		if color == tileColorBlack {
			t.tiles[loc] = tileColorWhite
		} else {
			t.tiles[loc] = tileColorBlack
		}
	} else {
		t.tiles[loc] = tileColorBlack
	}
}

func (t *hexagonTiling) BlackTiles() int {
	nBlack := 0
	for _, color := range t.tiles {
		if color == tileColorBlack {
			nBlack++
		}
	}
	return nBlack
}

func (t *hexagonTiling) ChangeTilesForNextDay() {
	newTiles := map[grid.Point]tileColor{}

	for loc, color := range t.tiles {
		if color == tileColorBlack {
			neighbors := t.neighbors(loc)
			nBlack := 0
			for _, neighbor := range neighbors {
				if neighborColor, ok := t.tiles[neighbor]; ok && neighborColor == tileColorBlack {
					nBlack++
				}
			}

			if nBlack == 1 || nBlack == 2 {
				newTiles[loc] = tileColorBlack
			}
		}
	}

	whiteNeighbors := t.whiteNeighbors()
	for loc := range whiteNeighbors {
		neighbors := t.neighbors(loc)
		nBlack := 0
		for _, neighbor := range neighbors {
			if neighborColor, ok := t.tiles[neighbor]; ok && neighborColor == tileColorBlack {
				nBlack++
			}
		}

		if nBlack == 2 {
			newTiles[loc] = tileColorBlack
		}
	}

	t.tiles = newTiles
}

func (t *hexagonTiling) whiteNeighbors() map[grid.Point]struct{} {
	whiteNeighbors := map[grid.Point]struct{}{}
	for loc, color := range t.tiles {
		if color == tileColorBlack {
			neighbors := t.neighbors(loc)
			for _, neighbor := range neighbors {
				if color, ok := t.tiles[neighbor]; !ok || color == tileColorWhite {
					whiteNeighbors[neighbor] = struct{}{}
				}
			}
		}
	}
	return whiteNeighbors
}

func (t *hexagonTiling) neighbors(loc grid.Point) []grid.Point {
	neighbors := make([]grid.Point, 0, 6)
	for _, vector := range tileVectors {
		neighbor := grid.Point{I: loc.I + vector.I, J: loc.J + vector.J}
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func translateDirectionSequenceToLoc(directionSequence []direction) grid.Point {
	loc := grid.Point{I: 0, J: 0}
	for _, dir := range directionSequence {
		delta := tileVectors[dir]
		loc = grid.Point{I: loc.I + delta.I, J: loc.J + delta.J}
	}
	return loc
}
