package datastructures

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

	require.Equal(t, 3, stack.Size())
	require.Equal(t, third, stack.Peek())
	require.Equal(t, third, stack.Pop())
	require.Equal(t, second, stack.Pop())
	require.Equal(t, first, stack.Pop())
	require.Equal(t, "", stack.Peek())
	require.Equal(t, "", stack.Pop())
	require.Equal(t, 0, stack.Size())
}
