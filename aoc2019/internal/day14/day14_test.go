package day14

import (
	"reflect"
	"testing"
)

// 10 ORE => 10 A
// 1 ORE => 1 B
// 7 A, 1 B => 1 C
// 7 A, 1 C => 1 D
// 7 A, 1 D => 1 E
// 7 A, 1 E => 1 FUEL
var smallInput = "10 ORE => 10 A\n1 ORE => 1 B\n7 A, 1 B => 1 C\n7 A, 1 C => 1 D\n7 A, 1 D => 1 E\n7 A, 1 E => 1 FUEL"
var smallInputMap = map[string]recipe{
	"A":    recipe{yield: 10, reqs: []requirement{requirement{label: "ORE", quantity: 10}}},
	"B":    recipe{yield: 1, reqs: []requirement{requirement{label: "ORE", quantity: 1}}},
	"C":    recipe{yield: 1, reqs: []requirement{requirement{label: "A", quantity: 7}, requirement{label: "B", quantity: 1}}},
	"D":    recipe{yield: 1, reqs: []requirement{requirement{label: "A", quantity: 7}, requirement{label: "C", quantity: 1}}},
	"E":    recipe{yield: 1, reqs: []requirement{requirement{label: "A", quantity: 7}, requirement{label: "D", quantity: 1}}},
	"FUEL": recipe{yield: 1, reqs: []requirement{requirement{label: "A", quantity: 7}, requirement{label: "E", quantity: 1}}},
}

var input13312 = "157 ORE => 5 NZVS\n165 ORE => 6 DCFZ\n44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL\n12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ\n179 ORE => 7 PSHF\n177 ORE => 5 HKGWZ\n7 DCFZ, 7 PSHF => 2 XJWVT\n165 ORE => 2 GPVTF\n3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT"

var input180697 = "2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG\n17 NVRVD, 3 JNWZP => 8 VPVL\n53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL\n22 VJHF, 37 MNCFX => 5 FWMGM\n139 ORE => 4 NVRVD\n144 ORE => 7 JNWZP\n5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC\n5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV\n145 ORE => 6 MNCFX\n1 NVRVD => 8 CXFTF\n1 VJHF, 6 MNCFX => 4 RFSQX\n176 ORE => 6 VJHF"

var input2210736 = "171 ORE => 8 CNZTR\n7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL\n114 ORE => 4 BHXH\n14 VRPVC => 6 BMBT\n6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL\n6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT\n15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW\n13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW\n5 BMBT => 4 WPTQ\n189 ORE => 9 KTJDG\n1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP\n12 VRPVC, 27 CNZTR => 2 XDBXC\n15 KTJDG, 12 BHXH => 5 XCVML\n3 BHXH, 2 VRPVC => 7 MZWV\n121 ORE => 7 VRPVC\n7 XCVML => 6 RJRHP\n5 BHXH, 4 VRPVC => 5 LTCX"

func TestInputToMap(t *testing.T) {
	input := smallInput
	expected := smallInputMap
	result := inputToMap(input)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestComputeOreRequirementForFuel(t *testing.T) {
	input := smallInputMap
	expected := 31
	result := computeOreRequirementForFuel(input)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestComputeOreRequirementForFuel13312(t *testing.T) {
	recipeBook := inputToMap(input13312)
	expected := 13312
	result := computeOreRequirementForFuel(recipeBook)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestComputeOreRequirementForFuel180697(t *testing.T) {
	recipeBook := inputToMap(input180697)
	expected := 180697
	result := computeOreRequirementForFuel(recipeBook)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestComputeOreRequirementForFuel2210736(t *testing.T) {
	recipeBook := inputToMap(input2210736)
	expected := 2210736
	result := computeOreRequirementForFuel(recipeBook)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestMaxFuelForOre13312(t *testing.T) {
	recipeBook := inputToMap(input13312)
	var input int64 = 1000000000000
	var expected int64 = 82892753
	result := computeMaxFuelForOre(recipeBook, input)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestMaxFuelForOre180697(t *testing.T) {
	recipeBook := inputToMap(input180697)
	var input int64 = 1000000000000
	var expected int64 = 5586022
	result := computeMaxFuelForOre(recipeBook, input)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestMaxFuelForOre2210736(t *testing.T) {
	recipeBook := inputToMap(input2210736)
	var input int64 = 1000000000000
	var expected int64 = 460664
	result := computeMaxFuelForOre(recipeBook, input)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}