// Possible Improvement: The handling of the "right" and "left" instructions
//                       for part 2. Maybe rotation matrices? Or repeated 90 degree turns?

package day12

import (
	"math"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	instructions, err := parseInstructions(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing input")
	}

	ferry := newShip()
	for i, instruction := range instructions {
		if err := ferry.TakeInstruction(instruction.key, instruction.value); err != nil {
			return "", errors.Wrapf(err, "when executing instruction %d", i)
		}
	}

	ferryPosition := ferry.Position()
	manhattanDistanceTraveled := int(math.Abs(float64(ferryPosition.I)) + math.Abs(float64(ferryPosition.J)))

	return strconv.Itoa(manhattanDistanceTraveled), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	instructions, err := parseInstructions(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing input")
	}

	ferry := newGuidedShip(grid.Point{I: 0, J: 0}, grid.Point{I: -1, J: 10})
	for i, instruction := range instructions {
		if err := ferry.TakeInstruction(instruction.key, instruction.value); err != nil {
			return "", errors.Wrapf(err, "when executing instruction %d", i)
		}
	}

	ferryPosition := ferry.Position()
	manhattanDistanceTraveled := int(math.Abs(float64(ferryPosition.I)) + math.Abs(float64(ferryPosition.J)))

	return strconv.Itoa(manhattanDistanceTraveled), nil
}

const (
	north = 'N'
	south = 'S'
	east  = 'E'
	west  = 'W'

	left  = 'L'
	right = 'R'

	forward = 'F'
)

type instruction struct {
	key   rune
	value int
}

func parseInstructions(input string) ([]instruction, error) {
	lines := strings.Split(input, "\n")

	instructions := make([]instruction, len(lines))

	for i, line := range lines {
		key := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, errors.Wrapf(err, "instruction %d didn't have a number", i)
		}
		switch key {
		case north, south, east, west, left, right, forward:
		default:
			return nil, errors.Errorf("%v was not a recognized instruction key", key)
		}

		instructions[i].key = rune(key)
		instructions[i].value = value
	}

	return instructions, nil
}

type ship struct {
	position    grid.Point
	orientation int
}

func newShip() *ship {
	return &ship{
		position:    grid.Point{I: 0, J: 0},
		orientation: 0,
	}
}

func (s *ship) TakeInstruction(key rune, value int) error {
	switch key {
	case north:
		s.position.I -= value
	case south:
		s.position.I += value
	case east:
		s.position.J += value
	case west:
		s.position.J -= value
	case right, left:
		newOrientation, err := findNewHeading(s.orientation, key, value)
		if err != nil {
			return err
		}
		s.orientation = newOrientation
	case forward:
		switch s.orientation {
		case 0:
			s.TakeInstruction(east, value)
		case 90:
			s.TakeInstruction(north, value)
		case 180:
			s.TakeInstruction(west, value)
		case 270:
			s.TakeInstruction(south, value)
		default:
			return errors.Errorf("ship is not on a grid world heading: %d", s.orientation)
		}
	default:
		return errors.Errorf("unrecognized instruction key: %v", key)
	}

	return nil
}

func (s *ship) Position() grid.Point {
	return s.position
}

func findNewHeading(currentOrientaiton int, turnDirection rune, magnitude int) (int, error) {
	var newOrientation int
	switch turnDirection {
	case left:
		newOrientation = (currentOrientaiton + magnitude) % 360
	case right:
		newOrientation = (currentOrientaiton - magnitude) % 360
		if newOrientation < 0 {
			newOrientation = 360 + newOrientation
		}
	default:
		return newOrientation, errors.Errorf("unrecognized turn direction: %v", turnDirection)
	}
	return newOrientation, nil
}

// 19427 too high

type guidedShip struct {
	position grid.Point
	waypoint grid.Point
}

func newGuidedShip(position grid.Point, waypoint grid.Point) *guidedShip {
	return &guidedShip{
		position: position,
		waypoint: waypoint,
	}
}

func (s *guidedShip) TakeInstruction(key rune, value int) error {
	switch key {
	case north:
		s.waypoint.I -= value
	case south:
		s.waypoint.I += value
	case east:
		s.waypoint.J += value
	case west:
		s.waypoint.J -= value
	case right, left:
		switch value {
		case 180:
			s.waypoint = grid.Point{
				I: -s.waypoint.I,
				J: -s.waypoint.J,
			}
		case 90, 270:
			quadrant := findQuadrant(s.waypoint)

			quadrantIncrement := 1
			if value == 270 {
				quadrantIncrement = 3
			}

			switch key {
			case right:
				quadrant = (4 + quadrant - quadrantIncrement) % 4
			case left:
				quadrant = (quadrant + quadrantIncrement) % 4
			}

			iSign, jSign := 0, 0
			switch quadrant {
			case 0:
				iSign, jSign = -1, 1
			case 1:
				iSign, jSign = -1, -1
			case 2:
				iSign, jSign = 1, -1
			case 3:
				iSign, jSign = 1, 1
			}

			absWaypointI, absWaypointJ := int(math.Abs(float64(s.waypoint.I))), int(math.Abs(float64(s.waypoint.J)))

			s.waypoint = grid.Point{
				I: iSign * absWaypointJ,
				J: jSign * absWaypointI,
			}

		default:
			return errors.Errorf("invalid rotation specified: %d", value)
		}
	case forward:
		s.position = grid.Point{
			I: s.position.I + (value * s.waypoint.I),
			J: s.position.J + (value * s.waypoint.J),
		}
	default:
		return errors.Errorf("unrecognized instruction key: %v", key)
	}

	return nil
}

func (s *guidedShip) Position() grid.Point {
	return s.position
}

func (s *guidedShip) Waypoint() grid.Point {
	return s.waypoint
}

func findQuadrant(point grid.Point) int {
	switch {
	case point.I <= 0 && 0 <= point.J:
		return 0
	case point.I <= 0 && point.J < 0:
		return 1
	case 0 < point.I && point.J < 0:
		return 2
	default:
		return 3
	}
}
