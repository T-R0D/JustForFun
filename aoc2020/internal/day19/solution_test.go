package day19

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ruleInput1 = `0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`

var parsedRuleInput1 = map[int]ruleNode{
	0: {
		kind:                ruleKindComposite,
		compositeChildRules: []int{1, 2},
	},
	1: {
		kind:   ruleKindTerminal,
		letter: "a",
	},
	2: {
		kind:             ruleKindBranched,
		branchRulesLeft:  []int{1, 3},
		branchRulesRight: []int{3, 1},
	},
	3: {
		kind:   ruleKindTerminal,
		letter: "b",
	},
}

const ruleInput2 = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`

var parsedRuleInput2 = map[int]ruleNode{
	0: {
		kind:                ruleKindComposite,
		compositeChildRules: []int{4, 1, 5},
	},
	1: {
		kind:             ruleKindBranched,
		branchRulesLeft:  []int{2, 3},
		branchRulesRight: []int{3, 2},
	},
	2: {
		kind:             ruleKindBranched,
		branchRulesLeft:  []int{4, 4},
		branchRulesRight: []int{5, 5},
	},
	3: {
		kind:             ruleKindBranched,
		branchRulesLeft:  []int{4, 5},
		branchRulesRight: []int{5, 4},
	},
	4: {
		kind:   ruleKindTerminal,
		letter: "a",
	},
	5: {
		kind:   ruleKindTerminal,
		letter: "b",
	},
}

func TestParseRules(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedOutput map[int]ruleNode
	}{
		{
			name:           "rule set 1 parses correctly",
			input:          ruleInput1,
			expectedOutput: parsedRuleInput1,
		},
		{
			name:           "rule set 2 parses correctly",
			input:          ruleInput2,
			expectedOutput: parsedRuleInput2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput, err := parseRuleSet(tc.input)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedOutput, actualOutput)
		})
	}
}

func TestGenerateRuleMatchingStrings(t *testing.T) {
	testCases := []struct {
		name                    string
		rules                   map[int]ruleNode
		expectedMatchingStrings map[string]struct{}
	}{
		{
			name:  "rule set 1 has 2 matching strings",
			rules: parsedRuleInput1,
			expectedMatchingStrings: map[string]struct{}{
				"aab": {},
				"aba": {},
			},
		},
		{
			name:  "rule set 2 has 8 matching strings",
			rules: parsedRuleInput2,
			expectedMatchingStrings: map[string]struct{}{
				"aaaabb": {},
				"aaabab": {},
				"abbabb": {},
				"abbbab": {},
				"aabaab": {},
				"aabbbb": {},
				"abaaab": {},
				"ababbb": {},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualMatchingStrings := generateRuleMatchingStrings(tc.rules)

			assert.Equal(t, tc.expectedMatchingStrings, actualMatchingStrings)
		})
	}
}

func TestGrammarYieldsString(t *testing.T) {
	testCases := []struct{
		grammarName string
			grammarDefinition map[int]ruleNode
			candidateString string
	}{
		{
			grammarName:       "1",
			grammarDefinition: parsedRuleInput1,
			candidateString:   "aab",
		},
		{
			grammarName:       "1",
			grammarDefinition: parsedRuleInput1,
			candidateString:   "aba",
		},
		{
			grammarName:       "2",
			grammarDefinition: parsedRuleInput2,
			candidateString:   "aaaabb",
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("grammar %s yields string %s", tc.grammarName, tc.candidateString)
		t.Run(name, func(t *testing.T){
			checker := newMessageChecker(tc.grammarDefinition)

			yields := checker.grammarYieldsString(0, tc.candidateString)

			assert.True(t, yields)
		})
	}
}

func TestGrammarDoesNotYieldString(t *testing.T) {
	testCases := []struct{
		grammarName string
			grammarDefinition map[int]ruleNode
			candidateString string
	}{
		{
			grammarName:       "1",
			grammarDefinition: parsedRuleInput1,
			candidateString:   "aaa",
		},
		{
			grammarName:       "1",
			grammarDefinition: parsedRuleInput1,
			candidateString:   "aaa",
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("grammar %s does not yield string %s", tc.grammarName, tc.candidateString)
		t.Run(name, func(t *testing.T){
			checker := newMessageChecker(tc.grammarDefinition)

			yields := checker.grammarYieldsString(0, tc.candidateString)

			assert.False(t, yields)
		})
	}
}

const myExample = `0: 5 1 6
1: 2
2: 3 4 | 4 3
3: 5 5 | 6 6
4: 5 6 | 6 5
5: "a"
6: "b"

aaaabb`

func TestMyExample(t *testing.T) {
	parsedInput, err := parseInput(myExample)
	assert.NoError(t, err)

	checker := newMessageChecker(parsedInput.rules)

	result := checker.grammarYieldsString(0, "aaaabb")

	assert.True(t, result)
}

func TestMediumProblem(t *testing.T) {
	parsedInput, err := parseInput(mediumProblem)
	assert.NoError(t, err)

	checker := newMessageChecker(parsedInput.rules)

	result := checker.grammarYieldsString(0, "bbabbbbaabaabba")

	assert.True(t, result)
}

const mediumProblem = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`
