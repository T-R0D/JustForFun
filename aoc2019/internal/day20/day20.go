package day20

import (
	"aoc2019/internal/location"
	"container/list"
	"fmt"
	"io/ioutil"
	"strings"
)

type Solver struct {
	InputPath string
}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	input, err := s.readInput()
	if err != nil {
		return nil, err
	}
	w := worldFromInput(input)
	path := findShortestPathThroughMaze(w)
	return len(path) - 1, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	input, err := s.readInput()
	if err != nil {
		return nil, err
	}
	w := worldFromInput(input)
	path := findShortestPathThroughRecursiveMaze(w)

	// 4500 too low.

	return len(path) - 1, nil
}

func (s *Solver) readInput() (string, error) {
	bytes, err := ioutil.ReadFile(s.InputPath)
	if err != nil {
		return "", err
	}

	input := string(bytes)
	return input, nil
}

const (
	EMPTY = int(' ')
	PATH  = int('.')
	WALL  = int('#')
)

type world struct {
	entrance    location.Point
	exit        location.Point
	m           map[location.Point]int
	h           int
	portalPairs map[string]portalPair
	portals     map[location.Point]location.Point
	w           int
}

type portalPair struct {
	a location.Point
	b location.Point
}

func worldFromInput(input string) *world {
	m := make(map[location.Point]int)
	width := 0
	lines := strings.Split(input, "\n")
	height := len(lines)
	for i, line := range lines {
		width = len(line)
		for j, c := range line {
			obj := int(c)
			m[location.Point{X: j, Y: i}] = obj
		}
	}

	var mazeEntrance location.Point
	var mazeExit location.Point
	entrances := make(map[string]location.Point)
	exits := make(map[string]location.Point)
	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			loc := location.Point{X: j, Y: i}
			if obj, ok := m[loc]; ok && isLetter(obj) {
				if surroundingSpacesContainsPath(loc, m) {
					pathLoc := getAdjacentPath(loc, m)
					obj2 := findLetterOppositeOfPath(loc, m)
					if obj == obj2 && obj == int('A') {
						mazeEntrance = pathLoc
					} else if obj == obj2 && obj == int('Z') {
						mazeExit = pathLoc
					} else {
						key := lettersToMapKey(obj, obj2, loc, pathLoc)
						if _, ok := entrances[key]; ok {
							exits[key] = pathLoc
						} else {
							entrances[key] = pathLoc
						}
					}
				}
			}
		}
	}

	portalPairs := make(map[string]portalPair)
	portals := make(map[location.Point]location.Point)
	for key, loc1 := range entrances {
		if loc2, ok := exits[key]; ok {
			portalPairs[key] = portalPair{
				a: loc1,
				b: loc2,
			}
			portals[loc1] = loc2
			portals[loc2] = loc1
		} else {
			panic(fmt.Sprintf("%s in entrances but not matched in exits", key))
		}
	}

	return &world{
		entrance:    mazeEntrance,
		exit:        mazeExit,
		m:           m,
		h:           height,
		portalPairs: portalPairs,
		portals:     portals,
		w:           width,
	}
}

func isLetter(x int) bool {
	return int('A') <= x && x <= int('Z')
}

func surroundingSpacesContainsPath(loc location.Point, m map[location.Point]int) bool {
	adjacentLoc := location.Point{X: loc.X + 1, Y: loc.Y}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return true
	}

	adjacentLoc = location.Point{X: loc.X - 1, Y: loc.Y}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return true
	}

	adjacentLoc = location.Point{X: loc.X, Y: loc.Y + 1}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return true
	}

	adjacentLoc = location.Point{X: loc.X, Y: loc.Y - 1}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return true
	}

	return false
}

func getAdjacentPath(loc location.Point, m map[location.Point]int) location.Point {
	adjacentLoc := location.Point{X: loc.X + 1, Y: loc.Y}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return adjacentLoc
	}

	adjacentLoc = location.Point{X: loc.X - 1, Y: loc.Y}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return adjacentLoc
	}

	adjacentLoc = location.Point{X: loc.X, Y: loc.Y + 1}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return adjacentLoc
	}

	adjacentLoc = location.Point{X: loc.X, Y: loc.Y - 1}
	if obj, ok := m[adjacentLoc]; ok && obj == PATH {
		return adjacentLoc
	}

	return location.Point{X: -1, Y: -1}
}

func findLetterOppositeOfPath(loc location.Point, m map[location.Point]int) int {
	north := location.Point{X: loc.X, Y: loc.Y - 1}
	south := location.Point{X: loc.X, Y: loc.Y + 1}
	east := location.Point{X: loc.X + 1, Y: loc.Y}
	west := location.Point{X: loc.X - 1, Y: loc.Y}

	if obj, ok := m[north]; ok && obj == PATH {
		if letter, ok := m[south]; ok {
			return letter
		} else {
			return -1
		}
	}

	if obj, ok := m[south]; ok && obj == PATH {
		if letter, ok := m[north]; ok {
			return letter
		} else {
			return -1
		}
	}

	if obj, ok := m[east]; ok && obj == PATH {
		if letter, ok := m[west]; ok {
			return letter
		} else {
			return -1
		}
	}

	if obj, ok := m[west]; ok && obj == PATH {
		if letter, ok := m[east]; ok {
			return letter
		} else {
			return -1
		}
	}

	return -1
}

