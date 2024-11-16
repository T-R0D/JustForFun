package queue

import (
	"slices"
	"testing"
)

func TestPriorityPopsItemsInSortedOrder(t *testing.T) {
	type Item struct {
		priority int
		value    string
	}

	items := []Item{
		{
			priority: 2,
			value:    "two",
		},
		{
			priority: 3,
			value:    "three",
		},
		{
			priority: 1,
			value:    "one",
		},
	}

	priority := NewPriority(func(a Item, b Item) bool { return a.priority < b.priority })

	for _, item := range items {
		priority.Push(item)
	}

	slices.SortFunc(items, func(a Item, b Item) int {
		return a.priority - b.priority
	})

	for i, expectedItem := range items {
		item, ok := priority.Pop()

		if !ok {
			t.Errorf("Expected popping item %d to be ok; got %v", i, ok)
		}

		if item.priority != expectedItem.priority || item.value != expectedItem.value {
			t.Errorf("Expected popped item %d would be %v; got %v", i, expectedItem, item)
		}
	}
}

func TestEmptyPriorityPopReturnsZeroValueAndFalse(t *testing.T) {
	priority := NewPriority(func(a int, b int) bool { return a < b })

	_, ok := priority.Pop()

	if ok {
		t.Errorf("Expected pop from empty queue to return zero value and false; got %v", ok)
	}
}
