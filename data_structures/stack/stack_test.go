package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	stack := New[string]()

	first, second, third := "mango", "apple", "water"

	stack.Push(first)
	stack.Push(second)
	stack.Push(third)

	var elem string

	require.Equal(t, 3, stack.Size())
	elem, _ = stack.Peek()
	require.Equal(t, third, elem)
	elem, _ = stack.Pop()
	require.Equal(t, third, elem)

	stack.Pop()
	stack.Pop()

	require.Equal(t, 0, stack.Size())
}
