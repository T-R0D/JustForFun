package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	instructions, err := parseProgram(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing the program from the inupt")
	}

	emulator, err := newSeaportComputerEmulator()
	if err != nil {
		return "", err
	}

	for i, instruction := range instructions {
		err := emulator.TakeInstruction(instruction)
		if err != nil {
			return "", errors.Wrapf(err, "executing instruction %d", i)
		}
	}

	memorySum := emulator.MemorySum()

	return fmt.Sprintf("%d", memorySum), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	instructions, err := parseProgram(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing the program from the inupt")
	}

	emulator, err := newSeaportComputerEmulatorV2()
	if err != nil {
		return "", err
	}

	for i, instruction := range instructions {
		err := emulator.TakeInstruction(instruction)
		if err != nil {
			return "", errors.Wrapf(err, "executing instruction %d", i)
		}
	}

	memorySum := emulator.MemorySum()

	return fmt.Sprintf("%d", memorySum), nil
}

type instructionKind int

const (
	instructionKindMask  = iota
	instructionKindWrite = iota
)

type emulatorInstruction struct {
	Kind  instructionKind
	Mask  *maskInstruction
	Write *writeInstruction
}

type maskInstruction struct {
	Mask string
}

type writeInstruction struct {
	Address uint64
	Value   uint64
}

