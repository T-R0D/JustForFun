package day04

import (
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/queue"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	grid := parseGrid(input)

	accessibleRolls := countAccessiblePaperRolls(grid)

	return strconv.Itoa(accessibleRolls), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	runeGrid := parseGrid(input)

	countGrid := gridToNeighborCountsGrid(runeGrid)

	removableRolls := countRemovablePaperRollsCascading(countGrid)

	return strconv.Itoa(removableRolls), nil
}

const (
	empty     = '.'
	paperRoll = '@'
)

const (
	emptyNeighborCell     = -1
	blockingNeighborCount = 4
)

var deltas = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func parseGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		row := []rune(line)
		grid = append(grid, row)
	}

	return grid
}

func countAccessiblePaperRolls(grid [][]rune) int {
	accessibleRolls := 0
	for i := range len(grid) {
	RowSearch:
		for j := range len(grid[i]) {
			if grid[i][j] != paperRoll {
				continue RowSearch
			}

			adjacentRolls := 0

		CandidateSearch:
			for _, delta := range deltas {
				y, x := i+delta[0], j+delta[1]
				if y < 0 || len(grid) <= y || x < 0 || len(grid[i]) <= x {
					continue CandidateSearch
				}

				if grid[y][x] == paperRoll {
					adjacentRolls += 1
				}
			}

			if adjacentRolls < blockingNeighborCount {
				accessibleRolls += 1
			}
		}
	}

	return accessibleRolls
}

func gridToNeighborCountsGrid(grid [][]rune) [][]int {
	countGrid := make([][]int, 0, len(grid))
	for i, row := range grid {
		countRow := make([]int, 0, len(row))
	RowScan:
		for j, el := range row {
			if el == empty {
				countRow = append(countRow, -1)
				continue RowScan
			}

			neighbors := 0
		NeighborScan:
			for _, delta := range deltas {
				y, x := i+delta[0], j+delta[1]
				if y < 0 || len(grid) <= y || x < 0 || len(grid[i]) <= x {
					continue NeighborScan
				}

				if grid[y][x] == paperRoll {
					neighbors += 1
				}
			}

			countRow = append(countRow, neighbors)
		}
		countGrid = append(countGrid, countRow)
	}

	return countGrid
}

func countRemovablePaperRollsCascading(countGrid [][]int) int {
	grid := make([][]int, len(countGrid))
	for i, row := range countGrid {
		grid[i] = make([]int, len(row))
		copy(grid[i], row)
	}

	toVisit := queue.NewLifo[[]int]()
	for i, row := range grid {
		for j, count := range row {
			if count != emptyNeighborCell && count < blockingNeighborCount {
				toVisit.Push([]int{i, j})
			}
		}
	}

	removed := 0
RemovalScan:
	for toVisit.Len() > 0 {
		target, ok := toVisit.Pop()
		if !ok {
			continue RemovalScan
		}

		i, j := target[0], target[1]

		if grid[i][j] == emptyNeighborCell {
			continue RemovalScan
		}

		grid[i][j] = emptyNeighborCell
		removed += 1

	NeighborScan:
		for _, delta := range deltas {
			y, x := i+delta[0], j+delta[1]
			if y < 0 || len(grid) <= y || x < 0 || len(grid[i]) <= x {
				continue NeighborScan
			}

			if grid[y][x] == emptyNeighborCell {
				continue NeighborScan
			}

			grid[y][x] -= 1

			if grid[y][x] < blockingNeighborCount {
				toVisit.Push([]int{y, x})
			}
		}
	}

	return removed
}
