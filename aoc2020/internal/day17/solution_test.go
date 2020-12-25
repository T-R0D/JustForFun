package day17

import (
	"fmt"
	"testing"

	"github.com/T-R0D/aoc2020/internal/grid"
	"github.com/stretchr/testify/assert"
)

const sampleInitialConfiguration = ".#.\n..#\n###"

var initialActiveCubes = map[grid.Point3]bool{
	{X: 1, Y: 0, Z: 0}: true,
	{X: 2, Y: 1, Z: 0}: true,
	{X: 0, Y: 2, Z: 0}: true,
	{X: 1, Y: 2, Z: 0}: true,
	{X: 2, Y: 2, Z: 0}: true,
}

func TestParseInitialConfiguration(t *testing.T) {
	activeCubes, err := parseInitialConfiguration(sampleInitialConfiguration)

	assert.NoError(t, err)
	assert.Equal(t, initialActiveCubes, activeCubes)
}

func TestEnergySourceSimulatorCycle(t *testing.T) {
	testCases := []struct {
		nCyclesToPerform           int
		expectedNActiveAfterCycles int
	}{
		{
			nCyclesToPerform:           0,
			expectedNActiveAfterCycles: 5,
		},
		{
			nCyclesToPerform:           1,
			expectedNActiveAfterCycles: 11,
		},
		{
			nCyclesToPerform:           2,
			expectedNActiveAfterCycles: 21,
		},
		{
			nCyclesToPerform:           3,
			expectedNActiveAfterCycles: 38,
		},
		{
			nCyclesToPerform:           6,
			expectedNActiveAfterCycles: 112,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("after %d cycles, %d cubes are active", tc.nCyclesToPerform, tc.expectedNActiveAfterCycles)
		t.Run(name, func(t *testing.T) {
			// Arrange.
			simulator := newEnergySourceSimulator(initialActiveCubes)

			// Act.
			for i := 0; i < tc.nCyclesToPerform; i++ {
				simulator.Cycle()
			}

			fmt.Println(simulator.activeCubes)

			// Assert.
			assert.Equal(t, tc.expectedNActiveAfterCycles, simulator.NActiveCubes())
		})
	}
}

func TestEnergySourceSimulator4DCycle(t *testing.T) {
	testCases := []struct {
		nCyclesToPerform           int
		expectedNActiveAfterCycles int
	}{
		{
			nCyclesToPerform:           0,
			expectedNActiveAfterCycles: 5,
		},
		{
			nCyclesToPerform:           1,
			expectedNActiveAfterCycles: 29,
		},
		{
			nCyclesToPerform:           2,
			expectedNActiveAfterCycles: 60,
		},
		{
			nCyclesToPerform:           6,
			expectedNActiveAfterCycles: 848,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("after %d cycles, %d cubes are active", tc.nCyclesToPerform, tc.expectedNActiveAfterCycles)
		t.Run(name, func(t *testing.T) {
			// Arrange.
			simulator := newEnergySourceSimulator4D(initialActiveCubes)

			// Act.
			for i := 0; i < tc.nCyclesToPerform; i++ {
				simulator.Cycle()
			}

			fmt.Println(simulator.activeCubes)

			// Assert.
			assert.Equal(t, tc.expectedNActiveAfterCycles, simulator.NActiveCubes())
		})
	}
}