func parseProgram(input string) ([]emulatorInstruction, error) {
	lines := strings.Split(input, "\n")

	maskInstructionRegexp, err := regexp.Compile(`^mask = (?P<mask>[X10]{36})$`)
	if err != nil {
		return nil, errors.Wrap(err, "compiling RE for mask instructions")
	}

	writeInstructionRegexp, err := regexp.Compile(`^mem\[(?P<address>\d+)\] = (?P<value>0|[1-9]\d*)$`)
	if err != nil {
		return nil, errors.Wrap(err, "compiling RE for mem instructions")
	}

	instructions := make([]emulatorInstruction, 0, len(lines))
	for i, line := range lines {
		instruction := emulatorInstruction{}
		if match := maskInstructionRegexp.FindStringSubmatch(line); len(match) > 0 {
			instruction.Kind = instructionKindMask
			instruction.Mask = &maskInstruction{
				Mask: match[1],
			}
		} else if match := writeInstructionRegexp.FindStringSubmatch(line); len(match) > 0 {
			addressStr, valueStr := match[1], match[2]

			address, err := strconv.Atoi(addressStr)
			if err != nil {
				return nil, errors.Wrap(err, "converting memory address")
			}

			value, err := strconv.Atoi(valueStr)
			if err != nil {
				return nil, errors.Wrap(err, "converting writeValue")
			}

			instruction.Kind = instructionKindWrite
			instruction.Write = &writeInstruction{
				Address: uint64(address),
				Value:   uint64(value),
			}
		} else {
			return nil, errors.Wrapf(err, "line %d did not match any instruction regex", i)
		}

		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

type seaportComputerEmulator struct {
	maskApplicator *maskApplicator
	memory         map[uint64]uint64
}

func newSeaportComputerEmulator() (*seaportComputerEmulator, error) {
	applicator, err := newMaskApplicator(noEffectMask)
	if err != nil {
		return nil, errors.Wrap(err, "constructing a new seaportComputerEmulator")
	}

	return &seaportComputerEmulator{
		maskApplicator: applicator,
		memory:         map[uint64]uint64{},
	}, nil
}

func (e *seaportComputerEmulator) TakeInstruction(instruction emulatorInstruction) error {
	switch instruction.Kind {
	case instructionKindMask:
		return e.setNewMaskApplicator(*instruction.Mask)
	case instructionKindWrite:
		return e.writeToMemory(*instruction.Write)
	default:
		return errors.Errorf("invalid instruction kind supplied to TakeInstruction: %v", instruction.Kind)
	}
}

func (e *seaportComputerEmulator) setNewMaskApplicator(instruction maskInstruction) error {
	applicator, err := newMaskApplicator(instruction.Mask)
	if err != nil {
		return errors.Wrap(err, "executing setNewMaskApplicator instruction")
	}
	e.maskApplicator = applicator
	return nil
}

func (e *seaportComputerEmulator) writeToMemory(instruction writeInstruction) error {
	valueToWrite := e.maskApplicator.Apply(instruction.Value)
	e.memory[instruction.Address] = valueToWrite
	return nil
}

func (e *seaportComputerEmulator) MemorySum() uint64 {
	sum := uint64(0)
	for _, value := range e.memory {
		sum += value
	}
	return sum
}

const (
	noEffectMask = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
)

type maskApplicator struct {
	representation string
	andBits        uint64
	orBits         uint64
}

func newMaskApplicator(representation string) (*maskApplicator, error) {
	if matched, err := regexp.MatchString(`[X10]{36}`, representation); !matched || err != nil {
		return nil, errors.New("representation does not fit the regular expression `[X10]{36}`")
	}

	andBits := ^uint64(0)
	orBits := uint64(0)
	for _, bit := range representation {
		andBits <<= 1
		orBits <<= 1

		switch bit {
		case '1':
			andBits |= 1
			orBits |= 1
		case '0':
			andBits |= 0
			orBits |= 0
		case 'X':
			andBits |= 1
			orBits |= 0
		}
	}

	return &maskApplicator{
		representation: representation,
		andBits:        andBits,
		orBits:         orBits,
	}, nil
}

func (a *maskApplicator) Apply(in uint64) uint64 {
	out := in
	out &= a.andBits
	out |= a.orBits
	return out
}

type seaportComputerEmulatorV2 struct {
	maskApplicator *maskApplicatorV2
	memory         map[uint64]uint64
}

func newSeaportComputerEmulatorV2() (*seaportComputerEmulatorV2, error) {
	applicator, err := newMaskApplicatorV2(noEffectMask)
	if err != nil {
		return nil, errors.Wrap(err, "constructing a new seaportComputerEmulator")
	}

	return &seaportComputerEmulatorV2{
		maskApplicator: applicator,
		memory:         map[uint64]uint64{},
	}, nil
}

func (e *seaportComputerEmulatorV2) TakeInstruction(instruction emulatorInstruction) error {
	switch instruction.Kind {
	case instructionKindMask:
		return e.setNewMaskApplicator(*instruction.Mask)
	case instructionKindWrite:
		return e.writeToMemory(*instruction.Write)
	default:
		return errors.Errorf("invalid instruction kind supplied to TakeInstruction: %v", instruction.Kind)
	}
}

func (e *seaportComputerEmulatorV2) setNewMaskApplicator(instruction maskInstruction) error {
	applicator, err := newMaskApplicatorV2(instruction.Mask)
	if err != nil {
		return errors.Wrap(err, "executing setNewMaskApplicator instruction")
	}
	e.maskApplicator = applicator
	return nil
}

func (e *seaportComputerEmulatorV2) writeToMemory(instruction writeInstruction) error {
	addresses := e.maskApplicator.Apply(instruction.Address)
	for _, address := range addresses {
		e.memory[address] = instruction.Value
	}
	return nil
}

func (e *seaportComputerEmulatorV2) MemorySum() uint64 {
	sum := uint64(0)
	for _, value := range e.memory {
		sum += value
	}
	return sum
}

type maskApplicatorV2 struct {
	representation string
	orBits         uint64
	floatMasks     []floatMask
}

type floatMask struct {
	andMask uint64
	orMask  uint64
}

func newMaskApplicatorV2(representation string) (*maskApplicatorV2, error) {
	if matched, err := regexp.MatchString(`[X10]{36}`, representation); !matched || err != nil {
		return nil, errors.New("representation does not fit the regular expression `[X10]{36}`")
	}

	orBits := uint64(0)
	floatMasks := []floatMask{}
	for i, bit := range representation {
		orBits <<= 1

		switch bit {
		case '1':
			orBits |= 1
		case '0':
			orBits |= 0
		case 'X':
			whichBit := (36 - 1 - i)
			andMask := ^uint64(0) ^ (1 << whichBit)
			orMask := uint64(0) | (1 << whichBit)
			floatMasks = append(floatMasks, floatMask{andMask: andMask, orMask: orMask})
		}
	}

	return &maskApplicatorV2{
		representation: representation,
		orBits:         orBits,
		floatMasks:     floatMasks,
	}, nil
}

func (a *maskApplicatorV2) Apply(in uint64) []uint64 {
	outs := []uint64{in | a.orBits}

	for _, mask := range a.floatMasks {
		newOuts := make([]uint64, 0, 2*len(outs))
		for _, previousOut := range outs {
			newOuts = append(newOuts, previousOut&mask.andMask)
			newOuts = append(newOuts, previousOut|mask.orMask)
		}
		outs = newOuts
	}

	return outs
}
