package day18

import (
	"aoc2019/internal/location"
	"reflect"
	"testing"
)

const smallExample = "#########\n#b.A.@.a#\n#########"

func TestInputToWorldSmallExample(t *testing.T) {
	input := smallExample
	result := inputToWorld(input)
	expectedEntrance := location.Point{X: 5, Y: 1}
	expectedKeys := map[int]location.Point{
		int('a'): location.Point{X: 7, Y: 1},
		int('b'): location.Point{X: 1, Y: 1},
	}
	expectedHeldKeys := map[int]struct{}{}
	expectedDoors := map[int]location.Point{
		int('A'): location.Point{X: 3, Y: 1},
	}

	if !reflect.DeepEqual(result.entrance, expectedEntrance) {
		t.Fatalf("Expected:\n%v,\ngot:\n%v", expectedEntrance, result.entrance)
	} else if !reflect.DeepEqual(result.keys, expectedKeys) {
		t.Fatalf("Expected:\n%v,\ngot:\n%v", expectedKeys, result.keys)
	} else if !reflect.DeepEqual(result.heldKeys, expectedHeldKeys) {
		t.Fatalf("Expected:\n%v,\ngot:\n%v", expectedHeldKeys, result.heldKeys)
	} else if !reflect.DeepEqual(result.doors, expectedDoors) {
		t.Fatalf("Expected:\n%v,\ngot:\n%v", expectedDoors, result.doors)
	}
}

func TestEfficientlyCollectKeysSmallExample(t *testing.T) {
	input := smallExample
	w := inputToWorld(input)
	result := findAllKeys(w)
	expected := 8
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestEfficientlyCollectKeys2(t *testing.T) {
	input := "########################\n#f.D.E.e.C.b.A.@.a.B.c.#\n######################.#\n#d.....................#\n########################"
	w := inputToWorld(input)
	result := findAllKeys(w)
	expected := 86
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestEfficientlyCollectKeys3(t *testing.T) {
	input := "########################\n#...............b.C.D.f#\n#.######################\n#.....@.a.B.c.d.A.e.F.g#\n########################"
	w := inputToWorld(input)
	result := findAllKeys(w)
	expected := 132
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestEfficientlyCollectKeys4(t *testing.T) {
	input := "#################\n#i.G..c...e..H.p#\n########.########\n#j.A..b...f..D.o#\n########@########\n#k.E..a...g..B.n#\n########.########\n#l.F..d...h..C.m#\n#################"
	w := inputToWorld(input)
	result := findAllKeys(w)
	expected := 136
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestEfficientlyCollectKeys5(t *testing.T) {
	input := "########################\n#@..............ac.GI.b#\n###d#e#f################\n###A#B#C################\n###g#h#i################\n########################"
	w := inputToWorld(input)
	result := findAllKeys(w)
	expected := 81
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCollectKeysWith4Bots(t *testing.T) {
	input := "#######\n#a.#Cd#\n##...## \n##.@.##\n##...##\n#cB#Ab#\n#######"
	w := inputToWorld(input)
	w.updateTo4BotWorld()
	result := findAllKeys2(w)
	expected := 8
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCollectKeysWith4Bots2(t *testing.T) {
	input := "###############\n#d.ABC.#.....a#\n######...######\n######.@.######\n######...######\n#b.....#.....c#\n###############"
	w := inputToWorld(input)
	w.updateTo4BotWorld()
	result := findAllKeys2(w)
	expected := 24
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCollectKeysWith4Bots3(t *testing.T) {
	input := "#############\n#DcBa.#.GhKl#\n#.###...#I###\n#e#d#.@.#j#k#\n###C#...###J#\n#fEbA.#.FgHi#\n#############"
	w := inputToWorld(input)
	w.updateTo4BotWorld()
	result := findAllKeys2(w)
	expected := 32
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}

func TestCollectKeysWith4Bots4(t *testing.T) {
	input := "#############\n#g#f.D#..h#l#\n#F###e#E###.#\n#dCba...BcIJ#\n#####.@.#####\n#nK.L...G...#\n#M###N#H###.#\n#o#m..#i#jk.#\n#############"
	w := inputToWorld(input)
	w.updateTo4BotWorld()
	result := findAllKeys2(w)
	expected := 72
	if result != expected {
		t.Fatalf("Expected: %v, got: %v", expected, result)
	}
}
