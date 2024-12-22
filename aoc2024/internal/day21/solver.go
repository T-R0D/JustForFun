package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

// 173652 too high
// 170432 too high
// 169164 too high
// 168568 not right
func (s *Solver) SolvePartOne(input string) (string, error) {
	codes := parseCodes(input)

	totalComplexity, err := findComplexityScoreOfEnteringCodes(codes)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(totalComplexity, 10), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	return "", fmt.Errorf("part 2 not implemented")
}

type vec2 [2]int

func (v vec2) Plus(that vec2) vec2 {
	return vec2{v[0] + that[0], v[1] + that[1]}
}

func (v vec2) Distance(that vec2) int {
	y := v[0] - that[0]
	if y < 0 {
		y = that[0] - v[0]
	}

	x := v[1] - that[1]
	if x < 0 {
		x = that[1] - v[1]
	}

	return x + y
}

type unlockCode []rune

func (u unlockCode) NumericValue() (int64, error) {
	return strconv.ParseInt(string(u[:len(u)-1]), 10, 64)
}

func parseCodes(input string) []unlockCode {
	lines := strings.Split(input, "\n")
	codes := make([]unlockCode, 0, len(lines))
	for _, line := range lines {
		codes = append(codes, []rune(line))
	}

	return codes
}

var numberKeyLocations map[rune]vec2 = map[rune]vec2{
	'7': {0, 0},
	'8': {0, 1},
	'9': {0, 2},
	'4': {1, 0},
	'5': {1, 1},
	'6': {1, 2},
	'1': {2, 0},
	'2': {2, 1},
	'3': {2, 2},
	'0': {3, 1},
	'A': {3, 2},
}

var directionalKeyLocations map[rune]vec2 = map[rune]vec2{
	'^': {0, 1},
	'A': {0, 2},
	'<': {1, 0},
	'v': {1, 1},
	'>': {1, 2},
}

func findComplexityScoreOfEnteringCodes(codes []unlockCode) (int64, error) {
	numberPadDirectionRanks := map[rune]int{
		'<': 1,
		'>': 3,
		'v': 4,
		'^': 2,
	}

	numberPadNavigations := newKeyToNavigationMap(numberKeyLocations, map[rune]int{})
	directionPadNavigations := newKeyToNavigationMap(directionalKeyLocations, numberPadDirectionRanks)
	// directionPadNavigations = newRefinedKeyToNavigationMap(directionalKeyLocations, directionPadNavigations)
	// numberPadNavigations = newRefinedKeyToNavigationMap(numberKeyLocations, directionPadNavigations)

	// pairs := [][]rune{
	// 	{'3', '7'},
	// 	{'1', '9'},
	// 	{'A', '8'},
	// 	{'A', '7'},
	// 	{'A', '5'},
	// 	{'A', '4'},
	// 	{'A', '1'},
	// }
	// for _, pair := range pairs {
	// 	func(src rune, dst rune) {
	// 		fmt.Printf("%c -> %c | %s\n", src, dst, string(numberPadNavigations[src][dst]))
	// 		fmt.Printf("%c -> %c | %s\n", dst, src, string(numberPadNavigations[dst][src]))
	// 	}(pair[0], pair[1])
	// }

	// numberPadNavigations['3']['7'] = []rune{'<', '<', '^', '^'}
	// numberPadNavigations['1']['9'] = []rune{'>', '>', '^', '^'}
	// numberPadNavigations['A']['8'] = []rune{'<', '^', '^', '^'}
	// numberPadNavigations['A']['5'] = []rune{'<', '^', '^'}
	// // numberPadNavigations['A']['4'] = []rune{'<', '^', '^'}
	// numberPadNavigations['A']['1'] = []rune{'<', '<', '^'}
	// // numberPadNavigations['7']['3'] = []rune{'>', '>', 'v', 'v'}
	// // numberPadNavigations['9']['1'] = []rune{'v', 'v', '<', '<'}
	// // numberPadNavigations['8']['A'] = []rune{'>', 'v', 'v', 'v'}
	// numberPadNavigations['5']['A'] = []rune{'>', 'v', 'v'}

	totalComplexity := int64(0)
	for _, code := range codes {
		value, err := code.NumericValue()
		if err != nil {
			return 0, err
		}

		// fmt.Printf("code: %v\n", string(code))

		sequence := findSequenceToEnterCode(numberPadNavigations, directionPadNavigations, code)

		// fmt.Printf("%d * %d = %d\n", value, len(sequence), value*int64(len(sequence)))

		totalComplexity += value * int64(len(sequence))
	}

	return totalComplexity, nil
}

