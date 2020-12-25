// Possible Improvement: Just convert all of the seat IDs to binary and work with that.
// Possible Improvement: Find the empty seat in part 2 by finding the sum of all seat IDs
//                       in the filled range and comparing that to the sum of seats that are filled.

package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	rawBoardingIDs := strings.Split(input, "\n")

	highestSeatID := -1
	for i, rawBoardingID := range rawBoardingIDs {
		boardingID, err := boardingIDFromString(rawBoardingID)
		if err != nil {
			return "", errors.Wrapf(err, "error processing boarding pass %d", i)
		}

		if seatID := boardingID.SeatID(); seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	if highestSeatID == -1 {
		return "", errors.New("didn't find a valid seat ID")
	}

	return strconv.Itoa(highestSeatID), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	rawBoardingIDs := strings.Split(input, "\n")

	takenSeatIDs := make([]int, 0, len(rawBoardingIDs))
	for i, rawBoardingID := range rawBoardingIDs {
		boardingID, err := boardingIDFromString(rawBoardingID)
		if err != nil {
			return "", errors.Wrapf(err, "error processing boarding pass %d", i)
		}
		takenSeatIDs = append(takenSeatIDs, boardingID.SeatID())
	}

	sort.Ints(takenSeatIDs)

	previousID := takenSeatIDs[0]
	for _, currentID := range takenSeatIDs[0:] {
		if currentID == previousID+2 {
			return strconv.Itoa(previousID + 1), nil
		}

		previousID = currentID
	}

	return "", errors.New("no gap in contiguous IDs found")
}

const (
	back  = 'B'
	front = 'F'

	left  = 'L'
	right = 'R'
)

type boardingID struct {
	column         int
	representation string
	row            int
}

func boardingIDFromString(candidate string) (*boardingID, error) {
	if strLen := len(candidate); strLen != 10 {
		return nil, errors.Errorf("boarding ID has invalid length: %d", strLen)
	}

	rowNumber, err := findRowNumber(candidate[0:7])
	if err != nil {
		return nil, err
	}

	columnNumber, err := findColumnNumber(candidate[7:])
	if err != nil {
		return nil, err
	}

	return &boardingID{
		column:         columnNumber,
		representation: candidate,
		row:            rowNumber,
	}, nil
}

func (b *boardingID) String() string {
	return b.representation
}

func (b *boardingID) SeatID() int {
	return (b.row * 8) + b.column
}

func findRowNumber(boardingIDRowSegment string) (int, error) {
	return findBinaryPartition(boardingIDRowSegment, 0, 127, front, back)
}

func findColumnNumber(boardingIDColumnSegment string) (int, error) {
	return findBinaryPartition(boardingIDColumnSegment, 0, 7, left, right)
}

func findBinaryPartition(partitionIDSegment string, lowerBound int, upperBound int, lowerSignal rune, upperSignal rune) (int, error) {
	for _, r := range partitionIDSegment {
		switch r {
		case lowerSignal:
			lowerBound, upperBound = findBinaryPartitionStep(lowerBound, upperBound, true)
		case upperSignal:
			lowerBound, upperBound = findBinaryPartitionStep(lowerBound, upperBound, false)
		default:
			return 0, errors.Errorf("invalid character in ID segment: %v", r)
		}
		if upperBound < lowerBound {
			return 0, errors.Errorf("something went wrong with the binary partition finding, lower=%d, upper=%d", lowerBound, upperBound)
		}
	}

	if upperBound != lowerBound {
		return 0, errors.Errorf("something went wrong with the binary partition finding, lower=%d, upper=%d", lowerBound, upperBound)
	}

	return lowerBound, nil
}

func findBinaryPartitionStep(lowerBound int, upperBound int, takeLowerHalf bool) (int, int) {
	if takeLowerHalf {
		return lowerBound, (lowerBound + upperBound) / 2
	}
	return (lowerBound+upperBound)/2 + 1, upperBound
}
