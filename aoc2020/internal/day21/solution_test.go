package day21

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFoodLabel(t *testing.T) {
	testCases := []struct {
		inputLine         string
		expectedFoodLabel foodLabel
	}{
		{
			inputLine: "mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
			expectedFoodLabel: foodLabel{
				ingredients: []string{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				allergens:   []string{"dairy", "fish"},
			},
		},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			label, err := parseFoodLabel(tc.inputLine)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedFoodLabel, label)
		})
	}
}

const smallInput = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

var safeIngredients = map[string]struct{}{
	"kfcds": {}, "nhms": {}, "sbzzf": {}, "trh": {},
}

func TestFindSafeIngredients(t *testing.T) {
	labels, err := readFoodList(smallInput)
	assert.NoError(t, err)
	frequencies := findFrequencies(labels)

	ingredientsCharacterization := characterizeIngredients(labels, frequencies)

	assert.Equal(t, safeIngredients, ingredientsCharacterization.safeIngredients)
}
