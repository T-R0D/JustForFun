package counter

import (
	"testing"
)

func TestCounterTracksNumberOfTrackedItems(t *testing.T) {
	c := New[int]()

	if l := c.Len(); l != 0 {
		t.Errorf("Expected counter to have Len of 0 before adding items; got %d", l)
	}

	c.Set(1, 3)
	c.Set(3, 5)

	if l := c.Len(); l != 2 {
		t.Errorf("Expected counter to have Len of 2 after adding items; got %d", l)
	}
}

func TestCounterTracksSetValues(t *testing.T) {
	c := New[int]()

	for i := range 5 {
		c.Set(i, i*2)
	}

	for i := range 5 {
		if count, exists := c.Get(i); !exists {
			t.Errorf("Expected %d to have a count set; there was none", i)
		} else if count != i*2 {
			t.Errorf("Expected value %d to have a count of %d; got %d instead", i, i*2, count)
		}
	}
}

func TestCounterTracksIncrementedValues(t *testing.T) {
	c := New[int]()

	if exists := c.Increment(3, 2); exists {
		t.Errorf("Expected counter to report that count for 3 was not tracked yet; it reported that 3 was already present")
	}

	if exists := c.Increment(3, 3); !exists {
		t.Errorf("Expected counter to report that count for 3 was already tracked; it reported that 3 was not already present")
	}

	if count, exists := c.Get(3); !exists {
		t.Errorf("Expected 3 to have a count set; there was none")
	} else if count != 5 {
		t.Errorf("Expected value 3 to have a count of 5; got %d instead", count)
	}
}

func TestCounterSumsTrackedValues(t *testing.T) {
	c := New[int]()

	for i := range 6 {
		c.Increment(i, i)
	}

	if sum := c.Sum(); sum != 15 {
		t.Errorf("Expected counter to have a sum of 15; got %d", sum)
	}
}

func TestCounterIteratesTrackedCounts(t *testing.T) {
	c := New[int]()

	expectedCounts := []int{0, 2, 1, 0}
	c.Set(0, 1)
	c.Set(1, 2)
	c.Set(2, 1)

	actualCounts := append([]int{}, expectedCounts...)
	for countValue := range c.Counts() {
		actualCounts[countValue] -= 1
	}

	for i, countedTimes := range actualCounts {
		if countedTimes < 0 {
			t.Errorf("Expected %d to be counted %d times; it was counted %d too many times", i, expectedCounts[i], -countedTimes)
		}

		if 0 < countedTimes {
			t.Errorf("Expected %d to be counted %d times; it was counted %d too few times", i, expectedCounts[i], countedTimes)
		}
	}
}

func TestCounterIteratesItems(t *testing.T) {
	c := New[int]()

	expectedValues := map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 3,
		4: 5,
		5: 8,
	}

	for k, v := range expectedValues {
		c.Increment(k, v)
	}

	for k, v := range c.Items() {
		if expected, exists := expectedValues[k]; !exists {
			t.Errorf("Expected counter to not have key %d; it did", k)
		} else if v != expected {
			t.Errorf("Expected counter to have value of %d for key %d; got %d", expected, k, v)
		}

		delete(expectedValues, k)
	}

	if len(expectedValues) > 0 {
		t.Errorf("Expected all items inserted into counter to be iterated; the following were missed: %v", expectedValues)
	}
}
