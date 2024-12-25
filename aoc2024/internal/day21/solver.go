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

	totalComplexity, err := findComplexityScoreOfEnteringCodes(codes, 3)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", totalComplexity), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	codes := parseCodes(input)

	totalComplexity, err := findComplexityScoreOfEnteringCodes(codes, 26)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", totalComplexity), nil
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

func findComplexityScoreOfEnteringCodes(codes []unlockCode, nDirectionPadLayers int) (int64, error) {
	numberPadNavigationOptions := newKeyToNavigationMap(numberKeyLocations)
	directionPadNavigationOptions := newKeyToNavigationMap(directionalKeyLocations)

	totalComplexity := int64(0)
	for _, code := range codes {
		value, err := code.NumericValue()
		if err != nil {
			return 0, err
		}

		sequenceLength := findFinalSequenceLengthToEnterCode(
			numberPadNavigationOptions,
			directionPadNavigationOptions,
			nDirectionPadLayers,
			code)

		totalComplexity += value * int64(sequenceLength)
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

type transformationLayer struct {
	Transformation keyToKeyNavigationOptionsLookup
	Cache          transformationCache
}

type transformationCache map[transformationCacheKey]transformationCacheEntry

type transformationCacheKey struct {
	RemoteSequence string
	StartKey       rune
}

type transformationCacheEntry struct {
	Sequence          []rune
	EndSequenceLength int
}

func findFinalSequenceLengthToEnterCode(
	numberPadNavigationOptions keyToKeyNavigationOptionsLookup,
	directionPadNavigationOptions keyToKeyNavigationOptionsLookup,
	nDirectionPadLayers int,
	code unlockCode) int {

	mappingLayers := []transformationLayer{
		{Transformation: numberPadNavigationOptions, Cache: transformationCache{}},
	}

	for range nDirectionPadLayers - 1 {
		mappingLayers = append(mappingLayers, transformationLayer{
			Transformation: directionPadNavigationOptions,
			Cache:          transformationCache{},
		})
	}

	primeCaches(mappingLayers, code)

	return getFinalSequenceLength(code, mappingLayers)
}

func primeCaches(mappingLayers []transformationLayer, remoteSequence []rune) {
	type stackFrame struct {
		RemoteSequence []rune
		StartKey       rune
		LayerIndex     int
		Visited        bool
	}

	stack := queue.NewLifo[stackFrame]()
	stack.Push(stackFrame{
		RemoteSequence: remoteSequence,
		StartKey:       'A',
		LayerIndex:     0,
		Visited:        false,
	})

	for stack.Len() > 0 {
		frame, ok := stack.Pop()
		if !ok {
			break
		}

		if frame.LayerIndex >= len(mappingLayers) {
			continue
		}

		layer := &(mappingLayers[frame.LayerIndex])

		cacheKey := transformationCacheKey{
			RemoteSequence: string(frame.RemoteSequence),
			StartKey:       frame.StartKey,
		}
		if _, ok := layer.Cache[cacheKey]; ok {
			continue
		}

		candidateSequences := generateLocalSequenceCombinations(frame.RemoteSequence, layer)

		if frame.Visited {
			if frame.LayerIndex < len(mappingLayers)-1 {
				bestCandidate := candidateSequences[0]
				bestEndSequenceLength := math.MaxInt
				for _, candidate := range candidateSequences {
					resultingEndLength := 0
					for j, subsequence := range candidate {
						startKey := 'A'
						if frame.LayerIndex == 0 && j > 0 {
							startKey = frame.RemoteSequence[j-1]
						}

						key := transformationCacheKey{
							RemoteSequence: string(subsequence),
							StartKey:       startKey,
						}

						cacheEntry, ok := mappingLayers[frame.LayerIndex+1].Cache[key]
						if !ok {
							panic("how was the cache not primed?")
						}

						resultingEndLength += cacheEntry.EndSequenceLength
					}

					if resultingEndLength < bestEndSequenceLength {
						bestEndSequenceLength = resultingEndLength
						bestCandidate = candidate
					}
				}

				layer.Cache[cacheKey] = transformationCacheEntry{
					Sequence:          flattenSequenceOfSubsequences(bestCandidate),
					EndSequenceLength: bestEndSequenceLength,
				}
			} else {
				bestCandidate := flattenSequenceOfSubsequences(candidateSequences[0])
				bestEndSequenceLength := len(bestCandidate)
				for _, candidate := range candidateSequences {
					flattenedCandidate := flattenSequenceOfSubsequences(candidate)
					if len(flattenedCandidate) < len(bestCandidate) {
						bestCandidate = flattenedCandidate
						bestEndSequenceLength = len(bestCandidate)
					}
				}

				layer.Cache[cacheKey] = transformationCacheEntry{
					Sequence:          bestCandidate,
					EndSequenceLength: bestEndSequenceLength,
				}
			}

			continue
		}

		frame.Visited = true
		stack.Push(stackFrame{
			RemoteSequence: frame.RemoteSequence,
			StartKey:       frame.StartKey,
			LayerIndex:     frame.LayerIndex,
			Visited:        true,
		})

		for _, candidate := range candidateSequences {
			for j, subsequence := range candidate {
				startKey := 'A'
				if frame.LayerIndex == 0 && j > 0 {
					startKey = frame.RemoteSequence[j-1]
				}

				localSequence := append(make([]rune, 0, len(subsequence)), subsequence...)
				stack.Push(stackFrame{
					RemoteSequence: localSequence,
					StartKey:       startKey,
					LayerIndex:     frame.LayerIndex + 1,
					Visited:        false,
				})
			}
		}
	}
}

func generateLocalSequenceCombinations(remoteSequence []rune, layer *transformationLayer) [][][]rune {
	type stackFrame struct {
		LayerIndex int
		Sequence   [][]rune
	}

	stack := queue.NewLifo[stackFrame]()
	stack.Push(stackFrame{LayerIndex: 0, Sequence: [][]rune{}})

	combinations := [][][]rune{}
	for stack.Len() > 0 {
		frame, ok := stack.Pop()
		if !ok {
			break
		}

		if frame.LayerIndex >= len(remoteSequence) {
			combinations = append(combinations, frame.Sequence)
			continue
		}

		remoteSrc := 'A'
		remoteDst := remoteSequence[frame.LayerIndex]
		if frame.LayerIndex > 0 {
			remoteSrc = remoteSequence[frame.LayerIndex-1]
		}

		for _, candidate := range layer.Transformation[remoteSrc][remoteDst] {
			extendedSequence := append(make([][]rune, 0, len(frame.Sequence)+len(candidate)+1), frame.Sequence...)
			nextSubsequence := append(make([]rune, 0, len(candidate)), candidate...)
			nextSubsequence = append(nextSubsequence, 'A')
			extendedSequence = append(extendedSequence, nextSubsequence)

			stack.Push(stackFrame{LayerIndex: frame.LayerIndex + 1, Sequence: extendedSequence})
		}
	}

	return combinations
}

func getFinalSequenceLength(remoteSequence []rune, primedLayers []transformationLayer) int {
	cacheKey := transformationCacheKey{
		RemoteSequence: string(remoteSequence),
		StartKey:       'A',
	}
	entry, ok := primedLayers[0].Cache[cacheKey]
	if !ok {
		panic("the caches must not be primed...")
	}

	return entry.EndSequenceLength
}

func flattenSequenceOfSubsequences(s [][]rune) []rune {
	flattenedSequence := make([]rune, 0, len(s)*20)
	for _, subsequence := range s {
		flattenedSequence = append(flattenedSequence, subsequence...)
	}
	return flattenedSequence
}
