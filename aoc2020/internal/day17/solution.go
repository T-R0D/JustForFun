// Possible Improvement: Try using big arrays instead of maps and just store the
//                       whole space. I'm skeptical this provides huge improvement,
//                       but might be fun to try.

package day17

import (
	"strconv"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	activeCubes, err := parseInitialConfiguration(input)
	if err != nil {
		return "", err
	}

	simulator := newEnergySourceSimulator(activeCubes)

	for i := 0; i < numBootCycles; i++ {
		simulator.Cycle()
	}

	return strconv.Itoa(simulator.NActiveCubes()), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	activeCubes, err := parseInitialConfiguration(input)
	if err != nil {
		return "", err
	}

	simulator := newEnergySourceSimulator4D(activeCubes)

	for i := 0; i < numBootCycles; i++ {
		simulator.Cycle()
	}

	return strconv.Itoa(simulator.NActiveCubes()), nil
}

const (
	activeCubeSymbol   = '#'
	inactiveCubeSymbol = '.'
)

const (
	numBootCycles = 6
)

func parseInitialConfiguration(startingConfiguration string) (map[grid.Point3]bool, error) {
	activeCubes := map[grid.Point3]bool{}

	i, j := 0, 0
	for _, r := range startingConfiguration {
		switch r {
		case activeCubeSymbol:
			activeCubes[grid.Point3{X: j, Y: i, Z: 0}] = true
			j++
		case inactiveCubeSymbol:
			j++
		case '\n':
			i++
			j = 0
		default:
			return nil, errors.Errorf("unrecognized symbol %c at %d %d", r, i, j)
		}
	}

	return activeCubes, nil
}

type energySourceSimulator struct {
	activeCubes       map[grid.Point3]bool
	inactiveNeighbors map[grid.Point3]bool
}

func newEnergySourceSimulator(activeCubes map[grid.Point3]bool) *energySourceSimulator {
	activeCubesCopy := map[grid.Point3]bool{}
	for location := range activeCubes {
		activeCubesCopy[location] = true
	}

	inactiveNeighbors := findInactiveNeighbors(activeCubes)

	return &energySourceSimulator{
		activeCubes:       activeCubes,
		inactiveNeighbors: inactiveNeighbors,
	}
}

func findInactiveNeighbors(activeCubes map[grid.Point3]bool) map[grid.Point3]bool {
	inactiveNeighbors := map[grid.Point3]bool{}
	for activeLocation := range activeCubes {
		for x := activeLocation.X - 1; x <= activeLocation.X+1; x++ {
			for y := activeLocation.Y - 1; y <= activeLocation.Y+1; y++ {
				for z := activeLocation.Z - 1; z <= activeLocation.Z+1; z++ {
					location := grid.Point3{X: x, Y: y, Z: z}
					if location == activeLocation {
						continue
					} else if _, ok := activeCubes[location]; !ok {
						inactiveNeighbors[location] = false
					}
				}
			}
		}
	}
	return inactiveNeighbors
}

func (s *energySourceSimulator) Cycle() {
	newActiveCubes := map[grid.Point3]bool{}

	for activeCube := range s.activeCubes {
		neighbors := generateNeighbors(activeCube)
		numActiveNeighbors := 0
		for _, neighbor := range neighbors {
			if _, ok := s.activeCubes[neighbor]; ok {
				numActiveNeighbors++
			}
		}

		if numActiveNeighbors == 2 || numActiveNeighbors == 3 {
			newActiveCubes[activeCube] = true
		}
	}

	for inactiveNeighbor := range s.inactiveNeighbors {
		neighbors := generateNeighbors(inactiveNeighbor)
		numActiveNeighbors := 0
		for _, neighbor := range neighbors {
			if _, ok := s.activeCubes[neighbor]; ok {
				numActiveNeighbors++
			}
		}

		if numActiveNeighbors == 3 {
			newActiveCubes[inactiveNeighbor] = true
		}
	}

	s.activeCubes = newActiveCubes
	s.inactiveNeighbors = findInactiveNeighbors(newActiveCubes)
}

func (s *energySourceSimulator) NActiveCubes() int {
	return len(s.activeCubes)
}

func generateNeighbors(targetLocation grid.Point3) []grid.Point3 {
	neighbors := make([]grid.Point3, 0, 26)
	for x := targetLocation.X - 1; x <= targetLocation.X+1; x++ {
		for y := targetLocation.Y - 1; y <= targetLocation.Y+1; y++ {
			for z := targetLocation.Z - 1; z <= targetLocation.Z+1; z++ {
				location := grid.Point3{X: x, Y: y, Z: z}
				if location == targetLocation {
					continue
				}
				neighbors = append(neighbors, location)
			}
		}
	}
	return neighbors
}

type energySourceSimulator4D struct {
	activeCubes       map[grid.Point4]bool
}

func newEnergySourceSimulator4D(activeCubes map[grid.Point3]bool) *energySourceSimulator4D {
	activeCubesCopy := map[grid.Point4]bool{}
	for location := range activeCubes {
		location4D := grid.Point4{W: 0, X: location.X, Y: location.Y, Z: location.Z}
		activeCubesCopy[location4D] = true
	}


	return &energySourceSimulator4D{
		activeCubes:       activeCubesCopy,
	}
}

func (s *energySourceSimulator4D) Cycle() {
	newActiveCubes := map[grid.Point4]bool{}
	inactiveNeighbors := map[grid.Point4]bool{}


	for activeCube := range s.activeCubes {
		neighbors := generateNeighbors4D(activeCube)
		numActiveNeighbors := 0
		for _, neighbor := range neighbors {
			if _, ok := s.activeCubes[neighbor]; ok {
				numActiveNeighbors++
			} else {
				inactiveNeighbors[neighbor] = false
			}
		}

		if numActiveNeighbors == 2 || numActiveNeighbors == 3 {
			newActiveCubes[activeCube] = true
		}
	}

	for inactiveNeighbor := range inactiveNeighbors {
		neighbors := generateNeighbors4D(inactiveNeighbor)
		numActiveNeighbors := 0
		for _, neighbor := range neighbors {
			if _, ok := s.activeCubes[neighbor]; ok {
				numActiveNeighbors++
			}
		}

		if numActiveNeighbors == 3 {
			newActiveCubes[inactiveNeighbor] = true
		}
	}

	s.activeCubes = newActiveCubes
}

func (s *energySourceSimulator4D) NActiveCubes() int {
	return len(s.activeCubes)
}

func findInactiveNeighbors4D(activeCubes map[grid.Point4]bool) map[grid.Point4]bool {
	inactiveNeighbors := map[grid.Point4]bool{}

	for active := range activeCubes {
		neighbors := generateNeighbors4D(active)
		for _, neighbor := range neighbors {
			if _, ok := activeCubes[neighbor]; !ok {
				inactiveNeighbors[neighbor] = false
			}
		}
	}

	return inactiveNeighbors
}

func generateNeighbors4D(targetLocation grid.Point4) []grid.Point4 {
	neighbors := make([]grid.Point4, 0, 80)
	for w := targetLocation.W - 1; w <= targetLocation.W+1; w++ {
		for x := targetLocation.X - 1; x <= targetLocation.X+1; x++ {
			for y := targetLocation.Y - 1; y <= targetLocation.Y+1; y++ {
				for z := targetLocation.Z - 1; z <= targetLocation.Z+1; z++ {
					location := grid.Point4{W: w, X: x, Y: y, Z: z}
					if location == targetLocation {
						continue
					}
					neighbors = append(neighbors, location)
				}
			}
		}
	}
	return neighbors
}
