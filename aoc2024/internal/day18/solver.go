package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	bytefallLocations, err := parseBytefallLocations(input)
	if err != nil {
		return "", err
	}

	corruptedLocations := findCorruptedLocations(bytefallLocations[:bytesPerKilobyte])
	shortestSafePath, ok := findShortestPathThroughMemoryGrid(gridM, gridN, corruptedLocations)
	if !ok {
		return "", fmt.Errorf("could not find a valid path through the memory grid")
	}

	return strconv.Itoa(len(shortestSafePath) - 1), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	bytefallLocations, err := parseBytefallLocations(input)
	if err != nil {
		return "", err
	}

	location, ok := findFirstBytefallToBlockPathToExit(gridM, gridN, bytefallLocations)
	if !ok {
		return "", fmt.Errorf("the path was never blocked")
	}

	return location.Swap().String(), nil
}

const (
	bytesPerKilobyte = 1024

	gridM = 71
	gridN = 71
)

type vec2 [2]int

func (v vec2) Plus(that vec2) vec2 {
	return vec2{v[0] + that[0], v[1] + that[1]}
}

func (v vec2) Swap() vec2 {
	return vec2{v[1], v[0]}
}

func (v vec2) String() string {
	values := []string{}
	for i := range v {
		values = append(values, strconv.Itoa(v[i]))
	}
	return strings.Join(values, ",")
}

func parseBytefallLocations(input string) ([]vec2, error) {
	lines := strings.Split(input, "\n")
	bytefallLocations := make([]vec2, 0, len(lines))
	for i, line := range lines {
		location := vec2{}
		for j, coordinate := range strings.Split(line, ",") {
			value, err := strconv.Atoi(coordinate)
			if err != nil {
				return []vec2{}, fmt.Errorf("line %d, item %d was not a number; %w", i, j, err)
			}
			location[2-j-1] = value
		}
		bytefallLocations = append(bytefallLocations, location)
	}

	return bytefallLocations, nil
}

func findCorruptedLocations(bytefalls []vec2) map[vec2]struct{} {
	corruptedLocations := map[vec2]struct{}{}
	for _, bytefall := range bytefalls {
		corruptedLocations[bytefall] = struct{}{}
	}

	return corruptedLocations
}

func findShortestPathThroughMemoryGrid(m int, n int, corruptedLocations map[vec2]struct{}) ([]vec2, bool) {
	type searchState struct {
		Location vec2
		Path     []vec2
	}

	frontier := queue.NewFifo[searchState]()
	frontier.Push(searchState{Location: vec2{0, 0}, Path: []vec2{{0, 0}}})

	seen := map[vec2]struct{}{}

	goal := vec2{m - 1, n - 1}

	deltas := []vec2{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		if state.Location[0] < 0 || m <= state.Location[0] ||
			state.Location[1] < 0 || n <= state.Location[1] {

			continue
		}

		if _, ok := corruptedLocations[state.Location]; ok {
			continue
		}

		if _, ok := seen[state.Location]; ok {
			continue
		}

		if state.Location == goal {
			return state.Path, true
		}

		for _, delta := range deltas {
			newLocation := vec2{state.Location[0] + delta[0], state.Location[1] + delta[1]}
			newPath := append([]vec2{}, state.Path...)
			newPath = append(newPath, newLocation)
			frontier.Push(searchState{Location: newLocation, Path: newPath})
		}

		seen[state.Location] = struct{}{}
	}

	return []vec2{}, false
}

func findFirstBytefallToBlockPathToExit(m int, n int, bytefallLocations []vec2) (vec2, bool) {
	low, high := 0, len(bytefallLocations)-1

	for low <= high {

		test := (low + high) / 2

		corruptedLocations := map[vec2]struct{}{}

		for _, bytefallLocation := range bytefallLocations[:test] {
			corruptedLocations[bytefallLocation] = struct{}{}
		}

		_, ok := findShortestPathThroughMemoryGrid(m, n, corruptedLocations)

		if low == high && ok {
			return bytefallLocations[test], true
		} else if ok {
			low = test + 1
		} else {
			high = test - 1
		}
	}

	return vec2{0, 0}, false
}
