package day04

import (
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/validation"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	passportsAsTextLines := cleanBatchFileContents(input)

	nValidishPassports := 0
	optionalFields := []passportField{passportFieldCountryID}
	for i, passportLine := range passportsAsTextLines {
		if valid, err := passportHasAllFields(passportLine, optionalFields); err != nil {
			return "", errors.Wrapf(err, "the %dth passport had an error", i)
		} else if valid {
			nValidishPassports++
		}
	}

	return strconv.Itoa(nValidishPassports), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	passportsAsTextLines := cleanBatchFileContents(input)

	nValidishPassports := 0
	optionalFields := []passportField{passportFieldCountryID}
	for i, passportLine := range passportsAsTextLines {
		if valid, err := passportIsValid(passportLine, optionalFields); err != nil {
			return "", errors.Wrapf(err, "the %dth passport had an error", i)
		} else if valid {
			nValidishPassports++
		}
	}

	return strconv.Itoa(nValidishPassports), nil
}

type passportField string

const (
	passportFieldBirthYear      passportField = "byr"
	passportFieldIssueYear      passportField = "iyr"
	passportFieldExpirationYear passportField = "eyr"
	passportFieldHeight         passportField = "hgt"
	passportFieldHairColor      passportField = "hcl"
	passportFieldEyeColor       passportField = "ecl"
	passportFieldPassportID     passportField = "pid"
	passportFieldCountryID      passportField = "cid"
)

func passportFieldFromString(candidate string) (passportField, error) {
	candidateAsPassportField := passportField(candidate)
	switch candidateAsPassportField {
	case
		passportFieldBirthYear,
		passportFieldIssueYear,
		passportFieldExpirationYear,
		passportFieldHeight,
		passportFieldHairColor,
		passportFieldEyeColor,
		passportFieldPassportID,
		passportFieldCountryID:
		return candidateAsPassportField, nil
	default:
		return candidateAsPassportField, errors.Errorf("%s is not a member of the passportField enum", candidate)
	}
}

func cleanBatchFileContents(input string) []string {
	formattedInputBuilder := strings.Builder{}
	newLineJustSeen := false
	for _, r := range input {
		switch r {
		case '\n':
			if newLineJustSeen {
				formattedInputBuilder.WriteRune(r)
				newLineJustSeen = false
			} else {
				newLineJustSeen = true
			}
		default:
			if newLineJustSeen {
				formattedInputBuilder.WriteRune(' ')
				newLineJustSeen = false
			}
			formattedInputBuilder.WriteRune(r)
		}
	}
	formattedInput := formattedInputBuilder.String()
	return strings.Split(formattedInput, "\n")
}

func mapFromPassportLine(line string) (map[passportField]string, error) {
	keyValuePairs := strings.Split(line, " ")

	passportMap := map[passportField]string{}
	for _, kvPair := range keyValuePairs {
		kvParts := strings.Split(kvPair, ":")
		key, err := passportFieldFromString(kvParts[0])
		if err != nil {
			return passportMap, errors.Wrap(err, "invalid key encountered")
		}
		passportMap[key] = kvParts[1]
	}
	return passportMap, nil
}

func passportHasAllFields(rawData string, optionalFields []passportField) (bool, error) {
	fieldsNotYetSeen := map[passportField]struct{}{
		passportFieldBirthYear:      {},
		passportFieldIssueYear:      {},
		passportFieldExpirationYear: {},
		passportFieldHeight:         {},
		passportFieldHairColor:      {},
		passportFieldEyeColor:       {},
		passportFieldPassportID:     {},
		passportFieldCountryID:      {},
	}

	for _, field := range optionalFields {
		delete(fieldsNotYetSeen, field)
	}

	passportMap, err := mapFromPassportLine(rawData)
	if err != nil {
		return false, err
	}

	for field := range passportMap {
		delete(fieldsNotYetSeen, field)
	}

	if len(fieldsNotYetSeen) > 0 {
		return false, err
	}
	return true, err
}

func passportIsValid(rawData string, optionalFields []passportField) (bool, error) {
	fieldsNotYetSeen := map[passportField]struct{}{
		passportFieldBirthYear:      {},
		passportFieldIssueYear:      {},
		passportFieldExpirationYear: {},
		passportFieldHeight:         {},
		passportFieldHairColor:      {},
		passportFieldEyeColor:       {},
		passportFieldPassportID:     {},
		passportFieldCountryID:      {},
	}

	for _, field := range optionalFields {
		delete(fieldsNotYetSeen, field)
	}

	passportMap, err := mapFromPassportLine(rawData)
	if err != nil {
		return false, err
	}

	for field, rawValue := range passportMap {
		validator, ok := passportFieldValidators[field]
		if !ok {
			continue
		}


		if fieldIsValid := validator(rawValue); !fieldIsValid {
			return false, nil
		}

		delete(fieldsNotYetSeen, field)
	}

	if len(fieldsNotYetSeen) > 0 {
		return false, err
	}
	return true, err
}

var passportFieldValidators = map[passportField]func(string) bool{
	passportFieldBirthYear:      birthYearIsValid,
	passportFieldIssueYear:      issueYearIsValid,
	passportFieldExpirationYear: expirationYearIsValid,
	passportFieldHeight:         heightIsValid,
	passportFieldHairColor:      hairColorIsValid,
	passportFieldEyeColor:       eyeColorIsValid,
	passportFieldPassportID:     passportIDIsValid,
}

func birthYearIsValid(rawBirthYear string) bool {
	return passportValueInAcceptableRange(rawBirthYear, 1920, 2003)
}

func issueYearIsValid(rawIssueYear string) bool {
	return passportValueInAcceptableRange(rawIssueYear, 2010, 2021)
}

func expirationYearIsValid(rawExpirationYear string) bool {
	return passportValueInAcceptableRange(rawExpirationYear, 2020, 2031)
}

func heightIsValid(rawHeight string) bool {
	if strings.HasSuffix(rawHeight, "cm") {
		rawHeightNumber := strings.Replace(rawHeight, "cm", "", 1)
		return passportValueInAcceptableRange(rawHeightNumber, 150, 194)
	} else if strings.HasSuffix(rawHeight, "in") {
		rawHeightNumber := strings.Replace(rawHeight, "in", "", 1)
		return passportValueInAcceptableRange(rawHeightNumber, 59, 77)
	}

	return false
}

func hairColorIsValid(rawHairColor string) bool {
	if len(rawHairColor) != 7 {
		return false
	}

	if rawHairColor[0] != '#' {
		return false
	}

	for _, r := range rawHairColor[1:] {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f':
			continue
		default:
			return false
		}
	}

	return true
}

func eyeColorIsValid(rawEyeColor string) bool {
	switch rawEyeColor {
	case
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth":
		return true
	default:
		return false
	}
}

func passportIDIsValid(rawPassportID string)  bool {
	if len(rawPassportID) != 9 {
		return false
	}

	for _, r := range rawPassportID {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			return false
		}
	}

	return true
}

func passportValueInAcceptableRange(rawValue string, lowerBound int, upperBound int) bool {
	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return false
	}
	return validation.ValueIsInRange(value, lowerBound, upperBound)
}
