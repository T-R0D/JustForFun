package day12

import (
	"aoc2019/internal/aocmath"
	"fmt"
	"math"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	bodies, err := inputToBodies(input)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 1000; i++ {
		nBodySimulateTimeStep(bodies)
	}

	return computeSystemEnergy(bodies), nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	bodies, err := inputToBodies(input)
	if err != nil {
		return nil, err
	}
	originalBodies, err := inputToBodies(input)
	if err != nil {
		return nil, err
	}

	stepsToRepeat := vector{-1, -1, -1}
	step := 0
	for stepsToRepeat[0] == -1 || stepsToRepeat[1] == -1 || stepsToRepeat[2] == -1 {
		step++
		nBodySimulateTimeStep(bodies)

		if stepsToRepeat[0] == -1 && bodiesAtSameStateForAxis(bodies, originalBodies, 0) {
			stepsToRepeat[0] = step
		}

		if stepsToRepeat[1] == -1 && bodiesAtSameStateForAxis(bodies, originalBodies, 1) {
			stepsToRepeat[1] = step
		}

		if stepsToRepeat[2] == -1 && bodiesAtSameStateForAxis(bodies, originalBodies, 2) {
			stepsToRepeat[2] = step
		}
	}

	return lcm(stepsToRepeat[0], lcm(stepsToRepeat[1], stepsToRepeat[2])), nil
}

const (
	DIM = 3
)

type vector [DIM]int

type body struct {
	P vector
	V vector
}

func inputToBodies(input string) ([]*body, error) {
	positions, err := inputToPositionVectors(input)
	if err != nil {
		return nil, err
	}

	return initBodies(positions), nil
}

func inputToPositionVectors(input string) ([]vector, error) {
	vecStrs := strings.Split(input, "\n")

	vecs := make([]vector, len(vecStrs))
	for i, vecStr := range vecStrs {
		x, y, z := 0, 0, 0
		_, err := fmt.Sscanf(vecStr, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		if err != nil {
			return nil, err
		}
		vecs[i] = vector{x, y, z}
	}
	return vecs, nil
}

func initBodies(positions []vector) []*body {
	bodies := make([]*body, len(positions))
	for i, p := range positions {
		bodies[i] = newBody(p)
	}
	return bodies
}

func newBody(position vector) *body {
	return &body{
		P: position,
		V: vector{0, 0, 0},
	}
}

func nBodySimulateTimeStep(bodies []*body) {
	for i, b1 := range bodies {
		for j := i + 1; j < len(bodies); j++ {
			b2 := bodies[j]
			updateVelocity(b1, b2)
		}
	}

	for _, b := range bodies {
		b.updatePosition()
	}
}

func updateVelocity(b1, b2 *body) {
	for i := 0; i < DIM; i++ {
		if b1.P[i] == b2.P[i] {
			continue
		}

		if b1.P[i] < b2.P[i] {
			b1.V[i] += 1
			b2.V[i] -= 1
		} else {
			b1.V[i] -= 1
			b2.V[i] += 1
		}
	}
}

func (b *body) updatePosition() {
	for i := 0; i < DIM; i++ {
		b.P[i] += b.V[i]
	}
}

func computeSystemEnergy(bodies []*body) int {
	te := 0
	for _, b := range bodies {
		te += b.computeTotalEnergy()
	}
	return te
}

func (b *body) computeTotalEnergy() int {
	pe := 0
	ke := 0
	for i := 0; i < DIM; i++ {
		pe += int(math.Abs(float64(b.P[i])))
		ke += int(math.Abs(float64(b.V[i])))
	}
	return pe * ke
}

func bodiesAtSameStateForAxis(bodies, bodies2 []*body, axis int) bool {
	for i, b := range bodies {
		b2 := bodies2[i]
		if b.P[axis] != b2.P[axis] || b.V[axis] != b2.V[axis] {
			return false
		}
	}
	return true
}

func lcm(a, b int) int {
	return (a * b) / aocmath.GreatestCommonDivisor(a, b)
}
