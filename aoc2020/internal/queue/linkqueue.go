package queue

import (
	"fmt"

	"github.com/pkg/errors"
)

// LinkQueue implements the Queue interface via a linked list representation.
type LinkQueue struct {
	head *node
	length int
	tail *node
}

type node struct {
	data interface{}
	next *node
}

// NewLinkQueue creates a new, empty LinkQueue.
func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		head: nil,
		length: 0,
		tail: nil,
	}
}

// AppendRight adds an item to the queue's tail.
func (lq *LinkQueue) AppendRight(x interface{}) {
	newNode := &node{
		data: x,
		next: nil,
	}

	if lq.head == nil {
		lq.head = newNode
	} else {
		lq.tail.next = newNode
	}
	lq.tail = newNode

	lq.length++
}
	
// Len provides the number of items in the queue.
func (lq *LinkQueue) Len() int {
	return lq.length
}

// PopLeft removes the item at the head of the queue and returns it.
// An error is returned if the queue is empty.
func (lq *LinkQueue) PopLeft() (interface{}, error) {
	if lq.head == nil {
		return nil, errors.New("cannot pop from empty queue")
	}

	ret := lq.head.data
	lq.head = lq.head.next

	lq.length--

	return ret, nil
}

// String stringifies the contents of the queue.
func (lq *LinkQueue) String() string {
	contents := make([]interface{}, 0, lq.Len())
	current := lq.head
	for current != nil {
		contents = append(contents, current.data)
		current = current.next
	}
	return fmt.Sprintf("%v", contents)
}
