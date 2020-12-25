package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayCombat(t *testing.T) {
	startingHands := [][]uint64{
		{9, 2, 6, 3, 1},
		{5, 8, 4, 7, 10},
	}

	resultingHands, err := playCombat(startingHands)
	assert.NoError(t, err)

	expectedResultingHands := [][]uint64{
		{},
		{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
	}
	assert.Equal(t, expectedResultingHands, resultingHands)
}

func TestScoreHand(t *testing.T) {
	hand := []uint64{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}

	score := scoreHand(hand)

	assert.Equal(t, uint64(306), score)
}

func TestPlayRecursiveCombat(t *testing.T) {
	startingHands := [][]uint64{
		{9, 2, 6, 3, 1},
		{5, 8, 4, 7, 10},
	}

	seenGames := map[string]int{}
	resultingHands, err := playRecursiveCombat(startingHands, seenGames)

	expectedResultingHands := [][]uint64{
		{},
		{7, 5, 6, 2, 4, 1, 10, 8, 9, 3},
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedResultingHands, resultingHands)
}
