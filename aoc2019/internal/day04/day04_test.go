package day04

import (
	"testing"
)

func TestMonotonicallyIncreasing(t *testing.T) {
	testMonotonicallyIncreasing(t, []int{1, 1, 1, 1, 1, 1}, true)

	testMonotonicallyIncreasing(t, []int{2, 2, 3, 4, 5, 0}, false)

	testMonotonicallyIncreasing(t, []int{1, 2, 3, 7, 8, 9}, true)
}

func testMonotonicallyIncreasing(t *testing.T, in []int, expected bool) {
	result := digitsAreInIncreasingOrder(in)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestHasAdjacentPair(t *testing.T) {
	testHasAdjacentPair(t, []int{1, 1, 1, 1, 1, 1}, true)

	testHasAdjacentPair(t, []int{2, 2, 3, 4, 5, 0}, true)

	testHasAdjacentPair(t, []int{1, 2, 3, 7, 8, 9}, false)
}

func testHasAdjacentPair(t *testing.T, in []int, expected bool) {
	result := codeHasAdjacentPairOfDigits(in)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestHasStrictPair(t *testing.T) {
	testHasStrictPair(t, []int{2, 2, 3, 4, 5, 0}, true)

	testHasStrictPair(t, []int{1, 2, 3, 7, 8, 9}, false)

	testHasStrictPair(t, []int{1, 1, 2, 2, 3, 3}, true)

	testHasStrictPair(t, []int{1, 2, 3, 4, 4, 4}, false)

	testHasStrictPair(t, []int{1, 1, 1, 1, 2, 2}, true)
}

func testHasStrictPair(t *testing.T, in []int, expected bool) {
	result := codeHasStrictPairOfDigits(in)
	if result != expected {
		t.Fatalf("For input %v, expected: %v, got: %v", in, expected, result)
	}
}
