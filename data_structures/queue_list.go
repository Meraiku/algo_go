package datastructures

import (
	"fmt"
)

type QueueList[T any] struct {
	items *LinkedList[T]
}

func NewQueueList[T any]() *QueueList[T] {
	return &QueueList[T]{
		items: NewLinkedList[T](),
	}
}

func (q *QueueList[T]) Push(item T) {
	q.items.AddTail(item)
}

func (q *QueueList[T]) Pop() (T, error) {

	out, err := q.items.PopHead()
	if err != nil {
		return out, fmt.Errorf("queue pop: %w", ErrNoElements)
	}

	return out, nil
}

func (q *QueueList[T]) Peek() (T, error) {

	out, err := q.items.PeekHead()
	if err != nil {
		return out, fmt.Errorf("queue peek: %w", ErrNoElements)
	}

	return out, nil
}
