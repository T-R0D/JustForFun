package day01

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	sequence, err := parseDialSequence(input)
	if err != nil {
		return "", err
	}

	currentPosition := dialStart
	timesPointedAtZero := 0
	for _, instr := range sequence {
		if instr.direction == left {
			currentPosition = (currentPosition + dialNotches - instr.clicks) % dialNotches
		} else {
			currentPosition = (currentPosition + instr.clicks) % dialNotches
		}

		if currentPosition == dialZero {
			timesPointedAtZero += 1
		}
	}

	return strconv.Itoa(timesPointedAtZero), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	sequence, err := parseDialSequence(input)
	if err != nil {
		return "", err
	}

	currentPosition := dialStart
	timesPointedAtZero := 0
	for _, instr := range sequence {
		fullRotations, clicks := instr.clicks/dialNotches, instr.clicks%dialNotches

		timesPointedAtZero += fullRotations

		if clicks == 0 {
			continue
		}

		previousPosition := currentPosition

		if instr.direction == left {
			currentPosition = (currentPosition + dialNotches - clicks) % dialNotches

			if (previousPosition < currentPosition || currentPosition == dialZero) && previousPosition != dialZero {
				timesPointedAtZero += 1
			}
		} else {
			currentPosition = (currentPosition + clicks) % dialNotches

			if (previousPosition > currentPosition || currentPosition == dialZero) && previousPosition != dialZero {
				timesPointedAtZero += 1
			}
		}
	}

	return strconv.Itoa(timesPointedAtZero), nil
}

const (
	dialZero    = 0
	dialStart   = 50
	dialNotches = 100
)

type dialDirection rune

const (
	left  dialDirection = 'L'
	right dialDirection = 'R'
)

type instruction struct {
	direction dialDirection
	clicks    int
}

func parseDialSequence(input string) ([]instruction, error) {
	lines := strings.Split(input, "\n")

	instructions := make([]instruction, 0, len(lines))
	for _, line := range lines {
		instr, err := parseLine(line)
		if err != nil {
			return []instruction{}, err
		}
		instructions = append(instructions, instr)
	}

	return instructions, nil
}

func parseLine(line string) (instruction, error) {
	dir := rune(line[0])
	if dir != 'L' && dir != 'R' {
		return instruction{}, fmt.Errorf("%c is not a valid direction", dir)
	}

	clicks, err := strconv.Atoi(line[1:])
	if err != nil {
		return instruction{}, err
	}

	return instruction{
		direction: dialDirection(dir),
		clicks:    clicks,
	}, nil
}
