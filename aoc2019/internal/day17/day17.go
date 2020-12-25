package day17

import (
	"aoc2019/internal/intcode"
	"aoc2019/internal/location"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	sw := newScaffoldWorld()
	sw.snapshotWorld(input)
	intersections := sw.tracePathForIntersections()
	checksum := checksumIntersections(intersections)
	return checksum, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	sw := newScaffoldWorld()
	sw.snapshotWorld(input)
	path := sw.tracePathForPath()
	fmt.Println(path)

	c := intcode.NewComputer()
	c.InputProgram(input)
	c.UpdateProgram(map[int]int{0: 2})
	c.SetInterruptibleMode()
	for c.GetState() != intcode.STATE_AWAITING_INPUT {
		c.RunProgram()
	}

	compressedPath, dictionary := compressPath(path)

	compressedPath2 := strings.Join(compressedPath, ",")

	fmt.Println("inputting main program")
	inputMovementLine(c, compressedPath2) // Main program.
	fmt.Println("defining subroutine A")
	collectOuptutMessage(c)
	inputMovementLine(c, dictionary[SUB_A]) // Definition of subroutine A.
	fmt.Println("defining subroutine B")
	collectOuptutMessage(c)
	inputMovementLine(c, dictionary[SUB_B]) // Definition of subroutine B.
	fmt.Println("defining subroutine C")
	collectOuptutMessage(c)
	inputMovementLine(c, dictionary[SUB_C]) // Definition of subroutine C.
	fmt.Println("toggling snapshotting")
	collectOuptutMessage(c)
	inputMovementLine(c, "n") // Toggle snapshotting.
	collectOuptutMessage(c)

	// for c.GetState() == intcode.STATE_AWAITING_OUTPUT {
	// 	output := c.CollectOutput()
	// 	fmt.Printf("output: %c\n", output)
	// 	c.RunProgram()
	// }

	output := c.CollectOutput()
	fmt.Printf("output: %c\n", output)
	c.RunProgram()
	output = c.CollectOutput()
	fmt.Printf("output: %c\n", output)

	// 762405

	return output, nil
}

func getInputLayout(input string) (map[location.Point]int, string) {
	c := intcode.NewComputer()
	c.InputProgram(input)
	c.RunProgram()
	out := c.GetBatchOutput()

	loc := location.Point{X: 0, Y: 0}
	m := make(map[location.Point]int)
	var b strings.Builder
	for _, v := range out {
		m[loc] = v
		b.WriteByte(byte(v))
		if v == OBJ_NEWLINE {
			loc = location.Point{X: 0, Y: loc.Y + 1}
		} else {
			loc = location.Point{X: loc.X + 1, Y: loc.Y}
		}
	}
	return m, b.String()
}

func checksumIntersections(intersctions map[location.Point]struct{}) int {
	checksum := 0
	for point := range intersctions {
		checksum += point.X * point.Y
	}
	return checksum
}

const (
	OBJ_SCAFFOLD = 35
	OBJ_EMPTY    = 46
	OBJ_NEWLINE  = 10
	OBJ_LOSTBOT  = int('X')
	OBJ_NORTHBOT = int('^')
	OBJ_SOUTHBOT = int('v')
	OBJ_EASTBOT  = int('>')
	OBJ_WESTBOT  = int('<')

	DIR_NORTH = 1
	DIR_SOUTH = 2
	DIR_WEST  = 3
	DIR_EAST  = 4
)

type scaffoldWorld struct {
	c         *intcode.Computer
	vacuumLoc location.Point
	vacuumDir int
	worldMap  map[location.Point]int
	repr      string
}

func newScaffoldWorld() *scaffoldWorld {
	return &scaffoldWorld{
		vacuumLoc: location.Point{X: 0, Y: 0},
		vacuumDir: DIR_NORTH,
		worldMap:  make(map[location.Point]int),
		repr:      "",
	}
}

