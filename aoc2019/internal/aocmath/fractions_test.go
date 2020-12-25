package aocmath

import (
	"testing"
)

func TestGreatestCommonDivisor(t *testing.T) {
	a := GreatestCommonDivisor(1, 2)
	assertEqual(t, 1, a)

	a = GreatestCommonDivisor(3, 5)
	assertEqual(t, 1, a)

	a = GreatestCommonDivisor(9, 3)
	assertEqual(t, 3, a)

	a = GreatestCommonDivisor(6, 14)
	assertEqual(t, 2, a)
}

func TestReduceFraction(t *testing.T) {
	n, d := ReduceFraction(1, 2)
	assertEqual(t, 1, n)
	assertEqual(t, 2, d)

	n, d = ReduceFraction(3, 5)
	assertEqual(t, 3, n)
	assertEqual(t, 5, d)

	n, d = ReduceFraction(9, 3)
	assertEqual(t, 3, n)
	assertEqual(t, 1, d)

	n, d = ReduceFraction(6, 14)
	assertEqual(t, 3, n)
	assertEqual(t, 7, d)

	n, d = ReduceFraction(-6, 14)
	assertEqual(t, -3, n)
	assertEqual(t, 7, d)
}

func assertEqual(t *testing.T, e, a int) {
	if e != a {
		t.Fatalf("Expected: %v, got: %v", e, a)
	}
}
