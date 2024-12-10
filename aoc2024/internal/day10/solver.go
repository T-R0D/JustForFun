package day10

import (
	"fmt"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	mapWithMarkedTrailheads, err := parseMap(input)
	if err != nil {
		return "", err
	}

	totalScore, _ := scoreAndRateAllTrailheads(mapWithMarkedTrailheads)

	return fmt.Sprintf("%d", totalScore), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	mapWithMarkedTrailheads, err := parseMap(input)
	if err != nil {
		return "", err
	}

	_, totalRating := scoreAndRateAllTrailheads(mapWithMarkedTrailheads)

	return fmt.Sprintf("%d", totalRating), nil
}

type mapWithMarkedTrailheads struct {
	topoMap    [][]int
	trailheads []coordinatePair
}

type coordinatePair [2]int

func parseMap(input string) (mapWithMarkedTrailheads, error) {
	lines := strings.Split(input, "\n")
	topoMap := make([][]int, 0, len(lines))
	trailheads := []coordinatePair{}
	for i, line := range lines {
		mapRow := make([]int, len(line))
		for j, elevationR := range line {
			elevation := int(elevationR - '0')
			if elevation < 0 || 9 < elevation {
				return mapWithMarkedTrailheads{}, fmt.Errorf("location (%d, %d) was not a number (%c)", i, j, elevationR)
			}

			mapRow[j] = elevation
			if elevation == 0 {
				trailheads = append(trailheads, coordinatePair{i, j})
			}
		}

		topoMap = append(topoMap, mapRow)
	}

	return mapWithMarkedTrailheads{topoMap: topoMap, trailheads: trailheads}, nil
}

func scoreAndRateAllTrailheads(mapWithTrailheads mapWithMarkedTrailheads) (int, int) {
	totalScore := 0
	totalRating := 0
	for _, trailhead := range mapWithTrailheads.trailheads {
		score, rating := findTrailheadScoreAndRating(mapWithTrailheads.topoMap, trailhead)
		totalScore += score
		totalRating += rating
	}

	return totalScore, totalRating
}

func findTrailheadScoreAndRating(topoMap [][]int, trailhead coordinatePair) (int, int) {
	type walkState struct {
		location  coordinatePair
		elevation int
	}
	frontier := queue.NewLifo[walkState]()
	frontier.Push(walkState{
		location:  trailhead,
		elevation: 0,
	})

	deltas := []coordinatePair{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	peaks := map[coordinatePair]int{}
	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		if state.elevation == 9 {
			if previousTrails, ok := peaks[state.location]; ok {
				peaks[state.location] = previousTrails + 1
			} else {
				peaks[state.location] = 1
			}
			continue
		}

	NEXT_STATES:
		for _, delta := range deltas {
			i, j := state.location[0]+delta[0], state.location[1]+delta[1]

			if i < 0 || len(topoMap) <= i || j < 0 || len(topoMap[0]) <= j {
				continue NEXT_STATES
			}

			nextElevation := topoMap[i][j]
			if nextElevation-state.elevation != 1 {
				continue NEXT_STATES
			}

			frontier.Push(walkState{
				location:  coordinatePair{i, j},
				elevation: nextElevation,
			})
		}
	}

	rating := 0
	for _, trailsToPeak := range peaks {
		rating += trailsToPeak
	}

	return len(peaks), rating
}
