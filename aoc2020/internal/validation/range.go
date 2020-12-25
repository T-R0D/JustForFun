package validation

// ValueIsInRange tests to see if a value is in the given range,
// inclusive of the lower bound, exclusive of the upper bound.
func ValueIsInRange(candidate int, lowerBound int, upperBound int) bool {
	return lowerBound <= candidate && candidate < upperBound
}
