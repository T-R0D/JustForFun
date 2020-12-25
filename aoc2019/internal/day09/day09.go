package day09

import (
	"aoc2019/internal/intcode"
	"fmt"
)

const (
	kTEST_MODE         = 1
	kSENSOR_BOOST_MODE = 2
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	c := intcode.NewComputer()

	err := c.InputProgram(input)
	if err != nil {
		return nil, err
	}

	c.AddBatchInput(kTEST_MODE)

	state, err := c.RunProgram()
	if err != nil {
		return nil, err
	} else if state != intcode.STATE_RUN_COMPLETE {
		return nil, fmt.Errorf("Expected run to be complete, not in state %v", state)
	}

	output := c.GetBatchOutput()
	if len(output) != 1 {
		return nil, fmt.Errorf("Expected precisely 1 output, got %d", len(output))
	}

	return output[0], nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	c := intcode.NewComputer()

	err := c.InputProgram(input)
	if err != nil {
		return nil, err
	}

	c.AddBatchInput(kSENSOR_BOOST_MODE)

	state, err := c.RunProgram()
	if err != nil {
		return nil, err
	} else if state != intcode.STATE_RUN_COMPLETE {
		return nil, fmt.Errorf("Expected run to be complete, not in state %v", state)
	}

	output := c.GetBatchOutput()
	if len(output) != 1 {
		return nil, fmt.Errorf("Expected precisely 1 output, got %d", len(output))
	}

	return output[0], nil
}
