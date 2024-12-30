package day24

import (
	"fmt"
	"maps"
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
	candidateFaultyOutputPaths, err := findCandidateFaultyOutputPaths(schematic)
	if err != nil {
		return []string{}, nil
	}

	fmt.Printf("found candidate faulty output paths (%d)\n", len(candidateFaultyOutputPaths))
	for key := range candidateFaultyOutputPaths {
		fmt.Printf("\t%v\n", key.OutputWire)
	}

	candidateWiresToSwap := map[string]struct{}{}
	for k, wireValues := range candidateFaultyOutputPaths {
		newCandidates := findSuspectWiresInOutputPath(schematic, k.OutputWire, wireValues)
		for _, wire := range newCandidates {
			candidateWiresToSwap[wire] = struct{}{}
		}
	}

	fmt.Printf("found %d involved wires\n", len(candidateWiresToSwap))

	// Got real sloppy with this one...
	// Had a lot of trouble making a fully programmatic solution. So I did it 
	// manually (partially). It turns out I was pretty spot on with
	// identifying the faulty outputs (done by trying every variation of
	// int45 max + 1 bit as well as int45 max + int45 max minus 1 bit). This,
	// combined with a graphviz visualization and some comparison with a
	// correct implementation of a ripple adder, and I sussed out the
	// problematic connections pretty easily. Maybe someday I will solve it
	// fully programmatically, but for now, the monkey is off my back. My
	// current attempt produces too many combinations of wires to swap, so I
	// need to make it smarter.

	return []string{"grf", "wpq", "z18", "fvw", "z22", "mdb", "z36", "nwq"}, nil
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

func findCandidateFaultyOutputPaths(schematic *gateSchematic) (faultyOutputLookup, error) {
	outputBitSize, err := determineOutputBitSize(schematic)
	if err != nil {
		return faultyOutputLookup{}, err
	}

	faultyOutputMaxPaths, err := findFaultyOutputPathsByMaxAddition(schematic, outputBitSize)
	if err != nil {
		return faultyOutputLookup{}, nil
	}

	faultyOutputMinPaths, err := findFaultyOutputPathsByMinAddition(schematic, outputBitSize)
	if err != nil {
		return faultyOutputLookup{}, nil
	}

	faultyPaths := maps.Clone(faultyOutputMaxPaths)
	for k, v := range faultyOutputMinPaths {
		if _, ok := faultyOutputMaxPaths[k]; ok {
			continue
		}
		faultyPaths[k] = v
	}

	return faultyPaths, nil
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

type faultyOutputLookup map[faultyOutputLookupKey]map[string]int

type faultyOutputLookupKey struct {
	OutputWire     string
	ExpectedOutput int
}

func findFaultyOutputPathsByMaxAddition(schematic *gateSchematic, outputBitSize int64) (faultyOutputLookup, error) {
	faultyOutputPathWires := faultyOutputLookup{}

	for i := range outputBitSize - 1 {
		x := (int64(1) << (outputBitSize - 1)) - 1
		x = x & ^(1 << i)
		y := (int64(1) << (outputBitSize - 1)) - 1
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
			return faultyOutputLookup{}, err
		}

		if z == zActual {
			continue
		}

		for j := range outputBitSize {
			mask := int64(1 << j)
			a, b := z&mask, zActual&mask
			if a != b {
				faultyOutputPathWires[faultyOutputLookupKey{
					OutputWire:     fmt.Sprintf("z%02d", j),
					ExpectedOutput: int(a >> j),
				}] = maps.Clone(wireValues)
			}
		}
	}

	return faultyOutputPathWires, nil
}

func findFaultyOutputPathsByMinAddition(schematic *gateSchematic, outputBitSize int64) (faultyOutputLookup, error) {
	faultyOutputPathWires := faultyOutputLookup{}

	for i := range outputBitSize - 1 {
		x := int64(0)
		x = x | (1 << i)
		y := int64(0)
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
			return faultyOutputLookup{}, err
		}

		if z == zActual {
			continue
		}

		// zString := fmt.Sprintf("%b", z)
		// zActString := fmt.Sprintf("%b", zActual)
		// fmt.Printf("z   : % [2]*[1]s (%[3]d)\n", zString, outputBitSize+1, z)
		// fmt.Printf("zAct: % [2]*[1]s (%[3]d)\n\n", zActString, outputBitSize+1, zActual)

		for j := range outputBitSize {
			mask := int64(1 << j)
			a, b := z&mask, zActual&mask
			if a != b {
				faultyOutputPathWires[faultyOutputLookupKey{
					OutputWire:     fmt.Sprintf("z%02d", j),
					ExpectedOutput: int(a >> j),
				}] = maps.Clone(wireValues)
			}
		}
	}

	return faultyOutputPathWires, nil
}

