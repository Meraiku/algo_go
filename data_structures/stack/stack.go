package stack

import (
	"sync"

	datastructures "github.com/meraiku/algo_go/data_structures"
)

type Stack[T any] struct {
	items []T
	mu    sync.RWMutex
}

func (s *Stack[T]) Append(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.items == nil {
		s.items = []T{item}
		return
	}

	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var out T

	if len(s.items) == 0 {
		return out, datastructures.ErrNoElements
	}

	out = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return out, nil
}

func (s *Stack[T]) Top() (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var out T

	if len(s.items) == 0 {
		return out, datastructures.ErrNoElements
	}

	out = s.items[len(s.items)-1]

	return out, nil
}
