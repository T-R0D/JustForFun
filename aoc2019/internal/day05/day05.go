package day05

import (
	"aoc2019/internal/intcode"
	"fmt"
)

const (
	AC_UNIT_ID          = 1
	THERMAL_RADIATOR_ID = 5
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	c := intcode.NewComputer()
	err := c.InputProgram(input)
	if err != nil {
		return nil, err
	}
	c.AddBatchInput(AC_UNIT_ID)
	_, err = c.RunProgram()
	if err != nil {
		return nil, err
	}
	output := c.GetBatchOutput()
	if errorsReported, err := outputReportsErrors(output); errorsReported || err != nil {
		return nil, fmt.Errorf("Program's output did not indicate error free run")
	}

	return output[len(output)-1], nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	c := intcode.NewComputer()
	err := c.InputProgram(input)
	if err != nil {
		return nil, err
	}
	c.AddBatchInput(THERMAL_RADIATOR_ID)
	_, err = c.RunProgram()
	if err != nil {
		return nil, err
	}
	output := c.GetBatchOutput()
	if len(output) != 1 {
		return nil, fmt.Errorf("Program did not output precisely 1 value - got %v", len(output))
	}

	return output[0], nil
}

func outputReportsErrors(output []int) (bool, error) {
	if len(output) < 2 {
		return false, fmt.Errorf("Output too short to be capable of reporting errors")
	}

	errorReports := output[:len(output)-1]
	for _, report := range errorReports {
		if report != 0 {
			return true, nil
		}
	}

	return false, nil
}
