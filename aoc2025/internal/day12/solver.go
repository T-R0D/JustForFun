package day12

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/set"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	specifications, err := parseSpecifications(input)
	if err != nil {
		return "", err
	}

	treesThatCanFitPresents := 0
	for _, treeSpec := range specifications.TreeSpecs {
		if !treeAreaIsGreaterThanGiftArea(treeSpec, specifications.GiftSpecs) {
			continue
		}

		if recursiveSearchForGiftPlacementSolution(treeSpec, specifications.GiftSpecs) {
			treesThatCanFitPresents += 1
		}
	}

	return strconv.Itoa(treesThatCanFitPresents), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	return "Merry Christmas!", nil
}

type specificationsPair struct {
	GiftSpecs []giftShapeSpec
	TreeSpecs []treeAreaSpec
}

const giftWidth = 3

type giftShapeSpec [giftWidth][giftWidth]rune

type treeAreaSpec struct {
	Depth      int
	Width      int
	Quantities []int
}

func parseSpecifications(input string) (specificationsPair, error) {
	parts := strings.Split(input, "\n\n")

	giftSpecs := make([]giftShapeSpec, 0, len(parts)-1)
	for _, chunk := range parts[:len(parts)-1] {
		lines := strings.Split(chunk, "\n")
		schematic := [3][3]rune{}
		for i, line := range lines[1:] {
			for j, r := range line {
				schematic[i][j] = r
			}
		}

		giftSpecs = append(giftSpecs, schematic)
	}

	treeSpecs := []treeAreaSpec{}
	for _, line := range strings.Split(parts[len(parts)-1], "\n") {
		chunks := strings.Split(line, ": ")

		areaPartStrs := strings.Split(chunks[0], "x")
		areaDimensions := [2]int{}
		for i, str := range areaPartStrs {
			val, err := strconv.Atoi(str)
			if err != nil {
				return specificationsPair{}, err
			}

			areaDimensions[i] = val
		}

		quantityStrs := strings.Split(chunks[1], " ")
		quantities := make([]int, 0, len(quantityStrs))
		for _, str := range quantityStrs {
			val, err := strconv.Atoi(str)
			if err != nil {
				return specificationsPair{}, err
			}

			quantities = append(quantities, val)
		}

		treeSpecs = append(treeSpecs, treeAreaSpec{
			Depth:      areaDimensions[0],
			Width:      areaDimensions[1],
			Quantities: quantities,
		})
	}

	return specificationsPair{
		GiftSpecs: giftSpecs,
		TreeSpecs: treeSpecs,
	}, nil
}

func treeAreaIsGreaterThanGiftArea(treeSpec treeAreaSpec, giftSpecs []giftShapeSpec) bool {
	giftAreas := make([]int, 0, len(giftSpecs))
	for _, spec := range giftSpecs {
		coveredArea := 0
		for i := range giftWidth {
			for j := range giftWidth {
				if spec[i][j] == occupiedSpace {
					coveredArea += 1
				}
			}
		}
		giftAreas = append(giftAreas, coveredArea)
	}

	treeArea := treeSpec.Depth * treeSpec.Width
	minimumGiftAreaRequired := 0
	for giftKey, quantity := range treeSpec.Quantities {
		minimumGiftAreaRequired += quantity * giftAreas[giftKey]
	}

	return treeArea > minimumGiftAreaRequired
}

func recursiveSearchForGiftPlacementSolution(treeSpec treeAreaSpec, giftSpecs []giftShapeSpec) bool {
	orientedGiftSpecs := make([][]giftShapeSpec, 0, len(giftSpecs))
	for _, spec := range giftSpecs {
		orientedGiftSpecs = append(orientedGiftSpecs, generateGiftOrientations(spec))
	}

	placementField := make([][]rune, 0, treeSpec.Depth)
	for range treeSpec.Depth {
		row := slices.Repeat([]rune{freeSpace}, treeSpec.Width)
		placementField = append(placementField, row)
	}

	quantities := append([]int{}, treeSpec.Quantities...)

	memo := set.New[string]()

	return recursiveSearchForGiftPlacementSolutionInner(memo, treeSpec, orientedGiftSpecs, placementField, quantities)
}

func recursiveSearchForGiftPlacementSolutionInner(memo set.Set[string], treeSpec treeAreaSpec, orientedGiftSpecs [][]giftShapeSpec, placementField [][]rune, quantities []int) bool {
	if vectorIsAllZeroes(quantities) {
		return true
	}

	for giftKey, remaining := range quantities {
		if remaining == 0 {
			continue
		}
		for _, spec := range orientedGiftSpecs[giftKey] {
			for i := range treeSpec.Depth - giftWidth + 1 {
			Placement:
				for j := range treeSpec.Width - giftWidth + 1 {
					for i2 := range giftWidth {
						for j2 := range giftWidth {
							if spec[i2][j2] == occupiedSpace && placementField[i+i2][j+j2] == occupiedSpace {
								continue Placement
							}
						}
					}

					quantities[giftKey] -= 1
					for i2 := range giftWidth {
						for j2 := range giftWidth {
							if spec[i2][j2] == occupiedSpace {
								placementField[i+i2][j+j2] = occupiedSpace
							}
						}
					}

					solutionFound := recursiveSearchForGiftPlacementSolutionInner(memo, treeSpec, orientedGiftSpecs, placementField, quantities)
					if solutionFound {
						return true
					}

					quantities[giftKey] += 1
					for i2 := range giftWidth {
						for j2 := range giftWidth {
							if spec[i2][j2] == occupiedSpace {
								placementField[i+i2][j+j2] = freeSpace
							}
						}
					}
				}
			}
		}
	}

	return false
}

func generateGiftOrientations(spec giftShapeSpec) []giftShapeSpec {
	discoveredOrientations := set.New[string]()

	specs := make([]giftShapeSpec, 0, 4)
	specs = append(specs, spec)
	discoveredOrientations.Add(fmt.Sprintf("%v", spec))
	previous := spec
	for r := 1; r < 4; r += 1 {
		next := giftShapeSpec{}
		for i := range giftWidth {
			for j := range giftWidth {
				next[giftWidth-j-1][i] = previous[i][j]
			}
		}

		specStr := fmt.Sprintf("%v", next)
		if !discoveredOrientations.Contains(specStr) {
			specs = append(specs, next)
			discoveredOrientations.Add(specStr)
		}

		previous = next
	}

	return specs
}

func vectorIsAllZeroes(v []int) bool {
	for _, value := range v {
		if value != 0 {
			return false
		}
	}
	return true
}

const (
	freeSpace     = '.'
	occupiedSpace = '#'
)
