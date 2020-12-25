package equal

// IntArrayEqual determines if two int arrays have the same contents
func IntArrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		} 
	}

	return true
}
