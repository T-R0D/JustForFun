package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleProgram = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestParseProgram(t *testing.T) {
	expectedProgram := []instruction{
		{key: instructionKeyNop, value: 0},
		{key: instructionKeyAcc, value: 1},
		{key: instructionKeyJmp, value: 4},
		{key: instructionKeyAcc, value: 3},
		{key: instructionKeyJmp, value: -3},
		{key: instructionKeyAcc, value: -99},
		{key: instructionKeyAcc, value: 1},
		{key: instructionKeyJmp, value: -4},
		{key: instructionKeyAcc, value: 6},
	}

	actualProgram := parseProgram(sampleProgram)

	assert.Equal(t, expectedProgram, actualProgram)
}

func TestRunUntilLoopEncountered(t *testing.T) {
	program := parseProgram(sampleProgram)
	sut := newHandheldConsoleEmulator(program)

	sut.RunUntilLoopEncountered()
	actualAccumulatorValue := sut.Accumulator()

	assert.Equal(t, 5, actualAccumulatorValue)
}
