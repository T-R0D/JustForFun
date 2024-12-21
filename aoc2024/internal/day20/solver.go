package day20

import (
	"fmt"
	"math"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	track := parseRacetrack(input)

	nPathsThatSignificantlyBeatBaseline := countPathsWithSignificantGain(
		track, oldCheatDuration, significantImprovementPicoSeconds)

	return fmt.Sprintf("%d", nPathsThatSignificantlyBeatBaseline), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	track := parseRacetrack(input)

	nPathsThatSignificantlyBeatBaseline := countPathsWithSignificantGain(
		track, newCheatDuration, significantImprovementPicoSeconds)

	return fmt.Sprintf("%d", nPathsThatSignificantlyBeatBaseline), nil
}

const (
	oldCheatDuration = 2
	newCheatDuration = 20

	significantImprovementPicoSeconds = 100
)

type vec2 [2]int

func (v vec2) Plus(that vec2) vec2 {
	return vec2{v[0] + that[0], v[1] + that[1]}
}

type racetrack struct {
	Grid  [][]rune
	M     int
	N     int
	Start vec2
	End   vec2
}

func parseRacetrack(input string) *racetrack {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, 0, len(rows))
	start := vec2{}
	end := vec2{}
	for i, row := range rows {
		gridRow := make([]rune, len(row))
		for j, sprite := range row {
			gridRow[j] = sprite

			if sprite == 'S' {
				start = vec2{i, j}
			} else if sprite == 'E' {
				end = vec2{i, j}
			}
		}
		grid = append(grid, gridRow)
	}

	return &racetrack{
		Grid:  grid,
		M:     len(grid),
		N:     len(grid[0]),
		Start: start,
		End:   end,
	}
}

func countPathsWithSignificantGain(
	track *racetrack, cheatDuration int, significantImprovement int) int {

	baseline := findPathThroughRacetrack(track)

	lengthCounts := findCheatingPathLengthsThroughRacetrack(
		track, baseline, cheatDuration)

	nPathsThatImproveOnBaselineSignificantly := 0
	for pathLength, count := range lengthCounts {
		if len(baseline)-pathLength >= significantImprovement {
			nPathsThatImproveOnBaselineSignificantly += count
		}
	}

	return nPathsThatImproveOnBaselineSignificantly
}

func findPathThroughRacetrack(track *racetrack) []vec2 {
	type searchState struct {
		Location vec2
		Path     []vec2
	}

	frontier := queue.NewFifo[*searchState]()
	frontier.Push(&searchState{
		Location: track.Start,
		Path:     []vec2{track.Start},
	})

	type searchEssence struct {
		Location vec2
	}
	seen := map[searchEssence]struct{}{}

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

		if state.Location[0] < 0 || track.M <= state.Location[0] ||
			state.Location[1] < 0 || track.N <= state.Location[1] {

			continue
		}

		essence := searchEssence{
			Location: state.Location,
		}
		if _, ok := seen[essence]; ok {
			continue
		}
		seen[essence] = struct{}{}

		if track.Grid[state.Location[0]][state.Location[1]] == '#' {
			continue
		}

		if state.Location == track.End {
			return state.Path
		}

		for _, delta := range deltas {
			newPosition := state.Location.Plus(delta)
			newPath := append([]vec2{}, state.Path...)
			newPath = append(newPath, newPosition)
			frontier.Push(&searchState{
				Location: newPosition,
				Path:     newPath,
			})
		}
	}

	return []vec2{}
}

func findCheatingPathLengthsThroughRacetrack(track *racetrack, nonCheatPath []vec2, cheatDuration int) map[int]int {
	stepsRemaining := map[vec2]int{}
	for i, location := range nonCheatPath {
		stepsRemaining[location] = len(nonCheatPath) - i - 1
	}

	pathLengths := map[int]int{}

	for stepsBeforeCheat, cheatStartLocation := range nonCheatPath {
		for i := -cheatDuration; i <= cheatDuration; i += 1 {
			for j := -cheatDuration; j <= cheatDuration; j += 1 {
				currentLocation := cheatStartLocation.Plus(vec2{i, j})

				if currentLocation[0] < 0 || track.M <= currentLocation[0] ||
					currentLocation[1] < 0 || track.N <= currentLocation[1] {

					continue
				}

				if i == 0 && j == 0 {
					continue
				}

				cheatSteps := int(math.Abs(float64(i)) + math.Abs(float64(j)))
				if cheatSteps > cheatDuration {
					continue
				}

				if stepsToGo, ok := stepsRemaining[currentLocation]; ok {
					totalSteps := stepsBeforeCheat + cheatSteps + stepsToGo

					if count, ok := pathLengths[totalSteps]; ok {
						pathLengths[totalSteps] = count + 1
					} else {
						pathLengths[totalSteps] = 1
					}
				}
			}
		}
	}

	return pathLengths
}
