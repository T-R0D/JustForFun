package day21

import (
	"fmt"
	"math"
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
	numberPadNavigations := newKeyToNavigationMap(numberKeyLocations)
	directionPadNavigations := newKeyToNavigationMap(directionalKeyLocations)

	totalComplexity := int64(0)
	for _, code := range codes {
		value, err := code.NumericValue()
		if err != nil {
			return 0, err
		}

		sequence := findFinalSequenceLengthToEnterCode(numberPadNavigations, directionPadNavigations, code)

		totalComplexity += value * int64(len(sequence))
	}

	return totalComplexity, nil
}

type keyToKeyNavigationOptionsLookup map[rune]keyToNavigationOptionsLookup

type keyToNavigationOptionsLookup map[rune]navigationOptionList

type navigationOptionList [][]rune

func newKeyToNavigationMap(keyLocations map[rune]vec2) keyToKeyNavigationOptionsLookup {
	keyToNavigation := keyToKeyNavigationOptionsLookup{}
	for src := range keyLocations {
		keyToNavigationOptions := keyToNavigationOptionsLookup{}
		for dst := range keyLocations {
			keyToNavigationOptions[dst] = navigateFromSrcKeyToDstKey(keyLocations, src, dst)
		}

		keyToNavigation[src] = keyToNavigationOptions
	}

	return keyToNavigation
}

func navigateFromSrcKeyToDstKey(keyLocations map[rune]vec2, src rune, dst rune) navigationOptionList {
	locationToKey := map[vec2]rune{}
	for key, location := range keyLocations {
		locationToKey[location] = key
	}

	srcLocation := keyLocations[src]
	dstLocation := keyLocations[dst]

	type searchState struct {
		Location vec2
		Path     []rune
		RunScore int
	}

	cmpSearchState := func(a searchState, b searchState) bool {
		if len(a.Path) != len(b.Path) {
			return len(a.Path) < len(b.Path)
		}

		return a.RunScore > b.RunScore
	}

	frontier := queue.NewPriority[searchState](cmpSearchState)
	frontier.Push(searchState{
		Location: srcLocation,
		Path:     []rune{},
		RunScore: 0,
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

	options := navigationOptionList{}

	targetPathLength := srcLocation.Distance(dstLocation)
	targetRunScore := (int(math.Abs(float64(srcLocation[0]-dstLocation[0]))) - 1 +
		int(math.Abs(float64(srcLocation[1]-dstLocation[1]))) - 1)

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		if len(state.Path) > targetPathLength {
			continue
		}

		if len(state.Path) == targetPathLength && state.RunScore < targetRunScore {
			continue
		}

		if key, ok := locationToKey[state.Location]; !ok {
			continue
		} else if key == dst {
			options = append(options, state.Path)
			continue
		}

		for _, d := range deltas {
			newLocation := state.Location.Plus(d.Delta)
			newPath := append([]rune{}, state.Path...)
			newPath = append(newPath, d.Direction)
			frontier.Push(searchState{
				Location: newLocation,
				Path:     newPath,
				RunScore: getPathRunScore(newPath),
			})
		}
	}

	return options
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

type mappingLayer struct {
	Transformation keyToKeyNavigationOptionsLookup
	Cache          map[string][]rune
}

func findFinalSequenceLengthToEnterCode(
	numberPadNavigationOptions keyToKeyNavigationOptionsLookup,
	directionPadNavigationOptions keyToKeyNavigationOptionsLookup,
	code unlockCode) []rune {

	mappingLayers := []mappingLayer{
		{Transformation: numberPadNavigationOptions, Cache: map[string][]rune{}},
		{Transformation: directionPadNavigationOptions, Cache: map[string][]rune{}},
		{Transformation: directionPadNavigationOptions, Cache: map[string][]rune{}},
	}

	primeCaches(mappingLayers, 0, code)

	return recreateFinalSequence(code, mappingLayers, 0)
}

func primeCaches(
	mappingLayers []mappingLayer,
	currentLayer int,
	remoteSequence []rune) []rune {

	if currentLayer >= len(mappingLayers) {
		return []rune{}
	}

	layer := mappingLayers[currentLayer]
	key := string(remoteSequence)

	if sequence, ok := layer.Cache[key]; ok {
		fmt.Printf("remoteSequence: % 10s | resultingSequence: % 20s (cache hit)\n", string(remoteSequence), string(sequence))

		return sequence
	}

	candidateSequences := generateLocalSequenceCombinations(remoteSequence, 0, layer)

	for _, candidateSequence := range candidateSequences {
		
	}


	finalSequence := []rune{}
	currentRemoteKey := 'A'
	bestLevelAboveSequenceLengthLen := 999_999
	for _, element := range remoteSequence {
		bestLocalSequence := append([]rune{}, layer.Transformation[currentRemoteKey][element][0]...)
		bestLocalSequence = append(bestLocalSequence, 'A')
		for _, candidate := range layer.Transformation[currentRemoteKey][element] {
			candidateLocalSequence := append([]rune{}, candidate...)
			candidateLocalSequence = append(candidateLocalSequence, 'A')
			levelAboveSequence := primeCaches(mappingLayers, currentLayer+1, candidateLocalSequence)

			if len(levelAboveSequence) < bestLevelAboveSequenceLengthLen {
				bestLocalSequence = candidateLocalSequence
				bestLevelAboveSequenceLengthLen = len(levelAboveSequence)
			}
		}

		finalSequence = append(finalSequence, bestLocalSequence...)

		currentRemoteKey = element
	}

	layer.Cache[key] = finalSequence

	fmt.Printf("remoteSequence: % 10s | resultingSequence: % 20s\n", string(remoteSequence), string(finalSequence))

	return finalSequence
}

func recreateFinalSequence(remoteSequence []rune, layers []mappingLayer, i int) []rune {
	fmt.Printf("layer %d, remoteSequence: '%s'\n", i, string(remoteSequence))
	for key, result := range layers[i].Cache {
		fmt.Printf("\t%s -> %s\n", key, string(result))
	}

	if i >= len(layers) {
		return remoteSequence
	}

	localSequence := layers[i].Cache[string(remoteSequence)]

	return recreateFinalSequence(localSequence, layers, i+1)
}

func generateLocalSequenceCombinations(remoteSequence []rune, i int, layer mappingLayer) [][]rune {
	if i >= len(remoteSequence) {
		return [][]rune{}
	}

	sequencesSoFar := [][]rune{}
	remoteSrc := 'A'
	remoteDst := remoteSequence[i]
	if i > 0 {
		remoteSrc = remoteSequence[i - 1]
	}

	for _, navigationOption := range layer.Transformation[remoteSrc][remoteDst] {
		navigation := append([]rune{}, navigationOption...)
		navigation = append(navigation, 'A')
		sequencesSoFar = append(sequencesSoFar, navigation)
	}

	sequenceSuffixes := generateLocalSequenceCombinations(remoteSequence, i+1, layer)

	finalSequences := make([][]rune, 0, len(sequencesSoFar) * len(sequenceSuffixes))
	for _, sequenceSoFar := range sequencesSoFar {
		for _, suffix := range sequenceSuffixes {
			sequence := append([]rune{}, sequenceSoFar...)
			sequence = append(sequence, suffix...)
			finalSequences = append(finalSequences, sequence)
		}
	}

	return finalSequences
}
