package day07

import (
	"aoc2019/internal/intcode"
	"fmt"
	"math"
)

const (
	kN_AMPLIFIERS = 5
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	maxOutput := 0
	phaseSettingPermutations := generatePermutations([]int{0, 1, 2, 3, 4})
	for _, settingPermutation := range phaseSettingPermutations {
		amplifierOutput := 0
		var err error
		for _, phaseSetting := range settingPermutation {
			amplifierOutput, err = runAmplifierControlProgram(input, phaseSetting, amplifierOutput)
			if err != nil {
				return nil, err
			}
		}
		maxOutput = int(math.Max(float64(maxOutput), float64(amplifierOutput)))
	}
	return maxOutput, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	maxOutput := 0
	phaseSettingPermutations := generatePermutations([]int{5, 6, 7, 8, 9})
	for _, settingPermutation := range phaseSettingPermutations {
		computers, err := initComputers(input, settingPermutation)
		if err != nil {
			return nil, err
		}

		runHalted := false
		amplifierOutput := 0
		expectingHalt := false
		for !runHalted {
			for i, c := range computers {
				state := c.GetState()
				if state == intcode.STATE_AWAITING_INPUT {
					c.Input(amplifierOutput)
					state, err = c.RunProgram()
					if state != intcode.STATE_AWAITING_OUTPUT {
						if err == nil {
							err = fmt.Errorf("Expected program to produce output after being given an input.")
						}
						return nil, err
					}
					amplifierOutput = c.CollectOutput()
					// The state will get picked up on the next loop iteration
					// for this computer.
					_, err = c.RunProgram()
					if err != nil {
						return nil, err
					}
				} else if i == 0 && state == intcode.STATE_RUN_COMPLETE {
					expectingHalt = true
					continue
				} else if expectingHalt && state == intcode.STATE_RUN_COMPLETE {
					if i == len(computers)-1 {
						runHalted = true
					}
				} else {
					return nil, fmt.Errorf("Unexpected program state from computer %d: %v", i, state)
				}
			}
		}

		maxOutput = int(math.Max(float64(maxOutput), float64(amplifierOutput)))
	}
	return maxOutput, nil
}

func runAmplifierControlProgram(prog string, phaseSetting, prevAmplifierOutput int) (output int, err error) {
	c := intcode.NewComputer()
	c.InputProgram(prog)
	c.AddBatchInputs([]int{phaseSetting, prevAmplifierOutput})
	state, err := c.RunProgram()
	if err != nil && state == intcode.STATE_RUN_COMPLETE {
		return
	} else if err == nil {
		err = fmt.Errorf("Expected a complete program run - that didn't happen.")
		return
	}
	outputs := c.GetBatchOutput()
	if len(outputs) != 1 {
		err = fmt.Errorf("Unexpected number of outputs; expected only one, got %v", outputs)
	}
	output = outputs[0]
	return
}

func generatePermutations(elements []int) [][]int {
	r := [][]int{}

	b := make([]int, len(elements))
	generatePermutationHelper(&r, elements, b, 0)

	return r
}

func generatePermutationHelper(r *[][]int, elements []int, b []int, i int) {
	if len(elements) == 0 {
		rb := make([]int, len(b))
		copy(rb, b)
		*r = append(*r, rb)
		return
	}

	for j := 0; j < len(elements); j++ {
		e := elements[j]
		b[i] = e
		remainingElements := produceReducedSet(elements, j)

		generatePermutationHelper(r, remainingElements, b, i+1)
	}
}

func produceReducedSet(elements []int, i int) []int {
	r := make([]int, 0, len(elements)-1)
	for j, v := range elements {
		if j == i {
			continue
		}
		r = append(r, v)
	}
	return r
}

func initComputers(prog string, settingPermutation []int) ([]*intcode.Computer, error) {
	computers := make([]*intcode.Computer, kN_AMPLIFIERS)
	for i := 0; i < kN_AMPLIFIERS; i++ {
		c := intcode.NewComputer()
		c.SetInterruptibleMode()
		c.InputProgram(prog)

		state, err := c.RunProgram()
		if state != intcode.STATE_AWAITING_INPUT || err != nil {
			if err == nil {
				err = fmt.Errorf("Expected computer to be awaiting initial input (the phase setting)")
			}
			return nil, err
		}

		c.Input(settingPermutation[i])

		state, err = c.RunProgram()
		if state != intcode.STATE_AWAITING_INPUT || err != nil {
			if err == nil {
				err = fmt.Errorf("Expected computer to be awaiting input (the first 'output')")
			}
			return nil, err
		}

		computers[i] = c
	}
	return computers, nil
}
