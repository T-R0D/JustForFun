package day16

import (
	"reflect"
	"testing"
)

func TestRepeatingPatternOrder0(t *testing.T) {
	expected := []int{1, 0, -1, 0}
	result := newRepeatingPattern([]int{0, 1, 0, -1}, 0).toIntArray()
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestRepeatingPatternOrder2(t *testing.T) {
	expected := []int{0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1, 0}
	result := newRepeatingPattern([]int{0, 1, 0, -1}, 2).toIntArray()
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestRepeatingPatternGetOrder0(t *testing.T) {
	r := newRepeatingPattern([]int{0,1,2,3,4}, 0)
	expected := 1
	result := r.get(0)
	if result != expected {
		t.Fatalf("Expected: %v, got %v", expected, result)
	}

	expected = 3
	result = r.get(2)
	if result != expected {
		t.Fatalf("Expected: %v, got %v", expected, result)
	}

	expected = 4
	result = r.get(8)
	if result != expected {
		t.Fatalf("Expected: %v, got %v", expected, result)
	}
}

func TestRepeatingPatternGetOrder2(t *testing.T) {
	//  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  3,  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  4,  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  4,  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  4,
	//  0   1,  2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47                 

	r := newRepeatingPattern([]int{3,2,1,0}, 2)
	
	expected := 3
	result := r.get(0)
	if result != expected {
		t.Fatalf("Expected: %v, got %v", expected, result)
	}
	
	expected = 2
	result = r.get(2)
	if result != expected {
		t.Fatalf("Expected: %v, got %v", expected, result)
	}
	
	expected = 0
	result = r.get(44)
	if result != expected {
		t.Fatalf("Expected: %v, got %v", expected, result)
	}
	
	expecteds := []int{3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  3,  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  3,  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  3,  3,  3,  2,  2,  2,  1,  1,  1,  0,  0,  0,  3}
	for i, expected := range expecteds {
		result := r.get(i)
		if result != expected {
			t.Fatalf("Expected: %v, got %v", expected, result)
		}
	}
}

func TestCleanSignalSmall(t *testing.T) {
	inputSignal := "12345678"
	signal := inputToIntArray(inputSignal)
	signal = cleanSignal(signal, 1)
	expected := "48226158"
	result := getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}

	signal = cleanSignal(signal, 1)
	expected = "34040438"
	result = getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}

	signal = cleanSignal(signal, 1)
	expected = "03415518"
	result = getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}

	signal = cleanSignal(signal, 1)
	expected = "01029498"
	result = getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCleanSignal(t *testing.T) {
	inputSignal := "80871224585914546619083218645595"
	signal := inputToIntArray(inputSignal)
	signal = cleanSignal(signal, 100)
	expected := "24176176"
	result := getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCleanSignal2(t *testing.T) {
	inputSignal := "19617804207202209144916044189917"
	signal := inputToIntArray(inputSignal)
	signal = cleanSignal(signal, 100)
	expected := "73745418"
	result := getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCleanSignal3(t *testing.T) {
	inputSignal := "69317163492948606335995924319873"
	signal := inputToIntArray(inputSignal)
	signal = cleanSignal(signal, 100)
	expected := "52432133"
	result := getDigitsFromSignal(signal, 0, 8)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestSoutionPart2(t *testing.T) {
	input := "03036732577212944063491565474664"
	expected := "84462026"
	s := Solver{}
	result, err := s.SolvePart2(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}
