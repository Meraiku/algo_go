package search

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

func Binary[T constraints.Ordered](arr []T, val T) (int, error) {
	low, high := 0, len(arr)-1

	loop := 0

	for low <= high {
		mid := (low + high) / 2

		switch {
		case arr[mid] == val:
			fmt.Printf("done in %d iterations\n", loop)
			return mid, nil
		case arr[mid] > val:
			high = mid - 1
		case arr[mid] < val:
			low = mid + 1
		}
		loop++
	}

	return -1, errors.New("value is not found")
}
