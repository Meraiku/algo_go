package sort

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/meraiku/algo_go/array"
	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {

	arr := array.Get()

	exp := make([]int, len(arr))
	copy(exp, arr)

	slices.Sort(exp)

	start := time.Now()
	out := QuickSort(arr)
	fmt.Printf("Quick Sort O(n log n) done in: %.2f Seconds\n", time.Since(start).Seconds())

	require.Equal(t, exp, out)
}
