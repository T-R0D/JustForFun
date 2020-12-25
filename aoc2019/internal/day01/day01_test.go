package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	in := 12
	expected := 2
	result := requiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}

	in = 14
	expected = 2
	result = requiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}

	in = 1969
	expected = 654
	result = requiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}

	in = 100756
	expected = 33583
	result = requiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	in := 14
	expected := 2
	result := actualRequiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}

	in = 1969
	expected = 966
	result = actualRequiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}

	in = 100756
	expected = 50346
	result = actualRequiredFuelForPart(in)
	if result != expected {
		t.Fatalf("Expected: %d - got: %d", expected, result)
	}
}
