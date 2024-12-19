package queue

type Fifo[T any] struct {
	data  []T
	first int
	last  int
}

func NewFifo[T any]() Fifo[T] {
	return NewFifoWithCapacity[T](100)
}

func NewFifoWithCapacity[T any](capacity int) Fifo[T] {
	return Fifo[T]{
		data:  make([]T, capacity),
		first: 0,
		last:  0,
	}
}

func (self *Fifo[T]) Push(x T) {
	if self.Len() == len(self.data)-1 {
		self.reallocate(2 * len(self.data))
	}

	self.data[self.last] = x
	self.last = (self.last + 1) % len(self.data)
}

func (self *Fifo[T]) Pop() (T, bool) {
	var ret T
	if self.Len() == 0 {
		return ret, false
	}

	ret = self.data[self.first]
	self.first = (self.first + 1) % len(self.data)
	return ret, true
}

func (self *Fifo[T]) Len() int {
	if self.first > self.last {
		return len(self.data) - (self.first - self.last)
	}

	return self.last - self.first
}

func (self *Fifo[T]) reallocate(newCapacity int) {
	oldLen := self.Len()
	newData := make([]T, newCapacity)
	for i := 0; i < oldLen; i += 1 {
		newData[i] = self.data[(self.first+i)%oldLen]
	}
	self.data = newData
	self.first = 0
	self.last = oldLen
}