type keyToKeyNavigationLookup map[rune]navigationLookup

type navigationLookup map[rune][]rune

func newKeyToNavigationMap(keyLocations map[rune]vec2, directionRanks map[rune]int) keyToKeyNavigationLookup {
	keyToNavigation := keyToKeyNavigationLookup{}
	for src := range keyLocations {
		navigations := navigationLookup{}
		for dst := range keyLocations {
			navigations[dst] = navigateFromSrcKeyToDstKey(keyLocations, directionRanks, src, dst)
		}

		keyToNavigation[src] = navigations
	}

	return keyToNavigation
}

func newRefinedKeyToNavigationMap(keyLocations map[rune]vec2, translationLookup keyToKeyNavigationLookup) keyToKeyNavigationLookup {
	keyToNavigation := keyToKeyNavigationLookup{}
	for src := range keyLocations {
		navigations := navigationLookup{}
		for dst := range keyLocations {
			navigations[dst] = refinedNavigateFromSrcKeyToDstKey(keyLocations, translationLookup, src, dst)
		}

		keyToNavigation[src] = navigations
	}

	return keyToNavigation
}

func navigateFromSrcKeyToDstKey(keyLocations map[rune]vec2, directionRanks map[rune]int, src rune, dst rune) []rune {
	locationToKey := map[vec2]rune{}
	for key, location := range keyLocations {
		locationToKey[location] = key
	}

	type searchState struct {
		Location vec2
		Path     []rune
	}

	cmpSearchState := func(a searchState, b searchState) bool {
		if len(a.Path) != len(b.Path) {
			return len(a.Path) < len(b.Path)
		}

		runScoreA := getPathRunScore(a.Path)
		runScoreB := getPathRunScore(b.Path)

		if len(a.Path) == 0 {
			return false
		}

		if runScoreA != runScoreB {
			return runScoreA > runScoreB
		}

		firstDirRankA := directionRanks[a.Path[0]]
		firstDirRankB := directionRanks[b.Path[0]]
		return firstDirRankA < firstDirRankB
	}

	frontier := queue.NewPriority[searchState](cmpSearchState)
	frontier.Push(searchState{
		Location: keyLocations[src],
		Path:     []rune{},
	})

	type deltaAndDirection struct {
		Delta     vec2
		Direction rune
	}
	deltas := []deltaAndDirection{
		{Delta: vec2{0, -1}, Direction: '<'},
		{Delta: vec2{0, 1}, Direction: '>'},
		{Delta: vec2{-1, 0}, Direction: '^'},
		{Delta: vec2{1, 0}, Direction: 'v'},
	}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		if key, ok := locationToKey[state.Location]; !ok {
			continue
		} else if key == dst {
			if src == '3' && dst == '7' {
				return []rune{'<', '<', '^', '^'}
			}
			return state.Path
		}

		for _, d := range deltas {
			newLocation := state.Location.Plus(d.Delta)
			newPath := append([]rune{}, state.Path...)
			newPath = append(newPath, d.Direction)
			frontier.Push(searchState{
				Location: newLocation,
				Path:     newPath,
			})
		}
	}

	return []rune{}
}

