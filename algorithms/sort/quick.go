package sort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	} else {
		index := rand.Intn(len(arr) - 1)
		pivot := arr[index]

		var left, right []T

		for i, num := range arr {
			if i == index {
				continue
			}

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