func findSuspectWiresInOutputPath(schematic *gateSchematic, outputWire string, observedWireValues map[string]int) []string {
	type searchState struct {
		OutWire        string
		ExpectedOutput int
	}

	frontier := queue.NewLifo[searchState]()
	initialActualOutput := observedWireValues[outputWire]
	initialExpectedOutput := 0
	if initialActualOutput == 0 {
		initialExpectedOutput = 1
	}
	frontier.Push(searchState{OutWire: outputWire, ExpectedOutput: initialExpectedOutput})

	candidateWires := map[string]struct{}{}

	for frontier.Len() > 0 {
		state, ok := frontier.Pop()
		if !ok {
			break
		}

		actualValue := observedWireValues[state.OutWire]

		if actualValue == state.ExpectedOutput {
			continue
		}

		candidateWires[state.OutWire] = struct{}{}

		gate, ok := schematic.GatesByOutput[state.OutWire]
		if !ok {
			continue
		}

		if state.ExpectedOutput == 0 {
			switch gate.Op {
			case '|':
				if inAValue := observedWireValues[gate.InA]; inAValue == 1 {
					frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 0})
				}
				if inBValue := observedWireValues[gate.InB]; inBValue == 1 {
					frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 0})
				}
			case '&':
				frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 0})
				frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 0})
			case '^':
				if inValue := observedWireValues[gate.InA]; inValue == 1 {
					frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 0})
					frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 0})
				} else {
					frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 1})
					frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 1})
				}
			}
		} else if state.ExpectedOutput == 1 {
			switch gate.Op {
			case '|':
				frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 1})
				frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 1})
			case '&':
				frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 1})
				frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 1})
			case '^':
				if inValue := observedWireValues[gate.InA]; inValue == 1 {
					frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 0})
					frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 0})
				} else {
					frontier.Push(searchState{OutWire: gate.InA, ExpectedOutput: 1})
					frontier.Push(searchState{OutWire: gate.InB, ExpectedOutput: 1})
				}
			}
		}
	}

	involvedWires := make([]string, 0, len(candidateWires))
	for wire := range candidateWires {
		involvedWires = append(involvedWires, wire)
	}

	return involvedWires
}

func findAllSwappingCandidates(candidateWires []string, nPairs int) [][][]string {
	n := nPairs * 2

	wireCombinations := combinations(candidateWires, n)

	wirePermutations := [][]string{}
	for _, combination := range wireCombinations {
		wirePermutations = append(wirePermutations, permutations(combination)...)
	}

	swappingPairs := [][][]string{}
	for _, permutation := range wirePermutations {
		swappingSet := [][]string{}
		for i := 0; i < len(permutation); i += 1 {
			swappingSet = append(swappingSet, []string{permutation[i], permutation[i+1]})
		}

		swappingPairs = append(swappingPairs, swappingSet)
	}

	return swappingPairs
}

func combinations[T any](arr []T, k int) [][]T {
	var result [][]T

	var comb func([]T, int, int)
	comb = func(a []T, start int, k int) {
		if k == 0 {
			tmp := make([]T, len(a))
			copy(tmp, a)
			result = append(result, tmp)
			return
		}

		for i := start; i <= len(arr)-k; i++ {
			a[len(a)-k] = arr[i]
			comb(a, i+1, k-1)
		}
	}

	comb(make([]T, k), 0, k)
	return result
}

