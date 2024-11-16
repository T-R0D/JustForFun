package queue

type Lifo[T any] struct {
	data []T
}

func NewLifo[T any]() Lifo[T] {
	return NewLifoWithCapacity[T](100)
}

func NewLifoWithCapacity[T any](capacity int) Lifo[T] {
	return Lifo[T]{
		data: make([]T, 0, capacity),
	}
}

func (self *Lifo[T]) Push(x T) {
	self.data = append(self.data, x)
}

func (self *Lifo[T]) Pop() (T, bool) {
	var ret T
	if self.Len() == 0 {
		return ret, false
	}

	ret = self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]
	return ret, true
}

func (self *Lifo[T]) Len() int {
	return len(self.data)
}
