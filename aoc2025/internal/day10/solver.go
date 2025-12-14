package day10

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/queue"
	"github.com/T-R0D/aoc2025/v2/internal/set"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	configurations, err := parseMachineConfigurations(input)
	if err != nil {
		return "", err
	}

	totalButtonPresses := 0
	for _, configuration := range configurations {
		presses, err := findLengthOfShortestLightEnablingSequenceForMachine(configuration)
		if err != nil {
			return "", err
		}

		totalButtonPresses += presses
	}

	return strconv.Itoa(totalButtonPresses), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	configurations, err := parseMachineConfigurations(input)
	if err != nil {
		return "", err
	}

	totalButtonPresses := 0
	for i, configuration := range configurations {
		presses, err := findLengthOfShortestJoltageConfiguringSequenceForMachineWithLinearAlgebra(configuration)
		if err != nil {
			return "", fmt.Errorf("err processing machine %d: %w", i, err)
		}

		totalButtonPresses += presses
	}

	return strconv.Itoa(totalButtonPresses), nil
}

type machineConfiguration struct {
	nOutputs               int
	targetLightDisplay     int
	buttonWiringSchematics [][]int
	joltageRequirements    []int
}

func parseMachineConfigurations(input string) ([]machineConfiguration, error) {
	configurations := []machineConfiguration{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")

		nOutputs := 0
		targetLightDisplay := 0
		for i, r := range parts[0][1 : len(parts[0])-1] {
			nOutputs += 1
			if r == '#' {
				targetLightDisplay |= 1 << i
			}
		}

		buttonWiringSchematics := [][]int{}
		for _, part := range parts[1 : len(parts)-1] {
			numberStrs := strings.Split(strings.ReplaceAll(strings.ReplaceAll(part, "(", ""), ")", ""), ",")
			schematic := make([]int, 0, len(numberStrs))
			for _, str := range numberStrs {
				val, err := strconv.Atoi(str)
				if err != nil {
					return []machineConfiguration{}, err
				}

				schematic = append(schematic, val)
			}

			buttonWiringSchematics = append(buttonWiringSchematics, schematic)
		}

		joltageRequirements := make([]int, 0, len(buttonWiringSchematics))
		for _, str := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(parts[len(parts)-1], "{", ""), "}", ""), ",") {
			val, err := strconv.Atoi(str)
			if err != nil {
				return []machineConfiguration{}, err
			}

			joltageRequirements = append(joltageRequirements, val)
		}

		configurations = append(configurations, machineConfiguration{nOutputs, targetLightDisplay, buttonWiringSchematics, joltageRequirements})
	}

	return configurations, nil
}

func findLengthOfShortestLightEnablingSequenceForMachine(configuration machineConfiguration) (int, error) {
	type SearchState struct {
		lightDisplay int
		presses      int
	}

	frontier := queue.NewFifo[SearchState]()
	frontier.Push(SearchState{
		lightDisplay: 0,
		presses:      0,
	})

	seen := set.New[int]()

Search:
	for frontier.Len() > 0 {
		currentState, ok := frontier.Pop()
		if !ok {
			return 0, fmt.Errorf("somehow, the frontier had no items")
		}

		if currentState.lightDisplay == configuration.targetLightDisplay {
			return currentState.presses, nil
		}

		if seen.Contains(currentState.lightDisplay) {
			continue Search
		}

		for _, buttonConfig := range configuration.buttonWiringSchematics {
			mask := 0
			for _, lightToggled := range buttonConfig {
				mask |= 1 << lightToggled
			}

			frontier.Push(SearchState{
				lightDisplay: currentState.lightDisplay ^ mask,
				presses:      currentState.presses + 1,
			})
		}

		seen.Add(currentState.lightDisplay)
	}

	return 0, fmt.Errorf("unable to find a successful sequence")
}

func findLengthOfShortestJoltageConfiguringSequenceForMachineWithLinearAlgebra(configuration machineConfiguration) (int, error) {
	m, n := configuration.nOutputs, len(configuration.buttonWiringSchematics)

	a := make([][]int, 0, m)
	for range m {
		row := slices.Repeat([]int{0}, n)
		a = append(a, row)
	}

	b := cloneVector(configuration.joltageRequirements, m)

	buttonPressLimits := findButtonPressLimits(configuration)

	for j, buttonSchematic := range configuration.buttonWiringSchematics {
		for _, i := range buttonSchematic {
			a[i][j] = 1
		}
	}

	reductionResult, err := reducedMatrix(a, b, buttonPressLimits, m, n)
	if err != nil {
		return 0, err
	}

	return getMinRealisticButtonPresses(reductionResult)
}

