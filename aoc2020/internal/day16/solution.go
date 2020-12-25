// Improvement: Clean up this code. It needs it, big time. Make it testable,
//              add unit tests.
// Possible Improvement: On reddit, people are talking about needing a 64 bit
//                       int for part 2's answer. I guess I got lucky since
//                       Go's int is 64 bit on a 64 bit system.

package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	information, err := parseAllInformation(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing the input")
	}

	ticketScanningErrorSum := 0
	for _, ticket := range information.nearbyTickets {
		value := scanTicketForErrorValue(information.rules, ticket)
		if value != noErrorValue {
			ticketScanningErrorSum += value
		}
	}

	return strconv.Itoa(ticketScanningErrorSum), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	information, err := parseAllInformation(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing the input")
	}

	decoder := newRuleDecoder(information.rules)

	decoder.AnalyzeTickets(information.nearbyTickets)

	if !decoder.SolutionFound() {
		return "", errors.New("decoder failed to find a solution")
	}

	departureFieldIndices := decoder.IndicesOfDepartureValues()
	sumOfDepartureValues := 1
	for _, index := range departureFieldIndices {
		sumOfDepartureValues *= information.myTicket.values[index]
	}

	return strconv.Itoa(sumOfDepartureValues), nil
}

const (
	noErrorValue = -1
)

type collectedInformation struct {
	rules         map[string]ruleDef
	myTicket      trainTicket
	nearbyTickets []trainTicket
}

type ruleDef struct {
	ranges []ruleRange
}

type ruleRange struct {
	lowerBound int
	upperBound int
}

type trainTicket struct {
	values []int
}

func parseAllInformation(input string) (*collectedInformation, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 3 {
		return nil, errors.Errorf("wrong number of sections in input: %d", len(sections))
	}

	rules, err := parseRules(sections[0])
	if err != nil {
		return nil, err
	}

	myTicket, err := parseMyTicket(sections[1])
	if err != nil {
		return nil, err
	}

	nearbyTickets, err := parseNearbyTickets(sections[2])
	if err != nil {
		return nil, err
	}

	return &collectedInformation{
		rules:         rules,
		myTicket:      *myTicket,
		nearbyTickets: nearbyTickets,
	}, nil
}

func parseRules(rulesSection string) (map[string]ruleDef, error) {
	ruleRE, err := regexp.Compile(`^([\w ]+): (0|\d+)-(0|\d+) or (0|\d+)-(0|\d+)$`)
	if err != nil {
		return nil, errors.Wrap(err, "compiling the rule regex")
	}

	rules := map[string]ruleDef{}

	lines := strings.Split(rulesSection, "\n")
	for i, line := range lines {
		match := ruleRE.FindStringSubmatch(line)
		if len(match) != 6 {
			return nil, errors.Errorf("line %d failed to parse", i)
		}

		field := match[1]
		ruleRangeStrs := []string{match[2], match[3], match[4], match[5]}
		ruleRangeBounds := make([]int, len(ruleRangeStrs))
		for j, rangeStr := range ruleRangeStrs {
			value, err := strconv.Atoi(rangeStr)
			if err != nil {
				return nil, errors.Wrapf(err, "line %d, range value %d failed to parse", i, j)
			}
			ruleRangeBounds[j] = value
		}

		rules[field] = ruleDef{
			ranges: []ruleRange{
				{
					lowerBound: ruleRangeBounds[0],
					upperBound: ruleRangeBounds[1],
				},
				{
					lowerBound: ruleRangeBounds[2],
					upperBound: ruleRangeBounds[3],
				},
			},
		}
	}

	return rules, nil
}

func parseMyTicket(myTicketSection string) (*trainTicket, error) {
	lines := strings.Split(myTicketSection, "\n")
	if len(lines) != 2 {
		return nil, errors.Errorf("myTicket section had %d lines, not 2", len(lines))
	}

	values, err := parseTicketValues(lines[1])
	if err != nil {
		return nil, errors.Wrap(err, "parsing the myTicket section")
	}

	return &trainTicket{values: values}, nil
}

func parseNearbyTickets(nearbyTicketsSection string) ([]trainTicket, error) {
	lines := strings.Split(nearbyTicketsSection, "\n")

	tickets := make([]trainTicket, len(lines)-1)
	for i, line := range lines[1:] {
		values, err := parseTicketValues(line)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing nearby ticket %d", i)
		}
		tickets[i].values = values
	}

	return tickets, nil
}

func parseTicketValues(line string) ([]int, error) {
	valueStrs := strings.Split(line, ",")
	values := make([]int, len(valueStrs))
	for i, valueStr := range valueStrs {
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing value %d", i)
		}
		values[i] = value
	}
	return values, nil
}

