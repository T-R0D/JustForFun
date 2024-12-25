package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2024/v2/internal/queue"
)

type Solver struct{}

func (s *Solver) SolvePartOne(input string) (string, error) {
	schematic, err := parseGateSchematic(input)
	if err != nil {
		return "", err
	}

	result, err := findNumberComputedByCircuit(schematic)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(result, 10), nil
}

func (s *Solver) SolvePartTwo(input string) (string, error) {
	schematic, err := parseGateSchematic(input)
	if err != nil {
		return "", err
	}

	swappedWires, err := determineSwappedWires(schematic)
	if err != nil {
		return "", err
	}

	slices.Sort(swappedWires)
	return strings.Join(swappedWires, ","), nil
}

type gateSchematic struct {
	InitialWireValues map[string]int
	Gates             []logicGate
	GatesByOutput     map[string]logicGate
}

type logicGate struct {
	Op  rune
	InA string
	InB string
	Out string
}

func parseGateSchematic(input string) (*gateSchematic, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return nil, fmt.Errorf("there was not precisely 2 sections in the input")
	}

	initialWireValues, err := parseInitialWireValues(sections[0])
	if err != nil {
		return nil, err
	}

	gates, err := parseLogicGates(sections[1])
	if err != nil {
		return nil, err
	}

	outputToGates := make(map[string]logicGate, len(gates))
	for _, gate := range gates {
		outputToGates[gate.Out] = gate
	}

	return &gateSchematic{
		InitialWireValues: initialWireValues,
		Gates:             gates,
		GatesByOutput:     outputToGates,
	}, nil
}

func parseInitialWireValues(initialValueList string) (map[string]int, error) {
	lines := strings.Split(initialValueList, "\n")
	initialValues := make(map[string]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ": ")
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return map[string]int{}, fmt.Errorf("initial values line %d did not have a number; %w", i, err)
		}
		initialValues[parts[0]] = value
	}

	return initialValues, nil
}

func parseLogicGates(gateList string) ([]logicGate, error) {
	const (
		andSeparator = " AND "
		orSeparator  = " OR "
		xorSeparator = " XOR "
	)

	lines := strings.Split(gateList, "\n")
	gates := make([]logicGate, 0, len(lines))
	for _, line := range lines {
		op := '?'
		opString := ""
		if strings.Contains(line, orSeparator) {
			op = '|'
			opString = orSeparator
		} else if strings.Contains(line, andSeparator) {
			op = '&'
			opString = andSeparator
		} else if strings.Contains(line, xorSeparator) {
			op = '^'
			opString = xorSeparator
		}

		operationAndOut := strings.Split(line, " -> ")
		inputs := strings.Split(operationAndOut[0], opString)
		if len(inputs) != 2 {
			return []logicGate{}, fmt.Errorf("%s (from '%s') could not be split into 2 parts", operationAndOut[0], line)
		}

		gate := logicGate{
			Op:  op,
			InA: inputs[0],
			InB: inputs[1],
			Out: operationAndOut[1],
		}
		gates = append(gates, gate)
	}

	return gates, nil
}

