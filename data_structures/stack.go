package datastructures

import "sync"

type Stack[T any] struct {
	items []T
	mu    sync.RWMutex
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
		mu:    sync.RWMutex{},
	}
}

func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.items) <= 0 {
		var out T
		return out
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.items) <= 0 {
		var out T
		return out
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.items)
}
