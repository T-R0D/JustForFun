package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	info, err := parseUpdateInfo(input)
	if err != nil {
		return "", err
	}

	ruleConformingOrders := findRuleConformingPageOrders(info)

	sum, err := sumMiddlePageNumbers(ruleConformingOrders)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(sum), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	info, err := parseUpdateInfo(input)
	if err != nil {
		return "", err
	}

	ruleNonconformingOrders := findRuleNonconformingPageOrders(info)
	fixedPrintingOrders := fixBrokenPrintingOrders(info.PageOrderingRules, ruleNonconformingOrders)

	sum, err := sumMiddlePageNumbers(fixedPrintingOrders)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(sum), nil
}

func parseUpdateInfo(input string) (updateInfo, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return updateInfo{}, fmt.Errorf("expected 2 sections, got %d", len(sections))
	}

	ruleLines := strings.Split(sections[0], "\n")
	rules := map[int]intSet{}
	for i, line := range ruleLines {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return updateInfo{}, fmt.Errorf("line %d did not have 2 parts", i)
		}

		pre, err := strconv.Atoi(parts[0])
		if err != nil {
			return updateInfo{}, fmt.Errorf("line %d pre: %w", i, err)
		}

		post, err := strconv.Atoi(parts[1])
		if err != nil {
			return updateInfo{}, fmt.Errorf("line %d post: %w", i, err)
		}

		if postSet, ok := rules[pre]; ok {
			postSet[post] = struct{}{}
			rules[pre] = postSet
		} else {
			rules[pre] = intSet{post: struct{}{}}
		}
	}

	updateLines := strings.Split(sections[1], "\n")
	updateLists := make([][]int, 0, len(updateLines))
	for i, line := range updateLines {
		pageNumberStrs := strings.Split(line, ",")
		pageNumbers := make([]int, 0, len(pageNumberStrs))
		for j, str := range pageNumberStrs {
			pageNumber, err := strconv.Atoi(str)
			if err != nil {
				return updateInfo{}, fmt.Errorf("update list %d, item %d is not an int; %w", i, j, err)
			}
			pageNumbers = append(pageNumbers, pageNumber)
		}

		updateLists = append(updateLists, pageNumbers)
	}

	return updateInfo{
		PageOrderingRules:     rules,
		UpdatePageNumberLists: updateLists,
	}, nil
}

type updateInfo struct {
	PageOrderingRules     map[int]intSet
	UpdatePageNumberLists [][]int
}

type intSet map[int]struct{}

func findRuleConformingPageOrders(info updateInfo) [][]int {
	conformingPageOrders := [][]int{}
	for _, order := range info.UpdatePageNumberLists {
		if pageOrderFollowsRules(info.PageOrderingRules, order) {
			conformingPageOrders = append(conformingPageOrders, order)
		}
	}

	return conformingPageOrders
}

func findRuleNonconformingPageOrders(info updateInfo) [][]int {
	conformingPageOrders := [][]int{}
	for _, order := range info.UpdatePageNumberLists {
		if !pageOrderFollowsRules(info.PageOrderingRules, order) {
			conformingPageOrders = append(conformingPageOrders, order)
		}
	}

	return conformingPageOrders
}

func pageOrderFollowsRules(rules map[int]intSet, pageOrder []int) bool {
	seen := intSet{}
	for _, pageNumber := range pageOrder {
		disallowedPreviousNumbers := rules[pageNumber]
		for seenNumber := range seen {
			if _, ok := disallowedPreviousNumbers[seenNumber]; ok {
				return false
			}
		}

		seen[pageNumber] = struct{}{}
	}

	return true
}

func fixBrokenPrintingOrders(rules map[int]intSet, brokenPageOrders [][]int) [][]int {
	newOrders := make([][]int, 0, len(brokenPageOrders))
	for _, brokenOrder := range brokenPageOrders {
		fixedOrder := fixPrintingOrder(rules, brokenOrder)
		newOrders = append(newOrders, fixedOrder)
	}

	return newOrders
}

func fixPrintingOrder(rules map[int]intSet, order []int) []int {
	remainingPageNumbers := intSet{}
	for _, x := range order {
		remainingPageNumbers[x] = struct{}{}
	}

	fixedOrder := make([]int, 0, len(order))

	var visit func(int)
	visit = func(pageNumber int) {
		if _, ok := remainingPageNumbers[pageNumber]; !ok {
			return
		}

		if postSet, ok := rules[pageNumber]; ok {
			for x := range postSet {
				visit(x)
			}
		}

		fixedOrder = append(fixedOrder, pageNumber)
		delete(remainingPageNumbers, pageNumber)
	}

	for len(remainingPageNumbers) > 0 {
		for x := range remainingPageNumbers {
			visit(x)
		}
	}

	for i := range len(fixedOrder) / 2 {
		fixedOrder[i], fixedOrder[len(fixedOrder)-i-1] = fixedOrder[len(fixedOrder)-i-1], fixedOrder[i]
	}

	return fixedOrder
}

func sumMiddlePageNumbers(pageOrders [][]int) (int, error) {
	sum := 0
	for i, order := range pageOrders {
		if len(order)%2 != 1 {
			return 0, fmt.Errorf("page order %d (%v) was not odd length, had no middle number", i, order)
		}

		sum += order[len(order)/2]
	}

	return sum, nil
}
