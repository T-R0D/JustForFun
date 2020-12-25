// Improvement: Use the CYK algorithmn in part 1 too.
// Improvement: The CYK algorithm _REALLY_ needs some speedup somehow though.
// Possible Improvement: Can we optimize this CYK thing any more?
// Possible Improvement: I hear regexes work too according to Reddit.
//                       Warning: If you can't figure out lookahead, you may
//                       just have to try all possibilities for a reasonable number
//                       of recursions.

package day19

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2020/internal/equal"
	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	parsedInput, err := parseInput(input)
	if err != nil {
		return "", err
	}

	rule0Matches := generateRuleMatchingStrings(parsedInput.rules)[0]

	nMatches := 0
	for _, message := range parsedInput.messages {
		if _, ok := rule0Matches[message]; ok {
			nMatches++
		}
	}

	return strconv.Itoa(nMatches), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	parsedInput, err := parseInput(input)
	if err != nil {
		return "", err
	}

	parsedInput.rules[8] = ruleNode{
		kind:             ruleKindBranched,
		branchRulesLeft:  []int{42},
		branchRulesRight: []int{42, 8},
	}
	parsedInput.rules[11] = ruleNode{
		kind:             ruleKindBranched,
		branchRulesLeft:  []int{42, 31},
		branchRulesRight: []int{42, 11, 31},
	}

	checker := newMessageChecker(parsedInput.rules)

	nMatches := 0
	for i, message := range parsedInput.messages[0:] {
		fmt.Println("processing", i)

		if checker.grammarYieldsString(0, message) {
			nMatches++
		}
	}

	return strconv.Itoa(nMatches), nil
}

type parsedInputPair struct {
	rules    map[int]ruleNode
	messages []string
}

type ruleKind int

const (
	ruleKindTerminal      ruleKind = iota
	ruleKindComposite     ruleKind = iota
	ruleKindBranched      ruleKind = iota
	ruleKindShortBranched ruleKind = iota
)

type ruleNode struct {
	kind                ruleKind
	letter              string
	compositeChildRules []int
	branchRulesLeft     []int
	branchRulesRight    []int
}

func (r *ruleNode) Matches(composingRules []int) bool {
	switch r.kind {
	case ruleKindComposite:
		return equal.IntArrayEqual(r.compositeChildRules, composingRules)
	case ruleKindBranched:
		return equal.IntArrayEqual(r.branchRulesLeft, composingRules) ||
			equal.IntArrayEqual(r.branchRulesRight, composingRules)
	case ruleKindShortBranched:
		return equal.IntArrayEqual(r.branchRulesLeft, composingRules) ||
			equal.IntArrayEqual(r.branchRulesRight, composingRules)
	default:
		return false
	}
}

func (r *ruleNode) IsSingleMatch(composingRules []int) bool {
	if len(composingRules) != 1 {
		return false
	}

	return r.Matches(composingRules)
}

func parseInput(input string) (*parsedInputPair, error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		return nil, errors.Errorf("input had %d sections, not 2", len(parts))
	}

	ruleSet, err := parseRuleSet(parts[0])
	if err != nil {
		return nil, errors.Wrap(err, "parsing the rule section")
	}

	messages := strings.Split(parts[1], "\n")

	return &parsedInputPair{
		rules:    ruleSet,
		messages: messages,
	}, nil
}

