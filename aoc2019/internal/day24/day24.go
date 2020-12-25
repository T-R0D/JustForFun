package day24

import (
	"aoc2019/internal/location"
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	startingState := inputToMap(input)
	firstRepeatedLayout := simulateBugLife(startingState)
	return firstRepeatedLayout.BiodiversityScore(), nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	startingState := inputToMap(input)
	finalState := simulateBugLifeRecursive(startingState, 200)
	// 1779 too low
	return countBugs(finalState), nil
}

const (
	BUG             = int('#')
	EMPTY           = int('.')
	RECURSIVE_LEVEL = int('?')
)

func inputToMap(input string) layout {
	m := make(map[location.Point]int)
	for i, line := range strings.Split(input, "\n") {
		for j, obj := range line {
			m[location.Point{X: j, Y: i}] = int(obj)
		}
	}
	return layout(m)
}

type layout map[location.Point]int

func (l layout) String() string {
	var b strings.Builder
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b.WriteRune(rune(l[location.Point{X: j, Y: i}]))
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (l layout) generateNextState() layout {
	next := make(map[location.Point]int)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			adjacentBugs := 0
			if obj, ok := l[location.Point{X: j - 1, Y: i}]; ok && obj == BUG {
				adjacentBugs += 1
			}
			if obj, ok := l[location.Point{X: j + 1, Y: i}]; ok && obj == BUG {
				adjacentBugs += 1
			}
			if obj, ok := l[location.Point{X: j, Y: i - 1}]; ok && obj == BUG {
				adjacentBugs += 1
			}
			if obj, ok := l[location.Point{X: j, Y: i + 1}]; ok && obj == BUG {
				adjacentBugs += 1
			}

			currentLoc := location.Point{X: j, Y: i}
			currentObj := l[currentLoc]
			if currentObj == BUG && adjacentBugs != 1 {
				next[currentLoc] = EMPTY
			} else if currentObj == EMPTY && (adjacentBugs == 1 || adjacentBugs == 2) {
				next[currentLoc] = BUG
			} else {
				next[currentLoc] = l[currentLoc]
			}
		}
	}
	return layout(next)
}

func (l layout) BiodiversityScore() int {
	score := 0
	for c := 0; c < 25; c += 1 {
		i, j := c/5, c%5
		currentLoc := location.Point{X: j, Y: i}
		if obj := l[currentLoc]; obj == BUG {
			score += 1 << c
		}
	}
	return score
}

func createEmptyLayout() layout {
	l := layout{}
	for i := 0; i < 5; i += 1 {
		for j := 0; j < 5; j += 1 {
			if i == 2 && j == 2 {
				l[location.Point{X: j, Y: i}] = RECURSIVE_LEVEL
			} else {
				l[location.Point{X: j, Y: i}] = EMPTY
			}
		}
	}
	return l
}

func simulateBugLife(startingState layout) layout {
	seen := make(map[string]struct{})
	seen[startingState.String()] = struct{}{}

	state := startingState
	for {
		state = state.generateNextState()
		if _, ok := seen[state.String()]; ok {
			break
		}
		seen[state.String()] = struct{}{}
	}

	return state
}

