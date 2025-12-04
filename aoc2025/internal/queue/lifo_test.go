package queue

import (
	"testing"
)

func TestLifoRemovesItemsInReverseInsertionOrder(t *testing.T) {
	lifo := NewLifo[int]()

	for i := range 5 {
		lifo.Push(i)
	}

	for i := range 5 {
		x, ok := lifo.Pop()

		if !ok {
			t.Errorf("Expected pop of %d to be successful (ok)", i)
		}

		if x != 5-(i+1) {
			t.Errorf("Expected the %dth pop to produce %d; got %d", i, 5-(i+1), x)
		}
	}
}

func TestLifoEmptyPopReturnsFalse(t *testing.T) {
	lifo := NewLifo[int]()

	_, ok := lifo.Pop()

	if ok {
		t.Errorf("Expected pop from empty queue to produce zero value and false; got %v", ok)
	}
}