func parseRuleSet(rulesSection string) (map[int]ruleNode, error) {
	letterRuleRegexp, err := regexp.Compile(`^"(\w)"$`)
	if err != nil {
		return nil, err
	}
	compositeRuleRegexp, err := regexp.Compile(`^(\d+\s?)+$`)
	if err != nil {
		return nil, err
	}
	branchedRuleRegexp, err := regexp.Compile(`^(\d+) (\d+) \| (\d+) (\d+)$`)
	if err != nil {
		return nil, err
	}
	shortBranchedRuleRegexp, err := regexp.Compile(`^(\d+) \| (\d+)$`)
	if err != nil {
		return nil, err
	}

	rules := map[int]ruleNode{}
	lines := strings.Split(rulesSection, "\n")
	for i, line := range lines {
		idAndDefinition := strings.Split(line, ": ")
		if len(idAndDefinition) != 2 {
			return nil, errors.Errorf("could not get 2 parts for id and definition on line %d, got %d parts", i, len(idAndDefinition))
		}

		id, err := strconv.Atoi(idAndDefinition[0])
		if err != nil {
			return nil, errors.Wrapf(err, "getting rule id on line %d", i)
		}

		definition := idAndDefinition[1]
		rule := ruleNode{}
		if match := letterRuleRegexp.FindStringSubmatch(definition); len(match) > 0 {
			rule.kind = ruleKindTerminal
			rule.letter = match[1]
		} else if match := shortBranchedRuleRegexp.FindStringSubmatch(definition); len(match) > 0 {
			rule.kind = ruleKindShortBranched
			childIDs := make([]int, len(match[1:]))
			for j, m := range match[1:] {
				childID, err := strconv.Atoi(m)
				if err != nil {
					return nil, errors.Wrapf(err, "parsing branched rule line %d, value %d", i, j)
				}
				childIDs[j] = childID
			}
			rule.branchRulesLeft = []int{childIDs[0]}
			rule.branchRulesRight = []int{childIDs[1]}
		} else if match := branchedRuleRegexp.FindStringSubmatch(definition); len(match) > 0 {
			rule.kind = ruleKindBranched
			childIDs := make([]int, len(match[1:]))
			for j, m := range match[1:] {
				childID, err := strconv.Atoi(m)
				if err != nil {
					return nil, errors.Wrapf(err, "parsing branched rule line %d, value %d", i, j)
				}
				childIDs[j] = childID
			}
			rule.branchRulesLeft = []int{childIDs[0], childIDs[1]}
			rule.branchRulesRight = []int{childIDs[2], childIDs[3]}
		} else if match := compositeRuleRegexp.FindStringSubmatch(definition); len(match) > 0 {
			rule.kind = ruleKindComposite
			childIDStrs := strings.Split(definition, " ")
			childIDs := make([]int, len(childIDStrs))
			for j, childIDStr := range childIDStrs {
				childID, err := strconv.Atoi(childIDStr)
				if err != nil {
					return nil, errors.Wrapf(err, "parsing composite rule line %d, value %d", i, j)
				}
				childIDs[j] = childID
			}
			rule.compositeChildRules = childIDs
		} else {
			return nil, errors.Errorf("line %d didn't have a valid rule definition (%s)", i, definition)
		}

		rules[id] = rule
	}

	return rules, nil
}

type messageChecker struct {
	cykTable             [][][]int
	reverseGrammarLookup map[string][]int
	grammar              map[int]ruleNode
	matchesForSubstring  map[string][]int
}

func newMessageChecker(rules map[int]ruleNode) *messageChecker {
	reverseLookup := createRuleReverseLookup(rules)

	return &messageChecker{
		cykTable:             nil,
		reverseGrammarLookup: reverseLookup,
		grammar:              rules,
		matchesForSubstring:  map[string][]int{},
	}
}

func createRuleReverseLookup(rules map[int]ruleNode) map[string][]int {
	reverseLookup := map[string][]int{}
	for id, rule := range rules {
		switch rule.kind {
		case ruleKindTerminal:
			key := rule.letter
			if matches, ok := reverseLookup[key]; ok {
				reverseLookup[key] = append(matches, id)
			} else {
				reverseLookup[key] = []int{id}
			}
		case ruleKindComposite:
			key := fmt.Sprintf("%v", rule.compositeChildRules)
			if matches, ok := reverseLookup[key]; ok {
				reverseLookup[key] = append(matches, id)
			} else {
				reverseLookup[key] = []int{id}
			}
		case ruleKindBranched, ruleKindShortBranched:
			key := fmt.Sprintf("%v", rule.branchRulesLeft)
			if matches, ok := reverseLookup[key]; ok {
				reverseLookup[key] = append(matches, id)
			} else {
				reverseLookup[key] = []int{id}
			}

			key = fmt.Sprintf("%v", rule.branchRulesRight)
			if matches, ok := reverseLookup[key]; ok {
				reverseLookup[key] = append(matches, id)
			} else {
				reverseLookup[key] = []int{id}
			}
		}
	}
	return reverseLookup
}

