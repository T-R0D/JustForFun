package location

import (
	"math"
)

func ManhattanDistance(a, b Point) int {
	return int(math.Abs(float64(a.X)-float64(b.X)) + math.Abs(float64(a.Y)-float64(b.Y)))
}
