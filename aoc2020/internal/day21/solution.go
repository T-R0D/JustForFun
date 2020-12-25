package day21

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

// Solver solves the day's problem.
type Solver struct{}

// Part1 solves part 1 of the day's problem.
func (s *Solver) Part1(input string) (string, error) {
	labels, err := readFoodList(input)
	if err != nil {
		return "", err
	}

	frequencies := findFrequencies(labels)

	characterization := characterizeIngredients(labels, frequencies)

	sumOfAppearances := 0
	for ingredient := range characterization.safeIngredients {
		if freq, ok := frequencies.ingredients[ingredient]; ok {
			sumOfAppearances += freq
		}
	}

	return fmt.Sprintf("%d", sumOfAppearances), nil
}

// Part2 solves part 2 of the day's problem.
func (s *Solver) Part2(input string) (string, error) {
	labels, err := readFoodList(input)
	if err != nil {
		return "", err
	}

	frequencies := findFrequencies(labels)

	characterization := characterizeIngredients(labels, frequencies)

	canonicalDangerousingredientsList := createCanonicalDangerousIngredientList(characterization)

	return canonicalDangerousingredientsList, nil
}

type foodLabel struct {
	ingredients []string
	allergens   []string
}

func readFoodList(input string) ([]foodLabel, error) {
	lines := strings.Split(input, "\n")

	labels := make([]foodLabel, len(lines))
	for i, line := range lines {
		label, err := parseFoodLabel(line)
		if err != nil {
			return nil, errors.Wrapf(err, "processing line %d", i)
		}
		labels[i] = label
	}

	return labels, nil
}

func parseFoodLabel(line string) (foodLabel, error) {
	label := foodLabel{}

	ingredientListAndAllergenList := strings.Split(line, " (contains ")
	if len(ingredientListAndAllergenList) != 2 {
		return label, errors.New("line did not contain an ingredient list AND allergen list")
	}

	label.ingredients = strings.Split(ingredientListAndAllergenList[0], " ")
	label.allergens = strings.Split(strings.ReplaceAll(ingredientListAndAllergenList[1], ")", ""), ", ")

	return label, nil
}

type frequenciesPair struct {
	ingredients map[string]int
	allergens   map[string]int
}

func findFrequencies(labels []foodLabel) frequenciesPair {
	ingredientFrequency := map[string]int{}
	allergenFrequency := map[string]int{}
	for _, label := range labels {
		for _, ingredient := range label.ingredients {
			if freq, ok := ingredientFrequency[ingredient]; ok {
				ingredientFrequency[ingredient] = freq + 1
			} else {
				ingredientFrequency[ingredient] = 1
			}
		}

		for _, allergen := range label.allergens {
			if freq, ok := ingredientFrequency[allergen]; ok {
				allergenFrequency[allergen] = freq + 1
			} else {
				allergenFrequency[allergen] = 1
			}
		}
	}
	return frequenciesPair{
		ingredients: ingredientFrequency,
		allergens:   allergenFrequency,
	}
}

type ingredientCharacterization struct {
	safeIngredients              map[string]struct{}
	dangerousIngredients         map[string]string
	allergensAndTheirIngredients map[string]string
}

func characterizeIngredients(labels []foodLabel, frequencies frequenciesPair) *ingredientCharacterization {
	type setPair struct {
		ingredients stringSet
		allergens   stringSet
	}

	labelSets := make([]*setPair, len(labels))
	for i, label := range labels {
		ingredients := stringSet{}
		for _, ingredient := range label.ingredients {
			ingredients.Add(ingredient)
		}
		allergens := stringSet{}
		for _, allergen := range label.allergens {
			allergens.Add(allergen)
		}
		labelSets[i] = &setPair{
			ingredients: ingredients,
			allergens:   allergens,
		}
	}

	allergenCandidates := map[string]stringSet{}
	for _, labelSet := range labelSets {
		for allergen := range labelSet.allergens {
			if candidates, ok := allergenCandidates[allergen]; ok {
				allergenCandidates[allergen] = candidates.Intersection(labelSet.ingredients)
			} else {
				candidates := stringSet{}
				for ingredient := range labelSet.ingredients {
					candidates.Add(ingredient)
				}
				allergenCandidates[allergen] = candidates
			}
		}
	}

	ingredientsWithAllergen := map[string]string{}
	allergensWithIdentifiedIngredient := map[string]string{}
	for len(ingredientsWithAllergen) < len(allergenCandidates) {
		for allergen, candidates := range allergenCandidates {
			switch len(candidates) {
			case 1:
				for ingredient := range candidates {
					ingredientsWithAllergen[ingredient] = allergen
					allergensWithIdentifiedIngredient[allergen] = ingredient
				}
			default:
				for ingredient := range ingredientsWithAllergen {
					delete(candidates, ingredient)
				}
				allergenCandidates[allergen] = candidates
			}
		}
	}

	safeIngredients := stringSet{}
	for ingredient := range frequencies.ingredients {
		if _, ok := ingredientsWithAllergen[ingredient]; !ok {
			safeIngredients.Add(ingredient)
		}
	}

	return &ingredientCharacterization{
		safeIngredients:              map[string]struct{}(safeIngredients),
		dangerousIngredients:         ingredientsWithAllergen,
		allergensAndTheirIngredients: allergensWithIdentifiedIngredient,
	}
}

func createCanonicalDangerousIngredientList(characterization *ingredientCharacterization) string {
	allergens := make([]string, 0, len(characterization.allergensAndTheirIngredients))
	for allergen := range characterization.allergensAndTheirIngredients {
		allergens = append(allergens, allergen)
	}

	sort.Strings(allergens)

	ingredients := make([]string, len(characterization.dangerousIngredients))
	for i, allergen := range allergens {
		ingredients[i] = characterization.allergensAndTheirIngredients[allergen]
	}

	return strings.Join(ingredients, ",")
}

type stringSet map[string]struct{}

func (s stringSet) Add(newItem string) {
	s[newItem] = struct{}{}
}

func (s stringSet) Intersection(other stringSet) stringSet {
	out := stringSet{}
	for item := range s {
		if _, ok := other[item]; ok {
			out[item] = struct{}{}
		}
	}
	return out
}
