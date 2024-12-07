package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	equations, err := parseCalibrationEquations(input)
	if err != nil {
		return "", err
	}

	operators := []func(int64, int64) int64{add, multiply}
	possiblyValidSum, err := sumTestValuesOfPossiblyValidEquations(equations, operators)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(possiblyValidSum, 10), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	equations, err := parseCalibrationEquations(input)
	if err != nil {
		return "", err
	}

	operators := []func(int64, int64) int64{add, multiply, concat}
	possiblyValidSum, err := sumTestValuesOfPossiblyValidEquations(equations, operators)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(possiblyValidSum, 10), nil
}

func parseCalibrationEquations(input string) ([]calibrationEquation, error) {
	lines := strings.Split(input, "\n")
	equations := make([]calibrationEquation, 0, len(lines))
	for i, line := range lines {
		sides := strings.Split(line, ": ")
		if len(sides) != 2 {
			return []calibrationEquation{}, fmt.Errorf("line %d could not be split into 2 sides of an equation", i)
		}

		testValue, err := strconv.ParseInt(sides[0], 10, 64)
		if err != nil {
			return []calibrationEquation{}, fmt.Errorf("line %d LHS is not a number; %w", i, err)
		}

		rhsStrings := strings.Fields(sides[1])
		rhsValues := make([]int64, 0, len(rhsStrings))
		for j, rhsString := range rhsStrings {
			rhsValue, err := strconv.ParseInt(rhsString, 10, 64)
			if err != nil {
				return []calibrationEquation{}, fmt.Errorf("line %d, item %d on RHS is not a number; %w", i, j, err)
			}
			rhsValues = append(rhsValues, rhsValue)
		}

		equations = append(equations, calibrationEquation{testValue: testValue, rhsValues: rhsValues})
	}

	return equations, nil
}

type calibrationEquation struct {
	testValue int64
	rhsValues []int64
}

func sumTestValuesOfPossiblyValidEquations(equations []calibrationEquation, operators []func(int64, int64) int64) (int64, error) {
	possiblyValidSum := int64(0)
	for _, equation := range equations {
		if isValid, err := isPossiblyValid(equation, operators); err != nil {
			return 0, err
		} else if isValid {
			possiblyValidSum += equation.testValue
		}
	}

	return possiblyValidSum, nil
}

func isPossiblyValid(equation calibrationEquation, operators []func(int64, int64) int64) (bool, error) {
	frontier := queue.NewLifo[balanceState]()
	initialState := balanceState{
		valueSoFar:        equation.rhsValues[0],
		nextOperandIndex:  1,
		nextOperatorIndex: 0,
	}
	frontier.Push(initialState)

	for frontier.Len() != 0 {
		currentState, ok := frontier.Pop()
		if !ok {
			return false, fmt.Errorf("stack was empty while searching for equation balance")
		}

		allOperandsUsed := currentState.nextOperandIndex == len(equation.rhsValues)

		if allOperandsUsed && currentState.valueSoFar == equation.testValue {
			return true, nil
		}

		if allOperandsUsed {
			continue
		}

		if currentState.nextOperatorIndex < len(operators)-1 {
			nextState := balanceState{
				valueSoFar:        currentState.valueSoFar,
				nextOperandIndex:  currentState.nextOperandIndex,
				nextOperatorIndex: currentState.nextOperatorIndex + 1,
			}
			frontier.Push(nextState)
		}

		operator := operators[currentState.nextOperatorIndex]
		newValueSoFar := operator(currentState.valueSoFar, equation.rhsValues[currentState.nextOperandIndex])
		nextState := balanceState{
			valueSoFar:        newValueSoFar,
			nextOperandIndex:  currentState.nextOperandIndex + 1,
			nextOperatorIndex: 0,
		}
		frontier.Push(nextState)
	}

	return false, nil
}

type balanceState struct {
	valueSoFar        int64
	nextOperandIndex  int
	nextOperatorIndex int
}

func add(a int64, b int64) int64 {
	return a + b
}

func multiply(a int64, b int64) int64 {
	return a * b
}

func concat(a int64, b int64) int64 {
	powerOfTen := int64(1)
	for ; powerOfTen <= b; powerOfTen *= 10 {
	}
	return (a * powerOfTen) + b
}
