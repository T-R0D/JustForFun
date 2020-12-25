package day12

import (
	"reflect"
	"testing"
)

var inputStr = "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>"
var inputVecs = []vector{
	{-1, 0, 2},
	{2, -10, -7},
	{4, -8, 8},
	{3, 5, -1},
}

func TestInputToVectors(t *testing.T) {
	expected := inputVecs
	result, err := inputToPositionVectors(inputStr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestNBodyStep(t *testing.T) {
	bodies, err := inputToBodies(inputStr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []*body{
		&body{P: vector{-1, 0, 2}, V: vector{0, 0, 0}},
		&body{P: vector{2, -10, -7}, V: vector{0, 0, 0}},
		&body{P: vector{4, -8, 8}, V: vector{0, 0, 0}},
		&body{P: vector{3, 5, -1}, V: vector{0, 0, 0}},
	}
	if !reflect.DeepEqual(bodies, expected) {
		t.Fatalf("T=0 - Expected:\n%v\ngot:\n%v", expected, bodies)
	}

	nBodySimulateTimeStep(bodies)
	expected = []*body{
		&body{P: vector{2, -1, 1}, V: vector{3, -1, -1}},
		&body{P: vector{3, -7, -4}, V: vector{1, 3, 3}},
		&body{P: vector{1, -7, 5}, V: vector{-3, 1, -3}},
		&body{P: vector{2, 2, 0}, V: vector{-1, -3, 1}},
	}
	if !reflect.DeepEqual(bodies, expected) {
		t.Fatalf("T=1 - Expected:\n%v\ngot:\n%v", expected, bodies)
	}

	for i := 0; i < 9; i++ {
		nBodySimulateTimeStep(bodies)
	}
	expected = []*body{
		&body{P: vector{2, 1, -3}, V: vector{-3, -2, 1}},
		&body{P: vector{1, -8, 0}, V: vector{-1, 1, 3}},
		&body{P: vector{3, -6, 1}, V: vector{3, 2, -3}},
		&body{P: vector{2, 0, 4}, V: vector{1, -1, -1}},
	}
	if !reflect.DeepEqual(bodies, expected) {
		t.Fatalf("T=10 - Expected:\n%v\ngot:\n%v", expected, bodies)
	}
}

func TestComputeSystemEnergy(t *testing.T) {
	bodies := []*body{
		&body{P: vector{8, -12, -9}, V: vector{-7, 3, 0}},
		&body{P: vector{13, 16, -3}, V: vector{3, -11, -5}},
		&body{P: vector{-29, -11, -1}, V: vector{-3, 7, 4}},
		&body{P: vector{16, -13, 23}, V: vector{7, 1, 1}},
	}
	expected := 1940
	result := computeSystemEnergy(bodies)
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestBodiesAtSameStateForAxis(t *testing.T) {
	bodies := []*body{
		&body{P: vector{8, -12, -9}, V: vector{-7, 3, 0}},
		&body{P: vector{13, 16, -3}, V: vector{3, -11, -5}},
		&body{P: vector{-29, -11, -1}, V: vector{-3, 7, 4}},
		&body{P: vector{16, -13, 23}, V: vector{7, 1, 1}},
	}
	bodies2 := []*body{
		&body{P: vector{8, -12, -9}, V: vector{-7, 3, 0}},
		&body{P: vector{13, 16, -3}, V: vector{3, -11, -5}},
		&body{P: vector{-29, -11, -1}, V: vector{-3, 7, 4}},
		&body{P: vector{16, -13, 23}, V: vector{7, 1, 1}},
	}
	for i := 0; i < DIM; i++ {
		if !bodiesAtSameStateForAxis(bodies, bodies2, i) {
			t.Fatalf("Bodies were not same for axis %v", i)
		}
	}

	bodies = []*body{
		&body{P: vector{8, -12, -9}, V: vector{-7, 3, 0}},
		&body{P: vector{13, 16, -3}, V: vector{3, -11, -5}},
		&body{P: vector{-29, -11, -1}, V: vector{-3, 7, 4}},
		&body{P: vector{16, -13, 23}, V: vector{7, 1, 1}},
	}
	bodies2 = []*body{
		&body{P: vector{8, -12, -9}, V: vector{-7, 3, 0}},
		&body{P: vector{13, 16, -3}, V: vector{3, -11, -5}},
		&body{P: vector{-29, 0, -1}, V: vector{-3, 7, 4}},
		&body{P: vector{16, -13, 23}, V: vector{7, 1, 1}},
	}
	i := 2
	if !bodiesAtSameStateForAxis(bodies, bodies2, i) {
		t.Fatalf("Bodies should differ on axis %v", i)
	}
}

func TestSolvePart2(t *testing.T) {
	input := inputStr
	expected := 2772
	s := Solver{}
	result, err := s.SolvePart2(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}

func TestSolvePart2BigResult(t *testing.T) {
	input := "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>"
	expected := 4686774924
	s := Solver{}
	result, err := s.SolvePart2(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expected, result)
	}
}