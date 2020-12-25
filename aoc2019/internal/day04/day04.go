package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	start, end, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	nValid := 0
	for i := start; i < end; i++ {
		s, err := intToSlice(i)
		if err != nil {
			return nil, err
		}

		if codeIsValid(s) {
			nValid++
		}
	}

	return nValid, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	start, end, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	nValid := 0
	for i := start; i < end; i++ {
		s, err := intToSlice(i)
		if err != nil {
			return nil, err
		}

		if codeIsValidV2(s) {
			nValid++
		}
	}

	return nValid, nil
}

func parseInput(input string) (int, int, error) {
	parts := strings.Split(input, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Unexpected number of parts in input")
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("First part could not be converted to int")
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Second part could not be converted to int")
	}

	return start, end, nil
}

func intToSlice(i int) ([]int, error) {
	if i < 0 {
		return nil, fmt.Errorf("Negative numbers are not supported")
	} else if i == 0 {
		return []int{0}, nil
	}

	r := make([]int, 0)
	for i > 0 {
		r = append(r, i%10)
		i /= 10
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r, nil
}

func codeIsValid(code []int) bool {
	return digitsAreInIncreasingOrder(code) && codeHasAdjacentPairOfDigits(code)
}

func codeIsValidV2(code []int) bool {
	return digitsAreInIncreasingOrder(code) && codeHasStrictPairOfDigits(code)
}

func digitsAreInIncreasingOrder(code []int) bool {
	if code == nil || len(code) < 1 {
		return false
	}

	prevDigit := 0
	for _, d := range code {
		if prevDigit > d {
			return false
		}
		prevDigit = d
	}

	return true
}

func codeHasAdjacentPairOfDigits(code []int) bool {
	prevDigit := -1
	for _, d := range code {
		if d == prevDigit {
			return true
		}
		prevDigit = d
	}

	return false
}

func codeHasStrictPairOfDigits(code []int) bool {
	runLen := 0
	prevDigit := -1
	for _, d := range code {
		if d == prevDigit {
			runLen++
		} else if runLen == 2 {
			return true
		} else {
			runLen = 1
		}
		prevDigit = d
	}
	return runLen == 2
}

func codeHasEvenLengthRuns(code []int) bool {
	runLen := 0
	prevDigit := -1
	for _, d := range code {
		if d == prevDigit {
			runLen++
		} else if runLen > 2 && runLen%2 == 1 {
			return false
		} else {
			runLen = 1
		}
		prevDigit = d
	}
	return runLen == 1 || runLen%2 == 0
}
