package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	stones, err := parseInitialStoneConfiguration(input)
	if err != nil {
		return "", err
	}

	resultingStones := blinkAndObserveStones(stones, 25)

	return strconv.Itoa(len(resultingStones)), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	stones, err := parseInitialStoneConfiguration(input)
	if err != nil {
		return "", err
	}

	nStones := blinkAndObserveStoneCounts(stones, 75)

	return strconv.FormatUint(nStones, 10), nil
}

func parseInitialStoneConfiguration(input string) ([]uint64, error) {
	numberStrs := strings.Fields(input)
	stones := make([]uint64, len(numberStrs))
	for i, numberStr := range numberStrs {
		stone, err := strconv.ParseUint(numberStr, 10, 64)
		if err != nil {
			return []uint64{}, fmt.Errorf("'%s' is not a valid stone engraving number", numberStr)
		}

		stones[i] = stone
	}

	return stones, nil
}

func blinkAndObserveStones(initialStones []uint64, nBlinks int) []uint64 {
	stones := initialStones
	for range nBlinks {
		stones = simulateStoneChange(stones)
	}

	return stones
}

func simulateStoneChange(stones []uint64) []uint64 {
	nextStones := make([]uint64, 0, len(stones))
	for _, stone := range stones {
		if stone == 0 {
			nextStones = append(nextStones, 1)
		} else if nDigits := countDigits(stone); nDigits&0x01 == 0 {
			halvingPowerOfTen := 1
			for range nDigits / 2 {
				halvingPowerOfTen *= 10
			}

			nextStones = append(nextStones, stone/uint64(halvingPowerOfTen))
			nextStones = append(nextStones, stone%uint64(halvingPowerOfTen))
		} else {
			nextStones = append(nextStones, stone*2024)
		}
	}

	return nextStones
}

func countDigits(x uint64) int {
	powerOfTen := uint64(1)
	nDigits := 0
	for ; powerOfTen <= x; powerOfTen, nDigits = powerOfTen*10, nDigits+1 {
	}
	return nDigits
}

func blinkAndObserveStoneCounts(initialStones []uint64, nBlinks int) uint64 {
	stoneCounts := map[uint64]uint64{}
	for _, stone := range initialStones {
		if count, ok := stoneCounts[stone]; ok {
			stoneCounts[stone] = count + 1
		} else {
			stoneCounts[stone] = 1
		}
	}

	for range nBlinks {
		stoneCounts = simulateStoneSplitCounts(stoneCounts)
	}

	nStones := uint64(0)
	for _, count := range stoneCounts {
		nStones += count
	}

	return nStones
}

func simulateStoneSplitCounts(stoneCounts map[uint64]uint64) map[uint64]uint64 {
	nextCounts := map[uint64]uint64{}
	for initialStoneValue, initialCount := range stoneCounts {
		if initialStoneValue == 0 {
			newStoneValue := uint64(1)
			if count, ok := nextCounts[newStoneValue]; ok {
				nextCounts[newStoneValue] = count + initialCount
			} else {
				nextCounts[newStoneValue] = initialCount
			}
		} else if nDigits := countDigits(initialStoneValue); nDigits&0x01 == 0 {
			halvingPowerOfTen := 1
			for range nDigits / 2 {
				halvingPowerOfTen *= 10
			}

			leftStoneValue := initialStoneValue / uint64(halvingPowerOfTen)
			rightStoneValue := initialStoneValue % uint64(halvingPowerOfTen)

			for _, newStoneValue := range []uint64{leftStoneValue, rightStoneValue} {
				if count, ok := nextCounts[newStoneValue]; ok {
					nextCounts[newStoneValue] = count + initialCount
				} else {
					nextCounts[newStoneValue] = initialCount
				}
			}
		} else {
			newStoneValue := initialStoneValue * 2024
			if count, ok := nextCounts[newStoneValue]; ok {
				nextCounts[newStoneValue] = count + initialCount
			} else {
				nextCounts[newStoneValue] = initialCount
			}
		}
	}

	return nextCounts
}
