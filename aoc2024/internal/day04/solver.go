package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	wordSearch := parseWordSearch(input)

	nFound := countAllInWordSearch(wordSearch, []rune("XMAS"))

	return strconv.Itoa(nFound), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	wordSearch := parseWordSearch(input)

	nFound, err := countAllXTargetsInWordSearch(wordSearch, []rune("MAS"))

	return strconv.Itoa(nFound), err
}

func parseWordSearch(input string) [][]rune {
	lines := strings.Split(input, "\n")
	wordSearch := make([][]rune, len(lines))
	for i, line := range lines {
		wordSearch[i] = []rune(line)
	}
	return wordSearch
}

func countAllInWordSearch(wordSearch [][]rune, target []rune) int {
	deltas := [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	nFound := 0
	for i, row := range wordSearch {
		for j := range row {
			for _, delta := range deltas {
				if wordExists(wordSearch, target, []int{i, j}, delta) {
					nFound += 1
				}
			}
		}
	}

	return nFound
}

func wordExists(wordSearch [][]rune, target []rune, origin []int, delta []int) bool {
	for k, r := range target {
		i, j := origin[0]+k*delta[0], origin[1]+k*delta[1]

		if i < 0 || len(wordSearch) <= i || j < 0 || len(wordSearch[0]) <= j {
			return false
		}

		if wordSearch[i][j] != r {
			return false
		}
	}

	return true
}

func countAllXTargetsInWordSearch(wordSearch [][]rune, target []rune) (int, error) {
	nFound := 0
	for i, row := range wordSearch {
		for j := range row {
			if found, err := xTargetExists(wordSearch, target, []int{i, j}); err != nil {
				return 0, err
			} else if found {
				nFound += 1
			}
		}
	}

	return nFound, nil
}

func xTargetExists(wordSearch [][]rune, target []rune, origin []int) (bool, error) {
	if len(target)%2 != 1 {
		return false, fmt.Errorf("'%s' is not an odd length, and therefore has no middle to 'cross' over", string(target))
	}

	deltas := [][]int{
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}

	targetMidIndex := (len(target) / 2)

	nFound := 0
CROSS_SEARCH:
	for _, delta := range deltas {
		for k := range (len(target) / 2) + 1 {
			i1, j1 := origin[0]+k*delta[0], origin[1]+k*delta[1]
			i2, j2 := origin[0]+-k*delta[0], origin[1]+-k*delta[1]

			if i1 < 0 || len(wordSearch) <= i1 || j1 < 0 || len(wordSearch[0]) <= j1 {
				return false, nil
			}

			if i2 < 0 || len(wordSearch) <= i2 || j2 < 0 || len(wordSearch[0]) <= j2 {
				return false, nil
			}

			if wordSearch[i1][j1] != target[targetMidIndex-k] || wordSearch[i2][j2] != target[targetMidIndex+k] {
				continue CROSS_SEARCH
			}
		}

		nFound += 1
	}

	return nFound >= 2, nil
}
