package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	reports, err := parseReports(input)
	if err != nil {
		return "", err
	}

	nSafeReports := countSafeReports(reports, isSafeReport)

	return strconv.Itoa(nSafeReports), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	reports, err := parseReports(input)
	if err != nil {
		return "", err
	}

	nSafeReports := countSafeReports(reports, isSafeReportWithDampening)

	return strconv.Itoa(nSafeReports), nil
}

func parseReports(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	reports := make([][]int, 0, len(lines))
	for i, line := range lines {
		items := strings.Fields(line)
		report := make([]int, len(items))
		for j, item := range items {
			value, err := strconv.Atoi(item)
			if err != nil {
				return nil, fmt.Errorf("item %d %d, is not valid integer: %w", i, j, err)
			}
			report[j] = value
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func countSafeReports(reports [][]int, isSafe func([]int) bool) int {
	nSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			nSafe += 1
		}
	}

	return nSafe
}

func isSafeReport(report []int) bool {
	if isMonotonicWithinTolerance(report, 1, 3) ||
		isMonotonicWithinTolerance(report, -3, -1) {

		return true
	}

	return false
}

func isMonotonicWithinTolerance(report []int, lowerBound int, upperBound int) bool {
	for i := range len(report) - 1 {
		difference := report[i] - report[i+1]
		if !(lowerBound <= difference && difference <= upperBound) {
			return false
		}
	}

	return true
}

func isSafeReportWithDampening(report []int) bool {
	if isMonotonicWithinToleranceWithDampening(report, 1, 3) ||
		isMonotonicWithinToleranceWithDampening(report, -3, -1) {

		return true
	}

	return false
}

func isMonotonicWithinToleranceWithDampening(report []int, lowerBound int, upperBound int) bool {
	if len(report) <= 2 {
		return true
	}

	var dampenedIndex *int = nil

	for i := range len(report) - 1 {
		a, b := report[i], report[i+1]
		if dampenedIndex != nil && (*dampenedIndex == i) {
			a = report[i-1]
		}

		if withinTolerance(a, b, lowerBound, upperBound) {
			continue
		}

		if dampenedIndex != nil {
			return false
		}

		indexToDampen := -1
		if i == 0 {
			a, b, c := report[i], report[i+1], report[i+2]
			if withinTolerance(b, c, lowerBound, upperBound) {
				indexToDampen = i
			} else if withinTolerance(a, c, lowerBound, upperBound) {
				indexToDampen = i + 1
			} else {
				return false
			}
		} else if i <= len(report)-3 {
			a, b, c, d := report[i-1], report[i], report[i+1], report[i+2]
			if withinTolerance(b, d, lowerBound, upperBound) {
				indexToDampen = i + 1
			} else if withinTolerance(b, c, lowerBound, upperBound) {
				indexToDampen = i - 1
			} else if withinTolerance(a, c, lowerBound, upperBound) {
				indexToDampen = i
			} else {
				return false
			}
		} else {
			indexToDampen = i + 1
		}
		dampenedIndex = &indexToDampen
	}

	return true
}

func withinTolerance(a int, b int, lowerBound int, upperBound int) bool {
	difference := a - b
	return lowerBound <= difference && difference <= upperBound
}
