package day01

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	lists, err := parseLists(input)
	if err != nil {
		return "", err
	}

	totalDistance := findTotalDistanceBetweenLists(lists)

	return strconv.Itoa(totalDistance), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	lists, err := parseLists(input)
	if err != nil {
		return "", err
	}

	similarityScore := findSimilarityScore(lists)

	return strconv.Itoa(similarityScore), nil
}

func parseLists(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	lists := [][]int{
		make([]int, len(lines)),
		make([]int, len(lines)),
	}

	for i, line := range lines {
		listMembers := strings.Fields(line)
		for j := range 2 {
			value, err := strconv.Atoi(listMembers[j])
			if err != nil {
				return nil, fmt.Errorf("unable to parse list value %d-%d: %w", j, i, err)
			}

			lists[j][i] = value
		}
	}

	return lists, nil
}

func findTotalDistanceBetweenLists(lists [][]int) int {
	for _, list := range lists {
		slices.Sort(list)
	}

	totalDistance := 0
	for i := range lists[0] {
		a, b := lists[0][i], lists[1][i]
		if a > b {
			totalDistance += a - b
		} else {
			totalDistance += b - a
		}
	}

	return totalDistance
}

func findSimilarityScore(lists [][]int) int {
	counts := map[int]int{}
	for _, x := range lists[1] {
		if count, ok := counts[x]; ok {
			counts[x] = count + 1
		} else {
			counts[x] = 1
		}
	}

	totalSimilarityScore := 0
	for _, x := range lists[0] {
		if count, ok := counts[x]; ok {
			totalSimilarityScore += x * count
		}
	}

	return totalSimilarityScore
}
