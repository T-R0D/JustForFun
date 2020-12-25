package day14

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const v1Program = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

const v2Program = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

var v1ProgramInstructions = []emulatorInstruction{
	{
		Kind: instructionKindMask,
		Mask: &maskInstruction{
			Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		},
		Write: nil,
	},
	{
		Kind: instructionKindWrite,
		Mask: nil,
		Write: &writeInstruction{
			Address: 8,
			Value:   11,
		},
	},
	{
		Kind: instructionKindWrite,
		Mask: nil,
		Write: &writeInstruction{
			Address: 7,
			Value:   101,
		},
	},
	{
		Kind: instructionKindWrite,
		Mask: nil,
		Write: &writeInstruction{
			Address: 8,
			Value:   0,
		},
	},
}

var v2ProgramInstructions = []emulatorInstruction{
	{
		Kind: instructionKindMask,
		Mask: &maskInstruction{
			Mask: "000000000000000000000000000000X1001X",
		},
		Write: nil,
	},
	{
		Kind: instructionKindWrite,
		Mask: nil,
		Write: &writeInstruction{
			Address: 42,
			Value:   100,
		},
	},
	{
		Kind: instructionKindMask,
		Mask: &maskInstruction{
			Mask: "00000000000000000000000000000000X0XX",
		},
		Write: nil,
	},
	{
		Kind: instructionKindWrite,
		Mask: nil,
		Write: &writeInstruction{
			Address: 26,
			Value:   1,
		},
	},
}

func TestParseProgram(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput []emulatorInstruction
	}{
		{
			name:           "v1 Program parses correctly",
			input:          v1Program,
			expectedOutput: v1ProgramInstructions,
		},
		{
			name:           "v2 Program parses correctly",
			input:          v2Program,
			expectedOutput: v2ProgramInstructions,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			instructions, err := parseProgram(tc.input)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, instructions)
		})
	}
}

func TestMaskApplicatorApply(t *testing.T) {
	testCases := []struct {
		representation   string
		inValue          uint64
		expectedOutValue uint64
	}{
		{
			representation:   "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			inValue:          11,
			expectedOutValue: 73,
		},
		{
			representation:   "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			inValue:          101,
			expectedOutValue: 101,
		},
		{
			representation:   "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			inValue:          0,
			expectedOutValue: 64,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d &| %s -> %d", tc.inValue, tc.representation, tc.expectedOutValue)
		t.Run(name, func(t *testing.T) {
			// Arrange.
			applicator, err := newMaskApplicator(tc.representation)
			assert.NoError(t, err)

			// Act.
			outValue := applicator.Apply(tc.inValue)

			// Assert.
			assert.Equal(t, tc.expectedOutValue, outValue)
		})
	}
}

func TestSeaportComputerEmulator(t *testing.T) {
	// Arrange.
	emulator, err := newSeaportComputerEmulator()
	assert.NoError(t, err)

	// Act.
	for _, instruction := range v1ProgramInstructions {
		err := emulator.TakeInstruction(instruction)
		assert.NoError(t, err)
	}

	// Assert.
	assert.Equal(t, uint64(165), emulator.MemorySum())
}

func TestMemoryApplicatorV2(t *testing.T) {
	testCases := []struct {
		maskRepresentation string
		inValue            uint64
		expectedOutValues  []uint64
	}{
		{
			maskRepresentation: "000000000000000000000000000000X1001X",
			inValue:            42,
			expectedOutValues:  []uint64{26, 27, 58, 59},
		},
		{
			maskRepresentation: "00000000000000000000000000000000X0XX",
			inValue:            26,
			expectedOutValues:  []uint64{16, 17, 18, 19, 24, 25, 26, 27},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d &| %s produces %d values", tc.inValue, tc.maskRepresentation, len(tc.expectedOutValues))
		t.Run(name, func(t *testing.T) {
			// Arrange.
			applicator, err := newMaskApplicatorV2(tc.maskRepresentation)
			assert.NoError(t, err)

			// Act.
			outValues := applicator.Apply(tc.inValue)

			// Assert.
			assert.Equal(t, tc.expectedOutValues, outValues)
		})
	}
}

func TestSeaportComputerEmulatorV2(t *testing.T) {
	// Arrange.
	emulator, err := newSeaportComputerEmulatorV2()
	assert.NoError(t, err)

	// Act.
	for _, instruction := range v2ProgramInstructions {
		err := emulator.TakeInstruction(instruction)
		assert.NoError(t, err)
	}

	// Assert.
	assert.Equal(t, uint64(208), emulator.MemorySum())
}