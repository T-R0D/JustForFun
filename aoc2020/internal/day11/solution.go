// Possible Improvement: Use arrays instead of maps.
// Upping the Ante: Multiple threads.

package day11

import (
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/pkg/errors"
)

// Solver solves the day's problems.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	occupancyMap, err := parseMapIntoLookup(input)
	if err != nil {
		return "", err
	}
	mapChanged := true
	for mapChanged {
		newOccupancyMap := performSeatChangeIteration(occupancyMap)

		mapChanged = !occupancyMap.Equal(newOccupancyMap)
		occupancyMap = newOccupancyMap
	}

	return strconv.Itoa(occupancyMap.OccupiedSeats()), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	occupancyMap, err := parseMapIntoLookup(input)
	if err != nil {
		return "", err
	}

	I, J := occupancyMap.Bounds()

	mapChanged := true
	for mapChanged {
		newOccupancyMap := performTolerantSeatChangeIteration(occupancyMap, I, J)

		mapChanged = !occupancyMap.Equal(newOccupancyMap)
		occupancyMap = newOccupancyMap
	}

	return strconv.Itoa(occupancyMap.OccupiedSeats()), nil
}

const (
	emptyFloor   = '.'
	emptySeat    = 'L'
	occupiedSeat = '#'

	lineSeparator = '\n'
)

type seatOccupancyMap map[grid.Point]bool

func parseMapIntoLookup(input string) (seatOccupancyMap, error) {
	lookup := map[grid.Point]bool{}
	i, j := 0, 0
	for _, r := range input {
		switch r {
		case lineSeparator:
			j = 0
			i++
		case emptyFloor:
			j++
		case emptySeat:
			lookup[grid.Point{I: i, J: j}] = false
			j++
		case occupiedSeat:
			lookup[grid.Point{I: i, J: j}] = true
			j++
		default:
			return nil, errors.New("unrecognized symbol")
		}
	}
	return lookup, nil
}

func performSeatChangeIteration(seatOccupancies seatOccupancyMap) seatOccupancyMap {
	newSeatOccupancies := seatOccupancyMap{}

	for seat, occupied := range seatOccupancies {
		nOccupiedAdjacentSeats := countOccupiedAdjacentSeats(seatOccupancies, seat)

		if !occupied && nOccupiedAdjacentSeats == 0 {
			newSeatOccupancies[seat] = true
		} else if occupied && nOccupiedAdjacentSeats >= 5 {
			newSeatOccupancies[seat] = false
		} else {
			newSeatOccupancies[seat] = occupied
		}
	}

	return newSeatOccupancies
}

func performTolerantSeatChangeIteration(seatOccupancies seatOccupancyMap, I int, J int) seatOccupancyMap {
	newSeatOccupancies := seatOccupancyMap{}

	for seat, occupied := range seatOccupancies {
		nOccupiedVisibleSeats := countOccupiedVisibleSeats(seatOccupancies, I, J, seat)

		if !occupied && nOccupiedVisibleSeats == 0 {
			newSeatOccupancies[seat] = true
		} else if occupied && nOccupiedVisibleSeats >= 5 {
			newSeatOccupancies[seat] = false
		} else {
			newSeatOccupancies[seat] = occupied
		}
	}

	return newSeatOccupancies
}

func countOccupiedAdjacentSeats(seatOccupancies seatOccupancyMap, seat grid.Point) int {
	I, J := seat.I, seat.J
	occupiedAdjacentSeats := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			neigborSeat := grid.Point{I: I + i, J: J + j}
			if neigborSeatOccupied, ok := seatOccupancies[neigborSeat]; ok && neigborSeatOccupied {
				occupiedAdjacentSeats++
			}
		}
	}
	return occupiedAdjacentSeats
}

func countOccupiedVisibleSeats(seatOccupancies seatOccupancyMap, I int, J int, seat grid.Point) int {
	occupiedVisibleSeats := 0

	for i := seat.I - 1; i >= 0; i-- {
		possibleSeat := grid.Point{I: i, J: seat.J}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for i, j := seat.I-1, seat.J+1; i >= 0 && j < J; i, j = i-1, j+1 {
		possibleSeat := grid.Point{I: i, J: j}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for j := seat.J + 1; j <= J; j++ {
		possibleSeat := grid.Point{I: seat.I, J: j}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for i, j := seat.I+1, seat.J+1; i < I && j < J; i, j = i+1, j+1 {
		possibleSeat := grid.Point{I: i, J: j}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for i := seat.I + 1; i < I; i++ {
		possibleSeat := grid.Point{I: i, J: seat.J}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for i, j := seat.I+1, seat.J-1; i < I && j >= 0; i, j = i+1, j-1 {
		possibleSeat := grid.Point{I: i, J: j}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for j := seat.J - 1; j >= 0; j-- {
		possibleSeat := grid.Point{I: seat.I, J: j}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	for i, j := seat.I-1, seat.J-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		possibleSeat := grid.Point{I: i, J: j}
		if occupied, ok := seatOccupancies[possibleSeat]; ok {
			if occupied {
				occupiedVisibleSeats++
			}
			break
		}
	}

	return occupiedVisibleSeats
}

func (m seatOccupancyMap) Equal(other seatOccupancyMap) bool {
	if len(m) != len(other) {
		return false
	}

	for seat, occupied := range m {
		if otherOccupied, ok := other[seat]; !ok || occupied != otherOccupied {
			return false
		}
	}

	return true
}

func (m seatOccupancyMap) OccupiedSeats() int {
	occupiedSeats := 0
	for _, occupied := range m {
		if occupied {
			occupiedSeats++
		}
	}
	return occupiedSeats
}

func (m seatOccupancyMap) Bounds() (int, int) {
	if len(m) == 0 {
		return 0, 0
	}

	I, J := 0, 0
	for seat := range m {
		if seat.I > I {
			I = seat.I
		}
		if seat.J > J {
			J = seat.J
		}
	}
	I++
	J++

	return I, J
}

func (m seatOccupancyMap) String() string {
	if len(m) == 0 {
		return ""
	}

	I, J := 0, 0
	for seat := range m {
		if seat.I > I {
			I = seat.I
		}
		if seat.J > J {
			J = seat.J
		}
	}
	I++
	J++

	builder := strings.Builder{}
	for i := 0; i < I; i++ {
		for j := 0; j < J; j++ {
			seat := grid.Point{I: i, J: j}
			if occupied, ok := m[seat]; ok {
				if occupied {
					builder.WriteRune(occupiedSeat)
				} else {
					builder.WriteRune(emptySeat)
				}
			} else {
				builder.WriteRune(emptyFloor)
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
