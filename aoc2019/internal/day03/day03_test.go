package day03

import (
	"aoc2019/internal/location"
	"reflect"
	"testing"
)

func TestInputToMaps(t *testing.T) {
	testInputToMaps(t,
		"R1,U1\nL1,D1",
		[]pointSet{
			pointSet{
				location.Point{X: 1, Y: 0}: struct{}{},
				location.Point{X: 1, Y: 1}: struct{}{},
			},
			pointSet{
				location.Point{X: -1, Y: 0}:  struct{}{},
				location.Point{X: -1, Y: -1}: struct{}{},
			},
		})

	testInputToMaps(t,
		"R1,R1,U2,L1,D1\nL1,R2,D1,U2,L3",
		[]pointSet{
			pointSet{
				location.Point{X: 1, Y: 0}: struct{}{},
				location.Point{X: 2, Y: 0}: struct{}{},
				location.Point{X: 2, Y: 1}: struct{}{},
				location.Point{X: 2, Y: 2}: struct{}{},
				location.Point{X: 1, Y: 2}: struct{}{},
				location.Point{X: 1, Y: 1}: struct{}{},
			},
			pointSet{
				location.Point{X: -1, Y: 0}: struct{}{},
				location.Point{X: 0, Y: 0}:  struct{}{},
				location.Point{X: 1, Y: 0}:  struct{}{},
				location.Point{X: 1, Y: -1}: struct{}{},
				location.Point{X: 1, Y: 1}:  struct{}{},
				location.Point{X: 0, Y: 1}:  struct{}{},
				location.Point{X: -1, Y: 1}: struct{}{},
				location.Point{X: -2, Y: 1}: struct{}{},
			},
		})
}

func testInputToMaps(t *testing.T, in string, expected []pointSet) {
	result, err := inputToMaps(in)
	if nil != err || !reflect.DeepEqual(result, expected) {
		t.Fatalf("Failure - expected: %v, got: %v; err: %v", expected, result, err)
	}
}

func TestFindDistanceToNearestIntersection(t *testing.T) {
	_, err := findDistanceToIntersectionNearestOrigin([]pointSet{
		pointSet{
			location.Point{X: 1, Y: 0}: struct{}{},
			location.Point{X: 1, Y: 1}: struct{}{},
		},
		pointSet{
			location.Point{X: -1, Y: 0}:  struct{}{},
			location.Point{X: -1, Y: -1}: struct{}{},
		},
	})
	if nil == err {
		t.Fatalf("Expected failure, no intersection should be found")
	}

	testFindDistanceToNearestIntersection(t, []pointSet{
		pointSet{
			location.Point{X: 1, Y: 0}: struct{}{},
			location.Point{X: 2, Y: 0}: struct{}{},
			location.Point{X: 2, Y: 1}: struct{}{},
			location.Point{X: 2, Y: 2}: struct{}{},
			location.Point{X: 1, Y: 2}: struct{}{},
			location.Point{X: 1, Y: 1}: struct{}{},
		},
		pointSet{
			location.Point{X: -1, Y: 0}: struct{}{},
			location.Point{X: 0, Y: 0}:  struct{}{},
			location.Point{X: 1, Y: 0}:  struct{}{},
			location.Point{X: 1, Y: -1}: struct{}{},
			location.Point{X: 1, Y: 1}:  struct{}{},
			location.Point{X: 0, Y: 1}:  struct{}{},
			location.Point{X: -1, Y: 1}: struct{}{},
			location.Point{X: -2, Y: 1}: struct{}{},
		},
	}, 1)
}

func testFindDistanceToNearestIntersection(t *testing.T, input []pointSet, expected int) {
	result, err := findDistanceToIntersectionNearestOrigin(input)
	if nil != err {
		t.Fatalf("Expected failure, no intersection should be found")
	} else if result != expected {
		t.Fatalf("Expected: %d, got %d", expected, result)
	}
}

func TestSolvePart1(t *testing.T) {
	testSolvePart1(t, "R8,U5,L5,D3\nU7,R6,D4,L4", 6)

	testSolvePart1(t, "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159)

	testSolvePart1(t, "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135)
}

func testSolvePart1(t *testing.T, input string, expected int) {
	s := Solver{}
	result, err := s.SolvePart1(input)
	if err != nil {
		t.Fatalf("Error finding nearest intersection: %v", err)
	} else if result != expected {
		t.Fatalf("Failure - expected: %d, got: %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	testSolvePart2(t, "R8,U5,L5,D3\nU7,R6,D4,L4", 30)

	testSolvePart2(t, "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 610)

	testSolvePart2(t, "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410)
}

func testSolvePart2(t *testing.T, input string, expected int) {
	s := Solver{}
	result, err := s.SolvePart2(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}