func (s *scaffoldWorld) snapshotWorld(snapshotProg string) {
	s.c = intcode.NewComputer()
	s.c.InputProgram(snapshotProg)
	s.c.RunProgram()
	out := s.c.GetBatchOutput()

	loc := location.Point{X: 0, Y: 0}
	m := make(map[location.Point]int)
	var b strings.Builder
	for _, v := range out {
		m[loc] = v
		b.WriteByte(byte(v))

		if v == OBJ_EMPTY {
			delete(m, loc)
		}

		switch v {
		case OBJ_NORTHBOT, OBJ_SOUTHBOT, OBJ_EASTBOT, OBJ_WESTBOT:
			s.vacuumLoc = loc
			switch v {
			case OBJ_NORTHBOT:
				s.vacuumDir = DIR_NORTH
			case OBJ_SOUTHBOT:
				s.vacuumDir = DIR_SOUTH
			case OBJ_EASTBOT:
				s.vacuumDir = DIR_EAST
			case OBJ_WESTBOT:
				s.vacuumDir = DIR_WEST
			}
		}

		if v == OBJ_NEWLINE {
			loc = location.Point{X: 0, Y: loc.Y + 1}
		} else {
			loc = location.Point{X: loc.X + 1, Y: loc.Y}
		}
	}

	s.worldMap = m
	s.repr = b.String()
}

func (s *scaffoldWorld) tracePathForIntersections() map[location.Point]struct{} {
	visited := make(map[location.Point]struct{})
	intersections := make(map[location.Point]struct{})
	for s.vacuumCanMove() {
		switch {
		case s.vacuumCanMoveAhead():
			nextLoc := s.getAheadSpace()
			if _, ok := visited[nextLoc]; ok {
				intersections[nextLoc] = struct{}{}
			} else {
				visited[nextLoc] = struct{}{}
			}
			s.vacuumLoc = nextLoc
		case s.vacuumCanMoveLeft():
			nextLoc := s.getLeftSpace()
			if _, ok := visited[nextLoc]; ok {
				intersections[nextLoc] = struct{}{}
			} else {
				visited[nextLoc] = struct{}{}
			}
			s.vacuumLoc = nextLoc
			s.vacuumDir = s.getTurnLeftDir()
		case s.vacuumCanMoveRight():
			nextLoc := s.getRightSpace()
			if _, ok := visited[nextLoc]; ok {
				intersections[nextLoc] = struct{}{}
			} else {
				visited[nextLoc] = struct{}{}
			}
			s.vacuumLoc = nextLoc
			s.vacuumDir = s.getTurnRightDir()
		}

	}
	return intersections
}

func (s *scaffoldWorld) tracePathForPath() []string {
	path := make([]string, 0)
	moveCounter := 0
	for s.vacuumCanMove() {
		switch {
		case s.vacuumCanMoveAhead():
			nextLoc := s.getAheadSpace()
			s.vacuumLoc = nextLoc
			moveCounter += 1
		case s.vacuumCanMoveLeft():
			nextLoc := s.getLeftSpace()
			s.vacuumLoc = nextLoc
			s.vacuumDir = s.getTurnLeftDir()
			if moveCounter > 0 {
				path = append(path, strconv.Itoa(moveCounter))
			}
			path = append(path, "L")
			moveCounter = 1
		case s.vacuumCanMoveRight():
			nextLoc := s.getRightSpace()
			s.vacuumLoc = nextLoc
			s.vacuumDir = s.getTurnRightDir()
			if moveCounter > 0 {
				path = append(path, strconv.Itoa(moveCounter))
			}
			path = append(path, "R")
			moveCounter = 1
		}
	}
	if moveCounter > 0 {
		path = append(path, strconv.Itoa(moveCounter))
	}
	return path
}

func (s *scaffoldWorld) vacuumCanMove() bool {
	return s.vacuumCanMoveAhead() || s.vacuumCanMoveLeft() || s.vacuumCanMoveRight()
}

func (s *scaffoldWorld) vacuumCanMoveAhead() bool {
	nextLoc := s.getAheadSpace()
	v, ok := s.worldMap[nextLoc]
	return ok && v == OBJ_SCAFFOLD
}

func (s *scaffoldWorld) vacuumCanMoveLeft() bool {
	nextLoc := s.getLeftSpace()
	v, ok := s.worldMap[nextLoc]
	return ok && v == OBJ_SCAFFOLD
}

func (s *scaffoldWorld) vacuumCanMoveRight() bool {
	nextLoc := s.getRightSpace()
	v, ok := s.worldMap[nextLoc]
	return ok && v == OBJ_SCAFFOLD
}

func (s *scaffoldWorld) getAheadSpace() location.Point {
	var nextLoc location.Point
	switch s.vacuumDir {
	case DIR_NORTH:
		nextLoc = location.Point{X: s.vacuumLoc.X, Y: s.vacuumLoc.Y - 1}
	case DIR_SOUTH:
		nextLoc = location.Point{X: s.vacuumLoc.X, Y: s.vacuumLoc.Y + 1}
	case DIR_EAST:
		nextLoc = location.Point{X: s.vacuumLoc.X + 1, Y: s.vacuumLoc.Y}
	case DIR_WEST:
		nextLoc = location.Point{X: s.vacuumLoc.X - 1, Y: s.vacuumLoc.Y}
	}
	return nextLoc
}

