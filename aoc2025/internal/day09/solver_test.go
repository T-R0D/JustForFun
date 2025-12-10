package day09

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct{
		name     string
		input    string
		expected string
	}{
		{
			name: "Part one finds the largest possible area cornered by two tiles",
			input: testTilePlacements,
			expected: "50",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			solver := Solver{}

			result, err := solver.SolvePartOne(tc.input)
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
	testCases := []struct{
		name     string
		input    string
		expected string
	}{
		// {
		// 	name: "Part two finds the largest possible area cornered by two tiles internal to the polygon",
		// 	input: testTilePlacements,
		// 	expected: "24",
		// },
		{
			name: "Part two finds the largest possible area cornered by two tiles internal to the cross",
			input: crossPlacement,
			expected: "56",
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			solver := Solver{}

			result, err := solver.SolvePartTwo(tc.input)
			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %s, wanted %s", result, tc.expected)
			}
		})
	}
}


/*
..............
.......#...#..
..............
..O....#......
..............
..#......O....
..............
.........#.#..
..............
24
*/
const testTilePlacements = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`


/*
..............
...O...#..#...
..............
.#.#......#.#.
..............
.#.#......#.#.
..............
...#..#...O...
..............
56
*/
const crossPlacement = `3,1
7,1
10,1
10,3
12,3
12,5
10,5
10,7
6,7
3,7
3,5
1,5
1,3
3,3`