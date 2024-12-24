package day21

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

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

const (
	numberKeypadLayout = `789
456
123
 0A`

	directionalKeypadLayout = ` ^A
<v>`
)

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

func findComplexityScoreOfEnteringCodes(codes []unlockCode, nRobots int) (int64, error) {
	numberPadNavigationOptions := newKeyToNavigationMap(numberKeyLocations)
	directionPadNavigationOptions := newKeyToNavigationMap(directionalKeyLocations)

	totalComplexity := int64(0)
	for _, code := range codes {
		value, err := code.NumericValue()
		if err != nil {
			return 0, err
		}

		sequence := findFinalSequenceLengthToEnterCode(numberPadNavigationOptions, directionPadNavigationOptions, code)

		totalComplexity += value * int64(len(sequence))

		// fmt.Printf("totalComplexity: %v (= %d * %d)\n", totalComplexity, value, len(sequence))
	}

	return totalComplexity, nil
}

type keypadSim struct {
	NavigationOptions keyToKeyNavigationOptionsLookup
	KeyLocations      map[rune]vec2
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

type keypadLayer struct {
	Keypad keypadSim
	cache  map[string]transformationCacheEntry
}

type mappingLayer struct {
	Transformation keyToKeyNavigationOptionsLookup
	Cache          map[string]transformationCacheEntry
}

type transformationCacheEntry struct {
	Sequence          []rune
	EndSequenceLength int
}

func findFinalSequenceLengthToEnterCode(
	numberPadNavigationOptions keyToKeyNavigationOptionsLookup,
	directionPadNavigationOptions keyToKeyNavigationOptionsLookup,
	code unlockCode) []rune {

	mappingLayers := []mappingLayer{
		{Transformation: numberPadNavigationOptions, Cache: map[string]transformationCacheEntry{}},
		{Transformation: directionPadNavigationOptions, Cache: map[string]transformationCacheEntry{}},
		{Transformation: directionPadNavigationOptions, Cache: map[string]transformationCacheEntry{}},
	}

	primeCaches(mappingLayers, code)

	return recreateFinalSequence(code, mappingLayers)
}

func primeCaches(mappingLayers []mappingLayer, remoteSequence []rune) {
	type stackFrame struct {
		RemoteSequence []rune
		I              int
		Visited        bool
	}

	stack := queue.NewLifo[stackFrame]()
	stack.Push(stackFrame{RemoteSequence: remoteSequence, I: 0, Visited: false})

	for stack.Len() > 0 {
		frame, ok := stack.Pop()
		if !ok {
			break
		}

		if frame.I >= len(mappingLayers) {
			continue
		}

		layer := &(mappingLayers[frame.I])

		if _, ok := layer.Cache[string(frame.RemoteSequence)]; ok {
			continue
		}

		candidateSequences := generateLocalSequenceCombinations(frame.RemoteSequence, layer)

		if frame.Visited {
			bestCandidate := candidateSequences[0]
			bestEndSequenceLength := math.MaxInt
			if frame.I < len(mappingLayers)-1 {
				for _, candidate := range candidateSequences {
					if cacheEntry, ok := mappingLayers[frame.I+1].Cache[string(candidate)]; ok && cacheEntry.EndSequenceLength < bestEndSequenceLength {
						bestEndSequenceLength = cacheEntry.EndSequenceLength
						bestCandidate = candidate
					}
				}
			} else {
				bestEndSequenceLength = len(bestCandidate)
				for _, candidate := range candidateSequences {
					if len(candidate) < len(bestCandidate) {
						bestCandidate = candidate
						bestEndSequenceLength = len(bestCandidate)
					}
				}
			}

			layer.Cache[string(frame.RemoteSequence)] = transformationCacheEntry{
				Sequence:          bestCandidate,
				EndSequenceLength: bestEndSequenceLength,
			}
			continue
		}

		frame.Visited = true
		stack.Push(stackFrame{RemoteSequence: frame.RemoteSequence, I: frame.I, Visited: true})

		for _, candidate := range candidateSequences {
			localSequence := append(make([]rune, 0, len(candidate)), candidate...)

			stack.Push(stackFrame{
				RemoteSequence: localSequence,
				I:              frame.I + 1,
				Visited:        false,
			})
		}
	}
}

func generateLocalSequenceCombinations(remoteSequence []rune, layer *mappingLayer) [][]rune {
	type stackFrame struct {
		I        int
		Sequence []rune
	}

	stack := queue.NewLifo[stackFrame]()
	stack.Push(stackFrame{I: 0, Sequence: []rune{}})

	combinations := [][]rune{}
	for stack.Len() > 0 {
		frame, ok := stack.Pop()
		if !ok {
			break
		}

		if frame.I >= len(remoteSequence) {
			combinations = append(combinations, frame.Sequence)
			continue
		}

		remoteSrc := 'A'
		remoteDst := remoteSequence[frame.I]
		if frame.I > 0 {
			remoteSrc = remoteSequence[frame.I-1]
		}

		for _, candidate := range layer.Transformation[remoteSrc][remoteDst] {
			extendedSequence := append(make([]rune, 0, len(frame.Sequence)+len(candidate)+1), frame.Sequence...)
			extendedSequence = append(extendedSequence, candidate...)
			extendedSequence = append(extendedSequence, 'A')

			stack.Push(stackFrame{I: frame.I + 1, Sequence: extendedSequence})
		}
	}

	return combinations
}

func recreateFinalSequence(remoteSequence []rune, layers []mappingLayer) []rune {
	sequence := remoteSequence
	for i := range len(layers) {
		// fmt.Printf("%d: %s\n", i, string(sequence))
		// for key, value := range layers[i].Cache {
		// 	fmt.Printf("\t%s -> (%d) %s\n", string(key), value.EndSequenceLength, string(value.Sequence))
		// }

		sequence = layers[i].Cache[string(sequence)].Sequence
	}

	return sequence
}