func scanTicketForErrorValue(rules map[string]ruleDef, ticket trainTicket) int {
	for _, value := range ticket.values {
		valueMatchedARule := false
		for _, rule := range rules {
			valueFellInARange := false
			for _, ruleRange := range rule.ranges {
				if ruleRange.lowerBound <= value && value <= ruleRange.upperBound {
					valueFellInARange = true
					break
				}
			}

			if valueFellInARange {
				valueMatchedARule = true
				break
			}
		}

		if !valueMatchedARule {
			return value
		}
	}

	return noErrorValue
}

type ruleDecoder struct {
	rules             map[string]ruleDef
	rulePossibilities []map[string]struct{}
}

func newRuleDecoder(rules map[string]ruleDef) *ruleDecoder {
	rulePossibilities := make([]map[string]struct{}, len(rules))
	for i := range rulePossibilities {
		possibilities := map[string]struct{}{}
		for fieldName := range rules {
			possibilities[fieldName] = struct{}{}
		}
		rulePossibilities[i] = possibilities
	}

	return &ruleDecoder{
		rules:             rules,
		rulePossibilities: rulePossibilities,
	}
}

func (d *ruleDecoder) AnalyzeTickets(tickets []trainTicket) error {
	for _, ticket := range tickets {
		d.AnalyzeTicket(ticket)
	}

	return nil
}

func (d *ruleDecoder) AnalyzeTicket(ticket trainTicket) bool {
	if scanTicketForErrorValue(d.rules, ticket) != noErrorValue {
		return false
	}

	for i, value := range ticket.values {
		nonMatchingRules := []string{}
		valueMatchedField := false
		for fieldName := range d.rulePossibilities[i] {
			if d.valueMatchesField(fieldName, value) {
				valueMatchedField = true
			} else {
				nonMatchingRules = append(nonMatchingRules, fieldName)
			}
		}

		if valueMatchedField {
			d.eliminateValues(i, nonMatchingRules)
			d.eliminateAcrossSlots()
		} else {
			panic("value didn't match any available fields")
		}
	}

	return true
}

func (d *ruleDecoder) valueMatchesField(fieldName string, value int) bool {
	rule, ok := d.rules[fieldName]
	if !ok {
		panic("invalid field name looked up")
	}

	return d.valueMatchesRule(rule, value)
}

func (d *ruleDecoder) valueMatchesRule(rule ruleDef, value int) bool {
	for _, ruleRange := range rule.ranges {
		if ruleRange.lowerBound <= value && value <= ruleRange.upperBound {
			return true
		}
	}
	return false
}

func (d *ruleDecoder) eliminateValues(listPosition int, impossibleFields []string) {
	possibilities := d.rulePossibilities[listPosition]
	for _, impossibleField := range impossibleFields {
		delete(possibilities, impossibleField)
	}

	if len(possibilities) == 1 {
		for fieldName := range possibilities {
			for i := range d.rulePossibilities {
				if i == listPosition {
					continue
				}
				delete(d.rulePossibilities[i], fieldName)
			}
		}
	}
}

func (d *ruleDecoder) eliminateAcrossSlots() {
	roundsOfEliminationToBeMadeStill := 0
	for listPosition, possibilities := range d.rulePossibilities {
		if len(possibilities) == 1 {
			for fieldName := range possibilities {
				for i, otherPossibilities := range d.rulePossibilities {
					if i == listPosition {
						continue
					}

					previousLen := len(otherPossibilities)
					delete(otherPossibilities, fieldName)
					currentLen := len(otherPossibilities)
					if previousLen > 1 && currentLen == 1 {
						roundsOfEliminationToBeMadeStill++
					}
				}
			}
		}
	}

	for i := 0; i < roundsOfEliminationToBeMadeStill; i++ {
		d.eliminateAcrossSlots()
	}
}

func (d *ruleDecoder) SolutionFound() bool {
	for _, possibilities := range d.rulePossibilities {
		if len(possibilities) != 1 {
			return false
		}
	}
	return true
}

func (d *ruleDecoder) IndicesOfDepartureValues() []int {
	indices := []int{}
	for i, possibilities := range d.rulePossibilities {
		for fieldName := range possibilities {
			if strings.HasPrefix(fieldName, "departure") {
				indices = append(indices, i)
			}
		}
	}
	return indices
}

func (d *ruleDecoder) PossibilitiesString() string {
	builder := strings.Builder{}
	for i, possibilities := range d.rulePossibilities {
		builder.WriteString(fmt.Sprintf("%d: [", i))
		for fieldName := range possibilities {
			builder.WriteString(fmt.Sprintf("%s ", fieldName))
		}
		builder.WriteString("]\n")
	}
	return builder.String()
}
