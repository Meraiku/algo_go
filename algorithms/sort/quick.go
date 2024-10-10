package sort

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]

		var left, right []T

		for _, num := range arr[1:] {

			if num <= pivot {
				left = append(left, num)
			} else {
				right = append(right, num)
			}
		}

		out := append(QuickSort(left), pivot)
		out = append(out, QuickSort(right)...)
		return out
	}
}
