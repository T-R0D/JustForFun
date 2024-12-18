package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	initialRegistersAndProgram, err := parseInitialRegistersAndProgram(input)
	if err != nil {
		return "", err
	}

	emulator := NewEmulator(initialRegistersAndProgram.Registers, initialRegistersAndProgram.Program)
	output, _ := emulator.Execute(false)

	return output, nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	initialRegistersAndProgram, err := parseInitialRegistersAndProgram(input)
	if err != nil {
		return "", err
	}

	quineInitValue, ok := findQuineInitValue(
		initialRegistersAndProgram.Program, actualInputArbitraryOperation)
	if !ok {
		return "", fmt.Errorf("a quine-producing initial value was not found")
	}

	return strconv.Itoa(quineInitValue), nil
}

type initialRegistersAndProgram struct {
	Registers [3]int
	Program   []int
}

func parseInitialRegistersAndProgram(input string) (*initialRegistersAndProgram, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return nil, fmt.Errorf("there were not exactly 2 sections")
	}

	registers, err := parseInitialRegisters(sections[0])
	if err != nil {
		return nil, err
	}

	program, err := parseProgram(sections[1])
	if err != nil {
		return nil, err
	}

	return &initialRegistersAndProgram{
		Registers: registers,
		Program:   program,
	}, nil
}

func parseInitialRegisters(section string) ([3]int, error) {
	var registers [3]int
	lines := strings.Split(section, "\n")
	if len(lines) != 3 {
		return registers, fmt.Errorf("there were not 3 lines in the register section")
	}

	labels := []rune{'A', 'B', 'C'}
	for i, line := range lines {
		valueString := strings.ReplaceAll(line, fmt.Sprintf("Register %c: ", labels[i]), "")
		value, err := strconv.Atoi(valueString)
		if err != nil {
			return registers, err
		}
		registers[i] = value
	}

	return registers, nil
}

func parseProgram(section string) ([]int, error) {
	codeStrings := strings.Split(strings.ReplaceAll(section, "Program: ", ""), ",")
	program := make([]int, len(codeStrings))
	for i, codeString := range codeStrings {
		value, err := strconv.Atoi(codeString)
		if err != nil {
			return []int{}, err
		}
		program[i] = value
	}

	return program, nil
}

func findQuineInitValue(program []int, arbitraryOperation func(int) int) (int, bool) {
	// This one was tough for me. It's a "reverse engineering" one, but I was
	// unable to completely reverse engineer it.
	// In essence, the program in my input takes a starting value, performs
	// some arbitrary operations (divisions and XORs) with that value, outputs
	// the result, then reduces the starting value by integer dividing by
	// 8 (2 to the power 3), and repeats if the initial value was not reduced
	// to zero. I'd guess this is what most inputs did, with some variation
	// in the "arbitrary operations". That or the constants were different (my
	// input had XORs with 3 a couple times and the divisor was 2 raised to
	// the 3rd).
	// The registers were used as follows:
	//
	// `A` was the initial value/value that carried over from iteration to
	// iteration. It was the one used in the JNZ test at the end of the
	// program.
	//
	// `B` and `C` were reinitialized by conducting and held the results of
	// the "arbitrary operations". Additionally, the result in `B` is what
	// went to the output. Because of the registers being wiped each iteration
	// their starting values did not matter, so I suspect all inputs had them
	// starting at 0.
	//
	// In general, an iteration was something along the lines of taking
	// `A` % 8, doing the arbitrary operations, outputting the result, and
	// reducing `A` by dividing by 8. This means that we could kind of build
	// the program in reverse order, going from the last output/instruction to
	// the first by taking a dividend, multiplying it by 8, and then adding
	// a remainder (computed by reversing the "arbitrary operations"). The
	// resulting divisor (at the top of the program) for one iteration becomes
	// the quotient at the bottom of the iteration that should come before in
	// a proper execution.
	//
	// However, I couldn't figure out how to reverse the operations. This led
	// me to a more heavy handed approach, where I did a sort of DFS that
	// started with all divisors/quotients ([0, 7]), then tested the result
	// of executing the "arbitrary operations" on that divisor/quotient
	// (also [0, 7]), and testing if the result matched the corresponding
	// output/instruction in order to proceed on that "search branch". If
	// there was a match, we multiply the divisor/quotient by 8, and try
	// every possible remainder (again,[0, 7]) for the next
	// instruction/output. Repeat until we've matched the whole
	// instruction/output list (in reverse order).
	//
	// Finally, just to double check myself, I ran the solution through
	// my `Emulator` in a mode that would abort if a quine was not produced
	// (well, perhaps not technically a quine, but close enough).

	if len(program) == 0 {
		return 0, false
	}

	type searchState struct {
		Dividend int
		Depth    int
	}
	frontier := queue.NewLifo[searchState]()
	for d := 1; d < 8; d += 1 {
		frontier.Push(searchState{Dividend: d, Depth: len(program) - 1})
	}

	solution := 0
	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		if state.Depth > len(program) {
			continue
		}

		output := arbitraryOperation(state.Dividend)

		if output == program[0] && state.Depth == 0 {
			solution = state.Dividend
			continue
		}

		if output != program[state.Depth] {
			continue
		}

		for r := range 8 {
			dividend := (state.Dividend * 8) + r
			frontier.Push(searchState{Dividend: dividend, Depth: state.Depth - 1})
		}
	}

	registers := [3]int{solution, 0, 0}
	emulator := NewEmulator(registers, program)
	_, quineFound := emulator.Execute(true)
	if quineFound {
		return solution, true
	}

	return 0, false
}

