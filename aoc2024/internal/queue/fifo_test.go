package queue

import (
	"testing"
)

func TestFifoRemovesItemsInInsertionOrder(t *testing.T) {
	fifo := NewFifo[int]()

	for i := range 5 {
		fifo.Push(i)
	}

	for i := range 5 {
		x, ok := fifo.Pop()

		if !ok {
			t.Errorf("Expected pop of %d to be successful (ok)", i)
		}

		if x != i {
			t.Errorf("Expected the %dth pop to produce %d; got %d", i, i, x)
		}
	}
}

func TestFifoWithReallocationRemovesItemsInInsertionOrder(t *testing.T) {
	fifo := NewFifoWithCapacity[int](3)

	for i := range 5 {
		fifo.Push(i)
	}

	for i := range 5 {
		x, ok := fifo.Pop()

		if !ok {
			t.Errorf("Expected pop of %d to be successful (ok)", i)
		}

		if x != i {
			t.Errorf("Expected the %dth pop to produce %d; got %d", i, i, x)
		}
	}
}

func TestFifoEmptyPopReturnsFalse(t *testing.T) {
	fifo := NewFifo[int]()

	_, ok := fifo.Pop()

	if ok {
		t.Errorf("Expected pop from empty queue to produce zero value and false; got %v", ok)
	}
}
