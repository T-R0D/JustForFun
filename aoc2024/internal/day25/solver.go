package day25

import (
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	schematics := parseLockAndKeySchematics(input)

	matchingPairs := countMatchingLockKeyPairs(schematics)

	return fmt.Sprintf("%d", matchingPairs), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	return "Merry Christmas!", nil
}

type lockAndKeySchematics struct {
	Locks [][]int
	Keys  [][]int
}

func parseLockAndKeySchematics(input string) *lockAndKeySchematics {
	schematics := strings.Split(input, "\n\n")
	locks := [][]int{}
	keys := [][]int{}
	for _, schematic := range schematics {
		lines := strings.Split(schematic, "\n")
		columns := make([]int, 5)
		if lines[0][0] == '#' {
			for i := 1; i < len(lines); i += 1 {
				line := lines[i]
				for j, r := range line {
					if r == '#' {
						columns[j] += 1
					}
				}
			}
			locks = append(locks, columns)
		} else {
			for i := len(lines) - 2; i >= 0; i -= 1 {
				line := lines[i]
				for j, r := range line {
					if r == '#' {
						columns[j] += 1
					}
				}
			}
			keys = append(keys, columns)
		}
	}

	return &lockAndKeySchematics{
		Locks: locks,
		Keys:  keys,
	}
}

func countMatchingLockKeyPairs(schematics *lockAndKeySchematics) int {
	matchingPairs := 0
	for _, lock := range schematics.Locks {
	NEXT_KEY:
		for _, key := range schematics.Keys {
			for i := range 5 {
				if lock[i]+key[i] > 5 {
					continue NEXT_KEY
				}
			}
			matchingPairs += 1
		}
	}

	return matchingPairs
}
