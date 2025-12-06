package day06

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	problems, err := parseWorksheetForHumanArithmetic(input)
	if err != nil {
		return "", err
	}

	checksum := 0
	for _, problem := range problems {
		checksum += problem.Solve()
	}

	return strconv.Itoa(checksum), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	problems, err := parseWorksheetForCephalopodArithmetic(input)
	if err != nil {
		return "", err
	}

	checksum := 0
	for _, problem := range problems {
		checksum += problem.Solve()
	}

	return strconv.Itoa(checksum), nil
}

type cephalopodOperator rune

const (
	cephalopodAdd cephalopodOperator = '+'
	cephalopodMul cephalopodOperator = '*'
)

type mathProblem struct {
	operands []int
	operator cephalopodOperator
}

func (this *mathProblem) Solve() int {
	if this.operator == cephalopodAdd {
		sum := 0
		for _, val := range this.operands {
			sum += val
		}
		return sum
	} else {
		product := 1
		for _, val := range this.operands {
			product *= val
		}
		return product
	}
}

func parseWorksheetForHumanArithmetic(input string) ([]mathProblem, error) {
	lines := strings.Split(input, "\n")
	if len(lines) < 1 {
		return []mathProblem{}, nil
	}

	operators := strings.Fields(lines[len(lines)-1])

	problems := make([]mathProblem, len(operators))
	for j, line := range lines[:len(lines)-1] {
		for i, value := range strings.Fields(line) {
			val, err := strconv.Atoi(value)
			if err != nil {
				return []mathProblem{}, fmt.Errorf("(%d, %d) %w", i, j, err)
			}

			problems[i].operands = append(problems[i].operands, val)
			problems[i].operator = cephalopodOperator(operators[i][0])
		}
	}

	return problems, nil
}

func parseWorksheetForCephalopodArithmetic(input string) ([]mathProblem, error) {
	runeGrid := worksheetToRuneGrid(input)
	if len(runeGrid) < 1 {
		return []mathProblem{}, nil
	}

	nOperands := len(runeGrid) - 1
	nProblems := len(strings.Fields(string(runeGrid[len(runeGrid)-1])))

	problems := make([]mathProblem, 0, nProblems)

	currentProblem := mathProblem{}
NextColumnToTheLeft:
	for j := len(runeGrid[0]) - 1; j >= 0; j -= 1 {
		digits := []rune{}
		onlySpacesEncountered := true
	NextRow:
		for i := 0; i < nOperands; i += 1 {
			r := runeGrid[i][j]

			if r == ' ' {
				continue NextRow
			}

			onlySpacesEncountered = false

			if unicode.IsDigit(r) {
				digits = append(digits, r)
			} else {
				return []mathProblem{}, fmt.Errorf("unexpected rune: '%c'", r)
			}
		}

		if onlySpacesEncountered {
			continue NextColumnToTheLeft
		}

		operand, err := runeSliceToInt(digits)
		if err != nil {
			return []mathProblem{}, err
		}
		currentProblem.operands = append(currentProblem.operands, operand)

		r := runeGrid[len(runeGrid)-1][j]
		if r == rune(cephalopodAdd) || r == rune(cephalopodMul) {
			currentProblem.operator = cephalopodOperator(r)
			problems = append(problems, currentProblem)

			currentProblem = mathProblem{}
		}
	}

	return problems, nil
}

func worksheetToRuneGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")

	runeGrid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		runeGrid = append(runeGrid, []rune(line))
	}

	return runeGrid
}

func runeSliceToInt(slice []rune) (int, error) {
	return strconv.Atoi(string(slice))
}