func (s *scaffoldWorld) getLeftSpace() location.Point {
	var nextLoc location.Point
	switch s.vacuumDir {
	case DIR_NORTH:
		nextLoc = location.Point{X: s.vacuumLoc.X - 1, Y: s.vacuumLoc.Y}
	case DIR_SOUTH:
		nextLoc = location.Point{X: s.vacuumLoc.X + 1, Y: s.vacuumLoc.Y}
	case DIR_EAST:
		nextLoc = location.Point{X: s.vacuumLoc.X, Y: s.vacuumLoc.Y - 1}
	case DIR_WEST:
		nextLoc = location.Point{X: s.vacuumLoc.X, Y: s.vacuumLoc.Y + 1}
	}
	return nextLoc
}

func (s *scaffoldWorld) getRightSpace() location.Point {
	var nextLoc location.Point
	switch s.vacuumDir {
	case DIR_NORTH:
		nextLoc = location.Point{X: s.vacuumLoc.X + 1, Y: s.vacuumLoc.Y}
	case DIR_SOUTH:
		nextLoc = location.Point{X: s.vacuumLoc.X - 1, Y: s.vacuumLoc.Y}
	case DIR_EAST:
		nextLoc = location.Point{X: s.vacuumLoc.X, Y: s.vacuumLoc.Y + 1}
	case DIR_WEST:
		nextLoc = location.Point{X: s.vacuumLoc.X, Y: s.vacuumLoc.Y - 1}
	}
	return nextLoc
}

func (s *scaffoldWorld) getTurnRightDir() int {
	switch s.vacuumDir {
	case DIR_NORTH:
		return DIR_EAST
	case DIR_SOUTH:
		return DIR_WEST
	case DIR_EAST:
		return DIR_SOUTH
	case DIR_WEST:
		return DIR_NORTH
	}
	panic("Vacuum not facing known direction...")
}

func (s *scaffoldWorld) getTurnLeftDir() int {
	switch s.vacuumDir {
	case DIR_NORTH:
		return DIR_WEST
	case DIR_SOUTH:
		return DIR_EAST
	case DIR_EAST:
		return DIR_NORTH
	case DIR_WEST:
		return DIR_SOUTH
	}
	panic("Vacuum not facing known direction...")
}

const (
	SUB_A   = int('A')
	SUB_B   = int('B')
	SUB_C   = int('C')
	SEP     = int(',')
	NEWLINE = int('\n')
)

func compressPath(path []string) (compressedPath []string, dictionary map[int]string) {

	// L,4,L,4,L,6,R,10,L,6,      A
	// L,4,L,4,L,6,R,10,L,6,      A
	// L,12,L,6,R,10,L,6,         B
	// R,8,R,10,L,6,              C
	// R,8,R,10,L,6,              C
	// L,4,L,4,L,6,R,10,L,6,      A
	// R,8,R,10,L,6,              C
	// L,12,L,6,R,10,L,6,         B
	// R,8,R,10,L,6,              C
	// L,12,L,6,R,10,L,6          B

	return []string{
			"A",
			"A",
			"B",
			"C",
			"C",
			"A",
			"C",
			"B",
			"C",
			"B",
		}, map[int]string{
			SUB_A: "L,4,L,4,L,6,R,10,L,6",
			SUB_B: "L,12,L,6,R,10,L,6",
			SUB_C: "R,8,R,10,L,6",
		}
}

func inputMovementLine(c *intcode.Computer, line string) {
	if state := c.GetState(); state != intcode.STATE_AWAITING_INPUT {
		panic(fmt.Sprintf("Computer should be awaiting input. In state %d", state))
	}

	for _, r := range line {
		inputASCIICode(c, int(r))
	}

	inputASCIICode(c, NEWLINE)
	// inputASCIICode(c, NEWLINE)
}

func inputASCIICode(c *intcode.Computer, code int) {
	if code == 10 {
		fmt.Println("Inputting: \\n")
	} else {
		fmt.Printf("Inputting: %c\n", rune(code))
	}

	c.Input(code)
	c.RunProgram()
}

func collectOuptutMessage(c *intcode.Computer) {
	fmt.Println("Collecting Output")
	var b strings.Builder
	for c.GetState() == intcode.STATE_AWAITING_OUTPUT {
		output := c.CollectOutput()
		b.WriteRune(rune(output))
		c.RunProgram()
	}
	fmt.Println(b.String())
}