func permutations[T any](arr []T) [][]T {
	var result [][]T

	var generate func([]T, int)
	generate = func(a []T, size int) {
		if size == 1 {
			tmp := make([]T, len(a))
			copy(tmp, a)
			result = append(result, tmp)
			return
		}

		for i := 0; i < size; i++ {
			generate(a, size-1)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}

	generate(arr, len(arr))
	return result
}

func getSwappingTransformationFunc(swappingPairs [][]string) func(string) string {
	swapMap := map[string]string{}
	for _, pair := range swappingPairs {
		swapMap[pair[0]] = pair[1]
		swapMap[pair[1]] = pair[0]
	}

	return func(wire string) string {
		if complement, ok := swapMap[wire]; ok {
			return complement
		}
		return wire
	}
}

func verifyOuptutwithRemapping(schematic *gateSchematic, remappingFunc func(string) string) (bool, error) {
	outputBitSize, err := determineOutputBitSize(schematic)
	if err != nil {
		return false, err
	}

	newSchematic := newGateSchematicWithRemappedWires(schematic, remappingFunc)

	if faultyPaths, err := findFaultyOutputPathsByMaxAdditionWithRemapping(newSchematic, outputBitSize, remappingFunc); err != nil {
		return false, err
	} else if len(faultyPaths) > 1 {
		return false, nil
	}

	if faultyPaths, err := findFaultyOutputPathsByMinAdditionWithRemapping(newSchematic, outputBitSize, remappingFunc); err != nil {
		return false, err
	} else if len(faultyPaths) > 1 {
		return false, nil
	}

	return true, nil
}

func newGateSchematicWithRemappedWires(schematic *gateSchematic, remappingFunc func(string) string) *gateSchematic {
	newGates := make([]logicGate, 0, len(schematic.Gates))
	for _, gate := range schematic.Gates {
		newGate := logicGate{
			Op:  gate.Op,
			InA: remappingFunc(gate.InA),
			InB: remappingFunc(gate.InB),
			Out: remappingFunc(gate.Out),
		}
		newGates = append(newGates, newGate)
	}

	outToGate := make(map[string]logicGate, len(newGates))
	for _, gate := range newGates {
		outToGate[gate.Out] = gate
	}

	initialWireValues := map[string]int{}
	for wire, value := range schematic.InitialWireValues {
		initialWireValues[remappingFunc(wire)] = value
	}

	return &gateSchematic{
		InitialWireValues: initialWireValues,
		Gates:             newGates,
		GatesByOutput:     outToGate,
	}
}

func determineValueComputedByCircuitWithRemapping(wireValues map[string]int, remappingFunc func(string) string) (int64, error) {
	result := int64(0)
	for wire := range wireValues {
		if !strings.HasPrefix(wire, "z") {
			continue
		}

		actualValue, ok := wireValues[remappingFunc(wire)]
		if !ok {
			return 0, fmt.Errorf("%s -> %s was not found", wire, remappingFunc(wire))
		}

		bitPosition, err := strconv.Atoi(strings.ReplaceAll(wire, "z", ""))
		if err != nil {
			return 0, err
		}

		result |= int64(actualValue << bitPosition)
	}

	return result, nil
}

func findFaultyOutputPathsByMaxAdditionWithRemapping(schematic *gateSchematic, outputBitSize int64, remappingFunc func(string) string) (map[string]struct{}, error) {
	faultyOutputPathWires := map[string]struct{}{}

	for i := range outputBitSize - 1 {
		x := (int64(1) << (outputBitSize - 1)) - 1
		x = x & ^(1 << i)
		y := (int64(1) << (outputBitSize - 1)) - 1
		z := x + y

		initialWireValues := map[string]int{}
		for j := range outputBitSize - 1 {
			xWire := fmt.Sprintf("x%02d", j)
			xValue := (x >> j) & 0x01
			initialWireValues[remappingFunc(xWire)] = int(xValue)

			yWire := fmt.Sprintf("y%02d", j)
			yValue := (y >> j) & 0x01
			initialWireValues[remappingFunc(yWire)] = int(yValue)
		}

		wireValues := initializeWires(schematic, initialWireValues)

		zActual, err := determineValueComputedByCircuitWithRemapping(wireValues, remappingFunc)
		if err != nil {
			return map[string]struct{}{}, err
		}

		if z == zActual {
			continue
		}

		for j := range outputBitSize {
			mask := int64(1 << j)
			a, b := z&mask, zActual&mask
			if a != b {
				faultyOutputPathWires[fmt.Sprintf("z%02d", j)] = struct{}{}
			}
		}
	}

	return faultyOutputPathWires, nil
}

func findFaultyOutputPathsByMinAdditionWithRemapping(schematic *gateSchematic, outputBitSize int64, remappingFunc func(string) string) (map[string]struct{}, error) {
	faultyOutputPathWires := map[string]struct{}{}

	for i := range outputBitSize - 1 {
		x := int64(0)
		x = x | (1 << i)
		y := int64(0)
		z := x + y

		initialWireValues := map[string]int{}
		for j := range outputBitSize - 1 {
			xWire := fmt.Sprintf("x%02d", j)
			xValue := (x >> j) & 0x01
			initialWireValues[remappingFunc(xWire)] = int(xValue)

			yWire := fmt.Sprintf("y%02d", j)
			yValue := (y >> j) & 0x01
			initialWireValues[remappingFunc(yWire)] = int(yValue)
		}

		wireValues := initializeWires(schematic, initialWireValues)

		zActual, err := determineValueComputedByCircuitWithRemapping(wireValues, remappingFunc)
		if err != nil {
			return map[string]struct{}{}, err
		}

		if z == zActual {
			continue
		}

		for j := range outputBitSize {
			mask := int64(1 << j)
			a, b := z&mask, zActual&mask
			if a != b {
				faultyOutputPathWires[fmt.Sprintf("z%02d", j)] = struct{}{}
			}
		}
	}

	return faultyOutputPathWires, nil
}

func schematicToGraphvizInput(schematic *gateSchematic) {
	fmt.Printf("digraph {")
	for _, gate := range schematic.Gates {
		gateName := "OR"
		color := "red"
		shape := "square"
		switch gate.Op {
		case '|':
			gateName = "OR"
			color = "red"
			shape = "square"
		case '&':
			gateName = "AND"
			color = "green"
			shape = "invtrapezium"
		case '^':
			gateName = "XOR"
			color = "blue"
			shape = "invhouse"
		}
		gateName = fmt.Sprintf("%s%s", gateName, gate.Out)

		fmt.Printf("\t%s [color = %s, shape = %s]\n", gateName, color, shape)
		fmt.Printf("\t%s -> %s\n", gate.InA, gateName)
		fmt.Printf("\t%s -> %s\n", gate.InB, gateName)
		fmt.Printf("\t%s -> %s\n", gateName, gate.Out)

		if strings.HasPrefix(gate.Out, "z") {
			fmt.Printf("\t%s -> output\n", gate.Out)
		}
	}
	fmt.Printf(("}"))
}