func (c *messageChecker) grammarYieldsString(startRule int, candidateStr string) bool {
	c.cykTable = make([][][]int, len(candidateStr))
	for i := 0; i < len(candidateStr); i++ {
		c.cykTable[i] = make([][]int, len(candidateStr)-i)
	}

	// Prime the table using the terminals.
	for i, r := range candidateStr {
		letter := string(r)
		if matchingRules, ok := c.reverseGrammarLookup[letter]; ok {
			c.cykTable[0][i] = matchingRules
			c.addInIdentityRules(0, i)
		} else {
			return false
		}
	}

	// Build the rest of the table up.
	for substringLength := 2; substringLength <= len(candidateStr); substringLength++ {
		l := substringLength - 1

		for s := 0; s <= len(candidateStr)-substringLength; s++ {
			substring := candidateStr[s:s+substringLength]
			if matchingRules, ok := c.matchesForSubstring[substring]; ok {
				c.cykTable[l][s] = matchingRules
				continue
			}

			var matchingRules []int

			// Try and match 2 partition rules.
			for partitionSize := 1; partitionSize <= substringLength-1; partitionSize++ {
				secondPartitionSize := substringLength - partitionSize

				var candidateMatches [][]int
				for _, match := range c.cykTable[partitionSize-1][s] {
					for _, match2 := range c.cykTable[secondPartitionSize-1][s+partitionSize] {
						candidateMatches = append(candidateMatches, []int{match, match2})
					}
				}

				for _, candidateMatch := range candidateMatches {
					key := c.keyFromMatch(candidateMatch)
					if matches, ok := c.reverseGrammarLookup[key]; ok {
						for _, match := range matches {
							matchingRules = append(matchingRules, match)
						}
					}
				}
			}

			// Try and match 3 partition rules.
			for firstPartitionSize := 1; firstPartitionSize <= substringLength-2; firstPartitionSize++ {
				for secondPartitionSize := 1; secondPartitionSize <= substringLength-(firstPartitionSize+1); secondPartitionSize++ {
					firstAndSecondPartitionSize := firstPartitionSize + secondPartitionSize
					thirdPartitionSize := substringLength - firstAndSecondPartitionSize

					var candidateMatches [][]int
					for _, match := range c.cykTable[firstPartitionSize-1][s] {
						for _, match2 := range c.cykTable[secondPartitionSize-1][s+firstPartitionSize] {
							for _, match3 := range c.cykTable[thirdPartitionSize-1][s+firstAndSecondPartitionSize] {
								candidateMatches = append(candidateMatches, []int{match, match2, match3})
							}
						}
					}

					for _, candidateMatch := range candidateMatches {
						key := c.keyFromMatch(candidateMatch)
						if matches, ok := c.reverseGrammarLookup[key]; ok {
							for _, match := range matches {
								matchingRules = append(matchingRules, match)
							}
						}
					}
				}
			}

			sort.Ints(matchingRules)

			c.cykTable[l][s] = matchingRules

			c.addInIdentityRules(l, s)

			c.matchesForSubstring[substring] = c.cykTable[l][s]
		}
	}

	entireStringRuleMatches := c.cykTable[len(candidateStr)-1][0]
	return arrayContains(entireStringRuleMatches, startRule)
}

func (c *messageChecker) addInIdentityRules(l int, s int) {
	matches := c.cykTable[l][s]

	needToCheckForIdentityRules := true
	for needToCheckForIdentityRules {
		needToCheckForIdentityRules = false

		for _, match := range matches {
			key := c.keyFromMatch([]int{match})
			if reverseLookupMatches, ok := c.reverseGrammarLookup[key]; ok {
				for _, reverseLookupMatch := range reverseLookupMatches {
					if arrayContains(matches, reverseLookupMatch) {
						continue
					}

					matches = append(matches, reverseLookupMatch)
					needToCheckForIdentityRules = true
				}
			}
		}
	}

	c.cykTable[l][s] = matches
}

func (c *messageChecker) keyFromMatch(match []int) string {
	return fmt.Sprintf("%v", match)
}

