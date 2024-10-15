package search

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Binary[T constraints.Ordered](arr []T, val T) (int, error) {
	low, high := 0, len(arr)-1

	loop := 0

	for low <= high {
		mid := (low + high) / 2

		switch {
		case arr[mid] == val:
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

func BinaryRecursive[T constraints.Ordered](arr []T, val T, low, high int) int {
	mid := (low + high) / 2
	if arr[mid] == val {
		return mid
	}
	if low == high {
		return -1
	}
	if arr[mid] < val {
		return BinaryRecursive(arr, val, mid+1, high)
	} else {
		return BinaryRecursive(arr, val, low, mid-1)
	}
}
