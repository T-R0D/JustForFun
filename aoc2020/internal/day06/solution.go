package day06

import (
	"strconv"
	"strings"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	responses := parseResponses(input)

	sumOfYesResponses := 0
	for _, groupResponses := range responses {
		sumOfYesResponses += countYesResponsesForGroup(groupResponses)
	}

	return strconv.Itoa(sumOfYesResponses), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	responses := parseResponses(input)

	sumOfAllYesResponses := 0
	for _, groupResponses := range responses {
		sumOfAllYesResponses += countAllYesResponsesForGroup(groupResponses)
	}

	return strconv.Itoa(sumOfAllYesResponses), nil
}

func parseResponses(allAnswers string) [][]string {
	allResponses := [][]string{}
	groupResponses := []string{}
	response := strings.Builder{}
	newlineJustSeen := false
	for _, r := range allAnswers {
		switch r {
		case '\n':
			if !newlineJustSeen {
				newlineJustSeen = true
				groupResponses = append(groupResponses, response.String())
				response = strings.Builder{}
			} else {
				allResponses = append(allResponses, groupResponses)
				groupResponses = []string{}
				response = strings.Builder{}
				newlineJustSeen = false
			}
		default:
			response.WriteRune(r)
			newlineJustSeen = false
		}
	}
	groupResponses = append(groupResponses, response.String())
	allResponses = append(allResponses, groupResponses)

	return allResponses
}

func countYesResponsesForGroup(groupResponses []string) int {
	yesResponses := make([]int, 26)

	for _, response := range groupResponses {
		for _, r := range response {
			index := int(r - 'a')
			yesResponses[index] = 1
		}
	}

	totalYesResponses := 0
	for _, responseValue := range yesResponses {
		totalYesResponses += responseValue
	}

	return totalYesResponses
}

func countAllYesResponsesForGroup(groupResponses []string) int {
	groupSize := len(groupResponses)

	yesResponses := make([]int, 26)

	for _, response := range groupResponses {
		for _, r := range response {
			index := int(r - 'a')
			yesResponses[index]++
		}
	}

	totalAllYesResponses := 0
	for _, responseValue := range yesResponses {
		if responseValue == groupSize {
			totalAllYesResponses++
		}
	}

	return totalAllYesResponses
}
