package day14

import (
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

func (s *Solver) SolvePart1(input string) (interface{}, error) {
	recipeBook := inputToMap(input)
	oreRequired := computeOreRequirementForFuel(recipeBook)
	return oreRequired, nil
}

func (s *Solver) SolvePart2(input string) (interface{}, error) {
	recipeBook := inputToMap(input)
	fuelPossible := computeMaxFuelForOre(recipeBook, 1000000000000)

	//3687787 too high
	// 3687786 is correct...
	// Something about what you set the bulk processing factor to...
	// Hella lame.
	return fuelPossible, nil
}

const (
	LAB_ORE = "ORE"
	LAB_FUE = "FUEL"
)

type recipe struct {
	yield int
	reqs  []requirement
}

type requirement struct {
	label    string
	quantity int
}

func inputToMap(input string) map[string]recipe {
	m := make(map[string]recipe)
	recipes := strings.Split(input, "\n")
	for _, r := range recipes {
		inAndOut := strings.Split(r, " => ")
		in := inAndOut[0]
		out := inAndOut[1]

		ins := strings.Split(in, ", ")
		reqs := make([]requirement, len(ins))
		for i, r := range ins {
			parts := strings.Split(r, " ")
			n, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			reqs[i] = requirement{
				label:    parts[1],
				quantity: n,
			}
		}

		parts := strings.Split(out, " ")
		name := parts[1]
		y, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		m[name] = recipe{
			yield: y,
			reqs:  reqs,
		}
	}
	return m
}

func computeOreRequirementForFuel(recipeBook map[string]recipe) int {
	oreRequired := 0
	surplusMaterials := make(map[string]int)
	needProcessing := []requirement{requirement{quantity: 1, label: LAB_FUE}}

	for len(needProcessing) > 0 {
		req := needProcessing[len(needProcessing)-1]
		needProcessing = needProcessing[:len(needProcessing)-1]

		// If the requirement is ore, just add it to our running total.
		if req.label == LAB_ORE {
			oreRequired += req.quantity
			surplusMaterials[LAB_ORE] += req.quantity
			continue
		}

		// If we have leftover materials from previous reactions,
		// use them first.
		surplus, ok := surplusMaterials[req.label]
		if !ok {
			surplus = 0
			surplusMaterials[req.label] = 0
		}

		if req.quantity >= surplus {
			req.quantity -= surplus
			surplusMaterials[req.label] -= surplus
		} else {
			surplus -= req.quantity
			surplusMaterials[req.label] = surplus
			continue
		}

		// Compute how many times we need to perform the reaction to get the
		// required amount of the current ingredient.
		recipe := recipeBook[req.label]
		times := int(math.Ceil(float64(req.quantity) / float64(recipe.yield)))

		// Pass down the requirements for more materials
		for _, r := range recipe.reqs {
			needProcessing = append(needProcessing, requirement{quantity: times * r.quantity, label: r.label})
		}

		// If this reaction produces a surplus, hang on to it.
		surplusMaterials[req.label] += (times * recipe.yield) - req.quantity
	}
	return oreRequired
}

func conmputeOreRequirementsForFuelWSurplus(recipeBook map[string]recipe) (int64, map[string]int64) {
	var oreRequired int64 = 0
	surplusMaterials := make(map[string]int64)
	needProcessing := []requirement{requirement{quantity: 1, label: LAB_FUE}}

	for len(needProcessing) > 0 {
		req := needProcessing[len(needProcessing)-1]
		needProcessing = needProcessing[:len(needProcessing)-1]

		// If the requirement is ore, just add it to our running total.
		if req.label == LAB_ORE {
			oreRequired += int64(req.quantity)
			surplusMaterials[LAB_ORE] += int64(req.quantity)
			continue
		}

		// If we have leftover materials from previous reactions,
		// use them first.
		surplus, ok := surplusMaterials[req.label]
		if !ok {
			surplus = 0
			surplusMaterials[req.label] = 0
		}

		if int64(req.quantity) >= surplus {
			req.quantity -= int(surplus)
			surplusMaterials[req.label] -= surplus
		} else {
			surplus -= int64(req.quantity)
			surplusMaterials[req.label] = surplus
			continue
		}

		// Compute how many times we need to perform the reaction to get the
		// required amount of the current ingredient.
		recipe := recipeBook[req.label]
		times := int(math.Ceil(float64(req.quantity) / float64(recipe.yield)))

		// Pass down the requirements for more materials
		for _, r := range recipe.reqs {
			needProcessing = append(needProcessing, requirement{quantity: times * r.quantity, label: r.label})
		}

		// If this reaction produces a surplus, hang on to it.
		surplusMaterials[req.label] += int64((times * recipe.yield) - req.quantity)
	}
	return oreRequired, surplusMaterials
}

func computeMaxFuelForOre(recipeBook map[string]recipe, givenOre int64) int64 {
	// costPerFuel, surplusMaterials := conmputeOreRequirementsForFuelWSurplus(recipeBook)

	surplusMaterials := map[string]int64{}
	fuelProduced := int64(0) //givenOre / int64(costPerFuel)
	remainingOre := givenOre //% int64(costPerFuel)
	// for k, v := range surplusMaterials {
	// 	surplusMaterials[k] = v * fuelProduced
	// }
	bulkProcessingFactor := 1 //5000

	// surplusMaterials := make(map[string]int64)
	// remainingOre := givenOre
	// fuelProduced := int64(0)

	for remainingOre > 0 {
		oreRequired := 0
		needProcessing := []requirement{requirement{quantity: bulkProcessingFactor, label: LAB_FUE}}

		for len(needProcessing) > 0 {
			req := needProcessing[len(needProcessing)-1]
			needProcessing = needProcessing[:len(needProcessing)-1]

			// If the requirement is ore, just add it to our running total.
			if req.label == LAB_ORE {
				oreRequired += req.quantity
				surplusMaterials[LAB_ORE] += int64(req.quantity)
				continue
			}

			// If we have leftover materials from previous reactions,
			// use them first.
			surplus, ok := surplusMaterials[req.label]
			if !ok {
				surplus = 0
				surplusMaterials[req.label] = 0
			}

			if int64(req.quantity) >= surplus {
				req.quantity -= int(surplus)
				surplusMaterials[req.label] -= surplus
			} else {
				surplus -= int64(req.quantity)
				surplusMaterials[req.label] = surplus
				continue
			}

			// Compute how many times we need to perform the reaction to get the
			// required amount of the current ingredient.
			recipe := recipeBook[req.label]
			times := int(math.Ceil(float64(req.quantity) / float64(recipe.yield)))

			// Pass down the requirements for more materials
			for _, r := range recipe.reqs {
				needProcessing = append(needProcessing, requirement{quantity: times * r.quantity, label: r.label})
			}

			// If this reaction produces a surplus, hang on to it.
			surplusMaterials[req.label] += int64((times * recipe.yield) - req.quantity)
		}

		if remainingOre >= int64(oreRequired) {
			remainingOre -= int64(oreRequired)
			fuelProduced += int64(bulkProcessingFactor)
		} else {
			break
		}
	}

	for remainingOre > 0 {
		oreRequired := 0
		needProcessing := []requirement{requirement{quantity: 1, label: LAB_FUE}}

		for len(needProcessing) > 0 {
			req := needProcessing[len(needProcessing)-1]
			needProcessing = needProcessing[:len(needProcessing)-1]

			// If the requirement is ore, just add it to our running total.
			if req.label == LAB_ORE {
				oreRequired += req.quantity
				surplusMaterials[LAB_ORE] += int64(req.quantity)
				continue
			}

			// If we have leftover materials from previous reactions,
			// use them first.
			surplus, ok := surplusMaterials[req.label]
			if !ok {
				surplus = 0
				surplusMaterials[req.label] = 0
			}

			if int64(req.quantity) >= surplus {
				req.quantity -= int(surplus)
				surplusMaterials[req.label] -= surplus
			} else {
				surplus -= int64(req.quantity)
				surplusMaterials[req.label] = surplus
				continue
			}

			// Compute how many times we need to perform the reaction to get the
			// required amount of the current ingredient.
			recipe := recipeBook[req.label]
			times := int(math.Ceil(float64(req.quantity) / float64(recipe.yield)))

			// Pass down the requirements for more materials
			for _, r := range recipe.reqs {
				needProcessing = append(needProcessing, requirement{quantity: times * r.quantity, label: r.label})
			}

			// If this reaction produces a surplus, hang on to it.
			surplusMaterials[req.label] += int64((times * recipe.yield) - req.quantity)
		}

		if remainingOre >= int64(oreRequired) {
			remainingOre -= int64(oreRequired)
			fuelProduced += 1
		} else {
			break
		}
	}

	return fuelProduced
}
