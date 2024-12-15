package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	warehouseAndRobotMovements, err := parseWarehouseLayoutAndRobotMovements(input)
	if err != nil {
		return "", err
	}

	movementStrategies := map[direction]movementStrategy{
		up:    checkStraightMoveFeasibility,
		down:  checkStraightMoveFeasibility,
		left:  checkStraightMoveFeasibility,
		right: checkStraightMoveFeasibility,
	}
	finalWarehouse := simulateRobotMovements(
		warehouseAndRobotMovements.Warehouse, warehouseAndRobotMovements.Movements, movementStrategies)
	totalGpsScore := scoreWarehouse(finalWarehouse)

	return strconv.Itoa(totalGpsScore), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	warehouseAndRobotMovements, err := parseWarehouseLayoutAndRobotMovements(input)
	if err != nil {
		return "", err
	}

	warehouse := scaleUpWarehouse(warehouseAndRobotMovements.Warehouse)

	movementStrategies := map[direction]movementStrategy{
		up:    checkBranchedMoveFeasibility,
		down:  checkBranchedMoveFeasibility,
		left:  checkStraightMoveFeasibility,
		right: checkStraightMoveFeasibility,
	}
	finalWarehouse := simulateRobotMovements(warehouse, warehouseAndRobotMovements.Movements, movementStrategies)
	totalGpsScore := scoreWarehouse(finalWarehouse)

	return strconv.Itoa(totalGpsScore), nil
}

type warehouseAndRobotMovements struct {
	Warehouse warehouseLayout
	Movements []direction
}

type warehouseLayout struct {
	Grid          [][]sprite
	M             int
	N             int
	RobotLocation vec2
}

type sprite rune

const (
	robot    sprite = '@'
	box      sprite = 'O'
	boxLeft  sprite = '['
	boxRight sprite = ']'
	empty    sprite = '.'
	wall     sprite = '#'
)

type direction rune

const (
	up    direction = '^'
	down  direction = 'v'
	left  direction = '<'
	right direction = '>'
)

type vec2 [2]int

func (v vec2) plus(that vec2) vec2 {
	return vec2{v[0] + that[0], v[1] + that[1]}
}

type dstSrcPair struct {
	Dst vec2
	Src vec2
}

type movementStrategy func(warehouseLayout, vec2) ([]dstSrcPair, bool)

var deltas = map[direction]vec2{
	up:    {-1, 0},
	down:  {1, 0},
	left:  {0, -1},
	right: {0, 1},
}

func parseWarehouseLayoutAndRobotMovements(input string) (warehouseAndRobotMovements, error) {
	section := strings.Split(input, "\n\n")
	if len(section) != 2 {
		return warehouseAndRobotMovements{}, fmt.Errorf("there were not 2 sections in the input")
	}

	warehouse := parseWarehouseLayout(section[0])
	movements := parseMovements(section[1])

	return warehouseAndRobotMovements{
		Warehouse: warehouse,
		Movements: movements,
	}, nil
}

func parseWarehouseLayout(section string) warehouseLayout {
	rows := strings.Split(section, "\n")
	m, n := len(rows), 0
	grid := make([][]sprite, 0, m)
	var robotLocation vec2
	for i, row := range rows {
		n = len(row)
		for j, sprite := range row {
			if sprite == rune(robot) {
				robotLocation = vec2{i, j}
			}
		}
		grid = append(grid, []sprite(row))
	}

	return warehouseLayout{
		Grid:          grid,
		M:             m,
		N:             n,
		RobotLocation: robotLocation,
	}
}

func parseMovements(section string) []direction {
	return []direction(strings.ReplaceAll(section, "\n", ""))
}

func scaleUpWarehouse(warehouse warehouseLayout) warehouseLayout {
	newN := warehouse.N * 2
	newGrid := make([][]sprite, 0, warehouse.M)
	var newRobotLocation vec2
	for i := range warehouse.M {
		newRow := make([]sprite, 0, newN)
		for j := range warehouse.N {
			s := warehouse.Grid[i][j]
			switch s {
			case wall:
				newRow = append(newRow, wall)
				newRow = append(newRow, wall)
			case box:
				newRow = append(newRow, boxLeft)
				newRow = append(newRow, boxRight)
			case empty:
				newRow = append(newRow, empty)
				newRow = append(newRow, empty)
			case robot:
				newRow = append(newRow, robot)
				newRow = append(newRow, empty)
				newRobotLocation = vec2{i, 2 * j}
			}
		}
		newGrid = append(newGrid, newRow)
	}

	return warehouseLayout{
		Grid:          newGrid,
		M:             warehouse.M,
		N:             newN,
		RobotLocation: newRobotLocation,
	}
}

