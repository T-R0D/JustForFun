// Question: Can we improve on brute force changing every possible operation
//           and running the program on each change in part 1?

package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	program := parseProgram(input)

	emulator := newHandheldConsoleEmulator(program)
	emulator.RunUntilLoopEncountered()

	return strconv.Itoa(emulator.Accumulator()), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	program := parseProgram(input)

	for i := range program {
		modificationMade, newProgram := switchNopToJmp(program, i)
		if !modificationMade {
			continue
		}

		emulator := newHandheldConsoleEmulator(newProgram)
		success := emulator.RunUntilLoopOrSuccess()

		if success {
			return strconv.Itoa(emulator.Accumulator()), nil
		}
	}

	for i := range program {
		modificationMade, newProgram := switchJmpToNop(program, i)
		if !modificationMade {
			continue
		}

		emulator := newHandheldConsoleEmulator(newProgram)
		success := emulator.RunUntilLoopOrSuccess()

		if success {
			return strconv.Itoa(emulator.Accumulator()), nil
		}
	}

	return "", errors.New("solution not found")
}

func parseProgram(input string) []instruction {
	lines := strings.Split(input, "\n")
	program := make([]instruction, 0, len(lines))
	for _, line := range lines {
		keyAndValue := strings.Split(line, " ")
		key := instructionKey(keyAndValue[0])
		value, err := strconv.Atoi(keyAndValue[1])
		if err != nil {
			panic(err)
		}
		program = append(program, instruction{key: key, value: value})
	}
	return program
}

type handheldConsoleEmulator struct {
	program              []instruction
	executedInstructions map[int]struct{}
	programCounter       int
	accumulator          int
}

type instructionKey string

const (
	instructionKeyAcc = "acc"
	instructionKeyJmp = "jmp"
	instructionKeyNop = "nop"
)

type instruction struct {
	key   instructionKey
	value int
}

func newHandheldConsoleEmulator(program []instruction) handheldConsoleEmulator {
	return handheldConsoleEmulator{
		program:              program,
		executedInstructions: map[int]struct{}{},
		programCounter:       0,
		accumulator:          0,
	}
}

func (h *handheldConsoleEmulator) RunUntilLoopEncountered() {
	for {
		if _, ok := h.executedInstructions[h.programCounter]; ok {
			break
		}

		nextInstruction := h.program[h.programCounter]

		h.executedInstructions[h.programCounter] = struct{}{}

		h.executeInstruction(nextInstruction)
	}
}

func (h *handheldConsoleEmulator) RunUntilLoopOrSuccess() bool {
	success := false

	for {
		if h.programCounter >= len(h.program) {
			success = true
			break
		}

		if _, ok := h.executedInstructions[h.programCounter]; ok {
			break
		}

		nextInstruction := h.program[h.programCounter]

		h.executedInstructions[h.programCounter] = struct{}{}

		h.executeInstruction(nextInstruction)
	}

	return success
}

func (h *handheldConsoleEmulator) executeInstruction(nextInstruction instruction) {
	switch nextInstruction.key {
	case instructionKeyAcc:
		h.accumulator += nextInstruction.value
		h.programCounter++
	case instructionKeyJmp:
		h.programCounter += nextInstruction.value
	case instructionKeyNop:
		h.programCounter++
	default:
		panic("unrecognized instruction" + string(nextInstruction.key))
	}
}

func (h *handheldConsoleEmulator) Accumulator() int {
	return h.accumulator
}

func (h *handheldConsoleEmulator) PrintDebugInfo() {
	fmt.Println("PC:", h.programCounter, "ACC:", h.accumulator, "EXE:", h.executedInstructions)
}

func switchNopToJmp(program []instruction, index int) (bool, []instruction) {
	if program[index].key != instructionKeyNop || program[index].value == 0 {
		return false, nil
	}

	newProgram := make([]instruction, len(program))
	copy(newProgram, program)
	newProgram[index].key = instructionKeyJmp

	return true, newProgram
}

func switchJmpToNop(program []instruction, index int) (bool, []instruction) {
	if program[index].key != instructionKeyJmp {
		return false, nil
	}

	newProgram := make([]instruction, len(program))
	copy(newProgram, program)
	newProgram[index].key = instructionKeyNop

	return true, newProgram
}
