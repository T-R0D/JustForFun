package day15

import (
	"aoc2019/internal/intcode"
	"aoc2019/internal/location"
	"container/list"
	"fmt"
	"math"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	shortestPath := findShortestPathToOxygenSupply(input)
	// -1 to not count the starting space.
	return len(shortestPath) - 1, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	worldMap, oxygenSource := mapTheWorld(input)

	return computeOxygenFillTime(worldMap, oxygenSource), nil
}

const (
	DIR_NORTH = 1
	DIR_SOUTH = 2
	DIR_WEST  = 3
	DIR_EAST  = 4

	MAP_WALL   = 0
	MAP_EMPTY  = 1
	MAP_OXYGEN = 2
)

type droid struct {
	Computer *intcode.Computer
	Loc      location.Point
}

func newDroid(program string) *droid {
	c := intcode.NewComputer()
	c.SetInterruptibleMode()
	c.InputProgram(program)
	return &droid{
		Computer: c,
		Loc:      location.Point{X: 0, Y: 0},
	}
}

func (d *droid) Clone() *droid {
	return &droid{
		Computer: d.Computer.Clone(),
		Loc:      d.Loc,
	}
}

type searchState struct {
	d     *droid
	space int
	path  []location.Point
}

func findShortestPathToOxygenSupply(droidProgram string) []location.Point {
	q := list.New()
	d := newDroid(droidProgram)
	state, err := d.Computer.RunProgram()
	if err != nil {
		panic(err)
	} else if state != intcode.STATE_AWAITING_INPUT {
		panic("Droid should be awaiting input.")
	}
	q.PushBack(&searchState{
		d:     d,
		path:  []location.Point{location.Point{X: 0, Y: 0}},
		space: MAP_EMPTY,
	})

	worldMap := map[location.Point]int{}

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(*searchState)

		if current.space == MAP_OXYGEN {
			return current.path
		}

		if _, ok := worldMap[current.d.Loc]; ok {
			continue
		}

		q.PushBack(produceNewSearchState(worldMap, current, DIR_NORTH))
		q.PushBack(produceNewSearchState(worldMap, current, DIR_SOUTH))
		q.PushBack(produceNewSearchState(worldMap, current, DIR_WEST))
		q.PushBack(produceNewSearchState(worldMap, current, DIR_EAST))

		worldMap[current.d.Loc] = current.space
	}

	return nil
}

func produceNewSearchState(worldMap map[location.Point]int, current *searchState, direction int) *searchState {
	newDroid := current.d.Clone()
	newPath := append(current.path[:0:0], current.path...)
	var space int

	if newDroid.Computer.GetState() != intcode.STATE_AWAITING_INPUT {
		panic(fmt.Errorf("Expected droid to be awaiting input. state=%d", newDroid.Computer.GetState()))
	}
	newDroid.Computer.Input(direction)
	state, err := newDroid.Computer.RunProgram()
	if err != nil {
		panic(err)
	} else if state != intcode.STATE_AWAITING_OUTPUT {
		panic("Expected droid to have output to give.")
	}
	space = newDroid.Computer.CollectOutput()
	newDroid.Computer.RunProgram()

	var newLoc location.Point
	switch direction {
	case DIR_NORTH:
		newLoc = location.Point{
			X: newDroid.Loc.X,
			Y: newDroid.Loc.Y + 1,
		}
	case DIR_SOUTH:
		newLoc = location.Point{
			X: newDroid.Loc.X,
			Y: newDroid.Loc.Y - 1,
		}
	case DIR_WEST:
		newLoc = location.Point{
			X: newDroid.Loc.X - 1,
			Y: newDroid.Loc.Y,
		}
	case DIR_EAST:
		newLoc = location.Point{
			X: newDroid.Loc.X + 1,
			Y: newDroid.Loc.Y,
		}
	}

	if space == MAP_WALL {
		worldMap[newLoc] = MAP_WALL
	} else if space == MAP_EMPTY || space == MAP_OXYGEN {
		newDroid.Loc = newLoc
		newPath = append(newPath, newLoc)
	} else {
		panic("Unknown value occupying space.")
	}

	return &searchState{
		d:     newDroid,
		space: space,
		path:  newPath,
	}
}

func mapTheWorld(droidProgram string) (map[location.Point]int, location.Point) {
	q := list.New()
	d := newDroid(droidProgram)
	state, err := d.Computer.RunProgram()
	if err != nil {
		panic(err)
	} else if state != intcode.STATE_AWAITING_INPUT {
		panic("Droid should be awaiting input.")
	}
	q.PushBack(&searchState{
		d:     d,
		path:  []location.Point{location.Point{X: 0, Y: 0}},
		space: MAP_EMPTY,
	})

	worldMap := map[location.Point]int{}
	var oxygenSource location.Point

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(*searchState)

		if current.space == MAP_OXYGEN {
			oxygenSource = current.d.Loc
		}

		if _, ok := worldMap[current.d.Loc]; ok {
			continue
		}

		q.PushBack(produceNewSearchState(worldMap, current, DIR_NORTH))
		q.PushBack(produceNewSearchState(worldMap, current, DIR_SOUTH))
		q.PushBack(produceNewSearchState(worldMap, current, DIR_WEST))
		q.PushBack(produceNewSearchState(worldMap, current, DIR_EAST))

		worldMap[current.d.Loc] = current.space
	}

	return worldMap, oxygenSource
}

type oxygenSpreadState struct {
	T   int
	Loc location.Point
}

func computeOxygenFillTime(worldMap map[location.Point]int, oxygenSource location.Point) int {
	q := list.New()
	q.PushBack(&oxygenSpreadState{
		T:   0,
		Loc: oxygenSource,
	})

	visitedAreas := map[location.Point]struct{}{}
	lastTimeStep := 0

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(*oxygenSpreadState)

		if _, ok := visitedAreas[current.Loc]; ok {
			continue
		}

		if space, ok := worldMap[current.Loc]; ok && space == MAP_WALL {
			continue
		}

		q.PushBack(&oxygenSpreadState{
			Loc: location.Point{X: current.Loc.X + 1, Y: current.Loc.Y},
			T:   current.T + 1,
		})
		q.PushBack(&oxygenSpreadState{
			Loc: location.Point{X: current.Loc.X - 1, Y: current.Loc.Y},
			T:   current.T + 1,
		})
		q.PushBack(&oxygenSpreadState{
			Loc: location.Point{X: current.Loc.X, Y: current.Loc.Y + 1},
			T:   current.T + 1,
		})
		q.PushBack(&oxygenSpreadState{
			Loc: location.Point{X: current.Loc.X, Y: current.Loc.Y - 1},
			T:   current.T + 1,
		})

		visitedAreas[current.Loc] = struct{}{}
		lastTimeStep = int(math.Max(float64(lastTimeStep), float64(current.T)))
	}

	return lastTimeStep
}
