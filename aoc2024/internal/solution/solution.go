package solution

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/T-R0D/aoc2024/v2/internal/day01"
	"github.com/T-R0D/aoc2024/v2/internal/day02"
	"github.com/T-R0D/aoc2024/v2/internal/day03"
	"github.com/T-R0D/aoc2024/v2/internal/day04"
	"github.com/T-R0D/aoc2024/v2/internal/day05"
	"github.com/T-R0D/aoc2024/v2/internal/day06"
	"github.com/T-R0D/aoc2024/v2/internal/day07"
	"github.com/T-R0D/aoc2024/v2/internal/day08"
	"github.com/T-R0D/aoc2024/v2/internal/day09"
	"github.com/T-R0D/aoc2024/v2/internal/day10"
	"github.com/T-R0D/aoc2024/v2/internal/day11"
	"github.com/T-R0D/aoc2024/v2/internal/day12"
	"github.com/T-R0D/aoc2024/v2/internal/day13"
	"github.com/T-R0D/aoc2024/v2/internal/day14"
	"github.com/T-R0D/aoc2024/v2/internal/day15"
	"github.com/T-R0D/aoc2024/v2/internal/day16"
	"github.com/T-R0D/aoc2024/v2/internal/day17"
	"github.com/T-R0D/aoc2024/v2/internal/day18"
	"github.com/T-R0D/aoc2024/v2/internal/day19"
	"github.com/T-R0D/aoc2024/v2/internal/day20"
	"github.com/T-R0D/aoc2024/v2/internal/day21"
	"github.com/T-R0D/aoc2024/v2/internal/day22"
	"github.com/T-R0D/aoc2024/v2/internal/day23"
	"github.com/T-R0D/aoc2024/v2/internal/day24"
	"github.com/T-R0D/aoc2024/v2/internal/day25"
)

type problemSolver interface {
	SolvePartOne(input string) (string, error)
	SolvePartTwo(input string) (string, error)
}

func Run(inputPath string, day int, part int) error {
	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("reading input file '%s': %w", inputPath, err)
	}

	var solver problemSolver = nil
	switch day {
	case 1:
		solver = &day01.Solver{}
	case 2:
		solver = &day02.Solver{}
	case 3:
		solver = &day03.Solver{}
	case 4:
		solver = &day04.Solver{}
	case 5:
		solver = &day05.Solver{}
	case 6:
		solver = &day06.Solver{}
	case 7:
		solver = &day07.Solver{}
	case 8:
		solver = &day08.Solver{}
	case 9:
		solver = &day09.Solver{}
	case 10:
		solver = &day10.Solver{}
	case 11:
		solver = &day11.Solver{}
	case 12:
		solver = &day12.Solver{}
	case 13:
		solver = &day13.Solver{}
	case 14:
		solver = &day14.Solver{}
	case 15:
		solver = &day15.Solver{}
	case 16:
		solver = &day16.Solver{}
	case 17:
		solver = &day17.Solver{}
	case 18:
		solver = &day18.Solver{}
	case 19:
		solver = &day19.Solver{}
	case 20:
		solver = &day20.Solver{}
	case 21:
		solver = &day21.Solver{}
	case 22:
		solver = &day22.Solver{}
	case 23:
		solver = &day23.Solver{}
	case 24:
		solver = &day24.Solver{}
	case 25:
		solver = &day25.Solver{}
	}

	var solve func(input string) (string, error) = nil
	switch part {
	case 1:
		solve = solver.SolvePartOne
	case 2:
		solve = solver.SolvePartTwo
	}

	input := strings.Trim(string(inputBytes), "\n")
	startTime := time.Now()
	result, err := solve(string(input))
	stopTime := time.Now()
	solveDuration := stopTime.Sub(startTime)

	if err != nil {
		fmt.Printf(
			"unable to solve %d.%d, errored after %dµs: %+v\n",
			day,
			part,
			solveDuration.Microseconds(),
			err)
		return nil
	}

	fmt.Printf("%s\n", result)
	fmt.Printf("(%dµs)\n", solveDuration.Microseconds())

	return nil
}
