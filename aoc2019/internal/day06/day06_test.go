package day06

import (
	"reflect"
	"testing"
)

var INPUT = "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
var INPUT_AS_CENTER_TO_ORBITERS = centerToOrbiters{
	"COM": []string{"B"},
	"G":   []string{"H"},
	"B":   []string{"C", "G"},
	"C":   []string{"D"},
	"D":   []string{"E", "I"},
	"J":   []string{"K"},
	"K":   []string{"L"},
	"E":   []string{"F", "J"},
}
var INPUT2 = "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
var INPUT_AS_ORBIT_GRAPH = orbitGraph{
	"COM": []string{"B"},
	"B":   []string{"COM", "C", "G"},
	"C":   []string{"B", "D"},
	"D":   []string{"C", "E", "I"},
	"E":   []string{"D", "F", "J"},
	"F":   []string{"E"},
	"G":   []string{"B", "H"},
	"H":   []string{"G"},
	"I":   []string{"D", "SAN"},
	"J":   []string{"E", "K"},
	"K":   []string{"J", "L", "YOU"},
	"L":   []string{"K"},
	"YOU": []string{"K"},
	"SAN": []string{"I"},
}

func TestInputToMap(t *testing.T) {
	expected := INPUT_AS_CENTER_TO_ORBITERS
	result, err := inputToCenterToOrbiters(INPUT)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCountDirectAndIndirectOrbits(t *testing.T) {
	input := INPUT_AS_CENTER_TO_ORBITERS
	expected := 42
	result := countDirectAndIndirectOrbits(input)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestInputToOrbitGraph(t *testing.T) {
	// Note that this works because the input is ordered nicely and we can get
	// away with using arrays instead of sets.
	expected := INPUT_AS_ORBIT_GRAPH
	result, err := inputToOrbitGraph(INPUT2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v, got:\n%v", expected, result)
	}
}

func Test(t *testing.T) {
	input := INPUT_AS_ORBIT_GRAPH
	expected := 4
	result := countOrbitalTransfersToMatchOrbit(input, MY_SPACECRAFT, SANTA)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}
