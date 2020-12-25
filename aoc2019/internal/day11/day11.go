package day11

import (
	"aoc2019/internal/intcode"
	"aoc2019/internal/location"
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	return countCoveredAreas(input)
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	return paintShipHull(input)
}

func countCoveredAreas(prog string) (int, error) {
	r := newHullPaintingRobot()
	r.InputProgram(prog)
	err := r.RunProgram()
	if err != nil {
		return 0, err
	}
	coveredAreas := r.GetCoveredAreas()
	return len(coveredAreas), nil
}

func paintShipHull(prog string) (string, error) {
	r := newHullPaintingRobot()
	r.InputProgram(prog)
	err := r.RunProgram()
	if err != nil {
		return "", err
	}
	coveredAreas := r.GetCoveredAreas()

	image := coveredAreasToImage(coveredAreas)

	return image, nil
}

func coveredAreasToImage(coveredAreas map[location.Point]color) string {
	minX, maxX := 0, 0
	minY, maxY := 0, 0
	for loc := range coveredAreas {
		if loc.X < minX {
			minX = loc.X
		}
		if loc.X > maxX {
			maxX = loc.X
		}
		if loc.Y < minY {
			minY = loc.Y
		}
		if loc.Y > maxY {
			maxY = loc.Y
		}
	}

	var b strings.Builder
	_, err := b.WriteRune('\n')
	if err != nil {
		panic("How did WriteRune return an error?!")
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c, ok := coveredAreas[location.Point{X: x, Y: y}]
			if !ok {
				c = COLOR_BLACK
			}

			var err error
			if c == COLOR_BLACK {
				_, err = b.WriteRune('.')
			} else if c == COLOR_WHITE {
				_, err = b.WriteRune('#')
			}
			if err != nil {
				panic("How did WriteRune return an error?!")
			}
		}
		_, err := b.WriteRune('\n')
		if err != nil {
			panic("How did WriteRune return an error?!")
		}
	}

	return b.String()
}

type color int

const (
	COLOR_BLACK = color(0)
	COLOR_WHITE = color(1)
)

type orientation int

const (
	DIR_LEFT  = 0
	DIR_RIGHT = 1
)

const (
	OR_NORTH = orientation(0)
	OR_EAST  = orientation(1)
	OR_SOUTH = orientation(2)
	OR_WEST  = orientation(3)
)

type hullPaintingRobot struct {
	computer     *intcode.Computer
	coveredAreas map[location.Point]color
	loc          location.Point
	orient       orientation
}

func newHullPaintingRobot() *hullPaintingRobot {
	r := &hullPaintingRobot{
		coveredAreas: map[location.Point]color{
			location.Point{X: 0, Y: 0}: COLOR_WHITE,
		},
		loc:    location.Point{X: 0, Y: 0},
		orient: OR_NORTH,
	}

	c := intcode.NewComputer()
	c.SetInterruptibleMode()
	r.computer = c

	return r
}

func (r *hullPaintingRobot) InputProgram(prog string) {
	r.computer.InputProgram(prog)
}

func (r *hullPaintingRobot) GetCoveredAreas() map[location.Point]color {
	return r.coveredAreas
}

func (r *hullPaintingRobot) RunProgram() error {
	state := r.computer.GetState()
	var err error
	state, err = r.computer.RunProgram()
	if err != nil {
		return err
	} else if state != intcode.STATE_AWAITING_INPUT {
		return fmt.Errorf("Expected program to require input. Instead it is in state %v", state)
	}

	for state == intcode.STATE_AWAITING_INPUT || state == intcode.STATE_AWAITING_OUTPUT {
		if state != intcode.STATE_AWAITING_INPUT {
			return fmt.Errorf("Expected program to require input. Instead it is in state %v", state)
		}
		observedColor, ok := r.coveredAreas[r.loc]
		if !ok {
			observedColor = COLOR_BLACK
		}
		r.computer.Input(int(observedColor))

		state, err = r.computer.RunProgram()
		if err != nil {
			return err
		} else if state != intcode.STATE_AWAITING_OUTPUT {
			return fmt.Errorf("Expected program to have output input. Instead it is in state %v", state)
		}
		colorToPaint := r.computer.CollectOutput()
		r.coveredAreas[r.loc] = color(colorToPaint)

		state, err = r.computer.RunProgram()
		if err != nil {
			return err
		} else if state != intcode.STATE_AWAITING_OUTPUT {
			return fmt.Errorf("Expected program to have output input. Instead it is in state %v", state)
		}
		directionToTurn := r.computer.CollectOutput()
		r.rotateAndAdvance(directionToTurn)

		state, err = r.computer.RunProgram()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *hullPaintingRobot) rotateAndAdvance(directionToTurn int) {
	r.rotate(directionToTurn)
	r.advance()
}

func (r *hullPaintingRobot) rotate(directionToTurn int) {
	if directionToTurn == DIR_LEFT {
		switch r.orient {
		case OR_NORTH:
			r.orient = OR_WEST
		case OR_WEST:
			r.orient = OR_SOUTH
		case OR_SOUTH:
			r.orient = OR_EAST
		case OR_EAST:
			r.orient = OR_NORTH
		}
	} else {
		switch r.orient {
		case OR_NORTH:
			r.orient = OR_EAST
		case OR_EAST:
			r.orient = OR_SOUTH
		case OR_SOUTH:
			r.orient = OR_WEST
		case OR_WEST:
			r.orient = OR_NORTH
		}
	}
}

func (r *hullPaintingRobot) advance() {
	loc := r.loc
	switch r.orient {
	case OR_NORTH:
		r.loc = location.Point{X: loc.X, Y: loc.Y - 1}
	case OR_EAST:
		r.loc = location.Point{X: loc.X + 1, Y: loc.Y}
	case OR_SOUTH:
		r.loc = location.Point{X: loc.X, Y: loc.Y + 1}
	case OR_WEST:
		r.loc = location.Point{X: loc.X - 1, Y: loc.Y}
	}
}