func findNumberComputedByCircuit(schematic *gateSchematic) (int64, error) {
	wireValues := initializeWires(schematic, schematic.InitialWireValues)

	result, err := determineValueComputedByCircuit(wireValues)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func determineSwappedWires(schematic *gateSchematic) ([]string, error) {
	outputBitSize, err := determineOutputBitSize(schematic)
	if err != nil {
		return []string{}, err
	}

	faultyOutputPaths, err := findFaultyOutputPaths(schematic, outputBitSize)
	if err != nil {
		return []string{}, nil
	}
	fmt.Printf("faultyOutputPaths: %v\n", len(faultyOutputPaths))

	return []string{}, nil
}

func initializeWires(schematic *gateSchematic, initialWireValues map[string]int) map[string]int {
	type searchState struct {
		GateId int
		Depth  int
	}

	cmpSearchState := func(a searchState, b searchState) bool {
		gateA := schematic.Gates[a.GateId]
		gateB := schematic.Gates[b.GateId]

		if a.Depth != b.Depth {
			return a.Depth > b.Depth
		}

		if gateA.Out == gateB.InA || gateA.Out == gateB.InB {
			return true
		} else if gateB.Out == gateA.InA || gateB.Out == gateA.InB {
			return false
		}

		return false
	}

	gatesInSearch := map[int]struct{}{}

	frontier := queue.NewPriority(cmpSearchState)
	for out := range schematic.GatesByOutput {
		if !strings.HasPrefix(out, "z") {
			continue
		}

	FIND_GATE_ID:
		for id, gate := range schematic.Gates {
			if gate.Out == out {
				frontier.Push(searchState{GateId: id, Depth: 0})

				gatesInSearch[id] = struct{}{}
				break FIND_GATE_ID
			}
		}
	}

	wireValues := map[string]int{}
	for wire, value := range initialWireValues {
		wireValues[wire] = value
	}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		gate := schematic.Gates[state.GateId]

		inAValue, inADone := wireValues[gate.InA]
		inBValue, inBDone := wireValues[gate.InB]

		if inADone && inBDone {
			outValue := 0
			switch gate.Op {
			case '|':
				outValue = inAValue | inBValue
			case '&':
				outValue = inAValue & inBValue
			case '^':
				outValue = inAValue ^ inBValue
			}

			wireValues[gate.Out] = outValue
			continue
		}

		if !inADone {
			nextGate := schematic.GatesByOutput[gate.InA]
			nextGateId := -1
			for id, x := range schematic.Gates {
				if nextGate == x {
					nextGateId = id
				}
			}

			frontier.Push(searchState{GateId: nextGateId, Depth: state.Depth + 1})

		}

		if !inBDone {
			nextGate := schematic.GatesByOutput[gate.InB]
			nextGateId := -1
			for id, x := range schematic.Gates {
				if nextGate == x {
					nextGateId = id
				}
			}

			frontier.Push(searchState{GateId: nextGateId, Depth: state.Depth + 1})
		}

		frontier.Push(searchState{GateId: state.GateId, Depth: state.Depth})
	}

	return wireValues
}

func determineValueComputedByCircuit(wireValues map[string]int) (int64, error) {
	result := int64(0)
	for wire, value := range wireValues {
		if !strings.HasPrefix(wire, "z") {
			continue
		}

		bitPosition, err := strconv.Atoi(strings.ReplaceAll(wire, "z", ""))
		if err != nil {
			return 0, err
		}

		result |= int64(value << bitPosition)
	}

	return result, nil
}

func determineOutputBitSize(schematic *gateSchematic) (int64, error) {
	maxBitPosition := int64(0)
	for wire := range schematic.GatesByOutput {
		if !strings.HasPrefix(wire, "z") {
			continue
		}

		bitPosition, err := strconv.ParseInt(strings.ReplaceAll(wire, "z", ""), 10, 64)
		if err != nil {
			return 0, err
		}

		if bitPosition > maxBitPosition {
			maxBitPosition = bitPosition
		}
	}

	return maxBitPosition + 1, nil
}

func findFaultyOutputPaths(schematic *gateSchematic, outputBitSize int64) ([]string, error) {
	faultyOutputPathWires := map[string]struct{}{}

	for i := range outputBitSize - 1 {
		x := (int64(1) << outputBitSize) - 1
		x = x & ^(1 << i)
		y := (int64(1) << outputBitSize) - 1
		z := x + y

		initialWireValues := map[string]int{}
		for j := range outputBitSize - 1 {

			xWire := fmt.Sprintf("x%02d", j)
			xValue := (x >> j) & 0x01
			initialWireValues[xWire] = int(xValue)

			yWire := fmt.Sprintf("y%02d", j)
			yValue := (y >> j) & 0x01
			initialWireValues[yWire] = int(yValue)
		}

		wireValues := initializeWires(schematic, initialWireValues)

		zActual, err := determineValueComputedByCircuit(wireValues)
		if err != nil {
			return []string{}, err
		}

		if z == zActual {
			continue
		}

		zString := fmt.Sprintf("%b", z)
		zActString := fmt.Sprintf("%b", zActual)
		fmt.Printf("z   : % [2]*[1]s\n", zString, outputBitSize+1)
		fmt.Printf("zAct: % [2]*[1]s\n\n", zActString, outputBitSize+1)

		for j := range outputBitSize {
			a, b := z&(1<<j), zActual&(1<<j)

			if a != b {

				faultyOutputPathWires[fmt.Sprintf("z%02d", j)] = struct{}{}
			}
		}
	}

	faultyOutputs := make([]string, 0, len(faultyOutputPathWires))
	for wire := range faultyOutputPathWires {
		faultyOutputs = append(faultyOutputs, wire)
	}
	slices.Sort(faultyOutputs)

	return faultyOutputs, nil
}
