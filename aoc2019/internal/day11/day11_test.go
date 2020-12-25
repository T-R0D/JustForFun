package day11

import (
	"testing"
)

func TestCountCoveredAreas(t *testing.T) {
	//                     1     0                0     0                1     0                1     0                0     1                1     0                1     0
	testProg := "3,100,104,1,104,0," + "3,100,104,0,104,0," + "3,100,104,1,104,0," + "3,100,104,1,104,0," + "3,100,104,0,104,1," + "3,100,104,1,104,0," + "3,100,104,1,104,0," + "99"
	expected := 6
	result, err := countCoveredAreas(testProg)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestPaintShipHull(t *testing.T) {
	//                     1     0                0     0                1     0                1     0                0     1                1     0                1     0
	testProg := "3,100,104,1,104,0," + "3,100,104,0,104,0," + "3,100,104,1,104,0," + "3,100,104,1,104,0," + "3,100,104,0,104,1," + "3,100,104,1,104,0," + "3,100,104,1,104,0," + "99"
	expected := "\n..#\n..#\n##.\n"
	result, err := paintShipHull(testProg)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if result != expected {
		t.Fatalf("Expected: %v got: %v", expected, result)
	}
}
