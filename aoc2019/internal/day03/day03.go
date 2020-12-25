package day03

import (
	"aoc2019/internal/location"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	kUP    = 'U'
	kDOWN  = 'D'
	kLEFT  = 'L'
	kRIGHT = 'R'
)

type pointSet map[location.Point]struct{}

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	ps, err := inputToMaps(input)
	if nil != err {
		return nil, err
	}

	distance, err := findDistanceToIntersectionNearestOrigin(ps)
	if nil != err {
		return nil, err
	}

	return distance, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	ws := strings.Split(input, "\n")
	wireASpec := ws[0]
	wireBSpec := ws[1]

	wireAPoints, err := wireSpecToPointSet(wireASpec)
	if err != nil {
		return nil, err
	}
	wireBPoints, err := wireSpecToPointSet(wireBSpec)
	if err != nil {
		return nil, err
	}

	intersections := findIntersections(wireAPoints, wireBPoints)

	wireAStepsToIntersections, err := findStepsToIntersections(wireASpec, intersections)
	if err != nil {
		return nil, err
	}
	wireBStepsToIntersections, err := findStepsToIntersections(wireBSpec, intersections)
	if err != nil {
		return nil, err
	}

	combinedSteps := -1
	for p, stepsA := range wireAStepsToIntersections {
		stepsB := wireBStepsToIntersections[p]

		if combinedSteps == -1 {
			combinedSteps = stepsA + stepsB
		} else {
			combinedSteps = int(math.Min(float64(combinedSteps), float64(stepsA+stepsB)))
		}
	}

	return combinedSteps, nil
}

func inputToMaps(i string) ([]pointSet, error) {
	wires := strings.Split(i, "\n")
	maps := make([]pointSet, len(wires))

	for i, wire := range wires {
		ps, err := wireSpecToPointSet(wire)
		if nil != err {
			return nil, err
		}
		maps[i] = ps
	}

	return maps, nil
}

func wireSpecToPointSet(spec string) (pointSet, error) {
	ps := make(pointSet)

	err := followWire(spec, func(_ int, point location.Point) {
		ps[point] = struct{}{}
	})
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func followWire(wireSpec string, onStep func(int, location.Point)) error {
	runs := strings.Split(wireSpec, ",")
	currentPos := location.Point{X: 0, Y: 0}
	step := 0
	for _, run := range runs {
		direction, length, err := parseRunSpec(run)
		if err != nil {
			return err
		}

		xInc, yInc, err := directionToIncrements(direction)

		for i := 0; i < length; i++ {
			newPos := location.Point{
				X: currentPos.X + xInc,
				Y: currentPos.Y + yInc,
			}
			step++
			onStep(step, newPos)
			currentPos = newPos
		}
	}

	return nil
}

func parseRunSpec(spec string) (direction byte, length int, err error) {
	direction = spec[0]
	length, err = strconv.Atoi(spec[1:])
	return
}

func directionToIncrements(d byte) (xInc, yInc int, err error) {
	xInc = 0
	yInc = 0
	err = nil
	switch d {
	case kRIGHT:
		xInc = 1
	case kLEFT:
		xInc = -1
	case kUP:
		yInc = 1
	case kDOWN:
		yInc = -1
	default:
		err = fmt.Errorf("Unknown run direction: %q", d)
	}
	return
}

func findDistanceToIntersectionNearestOrigin(wires []pointSet) (distance int, err error) {
	if 2 != len(wires) {
		return 0, fmt.Errorf("Expected 2 wires, not %d", len(wires))
	}

	wireA := wires[0]
	wireB := wires[1]
	origin := location.Point{X: 0, Y: 0}
	distance = 0
	err = fmt.Errorf("No intersections found.")
	for point := range wireA {
		if _, ok := wireB[point]; ok {
			d := location.ManhattanDistance(origin, point)
			if nil != err {
				err = nil
				distance = d
			} else {
				distance = int(math.Min(float64(distance), float64(d)))
			}
		}
	}

	return
}

func findIntersections(wireA, wireB pointSet) pointSet {
	intersections := pointSet{}
	for point := range wireA {
		if _, ok := wireB[point]; ok {
			intersections[point] = struct{}{}
		}
	}
	return intersections
}

func findStepsToIntersections(wireSpec string, intersections pointSet) (map[location.Point]int, error) {
	stepsToIntersections := make(map[location.Point]int)

	onStep := func(steps int, point location.Point) {
		if _, ok := intersections[point]; ok {
			if _, ok = stepsToIntersections[point]; !ok {
				stepsToIntersections[point] = steps
			}
		}
	}
	err := followWire(wireSpec, onStep)
	if err != nil {
		return nil, err
	}

	return stepsToIntersections, nil
}
