package queue

import (
	"container/heap"
)

type Priority[T any] struct {
	internal internalHeap[T]
}

func NewPriority[T comparable](cmp func(T, T) bool) Priority[T] {
	return NewPriorityWithCapacity[T](cmp, 100)
}

func NewPriorityWithCapacity[T comparable](cmp func(T, T) bool, capacity int) Priority[T] {
	return Priority[T]{
		internal: internalHeap[T]{
			data: make([]T, 0, capacity),
			cmp:  cmp,
		},
	}
}

func (self *Priority[T]) Push(x T) {
	heap.Push(&self.internal, x)
}

func (self *Priority[T]) Pop() (T, bool) {
	var ret T
	if self.Len() == 0 {
		return ret, false
	}

	ret = heap.Pop(&self.internal).(T)
	return ret, true
}

func (self *Priority[T]) Len() int {
	return self.internal.Len()
}

type internalHeap[T any] struct {
	data []T
	cmp  func(T, T) bool
}

func (self internalHeap[T]) Len() int {
	return len(self.data)
}

func (self internalHeap[T]) Less(i int, j int) bool {
	return self.cmp(self.data[i], self.data[j])
}

func (self internalHeap[T]) Swap(i int, j int) {
	self.data[i], self.data[j] = self.data[j], self.data[i]
}

func (self *internalHeap[T]) Push(x any) {
	self.data = append(self.data, x.(T))
}

func (self *internalHeap[T]) Pop() any {
	old := self.data
	n := len(old)
	x := old[n-1]
	self.data = old[:n-1]
	return x
}
