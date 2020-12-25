package day19

import (
	"aoc2019/internal/intcode"
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	nTractoredLocations := countTractoredLocations(input, 50, 50)
	return nTractoredLocations, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {

	// dumpBeamSection(input, 39885, 99915)

	x, y := findCoordinateWhereShipFitsInBeam(input, 100, 100)
	return (x * 10000) + y, nil
}

const (
	UNAFFECTED = 0
	TRACTORED  = 1
)

func countTractoredLocations(prog string, xRange, yRange int) int {

	// var b strings.Builder

	tractoredLocations := 0
	for x := 0; x < xRange; x += 1 {
		for y := 0; y < yRange; y += 1 {
			response := checkCoordinate(prog, x, y)
			if response == TRACTORED {
				tractoredLocations += 1
			}

			// if response == TRACTORED {
			// 	b.WriteRune('T')
			// } else {
			// 	b.WriteRune('.')
			// }
		}

		// b.WriteRune('\n')
	}

	// fmt.Println(b.String())

	return tractoredLocations
}

func checkCoordinate(prog string, x, y int) int {
	c := intcode.NewComputer()
	c.SetInterruptibleMode()
	c.InputProgram(prog)
	c.RunProgram()
	inputToComputer(c, x)
	inputToComputer(c, y)
	response := collectOutputFromComputer(c)
	return response
}

func inputToComputer(c *intcode.Computer, input int) {
	if c.GetState() != intcode.STATE_AWAITING_INPUT {
		panic("Expected computer to be awaiting input.")
	}
	c.Input(input)
	c.RunProgram()
}

func collectOutputFromComputer(c *intcode.Computer) int {
	if c.GetState() != intcode.STATE_AWAITING_OUTPUT {
		panic("Expected computer to be awaiting output.")
	}
	output := c.CollectOutput()
	c.RunProgram()
	return output
}

func findCoordinateWhereShipFitsInBeam(prog string, shipHeight, shipWidth int) (int, int) {
	row := shipWidth
	rowWidth := 0
	for rowWidth < shipWidth {
		row += 1
		rowWidth = getBeamWidthInRow(prog, row)
	}

	resX, resY := -1, -1
	shipFits := false
	for y := row; !shipFits; y += 1 {
		x := findLastXShipFitsAt(prog, shipWidth, y)
		if x == -1 {
			continue
		}

		// fmt.Printf("(%d, %d)\n", x, y)

		shipFits = shipFitsAtCoordinate(prog, x, y, shipHeight, shipWidth)
		if shipFits {
			resX, resY = x, y
		}
	}
	return resX, resY
}

func getBeamWidthInRow(prog string, row int) int {
	width := 0
	x := 0
	response := checkCoordinate(prog, x, row)
	for response == UNAFFECTED {
		response = checkCoordinate(prog, x, row)
		x += 1
	}
	for response == TRACTORED {
		width += 1
		x += 1
		response = checkCoordinate(prog, x, row)
	}
	return width
}

func findLastXShipFitsAt(prog string, shipWidth, y int) int {
	x := 0
	for checkCoordinate(prog, x, y) == UNAFFECTED {
		x += 1
	}
	shipFits := false
	for checkCoordinate(prog, x, y) == TRACTORED && checkCoordinate(prog, x+shipWidth-1, y) == TRACTORED {
		shipFits = true
		x += 1
	}
	if shipFits {
		return x - 1
	} else {
		return -1
	}
}

func shipFitsAtCoordinate(prog string, x, y, shipHeight, shipWidth int) bool {
	for i := 0; i < shipWidth; i += 1 {
		response := checkCoordinate(prog, x+i, y)

		// fmt.Printf("\tx = %d -> %d\n", x+i, response)

		if response == UNAFFECTED {
			return false
		}
	}

	for i := 0; i < shipHeight; i += 1 {
		response := checkCoordinate(prog, x, y+i)

		// fmt.Printf("\ty = %d -> %d\n", y+i, response)

		if response == UNAFFECTED {
			return false
		}
	}

	return true
}

func dumpBeamSection(prog string, x, y int) {
	var b strings.Builder
	for i := y; i < y+50; i += 1 {
		for j := x; j < x+150; j += 1 {
			if checkCoordinate(prog, j, i) == TRACTORED {
				b.WriteRune('T')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	fmt.Println(b.String())
}