func simulateRobotMovements(warehouse warehouseLayout, movements []direction, movementStrategies map[direction]movementStrategy) warehouseLayout {
	nextWarehouse := warehouse
	for _, movement := range movements {
		nextWarehouse = simulateRobotMovement(nextWarehouse, movement, movementStrategies[movement])
	}

	return nextWarehouse
}

func simulateRobotMovement(
	warehouse warehouseLayout, movement direction, checkMovement movementStrategy) warehouseLayout {

	delta := deltas[movement]
	nextWarehouse := warehouse
	if objectMoves, ok := checkMovement(warehouse, delta); ok {
		nextGrid := nextWarehouse.Grid
		for k := len(objectMoves) - 1; k >= 0; k -= 1 {
			dst, src := objectMoves[k].Dst, objectMoves[k].Src
			nextGrid[dst[0]][dst[1]] = nextGrid[src[0]][src[1]]
			nextGrid[src[0]][src[1]] = empty
		}

		nextRobotLocation := objectMoves[0].Dst
		nextWarehouse.Grid = nextGrid
		nextWarehouse.RobotLocation = nextRobotLocation
	}

	return nextWarehouse
}

func checkStraightMoveFeasibility(warehouse warehouseLayout, delta vec2) ([]dstSrcPair, bool) {
	origin := warehouse.RobotLocation
	objectMoves := []dstSrcPair{}
	for src, dst := origin, origin.plus(delta); coordinateInBounds(dst, warehouse.M, warehouse.N) &&
		warehouse.Grid[dst[0]][dst[1]] != empty; src, dst = dst, dst.plus(delta) {

		if warehouse.Grid[dst[0]][dst[1]] == wall {
			return []dstSrcPair{}, false
		}

		objectMoves = append(objectMoves, dstSrcPair{Dst: dst, Src: src})
	}

	var firstObjectToMoveLocation vec2
	if len(objectMoves) > 0 {
		firstObjectToMoveLocation = objectMoves[len(objectMoves)-1].Dst
	} else {
		firstObjectToMoveLocation = origin
	}
	candidateSpace := firstObjectToMoveLocation.plus(delta)

	if !(coordinateInBounds(candidateSpace, warehouse.M, warehouse.N) &&
		warehouse.Grid[candidateSpace[0]][candidateSpace[1]] == empty) {

		return []dstSrcPair{}, false
	}

	objectMoves = append(objectMoves, dstSrcPair{Dst: candidateSpace, Src: firstObjectToMoveLocation})

	return objectMoves, true
}

func checkBranchedMoveFeasibility(warehouse warehouseLayout, delta vec2) ([]dstSrcPair, bool) {
	type pushSearchState struct {
		Dst      vec2
		Src      vec2
		IsBranch bool
	}

	frontier := queue.NewFifo[pushSearchState]()
	frontier.Push(pushSearchState{
		Dst: warehouse.RobotLocation.plus(delta),
		Src: warehouse.RobotLocation,
	})
	seen := map[pushSearchState]struct{}{}

	objectMoves := []dstSrcPair{}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			return []dstSrcPair{}, false
		}

		if _, ok := seen[state]; ok {
			continue
		}

		if !coordinateInBounds(state.Dst, warehouse.M, warehouse.N) {
			return []dstSrcPair{}, false
		}

		dstObject := warehouse.Grid[state.Dst[0]][state.Dst[1]]

		if dstObject == wall {
			return []dstSrcPair{}, false
		}

		if dstObject == boxLeft {
			frontier.Push(pushSearchState{
				Dst: state.Dst.plus(delta),
				Src: state.Dst,
			})
			if !state.IsBranch {
				frontier.Push(pushSearchState{
					Dst: state.Dst.plus(vec2{0, 1}).plus(delta),
					Src: state.Dst.plus(vec2{0, 1}),
				})
			}
		} else if dstObject == boxRight {
			frontier.Push(pushSearchState{
				Dst: state.Dst.plus(delta),
				Src: state.Dst,
			})
			if !state.IsBranch {
				frontier.Push(pushSearchState{
					Dst: state.Dst.plus(vec2{0, -1}).plus(delta),
					Src: state.Dst.plus(vec2{0, -1}),
				})
			}
		}

		objectMoves = append(objectMoves, dstSrcPair{Dst: state.Dst, Src: state.Src})

		seen[state] = struct{}{}
	}

	return objectMoves, true
}

func scoreWarehouse(warehouse warehouseLayout) int {
	totalGps := 0
	for i := range warehouse.M {
		for j := range warehouse.N {
			if warehouse.Grid[i][j] == box || warehouse.Grid[i][j] == boxLeft {
				totalGps += (100 * i) + j
			}
		}
	}

	return totalGps
}

func coordinateInBounds(coordinate vec2, m int, n int) bool {
	return 0 <= coordinate[0] && coordinate[0] < m &&
		0 <= coordinate[1] && coordinate[1] < n
}
