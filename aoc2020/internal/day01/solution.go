// Possible improvement: solve generically by using a combination generating function.

package day01

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	expenses, err := convertInputToIntList(input)
	if err != nil {
		return "", err
	}

	expenseKeyedByDifference := map[int]int{}
	for _, expense := range expenses {
		expenseKeyedByDifference[magicYear2020 - expense] = expense
	}

	for _, expense := range expenses {
		if expense1, ok := expenseKeyedByDifference[expense]; ok {
			return strconv.Itoa(expense1 * expense), nil
		}
	}

	return "", errors.New("failed to find a pair of expenses that fit the criteria")
}

// Part2 solves part 2 of the day's problem.
// TODO: Solve this with the part 1 method (need a combination generating function).
func (s *Solver) Part2(input string) (string, error) {
	expenses, err := convertInputToIntList(input)
	if err != nil {
		return "", err
	}

	for i, expense1 := range expenses {
		for j, expense2 := range expenses {
			for k, expense3 := range expenses {
				if i == j || i == k || j == k {
					continue
				}

				if expense1+expense2+expense3 == magicYear2020 {
					result := expense1 * expense2 * expense3
					return strconv.Itoa(result), nil
				}
			}
		}
	}

	return "", errors.New("failed to find a triple of expenses that fit the criteria")
}

const (
	magicYear2020 = 2020
)

func convertInputToIntList(input string) ([]int, error) {
	intsAsStrings := strings.Split(input, "\n")

	ints := make([]int, len(intsAsStrings))
	for i, intAsString := range intsAsStrings {
		integer, err := strconv.Atoi(intAsString)
		if err != nil {
			return nil, err
		}
		ints[i] = integer
	}

	return ints, nil
}
