package day09

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/T-R0D/aoc2025/v2/internal/set"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	tilePlacements, err := parseTilePlacements(input)
	if err != nil {
		return "", err
	}

	maxRectangleArea := 0
	for i, tileA := range tilePlacements[:len(tilePlacements)-1] {
		for _, tileB := range tilePlacements[i+1:] {
			area := rectangleArea(tileA, tileB)

			if area > maxRectangleArea {
				maxRectangleArea = area
			}
		}
	}

	return strconv.Itoa(maxRectangleArea), nil
}

// 4345908 too low
func (this *Solver) SolvePartTwo(input string) (string, error) {
	tilePlacements, err := parseTilePlacements(input)
	if err != nil {
		return "", err
	}

	verticalSegments, horizontalSegments := catalogSideSegments(tilePlacements)
	// fmt.Printf("vertical segments: %v\n", verticalSegments)
	// fmt.Printf("horizontal segments: %v\n", horizontalSegments)

	xCoordinates, yCoordinates := getAllCoordinates(tilePlacements)
	// fmt.Printf("xs and ys: %v | %v\n", xCoordinates, yCoordinates)

	orderedTiles := getTilesClockwise(tilePlacements)
	// fmt.Printf("ordered tiles: %v\n", orderedTiles)

	verticalRanges := findVerticalRanges(verticalSegments, horizontalSegments, xCoordinates, yCoordinates, orderedTiles)
	// fmt.Printf("verticalRanges: %v\n", verticalRanges)

	largestArea := 0
	for i, a := range orderedTiles {
		for _, b := range orderedTiles[i+1:] {
			if checkRectangle(a, b, verticalRanges, xCoordinates) {
				area := rectangleArea(a, b)
				if area > largestArea {
					largestArea = area
				}
			}
		}
	}

	return strconv.Itoa(largestArea), nil
}

func parseTilePlacements(input string) ([][2]int, error) {
	placements := [][2]int{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return [][2]int{}, fmt.Errorf("unexpected number of coordinates on line; got %d", len(parts))
		}

		coordinates := [2]int{0, 0}
		for i, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				return [][2]int{}, err
			}

			coordinates[i] = value
		}
		coordinates[1] *= -1

		placements = append(placements, coordinates)
	}

	return placements, nil
}

