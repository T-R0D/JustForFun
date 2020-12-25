package day08

import (
	"reflect"
	"testing"
)

func TestInputStreamToLayers(t *testing.T) {
	input := "123456789012"
	expected := []layer{layer{1, 2, 3, 4, 5, 6}, layer{7, 8, 9, 0, 1, 2}}
	result, err := inputStreamToLayers(input, 3, 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCountDigits(t *testing.T) {
	input := layer{1, 2, 3, 4, 5, 6}

	expected := 0
	result := countDigits(input, 0)
	if result != expected {
		t.Fatalf("Counting 0s - expected: %v, got: %v", expected, result)
	}

	expected = 1
	result = countDigits(input, 1)
	if result != expected {
		t.Fatalf("Counting 1s - expected: %v, got: %v", expected, result)
	}

	expected = 1
	result = countDigits(input, 2)
	if result != expected {
		t.Fatalf("Counting 2s - expected: %v, got: %v", expected, result)
	}
}

func TestFindLayerWithFewestZeros(t *testing.T) {
	input := []layer{layer{1, 2, 3, 4, 5, 6}, layer{7, 8, 9, 0, 1, 2}}
	expected := 0
	result := findLayerWithFewestZeros(input, 6)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestDecodeImage(t *testing.T) {
	layers, err := inputStreamToLayers("0222112222120000", 2, 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := layer{0, 1, 1, 0}
	result := decodeImage(layers, (2 * 2))
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestToString(t *testing.T) {
	expected := "\nX.\n.X\n"
	l := layer{0, 1, 1, 0}
	result := l.ToString(2)
	if expected != result {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}
