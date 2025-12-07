package day07

import (
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/counter"
	"github.com/T-R0D/aoc2025/v2/internal/set"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	spec := parseManifoldSpec(input)

	nSplits := spec.CountBeamSplits()

	return strconv.Itoa(nSplits), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	spec := parseManifoldSpec(input)

	nSplits := spec.CountTimelines()

	return strconv.Itoa(nSplits), nil
}

type manifoldSpecification struct {
	depth     int
	source    [2]int
	splitters []set.Set[int]
}

func newManifoldSpecification() *manifoldSpecification {
	return &manifoldSpecification{
		depth:     0,
		source:    [2]int{0, 0},
		splitters: []set.Set[int]{},
	}
}

func parseManifoldSpec(input string) manifoldSpecification {
	spec := newManifoldSpecification()
	for i, line := range strings.Split(input, "\n") {
		spec.splitters = append(spec.splitters, set.New[int]())
		spec.depth += 1

		for j, val := range line {
			if val == 'S' {
				spec.source = [2]int{i, j}
			}

			if val == '^' {
				spec.splitters[i].Add(j)
			}
		}
	}

	return *spec
}

func (this *manifoldSpecification) CountBeamSplits() int {
	nSplits := 0
	currentBeamJs := set.New[int]()
	for i := range this.depth {
		nextBeamJs := set.New[int]()

		if this.source[0] == i {
			nextBeamJs.Add(this.source[1])
		}

	BeamScan:
		for beamJ := range currentBeamJs.All() {
			if !this.splitters[i].Contains(beamJ) {
				nextBeamJs.Add(beamJ)
				continue BeamScan
			}

			nextBeamJs.Add(beamJ - 1)
			nextBeamJs.Add(beamJ + 1)
			nSplits += 1
		}

		currentBeamJs = nextBeamJs
	}

	return nSplits
}

func (this *manifoldSpecification) CountTimelines() int {
	currentTimelinesByPosition := counter.New[int]()
	for i := range this.depth {
		nextTimelinesByPosition := counter.New[int]()

		if this.source[0] == i {
			nextTimelinesByPosition.Increment(this.source[1], 1)
		}

	BeamScan:
		for beamJ, nTimelines := range currentTimelinesByPosition.Items() {
			if !this.splitters[i].Contains(beamJ) {
				nextTimelinesByPosition.Increment(beamJ, nTimelines)
				continue BeamScan
			}

			nextTimelinesByPosition.Increment(beamJ-1, nTimelines)
			nextTimelinesByPosition.Increment(beamJ+1, nTimelines)
		}

		currentTimelinesByPosition = nextTimelinesByPosition
	}

	return currentTimelinesByPosition.Sum()
}
