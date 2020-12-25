package location

import (
	"aoc2019/internal/aocmath"
	"math"
)

type Vector struct {
	X, Y int
}

func NewVector(x, y int) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

func (v *Vector) Reduce() {
	if v.X == 0 && v.Y == 0 {
		return
	} else if v.X == 0 {
		v.Y = v.Y / int(math.Abs(float64(v.Y)))
		return
	} else if v.Y == 0 {
		v.X = v.X / int(math.Abs(float64(v.X)))
		return
	}

	v.X, v.Y = aocmath.ReduceFraction(v.X, v.Y)
}