func simulateBugLifeRecursive(startingState layout, iterations int) map[int]layout {
	startingState[location.Point{X: 2, Y: 2}] = RECURSIVE_LEVEL
	state := map[int]layout{0: startingState}

	for t := 0; t < iterations; t += 1 {
		nextState := make(map[int]layout)

		minZ, maxZ := 0, 0
		for z := range state {
			if z < minZ {
				minZ = z
			}
			if z > maxZ {
				maxZ = z
			}
		}

		state[maxZ+1] = createEmptyLayout()
		state[minZ-1] = createEmptyLayout()

		for z, layer := range state {
			newLayout := layout{}
			outer, outerOk := state[z-1]
			inner, innerOk := state[z+1]

			for i := 0; i < 5; i += 1 {
				for j := 0; j < 5; j += 1 {
					if i == 2 && j == 2 {
						newLayout[location.Point{X: 2, Y: 2}] = RECURSIVE_LEVEL
						continue
					}

					currentLoc := location.Point{X: j, Y: i}
					adjacentBugs := 0

					// Check the standard cardinal directions.
					adjacentLoc := location.Point{X: currentLoc.X - 1, Y: currentLoc.Y}
					if obj, ok := layer[adjacentLoc]; (ok && adjacentLoc != location.Point{X: 2, Y: 2}) && obj == BUG {
						adjacentBugs += 1
					}
					adjacentLoc = location.Point{X: currentLoc.X + 1, Y: currentLoc.Y}
					if obj, ok := layer[adjacentLoc]; (ok && adjacentLoc != location.Point{X: 2, Y: 2}) && obj == BUG {
						adjacentBugs += 1
					}
					adjacentLoc = location.Point{X: currentLoc.X, Y: currentLoc.Y - 1}
					if obj, ok := layer[adjacentLoc]; (ok && adjacentLoc != location.Point{X: 2, Y: 2}) && obj == BUG {
						adjacentBugs += 1
					}
					adjacentLoc = location.Point{X: currentLoc.X, Y: currentLoc.Y + 1}
					if obj, ok := layer[adjacentLoc]; (ok && adjacentLoc != location.Point{X: 2, Y: 2}) && obj == BUG {
						adjacentBugs += 1
					}

					// Check the inner level, if relevant:
					if (i == 1 && j == 2) && innerOk {
						for j2 := 0; j2 < 5; j2++ {
							if inner[location.Point{X: j2, Y: 0}] == BUG {
								adjacentBugs += 1
							}
						}
					} else if (i == 3 && j == 2) && innerOk {
						for j2 := 0; j2 < 5; j2++ {
							if inner[location.Point{X: j2, Y: 4}] == BUG {
								adjacentBugs += 1
							}
						}
					} else if (j == 1 && i == 2) && innerOk {
						for i2 := 0; i2 < 5; i2++ {
							if inner[location.Point{X: 0, Y: i2}] == BUG {
								adjacentBugs += 1
							}
						}
					} else if (j == 3 && i == 2) && innerOk {
						for i2 := 0; i2 < 5; i2++ {
							if inner[location.Point{X: 4, Y: i2}] == BUG {
								adjacentBugs += 1
							}
						}
					}

					// Check the outer level, if relevant:
					if i == 0 && outerOk && outer[location.Point{X: 2, Y: 1}] == BUG {
						adjacentBugs += 1
					} else if i == 4 && outerOk && outer[location.Point{X: 2, Y: 3}] == BUG {
						adjacentBugs += 1
					}

					if j == 0 && outerOk && outer[location.Point{X: 1, Y: 2}] == BUG {
						adjacentBugs += 1
					} else if j == 4 && outerOk && outer[location.Point{X: 3, Y: 2}] == BUG {
						adjacentBugs += 1
					}

					// Assign the value to the same location in the new grid.
					if layer[currentLoc] == BUG && adjacentBugs != 1 {
						newLayout[currentLoc] = EMPTY
					} else if layer[currentLoc] == EMPTY && (adjacentBugs == 1 || adjacentBugs == 2) {
						newLayout[currentLoc] = BUG
					} else {
						objRune := rune(layer[currentLoc])
						if false {
							fmt.Println(objRune)
						}
						newLayout[currentLoc] = layer[currentLoc]
					}
				}
			}
			nextState[z] = newLayout
		}

		state = nextState
	}
	return state
}

func countBugs(state map[int]layout) int {
	bugs := 0
	for _, layer := range state {
		for i := 0; i < 5; i += 1 {
			for j := 0; j < 5; j += 1 {
				if layer[location.Point{X: j, Y: i}] == BUG {
					bugs += 1
				}
			}
		}
	}
	return bugs
}
