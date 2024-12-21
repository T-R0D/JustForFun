package day20

import (
	"fmt"
	"testing"
)

const exampleRacetrack = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name            string
		input           string
		expected        string
		cheatDuration   int
		significantGain int
	}{
		{
			name:            "PartOne counts paths that are a significant improvement by cheating",
			input:           exampleRacetrack,
			expected:        "2",
			cheatDuration:   2,
			significantGain: 40,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := func() (string, error) {
				track := parseRacetrack(tc.input)

				nPathsThatSignificantlyBeatBaseline :=
					countPathsWithSignificantGain(track, tc.cheatDuration, tc.significantGain)

				return fmt.Sprintf("%d", nPathsThatSignificantlyBeatBaseline), nil
			}()

			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %s, wanted %s", result, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		name            string
		input           string
		expected        string
		cheatDuration   int
		significantGain int
	}{
		{
			name:            "PartTwo counts paths that are a significant improvement by cheating with updated rules",
			input:           exampleRacetrack,
			expected:        "29",
			cheatDuration:   20,
			significantGain: 72,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := func() (string, error) {
				track := parseRacetrack(tc.input)

				nPathsThatSignificantlyBeatBaseline :=
					countPathsWithSignificantGain(track, tc.cheatDuration, tc.significantGain)

				return fmt.Sprintf("%d", nPathsThatSignificantlyBeatBaseline), nil
			}()

			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %s, wanted %s", result, tc.expected)
			}
		})
	}
}
