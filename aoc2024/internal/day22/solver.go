package day22

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	initialSecretNumbers, err := parseInitialSecretNumbers(input)
	if err != nil {
		return "", err
	}

	total := sumTargetSecretNumbers(initialSecretNumbers, maxSecretIterations)

	return strconv.FormatInt(total, 10), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	initialSecretNumbers, err := parseInitialSecretNumbers(input)
	if err != nil {
		return "", err
	}

	maxBananas := findMaxBananasBySelling(initialSecretNumbers)

	return strconv.FormatInt(maxBananas, 10), nil
}

const (
	maxSecretIterations = 2000

	secretMultiplierA int64 = 64
	secretMultiplierB int64 = 2048
	secretDivisor     int64 = 32
	secretModulo      int64 = 16777216
)

func parseInitialSecretNumbers(input string) ([]int64, error) {
	lines := strings.Split(input, "\n")
	secretNumbers := make([]int64, len(lines))
	for i, line := range lines {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return []int64{}, fmt.Errorf("line %d did not contain a valid number; %w", i, err)
		}
		secretNumbers[i] = value
	}

	return secretNumbers, nil
}

func sumTargetSecretNumbers(initialSecretNumbers []int64, targetIterations int) int64 {
	total := int64(0)
	for _, initialNumber := range initialSecretNumbers {
		total += getNthSecretNumber(initialNumber, targetIterations)
	}

	return total
}

func getNthSecretNumber(seed int64, n int) int64 {
	secretNumber := seed
	for range n {
		secretNumber = nextSecretNumber(secretNumber)
	}

	return secretNumber
}

func nextSecretNumber(number int64) int64 {
	a := ((number * secretMultiplierA) ^ number) % secretModulo
	b := ((a / secretDivisor) ^ a) % secretModulo
	return ((b * secretMultiplierB) ^ b) % secretModulo
}

type changeSequenceT [4]int64

func findMaxBananasBySelling(initialSecretNumbers []int64) int64 {
	sequenceToFirstResults := make([]map[changeSequenceT]int64, 0, len(initialSecretNumbers))
	for _, initialSecretNumber := range initialSecretNumbers {
		sequenceToFirstResults = append(
			sequenceToFirstResults,
			getChangeSequenceToFirstResult(initialSecretNumber, maxSecretIterations))
	}

	sequenceSet := mergeSequencesIntoSet(sequenceToFirstResults)

	maxBananas := int64(0)
	for sequence := range sequenceSet {
		candidateBananas := int64(0)
		for _, sequenceToFirstResult := range sequenceToFirstResults {
			if bananasGained, ok := sequenceToFirstResult[sequence]; ok {
				candidateBananas += bananasGained
			}
		}

		if candidateBananas > maxBananas {
			maxBananas = candidateBananas
		}
	}

	return maxBananas
}

func getChangeSequenceToFirstResult(seed int64, n int) map[changeSequenceT]int64 {
	changes := make([]int64, n)
	pricesInBananas := make([]int64, n)
	secretNumber := seed
	for i := range n {
		previous := secretNumber
		secretNumber = nextSecretNumber(previous)
		changes[i] = (secretNumber % 10) - (previous % 10)
		pricesInBananas[i] = secretNumber % 10
	}

	changeSequenceToFirstValue := map[changeSequenceT]int64{}
	for i := 3; i < n; i += 1 {
		sequence := changeSequenceT{changes[i-3], changes[i-2], changes[i-1], changes[i]}
		if _, ok := changeSequenceToFirstValue[sequence]; ok {
			continue
		}

		changeSequenceToFirstValue[sequence] = pricesInBananas[i]
	}

	return changeSequenceToFirstValue
}

func mergeSequencesIntoSet(sequenceToFirstResults []map[changeSequenceT]int64) map[changeSequenceT]struct{} {
	sequenceSet := map[changeSequenceT]struct{}{}
	for _, sequenceToFirstResult := range sequenceToFirstResults {
		for sequence := range sequenceToFirstResult {
			sequenceSet[sequence] = struct{}{}
		}
	}

	return sequenceSet
}
