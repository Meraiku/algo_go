package sort

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](arr []T) {

	switching := true

	for switching {
		switching = false
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				switching = true
			}
		}
	}
}
