package search

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/meraiku/algo_go/array"
	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	n := 200

	arr := array.Get()
	arr[0] = n
	slices.Sort(arr)

	index, err := Binary(arr, n)

	require.NoError(t, err)
	require.True(t, index >= 0)
}

func TestBinarySearchError(t *testing.T) {
	n := rand.Int()

	arr := array.Get()
	slices.Sort(arr)

	index, err := Binary(arr, n)

	require.Error(t, err)
	require.True(t, index == -1)
}

func TestRecursiveBinarySearch(t *testing.T) {
	n := 200

	arr := array.Get()
	arr[0] = n
	slices.Sort(arr)

	index := BinaryRecursive(arr, n, 0, len(arr)-1)

	require.True(t, index >= 0)
}
