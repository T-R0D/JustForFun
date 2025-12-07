package counter

import (
	"iter"
)

type Counter[T comparable] struct {
	data map[T]int
}

func New[T comparable]() Counter[T] {
	return Counter[T]{
		data: map[T]int{},
	}
}

func (self *Counter[T]) Len() int {
	return len(self.data)
}

func (self *Counter[T]) Set(k T, v int) {
	self.data[k] = v
}

func (self *Counter[T]) Increment(k T, v int) bool {
	if prev, exists := self.data[k]; exists {
		self.data[k] = prev + v
		return exists
	}

	self.data[k] = v
	return false
}

func (self *Counter[T]) Get(k T) (int, bool) {
	if val, exists := self.data[k]; exists {
		return val, exists
	}

	return 0, false
}

func (self *Counter[T]) Sum() int {
	sum := 0
	for count := range self.Counts() {
		sum += count
	}

	return sum
}

func (self *Counter[T]) Counts() iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, val := range self.data {
			if !yield(val) {
				return
			}
		}
	}
}

func (self *Counter[T]) Items() iter.Seq2[T, int] {
	return func(yield func(T, int) bool) {
		for k, val := range self.data {
			if !yield(k, val) {
				return
			}
		}
	}
}
