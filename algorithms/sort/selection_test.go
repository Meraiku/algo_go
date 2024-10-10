package sort

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/meraiku/algo_go/array"
	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	arr := array.Get()

	exp := make([]int, len(arr))
	copy(exp, arr)

	slices.Sort(exp)

	start := time.Now()
	SelectionSort(arr)
	fmt.Printf("Selection sort O(n2) done in: %.2f Seconds\n", time.Since(start).Seconds())

	require.Equal(t, exp, arr)
}
