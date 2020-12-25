package day02

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Solver struct{}

func (s *Solver) Part1(input string) (string, error) {
	entries := strings.Split(input, "\n")
	nValidPasswords := 0

	for _, entry := range entries {
		if valid, err := hasValidPasswordByFrequency(entry); err != nil {
			return "", err
		} else if valid {
			nValidPasswords += 1
		}
	}

	return strconv.Itoa(nValidPasswords), nil
}

func (s *Solver) Part2(input string) (string, error) {
	entries := strings.Split(input, "\n")
	nValidPasswords := 0

	for _, entry := range entries {
		if valid, err := hasValidPasswordByPlacement(entry); err != nil {
			return "", err
		} else if valid {
			nValidPasswords += 1
		}
	}

	return strconv.Itoa(nValidPasswords), nil
}

func hasValidPasswordByFrequency(entryLine string) (bool, error) {
	entry, err := parseDBEntry(entryLine)
	if err != nil {
		return false, errors.Wrap(err, "unable to tell if entry contains valid password")
	}

	targetLetterCount := 0
	for _, r := range entry.password {
		if r == entry.requirement.letter {
			targetLetterCount += 1
		}
	}

	return entry.requirement.min <= targetLetterCount && targetLetterCount <= entry.requirement.max, nil
}

// !442
func hasValidPasswordByPlacement(entryLine string) (bool, error) {
	entry, err := parseDBEntry(entryLine)
	if err != nil {
		return false, errors.Wrap(err, "unable to tell if entry contains valid password")
	}

	password := entry.password
	requirement := entry.requirement
	passwordLen := len(password)
	firstPosition := requirement.min - 1
	lastPosition := requirement.max - 1

	if !isInRange(firstPosition, 0, passwordLen) {
		return false, errors.New("requirement didn't have a valid first placement")
	}

	firstPositionLetter := rune(password[firstPosition])

	var secondPositionLetter rune
	if isInRange(lastPosition, 0, passwordLen) {
		secondPositionLetter = rune(password[lastPosition])
	}

	hasTheRightLetterInAPosition := firstPositionLetter == requirement.letter || secondPositionLetter == requirement.letter
	hasPreciselyOneLetterInTheRightPosition := hasTheRightLetterInAPosition && firstPositionLetter != secondPositionLetter

	return hasPreciselyOneLetterInTheRightPosition, nil
}

type passwordDBEntry struct {
	requirement passwordRequirement
	password    string
}

type passwordRequirement struct {
	letter rune
	max    int
	min    int
}

func parseDBEntry(line string) (passwordDBEntry, error) {
	entry := passwordDBEntry{}

	parts := strings.Split(line, ": ")

	if len(parts) != 2 {
		return entry, errors.New("line didn't have 2 parts after splitting on ': '")
	}

	entry.password = parts[1]

	requirement, err := parsePasswordRequirement(parts[0])
	if err != nil {
		return entry, err
	}

	entry.requirement = requirement

	return entry, nil

}

func parsePasswordRequirement(req string) (passwordRequirement, error) {
	requirement := passwordRequirement{}

	parts := strings.Split(req, " ")

	if len(parts) != 2 {
		return requirement, errors.New("password requirement didn't have 2 parts after splitting on ' '")
	}

	requirement.letter = rune(parts[1][0])

	frequencyParts := strings.Split(parts[0], "-")

	if len(frequencyParts) != 2 {
		return requirement, errors.New("range specification in password requirement didn't have 2 parts after splitting on '-'")
	}

	min, err := strconv.Atoi(frequencyParts[0])
	if err != nil {
		return requirement, errors.Wrap(err, "min candidate could not be parsed")
	}
	max, err := strconv.Atoi(frequencyParts[1])
	if err != nil {
		return requirement, errors.Wrap(err, "max candidate could not be parsed")
	}

	requirement.min = min
	requirement.max = max

	return requirement, nil
}

func isInRange(candidate int, start int, end int) bool {
	return start <= candidate && candidate < end
}