func lettersToMapKey(l1, l2 int, l1Loc, pathLoc location.Point) string {

	diff := location.Point{X: l1Loc.X - pathLoc.X, Y: l1Loc.Y - pathLoc.Y}

	// Path is below or to the right:
	if (diff.X == 0 && diff.Y == -1) || (diff.X == -1 && diff.Y == 0) {
		return fmt.Sprintf("%c%c", rune(l2), rune(l1))
	} else if (diff.X == 0 && diff.Y == 1) || (diff.X == 1 && diff.Y == 0) {
		return fmt.Sprintf("%c%c", rune(l1), rune(l2))
	}

	panic("Path is not adjacent to l1")
}

type searchState struct {
	loc  location.Point
	path []location.Point
}

func newSearchState(newLoc location.Point, previousPath []location.Point) searchState {
	path := append(previousPath[:0:0], previousPath...)
	return searchState{
		loc:  newLoc,
		path: append(path, newLoc),
	}
}

func findShortestPathThroughMaze(w *world) []location.Point {
	q := list.New()
	q.PushBack(newSearchState(w.entrance, []location.Point{}))
	seen := make(map[location.Point]struct{})

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(searchState)

		if _, ok := seen[current.loc]; ok {
			continue
		}

		if current.loc == w.exit {
			return current.path
		}

		if nextLoc, ok := w.portals[current.loc]; ok {
			q.PushBack(newSearchState(nextLoc, current.path))
		}
		adjacentLoc := location.Point{X: current.loc.X + 1, Y: current.loc.Y}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newSearchState(adjacentLoc, current.path))
		}

		adjacentLoc = location.Point{X: current.loc.X - 1, Y: current.loc.Y}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newSearchState(adjacentLoc, current.path))
		}

		adjacentLoc = location.Point{X: current.loc.X, Y: current.loc.Y + 1}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newSearchState(adjacentLoc, current.path))
		}

		adjacentLoc = location.Point{X: current.loc.X, Y: current.loc.Y - 1}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newSearchState(adjacentLoc, current.path))
		}

		seen[current.loc] = struct{}{}
	}

	panic("Exit to maze not found!")
}

type mazeLocation struct {
	loc   location.Point
	level int
}

func (m mazeLocation) String() string {
	return fmt.Sprintf("X: % 3d  Y: % 3d  Z: %+ 3d", m.loc.X, m.loc.Y, m.level)
}

type recursiveSearchState struct {
	mazeLoc mazeLocation
	path    []mazeLocation
}

func newRecursiveSearchState(newLoc location.Point, newLevel int, previousPath []mazeLocation) recursiveSearchState {
	newMazeLoc := mazeLocation{
		loc:   newLoc,
		level: newLevel,
	}
	path := append(previousPath[:0:0], previousPath...)
	return recursiveSearchState{
		mazeLoc: newMazeLoc,
		path:    append(path, newMazeLoc),
	}
}

func findShortestPathThroughRecursiveMaze(w *world) []mazeLocation {
	q := list.New()
	q.PushBack(newRecursiveSearchState(w.entrance, 0, []mazeLocation{}))
	seen := make(map[mazeLocation]struct{})

	for q.Len() > 0 {
		f := q.Front()
		q.Remove(f)
		current := f.Value.(recursiveSearchState)

		if _, ok := seen[current.mazeLoc]; ok {
			continue
		}

		if current.mazeLoc.loc == w.exit && current.mazeLoc.level == 0 {
			return current.path
		}

		if nextLoc, ok := w.portals[current.mazeLoc.loc]; ok {
			var nextLevel int
			if isOuterPortal(current.mazeLoc.loc, w) {
				nextLevel = current.mazeLoc.level - 1
			} else {
				nextLevel = current.mazeLoc.level + 1
			}
			if nextLevel >= 0 {
				q.PushBack(newRecursiveSearchState(nextLoc, nextLevel, current.path))
			}
		}

		adjacentLoc := location.Point{X: current.mazeLoc.loc.X + 1, Y: current.mazeLoc.loc.Y}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newRecursiveSearchState(adjacentLoc, current.mazeLoc.level, current.path))
		}

		adjacentLoc = location.Point{X: current.mazeLoc.loc.X - 1, Y: current.mazeLoc.loc.Y}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newRecursiveSearchState(adjacentLoc, current.mazeLoc.level, current.path))
		}

		adjacentLoc = location.Point{X: current.mazeLoc.loc.X, Y: current.mazeLoc.loc.Y + 1}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newRecursiveSearchState(adjacentLoc, current.mazeLoc.level, current.path))
		}

		adjacentLoc = location.Point{X: current.mazeLoc.loc.X, Y: current.mazeLoc.loc.Y - 1}
		if obj, ok := w.m[adjacentLoc]; ok && obj == PATH {
			q.PushBack(newRecursiveSearchState(adjacentLoc, current.mazeLoc.level, current.path))
		}

		seen[current.mazeLoc] = struct{}{}
	}

	panic("Exit to maze not found!")
}

func isOuterPortal(loc location.Point, w *world) bool {
	return (loc.X <= 2 || w.w-3 <= loc.X) || (loc.Y <= 2 || w.h-3 <= loc.Y)
}
