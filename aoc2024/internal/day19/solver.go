package day19

import (
	"fmt"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	towelsAndTargets, err := parseTowelsAndTargets(input)
	if err != nil {
		return "", err
	}

	achievableDesigns := findAchievableDesigns(
		towelsAndTargets.Targets, towelsAndTargets.Towels)

	return fmt.Sprintf("%d", len(achievableDesigns)), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	towelsAndTargets, err := parseTowelsAndTargets(input)
	if err != nil {
		return "", err
	}

	nWaysToAchieveDesigns := countWaysToAchieveDesigns(
		towelsAndTargets.Targets, towelsAndTargets.Towels)

	return fmt.Sprintf("%d", nWaysToAchieveDesigns), nil
}

type towelsAndTargetsPair struct {
	Towels  [][]rune
	Targets [][]rune
}

func parseTowelsAndTargets(input string) (*towelsAndTargetsPair, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return nil, fmt.Errorf("there were not exactly 2 sections")
	}

	towelPatterns := parseTowelPatterns(sections[0])

	targetDesigns := parseDesigns(sections[1])

	return &towelsAndTargetsPair{
		Towels:  towelPatterns,
		Targets: targetDesigns,
	}, nil
}

func parseTowelPatterns(section string) [][]rune {
	towels := strings.Split(section, ", ")
	towelPatterns := make([][]rune, 0, len(towels))
	for _, towel := range towels {
		towelPatterns = append(towelPatterns, []rune(towel))
	}

	return towelPatterns
}

func parseDesigns(section string) [][]rune {
	lines := strings.Split(section, "\n")
	designs := make([][]rune, 0, len(lines))
	for _, line := range lines {
		designs = append(designs, []rune(line))
	}

	return designs
}

func findAchievableDesigns(designs [][]rune, patterns [][]rune) [][]rune {
	achievableDesigns := [][]rune{}

	for _, design := range designs {
		if isAchievableDesign(design, patterns) {
			achievableDesigns = append(achievableDesigns, design)
		}
	}

	return achievableDesigns
}

func isAchievableDesign(targetDesign []rune, patterns [][]rune) bool {
	frontier := queue.NewLifo[[]rune]()
	frontier.Push([]rune{})

SEARCH:
	for frontier.Len() > 0 {
		partialDesign, ok := frontier.Pop()
		if !ok {
			break SEARCH
		}

		if len(partialDesign) > len(targetDesign) {
			continue SEARCH
		}

		if len(partialDesign) == len(targetDesign) {
			return true
		}

	GENERATE_STATES:
		for _, pattern := range patterns {
			newPartialDesign := append([]rune{}, partialDesign...)
			newPartialDesign = append(newPartialDesign, pattern...)

			if len(newPartialDesign) > len(targetDesign) {
				continue GENERATE_STATES
			}
			for i := 0; i < len(newPartialDesign); i += 1 {
				if newPartialDesign[i] != targetDesign[i] {
					continue GENERATE_STATES
				}
			}
			frontier.Push(newPartialDesign)
		}
	}

	return false
}

func countWaysToAchieveDesigns(designs [][]rune, patterns [][]rune) int64 {
	nWaysToAchieveDesigns := int64(0)
	for _, design := range designs {
		nWaysToAchieveDesigns += countWaysToAchieveDesign(design, patterns)
	}

	return nWaysToAchieveDesigns
}

func countWaysToAchieveDesign(design []rune, patterns [][]rune) int64 {
	countTable := make([]int64, len(design))
	
	for i := range len(design) {
		TRY_PATTERN:
		for _, pattern := range patterns {
			if i + len(pattern) > len(design) {
				continue TRY_PATTERN
			}

			for j := range len(pattern) {
				if pattern[j] != design[i + j] {
					continue TRY_PATTERN
				}
			}

			if i <= 0 {
				countTable[i + len(pattern)-1] = 1
			} else {
				countTable[i + len(pattern)-1] += countTable[i-1]
			}
		}
	}

	return countTable[len(countTable)-1]
}
