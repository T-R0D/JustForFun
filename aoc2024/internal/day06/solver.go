package day06

import (
	"fmt"
	"maps"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	details := parseMapDetails(input)

	path := traceGuardPath(
		details.guardStartLocation,
		details.guardOrientation,
		details.obstacles,
		details.mapHeight,
		details.mapWidth)

	nUniqueLocations := countUniqueLocations(path[:len(path)-1])

	return fmt.Sprintf("%d", nUniqueLocations), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	details := parseMapDetails(input)

	loopCausingObstacleLocations := findLoopCausingObstacleLocations(details)

	return fmt.Sprintf("%d", len(loopCausingObstacleLocations)), nil
}

func parseMapDetails(input string) mapDetails {
	rows := strings.Split(input, "\n")
	mapHeight := len(rows)
	mapWidth := 0
	obstacles := map[[2]int]struct{}{}
	guardStartLocation := [2]int{-1, -1}
	guardOrientation := 'X'
	for i, row := range rows {
		mapWidth = len(row)
		for j, mark := range row {
			switch mark {
			case '#':
				obstacles[[2]int{i, j}] = struct{}{}
			case '^', '>', 'v', '<':
				guardStartLocation = [2]int{i, j}
				guardOrientation = mark
			}
		}
	}

	return mapDetails{
		mapHeight: mapHeight,
		mapWidth: mapWidth,
		obstacles:          obstacles,
		guardStartLocation: guardStartLocation,
		guardOrientation:   guardOrientation,
	}
}

type mapDetails struct {
	mapHeight int
	mapWidth int
	obstacles          map[[2]int]struct{}
	guardStartLocation [2]int
	guardOrientation   rune
}

func traceGuardPath(
	guardStartLocation [2]int,
	guardStartOrientation rune,
	obstacles map[[2]int]struct{},
	mapHeight int,
	mapWidth int) []positionRecord {

	path := []positionRecord{
		{location: guardStartLocation, orientation: guardStartOrientation},
	}
	uniquePositions := map[positionRecord]struct{}{
		{
			location:    guardStartLocation,
			orientation: guardStartOrientation,
		}: {},
	}

	deltas := map[rune][2]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}
	currentPosition := guardStartLocation
	currentOrientation := guardStartOrientation

	for loopFound := false; 0 <= currentPosition[0] && currentPosition[0] < mapHeight &&
		0 <= currentPosition[1] && currentPosition[1] < mapWidth &&
		!loopFound; {

		var candidatePosition [2]int
		startingOrientation := currentOrientation
		for {
			delta := deltas[currentOrientation]
			candidatePosition = [2]int{currentPosition[0] + delta[0], currentPosition[1] + delta[1]}

			if _, ok := obstacles[candidatePosition]; !ok {
				break
			}

			if _, ok := obstacles[candidatePosition]; ok {
				switch currentOrientation {
				case '^':
					currentOrientation = '>'
				case '>':
					currentOrientation = 'v'
				case 'v':
					currentOrientation = '<'
				case '<':
					currentOrientation = '^'
				}
			}

			if currentOrientation == startingOrientation {
				panic("Infinite spin surrounded by obstacles detected... somehow...")
			}
		}

		currentPosition = candidatePosition

		record := positionRecord{
			location:    candidatePosition,
			orientation: currentOrientation,
		}

		_, loopFound = uniquePositions[record]

		path = append(path, record)
		uniquePositions[record] = struct{}{}
	}

	return path
}

type positionRecord struct {
	location    [2]int
	orientation rune
}

func countUniqueLocations(path []positionRecord) int {
	uniqueLocations := map[[2]int]struct{}{}
	for _, record := range path {
		uniqueLocations[record.location] = struct{}{}
	}

	return len(uniqueLocations)
}

func findLoopCausingObstacleLocations(details mapDetails) map[[2]int]struct{} {
	originalPath := traceGuardPath(
		details.guardStartLocation,
		details.guardOrientation,
		details.obstacles,
		details.mapHeight,
		details.mapWidth)
	onMapOriginalPath := originalPath[1 : len(originalPath)-1]

	loopCausingObstacleLocations := map[[2]int]struct{}{}

	for _, candidatePosition := range onMapOriginalPath {
		candidateObstacleLocation := candidatePosition.location

		if _, ok := loopCausingObstacleLocations[candidateObstacleLocation]; ok {
			continue
		}

		updatedObstacles := maps.Clone(details.obstacles)
		updatedObstacles[candidateObstacleLocation] = struct{}{}

		path := traceGuardPath(
			details.guardStartLocation,
			details.guardOrientation,
			updatedObstacles,
			details.mapHeight,
			details.mapWidth)

		lastLocation := path[len(path)-1].location
		if 0 <= lastLocation[0] && lastLocation[0] < details.mapHeight &&
			0 <= lastLocation[1] && lastLocation[1] < details.mapWidth {

			loopCausingObstacleLocations[candidateObstacleLocation] = struct{}{}
		}
	}

	return loopCausingObstacleLocations
}
