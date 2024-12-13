package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	machines, err := parseArcadeInventory(input)
	if err != nil {
		return "", err
	}

	totalCost := findCostToWinAllPossiblePrizes(
		machines, maxTimesButtonCanBePressedForClosePrizes)

	return strconv.FormatInt(totalCost, 10), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	machines, err := parseArcadeInventory(input)
	if err != nil {
		return "", err
	}

	for i := range machines {
		machines[i].PrizeLocation[0] += measurementError
		machines[i].PrizeLocation[1] += measurementError
	}

	totalCost := findCostToWinAllPossiblePrizes(
		machines, effectivelyInfinitePressLimit)

	return strconv.FormatInt(totalCost, 10), nil
}

const (
	costToPressA = 3
	costToPressB = 1

	maxTimesButtonCanBePressedForClosePrizes = int64(100)
	effectivelyInfinitePressLimit            = int64(999_999_999_999_999)

	measurementError = int64(10_000_000_000_000)
)

type arcadeMachine struct {
	AButtonDeltas [2]int64
	BButtonDeltas [2]int64
	PrizeLocation [2]int64
}

func parseArcadeInventory(input string) ([]arcadeMachine, error) {
	sections := strings.Split(input, "\n\n")
	machines := make([]arcadeMachine, 0, len(sections))

	for i, section := range sections {
		lines := strings.Split(section, "\n")
		if len(lines) != 3 {
			return []arcadeMachine{}, fmt.Errorf("section %d did not have the correct number of lines", i)
		}

		aDeltas, err := parseButtonLine(lines[0], "A")
		if err != nil {
			return []arcadeMachine{}, fmt.Errorf("section %d: %w", i, err)
		}

		bDeltas, err := parseButtonLine(lines[1], "B")
		if err != nil {
			return []arcadeMachine{}, fmt.Errorf("section %d: %w", i, err)
		}

		prizeCoordinates, err := parsePrizeLine(lines[2])
		if err != nil {
			return []arcadeMachine{}, fmt.Errorf("section %d: %w", i, err)
		}

		machines = append(machines, arcadeMachine{
			AButtonDeltas: aDeltas,
			BButtonDeltas: bDeltas,
			PrizeLocation: prizeCoordinates,
		})
	}

	return machines, nil
}

func parseButtonLine(line string, buttonLabel string) ([2]int64, error) {
	deltasStr := strings.ReplaceAll(
		line, fmt.Sprintf("Button %s: ", buttonLabel), "")
	deltaStrs := strings.Split(
		strings.ReplaceAll(
			strings.ReplaceAll(deltasStr, "X+", ""),
			"Y+",
			""),
		", ")

	deltas := [2]int64{}
	for i, deltaStr := range deltaStrs {
		delta, err := strconv.ParseInt(deltaStr, 10, 64)
		if err != nil {
			return [2]int64{},
				fmt.Errorf("delta %d (%s) for button %s was not number; %w", i, deltaStr, buttonLabel, err)
		}
		deltas[i] = delta
	}

	return deltas, nil
}

func parsePrizeLine(line string) ([2]int64, error) {
	coordinateStrs := strings.Split(
		strings.ReplaceAll(
			strings.ReplaceAll(strings.ReplaceAll(line, "Prize: ", ""), "X=", ""),
			"Y=", ""),
		", ")

	coordinates := [2]int64{}
	for i, coordinateStr := range coordinateStrs {
		coordinate, err := strconv.ParseInt(coordinateStr, 10, 64)
		if err != nil {
			return [2]int64{},
				fmt.Errorf("coordinate %d (%s) for prize was not number; %w", i, coordinateStr, err)
		}
		coordinates[i] = coordinate
	}

	return coordinates, nil
}

func findCostToWinAllPossiblePrizes(
	machines []arcadeMachine, maxPressesPerButton int64) int64 {

	totalCost := int64(0)
	for _, machine := range machines {
		if cost, ok := findCostToWinPrize(&machine, maxPressesPerButton); ok {
			totalCost += cost
		}
	}

	return totalCost
}

func findCostToWinPrize(machine *arcadeMachine, maxPressesPerButton int64) (int64, bool) {
	result, ok := integerGaussianElimination(machine)
	if !ok {
		return 0, false
	}

	// With the given inputs, this requirement turned out to be a red herring,
	// but it is left for completeness.
	if result.A > maxPressesPerButton || result.B > maxPressesPerButton {
		return 0, false
	}
	
	return result.A*costToPressA + result.B*costToPressB, true
}

type gaussianEliminationResult struct {
	A int64
	B int64
}

// integerGaussianElimination solves a 2 variable system of equations with an
// integer solution.
// If the system does not have a strictly integer solution, then false
// is returned.
// The function transforms the `machine` into a matrix of the form:
// ```
// x_1  x_2  x_p
// y_1  y_2  y_p  
// ```
// in order to solve the for A and B in the equations
// `A * a_button_x_delta + B * b_button_x_delta = prize_location_x` and
// `A * a_button_y_delta + B * b_button_y_delta = prize_location_y`.
// The solution is achieved in a slightly roundabout way to remain
// integer-friendly. This means that rather than going directly for a strictly
// upper-triangular form, we achieve an upper-triangular form that does not
// have '1' coefficients on the diagonal. This is achieved through finding
// a common multiple of the first coefficient on each row, and then
// eliminating.
// After that, substitutions are used, checking if any integer divisions would
// result in a rounding error.
func integerGaussianElimination(machine *arcadeMachine) (gaussianEliminationResult, bool) {
	matrix := [2][3]int64{
		{machine.AButtonDeltas[0], machine.BButtonDeltas[0], machine.PrizeLocation[0]},
		{machine.AButtonDeltas[1], machine.BButtonDeltas[1], machine.PrizeLocation[1]},
	}

	// Step 1: Scale the rows so that the A's have the same value.
	// Do this by finding a common multiple of x_1 and y_1, scaling the rows,
	// then eliminating a row.
	// I find a common multiple instead of least common multiple because
	// I'm lazy, it is less algorithmically complex, and the inputs allow it
	// (if they were larger, we might need LCM).
	commonMultiple := matrix[0][0] * matrix[1][0]
	for i := range 2{
		multiplier := commonMultiple / matrix[i][0]
		for j := range 3 {
			matrix[i][j] *= multiplier
		}
	}

	// Step 2: Eliminate an A coefficient from the second row.
	for j := range 3 {
		matrix[1][j] -= matrix[0][j]
	}

	// Step 3: Find the value for B (if it is an integer solution).
	if matrix[1][2] % matrix[1][1] != 0 {
		return gaussianEliminationResult{}, false
	}
	b := matrix[1][2] / matrix[1][1]

	// Step 4: Find A through substitution of B (if it is an integer solution).
	firstRowLhs := matrix[0][2] - (matrix[0][1] * b)
	if firstRowLhs % matrix[0][0] != 0 {
		return gaussianEliminationResult{}, false
	}
	a := firstRowLhs / matrix[0][0]

	return gaussianEliminationResult{
		A: a,
		B: b,
	}, true
}
