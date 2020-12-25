// Question: I feel like the API made was hacky. Can it be improved?

package day09

import (
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/queue"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	xmasData, err := parseXMASData(input)
	if err != nil {
		return "", err
	}

	processor := newXMASProcessor(xmasData, specPreambleSize)

	for {
		if currentValueIsValid, err := processor.CurrentDataIsValidXMAS(); err != nil {
			return "", err
		} else if !currentValueIsValid {
			return strconv.Itoa(processor.CurrentData()), nil
		}

		if advanceOccurred, err := processor.Advance(); err != nil {
			return "", err
		} else if !advanceOccurred {
			return "", errors.Errorf("xmas code was exhausted without finding an error")
		}
	}
}

// Part2 solves part 2 of the day's problem.
// 291116327 t00 low
func (s *Solver) Part2(input string) (string, error) {
	xmasData, err := parseXMASData(input)
	if err != nil {
		return "", err
	}

	processor := newXMASProcessor(xmasData, specPreambleSize)

	for {
		if currentValueIsValid, err := processor.CurrentDataIsValidXMAS(); err != nil {
			return "", err
		} else if !currentValueIsValid {
			target := processor.CurrentData()

			pair, err := findMaxMinInSumRange(xmasData, target)
			if err != nil {
				return "", err
			}

			pairSum := pair.max + pair.min

			return strconv.Itoa(pairSum), nil
		}

		if advanceOccurred, err := processor.Advance(); err != nil {
			return "", err
		} else if !advanceOccurred {
			return "", errors.Errorf("xmas code was exhausted without finding an error")
		}
	}
}

const specPreambleSize = 25

func parseXMASData(input string) ([]int, error) {
	lines := strings.Split(input, "\n")
	xmasData := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.Wrapf(err, "line %d: %s could not be converted to int", i, line)
		}
		xmasData[i] = num
	}
	return xmasData, nil
}

type xmasProcessor struct {
	cursor          int
	data            []int
	preambleMembers map[int]struct{}
	preamble        queue.Queue
	preambleSize    int
}

func newXMASProcessor(data []int, preambleSize int) *xmasProcessor {
	processor := &xmasProcessor{
		cursor:          0,
		data:            data,
		preambleMembers: map[int]struct{}{},
		preamble:        nil,
		preambleSize:    preambleSize,
	}

	processor.Init()

	return processor
}

func (xp *xmasProcessor) Init() {
	xp.preamble = queue.NewLinkQueue()
	for _, number := range xp.data[0:xp.preambleSize] {
		xp.preamble.AppendRight(number)
		xp.preambleMembers[number] = struct{}{}
	}
	xp.cursor = xp.preambleSize
}

func (xp *xmasProcessor) Advance() (bool, error) {
	if xp.cursor+1 >= len(xp.data) {
		return false, nil
	}

	newestPreambleMember := xp.data[xp.cursor]
	xp.cursor++

	ejected, err := xp.preamble.PopLeft()
	if err != nil {
		return false, errors.New("error removing from preamble")
	}

	switch ejectedNumber := ejected.(type) {
	case int:
		delete(xp.preambleMembers, ejectedNumber)
	default:
		return false, errors.Errorf("got back something that wasn't an int from the preamble: %v", ejected)
	}

	xp.preambleMembers[newestPreambleMember] = struct{}{}
	xp.preamble.AppendRight(newestPreambleMember)

	return true, nil
}

func (xp *xmasProcessor) CurrentDataIsValidXMAS() (bool, error) {
	if xp.cursor >= len(xp.data) {
		return false, errors.New("xmasProcessor in an invalid state")
	}

	dataItem := xp.data[xp.cursor]
	for preambleMember := range xp.preambleMembers {
		difference := dataItem - preambleMember

		if difference == preambleMember {
			continue
		}

		if _, ok := xp.preambleMembers[difference]; ok {
			return true, nil
		}
	}

	return false, nil
}

func (xp *xmasProcessor) CurrentData() int {
	return xp.data[xp.cursor]
}

type xmasWeaknessPair struct {
	max int
	min int
}

func findMaxMinInSumRange(data []int, target int) (*xmasWeaknessPair, error) {
	if len(data) < 1 {
		return nil, errors.New("can't find sum range for empty data set")
	}

	max, min := -int(^uint(0)>>1)-1, int(^uint(0)>>1)

	start, end, currentSum := 0, 0, 0
	targetSumFound := false
	currentRange := queue.NewLinkQueue()

	for start < len(data) {
		for end < len(data) {
			currentSum += data[end]
			currentRange.AppendRight(data[end])
			if currentSum == target && start < end {
				targetSumFound = true
				break
			}
			end++
		}

		if targetSumFound {
			break
		}

		start++
		end = start
		currentSum = 0
		currentRange = queue.NewLinkQueue()
	}

	if !targetSumFound {
		return nil, errors.New("unable to find a range adding up to the target sum")
	}

	for currentRange.Len() > 0 {
		item, err := currentRange.PopLeft()
		if err != nil {
			return nil, errors.Wrap(err, "while finding min and max")
		}

		itemValue, ok := item.(int)
		if !ok {
			return nil, errors.New("someone stuck something that wasn't an int in the queue")
		}
		if itemValue < min {
			min = itemValue
		}
		if itemValue > max {
			max = itemValue
		}

	}

	return &xmasWeaknessPair{max: max, min: min}, nil
}
