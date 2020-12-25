package aocmath

import (
	"math"
)

func ReduceFraction(n, d int) (reducedNumerator int, reducedDenominator int) {
	gcd := GreatestCommonDivisor(n, d)
	return n / gcd, d / gcd
}

func GreatestCommonDivisor(a, b int) int {
	absA := int(math.Abs(float64(a)))
	absB := int(math.Abs(float64(b)))
	lesser := int(math.Min(float64(absA), float64(absB)))
	for i := lesser; i > 0; i-- {
		if absA % i == 0 && absB % i == 0 {
			return i
		} 
	}
	return 1
}