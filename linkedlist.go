package datastructures

import (
	"errors"
)

type LinkedList[T interface{}] struct {
	first *node[T]
	last  *node[T]
	len   uint32
}

type node[T interface{}] struct {
	Value T
	next  *node[T]
}

func (queue *LinkedList[T]) push(node *node[T]) {
	defer func() { queue.len++ }()
	if queue.first == nil {
		queue.first = node
		queue.first.next = nil
		queue.last = node
		return
	}
	last := queue.last
	last.next = node
	queue.last = node
}

func (queue *LinkedList[T]) Push(value T) {
	queue.push(&node[T]{Value: value})
}

func (queue *LinkedList[T]) Pop() (T, error) {
	if queue.len == 0 || queue.first == nil {
		var empty T
		return empty, errors.New("empty queue")
	}
	retVal := queue.first
	queue.first = retVal.next
	queue.len--
	if queue.len == 0 {
		queue.last = nil
	}
	return retVal.Value, nil
}

func (queue *LinkedList[T]) Len() uint32 {
	return queue.len
}

func (queue *LinkedList[T]) Clear() {
	len := queue.Len()
	for i := 0; i < int(len); i++ {
		queue.Pop()
	}
}
