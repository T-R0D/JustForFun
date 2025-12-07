package set

import (
	"iter"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{
		data: map[T]struct{}{},
	}
}

func (self *Set[T]) Len() int {
	return len(self.data)
}

func (self *Set[T]) Add(x T) bool {
	_, exists := self.data[x]

	self.data[x] = struct{}{}

	return exists
}

func (self *Set[T]) Remove(x T) bool {
	if _, exists := self.data[x]; exists {
		delete(self.data, x)
		return exists
	}

	return false
}

func (self *Set[T]) Contains(x T) bool {
	_, exists := self.data[x]
	return exists
}

func (self *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range self.data {
			if !yield(k) {
				return
			}
		}
	}
}