func findButtonPressLimits(configuration machineConfiguration) []int {
	buttonPressLimits := slices.Repeat([]int{math.MaxInt}, len(configuration.buttonWiringSchematics))

	for i, buttonConfig := range configuration.buttonWiringSchematics {
		minTargetValue := configuration.joltageRequirements[buttonConfig[0]]
		for _, output := range buttonConfig {
			targetOutputValue := configuration.joltageRequirements[output]
			if targetOutputValue < minTargetValue {

			}

		}
		buttonPressLimits[i] = minTargetValue
	}

	return buttonPressLimits
}

type reductionTuple struct {
	A                      [][]int
	B                      []int
	ButtonPressLimits      []int
	M                      int
	N                      int
	FreeVariableStartIndex int
}

func reducedMatrix(a [][]int, b []int, buttonPressLimits []int, m int, n int) (reductionTuple, error) {
RowEchelonFormation:
	for h := 0; h < m && h < n; h += 1 {
	PivotSearch:
		for {
			for i := h; i < m; i += 1 {
				if a[i][h] != 0 {
					if i != h {
						err := swapRow(a, b, m, n, h, i)
						if err != nil {
							return reductionTuple{}, err
						}
					}

					break PivotSearch
				}
			}

			for k := h + 1; k < n; k += 1 {
				for i := h; i < m; i += 1 {
					if a[i][k] != 0 {
						err := swapColumn(a, buttonPressLimits, m, n, h, k)
						if err != nil {
							return reductionTuple{}, err
						}

						err = swapRow(a, b, m, n, h, i)
						if err != nil {
							return reductionTuple{}, err
						}

						break PivotSearch
					}
				}
			}

			break RowEchelonFormation
		}

		for i := h + 1; i < m; i += 1 {
			reduceRow(a, b, m, n, h, i)
		}
	}

	newM := m
FindNewM:
	for i := newM - 1; i >= 0; i -= 1 {
		nonZeroFound := false
		for j := range n {
			if a[i][j] != 0 {
				nonZeroFound = true
			}
		}

		if !nonZeroFound && b[i] != 0 {
			return reductionTuple{}, fmt.Errorf("the system was inconsistent (i.e. 0 = k for some constant k where k != 0)")
		}

		if nonZeroFound {
			break FindNewM
		}

		newM -= 1
	}

	for h := newM - 1; h >= 0; h -= 1 {
		for i := h - 1; i >= 0; i -= 1 {
			reduceRow(a, b, m, n, h, i)
		}
	}

	return reductionTuple{
		A:                      cloneMatrix(a, newM, n),
		B:                      cloneVector(b, newM),
		ButtonPressLimits:      buttonPressLimits,
		M:                      newM,
		N:                      n,
		FreeVariableStartIndex: newM,
	}, nil
}

func cloneMatrix(a0 [][]int, m int, n int) [][]int {
	a := make([][]int, 0, m)
	for i := range m {
		row := cloneVector(a0[i], n)
		a = append(a, row)
	}
	return a
}

func cloneVector(v0 []int, n int) []int {
	v := make([]int, n)
	copy(v, v0)
	return v
}

func swapRow(a [][]int, b []int, m int, n int, i1 int, i2 int) error {
	if i1 < 0 || m <= i1 {
		return fmt.Errorf("i1 (%d) outside the bounds of m (%d)", i1, m)
	}

	if i2 < 0 || m <= i2 {
		return fmt.Errorf("i2 (%d) outside the bounds of m (%d)", i2, m)
	}

	for j := range n {
		a[i1][j], a[i2][j] = a[i2][j], a[i1][j]
	}

	b[i1], b[i2] = b[i2], b[i1]

	return nil
}

func swapColumn(a [][]int, buttonPressLimits []int, m int, n int, j1 int, j2 int) error {
	if j1 < 0 || n <= j1 {
		return fmt.Errorf("j1 (%d) outside the bounds of n (%d)", j1, n)
	}

	if j2 < 0 || n <= j2 {
		return fmt.Errorf("j2 (%d) outside the bounds of n (%d)", j2, n)
	}

	if j1 == j2 {
		return nil
	}

	for i := range m {
		a[i][j1], a[i][j2] = a[i][j2], a[i][j1]
	}

	buttonPressLimits[j1], buttonPressLimits[j2] = buttonPressLimits[j2], buttonPressLimits[j1]

	return nil
}

func reduceRow(a [][]int, b []int, m int, n int, h int, i int) error {
	if i < 0 || m <= i {
		return fmt.Errorf("0 < %d && %d <= %d is not true", i, m, i)
	}

	if a[h][h] == 0 {
		return fmt.Errorf("can't reduce a row with pivot of 0")
	}

	if a[i][h] == 0 {
		return nil
	}

	x := a[h][h]
	y := -1 * a[i][h]
	d := greatestCommonDivisor(x, y)
	for j := 0; j < n; j += 1 {
		a[i][j] = ((y * a[h][j]) + (x * a[i][j])) / d
	}
	b[i] = ((y * b[h]) + (x * b[i])) / d

	return nil
}

