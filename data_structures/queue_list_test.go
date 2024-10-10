package datastructures

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueueList(t *testing.T) {
	queue := NewQueueList[int]()

	queue.Push(5)

	for i := 0; i < 10; i++ {
		queue.Push(rand.Intn(1000))
	}

	val, err := queue.Pop()

	require.NoError(t, err)
	require.Equal(t, 5, val)
}
