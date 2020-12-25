// Possible Improvement: Make the representation a bidi graph, then do only one
//                       DFS in part one instead of multiple. Alternately, memoize.

package day07

import (
	"strconv"
	"strings"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	lookup := buildBagCapabilitiesLookup(input)

	nColorsThatCanHoldShinyGold := 0
	for color := range lookup {
		if canHoldShinyGold(color, lookup) {
			nColorsThatCanHoldShinyGold++
		}
	}

	return strconv.Itoa(nColorsThatCanHoldShinyGold), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	lookup := buildBagCapabilitiesLookup(input)

	bagCapacityMemoTable := map[string]int{}

	bagsHeldByShinyGoldBag := findNumberOfBagsHeld(myBag, lookup, bagCapacityMemoTable)

	return strconv.Itoa(bagsHeldByShinyGoldBag), nil
}

const (
	myBag = "shiny gold"
)

type bagCapabilitiesLookup map[string][]bagSlots

type bagSlots struct {
	color string
	count int
}

type capabilityDef struct {
	color    string
	contents []bagSlots
}

func buildBagCapabilitiesLookup(input string) bagCapabilitiesLookup {
	lines := strings.Split(input, "\n")

	lookup := bagCapabilitiesLookup{}

	for _, line := range lines {
		color, slots := parseLineOfInput(line)
		lookup[color] = slots
	}

	return lookup
}

func parseLineOfInput(line string) (string, []bagSlots) {
	colorAndSlots := strings.Split(line, " bags contain ")
	parentColor := colorAndSlots[0]
	slotsStr := colorAndSlots[1]

	slotsStr = strings.ReplaceAll(slotsStr, ".", "")
	slotsStr = strings.ReplaceAll(slotsStr, " bags", "")
	slotsStr = strings.ReplaceAll(slotsStr, " bag", "")

	slotDefs := strings.Split(slotsStr, ", ")
	slots := []bagSlots{}
	for _, def := range slotDefs {
		def = strings.Replace(def, " ", ".", 1)
		countAndColor := strings.Split(def, ".")

		if countAndColor[0] == "no" {
			continue
		}

		count, err := strconv.Atoi(countAndColor[0])
		if err != nil {
			panic(err)
		}

		slots = append(slots, bagSlots{color: countAndColor[1], count: count})
	}

	return parentColor, slots
}

func canHoldShinyGold(color string, lookup bagCapabilitiesLookup) bool {
	frontier := []string{color}
	seenColors := map[string]struct{}{}

	for len(frontier) > 0 {
		currentColor := frontier[0]
		frontier = frontier[1:]

		if _, ok := seenColors[currentColor]; ok {
			continue
		}

		contents := lookup[currentColor]
		for _, content := range contents {
			if content.color == myBag {
				return true
			}

			if _, ok := seenColors[content.color]; ok {
				continue
			}

			frontier = append(frontier, content.color)
		}

		seenColors[currentColor] = struct{}{}
	}

	return false
}

func findNumberOfBagsHeld(color string, lookup bagCapabilitiesLookup, bagCapacityMemoTable map[string]int) int {
	if numberHeld, ok := bagCapacityMemoTable[color]; ok {
		return numberHeld
	}

	numberHeld := 0
	slots := lookup[color]
	for _, slot := range slots {
		numberHeld += slot.count * (1 + findNumberOfBagsHeld(slot.color, lookup, bagCapacityMemoTable))
	}

	bagCapacityMemoTable[color] = numberHeld
	return numberHeld
}