func getMinRealisticButtonPresses(reduction reductionTuple) (int, error) {
	solutionParameters, err := getFreeVariableColumnsAndPivots(reduction)
	if err != nil {
		return 0, err
	}

	realWorldSolution, err := searchForMinimalRealisticSolution(solutionParameters, reduction.ButtonPressLimits)

	if err != nil {
		return 0, err
	}

	return sumVector(realWorldSolution), nil
}

type parameterizedSolutionTuple struct {
	Constants           []int
	FreeVariableColumns [][]int
	Pivots              []int
	N                   int
	NFree               int
}

func getFreeVariableColumnsAndPivots(reduction reductionTuple) (parameterizedSolutionTuple, error) {
	constants := slices.Repeat([]int{0}, reduction.N)
	pivots := slices.Repeat([]int{-1}, reduction.N)
	for i := range reduction.M {
		pivots[i] = reduction.A[i][i]
		constants[i] = reduction.B[i]
	}

	freeVariableColumns := make([][]int, 0, reduction.N)
	nFree := 0
	for j := reduction.FreeVariableStartIndex; j < reduction.N; j += 1 {
		column := make([]int, reduction.N)
		for i := range reduction.M {
			column[i] = reduction.A[i][j]
		}
		for i := reduction.M; i < reduction.N; i += 1 {
			if i-reduction.M == nFree {
				column[i] = 1
			} else {
				column[i] = 0
			}
		}

		freeVariableColumns = append(freeVariableColumns, column)
		nFree += 1
	}

	return parameterizedSolutionTuple{
		Constants:           constants,
		FreeVariableColumns: freeVariableColumns,
		Pivots:              pivots,
		N:                   reduction.N,
		NFree:               nFree,
	}, nil
}

func searchForMinimalRealisticSolution(solutionParameters parameterizedSolutionTuple, buttonPressLimits []int) ([]int, error) {
	initialCombination := slices.Repeat([]int{0}, solutionParameters.NFree)

	frontier := queue.NewFifo[[]int]()
	frontier.Push(initialCombination)

	seen := set.New[string]()

	solution := slices.Repeat([]int{0}, solutionParameters.N)
	minPressesRequired := math.MaxInt

	Search:
	for frontier.Len() > 0 {
		combination, ok := frontier.Pop()
		if !ok {
			return []int{}, fmt.Errorf("somehow, there was nothing int the frontier")
		}

		combinationString := vectorToString(combination)
		if seen.Contains(combinationString) {
			continue Search
		}

		for i, val := range combination {
			if val > buttonPressLimits[solutionParameters.N-solutionParameters.NFree+i] {
				continue Search
			}
		}

		s, ok := computeSolution(solutionParameters, combination)
		if totalPresses := sumVector(s); ok && totalPresses < minPressesRequired {
			minPressesRequired, solution = totalPresses, s
		}

		for i := range solutionParameters.NFree {
			nextCombination := cloneVector(combination, len(combination))
			nextCombination[i] += 1
			frontier.Push(nextCombination)
		}

		seen.Add(combinationString)
	}

	return solution, nil
}

func computeSolution(solutionParameters parameterizedSolutionTuple, combination []int) ([]int, bool) {
	s := slices.Repeat([]int{0}, solutionParameters.N)
	for i := range s {
		pivot := solutionParameters.Pivots[i]

		s[i] = solutionParameters.Constants[i]
		for j, freeVariableColumn := range solutionParameters.FreeVariableColumns {
			s[i] -= combination[j] * freeVariableColumn[i]
		}

		if s[i]%pivot != 0 {
			return []int{}, false
		}
		s[i] /= pivot
	}

	if vectorContainsNegatives(s) {
		return []int{}, false
	}

	return s, true
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func greatestCommonDivisor(a, b int) int {
	absA := intAbs(a)
	absB := intAbs(b)
	lesser := int(math.Min(float64(absA), float64(absB)))
	for i := lesser; i > 0; i-- {
		if absA%i == 0 && absB%i == 0 {
			return i
		}
	}
	return 1
}

func sumVector(v []int) int {
	sum := 0
	for _, val := range v {
		sum += val
	}
	return sum
}
func vectorContainsNegatives(v []int) bool {
	for _, val := range v {
		if val < 0 {
			return true
		}
	}

	return false
}

func vectorToString(v []int) string {
	strs := make([]string, 0, len(v))
	for _, val := range v {
		str := strconv.Itoa(val)
		strs = append(strs, str)
	}
	return strings.Join(strs, ",")
}

func printMatrix(event string, a [][]int, b []int) {
	fmt.Printf("%s\n", event)
	for i, row := range a {
		fmt.Print("[")
		for j := range row {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Print("]")

		fmt.Printf("[%4d]\n", b[i])
	}
	fmt.Print("\n")
}
