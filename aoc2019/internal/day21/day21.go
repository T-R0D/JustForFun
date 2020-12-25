package day21

import (
	"aoc2019/internal/intcode"
	"fmt"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	hullDamage := programAndWalkSpringDroid(input)
	return hullDamage, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	hullDamage := programAndRunSpringDroid(input)
	return hullDamage, nil
}

func programAndWalkSpringDroid(prog string) int {
	c := intcode.NewComputer()
	c.SetInterruptibleMode()
	c.InputProgram(prog)
	c.RunProgram()

	line := readLine(c)
	// fmt.Println(line)

	sprinscriptProg := []string{
		"NOT A T",
		"NOT B J",
		"OR T J",
		"NOT C T",
		"OR T J",
		"AND D J",
		"WALK",
	}

	for _, line := range sprinscriptProg {
		writeLine(c, line)
	}

	line = readLine(c) // newline
	line = readLine(c) // "Walking..."
	// fmt.Print(line)
	line = readLine(c) // newline

	if c.GetState() == intcode.STATE_AWAITING_OUTPUT {
		output := c.CollectOutput()
		if output > 255 {
			return output
		}

		var b strings.Builder
		for c.GetState() == intcode.STATE_AWAITING_OUTPUT {
			line = readLine(c)
			b.WriteString(line)
		}
		fmt.Print(b.String())
	}

	return 0
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

func programAndRunSpringDroid(prog string) int {
	c := intcode.NewComputer()
	c.SetInterruptibleMode()
	c.InputProgram(prog)
	c.RunProgram()

	line := readLine(c)
	// fmt.Println(line)

	// This could probably be improved through some boolean algebra tricks,
	// but I need to move on.
	sprinscriptProg := []string{
		// !A
		"NOT A J",

		// || (!B || !C) && D -> || !(B && C) && D
		"OR B T",
		"AND C T",
		"NOT T T",
		"AND D T",
		"OR T J",

		// || (!C && D && H)
		"NOT C T",
		"AND D T",
		"AND H T",
		"OR T J",

		// && !(D && (!E && !H)) -> && !(D && !(E || H)) -> && !D || (E || H)
		"NOT D T",
		"OR E T",
		"OR H T",
		"AND T J",

		"RUN",
	}

	for _, line := range sprinscriptProg {
		writeLine(c, line)
	}

	line = readLine(c) // newline
	line = readLine(c) // "Walking..."
	// fmt.Print(line)
	line = readLine(c) // newline

	if c.GetState() == intcode.STATE_AWAITING_OUTPUT {
		output := c.CollectOutput()
		if output > 255 {
			return output
		}

		var b strings.Builder
		for c.GetState() == intcode.STATE_AWAITING_OUTPUT {
			line = readLine(c)
			b.WriteString(line)
		}
		fmt.Print(b.String())
	}

	return 0
}
