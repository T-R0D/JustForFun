package day03

import (
	"fmt"
	"regexp"
	"strconv"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	corruptProgram := parseCorruptedProgram(input)

	intactInstructions, err := parseIntactMulInstructions(corruptProgram)
	if err != nil {
		return "", err
	}

	sum := executeAndSumInstructions(intactInstructions)

	return strconv.Itoa(sum), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	corruptProgram := parseCorruptedProgram(input)

	intactInstructions, err := parseIntactMulInstructionsWithConditionals(corruptProgram)
	if err != nil {
		return "", err
	}

	sum := executeAndSumInstructions(intactInstructions)

	return strconv.Itoa(sum), nil
}

func parseCorruptedProgram(input string) string {
	return input
}

func parseIntactMulInstructions(program string) ([]mul, error) {
	mulInstructionRegex, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return []mul{}, err
	}

	matches := mulInstructionRegex.FindAllStringSubmatch(program, -1)
	if matches == nil {
		return []mul{}, nil
	}

	intactInstructions := make([]mul, 0, len(matches))
	for i, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			return []mul{},
				fmt.Errorf(
					"uncorrupted mul instruction %d ('%s'), operand A could not be parsed: %w",
					i,
					match[0],
					err)
		}

		b, err := strconv.Atoi(match[2])
		if err != nil {
			return []mul{},
				fmt.Errorf(
					"uncorrupted mul instruction %d ('%s'), operand B could not be parsed: %w",
					i,
					match[0],
					err)
		}

		intactInstructions = append(intactInstructions, mul{enabled: true, a: a, b: b})
	}

	return intactInstructions, nil
}

func parseIntactMulInstructionsWithConditionals(program string) ([]mul, error) {
	mulInstructionRegex, err := regexp.Compile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return []mul{}, err
	}

	matches := mulInstructionRegex.FindAllStringSubmatch(program, -1)
	if matches == nil {
		return []mul{}, nil
	}

	intactInstructions := make([]mul, 0, len(matches))
	enabled := true
	for i, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				return []mul{},
					fmt.Errorf(
						"uncorrupted mul instruction %d ('%s'), operand A could not be parsed: %w",
						i,
						match[0],
						err)
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				return []mul{},
					fmt.Errorf(
						"uncorrupted mul instruction %d ('%s'), operand B could not be parsed: %w",
						i,
						match[0],
						err)
			}

			intactInstructions = append(intactInstructions, mul{enabled: enabled, a: a, b: b})
		}
	}

	return intactInstructions, nil
}

type mul struct {
	enabled bool
	a       int
	b       int
}

func (m *mul) Execute() int {
	if !m.enabled {
		return 0
	}

	return m.a * m.b
}

func executeAndSumInstructions(instructions []mul) int {
	sum := 0
	for _, instruction := range instructions {
		sum += instruction.Execute()
	}

	return sum
}
