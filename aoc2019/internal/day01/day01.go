package day01

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	partMasses := strings.Split(input, "\n")
	requiredFuel := 0
	for _, mass := range partMasses {
		imass, err := strconv.Atoi(mass)
		if err != nil {
			return nil, err
		}
		requiredFuel += requiredFuelForPart(imass)
	}

	return requiredFuel, nil
}

func requiredFuelForPart(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	partMasses := strings.Split(input, "\n")
	requiredFuel := 0
	for _, mass := range partMasses {
		imass, err := strconv.Atoi(mass)
		if err != nil {
			return nil, err
		}
		requiredFuel += actualRequiredFuelForPart(imass)
	}

	return requiredFuel, nil
}

func actualRequiredFuelForPart(mass int) int {
	requiredFuel := requiredFuelForPart(mass)
	if 0 == requiredFuel {
		return 0
	}
	return requiredFuel + actualRequiredFuelForPart(requiredFuel)
}
