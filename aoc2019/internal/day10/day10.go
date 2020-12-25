package day10

import (
	"aoc2019/internal/location"
	"fmt"
	"math"
	"reflect"
	"sort"
)

const (
	kQUAD_I   = 0
	kQUAD_II  = 1
	kQUAD_III = 2
	kQUAD_IV  = 3
)

type asteroidSet map[location.Point]struct{}

type relativeAsteroid struct {
	Loc  location.Point
	Dist int
}

type relativeAsteroidList []relativeAsteroid

func (ral relativeAsteroidList) Len() int           { return len(ral) }
func (ral relativeAsteroidList) Swap(i, j int)      { ral[i], ral[j] = ral[j], ral[i] }
func (ral relativeAsteroidList) Less(i, j int) bool { return ral[i].Dist < ral[j].Dist }

type relativeAsteroidMapping map[location.Vector]relativeAsteroidList

type vectorAndAngle struct {
	V location.Vector
	A float64
}

type vectorList []location.Vector

func (v vectorList) Len() int      { return len(v) }
func (v vectorList) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v vectorList) Less(i, j int) bool {
	a := v[i]
	b := v[j]

	if quadA, quadB := vecToQuadrant(a), vecToQuadrant(b); quadA != quadB {
		return quadA < quadB
	}

	return getAngle(a) < getAngle(b)
}

func vecToQuadrant(v location.Vector) int {
	if v.X >= 0 && v.Y <= 0 {
		return kQUAD_I
	} else if v.X >= 0 {
		return kQUAD_II
	} else if v.X < 0 && v.Y >= 0 {
		return kQUAD_III
	} else {
		return kQUAD_IV
	}
}

func getAngle(v location.Vector) float64 {
	return math.Atan2(float64(v.Y), float64(v.X))
}

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	as, err := inputToAsteroidSet(input)
	if err != nil {
		return nil, err
	}

	_, maxViewableAsteroids := findAsteroidWithBestView(as)

	return maxViewableAsteroids, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	as, err := inputToAsteroidSet(input)
	if err != nil {
		return nil, err
	}

	bestLoc, _ := findAsteroidWithBestView(as)

	ram := groupAsteroidsByReducedVector(bestLoc, as)

	orderOfDestruction := findOrderOfDestruction(ram)
	if len(orderOfDestruction) < 200 {
		return nil, fmt.Errorf("Expected 200 or more asteroids, not %d", len(orderOfDestruction))
	}

	twoHundredth := orderOfDestruction[199]

	return (100 * twoHundredth.X) + twoHundredth.Y, nil
}

func inputToAsteroidSet(input string) (asteroidSet, error) {
	as := asteroidSet{}
	x, y := 0, 0
	for _, r := range input {
		switch r {
		case '.':
			x += 1
		case '#':
			asteroid := location.Point{X: x, Y: y}
			as[asteroid] = struct{}{}
			x += 1
		case '\n':
			x = 0
			y += 1
		default:
			return nil, fmt.Errorf("Unrecognized rune in input: %v", r)
		}
	}
	return as, nil
}

func findAsteroidWithBestView(as asteroidSet) (location.Point, int) {
	var bestLoc location.Point
	mostVisibleAsteroids := -1

	for asteroid := range as {
		linesToOtherAsteroids := map[location.Vector]struct{}{}
		for otherAsteroid := range as {
			if reflect.DeepEqual(asteroid, otherAsteroid) {
				continue
			}

			v := location.NewVector(
				asteroid.X-otherAsteroid.X,
				asteroid.Y-otherAsteroid.Y)
			v.Reduce()
			linesToOtherAsteroids[*v] = struct{}{}
		}
		if len(linesToOtherAsteroids) > mostVisibleAsteroids {
			bestLoc = asteroid
			mostVisibleAsteroids = len(linesToOtherAsteroids)
		}
	}

	return bestLoc, mostVisibleAsteroids
}

func groupAsteroidsByReducedVector(origin location.Point, as asteroidSet) relativeAsteroidMapping {
	ram := relativeAsteroidMapping{}
	for asteroid := range as {
		if reflect.DeepEqual(origin, asteroid) {
			continue
		}

		v := location.NewVector(
			asteroid.X-origin.X,
			asteroid.Y-origin.Y)
		v.Reduce()
		dist := location.ManhattanDistance(origin, asteroid)
		relAsteroid := relativeAsteroid{
			Loc:  asteroid,
			Dist: dist,
		}

		asteroids, ok := ram[*v]
		if !ok {
			asteroids = relativeAsteroidList{relAsteroid}
		} else {
			asteroids = append(asteroids, relAsteroid)
			sort.Sort(asteroids)
		}
		ram[*v] = asteroids
	}
	return ram
}

func getVectorList(ram relativeAsteroidMapping) vectorList {
	vectors := make(vectorList, len(ram))
	i := 0
	for v := range ram {
		vectors[i] = v
		i++
	}
	return vectors
}

func findOrderOfDestruction(ram relativeAsteroidMapping) []location.Point {
	vectors := getVectorList(ram)
	sort.Sort(vectors)

	nAsteroids := 0
	for _, v := range ram {
		nAsteroids += len(v)
	}

	asteroidsInOrderOfDestruction := make([]location.Point, nAsteroids)
	i := 0
	depth := 0
	for i < nAsteroids {
		for _, v := range vectors {
			asteroidsOnLine := ram[v]
			if depth < len(asteroidsOnLine) {
				asteroidsInOrderOfDestruction[i] = asteroidsOnLine[depth].Loc
				i++
			}
		}
		depth++
	}
	return asteroidsInOrderOfDestruction
}
