// Possible Improvement: I bet there's some number theory or math that can 
//                       speed this up significantly...

package day15

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	startingNumbers, err := parseStartingNumbers(input)
	if err != nil {
		return "", err
	}

	gameSolver := newMemoryGameSolver()
	gameSolver.Init(startingNumbers)

	lastNumberSpoken := gameSolver.RunUntil(magicNumber2020)

	return strconv.Itoa(lastNumberSpoken), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	startingNumbers, err := parseStartingNumbers(input)
	if err != nil {
		return "", err
	}

	gameSolver := newMemoryGameSolver()
	gameSolver.Init(startingNumbers)

	lastNumberSpoken := gameSolver.RunUntil(magicNumber30000000)

	return strconv.Itoa(lastNumberSpoken), nil
}

func parseStartingNumbers(input string) ([]int, error) {
	numStrs := strings.Split(input, ",")
	startingNumbers := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, errors.Wrapf(err, "processing number %d", i)
		}
		startingNumbers[i] = num
	}
	return startingNumbers, nil
}

const (
	magicNumber2020     = 2020
	magicNumber30000000 = 30000000
)

const (
	numberSpokenForTheFirstTime = -1
)

type memoryGameSolver struct {
	lastNumberSpoken int
	spokenNumbers    map[int]spokenRecord
	t                int
}

type spokenRecord struct {
	t         int
	firstTime bool
}

func newMemoryGameSolver() *memoryGameSolver {
	return &memoryGameSolver{
		lastNumberSpoken: numberSpokenForTheFirstTime,
		spokenNumbers:    map[int]spokenRecord{},
		t:                0,
	}
}

func (gs *memoryGameSolver) Init(startingNumbers []int) {
	for i, value := range startingNumbers {
		gs.spokenNumbers[value] = spokenRecord{
			t:         i + 1,
			firstTime: true,
		}
		gs.lastNumberSpoken = value
		gs.t = i + 1
	}
}

func (gs *memoryGameSolver) RunUntil(until int) int {
	for gs.t < until {
		gs.t++

		record, ok := gs.spokenNumbers[gs.lastNumberSpoken]

		nextNumber := -1
		if !ok || record.t == gs.t-1 {
			nextNumber = 0
		} else {
			nextNumber = (gs.t - 1) - record.t
		}

		if !ok {
			gs.spokenNumbers[gs.lastNumberSpoken] = spokenRecord{
				t:         gs.t - 1,
				firstTime: true,
			}
		} else {
			gs.spokenNumbers[gs.lastNumberSpoken] = spokenRecord{
				t:         gs.t - 1,
				firstTime: false,
			}
		}

		gs.lastNumberSpoken = nextNumber
	}

	return gs.lastNumberSpoken
}
