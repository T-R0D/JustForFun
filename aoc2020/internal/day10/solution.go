// Possible Improvement: In part 1, sort the adapters in an array and build up linearly.
// Question: People on Reddit have been exploiting the structure of the input, maybe we
//           can do something like that?

package day10

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the problem for the day.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	adapters, err := parseInputIntoSet(input)
	if err != nil {
		return "", err
	}

	joltageDifferenceSummary, err := connectAllAdapters(adapters, maxAllowedJoltageDifference)
	if err != nil {
		return "", err
	}

	oneJoltageDifferences := joltageDifferenceSummary[1]
	// +1 for the built in adapter.
	threeJoltageDifferences := joltageDifferenceSummary[3] + 1

	return strconv.Itoa(oneJoltageDifferences * threeJoltageDifferences), nil
}

// Part2 solves part 1 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	adapters, err := parseInputIntoSet(input)
	if err != nil {
		return "", err
	}

	waysToLinkAdapters := findNumberOfWaysToLinkAdapters(adapters, maxAllowedJoltageDifference)

	return fmt.Sprintf("%d", waysToLinkAdapters), nil
}

type adapterSet map[int]struct{}

func parseInputIntoSet(input string) (adapterSet, error) {
	lines := strings.Split(input, "\n")
	adapters := adapterSet{}
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.Wrapf(err, "error processing line %d", i)
		}
		adapters[value] = struct{}{}
	}
	return adapters, nil
}

const (
	minAllowedJoltageDifference = 1
	maxAllowedJoltageDifference = 3

	builtInAdapterJoltageDifference = 3
)

type adapterResult struct {
	difference int
	value      int
}

func connectAllAdapters(adapters adapterSet, allowedJoltageDifference int) (map[int]int, error) {
	currentJoltage := 0
	adaptersUsed := 0
	joltageDifferenceSummary := map[int]int{}

	for {
		result, found := findNextAdapter(currentJoltage, adapters, allowedJoltageDifference)

		if !found {
			break
		}

		currentJoltage = result.value
		joltageDifferenceSummary[result.difference] = joltageDifferenceSummary[result.difference] + 1
		adaptersUsed++
	}

	if adaptersUsed != len(adapters) {
		return nil, errors.New("not all adapters were used")
	}

	return joltageDifferenceSummary, nil
}

func findNextAdapter(currentJoltage int, adapters adapterSet, allowedJoltageDifference int) (*adapterResult, bool) {
	for difference := 1; difference <= allowedJoltageDifference; difference++ {
		adapterJoltage := currentJoltage + difference
		if _, ok := adapters[adapterJoltage]; ok {
			return &adapterResult{difference: difference, value: adapterJoltage}, true
		}
	}
	return nil, false
}

func findNumberOfWaysToLinkAdapters(adapters adapterSet, allowedJoltageDifference int) int64 {
	if len(adapters) < 1 {
		return 0
	}

	// TODO: check for first adapter presence.
	
	adaptersInSortedOrder := getAdaptersInSortedOrder(adapters)
	adapterToPositionInSortedOrder := getAdapterToPositionInSortedOrderLookup(adaptersInSortedOrder)
	waysToReachJoltage := make([]int64, len(adaptersInSortedOrder))

	for i, adapter := range adaptersInSortedOrder {
		waysToReachPredecessors := int64(0)
		for difference := 1; difference <= allowedJoltageDifference; difference++ {
			previousCandidate := adapter - difference
			if previousCandidate == 0 {
				waysToReachPredecessors++
			} else if index, ok := adapterToPositionInSortedOrder[previousCandidate]; ok {
				waysToReachPredecessors += waysToReachJoltage[index]
			}
		}
		waysToReachJoltage[i] = waysToReachPredecessors
	}

	return waysToReachJoltage[len(waysToReachJoltage) - 1]
}

func getAdaptersInSortedOrder(adapters adapterSet) []int {
	adaptersInSortedOrder := make([]int, 0, len(adapters))
	for adapter := range adapters {
		adaptersInSortedOrder = append(adaptersInSortedOrder, adapter)
	}
	sort.Ints(adaptersInSortedOrder)
	return adaptersInSortedOrder
}

func getAdapterToPositionInSortedOrderLookup(adaptersInSortedOrder []int) map[int]int {
	lookup := map[int]int{}
	for i, adapter := range adaptersInSortedOrder {
		lookup[adapter] = i
	}
	return lookup
}