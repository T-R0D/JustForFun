package day06

import (
	"container/list"
	"fmt"
	"strings"
)

const (
	MY_SPACECRAFT = "YOU"
	SANTA         = "SAN"
)

type Solver struct{}

type centerToOrbiters map[string][]string

type orbitGraph map[string][]string

type dfsState struct {
	id    string
	depth int
}

type bfsState struct {
	id    string
	depth int
}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	orbitRelations, err := inputToCenterToOrbiters(input)
	if err != nil {
		return nil, err
	}
	return countDirectAndIndirectOrbits(orbitRelations), nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	og, err := inputToOrbitGraph(input)
	if err != nil {
		return nil, err
	}
	return countOrbitalTransfersToMatchOrbit(og, MY_SPACECRAFT, SANTA), nil
}

func inputToCenterToOrbiters(input string) (centerToOrbiters, error) {
	or := centerToOrbiters{}
	orbitSpecs := strings.Split(input, "\n")
	for _, spec := range orbitSpecs {
		parts := strings.Split(spec, ")")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Spec %q not split properly", spec)
		}
		center, orbiter := parts[0], parts[1]

		if orbiters, ok := or[center]; ok {
			orbiters = append(orbiters, orbiter)
			or[center] = orbiters
		} else {
			orbiters = []string{orbiter}
			or[center] = orbiters
		}
	}
	return or, nil
}

func countDirectAndIndirectOrbits(or centerToOrbiters) int {
	stack := []*dfsState{
		&dfsState{
			id:    "COM",
			depth: 0,
		},
	}
	nOrbits := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		orbiters := or[current.id]
		for i := 0; i < len(orbiters); i++ {
			stack = append(stack, &dfsState{
				id:    orbiters[i],
				depth: current.depth + 1,
			})
		}

		nOrbits += current.depth
	}

	return nOrbits
}

func inputToOrbitGraph(input string) (orbitGraph, error) {
	og := orbitGraph{}
	specs := strings.Split(input, "\n")
	for _, spec := range specs {
		parts := strings.Split(spec, ")")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Spec %q not split properly", spec)
		}
		center, orbiter := parts[0], parts[1]

		addRelationToMap(og, center, orbiter)
		addRelationToMap(og, orbiter, center)
	}

	return og, nil
}

func addRelationToMap(m orbitGraph, a, b string) {
	if relations, ok := m[a]; ok {
		relations = append(relations, b)
		m[a] = relations
	} else {
		relations = []string{b}
		m[a] = relations
	}
}

func countOrbitalTransfersToMatchOrbit(og orbitGraph, start, end string) int {
	// NOTE: this does not handle the cases 1) where you start by orbiting Santa
	// and 2) Santa starts orbiting you. This should be simple to handle, but
	// I'm bored with this problem and it isn't required for the solution, so...

	q := list.New()
	q.PushBack(&bfsState{
		id:    MY_SPACECRAFT,
		depth: 0,
	})
	explored := map[string]struct{}{}

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(*bfsState)

		if _, ok := explored[current.id]; ok {
			continue
		}

		if current.id == end {
			// "-2": 1 for the orbit that we are already in, and one for Santa's
			// orbit. The BFS will count 1 depth to find the center that we
			// are orbiting at the beginning and 1 depth to get from the center
			// Santa is orbiting to Santa.
			return current.depth - 2
		}

		explored[current.id] = struct{}{}

		if nextIds, ok := og[current.id]; ok {
			for _, id := range nextIds {
				q.PushBack(&bfsState{
					id:    id,
					depth: current.depth + 1,
				})
			}
		}
	}

	return 0
}
