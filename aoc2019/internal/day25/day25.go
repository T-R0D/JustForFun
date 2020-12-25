package day25

import (
	"aoc2019/internal/intcode"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	password := searchShipForPassword(input)
	return password, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	return nil, nil
}

type searchState struct {
}

func searchShipForPassword(prog string) string {
	reader := bufio.NewReader(os.Stdin)

	c := intcode.NewComputer()
	c.SetInterruptibleMode()
	c.InputProgram(prog)
	c.RunProgram()

	for {
		for c.GetState() == intcode.STATE_AWAITING_OUTPUT {
			line := readLine(c)
			fmt.Print(line)
		}

		inputLine, _ := reader.ReadString('\n')
		writeLine(c, inputLine)
	}

	return ""
}

func readLine(c *intcode.Computer) string {
	if c.GetState() != intcode.STATE_AWAITING_OUTPUT {
		panic("Computer should have output ready to read line from it")
	}

	var b strings.Builder
	for c.GetState() == intcode.STATE_AWAITING_OUTPUT {
		output := c.CollectOutput()
		b.WriteRune(rune(output))
		c.RunProgram()
		if output == int('\n') {
			break
		}
	}
	return b.String()
}

func writeLine(c *intcode.Computer, line string) {
	if c.GetState() != intcode.STATE_AWAITING_INPUT {
		panic("Computer should be awaiting input to write line.")
	}

	for _, r := range line {
		if c.GetState() != intcode.STATE_AWAITING_INPUT {
			panic("Computer should be awaiting input to write character in line.")
		}
		c.Input(int(r))
		c.RunProgram()
	}
	if line[len(line)-1] != '\n' {
		if c.GetState() != intcode.STATE_AWAITING_INPUT {
			panic("Computer should be awaiting input to write newline char.")
		}
		c.Input(int('\n'))
		c.RunProgram()
	}
}
