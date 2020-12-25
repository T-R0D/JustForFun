package main

import (
	"aoc2019/internal/day01"
	"aoc2019/internal/day02"
	"aoc2019/internal/day03"
	"aoc2019/internal/day04"
	"aoc2019/internal/day05"
	"aoc2019/internal/day06"
	"aoc2019/internal/day07"
	"aoc2019/internal/day08"
	"aoc2019/internal/day09"
	"aoc2019/internal/day10"
	"aoc2019/internal/day11"
	"aoc2019/internal/day12"
	"aoc2019/internal/day13"
	"aoc2019/internal/day14"
	"aoc2019/internal/day15"
	"aoc2019/internal/day16"
	"aoc2019/internal/day17"
	"aoc2019/internal/day18"
	"aoc2019/internal/day19"
	"aoc2019/internal/day20"
	"aoc2019/internal/day21"
	"aoc2019/internal/day22"
	"aoc2019/internal/day23"
	"aoc2019/internal/day24"
	"aoc2019/internal/day25"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type solver interface {
	SolvePart1(string) (interface{}, error)
	SolvePart2(string) (interface{}, error)
}

type params struct {
	Day       uint
	Part      uint
	InputPath string
}

func main() {
	params, err := parseFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Advent of Code 2019")
	solution, err := run(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Day %d part %d solution: %v\n", params.Day, params.Part, solution)
}

func parseFlags() (*params, error) {
	day := flag.Uint("day", 0, "The number of the day's solution to run (1-25).")
	part := flag.Uint("part", 0, "The part of the day's solution to run (1|2).")
	inputPath := flag.String("input", "", "The path of the file containing the input for the specified problem.")
	flag.Parse()

	if !(1 <= *day && *day <= 25) {
		return nil, fmt.Errorf("Invalid day - please enter a day whose solution has been implemented (1-25)")
	}

	if !(1 == *part || 2 == *part) {
		return nil, fmt.Errorf("Invalid part - please enter a value for part (1 or 2).")
	}

	if "" == *inputPath {
		return nil, fmt.Errorf("Input file not specified - please provide an input file.")
	}

	return &params{
		Day:       *day,
		Part:      *part,
		InputPath: *inputPath,
	}, nil
}

func run(p *params) (interface{}, error) {
	input, err := readInput(p.InputPath)
	if err != nil {
		return nil, err
	}

	var s solver
	switch p.Day {
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
		s = &day20.Solver{
			InputPath: p.InputPath,
		}
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
		return nil, fmt.Errorf("Solution not implemented for day %d", p.Day)
	}

	return runSolution(s, p, input)
}

func readInput(inputPath string) (string, error) {
	bytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return "", err
	}

	input := string(bytes)

	return strings.TrimSpace(input), nil
}

func runSolution(s solver, p *params, input string) (interface{}, error) {
	if 1 == p.Part {
		return s.SolvePart1(input)
	} else if 2 == p.Part {
		return s.SolvePart2(input)
	} else {
		return nil, fmt.Errorf("Invalid part - %d is not 1 or 2.", p.Part)
	}
}
