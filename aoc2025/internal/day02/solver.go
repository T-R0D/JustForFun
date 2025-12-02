package day02

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	ranges, err := parseIDRanges(input)
	if err != nil {
		return "", err
	}

	invalidIdSum := 0
	for _, r := range ranges {
		for x := r.first; x <= r.last; x += 1 {
			if isRepeatedNumber(x) {
				invalidIdSum += x
			}
		}
	}

	return strconv.Itoa(invalidIdSum), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	ranges, err := parseIDRanges(input)
	if err != nil {
		return "", err
	}

	invalidIdSum := 0
	for _, r := range ranges {
		for x := r.first; x <= r.last; x += 1 {
			if isSequenceOfRepeats(x) {
				invalidIdSum += x
			}
		}
	}

	return strconv.Itoa(invalidIdSum), nil
}

type idRange struct {
	first int
	last  int
}

func parseIDRanges(input string) ([]idRange, error) {
	ranges := strings.Split(input, ",")
	parsedRanges := make([]idRange, 0, len(ranges))

	for _, rangeStr := range ranges {
		parts := strings.Split(rangeStr, "-")
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			return []idRange{}, err
		}
		last, err := strconv.Atoi(parts[1])
		if err != nil {
			return []idRange{}, err
		}

		parsedRanges = append(parsedRanges, idRange{first, last})
	}

	return parsedRanges, nil
}

func isRepeatedNumber(x int) bool {
	digits := intToDigitArray(x)
	nDigits := len(digits)

	if nDigits&1 == 1 {
		return false
	}

	halfLen := nDigits / 2
	for i := range halfLen {
		if digits[i] != digits[i+halfLen] {
			return false
		}
	}

	return true
}

func isSequenceOfRepeats(x int) bool {
	digits := intToDigitArray(x)
	nDigits := len(digits)
	halfLen := nDigits / 2

NextSequenceCandidate:
	for sequenceLen := 1; sequenceLen <= halfLen; sequenceLen += 1 {
		if nDigits%sequenceLen != 0 {
			continue NextSequenceCandidate
		}

		for nextSequenceStart := sequenceLen; nextSequenceStart < nDigits; nextSequenceStart += sequenceLen {
			for i := range sequenceLen {
				if digits[i] != digits[nextSequenceStart + i] {
					continue NextSequenceCandidate
				}
			}
		}

		return true
	}

	return false
}

func intToDigitArray(x int) []int {
	if x == 0 {
		return []int{0}
	}

	digits := []int{}
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	return digits
}
