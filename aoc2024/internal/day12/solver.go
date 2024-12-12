package day12

import (
	"fmt"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	gardenMap := parseGardenMap(input)

	regions := mapRegions(gardenMap)
	fenceCost := findCostToFenceGarden(regions)

	return fmt.Sprintf("%d", fenceCost), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	gardenMap := parseGardenMap(input)

	regions := mapRegions(gardenMap)
	fenceCost := findCostToFenceGardenWithBulkDiscount(regions)

	return fmt.Sprintf("%d", fenceCost), nil
}

func parseGardenMap(input string) [][]rune {
	rows := strings.Split(input, "\n")
	gardenMap := make([][]rune, 0, len(rows))
	for _, row := range rows {
		gardenRow := make([]rune, len(row))
		for j, plot := range row {
			gardenRow[j] = plot
		}
		gardenMap = append(gardenMap, gardenRow)
	}

	return gardenMap
}

type locationSet map[[2]int]struct{}

func mapRegions(gardenMap [][]rune) []locationSet {
	regions := []locationSet{}
	for i, row := range gardenMap {
	NEXT_PLOT:
		for j, crop := range row {
			currentLocation := [2]int{i, j}
			for _, region := range regions {
				if _, ok := region[currentLocation]; ok {
					continue NEXT_PLOT
				}
			}

			newRegion := mapRegion(gardenMap, currentLocation, crop)
			regions = append(regions, newRegion)
		}
	}

	return regions
}

func mapRegion(gardenMap [][]rune, start [2]int, crop rune) locationSet {
	frontier := queue.NewFifo[[2]int]()
	frontier.Push(start)
	seen := locationSet{}
	deltas := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	region := locationSet{}

	for frontier.Len() > 0 {
		currentLocation, ok := frontier.Pop()
		if !ok {
			return region
		}

		if currentLocation[0] < 0 || len(gardenMap) <= currentLocation[0] ||
			currentLocation[1] < 0 || len(gardenMap[0]) <= currentLocation[1] {

			continue
		}

		if _, ok := seen[currentLocation]; ok {
			continue
		}

		if gardenMap[currentLocation[0]][currentLocation[1]] != crop {
			continue
		}

		for _, delta := range deltas {
			nextLocation := [2]int{currentLocation[0] + delta[0], currentLocation[1] + delta[1]}
			frontier.Push(nextLocation)
		}

		region[currentLocation] = struct{}{}
		seen[currentLocation] = struct{}{}
	}

	return region
}

func findCostToFenceGarden(regions []locationSet) int {
	cost := 0
	for _, region := range regions {
		cost += fenceCost(region)
	}

	return cost
}

func findCostToFenceGardenWithBulkDiscount(regions []locationSet) int {
	cost := 0
	for _, region := range regions {
		cost += fenceCostWithBulkDiscount(region)
	}

	return cost
}

func fenceCost(region locationSet) int {
	return regionArea(region) * regionPerimeter(region)
}

func fenceCostWithBulkDiscount(region locationSet) int {
	return regionArea(region) * regionSides(region)
}

func regionArea(region locationSet) int {
	return len(region)
}

func regionPerimeter(region locationSet) int {
	deltas := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	nExposed := 0
	for location := range region {
		for _, delta := range deltas {
			neighborLocation := [2]int{location[0] + delta[0], location[1] + delta[1]}
			if _, ok := region[neighborLocation]; !ok {
				nExposed += 1
			}
		}
	}

	return nExposed
}

func regionSides(region locationSet) int {
	locationsWithTopOpen := [][2]int{}
	for location := range region {
		topNeighbor := [2]int{location[0] - 1, location[1]}
		if _, ok := region[topNeighbor]; !ok {
			locationsWithTopOpen = append(locationsWithTopOpen, location)
		}
	}

	nSides := 0
	for len(locationsWithTopOpen) > 0 {
		start := locationsWithTopOpen[len(locationsWithTopOpen)-1]

		newSides, coveredCandidates := leftHandOnWallWalkToCountSides(region, start)

		nSides += newSides

		remainingLocationsWithTopOpen := make([][2]int, 0, len(locationsWithTopOpen) - len(coveredCandidates) + 1)
		for _, location := range locationsWithTopOpen {
			if _, ok := coveredCandidates[location]; !ok {
				remainingLocationsWithTopOpen = append(remainingLocationsWithTopOpen, location)
			}
		}
		locationsWithTopOpen = remainingLocationsWithTopOpen
	}

	return nSides
}

func leftHandOnWallWalkToCountSides(region locationSet, start [2]int) (int, locationSet) {
	deltas := map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'>': {0, 1},
		'<': {0, -1},
	}
	leftHandDirection := map[rune]rune{
		'>': '^',
		'v': '>',
		'<': 'v',
		'^': '<',
	}
	rightHandDirection := map[rune]rune{
		'>': 'v',
		'v': '<',
		'<': '^',
		'^': '>',
	}

	nSides := 0
	coveredLocationsWithTopOpen := locationSet{}
	currentLocation := start
	orientation := '>'
	firstStep := true

WALK:
	for !(currentLocation == start && orientation == '>') || firstStep {
		firstStep = false

		if orientation == '>' {
			coveredLocationsWithTopOpen[currentLocation] = struct{}{}
		}

		forwardDelta := deltas[orientation]
		forwardLocation := [2]int{currentLocation[0] + forwardDelta[0], currentLocation[1] + forwardDelta[1]}

		if _, forwardAvailable := region[forwardLocation]; !forwardAvailable {
			if orientation == '^' {
				coveredLocationsWithTopOpen[currentLocation] = struct{}{}
			}

			orientation = rightHandDirection[orientation]
			nSides += 1
			continue WALK
		}

		forwardAndLeftDelta := deltas[leftHandDirection[orientation]]
		forwardAndLeftLocation := [2]int{forwardLocation[0] + forwardAndLeftDelta[0], forwardLocation[1] + forwardAndLeftDelta[1]}
		if _, forwardAndLeftAvailable := region[forwardAndLeftLocation]; forwardAndLeftAvailable {
			orientation = leftHandDirection[orientation]
			currentLocation = forwardAndLeftLocation
			nSides += 1
			continue WALK
		}

		currentLocation = forwardLocation
	}

	return nSides, coveredLocationsWithTopOpen
}
