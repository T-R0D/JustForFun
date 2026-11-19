package set

import (
	"fmt"
	"iter"
)

type Disjoint[T comparable] struct {
	forest []node[T]
	lookup map[T]int
}

const (
	noRelationSentinel = -1
)

type node[T comparable] struct{
	parent int
	setSize int
	value *T
}

func NewDisjoint[T comparable]() Disjoint[T] {
	return Disjoint[T]{
		forest: []node[T]{},
		lookup: map[T]int{},
	}
}

func (self *Disjoint[T]) Add(x T) bool {
	if _, exists := self.lookup[x]; exists {
		return exists
	}

	self.lookup[x] = len(self.forest)
	self.forest = append(self.forest, node[T]{
		parent: noRelationSentinel,
		setSize: 1,
		value: &x,
	})

	return false
}

func (self *Disjoint[T]) Union(x T, y T) (bool, error) {
	if x == y {
		return false, nil
	}

	i, ok := self.lookup[x]
	if !ok {
		return false, fmt.Errorf("x is not present in the set")
	}

	j, ok := self.lookup[y]
	if !ok {
		return false, fmt.Errorf("y is not present in the set")
	}

	iNode := self.findRootNode(i)
	jNode := self.findRootNode(j)

	if iNode.setSize < jNode.setSize {
		iNode, jNode = jNode, iNode
	}

	jNode.parent = i
	iNode.setSize += jNode.setSize

	return true, nil
}

func (self *Disjoint[T]) All() iter.Seq[Set[T]] {
	return func(yield func(Set[T]) bool) {
		// for every parent, form a set
		s := New[T]()
		if !yield(s) {
			return
		}
	}
}

func (self *Disjoint[T]) findRootNode(i int) *node[T] {
	root := &self.forest[i]
	for root.parent != noRelationSentinel {
		root = &self.forest[root.parent]
	}

	return root
}
