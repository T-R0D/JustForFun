// Improvement: Be less risky about grabbing the first ID in the list in part 1.
//              We are relying too much on nice input.
// Possible Improvment: In part two, apply Chinese Remainder Theorem.

package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	earliestDepartureAndBusIDs, err := parseInput(input)
	if err != nil {
		return "", err
	}

	result := findBusWithIdealDepartureTime(earliestDepartureAndBusIDs.EarliestDepartureTime, earliestDepartureAndBusIDs.BusIDs)

	return strconv.Itoa(result.BusID * result.TimeToWait), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	earliestDepartureAndBusIDs, err := parseInput(input)
	if err != nil {
		return "", err
	}

	magicTimeT := findMagicTimeT(earliestDepartureAndBusIDs.BusIDs)
	return fmt.Sprintf("%d", magicTimeT), nil
}

const (
	brokenDownBusID = -1
)

type earliestDepartureBusIDsPair struct {
	EarliestDepartureTime int
	BusIDs                []int
}

func parseInput(input string) (*earliestDepartureBusIDsPair, error) {
	lines := strings.Split(input, "\n")

	if len(lines) != 2 {
		return nil, errors.Errorf("input did not have correct number of lines: %d", len(lines))
	}

	earliestDepartureTime, err := strconv.Atoi(lines[0])
	if err != nil {
		return nil, errors.Wrap(err, "parsing earliest departure time")
	}

	busIDStrings := strings.Split(lines[1], ",")
	busIDs := make([]int, len(busIDStrings))
	for i, busIDString := range busIDStrings {
		busID := brokenDownBusID
		if busIDString == "x" {
			busIDs[i] = busID
			continue
		}
		busID, err := strconv.Atoi(busIDString)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing bus ID string %d", i)
		}
		busIDs[i] = busID
	}

	return &earliestDepartureBusIDsPair{
		EarliestDepartureTime: earliestDepartureTime,
		BusIDs:                busIDs,
	}, nil
}

type busIDAndTimeToWaitPair struct {
	BusID      int
	TimeToWait int
}

func findBusWithIdealDepartureTime(earliestDepartureTime int, busIDs []int) busIDAndTimeToWaitPair {
	bestBus := busIDs[0]
	bestTimeToWait := findTimeToWaitForBus(earliestDepartureTime, busIDs[0])

	for _, busID := range busIDs[1:] {
		if busID == brokenDownBusID {
			continue
		}

		if timeToWait := findTimeToWaitForBus(earliestDepartureTime, busID); timeToWait < bestTimeToWait {
			bestBus = busID
			bestTimeToWait = timeToWait
		}
	}

	return busIDAndTimeToWaitPair{
		BusID:      bestBus,
		TimeToWait: bestTimeToWait,
	}
}

func findTimeToWaitForBus(earliestDepartureTime int, busID int) int {
	lastBusMissedBy := earliestDepartureTime % busID
	timeToWaitForNextRound := busID - lastBusMissedBy
	return timeToWaitForNextRound
}

type intervalOffsetPair struct {
	Interval uint64
	Offset   uint64
}

// TODO: Outline this algorithm in comments for future readers.
func findMagicTimeT(busSchedule []int) uint64 {
	intervalsAndOffsets := getIntervalAndOffsetList(busSchedule)

	magicTimeT := uint64(1)
	lcmOfProcessedBusIntervals := uint64(1)
	for _, bus := range intervalsAndOffsets {
		magicTimeT = findMagicTimeTForAdditionalBus(magicTimeT, lcmOfProcessedBusIntervals, bus)

		// Knowing that the input only contains prime intervals, we can compute the LCM
		// simply by multiplying.
		lcmOfProcessedBusIntervals *= bus.Interval
	}

	return magicTimeT
}

func getIntervalAndOffsetList(busSchedule []int) []intervalOffsetPair {
	intervalsAndOffsets := []intervalOffsetPair{}
	for i, busID := range busSchedule {
		if busID == brokenDownBusID {
			continue
		}
		intervalsAndOffsets = append(intervalsAndOffsets, intervalOffsetPair{
			Interval: uint64(busID),
			Offset:   uint64(i),
		})
	}
	return intervalsAndOffsets
}

func findMagicTimeTForAdditionalBus(existingMagicTimeT uint64, lcmOfExistingBuses uint64, nextBus intervalOffsetPair) uint64 {
	for testTimeT := existingMagicTimeT; true; testTimeT += lcmOfExistingBuses {
		if (testTimeT+nextBus.Offset)%nextBus.Interval == 0 {
			return testTimeT
		}
	}
	return 0
}