func rectangleArea(cornerA [2]int, cornerB [2]int) int {
	height := intAbs(cornerA[0]-cornerB[0]) + 1
	width := intAbs(cornerA[1]-cornerB[1]) + 1
	return height * width
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func catalogSideSegments(tilePlacements [][2]int) (map[int][][2]int, map[int][][2]int) {
	verticalSegments := map[int][][2]int{}
	horizontalSegments := map[int][][2]int{}

	for i := range tilePlacements {
		a, b := tilePlacements[i], tilePlacements[(i+1)%len(tilePlacements)]

		if a[0] == b[0] {
			x := a[0]
			y1, y2 := a[1], b[1]
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			newRange := [2]int{y1, y2}

			if ranges, exists := verticalSegments[x]; exists {
				verticalSegments[x] = append(ranges, newRange)
			} else {
				verticalSegments[x] = [][2]int{newRange}
			}
		} else {
			y := a[1]
			x1, x2 := a[0], b[0]
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			newRange := [2]int{x1, x2}

			if ranges, exists := horizontalSegments[y]; exists {
				horizontalSegments[y] = append(ranges, newRange)
			} else {
				horizontalSegments[y] = [][2]int{newRange}
			}
		}
	}

	for k, v := range verticalSegments {
		verticalSegments[k] = mergeRanges(v)
	}

	for k, v := range horizontalSegments {
		horizontalSegments[k] = mergeRanges(v)
	}

	return verticalSegments, horizontalSegments
}

func mergeRanges(originalRanges [][2]int) [][2]int {
	if len(originalRanges) < 1 {
		return originalRanges
	}

	ranges := make([][2]int, 0, len(originalRanges))
	ranges = append(ranges, originalRanges...)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	condensedRanges := [][2]int{}
	currentLower := ranges[0][0]
	currentUpper := ranges[0][1]
	for _, r := range ranges {
		if currentUpper < r[0] {
			condensedRanges = append(condensedRanges, [2]int{currentLower, currentUpper})

			currentLower, currentUpper = r[0], r[1]
			continue
		}

		if currentUpper < r[1] {
			currentUpper = r[1]
		}
	}
	condensedRanges = append(condensedRanges, [2]int{currentLower, currentUpper})

	return condensedRanges
}

func getAllCoordinates(tiles [][2]int) ([]int, []int) {
	xs, ys := set.New[int](), set.New[int]()

	for _, tile := range tiles {
		xs.Add(tile[0])
		ys.Add(tile[1])
	}

	return slices.Sorted(xs.All()), slices.Sorted(ys.All())
}

func getTilesClockwise(tiles [][2]int) [][2]int {
	nTiles := len(tiles)
	if nTiles < 2 {
		return tiles
	}

	leftmostUpperTileIndex := 0
	for i, tile := range tiles {
		if tiles[leftmostUpperTileIndex][1] < tile[1] {
			leftmostUpperTileIndex = i
			continue
		}

		if tile[1] == tiles[leftmostUpperTileIndex][1] && tile[0] < tiles[leftmostUpperTileIndex][0] {
			leftmostUpperTileIndex = i
			continue
		}
	}

	clockwiseOrdered := tiles[leftmostUpperTileIndex][0] <= tiles[(leftmostUpperTileIndex+nTiles+1)%nTiles][0]

	orderedTiles := make([][2]int, 0, nTiles)
	if clockwiseOrdered {
		for i := range nTiles {
			orderedTiles = append(orderedTiles, tiles[(i+leftmostUpperTileIndex)%nTiles])
		}
	} else {
		for i := nTiles - 1; i >= 0; i += 1 {
			orderedTiles = append(orderedTiles, tiles[(i+leftmostUpperTileIndex)%nTiles])
		}
	}

	return orderedTiles
}

func findVerticalRanges(
	verticalSegments map[int][][2]int,
	horizontalSegments map[int][][2]int,
	xCoordinates []int,
	yCoordinates []int,
	clockwiseTiles [][2]int) map[int][][2]int {

	verticalRanges := make(map[int][][2]int, len(xCoordinates))

	for _, x := range xCoordinates {
		insidePolygon := false

		start, end := 0, 0
	YScan:
		for _, y := range yCoordinates {
			// fmt.Printf("\t(%d, %d)\n", x, y)

			end = y

			_, coordinateOnHorizontalSegment := coordinateOnSomeSegment(y, x, horizontalSegments)
			if !coordinateOnHorizontalSegment {
				continue YScan
			}

			_, coordinateOnVerticalSegment := coordinateOnSomeSegment(x, y, verticalSegments)
			if coordinateOnHorizontalSegment && !coordinateOnVerticalSegment {
				if !insidePolygon {
					// fmt.Print("\tToggling to 'inside'\n")
					insidePolygon = true
					start = y
				} else {
					// fmt.Print("\tToggling to 'outside (not on vertical segment)'\n")

					insidePolygon = false
					ranges, exists := verticalRanges[x]
					if !exists {
						ranges = [][2]int{}
					}
					verticalRanges[x] = append(ranges, [2]int{start, end})
				}

				continue YScan
			}

			if insidePolygon {
				if !nextStepUpStaysInsidePolygon(x, y, clockwiseTiles) {
					// fmt.Print("\tToggling to 'outside (next step not inside)'\n")

					insidePolygon = false
					ranges, exists := verticalRanges[x]
					if !exists {
						ranges = [][2]int{}
					}
					verticalRanges[x] = append(ranges, [2]int{start, end})
				}

				continue YScan
			}

			// fmt.Print("\tToggling to 'inside'\n")
			insidePolygon = true
			start = y
		}
	}

	return verticalRanges
}

func coordinateOnSomeSegment(x int, y int, segments map[int][][2]int) ([2]int, bool) {
	relevantSegments, exists := segments[x]
	if !exists {
		return [2]int{}, false
	}

	// i, ok := slices.BinarySearchFunc(relevantSegments, y, func(segment [2]int, target int) int {
	// 	if target < segment[0] {
	// 		return 1
	// 	}

	// 	if segment[1] < target {
	// 		return -1
	// 	}

	// 	return 0
	// })

	for _, s := range relevantSegments {
		if s[0] <= y && y <= s[1] {
			return s, true
		}
	}

	return [2]int{}, false
}

func nextStepUpStaysInsidePolygon(x int, y int, clockwiseTiles [][2]int) bool {
	target := 0
	for i, tile := range clockwiseTiles {
		if tile[0] == x && tile[1] == y {
			target = i
			break
		}
	}
	nTiles := len(clockwiseTiles)
	z, a, b := clockwiseTiles[(target+nTiles-1)%len(clockwiseTiles)], clockwiseTiles[target], clockwiseTiles[(target+1)%len(clockwiseTiles)]

	// fmt.Printf("\t\tz, a, b: %v, %v, %v\n", z, a, b)

	if a[0] == z[0] {
		// A > B
		// ^
		// Z
		if z[1] < a[1] && a[0] < b[0] {
			// fmt.Print("\t\t1\n")
			return false
		}

		// B < A
		//     ^
		//     Z
		if z[1] < a[1] && b[0] < a[0] {
			// fmt.Print("\t\t2\n")
			return true
		}

		// Z
		// v
		// A > B
		if a[1] < z[1] && a[0] < b[0] {
			// fmt.Print("\t\t3\n")
			return true
		}

		//     Z
		//     v
		// B < A
		if a[1] < z[1] && b[0] < a[0] {
			// fmt.Print("\t\t4\n")
			return true
		}
	} else {

		//     B
		//     ^
		// Z > A
		if z[0] < a[0] && a[1] < b[1] {
			// fmt.Print("\t\t5\n")
			return true
		}

		// Z > A
		//     v
		//     B
		if z[0] < a[0] && b[1] < a[1] {
			// fmt.Print("\t\t6\n")
			return false
		}

		// B
		// ^
		// A < Z
		if a[0] < z[0] && a[1] < b[1] {
			// fmt.Print("\t\t7\n")
			return true
		}

		// A < Z
		// v
		// B
		if a[0] < z[0] && b[1] < a[1] {
			// fmt.Print("\t\t8\n")
			return true
		}
	}

	return false
}

func checkRectangle(a [2]int, b [2]int, verticalSegments map[int][][2]int, xCoordinates []int) bool {
	x1, x2 := a[0], b[0]
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	y1, y2 := a[1], b[1]
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	firstXIndex, ok := slices.BinarySearch(xCoordinates, x1)
	if !ok {
		panic("uh oh")
	}

	satisfied := true
XScan:
	for i := firstXIndex; i < len(xCoordinates) && xCoordinates[i] <= x2; i += 1 {
		x := xCoordinates[i]

		segments, ok := verticalSegments[x]
		if !ok {
			panic("oops")
		}

		for _, s := range segments {
			if s[0] <= y1 && y2 <= s[1] {
				continue XScan
			}
		}

		satisfied = false
		break XScan
	}

	return satisfied
}
