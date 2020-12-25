package grid

import (
	"math"
)

// ManhattanDistance computes the Manattan, or taxi cab, distance
// between two points on a grid.
func ManhattanDistance(a Point, b Point) int {
	x1, y1 := float64(a.I), float64(a.J)
	x2, y2 := float64(b.I), float64(b.J)

	distance := math.Abs(x1 - x2) + math.Abs(y1 - y2)

	return int(distance)
}
