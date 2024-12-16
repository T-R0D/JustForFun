package day16

import (
	"fmt"
	"math"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	maze := parseReindeerMaze(input)

	lowestScoreAndPaths, err := findLowestScoringPathsThroughMaze(maze)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", lowestScoreAndPaths.Score), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	maze := parseReindeerMaze(input)

	lowestScorePathUnion, err := findLocationsOnAllLowestScorePaths(maze)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", len(lowestScorePathUnion)), nil
}

type reindeerMaze struct {
	Grid  [][]rune
	M     int
	N     int
	Start vec2
	End   vec2
}

type vec2 [2]int

type pathAction rune

const (
	rotateClockwise        pathAction = ')'
	rotateCounterClockwise pathAction = '('
	moveUp                 pathAction = '^'
	moveDown               pathAction = 'v'
	moveLeft               pathAction = '<'
	moveRight              pathAction = '>'
)

func parseReindeerMaze(input string) reindeerMaze {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0, len(lines))
	var start vec2
	var end vec2
	for i, line := range lines {
		row := make([]rune, len(line))
		for j, sprite := range line {
			row[j] = sprite

			if sprite == 'S' {
				start = vec2{i, j}
			} else if sprite == 'E' {
				end = vec2{i, j}
			}
		}
		grid = append(grid, row)
	}

	return reindeerMaze{
		Grid:  grid,
		M:     len(grid),
		N:     len(grid[0]),
		Start: start,
		End:   end,
	}
}

type lowestScoreAndPath struct {
	Score int
	Path  []vec2
}

func findLowestScoringPathsThroughMaze(maze reindeerMaze) (lowestScoreAndPath, error) {
	const (
		movePenalty = 1
		turnPenalty = 1000
	)

	deltas := map[pathAction]vec2{
		moveUp:    {-1, 0},
		moveDown:  {1, 0},
		moveRight: {0, 1},
		moveLeft:  {0, -1},
	}

	clockwiseTurns := map[pathAction]pathAction{
		moveUp:    moveRight,
		moveRight: moveDown,
		moveDown:  moveLeft,
		moveLeft:  moveUp,
	}

	counterClockwiseTurns := map[pathAction]pathAction{
		moveUp:    moveLeft,
		moveLeft:  moveDown,
		moveDown:  moveRight,
		moveRight: moveUp,
	}

	turnMaps := []map[pathAction]pathAction{clockwiseTurns, counterClockwiseTurns}

	type navigationState struct {
		Location    vec2
		Orientation pathAction
		ScoreSoFar  int
		Path        []vec2
	}

	cmpNavigationStateForLowerScore := func(a navigationState, b navigationState) bool {
		return a.ScoreSoFar < b.ScoreSoFar
	}

	frontier := queue.NewPriority(cmpNavigationStateForLowerScore)
	frontier.Push(navigationState{
		Location:    maze.Start,
		Orientation: moveRight,
		ScoreSoFar:  0,
		Path:        []vec2{maze.Start},
	})

	type navigationStateEssence struct {
		Location    vec2
		Orientation pathAction
	}
	seen := map[navigationStateEssence]struct{}{}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			return lowestScoreAndPath{}, fmt.Errorf("somehow, there were no states to process")
		}

		essence := navigationStateEssence{
			Location:    state.Location,
			Orientation: state.Orientation,
		}
		if _, ok := seen[essence]; ok {
			continue
		}

		if state.Location[0] < 0 || maze.M <= state.Location[0] ||
			state.Location[1] < 0 || maze.N <= state.Location[1] {

			continue
		}

		sprite := maze.Grid[state.Location[0]][state.Location[1]]

		if sprite == '#' {
			continue
		}

		if sprite == 'E' {
			return lowestScoreAndPath{
				Score: state.ScoreSoFar,
				Path:  state.Path,
			}, nil
		}

		delta := deltas[state.Orientation]
		newLocation := vec2{state.Location[0] + delta[0], state.Location[1] + delta[1]}
		newPath := append([]vec2{}, state.Path...)
		newPath = append(newPath, newLocation)
		frontier.Push(navigationState{
			Location:    newLocation,
			Orientation: state.Orientation,
			ScoreSoFar:  state.ScoreSoFar + movePenalty,
			Path:        newPath,
		})

		for _, turnMap := range turnMaps {
			turn := turnMap[state.Orientation]
			frontier.Push(navigationState{
				Location:    state.Location,
				Orientation: turn,
				ScoreSoFar:  state.ScoreSoFar + turnPenalty,
				Path:        state.Path,
			})
		}

		seen[essence] = struct{}{}
	}

	return lowestScoreAndPath{}, fmt.Errorf("a path through the maze could not be found")
}