func actualInputArbitraryOperation(dividend int) int {
	b := ((dividend) & 0x07) ^ 3
	c := ((dividend) >> b) &0x07
	return (b ^ c) ^ 3
}

type Emulator struct {
	registers          [3]int
	instructionPointer int
	program            []int
	operations         [8]instruction
	outputBuffer       []int
}

type instruction func(operand int)

func NewEmulator(registerValues [3]int, newProgram []int) *Emulator {
	e := &Emulator{}
	for i := range e.registers {
		e.registers[i] = registerValues[i]
	}
	e.program = append([]int{}, newProgram...)
	e.operations = [8]instruction{
		e.adv,
		e.bxl,
		e.bst,
		e.jnz,
		e.bxc,
		e.out,
		e.bdv,
		e.cdv,
	}

	return e
}

func (e *Emulator) Execute(quineMode bool) (string, bool) {
	for 0 <= e.instructionPointer && e.instructionPointer+1 < len(e.program) {
		opCode := e.program[e.instructionPointer]
		operand := e.program[e.instructionPointer+1]
		op, operand := e.operations[opCode], operand
		op(operand)
		e.instructionPointer += 2

		if quineMode && opCode == 5 {
			lastOutputIndex := len(e.outputBuffer) - 1
			if lastOutputIndex >= len(e.program) ||
				e.outputBuffer[lastOutputIndex] != e.program[lastOutputIndex] {

				return "", false
			}
		}
	}

	if quineMode && len(e.outputBuffer) == len(e.program) {
		return e.JoinOutputIntoString(), true
	}

	return e.JoinOutputIntoString(), false
}

func (e *Emulator) adv(operand int) {
	e.registers[0] = e.registers[0] >> e.comboOperandValue(operand)
}

func (e *Emulator) bxl(operand int) {
	e.registers[1] = e.registers[1] ^ operand
}

func (e *Emulator) bst(operand int) {
	e.registers[1] = e.comboOperandValue(operand) & 0x07
}

func (e *Emulator) jnz(operand int) {
	if e.registers[0] == 0 {
		return
	}
	e.instructionPointer = operand - 2
}

func (e *Emulator) bxc(operand int) {
	e.registers[1] = e.registers[1] ^ e.registers[2]
}

func (e *Emulator) out(operand int) {
	e.outputBuffer = append(e.outputBuffer, e.comboOperandValue(operand)&0x07)
}

func (e *Emulator) bdv(operand int) {
	e.registers[1] = e.registers[0] >> e.comboOperandValue(operand)
}

func (e *Emulator) cdv(operand int) {
	e.registers[2] = e.registers[0] >> e.comboOperandValue(operand)
}

func (e *Emulator) comboOperandValue(operand int) int {
	if 0 <= operand && operand <= 3 {
		return operand
	} else if 4 <= operand && operand <= 6 {
		return e.registers[operand-4]
	} else {
		return operand
	}
}

func (e *Emulator) JoinOutputIntoString() string {
	out := []string{}
	for _, value := range e.outputBuffer {
		out = append(out, strconv.Itoa(value))
	}

	return strings.Join(out, ",")
}
