package day13

import (
	"aoc2019/internal/intcode"
	"aoc2019/internal/location"
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	cabinet := newArcadeCabinet()
	cabinet.LoadProgram(input)
	err := cabinet.RunProgramWithoutPlaying()
	if err != nil {
		return nil, err
	}
	screen := cabinet.GetScreenContents()
	nBlocks := countObjects(screen, OBJ_BLOCK)
	return nBlocks, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	cabinet := newArcadeCabinet()
	cabinet.LoadProgram(input)
	cabinet.HackToPlayForFree()
	err := cabinet.RunPerfectPlayProgram()
	if err != nil {
		return nil, err
	}
	score := cabinet.GetScore()
	return score, nil
}

func countObjects(screen map[location.Point]int, objType int) int {
	c := 0
	for _, v := range screen {
		if v == objType {
			c++
		}
	}
	return c
}

const (
	OBJ_EMPTY  = 0
	OBJ_WALL   = 1
	OBJ_BLOCK  = 2
	OBJ_PADDLE = 3
	OBJ_BALL   = 4

	JOYSTICK_RIGHT = 1
	JOYSTICK_STAY  = 0
	JOYSTICK_LEFT  = -1
)

type arcadeCabinet struct {
	ballLoc        location.Point
	c              *intcode.Computer
	paddleLoc      location.Point
	screen         map[location.Point]int
	segmentDisplay int
}

func newArcadeCabinet() *arcadeCabinet {
	return &arcadeCabinet{
		c:      intcode.NewComputer(),
		screen: make(map[location.Point]int),
	}
}

func (ac *arcadeCabinet) LoadProgram(prog string) {
	ac.c.InputProgram(prog)
}

func (ac *arcadeCabinet) RunProgramWithoutPlaying() error {
	_, err := ac.c.RunProgram()
	output := ac.c.GetBatchOutput()
	if len(output)%3 != 0 {
		return fmt.Errorf("Expected output to be in triples.")
	}

	for i := 0; i < len(output); i += 3 {
		x := output[i]
		y := output[i+1]
		obj := output[i+2]
		ac.screen[location.Point{X: x, Y: y}] = obj
	}

	return err
}

func (ac *arcadeCabinet) GetScreenContents() map[location.Point]int {
	return ac.screen
}

func (ac *arcadeCabinet) HackToPlayForFree() {
	ac.c.UpdateProgram(map[int]int{0: 2})
	ac.c.SetInterruptibleMode()
}

func (ac *arcadeCabinet) RunPerfectPlayProgram() error {
	state, err := ac.c.RunProgram()
	if err != nil {
		return err
	}
	for !(state == intcode.STATE_RUN_COMPLETE || state == intcode.STATE_RUNTIME_ERROR) {
		x, y, obj := 0, 0, 0
		for state == intcode.STATE_AWAITING_OUTPUT {
			x = ac.c.CollectOutput()

			state, err = ac.c.RunProgram()
			if err != nil {
				return err
			} else if state != intcode.STATE_AWAITING_OUTPUT {
				return fmt.Errorf("Expected program to be sending output")
			}
			y = ac.c.CollectOutput()

			state, err = ac.c.RunProgram()
			if err != nil {
				return err
			} else if state != intcode.STATE_AWAITING_OUTPUT {
				return fmt.Errorf("Expected program to be sending output")
			}
			obj = ac.c.CollectOutput()

			ac.paintObject(x, y, obj)

			state, err = ac.c.RunProgram()
			if err != nil {
				return err
			}
		}

		if state == intcode.STATE_AWAITING_INPUT {
			if ac.paddleLoc.X < ac.ballLoc.X {
				ac.c.Input(JOYSTICK_RIGHT)
			} else if ac.paddleLoc.X > ac.ballLoc.X {
				ac.c.Input(JOYSTICK_LEFT)
			} else {
				ac.c.Input(JOYSTICK_STAY)
			}

			state, err = ac.c.RunProgram()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ac *arcadeCabinet) paintObject(x, y, obj int) {
	if x == -1 && y == 0 {
		ac.segmentDisplay = obj
		return
	}

	loc := location.Point{X: x, Y: y}
	ac.screen[loc] = obj

	if obj == OBJ_BALL {
		ac.ballLoc = loc
	} else if obj == OBJ_PADDLE {
		ac.paddleLoc = loc
	}
}

func (ac *arcadeCabinet) GetScore() int {
	return ac.segmentDisplay
}

func getScreenImage(screen map[location.Point]int) string {
	minX, maxX := 0, 0
	minY, maxY := 0, 0
	for loc := range screen {
		if loc.X < minX {
			minX = loc.X
		}
		if loc.X > maxX {
			maxX = loc.X
		}
		if loc.Y < minY {
			minY = loc.Y
		}
		if loc.Y > maxY {
			maxY = loc.Y
		}
	}

	var b strings.Builder
	_, err := b.WriteRune('\n')
	if err != nil {
		panic("How did WriteRune return an error?!")
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			c, ok := screen[location.Point{X: x, Y: y}]
			if !ok {
				c = OBJ_EMPTY
			}

			var err error
			switch c {
			case OBJ_EMPTY:
				_, err = b.WriteRune(' ')
			case OBJ_WALL:
				_, err = b.WriteRune('X')
			case OBJ_BLOCK:
				_, err = b.WriteRune('#')
			case OBJ_PADDLE:
				_, err = b.WriteRune('T')
			case OBJ_BALL:
				_, err = b.WriteRune('*')
			default:
				panic("WTF: object not recognized")
			}
			if err != nil {
				panic("How did WriteRune return an error?!")
			}
		}
		_, err := b.WriteRune('\n')
		if err != nil {
			panic("How did WriteRune return an error?!")
		}
	}

	return b.String()
}
