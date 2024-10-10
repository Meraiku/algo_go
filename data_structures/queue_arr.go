package datastructures

import (
	"fmt"
	"sync"
)

type QueueArr[T any] struct {
	items []T
	mu    sync.Mutex
}

func NewQueueArr[T any]() *QueueArr[T] {
	return &QueueArr[T]{
		items: []T{},
		mu:    sync.Mutex{},
	}
}

func (q *QueueArr[T]) Push(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, item)
}

func (q *QueueArr[T]) Pop() (T, error) {
	l := len(q.items)
	var out T

	if l == 0 {
		return out, fmt.Errorf("queue pop: %w", ErrNoElements)
	}

	out = q.items[0]
	q.items = q.items[1:]

	return out, nil
}

func (q *QueueArr[T]) Peek() (T, error) {
	l := len(q.items)
	var out T

	if l == 0 {
		return out, fmt.Errorf("queue peek: %w", ErrNoElements)
	}

	out = q.items[0]

	return out, nil
}