func refinedNavigateFromSrcKeyToDstKey(keyLocations map[rune]vec2, translationLookup keyToKeyNavigationLookup, src rune, dst rune) []rune {
	locationToKey := map[vec2]rune{}
	for key, location := range keyLocations {
		locationToKey[location] = key
	}

	type searchState struct {
		Location        vec2
		Path            []rune
		TranslationCost int
	}

	cmpSearchState := func(a searchState, b searchState) bool {
		if len(a.Path) != len(b.Path) {
			return len(a.Path) < len(b.Path)
		}

		runScoreA := getPathRunScore(a.Path)
		runScoreB := getPathRunScore(b.Path)

		if len(a.Path) == 0 {
			return false
		}

		if runScoreA != runScoreB {
			return runScoreA > runScoreB
		}

		return a.TranslationCost < b.TranslationCost
	}

	frontier := queue.NewPriority[searchState](cmpSearchState)
	frontier.Push(searchState{
		Location:        keyLocations[src],
		Path:            []rune{},
		TranslationCost: 0,
	})

	type deltaAndDirection struct {
		Delta     vec2
		Direction rune
	}
	deltas := []deltaAndDirection{
		{Delta: vec2{-1, 0}, Direction: '^'},
		{Delta: vec2{1, 0}, Direction: 'v'},
		{Delta: vec2{0, -1}, Direction: '<'},
		{Delta: vec2{0, 1}, Direction: '>'},
	}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		if key, ok := locationToKey[state.Location]; !ok {
			continue
		} else if key == dst {
			if src == '3' && dst == '7' {
				return []rune{'<', '<', '^', '^'}
			}
			return state.Path
		}

		for _, d := range deltas {
			newLocation := state.Location.Plus(d.Delta)
			newPath := append([]rune{}, state.Path...)
			newPath = append(newPath, d.Direction)
			frontier.Push(searchState{
				Location:        newLocation,
				Path:            newPath,
				TranslationCost: getTranslationCost(translationLookup, newPath, 2),
			})
		}
	}

	return []rune{}
}

func getPathRunScore(path []rune) int {
	score := 0
	currentRun := 0
	for i := range len(path) - 1 {
		if path[i] == path[i+1] {
			currentRun += 1
		} else {
			score += currentRun
			currentRun = 0
		}
	}
	score += currentRun

	return score
}

func getTranslationCost(translationLookup keyToKeyNavigationLookup, newPath []rune, layersOfTranslation int) int {
	sequence := append([]rune{}, newPath...)
	for range layersOfTranslation {
		sequence = findSequenceOfKeyPressesToAchieveRemoteSequence(sequence, translationLookup)
	}

	return len(sequence)
}

func findSequenceToEnterCode(
	numberPadNavigations keyToKeyNavigationLookup,
	directionPadNavigations keyToKeyNavigationLookup,
	code unlockCode) []rune {

	mappingLayers := []keyToKeyNavigationLookup{
		numberPadNavigations,
		directionPadNavigations,
		directionPadNavigations,
	}

	nextSequence := []rune(code)
	for _, transformation := range mappingLayers {

		// fmt.Printf("\tremoteSequence: %v\n", string(nextSequence))

		nextSequence = findSequenceOfKeyPressesToAchieveRemoteSequence(
			nextSequence, transformation)

	}
	// fmt.Printf("\tfinalSequence : %v\n", string(nextSequence))

	return nextSequence
}

func findSequenceOfKeyPressesToAchieveRemoteSequence(
	remoteSequence []rune,
	keyToKeyNavigations keyToKeyNavigationLookup) []rune {

	sequenceOfSequences := make([][]rune, 0, len(remoteSequence))
	remoteCurrentKey := 'A'
	for _, element := range remoteSequence {
		subsequence := []rune{}
		subsequence = append(subsequence, keyToKeyNavigations[remoteCurrentKey][element]...)
		remoteCurrentKey = element
		subsequence = append(subsequence, 'A')

		// fmt.Printf("\tSequence to press %c:\n", element)
		// fmt.Printf("\t -> %s\n", string(subsequence))

		sequenceOfSequences = append(sequenceOfSequences, subsequence)
	}

	flattenedSequence := []rune{}
	for _, subsequence := range sequenceOfSequences {
		flattenedSequence = append(flattenedSequence, subsequence...)
	}

	return flattenedSequence
}
