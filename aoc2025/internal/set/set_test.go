package set

import (
	"testing"
)

func TestSetTracksItsOwnSize(t *testing.T) {
	s := New[int]()

	if l := s.Len(); l != 0 {
		t.Errorf("Expected set to have no length before adding items; got %d", l)
	}

	for i := range 3 {
		s.Add(i)

		if l := s.Len(); l != i+1 {
			t.Errorf("Expected set to have length %d; got %d", i+1, l)
		}
	}
}

func TestSetKeepsAddedItems(t *testing.T) {
	s := New[int]()

	for i := range 5 {
		if i&1 == 0 {
			if exists := s.Add(i); exists {
				t.Errorf("Expected set to report that %d was not already present; it did", i)
			}
		}
	}

	for i := range 5 {
		if i&1 == 0 {
			if !s.Contains(i) {
				t.Errorf("Expected set to contain %d; it did not", i)
			}
		} else {
			if s.Contains(i) {
				t.Errorf("Expected set not to contain %d; it did", i)
			}
		}
	}

	if exists := s.Add(2); !exists {
		t.Errorf("Expected set to report that 2 was already present in the set when adding; it did not")
	}
}

func TestSetDropsRemovedItems(t *testing.T) {
	s := New[int]()

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if removed := s.Remove(2); !removed {
		t.Errorf("Expected set to report that 2 was removed; it did not")
	}

	if s.Contains(2) {
		t.Errorf("Expected set to no longer contain 2; it did")
	}

	if removed := s.Remove(5); removed {
		t.Errorf("Expected set to report that it did not remove 5 (because it was not in the set); it did")
	}
}

func TestSetContainsReportsMembershipAccurately(t *testing.T) {
	s := New[int]()

	s.Add(5)

	if s.Contains(0) {
		t.Errorf("Expected set to not contain 0; it reports it does")
	}

	if !s.Contains(5) {
		t.Errorf("Expected set to contain 5; it reports it does not")
	}
}

func TestSetIsIterable(t *testing.T) {
	bitSet := uint(0)
	bitSet |= 1 << 3
	bitSet |= 1 << 5

	s := New[int]()
	s.Add(3)
	s.Add(5)

	for val := range s.All() {
		bitSet ^= 1 << val
	}

	if bitSet != 0 {
		t.Errorf("Iteration did not produce precisely the values added to the set")
	}
}

func TestUnionProducesTheUnionOfTwoSets(t *testing.T) {
	a := New[int]()
	b := New[int]()

	a.Add(1)
	a.Add(2)

	b.Add(2)
	b.Add(3)

	c := Union(a, b)

	for i := range 3 {
		if !c.Contains(i + 1) {
			t.Errorf("Expected union set to contain %d; it did not", i+1)
		}
	}

	if c.Contains(5) {
		t.Errorf("Expected union set not to contain 5; it did")
	}
}
