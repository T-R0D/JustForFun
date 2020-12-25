package day10

import (
	"aoc2019/internal/location"
	"reflect"
	"testing"
)

var smallInput = ".#..#\n.....\n#####\n....#\n...##"
var smallAsteroidSet = asteroidSet{
	location.Point{X: 1, Y: 0}: struct{}{},
	location.Point{X: 4, Y: 0}: struct{}{},
	location.Point{X: 0, Y: 2}: struct{}{},
	location.Point{X: 1, Y: 2}: struct{}{},
	location.Point{X: 2, Y: 2}: struct{}{},
	location.Point{X: 3, Y: 2}: struct{}{},
	location.Point{X: 4, Y: 2}: struct{}{},
	location.Point{X: 4, Y: 3}: struct{}{},
	location.Point{X: 3, Y: 4}: struct{}{},
	location.Point{X: 4, Y: 4}: struct{}{},
}

var mediumInput1 = "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####"

var mediumInput2 = "#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###."

var mediumInput3 = ".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#.."

var largeInput = ".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##"

var originOfContrivedSituation = location.Point{X: 2, Y: 2}
var ramOfContrivedSituation = relativeAsteroidMapping{
	location.Vector{X: -1, Y: 0}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 1, Y: 2},
			Dist: 1,
		},
		relativeAsteroid{
			Loc:  location.Point{X: 0, Y: 2},
			Dist: 2,
		},
	},
	location.Vector{X: 1, Y: 0}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 3, Y: 2},
			Dist: 1,
		},
		relativeAsteroid{
			Loc:  location.Point{X: 4, Y: 2},
			Dist: 2,
		},
	},
	location.Vector{X: -1, Y: -2}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 1, Y: 0},
			Dist: 3,
		},
	},
	location.Vector{X: 1, Y: -1}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 4, Y: 0},
			Dist: 4,
		},
	},
	location.Vector{X: 2, Y: 1}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 4, Y: 3},
			Dist: 3,
		},
	},
	location.Vector{X: 1, Y: 2}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 3, Y: 4},
			Dist: 3,
		},
	},
	location.Vector{X: 1, Y: 1}: relativeAsteroidList{
		relativeAsteroid{
			Loc:  location.Point{X: 4, Y: 4},
			Dist: 4,
		},
	},
}

func TestInputToAsteroidSet(t *testing.T) {
	expected := smallAsteroidSet
	result, err := inputToAsteroidSet(smallInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected:\n%v,\ngot:\n%v", expected, result)
	}
}

func TestFindAsteroidWithBestViewSmallInput(t *testing.T) {
	bestLocation, viewableAsteroids := findAsteroidWithBestView(smallAsteroidSet)
	expectedLocation := location.Point{X: 3, Y: 4}
	expectedViewable := 8
	if !reflect.DeepEqual(bestLocation, expectedLocation) {
		t.Fatalf("Location not correct - expected: %v, got: %v", expectedLocation, bestLocation)
	} else if viewableAsteroids != expectedViewable {
		t.Fatalf("Viewable asteroids not correct - expected: %v, got: %v", expectedViewable, viewableAsteroids)
	}
}

func TestFindAsteroidWithBestViewMediumInput1(t *testing.T) {
	expectedLocation := location.Point{X: 5, Y: 8}
	expectedViewable := 33
	as, err := inputToAsteroidSet(mediumInput1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	bestLocation, viewableAsteroids := findAsteroidWithBestView(as)
	if !reflect.DeepEqual(bestLocation, expectedLocation) {
		t.Fatalf("Location not correct - expected: %v, got: %v", expectedLocation, bestLocation)
	} else if viewableAsteroids != expectedViewable {
		t.Fatalf("Viewable asteroids not correct - expected: %v, got: %v", expectedViewable, viewableAsteroids)
	}
}

func TestFindAsteroidWithBestViewMediumInput2(t *testing.T) {
	expectedLocation := location.Point{X: 1, Y: 2}
	expectedViewable := 35
	as, err := inputToAsteroidSet(mediumInput2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	bestLocation, viewableAsteroids := findAsteroidWithBestView(as)
	if !reflect.DeepEqual(bestLocation, expectedLocation) {
		t.Fatalf("Location not correct - expected: %v, got: %v", expectedLocation, bestLocation)
	} else if viewableAsteroids != expectedViewable {
		t.Fatalf("Viewable asteroids not correct - expected: %v, got: %v", expectedViewable, viewableAsteroids)
	}
}

func TestFindAsteroidWithBestViewMediumInput3(t *testing.T) {
	expectedLocation := location.Point{X: 6, Y: 3}
	expectedViewable := 41
	as, err := inputToAsteroidSet(mediumInput3)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	bestLocation, viewableAsteroids := findAsteroidWithBestView(as)
	if !reflect.DeepEqual(bestLocation, expectedLocation) {
		t.Fatalf("Location not correct - expected: %v, got: %v", expectedLocation, bestLocation)
	} else if viewableAsteroids != expectedViewable {
		t.Fatalf("Viewable asteroids not correct - expected: %v, got: %v", expectedViewable, viewableAsteroids)
	}
}

func TestFindAsteroidWithBestViewLargeInput(t *testing.T) {
	expectedLocation := location.Point{X: 11, Y: 13}
	expectedViewable := 210
	as, err := inputToAsteroidSet(largeInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	bestLocation, viewableAsteroids := findAsteroidWithBestView(as)
	if !reflect.DeepEqual(bestLocation, expectedLocation) {
		t.Fatalf("Location not correct - expected: %v, got: %v", expectedLocation, bestLocation)
	} else if viewableAsteroids != expectedViewable {
		t.Fatalf("Viewable asteroids not correct - expected: %v, got: %v", expectedViewable, viewableAsteroids)
	}
}

func TestGroupAsteroidsByReducedVector(t *testing.T) {
	origin := originOfContrivedSituation
	as := smallAsteroidSet
	expected := ramOfContrivedSituation
	result := groupAsteroidsByReducedVector(origin, as)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected:\n%v,\ngot:\n%v", expected, result)
	}
}

func TestFindOrderOfDestruction(t *testing.T) {
	// .7..1
	// .....
	// 96028
	// ....3
	// ...54
	expected := []location.Point{
		{X: 4, Y: 0},
		{X: 3, Y: 2},
		{X: 4, Y: 3},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 1, Y: 2},
		{X: 1, Y: 0},
		{X: 4, Y: 2},
		{X: 0, Y: 2},
	}
	result := findOrderOfDestruction(ramOfContrivedSituation)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%v,\ngot:\n%v", expected, result)
	}
}
