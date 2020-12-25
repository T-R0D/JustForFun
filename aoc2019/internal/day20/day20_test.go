package day20

import (
	"aoc2019/internal/location"
	"fmt"
	"reflect"
	"testing"
)

const smallExample = `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `

const largeExample = `                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               `

const recursiveInterestingExample = `             Z L X W       C                 
             Z P Q B       K                 
  ###########.#.#.#.#######.###############  
  #...#.......#.#.......#.#.......#.#.#...#  
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  
  #.#...#.#.#...#.#.#...#...#...#.#.......#  
  #.###.#######.###.###.#.###.###.#.#######  
  #...#.......#.#...#...#.............#...#  
  #.#########.#######.#.#######.#######.###  
  #...#.#    F       R I       Z    #.#.#.#  
  #.###.#    D       E C       H    #.#.#.#  
  #.#...#                           #...#.#  
  #.###.#                           #.###.#  
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#  
CJ......#                           #.....#  
  #######                           #######  
  #.#....CK                         #......IC
  #.###.#                           #.###.#  
  #.....#                           #...#.#  
  ###.###                           #.#.#.#  
XF....#.#                         RF..#.#.#  
  #####.#                           #######  
  #......CJ                       NM..#...#  
  ###.#.#                           #.###.#  
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#  
  #.....#        F   Q       P      #.#.#.#  
  ###.###########.###.#######.#########.###  
  #.....#...#.....#.......#...#.....#.#...#  
  #####.#.###.#######.#######.###.###.#.#.#  
  #.......#.......#.#.#.#.#...#...#...#.#.#  
  #####.###.#####.#.#.#.#.###.###.#.###.###  
  #.......#.....#.#...#...............#...#  
  #############.#.#.###.###################  
               A O F   N                     
               A A D   M                     `

//    Walk from AA to XF (16 steps)
//    Recurse into level 1 through XF (1 step)
//    Walk from XF to CK (10 steps)
//    Recurse into level 2 through CK (1 step)
//    Walk from CK to ZH (14 steps)
//    Recurse into level 3 through ZH (1 step)
//    Walk from ZH to WB (10 steps)
//    Recurse into level 4 through WB (1 step)
//    Walk from WB to IC (10 steps)
//    Recurse into level 5 through IC (1 step)
//    Walk from IC to RF (10 steps)
//    Recurse into level 6 through RF (1 step)
//    Walk from RF to NM (8 steps)
//    Recurse into level 7 through NM (1 step)
//    Walk from NM to LP (12 steps)
//    Recurse into level 8 through LP (1 step)
//    Walk from LP to FD (24 steps)
//    Recurse into level 9 through FD (1 step)
//    Walk from FD to XQ (8 steps)
//    Recurse into level 10 through XQ (1 step)
//    Walk from XQ to WB (4 steps)
//    Return to level 9 through WB (1 step)
//    Walk from WB to ZH (10 steps)
//    Return to level 8 through ZH (1 step)
//    Walk from ZH to CK (14 steps)
//    Return to level 7 through CK (1 step)
//    Walk from CK to XF (10 steps)
//    Return to level 6 through XF (1 step)
//    Walk from XF to OA (14 steps)
//    Return to level 5 through OA (1 step)
//    Walk from OA to CJ (8 steps)
//    Return to level 4 through CJ (1 step)
//    Walk from CJ to RE (8 steps)
//    Return to level 3 through RE (1 step)
//    Walk from RE to IC (4 steps)
//    Recurse into level 4 through IC (1 step)
//    Walk from IC to RF (10 steps)
//    Recurse into level 5 through RF (1 step)
//    Walk from RF to NM (8 steps)
//    Recurse into level 6 through NM (1 step)
//    Walk from NM to LP (12 steps)
//    Recurse into level 7 through LP (1 step)
//    Walk from LP to FD (24 steps)
//    Recurse into level 8 through FD (1 step)
//    Walk from FD to XQ (8 steps)
//    Recurse into level 9 through XQ (1 step)
//    Walk from XQ to WB (4 steps)
//    Return to level 8 through WB (1 step)
//    Walk from WB to ZH (10 steps)
//    Return to level 7 through ZH (1 step)
//    Walk from ZH to CK (14 steps)
//    Return to level 6 through CK (1 step)
//    Walk from CK to XF (10 steps)
//    Return to level 5 through XF (1 step)
//    Walk from XF to OA (14 steps)
//    Return to level 4 through OA (1 step)
//    Walk from OA to CJ (8 steps)
//    Return to level 3 through CJ (1 step)
//    Walk from CJ to RE (8 steps)
//    Return to level 2 through RE (1 step)
//    Walk from RE to XQ (14 steps)
//    Return to level 1 through XQ (1 step)
//    Walk from XQ to FD (8 steps)
//    Return to level 0 through FD (1 step)
//    Walk from FD to ZZ (18 steps)