func findLocationsOnAllLowestScorePaths(maze reindeerMaze) (map[vec2]struct{}, error) {
	scoreGrid := make([][]int, 0, maze.M)
	for range maze.Grid {
		newRow := make([]int, maze.N)
		for j := range maze.N {
			newRow[j] = math.MaxInt
		}
		scoreGrid = append(scoreGrid, newRow)
	}

	const (
		movePenalty = 1
		turnPenalty = 1000
	)

	deltas := map[pathAction]vec2{
		moveUp:    {-1, 0},
		moveDown:  {1, 0},
		moveRight: {0, 1},
		moveLeft:  {0, -1},
	}

	clockwiseTurns := map[pathAction]pathAction{
		moveUp:    moveRight,
		moveRight: moveDown,
		moveDown:  moveLeft,
		moveLeft:  moveUp,
	}

	counterClockwiseTurns := map[pathAction]pathAction{
		moveUp:    moveLeft,
		moveLeft:  moveDown,
		moveDown:  moveRight,
		moveRight: moveUp,
	}

	turnMaps := []map[pathAction]pathAction{clockwiseTurns, counterClockwiseTurns}

	type navigationState struct {
		Location        vec2
		Orientation     pathAction
		ScoreSoFar      int
		Path            []vec2
		LastMoveWasTurn bool
	}

	cmpNavigationStateForLowerScore := func(a navigationState, b navigationState) bool {
		return a.ScoreSoFar < b.ScoreSoFar
	}

	frontier := queue.NewPriority(cmpNavigationStateForLowerScore)
	frontier.Push(navigationState{
		Location:        maze.Start,
		Orientation:     moveRight,
		ScoreSoFar:      0,
		Path:            []vec2{maze.Start},
		LastMoveWasTurn: false,
	})

	union := map[vec2]struct{}{}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			return map[vec2]struct{}{}, fmt.Errorf("somehow we didn't get a state to process")
		}

		if state.Location[0] < 0 || maze.M <= state.Location[0] ||
			state.Location[1] < 0 || maze.N <= state.Location[1] {

			continue
		}

		sprite := maze.Grid[state.Location[0]][state.Location[1]]

		if sprite == '#' {
			continue
		}

		currentScore := scoreGrid[state.Location[0]][state.Location[1]]

		if state.ScoreSoFar < currentScore {
			scoreGrid[state.Location[0]][state.Location[1]] = state.ScoreSoFar
		} else if state.ScoreSoFar-turnPenalty > currentScore && !state.LastMoveWasTurn {
			continue
		}

		if sprite == 'E' && state.ScoreSoFar == scoreGrid[state.Location[0]][state.Location[1]] {
			for _, location := range state.Path {
				union[location] = struct{}{}
			}
			continue
		}

		delta := deltas[state.Orientation]
		newLocation := vec2{state.Location[0] + delta[0], state.Location[1] + delta[1]}
		newPath := append([]vec2{}, state.Path...)
		newPath = append(newPath, newLocation)
		frontier.Push(navigationState{
			Location:        newLocation,
			Orientation:     state.Orientation,
			ScoreSoFar:      state.ScoreSoFar + movePenalty,
			Path:            newPath,
			LastMoveWasTurn: false,
		})

		if !state.LastMoveWasTurn {
			for _, turnMap := range turnMaps {
				turn := turnMap[state.Orientation]
				frontier.Push(navigationState{
					Location:        state.Location,
					Orientation:     turn,
					ScoreSoFar:      state.ScoreSoFar + turnPenalty,
					Path:            state.Path,
					LastMoveWasTurn: true,
				})
			}
		}
	}

	return union, nil
}