func (c *messageChecker) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("Total Rules: %d\n", len(c.reverseGrammarLookup)))

	for l := len(c.cykTable) - 1; l >= 0; l-- {
		tableSlice := c.cykTable[l]

		builder.WriteString(fmt.Sprintf("%2d:", l))

		for s := 0; s < len(tableSlice); s++ {
			str := fmt.Sprintf("%v                ", tableSlice[s])[:18]
			builder.WriteString(str)
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

func tableToString(table [][][]int) string {
	builder := strings.Builder{}

	for l := len(table) - 1; l >= 0; l-- {
		tableSlice := table[l]

		builder.WriteString(fmt.Sprintf("%2d:", l))

		for s := 0; s < len(tableSlice); s++ {
			str := fmt.Sprintf("%v        ", tableSlice[s])[:9]
			builder.WriteString(str)
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func arrayContains(arr []int, candidate int) bool {
	for _, x := range arr {
		if x == candidate {
			return true
		}
	}
	return false
}

// An attempt at Part 1 that just generates all of the strings.
type ruleTraversalState struct {
	ruleID    int
	rule      ruleNode
	nextChild int
}

func generateRuleMatchingStrings(rules map[int]ruleNode) []map[string]struct{} {
	ruleMatches := make([]map[string]struct{}, len(rules))

	frontier := []*ruleTraversalState{
		{
			ruleID:    0,
			rule:      rules[0],
			nextChild: 0,
		},
	}

	for len(frontier) > 0 {

		if len(frontier) > 1000 {
			panic("bailing cuz search got too deep")
		}

		currentState := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		switch currentState.rule.kind {
		case ruleKindTerminal:
			ruleMatches[currentState.ruleID] = map[string]struct{}{currentState.rule.letter: {}}
		case ruleKindComposite:
			if currentState.nextChild < len(currentState.rule.compositeChildRules) {
				nextChildID := currentState.rule.compositeChildRules[currentState.nextChild]
				currentState.nextChild++

				frontier = append(frontier, currentState)

				if ruleMatches[nextChildID] == nil {
					frontier = append(frontier, &ruleTraversalState{
						ruleID:    nextChildID,
						rule:      rules[nextChildID],
						nextChild: 0,
					})
				}
			} else {
				childRuleIDs := currentState.rule.compositeChildRules
				if len(childRuleIDs) == 1 {
					ruleMatches[currentState.ruleID] = ruleMatches[childRuleIDs[0]]
				} else if len(childRuleIDs) == 2 {
					candidates := map[string]struct{}{}
					for firstHalf := range ruleMatches[childRuleIDs[0]] {
						for secondHalf := range ruleMatches[childRuleIDs[1]] {
							candidates[firstHalf+secondHalf] = struct{}{}
						}
					}
					ruleMatches[currentState.ruleID] = candidates
				} else if len(childRuleIDs) == 3 {
					candidates := map[string]struct{}{}
					for beginning := range ruleMatches[childRuleIDs[0]] {
						for middle := range ruleMatches[childRuleIDs[1]] {
							for end := range ruleMatches[childRuleIDs[2]] {
								candidates[beginning+middle+end] = struct{}{}
							}
						}
					}
					ruleMatches[currentState.ruleID] = candidates
				} else {
					panic(" too many composite children")
				}
			}
		case ruleKindShortBranched:
			if currentState.nextChild < 2 {
				nextChildID := -1
				if currentState.nextChild < 1 {
					nextChildID = currentState.rule.branchRulesLeft[currentState.nextChild]
				} else {
					nextChildID = currentState.rule.branchRulesRight[currentState.nextChild-1]
				}

				currentState.nextChild++
				frontier = append(frontier, currentState)

				if ruleMatches[nextChildID] == nil {
					frontier = append(frontier, &ruleTraversalState{
						ruleID:    nextChildID,
						rule:      rules[nextChildID],
						nextChild: 0,
					})
				}
			} else {
				candidates := map[string]struct{}{}
				leftChildRuleIDs := currentState.rule.branchRulesLeft
				for candidate := range ruleMatches[leftChildRuleIDs[0]] {
					candidates[candidate] = struct{}{}
				}
				rightChildRuleIDs := currentState.rule.branchRulesRight
				for candidate := range ruleMatches[rightChildRuleIDs[0]] {
					candidates[candidate] = struct{}{}
				}
				ruleMatches[currentState.ruleID] = candidates
			}
		case ruleKindBranched:
			if currentState.nextChild < 4 {
				nextChildID := -1
				if currentState.nextChild < 2 {
					nextChildID = currentState.rule.branchRulesLeft[currentState.nextChild]
				} else {
					nextChildID = currentState.rule.branchRulesRight[currentState.nextChild-2]
				}

				currentState.nextChild++
				frontier = append(frontier, currentState)

				if ruleMatches[nextChildID] == nil {
					frontier = append(frontier, &ruleTraversalState{
						ruleID:    nextChildID,
						rule:      rules[nextChildID],
						nextChild: 0,
					})
				}
			} else {
				candidates := map[string]struct{}{}
				leftChildRuleIDs := currentState.rule.branchRulesLeft
				for firstHalf := range ruleMatches[leftChildRuleIDs[0]] {
					for secondHalf := range ruleMatches[leftChildRuleIDs[1]] {
						candidates[firstHalf+secondHalf] = struct{}{}
					}
				}
				rightChildRuleIDs := currentState.rule.branchRulesRight
				for firstHalf := range ruleMatches[rightChildRuleIDs[0]] {
					for secondHalf := range ruleMatches[rightChildRuleIDs[1]] {
						candidates[firstHalf+secondHalf] = struct{}{}
					}
				}
				ruleMatches[currentState.ruleID] = candidates
			}
		}
	}

	return ruleMatches
}
