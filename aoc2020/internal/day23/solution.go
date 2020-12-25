// Possible Improvement: I feel like it could maybe be faster somehow...

package day23

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver saves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	cupOrder, err := readCupOrder(input)
	if err != nil {
		return "", err
	}

	ring := newCupRing2(cupOrder, len(cupOrder))

	for i := 0; i < oneHundredIterations; i++ {
		ring.MakeMove()
	}

	return ring.ReadCupsStartingAfter(startAfterLabel), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	cupOrder, err := readCupOrder(input)
	if err != nil {
		return "", err
	}

	ring := newCupRing2(cupOrder, oneMillionCups)

	for i := 0; i < temMillionIterations; i++ {
		ring.MakeMove()
	}

	twoCupsLeftOfOne := ring.Read2CupsStartingAfter(startAfterLabel)

	product := uint64(1)
	for _, cup := range twoCupsLeftOfOne {
		product *= uint64(cup)
	}

	return fmt.Sprint(product), nil
}

func readCupOrder(input string) ([]int, error) {
	labelStrs := strings.Split(input, "")
	cupOrder := make([]int, len(labelStrs))
	for i, labelStr := range labelStrs {
		value, err := strconv.Atoi(labelStr)
		if err != nil {
			return nil, errors.Wrapf(err, "value at position %d is not a number", i)
		}
		cupOrder[i] = value
	}
	return cupOrder, nil
}

const (
	oneHundredIterations = 100
	temMillionIterations = 10000000

	startAfterLabel = 1
	oneMillionCups  = 1000000
)

type cupRing struct {
	current      *cupNode
	labels       []int
	lowestLabel  int
	highestLabel int
}

type cupNode struct {
	Label            int
	CounterClockwise *cupNode
	Clockwise        *cupNode
}

func newCupRing(labels []int) *cupRing {
	var current *cupNode = nil
	var previous *cupNode = nil
	for i, label := range labels {
		newest := &cupNode{
			Label:            label,
			CounterClockwise: nil,
			Clockwise:        nil,
		}

		if i == 0 {
			current = newest
		} else {
			newest.CounterClockwise = previous
			previous.Clockwise = newest
		}
		previous = newest
	}
	current.CounterClockwise = previous
	previous.Clockwise = current

	labelsCopy := make([]int, 0, len(labels))
	labelsCopy = append([]int{}, labels...)
	sort.Ints(labelsCopy)

	return &cupRing{
		current:      current,
		labels:       labelsCopy,
		lowestLabel:  labelsCopy[0],
		highestLabel: labelsCopy[len(labelsCopy)-1],
	}
}

func newBigCupRing(labels []int) *cupRing {
	labelsCopy := make([]int, 0, len(labels))
	labelsCopy = append([]int{}, labels...)
	sort.Ints(labelsCopy)

	var current *cupNode = nil
	var previous *cupNode = nil
	for i, label := range labels {
		newest := &cupNode{
			Label:            label,
			CounterClockwise: nil,
			Clockwise:        nil,
		}

		if i == 0 {
			current = newest
		} else {
			newest.CounterClockwise = previous
			previous.Clockwise = newest
		}
		previous = newest
	}
	for i := labelsCopy[len(labelsCopy)-1]; i <= oneMillionCups; i++ {
		newest := &cupNode{
			Label:            i,
			CounterClockwise: nil,
			Clockwise:        nil,
		}
		newest.CounterClockwise = previous
		previous.Clockwise = newest
		previous = newest
	}

	current.CounterClockwise = previous
	previous.Clockwise = current

	return &cupRing{
		current:      current,
		labels:       labelsCopy,
		lowestLabel:  labelsCopy[0],
		highestLabel: oneMillionCups,
	}
}

func (r *cupRing) MakeMove() {
	removedCups := r.removeThreeCupsRightOfCurrent()

	destinationLabel := r.current.Label - 1
	for destinationLabel == r.current.Label ||
		r.labelInRange(destinationLabel, removedCups) ||
		destinationLabel < r.labels[0] {

		if destinationLabel < r.lowestLabel {
			destinationLabel = r.highestLabel
		} else {
			destinationLabel--
		}
	}

	destination := r.findNodeWithLabel(destinationLabel)

	r.mergeCups(destination, removedCups)

	r.current = r.current.Clockwise
}

func (r *cupRing) removeThreeCupsRightOfCurrent() [2]*cupNode {
	start := r.current.Clockwise
	end := start.Clockwise.Clockwise

	r.current.Clockwise = end.Clockwise
	end.Clockwise.CounterClockwise = r.current

	start.CounterClockwise = nil
	end.Clockwise = nil

	return [2]*cupNode{start, end}
}

