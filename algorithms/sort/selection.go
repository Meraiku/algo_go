package sort

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](arr []T) {

	for i := range arr {
		smallestIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[smallestIndex] {
				smallestIndex = j
			}
		}
		arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
	}

}
