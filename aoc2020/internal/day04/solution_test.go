package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestPart1(t *testing.T) {
	// Arrange.
	s := Solver{}

	// Act.
	nValidishPassports, err := s.Part1(sampleInput)

	// Assert.
	assert.NoError(t, err)
	assert.Equal(t, "2", nValidishPassports)
}

func TestPassportHasAllFields(t *testing.T) {
	testCases := []struct{
		name string
		rawPassportData string
		hasAllFields bool
	}{
		{
			name: "complete passport has all fields",
			rawPassportData: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			hasAllFields: true,
		},
		{
			name: "passport missing non optional field is invalid",
			rawPassportData: "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			hasAllFields: false,
		},
		{
			name: "passport missing optional field is still valid",
			rawPassportData: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			hasAllFields: true,
		},
		{
			name: "passport missing optional and non optional fields is invalid",
			rawPassportData: "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
			hasAllFields: false,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T){
			// Arrange.

			// Act.
			actualHasAllFields, err := passportHasAllFields(tc.rawPassportData, []passportField{ passportFieldCountryID })

			// Assert.
			assert.NoError(t, err)
			assert.Equal(t, tc.hasAllFields, actualHasAllFields)
		})
	}
}