func (r *cupRing) labelInRange(label int, cupRange [2]*cupNode) bool {
	cursor := cupRange[0]
	for keepLooking := true; keepLooking; {
		if cursor.Label == label {
			return true
		}
		cursor = cursor.Clockwise
		keepLooking = cursor != nil && cursor != cupRange[0]
	}
	return false
}

func (r *cupRing) findNodeWithLabel(label int) *cupNode {
	cursor := r.current
	for keepLooking := true; keepLooking; {
		if cursor.Label == label {
			return cursor
		}
		cursor = cursor.Clockwise
		keepLooking = cursor != r.current
	}
	return nil
}

func (r *cupRing) mergeCups(destination *cupNode, removedCups [2]*cupNode) {
	start, end := removedCups[0], removedCups[1]

	end.Clockwise = destination.Clockwise
	destination.Clockwise.CounterClockwise = end
	destination.Clockwise = start
	start.CounterClockwise = destination
}

func (r *cupRing) ReadCupsStartingAfter(startAfter int) string {
	startAfterNode := r.findNodeWithLabel(startAfter)
	builder := strings.Builder{}
	cursor := startAfterNode.Clockwise
	for keepLooking := true; keepLooking; {
		builder.WriteString(strconv.Itoa(cursor.Label))
		cursor = cursor.Clockwise
		keepLooking = cursor != startAfterNode
	}
	return builder.String()
}

func (r *cupRing) Read2CupsStartingAfter(startingAfter int) [2]int {
	startingAfterNode := r.findNodeWithLabel(startingAfter)
	return [2]int{
		startingAfterNode.Clockwise.Label,
		startingAfterNode.Clockwise.Clockwise.Label,
	}
}

func (r *cupRing) String() string {
	builder := strings.Builder{}
	cursor := r.current
	for keepLooking := true; keepLooking; {
		builder.WriteString(strconv.Itoa(cursor.Label))
		cursor = cursor.Clockwise
		keepLooking = cursor != r.current
	}
	return builder.String()
}

type cupRing2 struct {
	cups         []int
	current      int
	highestLabel int
	lowestLabel  int
}

func newCupRing2(startingLabels []int, nCups int) *cupRing2 {
	sortedStartingLabels := []int{}
	sortedStartingLabels = append(sortedStartingLabels, startingLabels...)
	sort.Ints(sortedStartingLabels)

	cups := make([]int, nCups)
	for i, label := range startingLabels {
		cups[label-1] = startingLabels[(i+1)%len(startingLabels)]
	}
	if nCups > len(startingLabels) {
		cups[startingLabels[len(startingLabels)-1]-1] = sortedStartingLabels[len(sortedStartingLabels)-1] + 1
		for i := len(startingLabels); i < nCups; i++ {
			// +1 to get the label we are working on, +1 more to get the next one.
			cups[i] = i + 2
		}
		cups[nCups-1] = startingLabels[0]
	}

	return &cupRing2{
		cups:         cups,
		current:      startingLabels[0],
		highestLabel: nCups,
		lowestLabel:  sortedStartingLabels[0],
	}
}

func (r *cupRing2) MakeMove() {
	// Remove the cups from the ring; reaform the ring.
	firstRemovedCup := r.nextCup(r.current)
	lastRemovedCup := r.nextCup(r.nextCup(firstRemovedCup))
	r.setNextCup(r.current, r.nextCup(lastRemovedCup))

	// Select the destination cup.
	destination := r.current - 1
	for destination == r.current ||
		destination == firstRemovedCup ||
		destination == r.nextCup(firstRemovedCup) ||
		destination == lastRemovedCup ||
		destination < r.lowestLabel {

		if destination < r.lowestLabel {
			destination = r.highestLabel
		} else {
			destination--
		}
	}

	// Re-insert the removed cups.
	r.setNextCup(lastRemovedCup, r.nextCup(destination))
	r.setNextCup(destination, firstRemovedCup)

	// Update the current cup.
	r.current = r.nextCup(r.current)
}

func (r *cupRing2) nextCup(label int) int {
	return r.cups[label-1]
}

func (r *cupRing2) setNextCup(label int, nextCup int) {
	r.cups[label-1] = nextCup
}

func (r *cupRing2) ReadCupsStartingAfter(startingAfter int) string {
	builder := strings.Builder{}
	nextCup := r.nextCup(startingAfter)
	for keepGoing := true; keepGoing; {
		builder.WriteString(fmt.Sprint(nextCup))
		nextCup = r.nextCup(nextCup)
		keepGoing = nextCup != startingAfter
	}
	return builder.String()
}

func (r *cupRing2) Read2CupsStartingAfter(startingAfter int) [2]int {
	ret := [2]int{}
	ret[0] = r.nextCup(startingAfter)
	ret[1] = r.nextCup(ret[0])
	return ret
}
