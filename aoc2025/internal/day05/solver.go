package day05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	dbContents, err := parseDBContents(input)
	if err != nil {
		return "", err
	}

	condensedFreshRanges := mergeRanges(dbContents.freshRanges)

	nFreshIngredients := 0
	for _, id := range dbContents.availableIngredients {
		if valueInRanges(condensedFreshRanges, id) {
			nFreshIngredients += 1
		}
	}

	return strconv.Itoa(nFreshIngredients), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	dbContents, err := parseDBContents(input)
	if err != nil {
		return "", err
	}

	condensedFreshRanges := mergeRanges(dbContents.freshRanges)

	totalFreshIngredientIDs := sumRangeSpans(condensedFreshRanges)

	return strconv.Itoa(totalFreshIngredientIDs), nil
}

type inventoryDB struct {
	freshRanges          [][]int
	availableIngredients []int
}

func parseDBContents(input string) (inventoryDB, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return inventoryDB{}, fmt.Errorf("invalid number of sections: %d", len(sections))
	}

	idRangeStrs := strings.Split(sections[0], "\n")
	freshRanges := make([][]int, 0, len(idRangeStrs))
	for _, idRangeStr := range idRangeStrs {
		parts := strings.Split(idRangeStr, "-")
		if len(parts) != 2 {
			return inventoryDB{}, fmt.Errorf("invalid number of range parts: %d", len(parts))
		}

		idRange := make([]int, 0, len(parts))
		for _, part := range parts {
			id, err := strconv.Atoi(part)
			if err != nil {
				return inventoryDB{}, err
			}

			idRange = append(idRange, id)
		}

		freshRanges = append(freshRanges, idRange)
	}

	ingredientIDStrs := strings.Split(sections[1], "\n")
	availableIngredients := make([]int, 0, len(ingredientIDStrs))
	for _, idStr := range ingredientIDStrs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return inventoryDB{}, err
		}

		availableIngredients = append(availableIngredients, id)
	}

	return inventoryDB{freshRanges, availableIngredients}, nil
}

func mergeRanges(originalRanges [][]int) [][]int {
	if len(originalRanges) < 1 {
		return originalRanges
	}

	ranges := make([][]int, 0, len(originalRanges))
	ranges = append(ranges, originalRanges...)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	condensedRanges := [][]int{}
	currentLower := ranges[0][0]
	currentUpper := ranges[0][1]
	for _, r := range ranges {
		if currentUpper < r[0] {
			condensedRanges = append(condensedRanges, []int{currentLower, currentUpper})

			currentLower, currentUpper = r[0], r[1]
			continue
		}

		if currentUpper < r[1] {
			currentUpper = r[1]
		}
	}
	condensedRanges = append(condensedRanges, []int{currentLower, currentUpper})

	return condensedRanges
}

func valueInRanges(condensedRanges [][]int, value int) bool {
	if len(condensedRanges) < 1 {
		return false
	}

	for _, r := range condensedRanges {
		if r[0] <= value && value <= r[1] {
			return true
		}
	}

	return false
}

func sumRangeSpans(condensedRanges [][]int) int {
	sum := 0
	for _, r := range condensedRanges {
		sum += r[1] - r[0] + 1
	}

	return sum
}
