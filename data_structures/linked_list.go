package datastructures

import (
	"fmt"
	"sync"
)

type LinkedList[T any] struct {
	head   *Node[T]
	length int
	mu     sync.RWMutex
}

type Node[T any] struct {
	item T
	next *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		mu: sync.RWMutex{},
	}
}

func (l *LinkedList[T]) AddTail(item T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	node := &Node[T]{item: item}
	if l.head == nil {
		l.head = node
		return
	}
	temp := l.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = node

	l.length++
}

func (l *LinkedList[T]) AddHead(item T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	node := &Node[T]{item: item}
	if l.head == nil {
		l.head = node
		return
	}

	l.head, node.next = node, l.head
	l.length++
}

func (l *LinkedList[T]) PopTail() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var out T

	if l.head == nil {
		return out, fmt.Errorf("linked list pop tail: %w", ErrNoElements)
	}

	temp := l.head
	for temp.next.next != nil {
		temp = temp.next
	}
	out = temp.next.item
	temp.next = nil
	l.length--

	return out, nil
}

func (l *LinkedList[T]) PopHead() (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var out T

	if l.head == nil {
		return out, fmt.Errorf("linked list pop tail: %w", ErrNoElements)
	}

	out = l.head.item
	l.head = l.head.next
	l.length--

	return out, nil
}

func (l *LinkedList[T]) PeekHead() (T, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var out T
	if l.head == nil {
		return out, fmt.Errorf("linked list peek head: %w", ErrNoElements)
	}

	out = l.head.item

	return out, nil
}
