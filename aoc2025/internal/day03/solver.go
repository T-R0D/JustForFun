package day03

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	banks, err := parseBatteryBanks(input)
	if err != nil {
		return "", err
	}

	joltageSum := 0
	for _, bank := range banks {
		joltageSum += findMaxJoltage(bank, 2)
	}

	return strconv.Itoa(joltageSum), nil
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	banks, err := parseBatteryBanks(input)
	if err != nil {
		return "", err
	}

	joltageSum := 0
	for _, bank := range banks {
		joltageSum += findMaxJoltage(bank, 12)
	}

	return strconv.Itoa(joltageSum), nil
}

type batteryBank []int

func parseBatteryBanks(input string) ([]batteryBank, error) {
	bankLines := strings.Split(input, "\n")
	banks := make([]batteryBank, 0, len(bankLines))
	for _, line := range bankLines {
		individualJoltages := make([]int, 0, len(line))
		for _, digit := range strings.Split(line, "") {
			joltage, err := strconv.Atoi(digit)
			if err != nil {
				return []batteryBank{}, err
			}

			individualJoltages = append(individualJoltages, joltage)
		}

		banks = append(banks, individualJoltages)
	}

	return banks, nil
}

func findMaxJoltage(bank batteryBank, batteriesToEnable int) int {
	accumulatedJoltage := 0
	lastBatteryUsedIndex := -1
	for b := range batteriesToEnable {
		accumulatedJoltage *= 10

		lastBatteryUsedIndex += 1
		maxDigit := bank[lastBatteryUsedIndex]
		for i := lastBatteryUsedIndex + 1; i < (len(bank) - (batteriesToEnable - (b + 1))); i += 1 {
			if bank[i] > maxDigit {
				maxDigit = bank[i]
				lastBatteryUsedIndex = i
			}
		}

		accumulatedJoltage += maxDigit
	}

	return accumulatedJoltage
}
