package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/T-R0D/aoc2020/internal/day01"
	"github.com/T-R0D/aoc2020/internal/day02"
	"github.com/T-R0D/aoc2020/internal/day03"
	"github.com/T-R0D/aoc2020/internal/day04"
	"github.com/T-R0D/aoc2020/internal/day05"
	"github.com/T-R0D/aoc2020/internal/day06"
	"github.com/T-R0D/aoc2020/internal/day07"
	"github.com/T-R0D/aoc2020/internal/day08"
	"github.com/T-R0D/aoc2020/internal/day09"
	"github.com/T-R0D/aoc2020/internal/day10"
	"github.com/T-R0D/aoc2020/internal/day11"
	"github.com/T-R0D/aoc2020/internal/day12"
	"github.com/T-R0D/aoc2020/internal/day13"
	"github.com/T-R0D/aoc2020/internal/day14"
	"github.com/T-R0D/aoc2020/internal/day15"
	"github.com/T-R0D/aoc2020/internal/day16"
	"github.com/T-R0D/aoc2020/internal/day17"
	"github.com/T-R0D/aoc2020/internal/day18"
	"github.com/T-R0D/aoc2020/internal/day19"
	"github.com/T-R0D/aoc2020/internal/day20"
	"github.com/T-R0D/aoc2020/internal/day21"
	"github.com/T-R0D/aoc2020/internal/day22"
	"github.com/T-R0D/aoc2020/internal/day23"
	"github.com/T-R0D/aoc2020/internal/day24"
	"github.com/T-R0D/aoc2020/internal/day25"
	"github.com/T-R0D/aoc2020/internal/input"
	"github.com/pkg/errors"
)

func main() {
	args, err := parseArguments()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := runSolution(args); err != nil {
		fmt.Println(err.Error())
	}
}

const (
	invalidDay uint = 0
	bothParts  uint = 0

	firstDay uint = 1
	lastDay  uint = 25
)

type programArgs struct {
	day       uint
	part      uint
	inputPath string
}

func parseArguments() (programArgs, error) {
	day := flag.Uint("day", 0, "Which day's problem (1-25) should be run.")
	part := flag.Uint("part", 0, "Which part (1-2, 0 for both) for the day to run.")
	inputPath := flag.String("input", "", "The path to the text file with the day's input.")

	flag.Parse()

	args := programArgs{
		day:       *day,
		part:      *part,
		inputPath: *inputPath,
	}

	if args.day < firstDay || lastDay < args.day {
		return args, errors.Errorf("Invalid day specified: %d\n", day)
	} else if 2 < args.part {
		return args, errors.Errorf("Invalid part specified: %d\n", part)
	} else if args.inputPath == "" {
		return args, errors.Errorf("Input data file path not specified.")
	}

	return args, nil
}

type solver interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}

func runSolution(args programArgs) error {
	fmt.Printf("Solving day %d, ", args.day)
	if args.part == bothParts {
		fmt.Println("both parts")
	} else {
		fmt.Printf("part %d\n", args.part)
	}

	input, err := input.ReadInput(args.inputPath)
	if err != nil {
		return errors.Wrap(err, "unable to gather input")
	}

	var s solver
	switch args.day {
	case 1:
		s = &day01.Solver{}
	case 2:
		s = &day02.Solver{}
	case 3:
		s = &day03.Solver{}
	case 4:
		s = &day04.Solver{}
	case 5:
		s = &day05.Solver{}
	case 6:
		s = &day06.Solver{}
	case 7:
		s = &day07.Solver{}
	case 8:
		s = &day08.Solver{}
	case 9:
		s = &day09.Solver{}
	case 10:
		s = &day10.Solver{}
	case 11:
		s = &day11.Solver{}
	case 12:
		s = &day12.Solver{}
	case 13:
		s = &day13.Solver{}
	case 14:
		s = &day14.Solver{}
	case 15:
		s = &day15.Solver{}
	case 16:
		s = &day16.Solver{}
	case 17:
		s = &day17.Solver{}
	case 18:
		s = &day18.Solver{}
	case 19:
		s = &day19.Solver{}
	case 20:
		s = &day20.Solver{}
	case 21:
		s = &day21.Solver{}
	case 22:
		s = &day22.Solver{}
	case 23:
		s = &day23.Solver{}
	case 24:
		s = &day24.Solver{}
	case 25:
		s = &day25.Solver{}
	default:
		return errors.New("that day has not been solved yet")
	}

	if args.part == bothParts || args.part == 1 {
		start := time.Now()
		output, err := s.Part1(input)
		duration := time.Since(start)
		if err != nil {
			return err
		}
		fmt.Println(output)
		fmt.Printf("(duration: %dns = %fs)\n", duration.Nanoseconds(), float64(duration) / float64(time.Second))
	}
	if args.part == bothParts || args.part == 2 {
		start := time.Now()
		output, err := s.Part2(input)
		duration := time.Since(start)
		if err != nil {
			return err
		}
		fmt.Println(output)
		fmt.Printf("(duration: %dns = %fs)\n", duration.Nanoseconds(), float64(duration) / float64(time.Second))
	}

	return nil
}
