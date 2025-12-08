package day08

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/set"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	return this.solvePartOneInternal(input, 1000)
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	junctionBoxCoordinates, err := parseJunctionBoxCoordinates(input)
	if err != nil {
		return "", err
	}

	pairings := getParingsOrderedByDistanceAscending(junctionBoxCoordinates)

	lastPairing, err := findLastPairingToMakeCompleteCircuit(pairings)
	if err != nil {
		return "", err
	}

	distanceFromWall := junctionBoxCoordinates[lastPairing.I][0] * junctionBoxCoordinates[lastPairing.J][0]

	return strconv.Itoa(distanceFromWall), nil
}

func (this *Solver) solvePartOneInternal(input string, connections int) (string, error) {
	junctionBoxCoordinates, err := parseJunctionBoxCoordinates(input)
	if err != nil {
		return "", err
	}

	pairings := getParingsOrderedByDistanceAscending(junctionBoxCoordinates)
	if len(pairings) < connections {
		return "", fmt.Errorf("unable to make %d connections; there are only %d pairings", connections, len(pairings))
	}

	circuits, err := findCircuitsAfterNConnections(pairings, connections)
	if err != nil {
		return "", err
	}

	slices.SortFunc(circuits, func(a set.Set[int], b set.Set[int]) int {
		return b.Len() - a.Len()
	})

	top3SizeProduct := 1
	for _, circuit := range circuits[:3] {
		top3SizeProduct *= circuit.Len()
	}

	return strconv.Itoa(top3SizeProduct), nil
}

type point3 [3]int

func parseJunctionBoxCoordinates(input string) ([]point3, error) {
	lines := strings.Split(input, "\n")
	coordinateList := make([]point3, 0, len(lines))
	for _, line := range lines {
		coordinateStrs := strings.Split(line, ",")
		if len(coordinateStrs) != 3 {
			return []point3{}, fmt.Errorf("line did not have 3 coordinates; had %d", len(coordinateStrs))
		}

		coordinate := point3{}
		for i := range coordinate {
			value, err := strconv.Atoi(coordinateStrs[i])
			if err != nil {
				return []point3{}, err
			}

			coordinate[i] = value
		}

		coordinateList = append(coordinateList, coordinate)
	}

	return coordinateList, nil
}

type coordinatePairing struct {
	I              int
	J              int
	Distance       int
	CoordinateList []point3
}

func getParingsOrderedByDistanceAscending(coordinates []point3) []coordinatePairing {
	pairings := []coordinatePairing{}
	for i := 0; i < len(coordinates)-1; i += 1 {
		for j := i + 1; j < len(coordinates); j += 1 {
			pairings = append(pairings, coordinatePairing{
				I:              i,
				J:              j,
				Distance:       squaredEuclideanDistance(coordinates[i], coordinates[j]),
				CoordinateList: coordinates,
			})
		}
	}

	slices.SortFunc(pairings, func(a coordinatePairing, b coordinatePairing) int {
		return a.Distance - b.Distance
	})

	return pairings
}

func squaredEuclideanDistance(a point3, b point3) int {
	sum := 0
	for i := range 3 {
		difference := a[i] - b[i]
		sum += (difference * difference)
	}

	return sum
}

func findCircuitsAfterNConnections(pairings []coordinatePairing, nConnections int) ([]set.Set[int], error) {
	if len(pairings) == 0 {
		return []set.Set[int]{}, nil
	}

	circuits := make([]set.Set[int], 0, len(pairings[0].CoordinateList))
	for i := range pairings[0].CoordinateList {
		circuit := set.New[int]()
		circuit.Add(i)
		circuits = append(circuits, circuit)
	}

	for _, pairing := range pairings[:nConnections] {
		i, j := pairing.I, pairing.J

		iCircuitIndex := slices.IndexFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(i)
		})
		jCircuitIndex := slices.IndexFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(j)
		})

		if iCircuitIndex == -1 || jCircuitIndex == -1 {
			return []set.Set[int]{}, fmt.Errorf("somehow either i (%d) or j (%d) were not in the circuits", i, j)
		}

		if iCircuitIndex == jCircuitIndex {
			continue
		}

		newCircuit := set.Union(circuits[iCircuitIndex], circuits[jCircuitIndex])
		circuits = slices.DeleteFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(i)
		})
		circuits = slices.DeleteFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(j)
		})
		circuits = append(circuits, newCircuit)
	}

	return circuits, nil
}

func findLastPairingToMakeCompleteCircuit(pairings []coordinatePairing) (coordinatePairing, error) {
	if len(pairings) == 0 {
		return coordinatePairing{}, fmt.Errorf("there were no pairings to connect")
	}

	circuits := make([]set.Set[int], 0, len(pairings[0].CoordinateList))
	for i := range pairings[0].CoordinateList {
		circuit := set.New[int]()
		circuit.Add(i)
		circuits = append(circuits, circuit)
	}

	for _, pairing := range pairings {
		i, j := pairing.I, pairing.J

		iCircuitIndex := slices.IndexFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(i)
		})
		jCircuitIndex := slices.IndexFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(j)
		})

		if iCircuitIndex == -1 || jCircuitIndex == -1 {
			return coordinatePairing{}, fmt.Errorf("somehow either i (%d) or j (%d) were not in the circuits", i, j)
		}

		if iCircuitIndex == jCircuitIndex {
			continue
		}

		newCircuit := set.Union(circuits[iCircuitIndex], circuits[jCircuitIndex])
		circuits = slices.DeleteFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(i)
		})
		circuits = slices.DeleteFunc(circuits, func(circuit set.Set[int]) bool {
			return circuit.Contains(j)
		})
		circuits = append(circuits, newCircuit)

		if len(circuits) == 1 {
			return pairing, nil
		}
	}

	return coordinatePairing{}, fmt.Errorf("somehow, a complete circuit could not be made with all pairings")
}
