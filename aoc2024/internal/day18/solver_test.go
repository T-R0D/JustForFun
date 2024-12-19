package day18

import (
	"fmt"
	"strconv"
	"testing"
)

const exampleByteList = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestPartOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "PartOne finds the shortest safe path through the memory grid",
			input:    exampleByteList,
			expected: "22",
		},
	}

	const (
		testM = 7
		testN = 7

		testNBytefalls = 12
	)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := func() (string, error) {
				bytefallLocations, err := parseBytefallLocations(tc.input)
				if err != nil {
					return "", err
				}

				corruptedLocations := findCorruptedLocations(bytefallLocations[:testNBytefalls])
				shortestSafePath, ok := findShortestPathThroughMemoryGrid(testM, testN, corruptedLocations)
				if !ok {
					return "", fmt.Errorf("could not find a valid path through the memory grid")
				}

				return strconv.Itoa(len(shortestSafePath) - 1), nil
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
		name     string
		input    string
		expected string
	}{
		{
			name:     "PartTwo identifies the first bytefall location that cuts off the path to the exit",
			input:    exampleByteList,
			expected: "6,1",
		},
	}

	const (
		testM = 7
		testN = 7
	)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result, err := func() (string, error) {
				bytefallLocations, err := parseBytefallLocations(tc.input)
				if err != nil {
					return "", err
				}

				location, ok := findFirstBytefallToBlockPathToExit(testM, testN, bytefallLocations)
				if !ok {
					return "", fmt.Errorf("the path was never blocked")
				}

				return location.Swap().String(), nil
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
