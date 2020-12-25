package day07

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestBuildBagCapabilitiesLookup(t *testing.T) {
	lookup := buildBagCapabilitiesLookup(sampleInput)

	expectedLookup := bagCapabilitiesLookup{
		"light red":    []bagSlots{{color: "bright white", count: 1}, {color: "muted yellow", count: 2}},
		"dark orange":  []bagSlots{{color: "bright white", count: 3}, {color: "muted yellow", count: 4}},
		"bright white": []bagSlots{{color: "shiny gold", count: 1}},
		"muted yellow":  []bagSlots{{color: "shiny gold", count: 2}, {color: "faded blue", count: 9}},
		"shiny gold":  []bagSlots{{color: "dark olive", count: 1}, {color: "vibrant plum", count: 2}},
		"dark olive":  []bagSlots{{color: "faded blue", count: 3}, {color: "dotted black", count: 4}},
		"vibrant plum":  []bagSlots{{color: "faded blue", count: 5}, {color: "dotted black", count: 6}},
		"faded blue":  []bagSlots{},
		"dotted black":  []bagSlots{},
	}

	assert.Equal(t, expectedLookup, lookup)
}

func TestCanHoldShinyGold(t *testing.T) {
	testCases := []struct{
		startingColor string
		canHoldShinyGold bool
	}{
		{
			startingColor: "light red",
			canHoldShinyGold: true,
		},
		{
			startingColor: "dark orange",
			canHoldShinyGold: true,
		},
		{
			startingColor: "bright white",
			canHoldShinyGold: true,
		},
		{
			startingColor: "muted yellow",
			canHoldShinyGold: true,
		},
		{
			startingColor: "shiny gold",
			canHoldShinyGold: false,
		},
		{
			startingColor: "dark olive",
			canHoldShinyGold: false,
		},
		{
			startingColor: "vibrant plum",
			canHoldShinyGold: false,
		},
		{
			startingColor: "faded blue",
			canHoldShinyGold: false,
		},
		{
			startingColor: "dotted black",
			canHoldShinyGold: false,
		},
	}
	
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s can hold shiny gold: %v", tc.startingColor, tc.canHoldShinyGold), func(t *testing.T){
			lookup := buildBagCapabilitiesLookup(sampleInput)

			actualCanHoldShinyGold := canHoldShinyGold(tc.startingColor, lookup)

			assert.Equal(t, tc.canHoldShinyGold, actualCanHoldShinyGold)
		})
	}
}
