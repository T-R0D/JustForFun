package day03

import (
	"strconv"

	"github.com/T-R0D/aoc2020/internal/grid"

	"github.com/pkg/errors"
)

type Solver struct{}

func (s *Solver) Part1(input string) (string, error) {
	treemap, err := newTreeMap(input)
	if err != nil {
		return "", err
	}

	nTrees := treemap.CountTreesOnSlope(1, 3)

	return strconv.Itoa(nTrees), nil
}

func (s *Solver) Part2(input string) (string, error) {
	treemap, err := newTreeMap(input)
	if err != nil {
		return "", err
	}

	slopesToTry := []struct {
		rise int
		run  int
	}{
		{
			rise: 1,
			run:  1,
		},
		{
			rise: 1,
			run:  3,
		},
		{
			rise: 1,
			run:  5,
		},
		{
			rise: 1,
			run:  7,
		},
		{
			rise: 2,
			run:  1,
		},
	}

	treeProduct := 1

	for _, slope := range slopesToTry {
		treeProduct *= treemap.CountTreesOnSlope(slope.rise, slope.run)
	}

	return strconv.Itoa(treeProduct), nil
}

const (
	empty   = '.'
	tree    = '#'
	newline = '\n'

	run1  = 3
	rise1 = 1
)

type treeMap struct {
	Height int
	Trees  map[grid.Point]bool
	Width  int
}

func newTreeMap(input string) (*treeMap, error) {
	trees := map[grid.Point]bool{}
	i, j := 0, 0
	for _, r := range input {
		switch r {
		case tree:
			trees[grid.Point{I: i, J: j}] = true
			j += 1
		case empty:
			j += 1
		case newline:
			j = 0
			i += 1
		default:
			return nil, errors.Errorf("Invalid character encountered in map: %v", r)
		}
	}

	return &treeMap{
		Height: i + 1,
		Trees:  trees,
		Width:  j,
	}, nil
}

func (t *treeMap) CountTreesOnSlope(rise int, run int) int {
	nTrees := 0
	i, j := 0, 0

	for i < t.Height {
		if treePresent, ok := t.Trees[grid.Point{I: i, J: j}]; ok && treePresent {
			nTrees += 1
		}
		j = (j + run) % t.Width
		i += rise
	}

	return nTrees
}
