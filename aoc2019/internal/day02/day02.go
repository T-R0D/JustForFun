package day02

import (
	"aoc2019/internal/intcode"
	"fmt"
)

const (
	NOUN_ADDR = 1
	VERB_ADDR = 2

	PROGRAM_NOUN = 12
	ALARM_VERB   = 02

	DESIRED_OUTPUT = 19690720
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	c := intcode.NewComputer()

	err := c.InputProgram(input)
	if err != nil {
		return nil, err
	}

	err = c.UpdateProgram(map[int]int{
		NOUN_ADDR: PROGRAM_NOUN,
		VERB_ADDR: ALARM_VERB,
	})
	if err != nil {
		return nil, err
	}

	_, err = c.RunProgram()
	if err != nil {
		return nil, err
	}

	return c.GetProgramResult()
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			c := intcode.NewComputer()

			err := c.InputProgram(input)
			if err != nil {
				return nil, err
			}

			err = c.UpdateProgram(map[int]int{
				NOUN_ADDR: noun,
				VERB_ADDR: verb,
			})
			if err != nil {
				return nil, err
			}

			_, err = c.RunProgram()
			if err != nil {
				return nil, err
			}

			r, err := c.GetProgramResult()
			if err != nil {
				return nil, err
			}

			if DESIRED_OUTPUT == r {
				return (100 * noun) + verb, nil
			}
		}
	}
	return nil, fmt.Errorf("Solution not found")
}