func TestWorldFromInput(t *testing.T) {
	input := smallExample
	w := worldFromInput(input)

	expectedEntrance := location.Point{X: 9, Y: 2}
	if w.entrance != expectedEntrance {
		t.Fatalf("Expected: %v, got: %v", expectedEntrance, w.entrance)
	}

	expectedExit := location.Point{X: 13, Y: 16}
	if w.exit != expectedExit {
		t.Fatalf("Expected: %v, got: %v", expectedExit, w.exit)
	}

	expectedPortals := map[location.Point]location.Point{
		// B<->C
		location.Point{X: 9, Y: 6}: location.Point{X: 2, Y: 8},
		location.Point{X: 2, Y: 8}: location.Point{X: 9, Y: 6},
		// D<->E
		location.Point{X: 6, Y: 10}: location.Point{X: 2, Y: 13},
		location.Point{X: 2, Y: 13}: location.Point{X: 6, Y: 10},
		// F<->G
		location.Point{X: 11, Y: 12}: location.Point{X: 2, Y: 15},
		location.Point{X: 2, Y: 15}:  location.Point{X: 11, Y: 12},
	}
	if !reflect.DeepEqual(w.portals, expectedPortals) {
		t.Fatalf("Expected:\n%v\ngot:\n%v", expectedPortals, w.portals)
	}
}

func TestFindShortestPathSmall(t *testing.T) {
	w := worldFromInput(smallExample)
	path := findShortestPathThroughMaze(w)
	steps := len(path) - 1
	expected := 23
	if steps != expected {
		t.Fatalf("Expected: %v, got: %v", expected, steps)
	}
}

func TestFindShortestPathLarge(t *testing.T) {
	w := worldFromInput(largeExample)
	path := findShortestPathThroughMaze(w)
	steps := len(path) - 1
	expected := 58
	if steps != expected {
		t.Fatalf("Expected: %v, got: %v", expected, steps)
	}
}

func TestFindShortestPathRecursiveSmall(t *testing.T) {
	w := worldFromInput(smallExample)
	path := findShortestPathThroughRecursiveMaze(w)
	steps := len(path) - 1
	expected := 26
	if steps != expected {
		t.Fatalf("Expected: %v, got: %v", expected, steps)
	}
}

// func TestFindShortestPathRecursiveLarge(t *testing.T) {
// 	w := worldFromInput(largeExample)
// 	path := findShortestPathThroughRecursiveMaze(w)
// 	steps := len(path) - 1
// 	expected := 58
// 	if steps != expected {
// 		t.Fatalf("Expected: %v, got: %v", expected, steps)
// 	}
// }

func TestFindShortestPathRecursiveInteresting(t *testing.T) {
	w := worldFromInput(recursiveInterestingExample)
	path := findShortestPathThroughRecursiveMaze(w)
	steps := len(path) - 1
	expected := 396

	for i, step := range path {
		fmt.Printf("% 4d -   %s\n", i, step.String())
	}

	if steps != expected {
		t.Fatalf("Expected: %v, got: %v", expected, steps)
	}
}
